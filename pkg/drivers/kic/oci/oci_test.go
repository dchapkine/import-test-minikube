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

package oci

import (
	"os"
	"testing"
)

func TestPointToHostDockerDaemonEmpty(t *testing.T) {
	_ = os.Setenv("DOCKER_HOST", "foo_host")
	_ = os.Setenv("DOCKER_CERT_PATH", "foo_cert_path")
	_ = os.Setenv("DOCKER_TLS_VERIFY", "foo_tls_verify")

	_ = os.Unsetenv("MINIKUBE_EXISTING_DOCKER_HOST")
	_ = os.Unsetenv("MINIKUBE_EXISTING_DOCKER_CERT_PATH")
	_ = os.Unsetenv("MINIKUBE_EXISTING_DOCKER_TLS_VERIFY")

	if err := PointToHostDockerDaemon(); err != nil {
		t.Fatalf("failed to set docker environment: got %v", err)
	}

	for _, key := range []string{
		"DOCKER_HOST", "DOCKER_CERT_PATH", "DOCKER_TLS_VERIFY",
	} {
		if v, set := os.LookupEnv(key); set {
			t.Errorf("%v env variable should not be set. got: %v", key, v)
		}
	}
}

func TestPointToHostDockerDaemon(t *testing.T) {
	_ = os.Setenv("DOCKER_HOST", "foo_host")
	_ = os.Setenv("DOCKER_CERT_PATH", "foo_cert_path")
	_ = os.Setenv("DOCKER_TLS_VERIFY", "foo_tls_verify")

	_ = os.Setenv("MINIKUBE_EXISTING_DOCKER_HOST", "bar_host")
	_ = os.Setenv("MINIKUBE_EXISTING_DOCKER_CERT_PATH", "bar_cert_path")
	_ = os.Setenv("MINIKUBE_EXISTING_DOCKER_TLS_VERIFY", "bar_tls_verify")

	if err := PointToHostDockerDaemon(); err != nil {
		t.Fatalf("failed to set docker environment: got %v", err)
	}

	expected := []struct {
		key, value string
	}{
		{"DOCKER_HOST", "bar_host"},
		{"DOCKER_CERT_PATH", "bar_cert_path"},
		{"DOCKER_TLS_VERIFY", "bar_tls_verify"},
	}
	for _, exp := range expected {
		if v := os.Getenv(exp.key); v != exp.value {
			t.Errorf("invalid %v env variable. got: %v, want: %v", exp.value, v, exp.value)
		}
	}
}
