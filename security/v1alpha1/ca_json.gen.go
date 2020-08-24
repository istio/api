// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security/v1alpha1/ca.proto

// Keep this package for backward compatibility.

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "google/protobuf/struct.proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for IstioCertificateRequest
func (this *IstioCertificateRequest) MarshalJSON() ([]byte, error) {
	str, err := CaMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for IstioCertificateRequest
func (this *IstioCertificateRequest) UnmarshalJSON(b []byte) error {
	return CaUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for IstioCertificateResponse
func (this *IstioCertificateResponse) MarshalJSON() ([]byte, error) {
	str, err := CaMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for IstioCertificateResponse
func (this *IstioCertificateResponse) UnmarshalJSON(b []byte) error {
	return CaUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	CaMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	CaUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
