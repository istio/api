all: generate

########################
# docker_gen
########################

# Use a different generation mechanism when running from the
# image itself
protoc = protoc -I/protobuf -I/src -I.
protolock = protolock --ignore vendor
prototool = prototool
kubetype_gen = kubetype-gen
deepcopy_gen = deepcopy-gen
client_gen = client-gen
lister_gen = lister-gen
informer_gen = informer-gen
go_to_protobuf = go-to-protobuf

ifdef CIRCLECI
docker_run =
out_path = $(OUT_PATH)
else
all_img := gcr.io/istio-testing/api-build-tools:2019-02-04
pwd := $(shell pwd)
repo_mount := /src/istio.io/api
docker_run ?= docker run --rm -v $(pwd):$(repo_mount) -w $(repo_mount) $(all_img)
out_path = /src
endif

########################
# protoc_gen_gogo*
########################

gogo_plugin_prefix := --gogo_out=plugins=grpc,
gogofast_plugin_prefix := --gogofast_out=plugins=grpc,
gogoslick_plugin_prefix := --gogoslick_out=plugins=grpc,

########################
# protoc_gen_python
########################

protoc_gen_python_prefix := --python_out=,
protoc_gen_python_plugin := $(protoc_gen_python_prefix):./python/istio_api

comma := ,
empty:=
space := $(empty) $(empty)

importmaps := \
	gogoproto/gogo.proto=github.com/gogo/protobuf/gogoproto \
	google/protobuf/any.proto=github.com/gogo/protobuf/types \
	google/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor \
	google/protobuf/duration.proto=github.com/gogo/protobuf/types \
	google/protobuf/struct.proto=github.com/gogo/protobuf/types \
	google/protobuf/timestamp.proto=github.com/gogo/protobuf/types \
	google/protobuf/wrappers.proto=github.com/gogo/protobuf/types \
	google/rpc/status.proto=github.com/gogo/googleapis/google/rpc \
	google/rpc/code.proto=github.com/gogo/googleapis/google/rpc \
	google/rpc/error_details.proto=github.com/gogo/googleapis/google/rpc \

# generate mapping directive with M<proto>:<go pkg>, format for each proto file
mapping_with_spaces := $(foreach map,$(importmaps),M$(map),)
gogo_mapping := $(subst $(space),$(empty),$(mapping_with_spaces))

gogo_plugin := $(gogo_plugin_prefix)$(gogo_mapping):$(out_path)
gogofast_plugin := $(gogofast_plugin_prefix)$(gogo_mapping):$(out_path)
gogoslick_plugin := $(gogoslick_plugin_prefix)$(gogo_mapping):$(out_path)

########################
# protoc_gen_docs
########################

protoc_gen_docs_plugin := --docs_out=warnings=true,mode=html_fragment_with_front_matter:./

#####################
# Generation Rules
#####################

generate: \
	generate-mcp-go \
	generate-mcp-python \
	generate-mesh-go \
	generate-mesh-python \
	generate-mixer-go \
	generate-mixer-python \
	generate-routing-go \
	generate-routing-python \
	generate-rbac-go \
	generate-rbac-python \
	generate-authn-go \
	generate-authn-python \
	generate-envoy-go \
	generate-envoy-python

#####################
# mcp/...
#####################

config_mcp_path := mcp/v1alpha1
config_mcp_protos := $(shell find $(config_mcp_path) -type f -name '*.proto' ! -name 'generated.proto' | sort)
config_mcp_pb_gos := $(config_mcp_protos:.proto=.pb.go)
config_mcp_pb_pythons := $(config_mcp_protos:.proto=_pb2.py)
config_mcp_pb_doc := $(config_mcp_path)/istio.mcp.v1alpha1.pb.html

generate-mcp-go: $(config_mcp_pb_gos) $(config_mcp_pb_doc)

$(config_mcp_pb_gos) $(config_mcp_pb_doc): $(config_mcp_protos)
	## Generate mcp/v1alpha1/*.pb.go + $(config_mcp_pb_doc)
	@$(docker_run) $(protolock) status
	@$(docker_run) $(protoc) $(gogofast_plugin) $(protoc_gen_docs_plugin)$(config_mcp_path) $^

