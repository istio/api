// Code generated by protoc-gen-jsonshim. DO NOT EDIT.
package v1alpha1

import (
	bytes "bytes"
	jsonpb "github.com/golang/protobuf/jsonpb"
)

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

// MarshalJSON is a custom marshaler for Roots
func (this *Roots) MarshalJSON() ([]byte, error) {
	str, err := CaMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Roots
func (this *Roots) UnmarshalJSON(b []byte) error {
	return CaUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	CaMarshaler   = &jsonpb.Marshaler{}
	CaUnmarshaler = &jsonpb.Unmarshaler{AllowUnknownFields: true}
)
