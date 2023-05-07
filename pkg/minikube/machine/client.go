/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package machine

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/juju/fslock"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"
	"k8s.io/minikube/pkg/libmachine/libmachine"
	"k8s.io/minikube/pkg/libmachine/libmachine/auth"
	"k8s.io/minikube/pkg/libmachine/libmachine/cert"
	"k8s.io/minikube/pkg/libmachine/libmachine/check"
	"k8s.io/minikube/pkg/libmachine/libmachine/drivers"
	"k8s.io/minikube/pkg/libmachine/libmachine/drivers/plugin"
	"k8s.io/minikube/pkg/libmachine/libmachine/drivers/plugin/localbinary"
	"k8s.io/minikube/pkg/libmachine/libmachine/engine"
	"k8s.io/minikube/pkg/libmachine/libmachine/host"
	"k8s.io/minikube/pkg/libmachine/libmachine/mcnutils"
	"k8s.io/minikube/pkg/libmachine/libmachine/persist"
	"k8s.io/minikube/pkg/libmachine/libmachine/runner"
	lmssh "k8s.io/minikube/pkg/libmachine/libmachine/ssh"
	"k8s.io/minikube/pkg/libmachine/libmachine/state"
	"k8s.io/minikube/pkg/libmachine/libmachine/swarm"
	"k8s.io/minikube/pkg/libmachine/libmachine/version"
	"k8s.io/minikube/pkg/minikube/driver"
	"k8s.io/minikube/pkg/minikube/exit"
	"k8s.io/minikube/pkg/minikube/localpath"
	"k8s.io/minikube/pkg/minikube/out"
	"k8s.io/minikube/pkg/minikube/reason"
	"k8s.io/minikube/pkg/minikube/registry"
)

// NewRPCClient gets a new client.
func NewRPCClient(storePath, certsDir string) libmachine.API {
	c := libmachine.NewClient(storePath, certsDir)
	c.SSHClientType = lmssh.Native
	return c
}

// NewAPIClient gets a new client.
func NewAPIClient(miniHome ...string) (libmachine.API, error) {
	storePath := localpath.MiniPath()
	if len(miniHome) > 0 {
		storePath = miniHome[0]
	}
	certsDir := localpath.MakeMiniPath("certs")

	return &LocalClient{
		certsDir:     certsDir,
		storePath:    storePath,
		Filestore:    persist.NewFilestore(storePath, certsDir, certsDir),
		legacyClient: NewRPCClient(storePath, certsDir),
		flock:        fslock.New(localpath.MakeMiniPath("machine_client.lock")),
	}, nil
}

// LocalClient is a non-RPC implementation
// of the libmachine API
type LocalClient struct {
	certsDir  string
	storePath string
	*persist.Filestore
	legacyClient libmachine.API
	flock        *fslock.Lock
}

// NewHost creates a new Host
func (api *LocalClient) NewHost(drvName, cruntimeName string, rawDriver []byte) (*host.Host, error) {
	def := registry.Driver(drvName)
	if def.Empty() {
		return nil, fmt.Errorf("driver %q does not exist", drvName)
	}
	if def.Init == nil {
		return api.legacyClient.NewHost(drvName, cruntimeName, rawDriver)
	}
	d := def.Init()
	err := json.Unmarshal(rawDriver, d)
	if err != nil {
		return nil, errors.Wrapf(err, "Error getting driver %s", string(rawDriver))
	}

	return &host.Host{
		ConfigVersion: version.ConfigVersion,
		Name:          d.GetMachineName(),
		Driver:        d,
		DriverName:    d.DriverName(),
		HostOptions: &host.Options{
			AuthOptions: &auth.Options{
				CertDir:          api.certsDir,
				CaCertPath:       filepath.Join(api.certsDir, "ca.pem"),
				CaPrivateKeyPath: filepath.Join(api.certsDir, "ca-key.pem"),
				ClientCertPath:   filepath.Join(api.certsDir, "cert.pem"),
				ClientKeyPath:    filepath.Join(api.certsDir, "key.pem"),
				ServerCertPath:   filepath.Join(api.GetMachinesDir(), "server.pem"),
				ServerKeyPath:    filepath.Join(api.GetMachinesDir(), "server-key.pem"),
			},
			EngineOptions: &engine.Options{
				StorageDriver: "overlay2",
				TLSVerify:     true,
			},
			SwarmOptions: &swarm.Options{},
		},
	}, nil
}

