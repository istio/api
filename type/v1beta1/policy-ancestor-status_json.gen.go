// Code generated by protoc-gen-jsonshim. DO NOT EDIT.
package v1beta1

import (
	bytes "bytes"
	jsonpb "github.com/golang/protobuf/jsonpb"
)

// MarshalJSON is a custom marshaler for ParentReference
func (this *ParentReference) MarshalJSON() ([]byte, error) {
	str, err := PolicyAncestorStatusMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ParentReference
func (this *ParentReference) UnmarshalJSON(b []byte) error {
	return PolicyAncestorStatusUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for PolicyAncestorStatus
func (this *PolicyAncestorStatus) MarshalJSON() ([]byte, error) {
	str, err := PolicyAncestorStatusMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for PolicyAncestorStatus
func (this *PolicyAncestorStatus) UnmarshalJSON(b []byte) error {
	return PolicyAncestorStatusUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Condition
func (this *Condition) MarshalJSON() ([]byte, error) {
	str, err := PolicyAncestorStatusMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Condition
func (this *Condition) UnmarshalJSON(b []byte) error {
	return PolicyAncestorStatusUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	PolicyAncestorStatusMarshaler   = &jsonpb.Marshaler{}
	PolicyAncestorStatusUnmarshaler = &jsonpb.Unmarshaler{AllowUnknownFields: true}
)
