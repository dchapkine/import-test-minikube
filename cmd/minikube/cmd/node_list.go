/*
Copyright 2020 The Kubernetes Authors All rights reserved.

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
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"k8s.io/minikube/pkg/minikube/driver"
	"k8s.io/minikube/pkg/minikube/exit"
	"k8s.io/minikube/pkg/minikube/mustload"
)

var nodeListCmd = &cobra.Command{
	Use:   "list",
	Short: "List nodes.",
	Long:  "List existing Minikube nodes.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			exit.UsageT("Usage: minikube node list")
		}

		cname := ClusterFlagValue()
		_, cc := mustload.Partial(cname)

		if len(cc.Nodes) < 1 {
			glog.Warningf("Did not found any Minikube node.")
		} else {
			glog.Infof("%v", cc.Nodes)
		}

		for _, n := range cc.Nodes {
			machineName := driver.MachineName(*cc, n)
			fmt.Printf("%s\t%s\n", machineName, n.IP)
		}
		os.Exit(0)
	},
}

func init() {
	nodeCmd.AddCommand(nodeListCmd)
}
