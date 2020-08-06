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
out_path = /tmp

protoc = protoc -Icommon-protos -I.
protolock = protolock
protolock_release = /bin/bash scripts/check-release-locks.sh
prototool = prototool
annotations_prep = annotations_prep
htmlproofer = htmlproofer
cue = cue-gen -paths=common-protos
validate_crds = python3 $(repo_dir)/scripts/validate_crds.py

go_plugin_prefix := --go_out=plugins=grpc,
go_plugin := $(go_plugin_prefix):$(out_path)

########################
# protoc_gen_gogo*
########################

gogofast_plugin_prefix := --gogofast_out=plugins=grpc,
gogoslick_plugin_prefix := --gogoslick_out=plugins=grpc,

comma := ,
empty :=
space := $(empty) $(empty)

importmaps := \
	gogoproto/gogo.proto=github.com/gogo/protobuf/gogoproto \
	google/protobuf/any.proto=github.com/gogo/protobuf/types \
	google/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor \
	google/protobuf/duration.proto=github.com/gogo/protobuf/types \
	google/protobuf/struct.proto=github.com/gogo/protobuf/types \
	google/protobuf/timestamp.proto=github.com/gogo/protobuf/types \
	google/protobuf/wrappers.proto=github.com/gogo/protobuf/types \
	google/rpc/status.proto=istio.io/gogo-genproto/googleapis/google/rpc \
	google/rpc/code.proto=istio.io/gogo-genproto/googleapis/google/rpc \
	google/rpc/error_details.proto=istio.io/gogo-genproto/googleapis/google/rpc \
	google/api/field_behavior.proto=istio.io/gogo-genproto/googleapis/google/api \

# generate mapping directive with M<proto>:<go pkg>, format for each proto file
mapping_with_spaces := $(foreach map,$(importmaps),M$(map),)
gogo_mapping := $(subst $(space),$(empty),$(mapping_with_spaces))

gogofast_plugin := $(gogofast_plugin_prefix)$(gogo_mapping):$(out_path)
gogoslick_plugin := $(gogoslick_plugin_prefix)$(gogo_mapping):$(out_path)

########################
# protoc_gen_python
########################

python_output_path := python/istio_api
protoc_gen_python_prefix := --python_out=,
protoc_gen_python_plugin := $(protoc_gen_python_prefix):$(repo_dir)/$(python_output_path)

########################
# protoc_gen_docs
########################

protoc_gen_docs_plugin := --docs_out=warnings=true,dictionary=$(repo_dir)/dictionaries/en-US,custom_word_list=$(repo_dir)/dictionaries/custom.txt,mode=html_fragment_with_front_matter:$(repo_dir)/
protoc_gen_docs_plugin_per_file := --docs_out=warnings=true,dictionary=$(repo_dir)/dictionaries/en-US,custom_word_list=$(repo_dir)/dictionaries/custom.txt,per_file=true,mode=html_fragment_with_front_matter:$(repo_dir)/

########################
# protoc_gen_jsonshim
########################

protoc_gen_k8s_support_plugins := --jsonshim_out=$(gogo_mapping):$(out_path) --deepcopy_out=$(gogo_mapping):$(out_path)

#####################
# Generation Rules
#####################

gen: \
	generate-core \
    generate-type \
	generate-mcp \
	generate-mesh \
	generate-operator \
	generate-networking \
	generate-authn \
	generate-security \
	generate-analysis \
	generate-meta \
	generate-envoy \
	generate-annotations \
	generate-openapi-schema \
	generate-openapi-crd \
	tidy-go \
	mirror-licenses \

gen-check: clean gen check-clean-repo

#####################
# core/...
#####################

