# GOOGLEAPIS 
GOOGLEAPIS_SHA = c8c975543a134177cc41b64cbbf10b88fe66aa1d
GOOGLEAPIS_URL = https://raw.githubusercontent.com/googleapis/googleapis/$(GOOGLEAPIS_SHA)

# BUILD UP PROTOC ARGs
PROTO_PATH = --proto_path=protoc-tmp --proto_path=. --proto_path=vendor/github.com/gogo/protobuf

RPC_PATH = google/rpc
STATUS_PROTO = protoc-tmp/$(RPC_PATH)/status.proto
CODE_PROTO = protoc-tmp/$(RPC_PATH)/code.proto
ERR_PROTO = protoc-tmp/$(RPC_PATH)/error_details.proto

PROTOC = protoc-min-version -version=3.5.0

GOGOSLICK_PLUGIN_PREFIX = --gogoslick_out=plugins=grpc=
GOGO_PLUGIN_PREFIX = --gogo_out=plugins=grpc=
PLUGIN_SUFFIX = :.

# BASIC STANDARD MAPPINGS
MAPPING := Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,
MAPPING := $(MAPPING)Mgogoproto/gogo.proto=github.com/gogo/protobuf/gogoproto,
MAPPING := $(MAPPING)Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor,
MAPPING := $(MAPPING)Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,
MAPPING := $(MAPPING)Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,
MAPPING := $(MAPPING)Mgoogle/rpc/status.proto=istio.io/gogo-genproto/googleapis/google/rpc,
MAPPING := $(MAPPING)Mgoogle/rpc/code.proto=istio.io/gogo-genproto/googleapis/google/rpc,
MAPPING := $(MAPPING)Mgoogle/rpc/error_details.proto=istio.io/gogo-genproto/googleapis/google/rpc,

GOGOSLICK_PLUGIN = $(GOGOSLICK_PLUGIN_PREFIX)$(MAPPING)$(PLUGIN_SUFFIX)

# MIXER STUFF

## MIXER-PROTO SPECIFIC MAPPINGS
MIXER_MAPPING := $(MAPPING)Mmixer/v1/attributes.proto=istio.io/api/mixer/v1,Mmixer/v1/check.proto=istio.io/api/mixer/v1,Mmixer/v1/report.proto=istio.io/api/mixer/v1,Mmixer/v1/service.proto=istio.io/api/mixer/v1,
MIXER_MAPPING := $(MIXER_MAPPING)Mmixer/v1/config/client/api_spec.proto=istio.io/api/mixer/v1/config/client,Mmixer/v1/config/client/auth.proto=istio.io/api/mixer/v1/config/client,Mmixer/v1/config/client/client_config.proto=istio.io/api/mixer/v1/config/client,Mmixer/v1/config/client/quota.proto=istio.io/api/mixer/v1/config/client,Mmixer/v1/config/client/service.proto=istio.io/api/mixer/v1/config/client,
MIXER_MAPPING := $(MIXER_MAPPING)Mmixer/v1/config/descriptor/log_entry_descriptor.proto=istio.io/api/mixer/v1/config/descriptor,Mmixer/v1/config/descriptor/metric_descriptor.proto=istio.io/api/mixer/v1/config/descriptor,Mmixer/v1/config/descriptor/monitored_resource_descriptor.proto=istio.io/api/mixer/v1/config/descriptor,Mmixer/v1/config/descriptor/principal_descriptor.proto=istio.io/api/mixer/v1/config/descriptor,Mmixer/v1/config/descriptor/quota_descriptor.proto=istio.io/api/mixer/v1/config/descriptor,Mmixer/v1/config/descriptor/value_type.proto=istio.io/api/mixer/v1/config/descriptor,
MIXER_MAPPING := $(MIXER_MAPPING)Mmixer/v1/template/extensions.proto=istio.io/api/mixer/v1/template,Mmixer/v1/template/standard_types.proto=istio.io/api/mixer/v1/template,
MIXER_PLUGIN = $(GOGOSLICK_PLUGIN_PREFIX)$(MIXER_MAPPING)$(PLUGIN_SUFFIX)
ALT_MIXER_PLUGIN = $(GOGO_PLUGIN_PREFIX)$(MIXER_MAPPING)$(PLUGIN_SUFFIX)

## MIXER PROTOS 
MIXER_V1 = mixer/v1
MIXER_V1_PROTOS = $(MIXER_V1)/attributes.proto $(MIXER_V1)/check.proto $(MIXER_V1)/report.proto $(MIXER_V1)/service.proto
MIXER_V1_PB_GOS = $(MIXER_V1_PROTOS:.proto=.pb.go)

MIXER_CONFIG = $(MIXER_V1)/config
MIXER_CONFIG_CLIENT = $(MIXER_CONFIG)/client
MIXER_CONFIG_CLIENT_PROTOS = $(MIXER_CONFIG_CLIENT)/api_spec.proto $(MIXER_CONFIG_CLIENT)/auth.proto \
							$(MIXER_CONFIG_CLIENT)/client_config.proto $(MIXER_CONFIG_CLIENT)/quota.proto \
							$(MIXER_CONFIG_CLIENT)/service.proto
