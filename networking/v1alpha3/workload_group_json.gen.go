// Code generated by protoc-gen-jsonshim. DO NOT EDIT.
package v1alpha3

import (
	bytes "bytes"
	jsonpb "github.com/golang/protobuf/jsonpb"
)

// MarshalJSON is a custom marshaler for WorkloadGroup
func (this *WorkloadGroup) MarshalJSON() ([]byte, error) {
	str, err := WorkloadGroupMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WorkloadGroup
func (this *WorkloadGroup) UnmarshalJSON(b []byte) error {
	return WorkloadGroupUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for WorkloadGroup_ObjectMeta
func (this *WorkloadGroup_ObjectMeta) MarshalJSON() ([]byte, error) {
	str, err := WorkloadGroupMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WorkloadGroup_ObjectMeta
func (this *WorkloadGroup_ObjectMeta) UnmarshalJSON(b []byte) error {
	return WorkloadGroupUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ReadinessProbe
func (this *ReadinessProbe) MarshalJSON() ([]byte, error) {
	str, err := WorkloadGroupMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ReadinessProbe
func (this *ReadinessProbe) UnmarshalJSON(b []byte) error {
	return WorkloadGroupUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPHealthCheckConfig
func (this *HTTPHealthCheckConfig) MarshalJSON() ([]byte, error) {
	str, err := WorkloadGroupMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPHealthCheckConfig
func (this *HTTPHealthCheckConfig) UnmarshalJSON(b []byte) error {
	return WorkloadGroupUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for GrpcHealthCheckConfig
func (this *GrpcHealthCheckConfig) MarshalJSON() ([]byte, error) {
	str, err := WorkloadGroupMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GrpcHealthCheckConfig
func (this *GrpcHealthCheckConfig) UnmarshalJSON(b []byte) error {
	return WorkloadGroupUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPHeader
func (this *HTTPHeader) MarshalJSON() ([]byte, error) {
	str, err := WorkloadGroupMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPHeader
func (this *HTTPHeader) UnmarshalJSON(b []byte) error {
	return WorkloadGroupUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TCPHealthCheckConfig
func (this *TCPHealthCheckConfig) MarshalJSON() ([]byte, error) {
	str, err := WorkloadGroupMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TCPHealthCheckConfig
func (this *TCPHealthCheckConfig) UnmarshalJSON(b []byte) error {
	return WorkloadGroupUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ExecHealthCheckConfig
func (this *ExecHealthCheckConfig) MarshalJSON() ([]byte, error) {
	str, err := WorkloadGroupMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ExecHealthCheckConfig
func (this *ExecHealthCheckConfig) UnmarshalJSON(b []byte) error {
	return WorkloadGroupUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	WorkloadGroupMarshaler   = &jsonpb.Marshaler{}
	WorkloadGroupUnmarshaler = &jsonpb.Unmarshaler{AllowUnknownFields: true}
)
