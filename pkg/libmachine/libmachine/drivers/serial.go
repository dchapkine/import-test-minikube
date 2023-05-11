/*
Copyright 2023 The Kubernetes Authors All rights reserved.

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

package drivers

import (
	"sync"

	"encoding/json"

	"k8s.io/minikube/pkg/libmachine/libmachine/mcnflag"
	"k8s.io/minikube/pkg/libmachine/libmachine/state"
)

var stdLock = &sync.Mutex{}

// SerialDriver is a wrapper struct which is used to ensure that RPC calls
// to a driver only occur one at a time.
// Some providers, e.g. virtualbox, should not run driver operations at the
// same time as other driver instances of the same type. Otherwise, we scrape
// up against VirtualBox's own locking mechanisms.
//
// It would be preferable to simply have a lock around, say, the VBoxManage
// command, but with our current one-server-process-per-machine model it is
// impossible to dictate this locking on the server side.
type SerialDriver struct {
	Driver
	sync.Locker
}

func NewSerialDriver(innerDriver Driver) Driver {
	return newSerialDriverWithLock(innerDriver, stdLock)
}

func newSerialDriverWithLock(innerDriver Driver, lock sync.Locker) Driver {
	return &SerialDriver{
		Driver: innerDriver,
		Locker: lock,
	}
}

// CreateMachine a host using the driver's config
func (d *SerialDriver) CreateMachine() error {
	d.Lock()
	defer d.Unlock()
	return d.Driver.CreateMachine()
}

// DriverName returns the name of the driver as it is registered
func (d *SerialDriver) DriverName() string {
	d.Lock()
	defer d.Unlock()
	return d.Driver.DriverName()
}

// GetCreateFlags returns the mcnflag.Flag slice representing the flags
// that can be set, their descriptions and defaults.
func (d *SerialDriver) GetCreateFlags() []mcnflag.Flag {
	d.Lock()
	defer d.Unlock()
	return d.Driver.GetCreateFlags()
}

// GetIP returns an IP or hostname that this host is available at
// e.g. 1.2.3.4 or docker-host-d60b70a14d3a.cloudapp.net
func (d *SerialDriver) GetIP() (string, error) {
	d.Lock()
	defer d.Unlock()
	return d.Driver.GetIP()
}

// GetMachineName returns the name of the machine
func (d *SerialDriver) GetMachineName() string {
	d.Lock()
	defer d.Unlock()
	return d.Driver.GetMachineName()
}

// GetSSHHostname returns hostname for use with ssh
func (d *SerialDriver) GetSSHHostname() (string, error) {
	d.Lock()
	defer d.Unlock()
	return d.Driver.GetSSHHostname()
}

// GetSSHKeyPath returns key path for use with ssh
func (d *SerialDriver) GetSSHKeyPath() string {
	d.Lock()
	defer d.Unlock()
	return d.Driver.GetSSHKeyPath()
}

// GetSSHPort returns port for use with ssh
func (d *SerialDriver) GetSSHPort() (int, error) {
	d.Lock()
	defer d.Unlock()
	return d.Driver.GetSSHPort()
}

// GetSSHUsername returns username for use with ssh
func (d *SerialDriver) GetSSHUsername() string {
	d.Lock()
	defer d.Unlock()
	return d.Driver.GetSSHUsername()
}

// GetURL returns a Docker compatible host URL for connecting to this host
// e.g. tcp://1.2.3.4:2376
func (d *SerialDriver) GetURL() (string, error) {
	d.Lock()
	defer d.Unlock()
	return d.Driver.GetURL()
}

// GetMachineState returns the state that the host is in (running, stopped, etc)
func (d *SerialDriver) GetMachineState() (state.State, error) {
	d.Lock()
	defer d.Unlock()
	return d.Driver.GetMachineState()
}

// KillMachine stops a host forcefully
func (d *SerialDriver) KillMachine() error {
	d.Lock()
	defer d.Unlock()
	return d.Driver.KillMachine()
}

// PreCreateCheck allows for pre-create operations to make sure a driver is ready for creation
func (d *SerialDriver) PreCreateCheck() error {
	d.Lock()
	defer d.Unlock()
	return d.Driver.PreCreateCheck()
}

// RemoveMachine a host
func (d *SerialDriver) RemoveMachine() error {
	d.Lock()
	defer d.Unlock()
	return d.Driver.RemoveMachine()
}

// RestartMachine restarts a host. This may just call StopMachine(); StartMachine()
// if the provider does not have any special restart behaviour.
func (d *SerialDriver) RestartMachine() error {
	d.Lock()
	defer d.Unlock()
	return d.Driver.RestartMachine()
}

// SetConfigFromFlags configures the driver with the object that was returned
// by RegisterCreateFlags
func (d *SerialDriver) SetConfigFromFlags(opts DriverOptions) error {
	d.Lock()
	defer d.Unlock()
	return d.Driver.SetConfigFromFlags(opts)
}

// StartMachine a host
func (d *SerialDriver) StartMachine() error {
	d.Lock()
	defer d.Unlock()
	return d.Driver.StartMachine()
}

// StopMachine a host gracefully
func (d *SerialDriver) StopMachine() error {
	d.Lock()
	defer d.Unlock()
	return d.Driver.StopMachine()
}

func (d *SerialDriver) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Driver)
}
