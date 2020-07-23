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
	"bytes"
	"encoding/json"
	"testing"
)

func TestLegacyExitCode(t *testing.T) {
	var tests = []struct {
		name  string
		want  int
		state *LegacyStatus
	}{
		{"ok", 0, &LegacyStatus{Host: "Running", Kubelet: "Running", APIServer: "Running", Kubeconfig: Configured}},
		{"paused", 2, &LegacyStatus{Host: "Running", Kubelet: "Stopped", APIServer: "Paused", Kubeconfig: Configured}},
		{"down", 7, &LegacyStatus{Host: "Stopped", Kubelet: "Stopped", APIServer: "Stopped", Kubeconfig: Misconfigured}},
		{"missing", 7, &LegacyStatus{Host: "Nonexistent", Kubelet: "Nonexistent", APIServer: "Nonexistent", Kubeconfig: "Nonexistent"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := legacyExitCode([]*LegacyStatus{tc.state})
			if got != tc.want {
				t.Errorf("exitcode(%+v) = %d, want: %d", tc.state, got, tc.want)
			}
		})
	}
}

func TestLegacyStatusText(t *testing.T) {
	var tests = []struct {
		name  string
		state *LegacyStatus
		want  string
	}{
		{
			name:  "ok",
			state: &LegacyStatus{Name: "minikube", Host: "Running", Kubelet: "Running", APIServer: "Running", Kubeconfig: Configured},
			want:  "minikube\ntype: Control Plane\nhost: Running\nkubelet: Running\napiserver: Running\nkubeconfig: Configured\n\n",
		},
		{
			name:  "paused",
			state: &LegacyStatus{Name: "minikube", Host: "Running", Kubelet: "Stopped", APIServer: "Paused", Kubeconfig: Configured},
			want:  "minikube\ntype: Control Plane\nhost: Running\nkubelet: Stopped\napiserver: Paused\nkubeconfig: Configured\n\n",
		},
		{
			name:  "down",
			state: &LegacyStatus{Name: "minikube", Host: "Stopped", Kubelet: "Stopped", APIServer: "Stopped", Kubeconfig: Misconfigured},
			want:  "minikube\ntype: Control Plane\nhost: Stopped\nkubelet: Stopped\napiserver: Stopped\nkubeconfig: Misconfigured\n\n\nWARNING: Your kubectl is pointing to stale minikube-vm.\nTo fix the kubectl context, run `minikube update-context`\n",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var b bytes.Buffer
			err := legacyStatusText(tc.state, &b, legacyStatusTmpl)
			if err != nil {
				t.Errorf("text(%+v) error: %v", tc.state, err)
			}

			got := b.String()
			if got != tc.want {
				t.Errorf("text(%+v) = %q, want: %q", tc.state, got, tc.want)
			}
		})
	}
}

func TestLegacyStatusJSON(t *testing.T) {
	var tests = []struct {
		name  string
		state *LegacyStatus
	}{
		{"ok", &LegacyStatus{Host: "Running", Kubelet: "Running", APIServer: "Running", Kubeconfig: Configured}},
		{"paused", &LegacyStatus{Host: "Running", Kubelet: "Stopped", APIServer: "Paused", Kubeconfig: Configured}},
		{"down", &LegacyStatus{Host: "Stopped", Kubelet: "Stopped", APIServer: "Stopped", Kubeconfig: Misconfigured}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var b bytes.Buffer
			err := legacyStatusJSON([]*LegacyStatus{tc.state}, &b)
			if err != nil {
				t.Errorf("json(%+v) error: %v", tc.state, err)
			}

			st := &LegacyStatus{}
			if err := json.Unmarshal(b.Bytes(), st); err != nil {
				t.Errorf("json(%+v) unmarshal error: %v", tc.state, err)
			}
		})
	}
}
