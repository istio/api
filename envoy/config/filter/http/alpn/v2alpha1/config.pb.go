// Copyright 2018 Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: envoy/config/filter/http/alpn/v2alpha1/config.proto

// $title: ALPN filter for overriding ALPN for upstream TLS connections.

package v2alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Upstream protocols
type FilterConfig_Protocol int32

const (
	FilterConfig_HTTP10 FilterConfig_Protocol = 0
	FilterConfig_HTTP11 FilterConfig_Protocol = 1
	FilterConfig_HTTP2  FilterConfig_Protocol = 2
)

// Enum value maps for FilterConfig_Protocol.
var (
	FilterConfig_Protocol_name = map[int32]string{
		0: "HTTP10",
		1: "HTTP11",
		2: "HTTP2",
	}
	FilterConfig_Protocol_value = map[string]int32{
		"HTTP10": 0,
		"HTTP11": 1,
		"HTTP2":  2,
	}
)

func (x FilterConfig_Protocol) Enum() *FilterConfig_Protocol {
	p := new(FilterConfig_Protocol)
	*p = x
	return p
}

func (x FilterConfig_Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FilterConfig_Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_config_filter_http_alpn_v2alpha1_config_proto_enumTypes[0].Descriptor()
}

func (FilterConfig_Protocol) Type() protoreflect.EnumType {
	return &file_envoy_config_filter_http_alpn_v2alpha1_config_proto_enumTypes[0]
}

func (x FilterConfig_Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FilterConfig_Protocol.Descriptor instead.
func (FilterConfig_Protocol) EnumDescriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDescGZIP(), []int{0, 0}
}

// FilterConfig is the config for Istio-specific filter.
type FilterConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Map from upstream protocol to list of ALPN
	AlpnOverride  []*FilterConfig_AlpnOverride `protobuf:"bytes,1,rep,name=alpn_override,json=alpnOverride,proto3" json:"alpn_override,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FilterConfig) Reset() {
	*x = FilterConfig{}
	mi := &file_envoy_config_filter_http_alpn_v2alpha1_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FilterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterConfig) ProtoMessage() {}

func (x *FilterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_alpn_v2alpha1_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterConfig.ProtoReflect.Descriptor instead.
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDescGZIP(), []int{0}
}

func (x *FilterConfig) GetAlpnOverride() []*FilterConfig_AlpnOverride {
	if x != nil {
		return x.AlpnOverride
	}
	return nil
}

type FilterConfig_AlpnOverride struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Upstream protocol
	UpstreamProtocol FilterConfig_Protocol `protobuf:"varint,1,opt,name=upstream_protocol,json=upstreamProtocol,proto3,enum=istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig_Protocol" json:"upstream_protocol,omitempty"`
	// A list of ALPN that will override the ALPN for upstream TLS connections.
	AlpnOverride  []string `protobuf:"bytes,2,rep,name=alpn_override,json=alpnOverride,proto3" json:"alpn_override,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FilterConfig_AlpnOverride) Reset() {
	*x = FilterConfig_AlpnOverride{}
	mi := &file_envoy_config_filter_http_alpn_v2alpha1_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FilterConfig_AlpnOverride) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterConfig_AlpnOverride) ProtoMessage() {}

func (x *FilterConfig_AlpnOverride) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_alpn_v2alpha1_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterConfig_AlpnOverride.ProtoReflect.Descriptor instead.
func (*FilterConfig_AlpnOverride) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FilterConfig_AlpnOverride) GetUpstreamProtocol() FilterConfig_Protocol {
	if x != nil {
		return x.UpstreamProtocol
	}
	return FilterConfig_HTTP10
}

func (x *FilterConfig_AlpnOverride) GetAlpnOverride() []string {
	if x != nil {
		return x.AlpnOverride
	}
	return nil
}

var File_envoy_config_filter_http_alpn_v2alpha1_config_proto protoreflect.FileDescriptor

const file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDesc = "" +
	"\n" +
	"3envoy/config/filter/http/alpn/v2alpha1/config.proto\x12,istio.envoy.config.filter.http.alpn.v2alpha1\"\xd3\x02\n" +
	"\fFilterConfig\x12l\n" +
	"\ralpn_override\x18\x01 \x03(\v2G.istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.AlpnOverrideR\falpnOverride\x1a\xa5\x01\n" +
	"\fAlpnOverride\x12p\n" +
	"\x11upstream_protocol\x18\x01 \x01(\x0e2C.istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.ProtocolR\x10upstreamProtocol\x12#\n" +
	"\ralpn_override\x18\x02 \x03(\tR\falpnOverride\"-\n" +
	"\bProtocol\x12\n" +
	"\n" +
	"\x06HTTP10\x10\x00\x12\n" +
	"\n" +
	"\x06HTTP11\x10\x01\x12\t\n" +
	"\x05HTTP2\x10\x02B5Z3istio.io/api/envoy/config/filter/http/alpn/v2alpha1b\x06proto3"

var (
	file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDescOnce sync.Once
	file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDescData []byte
)

func file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDescGZIP() []byte {
	file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDescOnce.Do(func() {
		file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDesc), len(file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDesc)))
	})
	return file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDescData
}

var file_envoy_config_filter_http_alpn_v2alpha1_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_config_filter_http_alpn_v2alpha1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_config_filter_http_alpn_v2alpha1_config_proto_goTypes = []any{
	(FilterConfig_Protocol)(0),        // 0: istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.Protocol
	(*FilterConfig)(nil),              // 1: istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig
	(*FilterConfig_AlpnOverride)(nil), // 2: istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.AlpnOverride
}
var file_envoy_config_filter_http_alpn_v2alpha1_config_proto_depIdxs = []int32{
	2, // 0: istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.alpn_override:type_name -> istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.AlpnOverride
	0, // 1: istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.AlpnOverride.upstream_protocol:type_name -> istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.Protocol
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_config_filter_http_alpn_v2alpha1_config_proto_init() }
func file_envoy_config_filter_http_alpn_v2alpha1_config_proto_init() {
	if File_envoy_config_filter_http_alpn_v2alpha1_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDesc), len(file_envoy_config_filter_http_alpn_v2alpha1_config_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_filter_http_alpn_v2alpha1_config_proto_goTypes,
		DependencyIndexes: file_envoy_config_filter_http_alpn_v2alpha1_config_proto_depIdxs,
		EnumInfos:         file_envoy_config_filter_http_alpn_v2alpha1_config_proto_enumTypes,
		MessageInfos:      file_envoy_config_filter_http_alpn_v2alpha1_config_proto_msgTypes,
	}.Build()
	File_envoy_config_filter_http_alpn_v2alpha1_config_proto = out.File
	file_envoy_config_filter_http_alpn_v2alpha1_config_proto_goTypes = nil
	file_envoy_config_filter_http_alpn_v2alpha1_config_proto_depIdxs = nil
}
