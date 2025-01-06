// Code generated by protoc-gen-alias. DO NOT EDIT.
package v1

import "istio.io/api/telemetry/v1alpha1"

// <!-- crd generation tags
// +cue-gen:Telemetry:groupName:telemetry.istio.io
// +cue-gen:Telemetry:versions:v1alpha1,v1
// +cue-gen:Telemetry:storageVersion
// +cue-gen:Telemetry:annotations:helm.sh/resource-policy=keep
// +cue-gen:Telemetry:labels:app=istio-pilot,chart=istio,istio=telemetry,heritage=Tiller,release=istio
// +cue-gen:Telemetry:subresource:status
// +cue-gen:Telemetry:scope:Namespaced
// +cue-gen:Telemetry:resource:categories=istio-io,telemetry-istio-io,shortNames=telemetry,plural=telemetries
// +cue-gen:Telemetry:preserveUnknownFields:false
// +cue-gen:Telemetry:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp
// is a timestamp representing the server time when this object was created. It
// is not guaranteed to be set in happens-before order across separate
// operations. Clients may not set this value. It is represented in RFC3339 form
// and is in UTC. Populated by the system. Read-only. Null for lists. More info:
// https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=telemetry.istio.io/v1alpha1
// +genclient
// +k8s:deepcopy-gen=true
// -->
// +kubebuilder:validation:XValidation:message="only one of targetRefs or selector can be set",rule="oneof(self.selector, self.targetRef, self.targetRefs)"
type Telemetry = v1alpha1.Telemetry

// Tracing configures tracing behavior for workloads within a mesh.
// It can be used to enable/disable tracing, as well as to set sampling
// rates and custom tag extraction.
//
// Tracing configuration support overrides of the fields `providers`,
// `random_sampling_percentage`, `disable_span_reporting`, and `custom_tags` at
// each level in the configuration hierarchy, with missing values filled in
// from parent resources. However, when specified, `custom_tags` will
// fully replace any values provided by parent configuration.
type Tracing = v1alpha1.Tracing

// TracingSelector provides a coarse-grained ability to configure tracing
// behavior based on certain traffic metadata (such as traffic direction).
type Tracing_TracingSelector = v1alpha1.Tracing_TracingSelector

// CustomTag defines a tag to be added to a trace span that is based on
// an operator-supplied value. This value can either be a hard-coded value,
// a value taken from an environment variable known to the sidecar proxy, or
// from a request header.
//
// NOTE: when specified, `custom_tags` will fully replace any values provided
// by parent configuration.
type Tracing_CustomTag = v1alpha1.Tracing_CustomTag

// Literal adds the same, hard-coded value to each span.
type Tracing_CustomTag_Literal = v1alpha1.Tracing_CustomTag_Literal

// Environment adds the value of an environment variable to each span.
type Tracing_CustomTag_Environment = v1alpha1.Tracing_CustomTag_Environment

// RequestHeader adds the value of an header from the request to each
// span.
type Tracing_CustomTag_Header = v1alpha1.Tracing_CustomTag_Header
type Tracing_Literal = v1alpha1.Tracing_Literal
type Tracing_Environment = v1alpha1.Tracing_Environment
type Tracing_RequestHeader = v1alpha1.Tracing_RequestHeader

// Used to bind Telemetry configuration to specific providers for
// targeted customization.
type ProviderRef = v1alpha1.ProviderRef

// Metrics defines the workload-level overrides for metrics generation behavior
// within a mesh. It can be used to enable/disable metrics generation, as well
// as to customize the dimensions of the generated metrics.
type Metrics = v1alpha1.Metrics

// Provides a mechanism for matching metrics for the application of override
// behaviors.
type MetricSelector = v1alpha1.MetricSelector

// Curated list of known metric types that is supported by Istio metric
// providers. See also:
// https://istio.io/latest/docs/reference/config/metrics/#metrics
type MetricSelector_IstioMetric = v1alpha1.MetricSelector_IstioMetric

// Use of this enum indicates that the override should apply to all Istio
// default metrics.
const MetricSelector_ALL_METRICS MetricSelector_IstioMetric = v1alpha1.MetricSelector_ALL_METRICS

