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
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: type/v1beta1/selector.proto

// $title: Workload Selector
// $description: Definition of a workload selector.
// $location: https://istio.io/docs/reference/config/type/workload-selector.html

package v1beta1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// WorkloadMode allows selection of the role of the underlying workload in
// network traffic. A workload is considered as acting as a SERVER if it is
// the destination of the traffic (that is, traffic direction, from the
// perspective of the workload is *inbound*). If the workload is the source of
// the network traffic, it is considered to be in CLIENT mode (traffic is
// *outbound* from the workload).
type WorkloadMode int32

const (
	// Default value, which will be interpreted by its own usage.
	WorkloadMode_UNDEFINED WorkloadMode = 0
	// Selects for scenarios when the workload is the
	// source of the network traffic. In addition,
	// if the workload is a gateway, selects this.
	WorkloadMode_CLIENT WorkloadMode = 1
	// Selects for scenarios when the workload is the
	// destination of the network traffic.
	WorkloadMode_SERVER WorkloadMode = 2
	// Selects for scenarios when the workload is either the
	// source or destination of the network traffic.
	WorkloadMode_CLIENT_AND_SERVER WorkloadMode = 3
)

// Enum value maps for WorkloadMode.
var (
	WorkloadMode_name = map[int32]string{
		0: "UNDEFINED",
		1: "CLIENT",
		2: "SERVER",
		3: "CLIENT_AND_SERVER",
	}
	WorkloadMode_value = map[string]int32{
		"UNDEFINED":         0,
		"CLIENT":            1,
		"SERVER":            2,
		"CLIENT_AND_SERVER": 3,
	}
)

func (x WorkloadMode) Enum() *WorkloadMode {
	p := new(WorkloadMode)
	*p = x
	return p
}

func (x WorkloadMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkloadMode) Descriptor() protoreflect.EnumDescriptor {
	return file_type_v1beta1_selector_proto_enumTypes[0].Descriptor()
}

func (WorkloadMode) Type() protoreflect.EnumType {
	return &file_type_v1beta1_selector_proto_enumTypes[0]
}

func (x WorkloadMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkloadMode.Descriptor instead.
func (WorkloadMode) EnumDescriptor() ([]byte, []int) {
	return file_type_v1beta1_selector_proto_rawDescGZIP(), []int{0}
}

// WorkloadSelector specifies the criteria used to determine if a policy can be applied
// to a proxy. The matching criteria includes the metadata associated with a proxy,
// workload instance info such as labels attached to the pod/VM, or any other info
// that the proxy provides to Istio during the initial handshake. If multiple conditions are
// specified, all conditions need to match in order for the workload instance to be
// selected. Currently, only label based selection mechanism is supported.
type WorkloadSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// One or more labels that indicate a specific set of pods/VMs
	// on which a policy should be applied. The scope of label search is restricted to
	// the configuration namespace in which the resource is present.
	MatchLabels map[string]string `protobuf:"bytes,1,rep,name=match_labels,json=matchLabels,proto3" json:"match_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *WorkloadSelector) Reset() {
	*x = WorkloadSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_type_v1beta1_selector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadSelector) ProtoMessage() {}

func (x *WorkloadSelector) ProtoReflect() protoreflect.Message {
	mi := &file_type_v1beta1_selector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadSelector.ProtoReflect.Descriptor instead.
func (*WorkloadSelector) Descriptor() ([]byte, []int) {
	return file_type_v1beta1_selector_proto_rawDescGZIP(), []int{0}
}

func (x *WorkloadSelector) GetMatchLabels() map[string]string {
	if x != nil {
		return x.MatchLabels
	}
	return nil
}

// PortSelector is the criteria for specifying if a policy can be applied to
// a listener having a specific port.
type PortSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Port number
	Number uint32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PortSelector) Reset() {
	*x = PortSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_type_v1beta1_selector_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortSelector) ProtoMessage() {}

func (x *PortSelector) ProtoReflect() protoreflect.Message {
	mi := &file_type_v1beta1_selector_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortSelector.ProtoReflect.Descriptor instead.
func (*PortSelector) Descriptor() ([]byte, []int) {
	return file_type_v1beta1_selector_proto_rawDescGZIP(), []int{1}
}

func (x *PortSelector) GetNumber() uint32 {
	if x != nil {
		return x.Number
	}
	return 0
}

// PolicyTargetReferences specifies the targeted resource which the policy
// can be applied to. It must only target a single resource at a time, but it
// can be used to target larger resources such as Gateways that may apply to
// multiple child resources. The PolicyTargetReference will replace the
// WorkloadSelector for AuthorizationPolicies applied to ambient workloads.
type PolicyTargetReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// group is the group of the target resource.
	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	// kind is kind of the target resource.
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// name is the name of the target resource.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// namespace is the namespace of the referent. When unspecified, the local
	// namespace is inferred. Even when policy targets a resource in a different
	// namespace, it may only apply to traffic originating from the same
	// namespace as the policy.
	Namespace *string `protobuf:"bytes,4,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
}

