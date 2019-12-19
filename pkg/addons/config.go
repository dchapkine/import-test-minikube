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

package addons

import "k8s.io/minikube/pkg/minikube/config"

type setFn func(string, string, string) error

// Addon represents an addon
type Addon struct {
	name        string
	set         func(*config.MachineConfig, string, string) error
	validations []setFn
	callbacks   []setFn
}

// Addons is a list of all addons
var Addons = []*Addon{
	{
		name:      "addon-manager",
		set:       SetBool,
		callbacks: []setFn{EnableOrDisableAddon},
	},
	{
		name:      "dashboard",
		set:       SetBool,
		callbacks: []setFn{EnableOrDisableAddon},
	},

	{
		name:      "default-storageclass",
		set:       SetBool,
		callbacks: []setFn{EnableOrDisableStorageClasses},
	},
	{
		name:      "efk",
		set:       SetBool,
		callbacks: []setFn{EnableOrDisableAddon},
	},
	{
		name:      "freshpod",
		set:       SetBool,
		callbacks: []setFn{EnableOrDisableAddon},
	},
	{
		name:        "gvisor",
		set:         SetBool,
		validations: []setFn{IsContainerdRuntime},
		callbacks:   []setFn{EnableOrDisableAddon},
	},
	{
		name:      "helm-tiller",
		set:       SetBool,
		callbacks: []setFn{EnableOrDisableAddon},
	},
	{
		name:      "ingress",
		set:       SetBool,
		callbacks: []setFn{EnableOrDisableAddon},
	},
	{
		name:      "ingress-dns",
		set:       SetBool,
		callbacks: []setFn{EnableOrDisableAddon},
	},
	{
		name: "logviewer",
		set:  SetBool,
	},
	{
		name:      "metrics-server",
		set:       SetBool,
		callbacks: []setFn{EnableOrDisableAddon},
	},
	{
		name:      "nvidia-driver-installer",
		set:       SetBool,
		callbacks: []setFn{EnableOrDisableAddon},
	},
	{
		name:      "nvidia-gpu-device-plugin",
		set:       SetBool,
		callbacks: []setFn{EnableOrDisableAddon},
	},

	{
		name:      "registry",
		set:       SetBool,
		callbacks: []setFn{EnableOrDisableAddon},
	},
	{
		name:      "registry-creds",
		set:       SetBool,
		callbacks: []setFn{EnableOrDisableAddon},
	},
	{
		name:      "storage-provisioner",
		set:       SetBool,
		callbacks: []setFn{EnableOrDisableAddon},
	},
	{
		name:      "storage-provisioner-gluster",
		set:       SetBool,
		callbacks: []setFn{EnableOrDisableStorageClasses},
	},
}
