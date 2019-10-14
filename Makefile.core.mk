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
protoc_gen_docs_plugin_for_networking := --docs_out=warnings=true,dictionary=$(repo_dir)/dictionaries/en-US,custom_word_list=$(repo_dir)/dictionaries/custom.txt,per_file=true,mode=html_fragment_with_front_matter:$(repo_dir)/

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
	generate-mixer \
	generate-networking \
	generate-rbac \
	generate-authn \
	generate-security \
	generate-envoy \
	generate-policy \
	generate-annotations \
	generate-openapi-schema

gen-check: clean gen check-clean-repo

#####################
# core/...
#####################

core_v1alpha1_path := core/v1alpha1
core_v1alpha1_protos := $(wildcard $(core_v1alpha1_path)/*.proto)
core_v1alpha1_pb_gos := $(core_v1alpha1_protos:.proto=.pb.go)
core_v1alpha1_pb_pythons := $(patsubst $(core_v1alpha1_path)/%.proto,$(python_output_path)/$(core_v1alpha1_path)/%_pb2.py,$(core_v1alpha1_protos))
core_v1alpha1_pb_docs := $(core_v1alpha1_protos:.proto=.pb.html)
core_v1alpha1_openapi := $(core_v1alpha1_protos:.proto=.json)

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
type_v1beta1_openapi := $(type_v1beta1_protos:.proto=.json)
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
mcp_v1alpha1_openapi := $(mcp_v1alpha1_path)/istio.mcp.v1alpha1.json

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
mesh_v1alpha1_openapi := $(mesh_v1alpha1_path)/istio.mesh.v1alpha1.json

$(mesh_v1alpha1_pb_gos) $(mesh_v1alpha1_pb_doc) $(mesh_v1alpha1_pb_pythons): $(mesh_v1alpha1_protos)
	@$(protolock) status
	@$(protoc) $(gogofast_plugin) $(protoc_gen_docs_plugin)$(mesh_v1alpha1_path) $(protoc_gen_python_plugin) $^
	@cp -r /tmp/istio.io/api/mesh/* mesh

generate-mesh: $(mesh_v1alpha1_pb_gos) $(mesh_v1alpha1_pb_doc) $(mesh_v1alpha1_pb_pythons)

clean-mesh:
	@rm -fr $(mesh_v1alpha1_pb_gos) $(mesh_v1alpha1_pb_doc) $(mesh_v1alpha1_pb_pythons)

#####################
# policy/...
#####################

policy_v1beta1_path := policy/v1beta1
policy_v1beta1_protos := $(wildcard $(policy_v1beta1_path)/*.proto)
policy_v1beta1_pb_gos := $(policy_v1beta1_protos:.proto=.pb.go)
policy_v1beta1_pb_pythons := $(patsubst $(policy_v1beta1_path)/%.proto,$(python_output_path)/$(policy_v1beta1_path)/%_pb2.py,$(policy_v1beta1_protos))
policy_v1beta1_pb_doc := $(policy_v1beta1_path)/istio.policy.v1beta1.pb.html
policy_v1beta1_openapi := $(policy_v1beta1_path)/istio.policy.v1beta1.json
policy_v1beta1_k8s_gos := \
	$(patsubst $(policy_v1beta1_path)/%.proto,$(policy_v1beta1_path)/%_json.gen.go,$(shell grep -l "^ *oneof " $(policy_v1beta1_protos))) \
	$(patsubst $(policy_v1beta1_path)/%.proto,$(policy_v1beta1_path)/%_deepcopy.gen.go,$(shell grep -l "+kubetype-gen" $(policy_v1beta1_protos)))

$(policy_v1beta1_pb_gos) $(policy_v1beta1_pb_doc) $(policy_v1beta1_pb_pythons) $(policy_v1beta1_k8s_gos): $(policy_v1beta1_protos)
	@$(protolock) status
	@$(protoc) $(gogoslick_plugin) $(protoc_gen_k8s_support_plugins) $(protoc_gen_docs_plugin)$(policy_v1beta1_path) $(protoc_gen_python_plugin) $^
	@cp -r /tmp/istio.io/api/policy/* policy

generate-policy: $(policy_v1beta1_pb_gos) $(policy_v1beta1_pb_doc) $(policy_v1beta1_pb_pythons) $(policy_v1beta1_k8s_gos)

clean-policy:
	@rm -fr $(policy_v1beta1_pb_gos) policy/v1beta1/fixed_cfg.pb.go $(policy_v1beta1_pb_doc) $(policy_v1beta1_pb_pythons) $(policy_v1beta1_k8s_gos)

#####################
# mixer/...
#####################

mixer_v1_path := mixer/v1
mixer_v1_protos :=  $(wildcard $(mixer_v1_path)/*.proto)
mixer_v1_pb_gos := $(mixer_v1_protos:.proto=.pb.go)
mixer_v1_pb_pythons := $(patsubst $(mixer_v1_path)/%.proto,$(python_output_path)/$(mixer_v1_path)/%_pb2.py,$(mixer_v1_protos))
mixer_v1_openapi := $(mixer_v1_path)/istio.mixer.v1.json
mixer_v1_k8s_gos := \
	$(patsubst $(mixer_v1_path)/%.proto,$(mixer_v1_path)/%_json.gen.go,$(shell grep -l "^ *oneof " $(mixer_v1_protos))) \
	$(patsubst $(mixer_v1_path)/%.proto,$(mixer_v1_path)/%_deepcopy.gen.go,$(shell grep -l "+kubetype-gen" $(mixer_v1_protos)))

mixer_config_client_path := mixer/v1/config/client
mixer_config_client_protos := $(wildcard $(mixer_config_client_path)/*.proto)
mixer_config_client_pb_gos := $(mixer_config_client_protos:.proto=.pb.go)
mixer_config_client_pb_pythons := $(patsubst $(mixer_config_client_path)/%.proto,$(python_output_path)/$(mixer_client_config_path)/%_pb2.py,$(mixer_client_config_protos))
mixer_config_client_pb_doc := $(mixer_config_client_path)/istio.mixer.v1.config.client.pb.html
mixer_config_client_openapi := $(mixer_config_client_path)/istio.mixer.v1.config.client.json
mixer_config_client_k8s_gos := \
	$(patsubst $(mixer_config_client_path)/%.proto,$(mixer_config_client_path)/%_json.gen.go,$(shell grep -l "^ *oneof " $(mixer_config_client_protos))) \
	$(patsubst $(mixer_config_client_path)/%.proto,$(mixer_config_client_path)/%_deepcopy.gen.go,$(shell grep -l "+kubetype-gen" $(mixer_config_client_protos)))

mixer_adapter_model_v1beta1_path := mixer/adapter/model/v1beta1
mixer_adapter_model_v1beta1_protos := $(wildcard $(mixer_adapter_model_v1beta1_path)/*.proto)
mixer_adapter_model_v1beta1_pb_gos := $(mixer_adapter_model_v1beta1_protos:.proto=.pb.go)
mixer_adapter_model_v1beta1_pb_pythons := $(patsubst $(mixer_adapter_model_v1beta1_path)/%.proto,$(python_output_path)/$(mixer_client_config_path)/%_pb2.py,$(mixer_client_config_protos))
mixer_adapter_model_v1beta1_openapi := $(mixer_adapter_model_v1beta1_path)/istio.mixer.adapter.model.v1beta1.json

$(mixer_v1_pb_gos) $(mixer_v1_pb_pythons) $(mixer_v1_k8s_gos): $(mixer_v1_protos)
	@$(protolock) status
	@$(protoc) $(gogoslick_plugin) $(protoc_gen_k8s_support_plugins) $(protoc_gen_python_plugin) $^
	@cp -r /tmp/istio.io/api/mixer/* mixer

$(mixer_config_client_pb_gos) $(mixer_config_client_pb_doc) $(mixer_config_client_pb_pythons) $(mixer_config_client_k8s_gos): $(mixer_config_client_protos)
	@$(protolock) status
	@$(protoc) $(gogoslick_plugin) $(protoc_gen_k8s_support_plugins) $(protoc_gen_docs_plugin)$(mixer_config_client_path) $(protoc_gen_python_plugin) $^
	@cp -r /tmp/istio.io/api/mixer/* mixer

$(mixer_adapter_model_v1beta1_pb_gos) $(mixer_adapter_model_v1beta1_pb_pythons): $(mixer_adapter_model_v1beta1_protos)
	@$(protolock) status
	@$(protoc) $(gogoslick_plugin) $(protoc_gen_python_plugin)  $^
	@cp -r /tmp/istio.io/api/mixer/* mixer

generate-mixer: \
	$(mixer_v1_pb_gos) $(mixer_v1_pb_pythons) $(mixer_v1_k8s_gos) \
	$(mixer_config_client_pb_gos) $(mixer_config_client_pb_doc) $(mixer_config_client_pb_pythons) $(mixer_config_client_k8s_gos) \
	$(mixer_adapter_model_v1beta1_pb_gos) $(mixer_adapter_model_v1beta1_pb_pythons)

clean-mixer:
	@rm -fr $(mixer_v1_pb_gos) $(mixer_v1_pb_pythons) $(mixer_v1_k8s_gos)
	@rm -fr $(mixer_config_client_pb_gos) $(mixer_config_client_pb_doc) $(mixer_config_client_pb_pythons) $(mixer_config_client_k8s_gos)
	@rm -fr $(mixer_adapter_model_v1beta1_pb_gos) $(mixer_adapter_model_v1beta1_pb_pythons)

#####################
# networking/...
#####################

networking_v1alpha3_path := networking/v1alpha3
networking_v1alpha3_protos := $(wildcard $(networking_v1alpha3_path)/*.proto)
networking_v1alpha3_pb_gos := $(networking_v1alpha3_protos:.proto=.pb.go)
networking_v1alpha3_pb_pythons := $(patsubst $(networking_v1alpha3_path)/%.proto,$(python_output_path)/$(networking_v1alpha3_path)/%_pb2.py,$(networking_v1alpha3_protos))
networking_v1alpha3_pb_docs := $(networking_v1alpha3_protos:.proto=.pb.html)
networking_v1alpha3_openapi := $(networking_v1alpha3_protos:.proto=.json)
networking_v1alpha3_k8s_gos := \
	$(patsubst $(networking_v1alpha3_path)/%.proto,$(networking_v1alpha3_path)/%_json.gen.go,$(shell grep -l "^ *oneof " $(networking_v1alpha3_protos))) \
	$(patsubst $(networking_v1alpha3_path)/%.proto,$(networking_v1alpha3_path)/%_deepcopy.gen.go,$(shell grep -l "+kubetype-gen" $(networking_v1alpha3_protos)))

$(networking_v1alpha3_pb_gos) $(networking_v1alpha3_pb_docs) $(networking_v1alpha3_pb_pythons) $(networking_v1alpha3_k8s_gos): $(networking_v1alpha3_protos)
	@$(protolock) status
	@$(protoc) $(gogofast_plugin) $(protoc_gen_k8s_support_plugins) $(protoc_gen_docs_plugin_for_networking)$(networking_v1alpha3_path) $(protoc_gen_python_plugin) $^
	@cp -r /tmp/istio.io/api/networking/* networking

generate-networking: $(networking_v1alpha3_pb_gos) $(networking_v1alpha3_pb_docs) $(networking_v1alpha3_pb_pythons) $(networking_v1alpha3_k8s_gos)

clean-networking:
	@rm -fr $(networking_v1alpha3_pb_gos) $(networking_v1alpha3_pb_docs) $(networking_v1alpha3_pb_pythons) $(networking_v1alpha3_k8s_gos)

#####################
# rbac/...
#####################

rbac_v1alpha1_path := rbac/v1alpha1
rbac_v1alpha1_protos := $(wildcard $(rbac_v1alpha1_path)/*.proto)
rbac_v1alpha1_pb_gos := $(rbac_v1alpha1_protos:.proto=.pb.go)
rbac_v1alpha1_pb_pythons := $(patsubst $(rbac_v1alpha1_path)/%.proto,$(python_output_path)/$(rbac_v1alpha1_path)/%_pb2.py,$(rbac_v1alpha1_protos))
rbac_v1alpha1_pb_doc := $(rbac_v1alpha1_path)/istio.rbac.v1alpha1.pb.html
rbac_v1alpha1_openapi := $(rbac_v1alpha1_path)/istio.rbac.v1alpha1.json
rbac_v1alpha1_k8s_gos := \
	$(patsubst $(rbac_v1alpha1_path)/%.proto,$(rbac_v1alpha1_path)/%_json.gen.go,$(shell grep -l "^ *oneof " $(rbac_v1alpha1_protos))) \
	$(patsubst $(rbac_v1alpha1_path)/%.proto,$(rbac_v1alpha1_path)/%_deepcopy.gen.go,$(shell grep -l "+kubetype-gen" $(rbac_v1alpha1_protos)))

$(rbac_v1alpha1_pb_gos) $(rbac_v1alpha1_pb_doc) $(rbac_v1alpha1_pb_pythons) $(rbac_v1alpha1_k8s_gos): $(rbac_v1alpha1_protos)
	@$(protolock) status
	@$(protoc) $(gogofast_plugin) $(protoc_gen_k8s_support_plugins) $(protoc_gen_docs_plugin)$(rbac_v1alpha1_path) $(protoc_gen_python_plugin) $^
	@cp -r /tmp/istio.io/api/rbac/* rbac

generate-rbac: $(rbac_v1alpha1_pb_gos) $(rbac_v1alpha1_pb_doc) $(rbac_v1alpha1_protos) $(rbac_v1alpha1_k8s_gos)

clean-rbac:
	@rm -fr $(rbac_v1alpha1_pb_gos) $(rbac_v1alpha1_pb_doc) $(rbac_v1alpha1_pb_pythons) $(rbac_v1alpha1_k8s_gos)

#####################
# authentication/...
#####################

authn_v1alpha1_path := authentication/v1alpha1
authn_v1alpha1_protos := $(wildcard $(authn_v1alpha1_path)/*.proto)
authn_v1alpha1_pb_gos := $(authn_v1alpha1_protos:.proto=.pb.go)
authn_v1alpha1_pb_pythons := $(patsubst $(authn_v1alpha1_path)/%.proto,$(python_output_path)/$(authn_v1alpha1_path)/%_pb2.py,$(authn_v1alpha1_protos))
authn_v1alpha1_pb_doc := $(authn_v1alpha1_path)/istio.authentication.v1alpha1.pb.html
authn_v1alpha1_openapi := $(authn_v1alpha1_path)/istio.authentication.v1alpha1.json
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

security_v1beta1_path := security/v1beta1
security_v1beta1_protos := $(wildcard $(security_v1beta1_path)/*.proto)
security_v1beta1_pb_gos := $(security_v1beta1_protos:.proto=.pb.go)
security_v1beta1_pb_pythons := $(patsubst $(security_v1beta1_path)/%.proto,$(python_output_path)/$(security_v1beta1_path)/%_pb2.py,$(security_v1beta1_protos))
security_v1beta1_pb_doc := $(security_v1beta1_path)/istio.security.v1beta1.pb.html
security_v1beta1_openapi := $(security_v1beta1_protos:.proto=.json)
security_v1beta1_k8s_gos := \
	$(patsubst $(security_v1beta1_path)/%.proto,$(security_v1beta1_path)/%_json.gen.go,$(shell grep -l "^ *oneof " $(security_v1beta1_protos))) \
	$(patsubst $(security_v1beta1_path)/%.proto,$(security_v1beta1_path)/%_deepcopy.gen.go,$(shell grep -l "+kubetype-gen" $(security_v1beta1_protos)))


$(security_v1beta1_pb_gos) $(security_v1beta1_pb_doc) $(security_v1beta1_pb_pythons) $(security_v1beta1_k8s_gos): $(security_v1beta1_protos)
	@$(protolock) status
	@$(protoc) $(gogofast_plugin) $(protoc_gen_k8s_support_plugins) $(protoc_gen_docs_plugin)$(security_v1beta1_path) $(protoc_gen_python_plugin) $^
	@cp -r /tmp/istio.io/api/security/* security

generate-security: $(security_v1beta1_pb_gos) $(security_v1beta1_pb_doc) $(security_v1beta1_pb_pythons) $(security_v1beta1_k8s_gos)

clean-security:
	@rm -fr $(security_v1beta1_pb_gos) $(security_v1beta1_pb_docs) $(security_v1beta1_pb_pythons) $(security_v1beta1_k8s_gos)

#####################
# envoy/...
#####################

envoy_path := envoy
envoy_protos := $(shell find $(envoy_path) -type f -name '*.proto' | sort)
envoy_pb_gos := $(envoy_protos:.proto=.pb.go)
envoy_pb_pythons := $(patsubst $(envoy_path)/%.proto,$(python_output_path)/$(envoy_path)/%_pb2.py,$(envoy_protos))

$(envoy_pb_gos): %.pb.go : %.proto
	@$(protolock) status
	@$(protoc) $(gogofast_plugin) $<
	@cp -r /tmp/istio.io/api/envoy/* envoy

$(envoy_pb_pythons): $(envoy_protos)
	@$(protolock) status
	@$(protoc) $(protoc_gen_python_plugin) $^

generate-envoy: $(envoy_pb_gos) $(envoy_pb_pythons)

clean-envoy:
	@rm -fr $(envoy_pb_gos) $(envoy_pb_pythons)

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
	$(policy_v1beta1_protos) \
	$(mixer_v1_protos) \
	$(mixer_config_client_protos) \
	$(mixer_adapter_model_v1beta1_protos) \
	$(networking_v1alpha3_protos) \
	$(rbac_v1alpha1_protos) \
	$(authn_v1alpha1_protos) \
	$(security_v1beta1_protos) \
	$(type_v1beta1_protos)

all_openapi := \
	$(core_v1alpha1_openapi) \
	$(mcp_v1alpha1_openapi) \
	$(mesh_v1alpha1_openapi) \
	$(policy_v1beta1_openapi) \
	$(mixer_v1_openapi) \
	$(mixer_config_client_openapi) \
	$(mixer_adapter_model_v1beta1_openapi) \
	$(networking_v1alpha3_openapi) \
	$(rbac_v1alpha1_openapi) \
	$(authn_v1alpha1_openapi) \
	$(security_v1beta1_openapi) \
	$(type_v1beta1_openapi)

$(all_openapi): $(all_protos)
	@$(cue) -f=$(repo_dir)/cue.yaml

generate-openapi-schema: $(all_openapi)

clean-openapi-schema:
	@rm -fr $(all_openapi)

#####################
# Cleanup
#####################

clean: \
	clean-core \
	clean-mcp \
	clean-mesh \
	clean-mixer \
	clean-networking \
	clean-rbac \
	clean-authn \
	clean-envoy \
	clean-policy \
	clean-annotations \
	clean-openapi-schema \
	clean-security \
	clean-type

#####################
# CI System
#####################

presubmit: proto-commit lint release-lock-status
postsubmit: presubmit

#####################
# Common definitions
#####################

include common/Makefile.common.mk
