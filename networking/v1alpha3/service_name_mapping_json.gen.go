// Code generated by protoc-gen-jsonshim. DO NOT EDIT.
package v1alpha3

import (
	bytes "bytes"
	jsonpb "github.com/golang/protobuf/jsonpb"
)

// MarshalJSON is a custom marshaler for ServiceNameMapping
func (this *ServiceNameMapping) MarshalJSON() ([]byte, error) {
	str, err := ServiceNameMappingMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ServiceNameMapping
func (this *ServiceNameMapping) UnmarshalJSON(b []byte) error {
	return ServiceNameMappingUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ServiceNameMappingMarshaler   = &jsonpb.Marshaler{}
	ServiceNameMappingUnmarshaler = &jsonpb.Unmarshaler{AllowUnknownFields: true}
)