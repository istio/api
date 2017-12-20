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

## MIXER PROTOS 
MIXER_V1 = mixer/v1
ATTR_PROTO = $(MIXER_V1)/attributes.proto
CHECK_PROTO = $(MIXER_V1)/check.proto
REPORT_PROTO = $(MIXER_V1)/report.proto
SVC_PROTO = $(MIXER_V1)/service.proto
MIXER_V1_PROTOS = $(ATTR_PROTO) $(CHECK_PROTO) $(REPORT_PROTO) $(SVC_PROTO)
MIXER_V1_PB_GOS = $(MIXER_V1_PROTOS:.proto=.pb.go)

MIXER_CONFIG = $(MIXER_V1)/config
MIXER_CONFIG_CLIENT = $(MIXER_CONFIG)/client
API_SPEC_PROTO = $(MIXER_CONFIG_CLIENT)/api_spec.proto
AUTH_PROTO = $(MIXER_CONFIG_CLIENT)/auth.proto
CLIENT_CONFIG_PROTO = $(MIXER_CONFIG_CLIENT)/client_config.proto
QUOTA_PROTO = $(MIXER_CONFIG_CLIENT)/quota.proto
CONFIG_SVC_PROTO = $(MIXER_CONFIG_CLIENT)/service.proto
MIXER_CONFIG_CLIENT_PROTOS = $(API_SPEC_PROTO) $(AUTH_PROTO) $(CLIENT_CONFIG_PROTO) $(QUOTA_PROTO) $(CONFIG_SVC_PROTO)
MIXER_CONFIG_CLIENT_PB_GOS = $(MIXER_CONFIG_CLIENT_PROTOS:.proto=.pb.go)

MIXER_CONFIG_DESCRIPTOR = $(MIXER_CONFIG)/descriptor
LOG_ENTRY_PROTO = $(MIXER_CONFIG_DESCRIPTOR)/log_entry_descriptor.proto
METRIC_PROTO = $(MIXER_CONFIG_DESCRIPTOR)/metric_descriptor.proto
MR_PROTO = $(MIXER_CONFIG_DESCRIPTOR)/monitored_resource_descriptor.proto
PRINCIPAL_PROTO = $(MIXER_CONFIG_DESCRIPTOR)/principal_descriptor.proto
QUOTA_DESC_PROTO = $(MIXER_CONFIG_DESCRIPTOR)/quota_descriptor.proto
VALUE_TYPE_PROTO = $(MIXER_CONFIG_DESCRIPTOR)/value_type.proto
MIXER_CONFIG_DESCRIPTOR_PROTOS = $(LOG_ENTRY_PROTO) $(METRIC_PROTO) $(MR_PROTO) $(PRINCIPAL_PROTO) $(QUOTA_DESC_PROTO) $(VALUE_TYPE_PROTO)
MIXER_CONFIG_DESCRIPTOR_PB_GOS = $(MIXER_CONFIG_DESCRIPTOR_PROTOS:.proto=.pb.go)

MIXER_TEMPLATE = mixer/v1/template
TEMPLATE_EXT_PROTO = $(MIXER_TEMPLATE)/extensions.proto
STD_TYPES_PROTO = $(MIXER_TEMPLATE)/standard_types.proto
MIXER_TEMPLATE_PROTOS = $(TEMPLATE_EXT_PROTO) $(STD_TYPES_PROTO)
MIXER_TEMPLATE_PB_GOS = $(MIXER_TEMPLATE_PROTOS:.proto=.pb.go)

## MIXER-PROTO SPECIFIC MAPPINGS
MIXER_V1_PKG = istio.io/api/mixer/v1
CLIENT_PKG = istio.io/api/mixer/v1/config/client
DESCRIPTOR_PKG = istio.io/api/mixer/v1/config/descriptor
TEMPLATE_PKG = istio.io/api/mixer/v1/template

MIXER_MAPPING := $(MAPPING)M$(ATTR_PROTO)=$(MIXER_V1_PKG),M$(CHECK_PROTO)=$(MIXER_V1_PKG),M$(REPORT_PROTO)=$(MIXER_V1_PKG),M$(SVC_PROTO)=$(MIXER_V1_PKG),
MIXER_MAPPING := $(MIXER_MAPPING)M$(API_SPEC_PROTO)=$(CLIENT_PKG),M$(AUTH_PROTO)=$(CLIENT_PKG),M$(CLIENT_CONFIG_PROTO)=$(CLIENT_PKG),M$(QUOTA_PROTO)=$(CLIENT_PKG),M$(CONFIG_SVC_PROTO)=$(CLIENT_PKG),
MIXER_MAPPING := $(MIXER_MAPPING)M$(LOG_ENTRY_PROTO)=$(DESCRIPTOR_PKG),M$(METRIC_PROTO)=$(DESCRIPTOR_PKG),M$(MR_PROTO)=$(DESCRIPTOR_PKG),M$(PRINCIPAL_PROTO)=$(DESCRIPTOR_PKG),M$(QUOTA_DESC_PROTO)=$(DESCRIPTOR_PKG),M$(VALUE_TYPE_PROTO)=$(DESCRIPTOR_PKG),
MIXER_MAPPING := $(MIXER_MAPPING)M$(TEMPLATE_EXT_PROTO)=$(TEMPLATE_PKG),M$(STD_TYPES_PROTO)=$(TEMPLATE_PKG),
MIXER_PLUGIN = $(GOGOSLICK_PLUGIN_PREFIX)$(MIXER_MAPPING)$(PLUGIN_SUFFIX)
ALT_MIXER_PLUGIN = $(GOGO_PLUGIN_PREFIX)$(MIXER_MAPPING)$(PLUGIN_SUFFIX)

#####################
# Generation Rule
#####################
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
generate-mixer-go: install-deps download-googleapis-protos generate-mixer-v1-go generate-mixer-v1-config-go generate-mixer-v1-template-go

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