// Load a new client, creating driver
func (api *LocalClient) Load(name string) (*host.Host, error) {
	h, err := api.Filestore.Load(name)
	if err != nil {
		return nil, errors.Wrapf(err, "filestore %q", name)
	}

	def := registry.Driver(h.DriverName)
	if def.Empty() {
		return nil, fmt.Errorf("driver %q does not exist", h.DriverName)
	}
	if def.Init == nil {
		return api.legacyClient.Load(name)
	}
	h.Driver = def.Init()
	return h, json.Unmarshal(h.RawDriver, h.Driver)
}

// Close closes the client
func (api *LocalClient) Close() error {
	if api.legacyClient != nil {
		return api.legacyClient.Close()
	}
	return nil
}

// CommandRunner returns best available command runner for this host
func CommandRunner(h *host.Host) (runner.Runner, error) {
	// x7NOTE: we're leaving this responsibility to the driver itself
	return h.Driver.GetRunner()
}

// Create creates the host
func (api *LocalClient) Create(h *host.Host) error {
	klog.Infof("LocalClient.Create starting")
	start := time.Now()
	defer func() {
		klog.Infof("LocalClient.Create took %s", time.Since(start))
	}()

	def := registry.Driver(h.DriverName)
	if def.Empty() {
		return fmt.Errorf("driver %q does not exist", h.DriverName)
	}
	if def.Init == nil {
		// NOTE: This will call provision.DetectProvisioner
		return api.legacyClient.Create(h)
	}

	steps := []struct {
		name string
		f    func() error
	}{
		{
			"bootstrapping certificates",
			func() error {
				// Lock is needed to avoid race condition in parallel Docker-Env test because issue #10107.
				// CA cert and client cert should be generated atomically, otherwise might cause bad certificate error.
				lockErr := api.flock.LockWithTimeout(time.Second * 5)
				if lockErr != nil {
					return fmt.Errorf("failed to acquire bootstrap client lock: %v " + lockErr.Error())
				}
				defer func() {
					lockErr = api.flock.Unlock()
					if lockErr != nil {
						klog.Errorf("failed to release bootstrap cert client lock: %v", lockErr.Error())
					}
				}()
				certErr := cert.BootstrapCertificates(h.AuthOptions())
				return certErr
			},
		},
		{
			"precreate",
			h.Driver.PreCreateCheck,
		},
		{
			"saving",
			func() error {
				return api.Save(h)
			},
		},
		{
			"creating",
			h.Driver.CreateMachine,
		},
		{
			"waiting",
			func() error {
				if driver.BareMetal(h.Driver.DriverName()) {
					return nil
				}
				return mcnutils.WaitFor(drivers.MachineInState(h.Driver, state.Running))
			},
		},
		{
			"provisioning",
			func() error {
				// we're calling the provisioning method no matter what.
				// it is the provisioner's responsibility to check for
				// all the driver/machine special cases:
				// e.g. iso/container/unmanaged_machine
				return provisionMachine(h)
			},
		},
	}

	for _, step := range steps {
		if err := step.f(); err != nil {
			return errors.Wrap(err, step.name)
		}
	}

	return nil
}

// StartDriver starts the driver
func StartDriver() {
	cert.SetCertGenerator(&CertGenerator{})
	check.DefaultConnChecker = &ConnChecker{}
	if os.Getenv(localbinary.PluginEnvKey) == localbinary.PluginEnvVal {
		registerDriver(os.Getenv(localbinary.PluginEnvDriverName))
	}

	localbinary.CurrentBinaryIsDockerMachine = true
}

// ConnChecker can check the connection
type ConnChecker struct {
}

// Check checks the connection
func (cc *ConnChecker) Check(h *host.Host, _ bool) (string, *auth.Options, error) {
	authOptions := h.AuthOptions()
	dockerHost, err := h.Driver.GetURL()
	if err != nil {
		return "", &auth.Options{}, err
	}
	return dockerHost, authOptions, nil
}

// CertGenerator is used to override the default machine CertGenerator with a longer timeout.
type CertGenerator struct {
	cert.X509CertGenerator
}

// ValidateCertificate is a reimplementation of the default generator with a longer timeout.
func (cg *CertGenerator) ValidateCertificate(addr string, authOptions *auth.Options) (bool, error) {
	tlsConfig, err := cg.ReadTLSConfig(addr, authOptions)
	if err != nil {
		return false, err
	}

	dialer := &net.Dialer{
		Timeout: time.Second * 40,
	}

	_, err = tls.DialWithDialer(dialer, "tcp", addr, tlsConfig)
	if err != nil {
		return false, err
	}

	return true, nil
}

func registerDriver(drvName string) {
	def := registry.Driver(drvName)
	if def.Empty() {
		exit.Message(reason.Usage, "unsupported or missing driver: {{.name}}", out.V{"name": drvName})
	}
	plugin.RegisterDriver(def.Init())
}
