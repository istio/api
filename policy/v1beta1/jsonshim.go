package v1beta1

import (
    "bytes"
    "github.com/gogo/protobuf/jsonpb"
)

var (
	customMarshaler   = &jsonpb.Marshaler{}
	customUnmarshaler = &jsonpb.Unmarshaler{}
)

// MarshalJSON is a custom marshaler supporting oneof fields for Instance
func (i *Instance) MarshalJSON() ([]byte, error) {
	str, err := customMarshaler.MarshalToString(i)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for Instance
func (i *Instance) UnmarshalJSON(b []byte) error {
	return customUnmarshaler.Unmarshal(bytes.NewReader(b), i)
}

// MarshalJSON is a custom marshaler supporting oneof fields for Handler
func (h *Handler) MarshalJSON() ([]byte, error) {
	str, err := customMarshaler.MarshalToString(h)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for Handler
func (h *Handler) UnmarshalJSON(b []byte) error {
	return customUnmarshaler.Unmarshal(bytes.NewReader(b), h)
}
