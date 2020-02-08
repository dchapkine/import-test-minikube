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

package node

import (
	"errors"

	"github.com/golang/glog"
	"github.com/spf13/viper"
	"k8s.io/minikube/pkg/minikube/config"
)

// Add adds a new node config to an existing cluster.
func Add(cc *config.MachineConfig, name string, controlPlane bool, worker bool, k8sVersion string, profileName string) error {
	n := config.Node{
		Name:   name,
		Worker: true,
	}

	if controlPlane {
		n.ControlPlane = true
	}

	if worker {
		n.Worker = true
	}

	if k8sVersion != "" {
		n.KubernetesVersion = k8sVersion
	} else {
		n.KubernetesVersion = cc.KubernetesConfig.KubernetesVersion
	}

	cc.Nodes = append(cc.Nodes, n)
	err := config.SaveProfile(profileName, cc)
	if err != nil {
		return err
	}

	return Start(cc, &n, false)
}

// Delete stops and deletes the given node from the given cluster
func Delete(cc *config.MachineConfig, name string) error {
	nd, index, err := Retrieve(cc, name)
	if err != nil {
		return err
	}

	err = Stop(cc, nd)
	if err != nil {
		glog.Warningf("Failed to stop node %s. Will still try to delete.", name)
	}

	cc.Nodes = append(cc.Nodes[:index], cc.Nodes[index+1:]...)
	return config.SaveProfile(viper.GetString(config.MachineProfile), cc)
}

// Stop stops the node in the given cluster
func Stop(cc *config.MachineConfig, n *config.Node) error {
	return nil
}

// Start spins up a guest and starts the kubernetes node.
func Start(cc *config.MachineConfig, n *config.Node, primary bool) error {
	// Throw all the slop from cmd.start in here
	return nil
}

// Retrieve finds the node by name in the given cluster
func Retrieve(cc *config.MachineConfig, name string) (*config.Node, int, error) {
	for i, n := range cc.Nodes {
		if n.Name == name {
			return &n, i, nil
		}
	}

	return nil, -1, errors.New("Could not find node " + name)
}
