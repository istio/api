// Copyright 2019 Istio Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: envoy/extensions/stackdriver/config/v1alpha1/config.proto

// clang-format off
// $title: Stackdriver Config
// $description: Configuration for Stackdriver filter.
// $location: https://istio.io/docs/reference/config/proxy_extensions/stackdriver.html
// $weight: 20
// clang-format on

package v1alpha1

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Types of Access logs to export. Does not affect audit logging.
type PluginConfig_AccessLogging int32

const (
	// No Logs.
	PluginConfig_NONE PluginConfig_AccessLogging = 0
	// All logs including both success and error logs.
	PluginConfig_FULL PluginConfig_AccessLogging = 1
	// All error logs. This is currently only available for outbound/client side
	// logs. A request is classified as error when `status>=400 or
	// response_flag != "-"`
	PluginConfig_ERRORS_ONLY PluginConfig_AccessLogging = 2
)

// Enum value maps for PluginConfig_AccessLogging.
var (
	PluginConfig_AccessLogging_name = map[int32]string{
		0: "NONE",
		1: "FULL",
		2: "ERRORS_ONLY",
	}
	PluginConfig_AccessLogging_value = map[string]int32{
		"NONE":        0,
		"FULL":        1,
		"ERRORS_ONLY": 2,
	}
)

func (x PluginConfig_AccessLogging) Enum() *PluginConfig_AccessLogging {
	p := new(PluginConfig_AccessLogging)
	*p = x
	return p
}

func (x PluginConfig_AccessLogging) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PluginConfig_AccessLogging) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_enumTypes[0].Descriptor()
}

func (PluginConfig_AccessLogging) Type() protoreflect.EnumType {
	return &file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_enumTypes[0]
}

func (x PluginConfig_AccessLogging) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PluginConfig_AccessLogging.Descriptor instead.
func (PluginConfig_AccessLogging) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_rawDescGZIP(), []int{1, 0}
}

