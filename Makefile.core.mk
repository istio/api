# Copyright Istio Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

SHELL := /bin/bash

all: gen

########################
# setup
########################

repo_dir := .

protolock = protolock
protolock_release = /bin/bash scripts/check-release-locks.sh
annotations_prep = annotations_prep
htmlproofer = htmlproofer
cue = cue-gen -paths=common-protos

#####################
# Generation Rules
#####################

.PHONY: gen-proto
gen-proto:
	./gen.sh

.PHONY: clean
clean:
	./clean.sh

.PHONY: gen
gen: \
	clean \
	gen-proto \
	generate-annotations \
	generate-labels \
	mirror-licenses \
	tidy-go

gen-check: gen check-clean-repo

#####################
# annotation/...
#####################

annotations_path := annotation
annotations_pb_go := $(annotations_path)/annotations.gen.go
annotations_pb_doc := $(annotations_path)/annotations.pb.html
annotations_yaml := $(annotations_path)/annotations.yaml

$(annotations_pb_go) $(annotations_pb_doc): $(annotations_yaml)
	@$(annotations_prep) --input $(annotations_yaml) --output $(annotations_pb_go) --html_output $(annotations_pb_doc) --collection_type annotation

generate-annotations: $(annotations_pb_go) $(annotations_pb_doc)

clean-annotations:
	@rm -fr $(annotations_pb_go) $(annotations_pb_doc)

#####################
# label/...
#####################

labels_path := label
labels_pb_go := $(labels_path)/labels.gen.go
labels_pb_doc := $(labels_path)/labels.pb.html
labels_yaml := $(labels_path)/labels.yaml

$(labels_pb_go) $(labels_pb_doc): $(labels_yaml)
	@$(annotations_prep) --input $(labels_yaml) --output $(labels_pb_go) --html_output $(labels_pb_doc) --collection_type label

generate-labels: $(labels_pb_go) $(labels_pb_doc)

clean-labels:
	@rm -fr $(labels_pb_go) $(labels_pb_doc)

#####################
# Protolock
#####################

proto-commit:
	@$(protolock) commit

proto-commit-force:
	@$(protolock) commit --force

proto-status:
	@$(protolock) status

release-lock-status:
	@$(protolock_release)

#####################
# Misc
#####################

# lint-protos is different for istio/api. List all other lint-all targets and add local-lint-protos
local-lint-protos:
	@buf lint

lint: lint-dockerfiles lint-scripts lint-yaml lint-helm lint-copyright-banner lint-go lint-python lint-markdown lint-sass lint-typescript lint-licenses local-lint-protos
 	@$(htmlproofer) . --url-swap "istio.io:preliminary.istio.io" --assume-extension --check-html --check-external-hash --check-opengraph --timeframe 2d --storage-dir $(repo_dir)/.htmlproofer --url-ignore "/localhost/"

fmt: format-python

#####################
# CI System
#####################

presubmit: proto-commit lint release-lock-status
postsubmit: presubmit

#####################
# Common definitions
#####################

include common/Makefile.common.mk