func (x *PolicyTargetReference) Reset() {
	*x = PolicyTargetReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_type_v1beta1_selector_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyTargetReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyTargetReference) ProtoMessage() {}

func (x *PolicyTargetReference) ProtoReflect() protoreflect.Message {
	mi := &file_type_v1beta1_selector_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyTargetReference.ProtoReflect.Descriptor instead.
func (*PolicyTargetReference) Descriptor() ([]byte, []int) {
	return file_type_v1beta1_selector_proto_rawDescGZIP(), []int{2}
}

func (x *PolicyTargetReference) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *PolicyTargetReference) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *PolicyTargetReference) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PolicyTargetReference) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

var File_type_v1beta1_selector_proto protoreflect.FileDescriptor

var file_type_v1beta1_selector_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x69,
	0x73, 0x74, 0x69, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x5d, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e,
	0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x26, 0x0a, 0x0c, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x86,
	0x01, 0x0a, 0x15, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2a, 0x4c, 0x0a, 0x0c, 0x57, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46,
	0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x02, 0x12, 0x15,
	0x0a, 0x11, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x45, 0x52, 0x10, 0x03, 0x42, 0x1b, 0x5a, 0x19, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_type_v1beta1_selector_proto_rawDescOnce sync.Once
	file_type_v1beta1_selector_proto_rawDescData = file_type_v1beta1_selector_proto_rawDesc
)

func file_type_v1beta1_selector_proto_rawDescGZIP() []byte {
	file_type_v1beta1_selector_proto_rawDescOnce.Do(func() {
		file_type_v1beta1_selector_proto_rawDescData = protoimpl.X.CompressGZIP(file_type_v1beta1_selector_proto_rawDescData)
	})
	return file_type_v1beta1_selector_proto_rawDescData
}

var file_type_v1beta1_selector_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_type_v1beta1_selector_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_type_v1beta1_selector_proto_goTypes = []interface{}{
	(WorkloadMode)(0),             // 0: istio.type.v1beta1.WorkloadMode
	(*WorkloadSelector)(nil),      // 1: istio.type.v1beta1.WorkloadSelector
	(*PortSelector)(nil),          // 2: istio.type.v1beta1.PortSelector
	(*PolicyTargetReference)(nil), // 3: istio.type.v1beta1.PolicyTargetReference
	nil,                           // 4: istio.type.v1beta1.WorkloadSelector.MatchLabelsEntry
}
var file_type_v1beta1_selector_proto_depIdxs = []int32{
	4, // 0: istio.type.v1beta1.WorkloadSelector.match_labels:type_name -> istio.type.v1beta1.WorkloadSelector.MatchLabelsEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_type_v1beta1_selector_proto_init() }
func file_type_v1beta1_selector_proto_init() {
	if File_type_v1beta1_selector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_type_v1beta1_selector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadSelector); i {
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
		file_type_v1beta1_selector_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortSelector); i {
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
		file_type_v1beta1_selector_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyTargetReference); i {
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
	file_type_v1beta1_selector_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_type_v1beta1_selector_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_type_v1beta1_selector_proto_goTypes,
		DependencyIndexes: file_type_v1beta1_selector_proto_depIdxs,
		EnumInfos:         file_type_v1beta1_selector_proto_enumTypes,
		MessageInfos:      file_type_v1beta1_selector_proto_msgTypes,
	}.Build()
	File_type_v1beta1_selector_proto = out.File
	file_type_v1beta1_selector_proto_rawDesc = nil
	file_type_v1beta1_selector_proto_goTypes = nil
	file_type_v1beta1_selector_proto_depIdxs = nil
}
