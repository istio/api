// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: extensions/v1alpha1/wasm.proto

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "istio.io/api/networking/v1alpha3"
	_ "istio.io/api/type/v1beta1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for WasmPlugin
func (this *WasmPlugin) MarshalJSON() ([]byte, error) {
	str, err := WasmMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WasmPlugin
func (this *WasmPlugin) UnmarshalJSON(b []byte) error {
	return WasmUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	WasmMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	WasmUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{AllowUnknownFields: true}
)