// Custom instance configuration overrides.
// Provides a way to customize logs.
type CustomConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Optional) Collection of tag names and tag expressions to include in the
	// instance. Conflicts are resolved by the tag name by overriding previously
	// supplied values.
	Dimensions map[string]string `protobuf:"bytes,1,rep,name=dimensions,proto3" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// (Optional) A list of tags to remove.
	// Not implemented yet.
	// $hide_from_docs
	TagsToRemove []string `protobuf:"bytes,2,rep,name=tags_to_remove,json=tagsToRemove,proto3" json:"tags_to_remove,omitempty"`
}

func (x *CustomConfig) Reset() {
	*x = CustomConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomConfig) ProtoMessage() {}

func (x *CustomConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomConfig.ProtoReflect.Descriptor instead.
func (*CustomConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_rawDescGZIP(), []int{0}
}

func (x *CustomConfig) GetDimensions() map[string]string {
	if x != nil {
		return x.Dimensions
	}
	return nil
}

func (x *CustomConfig) GetTagsToRemove() []string {
	if x != nil {
		return x.TagsToRemove
	}
	return nil
}

// next id: 17
type PluginConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Controls whether to export server access log.
	// This is deprecated in favor of AccessLogging enum.
	//
	// Deprecated: Marked as deprecated in envoy/extensions/stackdriver/config/v1alpha1/config.proto.
	DisableServerAccessLogging bool `protobuf:"varint,1,opt,name=disable_server_access_logging,json=disableServerAccessLogging,proto3" json:"disable_server_access_logging,omitempty"`
	// Optional. Allows configuration of the size of the LogWrite request. The
	// size is in bytes, so that it allows for better performance. Default is 4MB.
	// The size of one log entry within LogWrite request is approx 1Kb.
	MaxLogBatchSizeInBytes int32 `protobuf:"varint,12,opt,name=max_log_batch_size_in_bytes,json=maxLogBatchSizeInBytes,proto3" json:"max_log_batch_size_in_bytes,omitempty"`
	// Optional. Allows configuration of the time between calls out to the
	// stackdriver logging service to report buffered LogWrite request.
	// Customers can choose to report more aggressively by keeping shorter report
	// interval if needed. Default is 10s.
	LogReportDuration *duration.Duration `protobuf:"bytes,13,opt,name=log_report_duration,json=logReportDuration,proto3" json:"log_report_duration,omitempty"`
	// Optional. Controls whether to export audit log.
	EnableAuditLog bool `protobuf:"varint,11,opt,name=enable_audit_log,json=enableAuditLog,proto3" json:"enable_audit_log,omitempty"`
	// Optional. FQDN of destination service that the request routed to, e.g.
	// productpage.default.svc.cluster.local. If not provided, request host header
	// will be used instead
	DestinationServiceName string `protobuf:"bytes,2,opt,name=destination_service_name,json=destinationServiceName,proto3" json:"destination_service_name,omitempty"`
	// Optional. Controls whether or not to export mesh edges to a mesh edges
	// service. This is disabled by default.
	// Deprecated -- Mesh edge reporting is no longer supported and this setting
	// is no-op.
	//
	// Deprecated: Marked as deprecated in envoy/extensions/stackdriver/config/v1alpha1/config.proto.
	EnableMeshEdgesReporting bool `protobuf:"varint,3,opt,name=enable_mesh_edges_reporting,json=enableMeshEdgesReporting,proto3" json:"enable_mesh_edges_reporting,omitempty"`
	// Optional. Allows configuration of the time between calls out to the mesh
	// edges service to report *NEW* edges. The minimum configurable duration is
	// `10s`. NOTE: This option ONLY configures the intermediate reporting of
	// novel edges. Once every `10m`, all edges observed in that 10m window are
	// reported and the local cache is cleared.
	// The default duration is `1m`. Any value greater than `10m` will result in
	// reporting every `10m`.
	// Deprecated -- Mesh edge reporting is no longer supported and this setting
	// is no-op.
	//
	// Deprecated: Marked as deprecated in envoy/extensions/stackdriver/config/v1alpha1/config.proto.
	MeshEdgesReportingDuration *duration.Duration `protobuf:"bytes,4,opt,name=mesh_edges_reporting_duration,json=meshEdgesReportingDuration,proto3" json:"mesh_edges_reporting_duration,omitempty"`
	// maximum size of the peer metadata cache.
	// A long lived proxy that connects with many transient peers can build up a
	// large cache. To turn off the cache, set this field to a negative value.
	MaxPeerCacheSize int32 `protobuf:"varint,5,opt,name=max_peer_cache_size,json=maxPeerCacheSize,proto3" json:"max_peer_cache_size,omitempty"`
	// Optional: Disable using host header as a fallback if destination service is
	// not available from the controlplane. Disable the fallback if the host
	// header originates outsides the mesh, like at ingress.
	DisableHostHeaderFallback bool `protobuf:"varint,6,opt,name=disable_host_header_fallback,json=disableHostHeaderFallback,proto3" json:"disable_host_header_fallback,omitempty"`
	// Optional. Allows configuration of the number of traffic assertions to batch
	// into a single request. Default is 100. Max is 1000.
	MaxEdgesBatchSize int32 `protobuf:"varint,7,opt,name=max_edges_batch_size,json=maxEdgesBatchSize,proto3" json:"max_edges_batch_size,omitempty"`
	// Optional. Allows disabling of reporting of the request and response size
	// metrics for HTTP traffic. Defaults to false (request and response size
	// metrics are enabled).
	// Deprecated -- use `metrics_overrides` instead.
	// if `metrics_overrides` is used, this value will be ignored.
	//
	// Deprecated: Marked as deprecated in envoy/extensions/stackdriver/config/v1alpha1/config.proto.
	DisableHttpSizeMetrics bool `protobuf:"varint,8,opt,name=disable_http_size_metrics,json=disableHttpSizeMetrics,proto3" json:"disable_http_size_metrics,omitempty"`
	// Optional. Allows enabling log compression for stackdriver access logs.
	EnableLogCompression *wrappers.BoolValue `protobuf:"bytes,9,opt,name=enable_log_compression,json=enableLogCompression,proto3" json:"enable_log_compression,omitempty"`
	// Optional. Controls what type of logs to export..
	AccessLogging PluginConfig_AccessLogging `protobuf:"varint,10,opt,name=access_logging,json=accessLogging,proto3,enum=stackdriver.config.v1alpha1.PluginConfig_AccessLogging" json:"access_logging,omitempty"`
	// CEL expression for filtering access logging. If the expression evaluates
	// to true, an access log entry will be generated. Otherwise, no access log
	// entry will be generated. If there are any type errors, the CEL expression
	// is evaluated as false. More details on type checking can be found
	// at https://kubernetes.io/docs/reference/using-api/cel/#type-checking.
	// A common error is referring to a non-existent field in the log entry.
	// It's crucial to note that in Envoy, the fields that appear in access log
	// entries can vary. This variation is influenced by several factors,
	// including the protocol in use (such as HTTP or TCP), the applied filters,
	// and the specific configuration of the Envoy instance. Therefore, when
	// using CEL expressions for filtering access logs, it's essential to ensure
	// that the expressions accurately refer to existing fields in the log entry.
	// The has() macro in CEL may be used in CEL expressions to check if a field
	// is accessible before attempting to access the field's value.
	// You can also quickly test CEL expressions at the CEL Playground
	// at https://playcel.undistro.io/.
	// NOTE: Audit logs ignore configured filters.
	AccessLoggingFilterExpression string `protobuf:"bytes,17,opt,name=access_logging_filter_expression,json=accessLoggingFilterExpression,proto3" json:"access_logging_filter_expression,omitempty"`
	// (Optional) Collection of tag names and tag expressions to include in the
	// logs. Conflicts are resolved by the tag name by overriding previously
	// supplied values. Does not apply to audit logs.
	// See
	// https://istio.io/latest/docs/tasks/observability/metrics/customize-metrics/#use-expressions-for-values
	// for more details about the expression language.
	CustomLogConfig *CustomConfig `protobuf:"bytes,14,opt,name=custom_log_config,json=customLogConfig,proto3" json:"custom_log_config,omitempty"`
	// Optional. Controls the metric expiry duration. If a metric time series is
	// not updated for the given duration, it will be purged from time series
	// cache as well as metric reporting. If this is not set or set to 0, time
	// series will never be expired. This option is useful to avoid unbounded
	// metric label explodes proxy memory.
	MetricExpiryDuration *duration.Duration `protobuf:"bytes,15,opt,name=metric_expiry_duration,json=metricExpiryDuration,proto3" json:"metric_expiry_duration,omitempty"`
	// Optional. Allows altering metrics behavior.
	// Metric names for specifying overloads drop the `istio.io/service` prefix.
	// Examples: `server/request_count`, `client/roundtrip_latencies`
	MetricsOverrides map[string]*MetricsOverride `protobuf:"bytes,16,rep,name=metrics_overrides,json=metricsOverrides,proto3" json:"metrics_overrides,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PluginConfig) Reset() {
	*x = PluginConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginConfig) ProtoMessage() {}