generate-mcp-python: $(config_mcp_pb_pythons)

$(config_mcp_pb_pythons): $(config_mcp_protos)
	## Generate python/istio_api/mcp/v1alpha1/*_pb2.py
	@$(docker_run) $(protolock) status
	@$(docker_run) $(protoc) $(protoc_gen_python_plugin) $^

clean-mcp:
	rm -f $(config_mcp_pb_gos)
	rm -f $(config_mcp_pb_doc)

#####################
# mesh/...
#####################

mesh_path := mesh/v1alpha1
mesh_protos := $(shell find $(mesh_path) -type f -name '*.proto' ! -name 'generated.proto' | sort)
mesh_pb_gos := $(mesh_protos:.proto=.pb.go)
mesh_pb_pythons := $(mesh_protos:.proto=_pb2.py)
mesh_pb_doc := $(mesh_path)/istio.mesh.v1alpha1.pb.html

generate-mesh-go: $(mesh_pb_gos) $(mesh_pb_doc)

$(mesh_pb_gos) $(mesh_pb_doc): $(mesh_protos)
	## Generate mesh/v1alpha1/*.pb.go + $(mesh_pb_doc)
	@$(docker_run) $(protolock) status
	@$(docker_run) $(protoc) $(gogofast_plugin) $(protoc_gen_docs_plugin)$(mesh_path) $^

generate-mesh-python: $(mesh_pb_pythons)

$(mesh_pb_pythons): $(mesh_protos)
	## Generate python/istio_api/mesh/v1alpha1/*_pb2.py
	@$(docker_run) $(protolock) status
	@$(docker_run) $(protoc) $(protoc_gen_python_plugin) $^

clean-mesh:
	rm -f $(mesh_pb_gos)
	rm -f $(mesh_pb_doc)

#####################
# mixer/...
#####################

mixer_v1_path := mixer/v1
mixer_v1_protos :=  $(shell find $(mixer_v1_path) -maxdepth 1 -type f -name '*.proto' ! -name 'generated.proto' | sort)
mixer_v1_pb_gos := $(mixer_v1_protos:.proto=.pb.go)
mixer_v1_pb_pythons := $(mixer_v1_protos:.proto=_pb2.py)
mixer_v1_pb_doc := $(mixer_v1_path)/istio.mixer.v1.pb.html

mixer_config_client_path := mixer/v1/config/client
mixer_config_client_protos := $(shell find $(mixer_config_client_path) -maxdepth 1 -type f -name '*.proto' ! -name 'generated.proto' | sort)
mixer_config_client_pb_gos := $(mixer_config_client_protos:.proto=.pb.go)
mixer_config_client_pb_pythons := $(mixer_config_client_protos:.proto=_pb2.py)
mixer_config_client_pb_doc := $(mixer_config_client_path)/istio.mixer.v1.config.client.pb.html

mixer_adapter_model_v1beta1_path := mixer/adapter/model/v1beta1
mixer_adapter_model_v1beta1_protos := $(shell find $(mixer_adapter_model_v1beta1_path) -maxdepth 1 -type f -name '*.proto' ! -name 'generated.proto' | sort)
mixer_adapter_model_v1beta1_pb_gos := $(mixer_adapter_model_v1beta1_protos:.proto=.pb.go)
mixer_adapter_model_v1beta1_pb_pythons := $(mixer_adapter_model_v1beta1_protos:.proto=_pb2.py)
mixer_adapter_model_v1beta1_pb_doc := $(mixer_adapter_model_v1beta1_path)/istio.mixer.adapter.model.v1beta1.pb.html

policy_v1beta1_path := policy/v1beta1
policy_v1beta1_protos := $(shell find $(policy_v1beta1_path) -maxdepth 1 -type f -name '*.proto' ! -name 'generated.proto' | sort)
policy_v1beta1_pb_gos := $(policy_v1beta1_protos:.proto=.pb.go)
policy_v1beta1_pb_pythons := $(policy_v1beta1_protos:.proto=_pb2.py)
policy_v1beta1_pb_doc := $(policy_v1beta1_path)/istio.policy.v1beta1.pb.html

