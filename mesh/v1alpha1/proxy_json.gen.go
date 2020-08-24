// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mesh/v1alpha1/proxy.proto

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "google/protobuf/duration.proto"
	_ "google/protobuf/wrappers.proto"
	_ "istio.io/api/networking/v1alpha3"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for Tracing
func (this *Tracing) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing
func (this *Tracing) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Tracing_Zipkin
func (this *Tracing_Zipkin) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing_Zipkin
func (this *Tracing_Zipkin) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Tracing_Lightstep
func (this *Tracing_Lightstep) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing_Lightstep
func (this *Tracing_Lightstep) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Tracing_Datadog
func (this *Tracing_Datadog) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing_Datadog
func (this *Tracing_Datadog) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Tracing_Stackdriver
func (this *Tracing_Stackdriver) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing_Stackdriver
func (this *Tracing_Stackdriver) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Tracing_OpenCensusAgent
func (this *Tracing_OpenCensusAgent) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing_OpenCensusAgent
func (this *Tracing_OpenCensusAgent) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Tracing_CustomTag
func (this *Tracing_CustomTag) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing_CustomTag
func (this *Tracing_CustomTag) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Tracing_Literal
func (this *Tracing_Literal) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing_Literal
func (this *Tracing_Literal) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Tracing_Environment
func (this *Tracing_Environment) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing_Environment
func (this *Tracing_Environment) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Tracing_RequestHeader
func (this *Tracing_RequestHeader) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Tracing_RequestHeader
func (this *Tracing_RequestHeader) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for SDS
func (this *SDS) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for SDS
func (this *SDS) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Topology
func (this *Topology) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Topology
func (this *Topology) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ProxyConfig
func (this *ProxyConfig) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ProxyConfig
func (this *ProxyConfig) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RemoteService
func (this *RemoteService) MarshalJSON() ([]byte, error) {
	str, err := ProxyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RemoteService
func (this *RemoteService) UnmarshalJSON(b []byte) error {
	return ProxyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ProxyMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	ProxyUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
