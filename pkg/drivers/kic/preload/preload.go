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

package preload

import (
	"fmt"
	"os"
	"path"

	"github.com/golang/glog"
	"github.com/jimmidyson/go-download"
	"k8s.io/minikube/pkg/minikube/localpath"
)

// TarballPath returns the path to the preloaded tarball
func TarballPath(k8sVersion string) string {
	targetDir := localpath.MakeMiniPath("cache", "preloaded-tarball")
	targetFilepath := path.Join(targetDir, fmt.Sprintf("preloaded-images-k8s-%s.tar.lz4", k8sVersion))
	return targetFilepath
}

// CacheTarball caches the preloaded images tarball on the host machine
func CacheTarball(k8sVersion string) error {
	targetFilepath := TarballPath(k8sVersion)

	if _, err := os.Stat(targetFilepath); err == nil {
		glog.Infof("Found %s in cache, skipping downloading", targetFilepath)
		return nil
	}

	url := fmt.Sprintf("https://storage.googleapis.com/minikube-docker-volume-tarballs/preloaded-images-k8s-%s.tar", k8sVersion)
	glog.Infof("Downloading %s to %s", url, targetFilepath)
	return download.ToFile(url, targetFilepath, download.FileOptions{Mkdirs: download.MkdirAll})
}

// UsingPreloadedVolume returns true if the preloaded tarball exists
func UsingPreloadedVolume(k8sVersion string) bool {
	path := TarballPath(k8sVersion)
	_, err := os.Stat(path)
	return err == nil
}