generate-mixer-go: \
	$(mixer_v1_pb_gos) $(mixer_v1_pb_doc) \
	$(mixer_config_client_pb_gos) $(mixer_config_client_pb_doc) \
	$(mixer_adapter_model_v1beta1_pb_gos) $(mixer_adapter_model_v1beta1_pb_doc) \
	$(policy_v1beta1_pb_gos) $(policy_v1beta1_pb_doc)

$(mixer_v1_pb_gos) $(mixer_v1_pb_doc): $(mixer_v1_protos)
	## Generate mixer/v1/*.pb.go + $(mixer_v1_pb_doc)
	@$(docker_run) $(protolock) status
	@$(docker_run) $(protoc) $(gogoslick_plugin) $(protoc_gen_docs_plugin)$(mixer_v1_path) $^

$(mixer_config_client_pb_gos) $(mixer_config_client_pb_doc): $(mixer_config_client_protos)
	## Generate mixer/v1/config/client/*.pb.go + $(mixer_config_client_pb_doc)
	@$(docker_run) $(protolock) status
	@$(docker_run) $(protoc) $(gogoslick_plugin) $(protoc_gen_docs_plugin)$(mixer_config_client_path) $^

$(mixer_adapter_model_v1beta1_pb_gos) $(mixer_adapter_model_v1beta1_pb_doc) : $(mixer_adapter_model_v1beta1_protos)
	## Generate mixer/adapter/model/v1beta1/*.pb.go + $(mixer_adapter_model_v1beta1_pb_doc)
	@$(docker_run) $(protolock) status
	@$(docker_run) $(protoc) $(gogoslick_plugin) $(protoc_gen_docs_plugin)$(mixer_adapter_model_v1beta1_path) $^

$(policy_v1beta1_pb_gos) $(policy_v1beta1_pb_doc) : $(policy_v1beta1_protos)
	## Generate policy/v1beta1/*.pb.go + $(policy_v1beta1_pb_doc)
	@$(docker_run) $(protolock) status
	@$(docker_run) $(protoc) $(gogoslick_plugin) $(protoc_gen_docs_plugin)$(policy_v1beta1_path) $^

generate-mixer-python: \
	$(mixer_v1_pb_pythons) \
	$(mixer_config_client_pb_pythons) \
	$(mixer_adapter_model_v1beta1_pb_pythons) \
	$(policy_v1beta1_pb_pythons)

$(mixer_v1_pb_pythons): $(mixer_v1_protos)
	## Generate python/istio_api/mixer/v1/*_pb2.py
	@$(docker_run) $(protolock) status
	@$(docker_run) $(protoc) $(protoc_gen_python_plugin) $^

$(mixer_config_client_pb_pythons): $(mixer_config_client_protos)
	## Generate python/istio_api/mixer/v1/config/client/*_pb2.py
	@$(docker_run) $(protolock) status
	@$(docker_run) $(protoc) $(protoc_gen_python_plugin) $^

$(mixer_adapter_model_v1beta1_pb_pythons): $(mixer_adapter_model_v1beta1_protos)
	## Generate python/istio_api/mixer/adapter/model/v1beta1/*_pb2.py
	@$(docker_run) $(protolock) status
	@$(docker_run) $(protoc) $(protoc_gen_python_plugin) $^

$(policy_v1beta1_pb_pythons): $(policy_v1beta1_protos)
	## Generate python/istio_api/policy/v1beta1/*_pb2.py
	@$(docker_run) $(protolock) status
	@$(docker_run) $(protoc) $(protoc_gen_python_plugin) $^

clean-mixer:
	rm -f $(mixer_v1_pb_gos) $(mixer_config_client_pb_gos) $(mixer_adapter_model_v1beta1_pb_gos) $(policy_v1beta1_pb_gos) policy/v1beta1/fixed_cfg.pb.go
	rm -f $(mixer_v1_pb_doc) $(mixer_config_client_pb_doc) $(mixer_adapter_model_v1beta1_pb_doc) $(policy_v1beta1_pb_doc)

#####################
# routing/...
#####################

