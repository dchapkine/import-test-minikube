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

package cni

import (
	"bytes"
	_ "embed"
	"fmt"
	"os/exec"
	"text/template"

	"github.com/blang/semver/v4"
	"github.com/icza/dyno"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"k8s.io/klog/v2"
	"k8s.io/minikube/pkg/minikube/config"
	"k8s.io/minikube/pkg/util"
)

// Generated by running `make update-cilium-version`
//
//go:embed cilium.yaml
var ciliumYaml string

// Cilium is the Cilium CNI manager
type Cilium struct {
	cc config.ClusterConfig
}

// String returns a string representation of this CNI
func (c Cilium) String() string {
	return "Cilium"
}

// CIDR returns the default CIDR used by this CNI
func (c Cilium) CIDR() string {
	return DefaultPodCIDR
}

// GenerateCiliumYAML generates the .yaml file
func (c Cilium) GenerateCiliumYAML() ([]byte, error) {

	// see issue #19683, older Kubernetes versions cannot recognize appArmorProfile fields
	k8sVersion, err := util.ParseKubernetesVersion(c.cc.KubernetesConfig.KubernetesVersion)
	if err == nil && k8sVersion.LT(semver.MustParse("1.30.0")) {
		if ciliumYaml, err = removeAppArorProfile(ciliumYaml); err != nil {
			return nil, err
		}
	}

	podCIDR := DefaultPodCIDR

	klog.Infof("Using pod CIDR: %s", podCIDR)

	opts := struct {
		PodSubnet string
	}{
		PodSubnet: podCIDR,
	}
	ciliumTmpl := template.Must(template.New("name").Parse(ciliumYaml))
	b := bytes.Buffer{}
	configTmpl := ciliumTmpl

	klog.Infof("cilium options: %+v", opts)
	if err := configTmpl.Execute(&b, opts); err != nil {
		return nil, err
	}
	klog.Infof("cilium config:\n%s\n", b.String())
	return b.Bytes(), nil
}

// Apply enables the CNI
func (c Cilium) Apply(r Runner) error {
	// see https://kubernetes.io/docs/tasks/administer-cluster/network-policy-provider/cilium-network-policy/
	if _, err := r.RunCmd(exec.Command("sudo", "/bin/bash", "-c", "grep 'bpffs /sys/fs/bpf' /proc/mounts || sudo mount bpffs -t bpf /sys/fs/bpf")); err != nil {
		return errors.Wrap(err, "bpf mount")
	}

	ciliumCfg, err := c.GenerateCiliumYAML()
	if err != nil {
		return errors.Wrap(err, "generating cilium cfg")
	}

	return applyManifest(c.cc, r, manifestAsset(ciliumCfg))
}

func removeAppArorProfile(ciliumConfig string) (string, error) {
	// remove all appArmorProfile fields
	decoder := yaml.NewDecoder(bytes.NewBufferString(ciliumConfig))
	var buffer bytes.Buffer
	encoder := yaml.NewEncoder(&buffer)
	for {
		obj := map[string]interface{}{}
		err := decoder.Decode(&obj)
		if err != nil && err.Error() == "EOF" {
			// we have unmarshaled all objects
			break
		} else if err != nil {
			return "", fmt.Errorf("failed to unmarshal yaml: %v", err)
		}
		if err := dyno.Delete(obj, "appArmorProfile", "spec", "template", "spec", "securityContext"); err != nil {
			return "", fmt.Errorf("failed to remove securityContext yaml: %v", err)
		}
		if err := encoder.Encode(obj); err != nil {
			return "", fmt.Errorf("failed to encode yaml")
		}

	}
	return buffer.String(), nil
}