// Counter of requests to/from an application, generated for HTTP, HTTP/2,
// and GRPC traffic.
//
// The Prometheus provider exports this metric as: `istio_requests_total`.
//
// The Stackdriver provider exports this metric as:
//
// - `istio.io/service/server/request_count` (SERVER mode)
// - `istio.io/service/client/request_count` (CLIENT mode)
const MetricSelector_REQUEST_COUNT MetricSelector_IstioMetric = v1alpha1.MetricSelector_REQUEST_COUNT

// Histogram of request durations, generated for HTTP, HTTP/2, and GRPC
// traffic.
//
// The Prometheus provider exports this metric as:
// `istio_request_duration_milliseconds`.
//
// The Stackdriver provider exports this metric as:
//
// - `istio.io/service/server/response_latencies` (SERVER mode)
// - `istio.io/service/client/roundtrip_latencies` (CLIENT mode)
const MetricSelector_REQUEST_DURATION MetricSelector_IstioMetric = v1alpha1.MetricSelector_REQUEST_DURATION

// Histogram of request body sizes, generated for HTTP, HTTP/2, and GRPC
// traffic.
//
// The Prometheus provider exports this metric as: `istio_request_bytes`.
//
// The Stackdriver provider exports this metric as:
//
// - `istio.io/service/server/request_bytes` (SERVER mode)
// - `istio.io/service/client/request_bytes` (CLIENT mode)
const MetricSelector_REQUEST_SIZE MetricSelector_IstioMetric = v1alpha1.MetricSelector_REQUEST_SIZE

// Histogram of response body sizes, generated for HTTP, HTTP/2, and GRPC
// traffic.
//
// The Prometheus provider exports this metric as: `istio_response_bytes`.
//
// The Stackdriver provider exports this metric as:
//
// - `istio.io/service/server/response_bytes` (SERVER mode)
// - `istio.io/service/client/response_bytes` (CLIENT mode)
const MetricSelector_RESPONSE_SIZE MetricSelector_IstioMetric = v1alpha1.MetricSelector_RESPONSE_SIZE

// Counter of TCP connections opened over lifetime of workload.
//
// The Prometheus provider exports this metric as:
// `istio_tcp_connections_opened_total`.
//
// The Stackdriver provider exports this metric as:
//
// - `istio.io/service/server/connection_open_count` (SERVER mode)
// - `istio.io/service/client/connection_open_count` (CLIENT mode)
const MetricSelector_TCP_OPENED_CONNECTIONS MetricSelector_IstioMetric = v1alpha1.MetricSelector_TCP_OPENED_CONNECTIONS

// Counter of TCP connections closed over lifetime of workload.
//
// The Prometheus provider exports this metric as:
// `istio_tcp_connections_closed_total`.
//
// The Stackdriver provider exports this metric as:
//
// - `istio.io/service/server/connection_close_count` (SERVER mode)
// - `istio.io/service/client/connection_close_count` (CLIENT mode)
const MetricSelector_TCP_CLOSED_CONNECTIONS MetricSelector_IstioMetric = v1alpha1.MetricSelector_TCP_CLOSED_CONNECTIONS

// Counter of bytes sent during a response over a TCP connection.
//
// The Prometheus provider exports this metric as:
// `istio_tcp_sent_bytes_total`.
//
// The Stackdriver provider exports this metric as:
//
// - `istio.io/service/server/sent_bytes_count` (SERVER mode)
// - `istio.io/service/client/sent_bytes_count` (CLIENT mode)
const MetricSelector_TCP_SENT_BYTES MetricSelector_IstioMetric = v1alpha1.MetricSelector_TCP_SENT_BYTES

// Counter of bytes received during a request over a TCP connection.
//
// The Prometheus provider exports this metric as:
// `istio_tcp_received_bytes_total`.
//
// The Stackdriver provider exports this metric as:
//
// - `istio.io/service/server/received_bytes_count` (SERVER mode)
// - `istio.io/service/client/received_bytes_count` (CLIENT mode)
const MetricSelector_TCP_RECEIVED_BYTES MetricSelector_IstioMetric = v1alpha1.MetricSelector_TCP_RECEIVED_BYTES