routing_v1alpha3_path := networking/v1alpha3
routing_v1alpha3_protos := $(shell find networking/v1alpha3 -type f -name '*.proto' ! -name 'generated.proto' | sort)
routing_v1alpha3_pb_gos := $(routing_v1alpha3_protos:.proto=.pb.go)
routing_v1alpha3_pb_pythons := $(routing_v1alpha3_protos:.proto=_pb2.py)
routing_v1alpha3_pb_doc := $(routing_v1alpha3_path)/istio.routing.v1alpha3.pb.html

generate-routing-go: $(routing_v1alpha3_pb_gos) $(routing_v1alpha3_pb_doc)

$(routing_v1alpha3_pb_gos) $(routing_v1alpha3_pb_doc): $(routing_v1alpha3_protos)
	## Generate networking/v1alpha3/*.pb.go
	@$(docker_run) $(protolock) status
	@$(docker_run) $(protoc) $(gogofast_plugin) $(protoc_gen_docs_plugin)$(routing_v1alpha3_path) $^

generate-routing-python: $(routing_v1alpha3_pb_pythons)

$(routing_v1alpha3_pb_pythons): $(routing_v1alpha3_protos)
	## Generate python/istio_api/networking/v1alpha3/*_pb2.py
	@$(docker_run) $(protolock) status
	@$(docker_run) $(protoc) $(protoc_gen_python_plugin) $^

clean-routing:
	rm -f $(routing_v1alpha3_pb_gos)
	rm -f $(routing_v1alpha3_pb_doc)

#####################
# rbac/...
#####################

rbac_v1alpha1_path := rbac/v1alpha1
rbac_v1alpha1_protos := $(shell find $(rbac_v1alpha1_path) -type f -name '*.proto' ! -name 'generated.proto' | sort)
rbac_v1alpha1_pb_gos := $(rbac_v1alpha1_protos:.proto=.pb.go)
rbac_v1alpha1_pb_pythons := $(rbac_v1alpha1_protos:.proto=_pb2.py)
rbac_v1alpha1_pb_doc := $(rbac_v1alpha1_path)/istio.rbac.v1alpha1.pb.html

generate-rbac-go: $(rbac_v1alpha1_pb_gos) $(rbac_v1alpha1_pb_doc)

$(rbac_v1alpha1_pb_gos) $(rbac_v1alpha1_pb_doc): $(rbac_v1alpha1_protos)
	## Generate rbac/v1alpha1/*.pb.go
	@$(docker_run) $(protolock) status
	@$(docker_run) $(protoc) $(gogofast_plugin) $(protoc_gen_docs_plugin)$(rbac_v1alpha1_path) $^

generate-rbac-python: $(rbac_v1alpha1_protos)

$(rbac_v1alpha1_pb_pythons): $(rbac_v1alpha1_protos)
	## Generate python/istio_api/rbac/v1alpha1/*_pb2.py
	@$(docker_run) $(protolock) status
	@$(docker_run) $(protoc) $(protoc_gen_python_plugin) $^

clean-rbac:
	rm -f $(rbac_v1alpha1_pb_gos)
	rm -f $(rbac_v1alpha1_pb_doc)


#####################
# authentication/...
#####################

authn_v1alpha1_path := authentication/v1alpha1
authn_v1alpha1_protos := $(shell find $(authn_v1alpha1_path) -type f -name '*.proto' ! -name 'generated.proto' | sort)
authn_v1alpha1_pb_gos := $(authn_v1alpha1_protos:.proto=.pb.go)
authn_v1alpha1_pb_pythons := $(authn_v1alpha1_protos:.proto=_pb2.py)
authn_v1alpha1_pb_doc := $(authn_v1alpha1_path)/istio.authentication.v1alpha1.pb.html

generate-authn-go: $(authn_v1alpha1_pb_gos) $(authn_v1alpha1_pb_doc)

$(authn_v1alpha1_pb_gos) $(authn_v1alpha1_pb_doc): $(authn_v1alpha1_protos)
	## Generate authentication/v1alpha1/*.pb.go
	@$(docker_run) $(protolock) status
	@$(docker_run) $(protoc) $(gogofast_plugin) $(protoc_gen_docs_plugin)$(authn_v1alpha1_path) $^