func (x *PluginConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginConfig.ProtoReflect.Descriptor instead.
func (*PluginConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_rawDescGZIP(), []int{1}
}

// Deprecated: Marked as deprecated in envoy/extensions/stackdriver/config/v1alpha1/config.proto.
func (x *PluginConfig) GetDisableServerAccessLogging() bool {
	if x != nil {
		return x.DisableServerAccessLogging
	}
	return false
}

func (x *PluginConfig) GetMaxLogBatchSizeInBytes() int32 {
	if x != nil {
		return x.MaxLogBatchSizeInBytes
	}
	return 0
}

func (x *PluginConfig) GetLogReportDuration() *duration.Duration {
	if x != nil {
		return x.LogReportDuration
	}
	return nil
}

func (x *PluginConfig) GetEnableAuditLog() bool {
	if x != nil {
		return x.EnableAuditLog
	}
	return false
}

func (x *PluginConfig) GetDestinationServiceName() string {
	if x != nil {
		return x.DestinationServiceName
	}
	return ""
}

// Deprecated: Marked as deprecated in envoy/extensions/stackdriver/config/v1alpha1/config.proto.
func (x *PluginConfig) GetEnableMeshEdgesReporting() bool {
	if x != nil {
		return x.EnableMeshEdgesReporting
	}
	return false
}

// Deprecated: Marked as deprecated in envoy/extensions/stackdriver/config/v1alpha1/config.proto.
func (x *PluginConfig) GetMeshEdgesReportingDuration() *duration.Duration {
	if x != nil {
		return x.MeshEdgesReportingDuration
	}
	return nil
}

func (x *PluginConfig) GetMaxPeerCacheSize() int32 {
	if x != nil {
		return x.MaxPeerCacheSize
	}
	return 0
}

func (x *PluginConfig) GetDisableHostHeaderFallback() bool {
	if x != nil {
		return x.DisableHostHeaderFallback
	}
	return false
}

func (x *PluginConfig) GetMaxEdgesBatchSize() int32 {
	if x != nil {
		return x.MaxEdgesBatchSize
	}
	return 0
}

// Deprecated: Marked as deprecated in envoy/extensions/stackdriver/config/v1alpha1/config.proto.
func (x *PluginConfig) GetDisableHttpSizeMetrics() bool {
	if x != nil {
		return x.DisableHttpSizeMetrics
	}
	return false
}

func (x *PluginConfig) GetEnableLogCompression() *wrappers.BoolValue {
	if x != nil {
		return x.EnableLogCompression
	}
	return nil
}

func (x *PluginConfig) GetAccessLogging() PluginConfig_AccessLogging {
	if x != nil {
		return x.AccessLogging
	}
	return PluginConfig_NONE
}

func (x *PluginConfig) GetAccessLoggingFilterExpression() string {
	if x != nil {
		return x.AccessLoggingFilterExpression
	}
	return ""
}

func (x *PluginConfig) GetCustomLogConfig() *CustomConfig {
	if x != nil {
		return x.CustomLogConfig
	}
	return nil
}

func (x *PluginConfig) GetMetricExpiryDuration() *duration.Duration {
	if x != nil {
		return x.MetricExpiryDuration
	}
	return nil
}

func (x *PluginConfig) GetMetricsOverrides() map[string]*MetricsOverride {
	if x != nil {
		return x.MetricsOverrides
	}
	return nil
}

// Provides behavior modifications for Cloud Monitoring metrics.
type MetricsOverride struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. If true, no data for the associated metric will be collected or
	// exported.
	Drop bool `protobuf:"varint,1,opt,name=drop,proto3" json:"drop,omitempty"`
	// Optional. Maps tag names to value expressions that will be used at
	// reporting time. If the tag name does not match a well-known tag for the
	// istio Cloud Monitoring metrics, the configuration will have no effect.
	TagOverrides map[string]string `protobuf:"bytes,2,rep,name=tag_overrides,json=tagOverrides,proto3" json:"tag_overrides,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MetricsOverride) Reset() {
	*x = MetricsOverride{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsOverride) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsOverride) ProtoMessage() {}

func (x *MetricsOverride) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsOverride.ProtoReflect.Descriptor instead.
func (*MetricsOverride) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_rawDescGZIP(), []int{2}
}

func (x *MetricsOverride) GetDrop() bool {
	if x != nil {
		return x.Drop
	}
	return false
}

func (x *MetricsOverride) GetTagOverrides() map[string]string {
	if x != nil {
		return x.TagOverrides
	}
	return nil
}

var File_envoy_extensions_stackdriver_config_v1alpha1_config_proto protoreflect.FileDescriptor

var file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_rawDesc = []byte{
	0x0a, 0x39, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x01, 0x0a, 0x0c, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x59, 0x0a, 0x0a, 0x64, 0x69, 0x6d,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x61, 0x67, 0x73, 0x5f, 0x74, 0x6f, 0x5f,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x61,
	0x67, 0x73, 0x54, 0x6f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x1a, 0x3d, 0x0a, 0x0f, 0x44, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x80, 0x0b, 0x0a, 0x0c, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x45, 0x0a, 0x1d, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x42, 0x02, 0x18, 0x01, 0x52, 0x1a, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x12, 0x3b, 0x0a, 0x1b, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x6d, 0x61, 0x78, 0x4c, 0x6f, 0x67, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x49, 0x6e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x49,
	0x0a, 0x13, 0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x4c, 0x6f, 0x67, 0x12, 0x38, 0x0a, 0x18, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a,
	0x1b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x68, 0x5f, 0x65, 0x64, 0x67,
	0x65, 0x73, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x52, 0x18, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x65,
	0x73, 0x68, 0x45, 0x64, 0x67, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x60, 0x0a, 0x1d, 0x6d, 0x65, 0x73, 0x68, 0x5f, 0x65, 0x64, 0x67, 0x65, 0x73, 0x5f, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x02, 0x18, 0x01, 0x52, 0x1a, 0x6d, 0x65, 0x73, 0x68, 0x45, 0x64, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x10, 0x6d, 0x61, 0x78, 0x50, 0x65, 0x65, 0x72, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x3f, 0x0a, 0x1c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x68, 0x6f, 0x73,
	0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x48, 0x6f, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x12, 0x2f, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x65, 0x64, 0x67, 0x65, 0x73, 0x5f,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x11, 0x6d, 0x61, 0x78, 0x45, 0x64, 0x67, 0x65, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x3d, 0x0a, 0x19, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x52, 0x16, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x48, 0x74, 0x74, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x12, 0x50, 0x0a, 0x16, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x67,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x14,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x5e, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x12, 0x47, 0x0a, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1d,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a,
	0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x6f, 0x67, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x4f, 0x0a, 0x16, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x14, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6c, 0x0a, 0x11, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x10, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69,
	0x64, 0x65, 0x73, 0x1a, 0x71, 0x0a, 0x15, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4f, 0x76,
	0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x42,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x34, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x53, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x02, 0x22, 0xcb, 0x01, 0x0a,
	0x0f, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x72, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x64, 0x72, 0x6f, 0x70, 0x12, 0x63, 0x0a, 0x0d, 0x74, 0x61, 0x67, 0x5f, 0x6f, 0x76, 0x65, 0x72,
	0x72, 0x69, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x54, 0x61, 0x67, 0x4f, 0x76, 0x65,
	0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x74, 0x61, 0x67,
	0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x1a, 0x3f, 0x0a, 0x11, 0x54, 0x61, 0x67,
	0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x3b, 0x5a, 0x39, 0x69, 0x73,
	0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_rawDescOnce sync.Once
	file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_rawDescData = file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_rawDesc
)

func file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_rawDescGZIP() []byte {
	file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_rawDescData)
	})
	return file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_rawDescData
}

var file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_goTypes = []any{
	(PluginConfig_AccessLogging)(0), // 0: stackdriver.config.v1alpha1.PluginConfig.AccessLogging
	(*CustomConfig)(nil),            // 1: stackdriver.config.v1alpha1.CustomConfig
	(*PluginConfig)(nil),            // 2: stackdriver.config.v1alpha1.PluginConfig
	(*MetricsOverride)(nil),         // 3: stackdriver.config.v1alpha1.MetricsOverride
	nil,                             // 4: stackdriver.config.v1alpha1.CustomConfig.DimensionsEntry
	nil,                             // 5: stackdriver.config.v1alpha1.PluginConfig.MetricsOverridesEntry
	nil,                             // 6: stackdriver.config.v1alpha1.MetricsOverride.TagOverridesEntry
	(*duration.Duration)(nil),       // 7: google.protobuf.Duration
	(*wrappers.BoolValue)(nil),      // 8: google.protobuf.BoolValue
}
var file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_depIdxs = []int32{
	4,  // 0: stackdriver.config.v1alpha1.CustomConfig.dimensions:type_name -> stackdriver.config.v1alpha1.CustomConfig.DimensionsEntry
	7,  // 1: stackdriver.config.v1alpha1.PluginConfig.log_report_duration:type_name -> google.protobuf.Duration
	7,  // 2: stackdriver.config.v1alpha1.PluginConfig.mesh_edges_reporting_duration:type_name -> google.protobuf.Duration
	8,  // 3: stackdriver.config.v1alpha1.PluginConfig.enable_log_compression:type_name -> google.protobuf.BoolValue
	0,  // 4: stackdriver.config.v1alpha1.PluginConfig.access_logging:type_name -> stackdriver.config.v1alpha1.PluginConfig.AccessLogging
	1,  // 5: stackdriver.config.v1alpha1.PluginConfig.custom_log_config:type_name -> stackdriver.config.v1alpha1.CustomConfig
	7,  // 6: stackdriver.config.v1alpha1.PluginConfig.metric_expiry_duration:type_name -> google.protobuf.Duration
	5,  // 7: stackdriver.config.v1alpha1.PluginConfig.metrics_overrides:type_name -> stackdriver.config.v1alpha1.PluginConfig.MetricsOverridesEntry
	6,  // 8: stackdriver.config.v1alpha1.MetricsOverride.tag_overrides:type_name -> stackdriver.config.v1alpha1.MetricsOverride.TagOverridesEntry
	3,  // 9: stackdriver.config.v1alpha1.PluginConfig.MetricsOverridesEntry.value:type_name -> stackdriver.config.v1alpha1.MetricsOverride
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_init() }
func file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_init() {
	if File_envoy_extensions_stackdriver_config_v1alpha1_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CustomConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PluginConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MetricsOverride); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_msgTypes,
	}.Build()
	File_envoy_extensions_stackdriver_config_v1alpha1_config_proto = out.File
	file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_rawDesc = nil
	file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_goTypes = nil
	file_envoy_extensions_stackdriver_config_v1alpha1_config_proto_depIdxs = nil
}