// Counter incremented for every gRPC messages sent from a client.
//
// The Prometheus provider exports this metric as:
// `istio_request_messages_total`
const MetricSelector_GRPC_REQUEST_MESSAGES MetricSelector_IstioMetric = v1alpha1.MetricSelector_GRPC_REQUEST_MESSAGES

// Counter incremented for every gRPC messages sent from a server.
//
// The Prometheus provider exports this metric as:
// `istio_response_messages_total`
const MetricSelector_GRPC_RESPONSE_MESSAGES MetricSelector_IstioMetric = v1alpha1.MetricSelector_GRPC_RESPONSE_MESSAGES

// One of the well-known [Istio Standard Metrics](https://istio.io/latest/docs/reference/config/metrics/).
type MetricSelector_Metric = v1alpha1.MetricSelector_Metric

// Allows free-form specification of a metric. No validation of custom
// metrics is provided.
// +kubebuilder:validation:MinLength=1
type MetricSelector_CustomMetric = v1alpha1.MetricSelector_CustomMetric

// MetricsOverrides defines custom metric generation behavior for an individual
// metric or the set of all standard metrics.
type MetricsOverrides = v1alpha1.MetricsOverrides

// TagOverride specifies an operation to perform on a metric dimension (also
// known as a `label`). Tags may be added, removed, or have their default
// values overridden.
// +kubebuilder:validation:XValidation:message="value must be set when operation is UPSERT",rule="default(self.operation, ”) == 'UPSERT' ? self.value != ” : true"
// +kubebuilder:validation:XValidation:message="value must not be set when operation is REMOVE",rule="default(self.operation, ”) == 'REMOVE' ? !has(self.value) : true"
type MetricsOverrides_TagOverride = v1alpha1.MetricsOverrides_TagOverride
type MetricsOverrides_TagOverride_Operation = v1alpha1.MetricsOverrides_TagOverride_Operation

// Insert or Update the tag with the provided value expression. The
// `value` field MUST be specified if `UPSERT` is used as the operation.
const MetricsOverrides_TagOverride_UPSERT MetricsOverrides_TagOverride_Operation = v1alpha1.MetricsOverrides_TagOverride_UPSERT

// Specifies that the tag should not be included in the metric when
// generated.
const MetricsOverrides_TagOverride_REMOVE MetricsOverrides_TagOverride_Operation = v1alpha1.MetricsOverrides_TagOverride_REMOVE

// Access logging defines the workload-level overrides for access log
// generation. It can be used to select provider or enable/disable access log
// generation for a workload.
type AccessLogging = v1alpha1.AccessLogging

// LogSelector provides a coarse-grained ability to configure logging behavior
// based on certain traffic metadata (such as traffic direction). LogSelector
// applies to traffic metadata which is not represented in the attribute set
// currently supported by [filters](https://istio.io/latest/docs/reference/config/telemetry/#AccessLogging-Filter).
// It allows control planes to limit the configuration sent to individual workloads.
// Finer-grained logging behavior can be further configured via `filter`.
type AccessLogging_LogSelector = v1alpha1.AccessLogging_LogSelector

// Allows specification of an access log filter.
type AccessLogging_Filter = v1alpha1.AccessLogging_Filter

// WorkloadMode allows selection of the role of the underlying workload in
// network traffic. A workload is considered as acting as a `SERVER` if it is
// the destination of the traffic (that is, traffic direction, from the
// perspective of the workload is *inbound*). If the workload is the source of
// the network traffic, it is considered to be in `CLIENT` mode (traffic is
// *outbound* from the workload).
type WorkloadMode = v1alpha1.WorkloadMode

// Selects for scenarios when the workload is either the
// source or destination of the network traffic.
const WorkloadMode_CLIENT_AND_SERVER WorkloadMode = v1alpha1.WorkloadMode_CLIENT_AND_SERVER

// Selects for scenarios when the workload is the
// source of the network traffic.
const WorkloadMode_CLIENT WorkloadMode = v1alpha1.WorkloadMode_CLIENT

// Selects for scenarios when the workload is the
// destination of the network traffic.
const WorkloadMode_SERVER WorkloadMode = v1alpha1.WorkloadMode_SERVER