MIXER_CONFIG_CLIENT_PB_GOS = $(MIXER_CONFIG_CLIENT_PROTOS:.proto=.pb.go)

MIXER_CONFIG_DESCRIPTOR = $(MIXER_CONFIG)/descriptor
MIXER_CONFIG_DESCRIPTOR_PROTOS = $(MIXER_CONFIG_DESCRIPTOR)/log_entry_descriptor.proto $(MIXER_CONFIG_DESCRIPTOR)/metric_descriptor.proto \
							$(MIXER_CONFIG_DESCRIPTOR)/monitored_resource_descriptor.proto $(MIXER_CONFIG_DESCRIPTOR)/principal_descriptor.proto \
							$(MIXER_CONFIG_DESCRIPTOR)/quota_descriptor.proto $(MIXER_CONFIG_DESCRIPTOR)/value_type.proto
MIXER_CONFIG_DESCRIPTOR_PB_GOS = $(MIXER_CONFIG_DESCRIPTOR_PROTOS:.proto=.pb.go)

MIXER_TEMPLATE = mixer/v1/template
MIXER_TEMPLATE_PROTOS = $(MIXER_TEMPLATE)/extensions.proto $(MIXER_TEMPLATE)/standard_types.proto
MIXER_TEMPLATE_PB_GOS = $(MIXER_TEMPLATE_PROTOS:.proto=.pb.go)

generate: generate-mixer-go

install-deps:
	# Installing generation deps
	@dep ensure
	@go install github.com/gogo/protobuf/protoc-gen-gogo
	@go install github.com/gogo/protobuf/protoc-gen-gogoslick	
	@go install github.com/gogo/protobuf/gogoreplace
	@go install github.com/gogo/protobuf/protoc-min-version

protoc.version:
	@echo `protoc --version` > protoc.version

protoc-tmp:
	@mkdir -p protoc-tmp

protoc-tmp/$(RPC_PATH): protoc-tmp
	@mkdir -p protoc-tmp/$(RPC_PATH)

download-googleapis-protos: $(STATUS_PROTO) $(CODE_PROTO) $(ERR_PROTO)

$(STATUS_PROTO): protoc-tmp/$(RPC_PATH)
	# Downloading google/rpc/status.proto
	@curl -sS $(GOOGLEAPIS_URL)/google/rpc/status.proto -o $(STATUS_PROTO)

$(CODE_PROTO): protoc-tmp/$(RPC_PATH)
	# Downloading google/rpc/code.proto
	@curl -sS $(GOOGLEAPIS_URL)/google/rpc/code.proto -o $(CODE_PROTO)

$(ERR_PROTO): protoc-tmp/$(RPC_PATH)
	# Downloading google/rpc/error_details.proto
	@curl -sS $(GOOGLEAPIS_URL)/google/rpc/error_details.proto -o $(ERR_PROTO)

# TODO: expand to support the other protos in this repo
generate-mixer-go: install-deps download-googleapis-protos generate-mixer-v1-go generate-mixer-v1-config-go	generate-mixer-v1-template-go

generate-mixer-v1-go: $(MIXER_V1_PB_GOS)

generate-mixer-v1-config-go: $(MIXER_CONFIG_CLIENT_PB_GOS) $(MIXER_CONFIG_DESCRIPTOR_PB_GOS)

generate-mixer-v1-template-go: $(MIXER_TEMPLATE_PB_GOS)

$(MIXER_V1_PB_GOS): $(MIXER_V1_PROTOS)
	## Generate mixer/v1/*.pb.go
	@$(PROTOC) $(PROTO_PATH) $(MIXER_PLUGIN) $^

$(MIXER_CONFIG_CLIENT_PB_GOS) : $(MIXER_CONFIG_CLIENT_PROTOS)
    ## Generate mixer/v1/config/client/*.pb.go
	@$(PROTOC) $(PROTO_PATH) $(MIXER_PLUGIN) $^

$(MIXER_CONFIG_DESCRIPTOR_PB_GOS) : $(MIXER_CONFIG_DESCRIPTOR_PROTOS)
    ## Generate mixer/v1/config/descriptor/*.pb.go
	@$(PROTOC) $(PROTO_PATH) $(MIXER_PLUGIN) $^

$(MIXER_TEMPLATE_PB_GOS) : $(MIXER_TEMPLATE_PROTOS)
    ## Generate mixer/v1/template/*.pb.go
	@$(PROTOC) $(PROTO_PATH) $(MIXER_PLUGIN) $^

mixer/v1/config/fixed_cfg.pb.go : mixer/v1/config/cfg.proto
	# Generate mixer/v1/config/fixed_cfg.pb.go (requires alternate plugin and sed scripting due to issues with google.protobuf.Struct)
	@$(PROTOC) $(PROTO_PATH) $(ALT_MIXER_PLUGIN) $^
	@sed -e 's/*google_protobuf.Struct/interface{}/g' -e 's/ValueType_VALUE_TYPE_UNSPECIFIED/VALUE_TYPE_UNSPECIFIED/g' mixer/v1/config/cfg.pb.go | goimports > mixer/v1/config/fixed_cfg.pb.go
	@rm mixer/v1/config/cfg.pb.go

clean:
	rm -rf protoc-tmp