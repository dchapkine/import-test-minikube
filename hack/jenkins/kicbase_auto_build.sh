#!/bin/bash

# Copyright 2021 The Kubernetes Authors All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -x -o pipefail

# Make sure docker is installed and configured
./hack/jenkins/installers/check_install_docker.sh
yes|gcloud auth configure-docker
docker login -u ${DOCKERHUB_USER} -p ${DOCKERHUB_PASS}

# Make sure gh is installed and configured
./hack/jenkins/installers/check_install_gh.sh

# Let's make sure we have the newest kicbase reference
curl -L https://github.com/kubernetes/minikube/raw/master/pkg/drivers/kic/types.go --output types-head.go
# kicbase tags are of the form VERSION-TIMESTAMP-PR, so this grep finds that TIMESTAMP in the middle
# if it doesn't exist, it will just return VERSION, which is covered in the if statement below
HEAD_KIC_TIMESTAMP=$(egrep "Version =" types-head.go | cut -d \" -f 2 | cut -d "-" -f 2)
CURRENT_KIC_TS=$(egrep "Version =" pkg/drivers/kic/types.go | cut -d \" -f 2 | cut -d "-" -f 2)
if [[ $HEAD_KIC_TIMESTAMP != v* ]]; then
	diff=$((CURRENT_KIC_TS-HEAD_KIC_TIMESTAMP))
	if [[ $CURRENT_KIC_TS == v* ]] || [ $diff -lt 0 ]; then
		gh pr comment ${ghprbPullId} --body "Hi ${ghprbPullAuthorLoginMention}, your kicbase info is out of date. Please rebase."
		exit 1
	fi
fi
rm types-head.go

# Setup variables
if [[ -z $KIC_VERSION ]]; then
	# Testing PRs here
	release=false
	now=$(date +%s)
	KV=$(egrep "Version =" pkg/drivers/kic/types.go | cut -d \" -f 2 | cut -d "-" -f 1)
	GCR_REPO=gcr.io/k8s-minikube/kicbase-builds
	DH_REPO=kicbase/build
	export KIC_VERSION=$KV-$now-$ghprbPullId
else
	# Actual kicbase release here
	release=true
	GCR_REPO=${GCR_REPO:-gcr.io/k8s-minikube/kicbase}
	DH_REPO=${DH_REPO:-kicbase/stable}
	export KIC_VERSION
fi
GCR_IMG=${GCR_REPO}:${KIC_VERSION}
DH_IMG=${DH_REPO}:${KIC_VERSION}
export KICBASE_IMAGE_REGISTRIES="${GCR_IMG} ${DH_IMG}"


# Build a new kicbase image
CIBUILD=yes make push-kic-base-image | tee kic-logs.txt

# Abort with error message if above command failed
ec=$?
if [ $ec -gt 0 ]; then
	if [ "$release" = false ]; then
		gh pr comment ${ghprbPullId} --body "Hi ${ghprbPullAuthorLoginMention}, building a new kicbase image failed.  
		See the logs at:
	       	```
		https://storage.cloud.google.com/minikube-builds/logs/${ghprbPullId}/${ghprbActualCommit:0:7}/kic_image_build.txt
		```
		"
	fi
	exit $ec
fi

# Retrieve the sha from the new image
docker pull $GCR_IMG
fullsha=$(docker inspect --format='{{index .RepoDigests 0}}' $GCR_IMG)
sha=$(echo ${fullsha} | cut -d ":" -f 2)
git config user.name "minikube-bot"
git config user.email "minikube-bot@google.com"


if [ "$release" = false ]; then
	# Update the user's PR with the newly built kicbase image.

	git remote add ${ghprbPullAuthorLogin} git@github.com:${ghprbPullAuthorLogin}/minikube.git
	git fetch ${ghprbPullAuthorLogin}
	git checkout -b ${ghprbPullAuthorLogin}-${ghprbSourceBranch} ${ghprbPullAuthorLogin}/${ghprbSourceBranch}

	sed -i "s|Version = .*|Version = \"${KIC_VERSION}\"|;s|baseImageSHA = .*|baseImageSHA = \"${sha}\"|;s|gcrRepo = .*|gcrRepo = \"${GCR_REPO}\"|;s|dockerhubRepo = .*|dockerhubRepo = \"${DH_REPO}\"|" pkg/drivers/kic/types.go; make generate-docs;

	git commit -am "Updating kicbase image to ${KIC_VERSION}"
	git push ${ghprbPullAuthorLogin} HEAD:${ghprbSourceBranch}

	message="Hi ${ghprbPullAuthorLoginMention}, we have updated your PR with the reference to newly built kicbase image. Pull the changes locally if you want to test with them or update your PR further."
	if [ $? -gt 0 ]; then
		message="Hi ${ghprbPullAuthorLoginMention}, we failed to push the reference to the kicbase to your PR. Please run the following command and push manually.

		sed -i 's|Version = .*|Version = \"${KIC_VERSION}\"|;s|baseImageSHA = .*|baseImageSHA = \"${sha}\"|;s|gcrRepo = .*|gcrRepo = \"${GCR_REPO}\"|;s|dockerhubRepo = .*|dockerhubRepo = \"${DH_REPO}\"|' pkg/drivers/kic/types.go; make generate-docs;
		
		"
	fi

	gh pr comment ${ghprbPullId} --body "${message}"
else
	# We're releasing, so open a new PR with the newly released kicbase

	CONTAINER_ID=$(docker create $GCR_IMG)
	IMG_TOOLS=$(mktemp)
	docker cp "$CONTAINER_ID:/tools.csv" "$IMG_TOOLS"
	docker rm "$CONTAINER_ID" > /dev/null
	OTHER_KIC_VERSION_COUNT=$(ls site/content/en/docs/releaseInfo/kicToolVersions \
		| sed "/^${KIC_VERSION}\.md$/ d; /_index\.md/ d" \
		| wc -l)
	TOOL_CELLS="$(< "$IMG_TOOLS" awk -F, 'NR>1 { printf "|%s|%s|\\n", $1, $2 }')"
	sed "s/{{.Version}}/${KIC_VERSION}/g; s/{{.Order}}/${OTHER_KIC_VERSION_COUNT}/g; s/{{.ToolCells}}/${TOOL_CELLS}/" ./hack/jenkins/misc/kicbase_tool_versions.md.tmpl \
		> "site/content/en/docs/releaseInfo/kicToolVersions/${KIC_VERSION}.md"
	rm $IMG_TOOLS

	branch=kicbase-release-${KIC_VERSION}
	git checkout -b ${branch}

	sed -i "s|Version = .*|Version = \"${KIC_VERSION}\"|;s|baseImageSHA = .*|baseImageSHA = \"${sha}\"|;s|gcrRepo = .*|gcrRepo = \"${GCR_REPO}\"|;s|dockerhubRepo = .*|dockerhubRepo = \"${DH_REPO}\"|" pkg/drivers/kic/types.go
	make generate-docs

	git add pkg/drivers/kic/types.go site/content/en/docs/commands/start.md "site/content/en/docs/releaseInfo/kicToolVersions/${KIC_VERSION}.md"
	git commit -m "Update kicbase to ${KIC_VERSION}"
	git remote add minikube-bot git@github.com:minikube-bot/minikube.git
	git push -f minikube-bot ${branch}

	gh pr create --fill --base master --head minikube-bot:${branch}
fi