core_v1alpha1_path := core/v1alpha1
core_v1alpha1_protos := $(wildcard $(core_v1alpha1_path)/*.proto)
core_v1alpha1_pb_gos := $(core_v1alpha1_protos:.proto=.pb.go)
core_v1alpha1_pb_pythons := $(patsubst $(core_v1alpha1_path)/%.proto,$(python_output_path)/$(core_v1alpha1_path)/%_pb2.py,$(core_v1alpha1_protos))
core_v1alpha1_pb_docs := $(core_v1alpha1_protos:.proto=.pb.html)
core_v1alpha1_openapi := $(core_v1alpha1_protos:.proto=.gen.json)

$(core_v1alpha1_pb_gos) $(core_v1alpha1_pb_docs) $(core_v1alpha1_pb_pythons): $(core_v1alpha1_protos)
	@$(protolock) status
	@$(protoc) $(gogofast_plugin)  $(protoc_gen_python_plugin) $(protoc_gen_docs_plugin)$(core_v1alpha1_path) $^
	@cp -r /tmp/istio.io/api/core/* core

generate-core: $(core_v1alpha1_pb_gos) $(core_v1alpha1_pb_docs) $(core_v1alpha1_pb_pythons)

clean-core:
	@rm -fr $(core_v1alpha1_pb_gos) $(core_v1alpha1_pb_docs) $(core_v1alpha1_pb_pythons)

#####################
# type/...
#####################

type_v1beta1_path := type/v1beta1
type_v1beta1_protos := $(wildcard $(type_v1beta1_path)/*.proto)
type_v1beta1_pb_gos := $(type_v1beta1_protos:.proto=.pb.go)
type_v1beta1_pb_pythons := $(patsubst $(type_v1beta1_path)/%.proto,$(python_output_path)/$(type_v1beta1_path)/%_pb2.py,$(type_v1beta1_protos))
type_v1beta1_pb_doc := $(type_v1beta1_path)/istio.type.v1beta1.pb.html
type_v1beta1_openapi := $(type_v1beta1_protos:.proto=.gen.json)
type_v1beta1_k8s_gos := \
	$(patsubst $(type_v1beta1_path)/%.proto,$(type_v1beta1_path)/%_json.gen.go,$(shell grep -l "^ *oneof " $(type_v1beta1_protos))) \
	$(patsubst $(type_v1beta1_path)/%.proto,$(type_v1beta1_path)/%_deepcopy.gen.go,$(shell grep -l "+kubetype-gen" $(type_v1beta1_protos)))

$(type_v1beta1_pb_gos) $(type_v1beta1_pb_doc) $(type_v1beta1_pb_pythons) $(type_v1beta1_k8s_gos): $(type_v1beta1_protos)
	@$(protolock) status
	@$(protoc) $(gogofast_plugin) $(protoc_gen_k8s_support_plugins) $(protoc_gen_docs_plugin)$(type_v1beta1_path) $(protoc_gen_python_plugin) $^
	@cp -r /tmp/istio.io/api/type/* type

generate-type: $(type_v1beta1_pb_gos) $(type_v1beta1_pb_doc) $(type_v1beta1_pb_pythons) $(type_v1beta1_k8s_gos)

clean-type:
	@rm -fr $(type_v1beta1_pb_gos) $(type_v1beta1_pb_docs) $(type_v1beta1_pb_pythons) $(type_v1beta1_k8s_gos)

#####################
# mcp/...
#####################

mcp_v1alpha1_path := mcp/v1alpha1
mcp_v1alpha1_protos := $(wildcard $(mcp_v1alpha1_path)/*.proto)
mcp_v1alpha1_pb_gos := $(mcp_v1alpha1_protos:.proto=.pb.go)
mcp_v1alpha1_pb_pythons := $(patsubst $(mcp_v1alpha1_path)/%.proto,$(python_output_path)/$(mcp_v1alpha1_path)/%_pb2.py,$(mcp_v1alpha1_protos))
mcp_v1alpha1_openapi := $(mcp_v1alpha1_path)/istio.mcp.v1alpha1.gen.json

$(mcp_v1alpha1_pb_gos) $(mcp_v1alpha1_pb_pythons): $(mcp_v1alpha1_protos)
	@$(protolock) status
	@$(protoc) $(gogofast_plugin) $(protoc_gen_python_plugin) $^
	@cp -r /tmp/istio.io/api/mcp/* mcp

generate-mcp: $(mcp_v1alpha1_pb_gos) $(mcp_v1alpha1_pb_doc) $(mcp_v1alpha1_pb_pythons)

clean-mcp:
	@rm -fr $(mcp_v1alpha1_pb_gos) $(mcp_v1alpha1_pb_pythons)

#####################
# mesh/...
#####################

mesh_v1alpha1_path := mesh/v1alpha1
mesh_v1alpha1_protos := $(wildcard $(mesh_v1alpha1_path)/*.proto)
mesh_v1alpha1_pb_gos := $(mesh_v1alpha1_protos:.proto=.pb.go)
mesh_v1alpha1_pb_pythons := $(patsubst $(mesh_v1alpha1_path)/%.proto,$(python_output_path)/$(mesh_v1alpha1_path)/%_pb2.py,$(mesh_v1alpha1_protos))
mesh_v1alpha1_pb_doc := $(mesh_v1alpha1_path)/istio.mesh.v1alpha1.pb.html
mesh_v1alpha1_openapi := $(mesh_v1alpha1_path)/istio.mesh.v1alpha1.gen.json

$(mesh_v1alpha1_pb_gos) $(mesh_v1alpha1_pb_doc) $(mesh_v1alpha1_pb_pythons): $(mesh_v1alpha1_protos)
	@$(protolock) status
	@$(protoc) $(gogofast_plugin) $(protoc_gen_docs_plugin)$(mesh_v1alpha1_path) $(protoc_gen_python_plugin) $^
	@cp -r /tmp/istio.io/api/mesh/* mesh
	@go run $(repo_dir)/operator/fixup_structs/main.go -f $(mesh_v1alpha1_path)/config.pb.go
	@go run $(repo_dir)/operator/fixup_structs/main.go -f $(mesh_v1alpha1_path)/network.pb.go
	@go run $(repo_dir)/operator/fixup_structs/main.go -f $(mesh_v1alpha1_path)/proxy.pb.go

generate-mesh: $(mesh_v1alpha1_pb_gos) $(mesh_v1alpha1_pb_doc) $(mesh_v1alpha1_pb_pythons)

clean-mesh:
	@rm -fr $(mesh_v1alpha1_pb_gos) $(mesh_v1alpha1_pb_doc) $(mesh_v1alpha1_pb_pythons)

#####################
# operator/...
#####################

operator_v1alpha1_path := operator/v1alpha1
operator_v1alpha1_protos := $(wildcard $(operator_v1alpha1_path)/*.proto)
operator_v1alpha1_pb_gos := $(operator_v1alpha1_protos:.proto=.pb.go)
operator_v1alpha1_pb_pythons := $(patsubst $(operator_v1alpha1_path)/%.proto,$(python_output_path)/$(operator_v1alpha1_path)/%_pb2.py,$(operator_v1alpha1_protos))
operator_v1alpha1_pb_doc := $(operator_v1alpha1_path)/istio.operator.v1alpha1.pb.html
operator_v1alpha1_k8s_gos := \
	$(patsubst $(operator_v1alpha1_path)/%.proto,$(operator_v1alpha1_path)/%_deepcopy.gen.go,$(shell grep -l "+kubetype-gen" $(operator_v1alpha1_protos)))


$(operator_v1alpha1_pb_gos) $(operator_v1alpha1_pb_doc) $(operator_v1alpha1_pb_pythons) $(operator_v1alpha1_k8s_gos): $(operator_v1alpha1_protos)
	@$(protolock) status
	@$(protoc) $(gogofast_plugin) $(protoc_gen_docs_plugin)$(operator_v1alpha1_path) $(protoc_gen_python_plugin) $^
	@cp -r /tmp/istio.io/api/operator/* operator
	@go run $(repo_dir)/operator/fixup_structs/main.go -f $(operator_v1alpha1_path)/operator.pb.go
	@sed -i 's|<key,value,effect>|\&lt\;key,value,effect\&gt\;|g' $(operator_v1alpha1_path)/istio.operator.v1alpha1.pb.html
	@sed -i 's|<operator>|\&lt\;operator\&gt\;|g' $(operator_v1alpha1_path)/istio.operator.v1alpha1.pb.html

generate-operator: $(operator_v1alpha1_pb_gos) $(operator_v1alpha1_pb_doc) $(operator_v1alpha1_pb_pythons) $(operator_v1alpha1_k8s_gos)

clean-operator:
	@rm -fr $(operator_v1alpha1_pb_gos) $(operator_v1alpha1_pb_doc) $(operator_v1alpha1_pb_pythons) $(operator_v1alpha1_k8s_gos)

#####################
# networking/...
#####################

networking_v1alpha3_path := networking/v1alpha3
networking_v1alpha3_protos := $(wildcard $(networking_v1alpha3_path)/*.proto)
networking_v1alpha3_pb_gos := $(networking_v1alpha3_protos:.proto=.pb.go)
networking_v1alpha3_pb_pythons := $(patsubst $(networking_v1alpha3_path)/%.proto,$(python_output_path)/$(networking_v1alpha3_path)/%_pb2.py,$(networking_v1alpha3_protos))
networking_v1alpha3_pb_docs := $(networking_v1alpha3_protos:.proto=.pb.html)
networking_v1alpha3_openapi := $(networking_v1alpha3_protos:.proto=.gen.json)
networking_v1alpha3_k8s_gos := \
	$(patsubst $(networking_v1alpha3_path)/%.proto,$(networking_v1alpha3_path)/%_json.gen.go,$(shell grep -l "^ *oneof " $(networking_v1alpha3_protos))) \
	$(patsubst $(networking_v1alpha3_path)/%.proto,$(networking_v1alpha3_path)/%_deepcopy.gen.go,$(shell grep -l "+kubetype-gen" $(networking_v1alpha3_protos)))

$(networking_v1alpha3_pb_gos) $(networking_v1alpha3_pb_docs) $(networking_v1alpha3_pb_pythons) $(networking_v1alpha3_k8s_gos): $(networking_v1alpha3_protos)
	@$(protolock) status
	@$(protoc) $(gogofast_plugin) $(protoc_gen_k8s_support_plugins) $(protoc_gen_docs_plugin_per_file)$(networking_v1alpha3_path) $(protoc_gen_python_plugin) $^
	@cp -r /tmp/istio.io/api/networking/v1alpha3/* networking/v1alpha3

networking_v1beta1_path := networking/v1beta1
networking_v1beta1_protos := $(wildcard $(networking_v1beta1_path)/*.proto)
networking_v1beta1_pb_gos := $(networking_v1beta1_protos:.proto=.pb.go)
networking_v1beta1_pb_pythons := $(patsubst $(networking_v1beta1_path)/%.proto,$(python_output_path)/$(networking_v1beta1_path)/%_pb2.py,$(networking_v1beta1_protos))
# v1beta1 docs are not generated as v1beta1 has the same fields as v1alpha3. Thus the target here is only for `make clean` purpose.
networking_v1beta1_pb_docs := $(networking_v1beta1_protos:.proto=.pb.html)
networking_v1beta1_openapi := $(networking_v1beta1_protos:.proto=.gen.json)
networking_v1beta1_k8s_gos := \
	$(patsubst $(networking_v1beta1_path)/%.proto,$(networking_v1beta1_path)/%_json.gen.go,$(shell grep -l "^ *oneof " $(networking_v1beta1_protos))) \
	$(patsubst $(networking_v1beta1_path)/%.proto,$(networking_v1beta1_path)/%_deepcopy.gen.go,$(shell grep -l "+kubetype-gen" $(networking_v1beta1_protos)))

$(networking_v1beta1_pb_gos) $(networking_v1beta1_pb_pythons) $(networking_v1beta1_k8s_gos): $(networking_v1beta1_protos)
	@$(protolock) status
	@$(protoc) $(gogofast_plugin) $(protoc_gen_k8s_support_plugins) $(protoc_gen_python_plugin) $^
	@cp -r /tmp/istio.io/api/networking/v1beta1/* networking/v1beta1

generate-networking: $(networking_v1alpha3_pb_gos) $(networking_v1alpha3_pb_docs) $(networking_v1alpha3_pb_pythons) $(networking_v1alpha3_k8s_gos) $(networking_v1beta1_pb_gos) $(networking_v1beta1_pb_pythons) $(networking_v1beta1_k8s_gos)

clean-networking:
	@rm -fr $(networking_v1alpha3_pb_gos) $(networking_v1alpha3_pb_docs) $(networking_v1alpha3_pb_pythons) $(networking_v1alpha3_k8s_gos) \
	$(networking_v1beta1_pb_gos) $(networking_v1beta1_pb_docs) $(networking_v1beta1_pb_pythons) $(networking_v1beta1_k8s_gos)

#####################
# authentication/...
#####################

authn_v1alpha1_path := authentication/v1alpha1
authn_v1alpha1_protos := $(wildcard $(authn_v1alpha1_path)/*.proto)
authn_v1alpha1_pb_gos := $(authn_v1alpha1_protos:.proto=.pb.go)
authn_v1alpha1_pb_pythons := $(patsubst $(authn_v1alpha1_path)/%.proto,$(python_output_path)/$(authn_v1alpha1_path)/%_pb2.py,$(authn_v1alpha1_protos))
authn_v1alpha1_pb_doc := $(authn_v1alpha1_path)/istio.authentication.v1alpha1.pb.html
authn_v1alpha1_openapi := $(authn_v1alpha1_path)/istio.authentication.v1alpha1.gen.json
authn_v1alpha1_k8s_gos := \
	$(patsubst $(authn_v1alpha1_path)/%.proto,$(authn_v1alpha1_path)/%_json.gen.go,$(shell grep -l "^ *oneof " $(authn_v1alpha1_protos))) \
	$(patsubst $(authn_v1alpha1_path)/%.proto,$(authn_v1alpha1_path)/%_deepcopy.gen.go,$(shell grep -l "+kubetype-gen" $(authn_v1alpha1_protos)))

$(authn_v1alpha1_pb_gos) $(authn_v1alpha1_pb_doc) $(authn_v1alpha1_pb_pythons) $(authn_v1alpha1_k8s_gos): $(authn_v1alpha1_protos)
	@$(protolock) status
	@$(protoc) $(gogofast_plugin) $(protoc_gen_k8s_support_plugins) $(protoc_gen_docs_plugin)$(authn_v1alpha1_path) $(protoc_gen_python_plugin) $^
	@cp -r /tmp/istio.io/api/authentication/* authentication

generate-authn: $(authn_v1alpha1_pb_gos) $(authn_v1alpha1_pb_doc) $(authn_v1alpha1_pb_pythons) $(authn_v1alpha1_k8s_gos)

clean-authn:
	@rm -fr $(authn_v1alpha1_pb_gos) $(authn_v1alpha1_pb_doc) $(authn_v1alpha1_pb_pythons) $(authn_v1alpha1_k8s_gos)

#####################
# security/...
#####################

security_v1alpha1_path := security/v1alpha1
security_v1alpha1_protos := $(wildcard $(security_v1alpha1_path)/*.proto)
security_v1alpha1_pb_gos := $(security_v1alpha1_protos:.proto=.pb.go)
security_v1alpha1_pb_pythons := $(patsubst $(security_v1alpha1_path)/%.proto,$(python_output_path)/$(security_v1alpha1_path)/%_pb2.py,$(security_v1alpha1_protos))
security_v1alpha1_pb_docs := $(security_v1alpha1_protos:.proto=.pb.html)
security_v1alpha1_openapi := $(security_v1alpha1_protos:.proto=.gen.json)
security_v1alpha1_k8s_gos := \
	$(patsubst $(security_v1alpha1_path)/%.proto,$(security_v1alpha1_path)/%_json.gen.go,$(shell grep -l "^ *oneof " $(security_v1alpha1_protos))) \
	$(patsubst $(security_v1alpha1_path)/%.proto,$(security_v1alpha1_path)/%_deepcopy.gen.go,$(shell grep -l "+kubetype-gen" $(security_v1alpha1_protos)))

$(security_v1alpha1_pb_gos) $(security_v1alpha1_pb_docs) $(security_v1alpha1_pb_pythons) $(security_v1alpha1_k8s_gos): $(security_v1alpha1_protos)
	@$(protolock) status
	@$(protoc) $(gogofast_plugin) $(protoc_gen_k8s_support_plugins) $(protoc_gen_docs_plugin_per_file)$(security_v1alpha1_path) $(protoc_gen_python_plugin) $^
	@cp -r /tmp/istio.io/api/security/* security

security_v1beta1_path := security/v1beta1
security_v1beta1_protos := $(wildcard $(security_v1beta1_path)/*.proto)
security_v1beta1_pb_gos := $(security_v1beta1_protos:.proto=.pb.go)
security_v1beta1_pb_pythons := $(patsubst $(security_v1beta1_path)/%.proto,$(python_output_path)/$(security_v1beta1_path)/%_pb2.py,$(security_v1beta1_protos))
security_v1beta1_pb_docs := $(security_v1beta1_protos:.proto=.pb.html)
security_v1beta1_openapi := $(security_v1beta1_protos:.proto=.gen.json)
security_v1beta1_k8s_gos := \
	$(patsubst $(security_v1beta1_path)/%.proto,$(security_v1beta1_path)/%_json.gen.go,$(shell grep -l "^ *oneof " $(security_v1beta1_protos))) \
	$(patsubst $(security_v1beta1_path)/%.proto,$(security_v1beta1_path)/%_deepcopy.gen.go,$(shell grep -l "+kubetype-gen" $(security_v1beta1_protos)))

$(security_v1beta1_pb_gos) $(security_v1beta1_pb_docs) $(security_v1beta1_pb_pythons) $(security_v1beta1_k8s_gos): $(security_v1beta1_protos)
	@$(protolock) status
	@$(protoc) $(gogofast_plugin) $(protoc_gen_k8s_support_plugins) $(protoc_gen_docs_plugin_per_file)$(security_v1beta1_path) $(protoc_gen_python_plugin) $^
	@cp -r /tmp/istio.io/api/security/* security

generate-security: $(security_v1beta1_pb_gos) $(security_v1beta1_pb_docs) $(security_v1beta1_pb_pythons) $(security_v1beta1_k8s_gos) $(security_v1alpha1_pb_gos) $(security_v1alpha1_pb_docs) $(security_v1alpha1_pb_pythons) $(security_v1alpha1_k8s_gos) 

clean-security:
	@rm -fr $(security_v1beta1_pb_gos) $(security_v1beta1_pb_docs) $(security_v1beta1_pb_pythons) $(security_v1beta1_k8s_gos)

#####################
# analysis/...
#####################

analysis_v1alpha1_path := analysis/v1alpha1
analysis_v1alpha1_protos := $(wildcard $(analysis_v1alpha1_path)/*.proto)
analysis_v1alpha1_pb_gos := $(analysis_v1alpha1_protos:.proto=.pb.go)
analysis_v1alpha1_pb_pythons := $(patsubst $(analysis_v1alpha1_path)/%.proto,$(python_output_path)/$(analysis_v1alpha1_path)/%_pb2.py,$(analysis_v1alpha1_protos))
analysis_v1alpha1_pb_docs := $(analysis_v1alpha1_protos:.proto=.pb.html)
analysis_v1alpha1_openapi := $(analysis_v1alpha1_protos:.proto=.gen.json)
analysis_v1alpha1_k8s_gos := \
	$(patsubst $(analysis_v1alpha1_path)/%.proto,$(analysis_v1alpha1_path)/%_json.gen.go,$(shell grep -l "^ *oneof " $(analysis_v1alpha1_protos))) \
	$(patsubst $(analysis_v1alpha1_path)/%.proto,$(analysis_v1alpha1_path)/%_deepcopy.gen.go,$(shell grep -l "+kubetype-gen" $(analysis_v1alpha1_protos)))

$(analysis_v1alpha1_pb_gos) $(analysis_v1alpha1_pb_docs) $(analysis_v1alpha1_pb_pythons) $(analysis_v1alpha1_k8s_gos): $(analysis_v1alpha1_protos)
	@$(protolock) status
	@$(protoc) $(gogofast_plugin) $(protoc_gen_k8s_support_plugins) $(protoc_gen_docs_plugin_per_file)$(analysis_v1alpha1_path) $(protoc_gen_python_plugin) $^
	@cp -r /tmp/istio.io/api/analysis/* analysis

generate-analysis: $(analysis_v1alpha1_pb_gos) $(analysis_v1alpha1_pb_docs) $(analysis_v1alpha1_pb_pythons) $(analysis_v1alpha1_k8s_gos)

clean-analysis:
	@rm -fr $(analysis_v1alpha1_pb_gos) $(analysis_v1alpha1_pb_docs) $(analysis_v1alpha1_pb_pythons) $(analysis_v1alpha1_k8s_gos)

#####################
# meta/...
#####################

meta_v1alpha1_path := meta/v1alpha1
meta_v1alpha1_protos := $(wildcard $(meta_v1alpha1_path)/*.proto)
meta_v1alpha1_pb_gos := $(meta_v1alpha1_protos:.proto=.pb.go)
meta_v1alpha1_pb_pythons := $(patsubst $(meta_v1alpha1_path)/%.proto,$(python_output_path)/$(meta_v1alpha1_path)/%_pb2.py,$(meta_v1alpha1_protos))
meta_v1alpha1_pb_docs := $(meta_v1alpha1_protos:.proto=.pb.html)
meta_v1alpha1_openapi := $(meta_v1alpha1_protos:.proto=.gen.json)
meta_v1alpha1_k8s_gos := \
	$(patsubst $(meta_v1alpha1_path)/%.proto,$(meta_v1alpha1_path)/%_json.gen.go,$(shell grep -l "^ *oneof " $(meta_v1alpha1_protos))) \
	$(patsubst $(meta_v1alpha1_path)/%.proto,$(meta_v1alpha1_path)/%_deepcopy.gen.go,$(shell grep -l "+kubetype-gen" $(meta_v1alpha1_protos)))

$(meta_v1alpha1_pb_gos) $(meta_v1alpha1_pb_docs) $(meta_v1alpha1_pb_pythons) $(meta_v1alpha1_k8s_gos): $(meta_v1alpha1_protos)
	@$(protolock) status
	@$(protoc) $(gogofast_plugin) $(protoc_gen_k8s_support_plugins) $(protoc_gen_docs_plugin_per_file)$(meta_v1alpha1_path) $(protoc_gen_python_plugin) $^
	@cp -r /tmp/istio.io/api/meta/* meta

generate-meta: $(meta_v1alpha1_pb_gos) $(meta_v1alpha1_pb_docs) $(meta_v1alpha1_pb_pythons) $(meta_v1alpha1_k8s_gos)

clean-meta:
	@rm -fr $(meta_v1alpha1_pb_gos) $(meta_v1alpha1_pb_docs) $(meta_v1alpha1_pb_pythons) $(meta_v1alpha1_k8s_gos)

#####################
# envoy/...
#####################

envoy_path := envoy
envoy_protos := $(shell find $(envoy_path) -type f -name '*.proto' | sort)
envoy_pb_pythons := $(patsubst $(envoy_path)/%.proto,$(python_output_path)/$(envoy_path)/%_pb2.py,$(envoy_protos))

$(envoy_pb_pythons): $(envoy_protos)
	@$(protolock) status
	@$(protoc) $(protoc_gen_python_plugin) $^

generate-envoy: $(envoy_pb_pythons)

clean-envoy:
	@rm -fr $(envoy_pb_pythons)

#####################
# annotation/...
#####################

annotations_path := annotation
annotations_pb_go := $(annotations_path)/annotations.gen.go
annotations_pb_doc := $(annotations_path)/annotations.pb.html

$(annotations_pb_go) $(annotations_pb_doc): $(annotations_path)/annotations.yaml
	@$(annotations_prep) --input $(annotations_path)/annotations.yaml --output $(annotations_pb_go) --html_output $(annotations_pb_doc)

generate-annotations: $(annotations_pb_go) $(annotations_pb_doc)

clean-annotations:
	@rm -fr $(annotations_pb_go) $(annotations_pb_doc)

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

lint: lint-all
 	@$(htmlproofer) . --url-swap "istio.io:preliminary.istio.io" --assume-extension --check-html --check-external-hash --check-opengraph --timeframe 2d --storage-dir $(repo_dir)/.htmlproofer --url-ignore "/localhost/"

fmt: format-python

#####################
# OpenAPI Schema
#####################

all_protos := \
	$(core_v1alpha1_protos) \
	$(mcp_v1alpha1_protos) \
	$(mesh_v1alpha1_protos) \
	$(operator_v1alpha1_protos) \
	$(networking_v1alpha3_protos) \
	$(networking_v1beta1_protos) \
	$(authn_v1alpha1_protos) \
	$(security_v1beta1_protos) \
	$(analysis_v1alpha1_protos) \
	$(meta_v1alpha1_protos) \
	$(type_v1beta1_protos)

all_openapi := \
	$(core_v1alpha1_openapi) \
	$(mcp_v1alpha1_openapi) \
	$(mesh_v1alpha1_openapi) \
	$(operator_v1alpha1_openapi) \
	$(networking_v1alpha3_openapi) \
	$(networking_v1beta1_openapi) \
	$(authn_v1alpha1_openapi) \
	$(security_v1beta1_openapi) \
	$(analysis_v1alpha1_openapi) \
	$(meta_v1alpha1_openapi) \
	$(type_v1beta1_openapi)

all_openapi_crd := kubernetes/customresourcedefinitions.gen.yaml

$(all_openapi): $(all_protos)
	@$(cue) -f=$(repo_dir)/cue.yaml

# The fields are added at the end to generate snake cases. This is a temporary solution to accommodate some wrong namings that currently exist.
$(all_openapi_crd): $(all_protos)
	@$(cue) -f=$(repo_dir)/cue.yaml --crd=true -snake=jwksUri,apiKeys,apiSpecs,includedPaths,jwtHeaders,triggerRules,excludedPaths,mirrorPercent
	# Some networking APIs need to have the same schema for their v1alpha3 and v1beta1 versions.
	@$(validate_crds) check_equal_schema --kinds VirtualService,DestinationRule,Gateway,Sidecar,ServiceEntry,WorkloadEntry --versions v1alpha3,v1beta1 --file $(all_openapi_crd)

generate-openapi-schema: $(all_openapi)

generate-openapi-crd: $(all_openapi_crd)

clean-openapi-schema:
	@rm -fr $(all_openapi)

clean-openapi-crd:
	@rm -f $(all_openapi_crd)

#####################
# Cleanup
#####################

clean: \
	clean-core \
	clean-mcp \
	clean-mesh \
	clean-operator \
	clean-networking \
	clean-authn \
	clean-envoy \
	clean-annotations \
	clean-openapi-schema \
	clean-security \
	clean-analysis \
	clean-meta \
	clean-type \
	clean-openapi-crd

#####################
# CI System
#####################

presubmit: proto-commit lint release-lock-status
postsubmit: presubmit

#####################
# Common definitions
#####################

include common/Makefile.common.mk
