/*
Copyright 2019 The Kubernetes Authors All rights reserved.

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

package ssh

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/pkg/errors"
	"k8s.io/klog/v2"
	pkgdrivers "k8s.io/minikube/pkg/drivers"
	"k8s.io/minikube/pkg/libmachine/libmachine/cruntime"
	"k8s.io/minikube/pkg/libmachine/libmachine/drivers"
	"k8s.io/minikube/pkg/libmachine/libmachine/engine"
	"k8s.io/minikube/pkg/libmachine/libmachine/log"
	"k8s.io/minikube/pkg/libmachine/libmachine/mcnutils"
	"k8s.io/minikube/pkg/libmachine/libmachine/runner"
	"k8s.io/minikube/pkg/libmachine/libmachine/state"
	"k8s.io/minikube/pkg/minikube/sysinit"
)

// Driver is a driver designed to run kubeadm w/o VM management.
// https://minikube.sigs.k8s.io/docs/reference/drivers/ssh/
type Driver struct {
	*drivers.BaseDriver
	*pkgdrivers.CommonDriver
	EnginePort int
	SSHKey     string
	runtime    cruntime.CRuntime
	runner     runner.Runner
}

// Config is configuration for the SSH driver
type Config struct {
	MachineName      string
	StorePath        string
	ContainerRuntime string
}

const (
	defaultTimeout = 15 * time.Second
)

// NewDriver creates and returns a new instance of the driver
func NewDriver(c Config) *Driver {
	d := &Driver{
		EnginePort: engine.DefaultPort,
		BaseDriver: &drivers.BaseDriver{
			MachineName: c.MachineName,
			StorePath:   c.StorePath,
		},
	}
	runner, err := d.GetRunner()
	if err != nil {
		klog.Fatalf("unable to create container runtime: %v", err)
	}
	runtime, err := cruntime.New(cruntime.Config{Type: c.ContainerRuntime, Runner: runner})
	// Libraries shouldn't panic, but there is no way for drivers to return error :(
	if err != nil {
		klog.Fatalf("unable to create container runtime: %v", err)
	}
	d.runtime = runtime
	d.runner = runner
	return d
}

// DriverName returns the name of the driver
func (d *Driver) DriverName() string {
	return "ssh"
}

// GetSSHHostname returns hostname for use with ssh
func (d *Driver) GetSSHHostname() (string, error) {
	return d.GetIP()
}

// GetSSHUsername returns username for use with ssh
func (d *Driver) GetSSHUsername() string {
	return d.SSHUser
}

// GetSSHKeyPath returns the key path for SSH
func (d *Driver) GetSSHKeyPath() string {
	return d.SSHKeyPath
}

// PreCreateCheck checks for correct privileges and dependencies
func (d *Driver) PreCreateCheck() error {
	if d.SSHKey != "" {
		if _, err := os.Stat(d.SSHKey); os.IsNotExist(err) {
			return fmt.Errorf("SSH key does not exist: %q", d.SSHKey)
		}

		key, err := os.ReadFile(d.SSHKey)
		if err != nil {
			return err
		}

		_, err = ssh.ParsePrivateKey(key)
		if err != nil {
			return errors.Wrapf(err, "SSH key does not parse: %q", d.SSHKey)
		}
	}

	return nil
}

// Create a host using the driver's config
func (d *Driver) CreateMachine() error {
	if d.SSHKey == "" {
		log.Info("No SSH key specified. Assuming an existing key at the default location.")
	} else {
		log.Info("Importing SSH key...")

		d.SSHKeyPath = d.ResolveStorePath(path.Base(d.SSHKey))
		if err := copySSHKey(d.SSHKey, d.SSHKeyPath); err != nil {
			return err
		}

		if err := copySSHKey(d.SSHKey+".pub", d.SSHKeyPath+".pub"); err != nil {
			log.Infof("Couldn't copy SSH public key : %s", err)
		}
	}

	if d.runtime.Name() == "Docker" {
		groups, err := d.runner.RunCmd(exec.Command("groups", d.GetSSHUsername()))
		if err != nil {
			return errors.Wrap(err, "groups")
		}
		if !strings.Contains(groups.Stdout.String(), "docker") {
			if _, err := d.runner.RunCmd(exec.Command("sudo", "usermod", "-aG", "docker", d.GetSSHUsername())); err != nil {
				return errors.Wrap(err, "usermod")
			}
		}
	}

	log.Debugf("IP: %s", d.IPAddress)

	return nil
}

// GetURL returns a Docker URL inside this host
func (d *Driver) GetURL() (string, error) {
	if err := drivers.MustBeRunning(d); err != nil {
		return "", err
	}

	ip, err := d.GetIP()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("tcp://%s", net.JoinHostPort(ip, strconv.Itoa(d.EnginePort))), nil
}

// GetMachineState returns the state that the host is in (running, stopped, etc)
func (d *Driver) GetMachineState() (state.State, error) {
	address := net.JoinHostPort(d.IPAddress, strconv.Itoa(d.SSHPort))

	_, err := net.DialTimeout("tcp", address, defaultTimeout)
	if err != nil {
		return state.Stopped, nil
	}

	return state.Running, nil
}

// StartMachine a host
func (d *Driver) StartMachine() error {
	return nil
}

// StopMachine a host gracefully, including any containers that we are managing.
func (d *Driver) StopMachine() error {
	if err := sysinit.New(d.runner).Stop("kubelet"); err != nil {
		klog.Warningf("couldn't stop kubelet. will continue with stop anyways: %v", err)
		if err := sysinit.New(d.runner).ForceStop("kubelet"); err != nil {
			klog.Warningf("couldn't force stop kubelet. will continue with stop anyways: %v", err)
		}
	}
	containers, err := d.runtime.ListContainers(cruntime.ListContainersOptions{})
	if err != nil {
		return errors.Wrap(err, "containers")
	}
	if len(containers) > 0 {
		if err := d.runtime.StopContainers(containers); err != nil {
			return errors.Wrap(err, "stop containers")
		}
	}
	klog.Infof("ssh driver is stopped!")
	return nil
}

// RestartMachine a host
func (d *Driver) RestartMachine() error {
	return sysinit.New(d.runner).Restart("kubelet")
}

// KillMachine stops a host forcefully, including any containers that we are managing.
func (d *Driver) KillMachine() error {
	if err := sysinit.New(d.runner).ForceStop("kubelet"); err != nil {
		klog.Warningf("couldn't force stop kubelet. will continue with kill anyways: %v", err)
	}

	// First try to gracefully stop containers
	containers, err := d.runtime.ListContainers(cruntime.ListContainersOptions{})
	if err != nil {
		return errors.Wrap(err, "containers")
	}
	if len(containers) == 0 {
		return nil
	}
	// Try to be graceful before sending SIGKILL everywhere.
	if err := d.runtime.StopContainers(containers); err != nil {
		return errors.Wrap(err, "stop")
	}

	containers, err = d.runtime.ListContainers(cruntime.ListContainersOptions{})
	if err != nil {
		return errors.Wrap(err, "containers")
	}
	if len(containers) == 0 {
		return nil
	}
	if err := d.runtime.KillContainers(containers); err != nil {
		return errors.Wrap(err, "kill")
	}
	return nil
}

// RemoveMachine a host, including any data which may have been written by it.
func (d *Driver) RemoveMachine() error {
	return nil
}

func copySSHKey(src, dst string) error {
	if err := mcnutils.CopyFile(src, dst); err != nil {
		return fmt.Errorf("unable to copy ssh key: %s", err)
	}

	if err := os.Chmod(dst, 0600); err != nil {
		return fmt.Errorf("unable to set permissions on the ssh key: %s", err)
	}

	return nil
}

// x7TODO:
// implement those
func (d *Driver) RunCmd(cmd *exec.Cmd) (*runner.RunResult, error) {
	return nil, nil
}

func (d *Driver) GetRunner() (runner.Runner, error) {
	return nil, nil
}
