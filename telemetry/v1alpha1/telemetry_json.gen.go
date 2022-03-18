// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: telemetry/v1alpha1/telemetry.proto

// Telemetry defines how the telemetry is generated for workloads within a mesh.
//
// For mesh level configuration, put the resource in root configuration
// namespace for your Istio installation *without* a workload selector.
//
// For any namespace, including the root configuration namespace, it is only
// valid to have a single workload selector-less Telemetry resource.
//
// For resources with a workload selector, it is only valid to have one resource
// selecting any given workload.
//
// The hierarchy of Telemetry configuration is as follows:
//
// 1. Workload-specific configuration
// 2. Namespace-specific configuration
// 3. Root namespace configuration
//
// Examples:
//
// Policy to enable random sampling for 10% of traffic:
// ```yaml
// apiVersion: telemetry.istio.io/v1alpha1
// kind: Telemetry
// metadata:
//   name: mesh-default
//   namespace: istio-system
// spec:
//   # no selector specified, applies to all workloads
//   tracing:
//   - randomSamplingPercentage: 10.00
// ```
//
// Policy to disable trace reporting for the "foo" workload (note: tracing
// context will still be propagated):
// ```yaml
// apiVersion: telemetry.istio.io/v1alpha1
// kind: Telemetry
// metadata:
//   name: foo-tracing
//   namespace: bar
// spec:
//   selector:
//     matchLabels:
//       service.istio.io/canonical-name: foo
//   tracing:
//   - disableSpanReporting: true
// ```
//
// Policy to select the alternate zipkin provider for trace reporting:
// ```yaml
// apiVersion: telemetry.istio.io/v1alpha1
// kind: Telemetry
// metadata:
//   name: foo-tracing-alternate
//   namespace: baz
// spec:
//   selector:
//     matchLabels:
//       service.istio.io/canonical-name: foo
//   tracing:
//   - providers:
//     - name: "zipkin-alternate"
//     randomSamplingPercentage: 10.00
// ```
//
// Policy to add a custom tag from a literal value:
// ```yaml
// apiVersion: telemetry.istio.io/v1alpha1
// kind: Telemetry
// metadata:
//   name: mesh-default
//   namespace: istio-system
// spec:
//   # no selector specified, applies to all workloads
//   tracing:
//   - randomSamplingPercentage: 10.00
//     customTags:
//       my_new_foo_tag:
//         literal:
//           value: "foo"
// ```
//
// Policy to disable server-side metrics for Stackdriver for an entire mesh:
// ```yaml
// apiVersion: telemetry.istio.io/v1alpha1
// kind: Telemetry
// metadata:
//   name: mesh-default
//   namespace: istio-system
// spec:
//   # no selector specified, applies to all workloads
//   metrics:
//   - providers:
//     - name: stackdriver
//     overrides:
//     - match:
//         metric: ALL_METRICS
//         mode: SERVER
//       disabled: true
// ```
//
// Policy to add dimensions to all Prometheus metrics for the `foo` namespace:
// ```yaml
// apiVersion: telemetry.istio.io/v1alpha1
// kind: Telemetry
// metadata:
//   name: namespace-metrics
//   namespace: foo
// spec:
//   # no selector specified, applies to all workloads in the namespace
//   metrics:
//   - providers:
//     - name: prometheus
//     overrides:
//     # match clause left off matches all istio metrics, client and server
//     - tagOverrides:
//         request_method:
//           value: "request.method"
//         request_host:
//           value: "request.host"
// ```
//
// Policy to remove the response_code dimension on some Prometheus metrics for
// the `bar.foo` workload:
// ```yaml
// apiVersion: telemetry.istio.io/v1alpha1
// kind: Telemetry
// metadata:
//   name: remove-response-code
//   namespace: foo
// spec:
//   selector:
//     matchLabels:
//       service.istio.io/canonical-name: bar
//   metrics:
//   - providers:
//     - name: prometheus
//     overrides:
//     - match:
//         metric: REQUEST_COUNT
//       tagOverrides:
//         response_code:
//           operation: REMOVE
//     - match:
//         metric: REQUEST_DURATION
//       tagOverrides:
//         response_code:
//           operation: REMOVE
//     - match:
//         metric: REQUEST_BYTES
//       tagOverrides:
//         response_code:
//           operation: REMOVE
//     - match:
//         metric: RESPONSE_BYTES
//       tagOverrides:
//         response_code:
//           operation: REMOVE
// ```
//
// Policy to enable access logging for the entire mesh:
// ```yaml
// apiVersion: telemetry.istio.io/v1alpha1
// kind: Telemetry
// metadata:
//   name: mesh-default
//   namespace: istio-system
// spec:
//   # no selector specified, applies to all workloads
//   accessLogging:
//   - providers:
//     - name: envoy
//     # By default, this turns on access logging (no need to set `disabled:
//     false`). # Unspecified `disabled` will be treated as `disabled: false`,
//     except in # cases where a parent configuration has marked as `disabled:
//     true`. In # those cases, `disabled: false` must be set explicitly to
//     override.
// ```
//
// Policy to disable access logging for the `foo` namespace:
// ```yaml
// apiVersion: telemetry.istio.io/v1alpha1
// kind: Telemetry
// metadata:
//   name: namespace-no-log
//   namespace: foo
// spec:
//   # no selector specified, applies to all workloads in the namespace
//   accessLogging:
//   - disabled: true
// ```
//

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "istio.io/api/type/v1beta1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for Telemetry
func (this *Telemetry) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Telemetry
func (this *Telemetry) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Tracing
func (this *Tracing) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing
func (this *Tracing) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Tracing_TracingSelector
func (this *Tracing_TracingSelector) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing_TracingSelector
func (this *Tracing_TracingSelector) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Tracing_CustomTag
func (this *Tracing_CustomTag) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing_CustomTag
func (this *Tracing_CustomTag) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Tracing_Literal
func (this *Tracing_Literal) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing_Literal
func (this *Tracing_Literal) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Tracing_Environment
func (this *Tracing_Environment) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing_Environment
func (this *Tracing_Environment) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Tracing_RequestHeader
func (this *Tracing_RequestHeader) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing_RequestHeader
func (this *Tracing_RequestHeader) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ProviderRef
func (this *ProviderRef) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ProviderRef
func (this *ProviderRef) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Metrics
func (this *Metrics) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Metrics
func (this *Metrics) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MetricSelector
func (this *MetricSelector) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MetricSelector
func (this *MetricSelector) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MetricsOverrides
func (this *MetricsOverrides) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MetricsOverrides
func (this *MetricsOverrides) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MetricsOverrides_TagOverride
func (this *MetricsOverrides_TagOverride) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MetricsOverrides_TagOverride
func (this *MetricsOverrides_TagOverride) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for AccessLogging
func (this *AccessLogging) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AccessLogging
func (this *AccessLogging) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for AccessLogging_LogSelector
func (this *AccessLogging_LogSelector) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AccessLogging_LogSelector
func (this *AccessLogging_LogSelector) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for AccessLogging_Filter
func (this *AccessLogging_Filter) MarshalJSON() ([]byte, error) {
	str, err := TelemetryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AccessLogging_Filter
func (this *AccessLogging_Filter) UnmarshalJSON(b []byte) error {
	return TelemetryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	TelemetryMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	TelemetryUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{AllowUnknownFields: true}
)
