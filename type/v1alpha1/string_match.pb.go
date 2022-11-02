// Copyright 2019 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: type/v1alpha1/string_match.proto

// $title: String Match
// $description: Definition how to match a given string.
// $location: https://istio.io/docs/reference/config/type/string-match.html

package v1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Describes how to match a given string. Match is case-sensitive.
type StringMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MatchType:
	//
	//	*StringMatch_Exact
	//	*StringMatch_Prefix
	//	*StringMatch_Regex
	MatchType isStringMatch_MatchType `protobuf_oneof:"match_type"`
}

func (x *StringMatch) Reset() {
	*x = StringMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_type_v1alpha1_string_match_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringMatch) ProtoMessage() {}

func (x *StringMatch) ProtoReflect() protoreflect.Message {
	mi := &file_type_v1alpha1_string_match_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringMatch.ProtoReflect.Descriptor instead.
func (*StringMatch) Descriptor() ([]byte, []int) {
	return file_type_v1alpha1_string_match_proto_rawDescGZIP(), []int{0}
}

func (m *StringMatch) GetMatchType() isStringMatch_MatchType {
	if m != nil {
		return m.MatchType
	}
	return nil
}

func (x *StringMatch) GetExact() string {
	if x, ok := x.GetMatchType().(*StringMatch_Exact); ok {
		return x.Exact
	}
	return ""
}

func (x *StringMatch) GetPrefix() string {
	if x, ok := x.GetMatchType().(*StringMatch_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (x *StringMatch) GetRegex() string {
	if x, ok := x.GetMatchType().(*StringMatch_Regex); ok {
		return x.Regex
	}
	return ""
}

type isStringMatch_MatchType interface {
	isStringMatch_MatchType()
}

type StringMatch_Exact struct {
	// exact string match
	Exact string `protobuf:"bytes,1,opt,name=exact,proto3,oneof"`
}

type StringMatch_Prefix struct {
	// prefix-based match
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof"`
}

type StringMatch_Regex struct {
	// RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
	Regex string `protobuf:"bytes,3,opt,name=regex,proto3,oneof"`
}

func (*StringMatch_Exact) isStringMatch_MatchType() {}

func (*StringMatch_Prefix) isStringMatch_MatchType() {}

func (*StringMatch_Regex) isStringMatch_MatchType() {}

var File_type_v1alpha1_string_match_proto protoreflect.FileDescriptor

var file_type_v1alpha1_string_match_proto_rawDesc = []byte{
	0x0a, 0x20, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x65, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x78, 0x61, 0x63, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x78, 0x61, 0x63, 0x74, 0x12, 0x18,
	0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x16, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65,
	0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78,
	0x42, 0x0c, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x1c,
	0x5a, 0x1a, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_type_v1alpha1_string_match_proto_rawDescOnce sync.Once
	file_type_v1alpha1_string_match_proto_rawDescData = file_type_v1alpha1_string_match_proto_rawDesc
)

func file_type_v1alpha1_string_match_proto_rawDescGZIP() []byte {
	file_type_v1alpha1_string_match_proto_rawDescOnce.Do(func() {
		file_type_v1alpha1_string_match_proto_rawDescData = protoimpl.X.CompressGZIP(file_type_v1alpha1_string_match_proto_rawDescData)
	})
	return file_type_v1alpha1_string_match_proto_rawDescData
}

var file_type_v1alpha1_string_match_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_type_v1alpha1_string_match_proto_goTypes = []interface{}{
	(*StringMatch)(nil), // 0: istio.type.v1alpha1.StringMatch
}
var file_type_v1alpha1_string_match_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_type_v1alpha1_string_match_proto_init() }
func file_type_v1alpha1_string_match_proto_init() {
	if File_type_v1alpha1_string_match_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_type_v1alpha1_string_match_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringMatch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_type_v1alpha1_string_match_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StringMatch_Exact)(nil),
		(*StringMatch_Prefix)(nil),
		(*StringMatch_Regex)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_type_v1alpha1_string_match_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_type_v1alpha1_string_match_proto_goTypes,
		DependencyIndexes: file_type_v1alpha1_string_match_proto_depIdxs,
		MessageInfos:      file_type_v1alpha1_string_match_proto_msgTypes,
	}.Build()
	File_type_v1alpha1_string_match_proto = out.File
	file_type_v1alpha1_string_match_proto_rawDesc = nil
	file_type_v1alpha1_string_match_proto_goTypes = nil
	file_type_v1alpha1_string_match_proto_depIdxs = nil
}