generate-authn-python: $(authn_v1alpha1_pb_pythons)

$(authn_v1alpha1_pb_pythons): $(authn_v1alpha1_protos)
	## Generate python/istio_api/authentication/v1alpha1/*_pb2.py
	@$(docker_run) $(protolock) status
	@$(docker_run) $(protoc) $(protoc_gen_python_plugin) $^

clean-authn:
	rm -f $(authn_v1alpha1_pb_gos)
	rm -f $(authn_v1alpha1_pb_doc)

#####################
# envoy/...
#####################

envoy_path := envoy
envoy_protos := $(shell find $(envoy_path) -type f -name '*.proto' ! -name 'generated.proto' | sort)
envoy_pb_gos := $(envoy_protos:.proto=.pb.go)
envoy_pb_pythons := $(envoy_protos:.proto=_pb2.py)

generate-envoy-go: $(envoy_pb_gos) $(envoy_pb_doc)

# Envoy APIs is internal APIs, documents is not required.
$(envoy_pb_gos): $(envoy_protos)
	## Generate envoy/*/*.pb.go
	@$(docker_run) $(protolock) status
	@$(docker_run) $(protoc) $(gogofast_plugin) $^

generate-envoy-python: $(envoy_pb_pythons)

# Envoy APIs is internal APIs, documents is not required.
$(envoy_pb_pythons): $(envoy_protos)
	## Generate envoy/*/*_pb2.py
	@$(docker_run) $(protolock) status
	@$(docker_run) $(protoc) $(protoc_gen_python_plugin) $^

clean-envoy:
	rm -f $(envoy_pb_gos)

#####################
# Kube APIs
#####################
src_dir := $(abspath $(dir $(firstword $(MAKEFILE_LIST))))
ifeq "$(shell echo ${src_dir} | grep '/istio.io/api$$')" ""
	$(error Source code must be located in .../istio.io/api directory.  Go package paths will not resolve correctly.)
endif

# define this to run code generation locally on your dev machine
# you'll need protoc in your PATH
ifdef LOCAL_KUBE_BUILD
kube_run =
install_gogo_protoc_gen = install-gogo-protoc-gen
install_k8s_code_generators = install-k8s-code-generators
go_build_generated = go build istio.io/api/pkg/kube/...
else
kube_run = $(docker_run)
go_build_generated =
endif

# kube code generator parameter values
kube_istio_source_packages = $(subst $(space),$(empty), \
	istio.io/api/authentication/..., \
	istio.io/api/mixer/..., \
	istio.io/api/networking/..., \
	istio.io/api/policy/..., \
	istio.io/api/rbac/... \
	)
kube_base_output_package = istio.io/api/pkg/kube
kube_api_base_package = $(kube_base_output_package)/apis
kube_api_packages = $(subst $(space),$(empty), \
	$(kube_api_base_package)/authentication/v1alpha1, \
	$(kube_api_base_package)/config/v1alpha2, \
	$(kube_api_base_package)/networking/v1alpha3, \
	$(kube_api_base_package)/rbac/v1alpha1 \
	)
kube_clientset_package = $(kube_base_output_package)/clientset
kube_clientset_name = versioned
kube_listers_package = $(kube_base_output_package)/listers
kube_informers_package = $(kube_base_output_package)/informers
kube_go_header_text = vendor/istio.io/tools/cmd/kubetype-gen/boilerplate.go.txt

# go-to-protobuf
# packages to scan for types that need protobuf un/marshalers
kube_proto_source_packages = $(kube_api_packages)
# packages to scan for dependencies.  generator looks for generated.proto files in these packages
# leading hyphen means include, but don't generate protos for those packages
# format is go-package=proto-package, default proto package is go package with path separator replaced with "."
kube_proto_api_packages = $(subst $(space),$(empty), \
	-istio.io/api/authentication/v1alpha1=istio.authentication.v1alpha1, \
	-istio.io/api/mixer/v1/config/client=istio.mixer.v1.config.client, \
	-istio.io/api/networking/v1alpha3=istio.networking.v1alpha3, \
	-istio.io/api/policy/v1beta1=istio.policy.v1beta1, \
	-istio.io/api/rbac/v1alpha1=istio.rbac.v1alpha1, \
	-k8s.io/apimachinery/pkg/apis/meta/v1, \
	-k8s.io/apimachinery/pkg/runtime/schema \
	)
