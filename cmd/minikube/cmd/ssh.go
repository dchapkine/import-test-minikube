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

package cmd

import (
	"os"

	"k8s.io/minikube/pkg/drivers"

	"github.com/spf13/cobra"
	"k8s.io/minikube/pkg/minikube/cluster"
	"k8s.io/minikube/pkg/minikube/config"
	"k8s.io/minikube/pkg/minikube/constants"
	"k8s.io/minikube/pkg/minikube/exit"
	"k8s.io/minikube/pkg/minikube/machine"
	"k8s.io/minikube/pkg/minikube/out"
)

// sshCmd represents the docker-ssh command
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Log into or run a command on a machine with SSH; similar to 'docker-machine ssh'",
	Long:  "Log into or run a command on a machine with SSH; similar to 'docker-machine ssh'.",
	Run: func(cmd *cobra.Command, args []string) {
		api, err := machine.NewAPIClient()
		if err != nil {
			exit.WithError("Error getting client", err)
		}
		defer api.Close()
		host, err := cluster.CheckIfHostExistsAndLoad(api, config.GetMachineName())
		if err != nil {
			exit.WithError("Error getting host", err)
		}
		if permissionError := drivers.ValidatePermissions(host.DriverName); permissionError != nil {
			exit.WithError("Permission denied", permissionError)
		}
		if host.Driver.DriverName() == constants.DriverNone {
			exit.UsageT("'none' driver does not support 'minikube ssh' command")
		}
		err = cluster.CreateSSHShell(api, args)
		if err != nil {
			// This is typically due to a non-zero exit code, so no need for flourish.
			out.ErrLn("ssh: %v", err)
			// It'd be nice if we could pass up the correct error code here :(
			os.Exit(exit.Failure)
		}
	},
}