# istio proto files include relative paths for google/rpc/*.proto and google/protobuf/*.proto
# we just need these to resolve so the istio .proto files load correctly.
# all the generated code should use the already generated .pb.go files from istio.
kube_proto_imports = $(subst $(space),$(empty), \
	./vendor, \
	./vendor/github.com/gogo/googleapis, \
	./vendor/github.com/gogo/protobuf/protobuf, \
	)

generate-kube-apis: generate-kube-apis-go generate-kube-apis-proto
	# verify that the generated files build with go
	$(go_build_generated)

generate-kube-apis-proto: $(install_gogo_protoc_gen)
	# generate proto files for kube api types
	@$(kube_run) $(go_to_protobuf) --packages $(kube_proto_source_packages) --apimachinery-packages $(kube_proto_api_packages) --proto-import $(kube_proto_imports) -h $(kube_go_header_text)

generate-kube-apis-go: $(install_k8s_code_generators)
	# generate kube api type wrappers for istio types
	@$(kube_run) $(kubetype_gen) --input-dirs $(kube_istio_source_packages) --output-package $(kube_api_base_package) -h $(kube_go_header_text)
	# generate deepcopy for istio types
	@$(kube_run) $(deepcopy_gen) --input-dirs $(kube_istio_source_packages) -O zz_generated.deepcopy  -h $(kube_go_header_text)
	# generate deepcopy for kube api types
	@$(kube_run) $(deepcopy_gen) --input-dirs $(kube_api_packages) -O zz_generated.deepcopy  -h $(kube_go_header_text)
	# generate clientsets for kube api types
	@$(kube_run) $(client_gen) --clientset-name $(kube_clientset_name) --input-base "" --input  $(kube_api_packages) --output-package $(kube_clientset_package) -h $(kube_go_header_text)
	# generate listers for kube api types
	@$(kube_run) $(lister_gen) --input-dirs $(kube_api_packages) --output-package $(kube_listers_package) -h $(kube_go_header_text)
	# generate informers for kube api types
	@$(kube_run) $(informer_gen) --input-dirs $(kube_api_packages) --versioned-clientset-package $(kube_clientset_package)/$(kube_clientset_name) --listers-package $(kube_listers_package) --output-package $(kube_informers_package) -h $(kube_go_header_text)

install-k8s-code-generators:
	# build kube api code generators
	@(cd vendor/k8s.io/code-generator && go install ./cmd/{defaulter-gen,client-gen,lister-gen,informer-gen,deepcopy-gen,go-to-protobuf})
	@(cd vendor/istio.io/tools/cmd && go install ./kubetype-gen)

install-gogo-protoc-gen:
	# build go-to-protobuf generator
	@(cd vendor/github.com/golang/protobuf && go install ./protoc-gen-go)
	@(cd vendor/github.com/gogo/protobuf && go install ./{protoc-gen-gofast,protoc-gen-gogo,protoc-gen-gogofast,protoc-gen-gogofaster,protoc-gen-gogoslick})

clean-kube-apis:
	# delete generated kube api files
	rm -rf pkg/kube/
	# delete generated deepcopy files for istio types
	find authentication mixer networking policy rbac -name zz_generated.deepcopy.go -delete

#####################
# Protolock
#####################

proto-commit:
	@$(docker_run) $(protolock) commit

proto-commit-force:
	@$(docker_run) $(protolock) commit --force

#####################
# Lint
#####################

lint:
	@$(docker_run) $(prototool) lint --protoc-bin-path=/usr/bin/protoc --protoc-wkt-path=/protobuf

#####################
# Cleanup
#####################

clean-python:
	rm -rf python/istio_api/*

clean: 	clean-mcp \
	clean-mesh \
	clean-mixer \
	clean-routing \
	clean-rbac \
	clean-authn \
	clean-envoy \
	clean-python
