// Copyright Istio Authors
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
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: security/v1alpha1/ca.proto

// Keep this package for backward compatibility.

package v1alpha1

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// Certificate request message. The authentication should be based on:
// 1. Bearer tokens carried in the side channel;
// 2. Client-side certificate via Mutual TLS handshake.
// Note: the service implementation is REQUIRED to verify the authenticated caller is authorize to
// all SANs in the CSR. The server side may overwrite any requested certificate field based on its
// policies.
type IstioCertificateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// PEM-encoded certificate request.
	// The public key in the CSR is used to generate the certificate,
	// and other fields in the generated certificate may be overwritten by the CA.
	Csr string `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
	// Optional: requested certificate validity period, in seconds.
	ValidityDuration int64 `protobuf:"varint,3,opt,name=validity_duration,json=validityDuration,proto3" json:"validity_duration,omitempty"`
	// $hide_from_docs
	// Optional: Opaque metadata provided by the XDS node to Istio.
	// Supported metadata: WorkloadName, WorkloadIP, ClusterID
	Metadata      *_struct.Struct `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IstioCertificateRequest) Reset() {
	*x = IstioCertificateRequest{}
	mi := &file_security_v1alpha1_ca_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IstioCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IstioCertificateRequest) ProtoMessage() {}

func (x *IstioCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_security_v1alpha1_ca_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IstioCertificateRequest.ProtoReflect.Descriptor instead.
func (*IstioCertificateRequest) Descriptor() ([]byte, []int) {
	return file_security_v1alpha1_ca_proto_rawDescGZIP(), []int{0}
}

func (x *IstioCertificateRequest) GetCsr() string {
	if x != nil {
		return x.Csr
	}
	return ""
}

func (x *IstioCertificateRequest) GetValidityDuration() int64 {
	if x != nil {
		return x.ValidityDuration
	}
	return 0
}

func (x *IstioCertificateRequest) GetMetadata() *_struct.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Certificate response message.
type IstioCertificateResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// PEM-encoded certificate chain.
	// The leaf cert is the first element, and the root cert is the last element.
	CertChain     []string `protobuf:"bytes,1,rep,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IstioCertificateResponse) Reset() {
	*x = IstioCertificateResponse{}
	mi := &file_security_v1alpha1_ca_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IstioCertificateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IstioCertificateResponse) ProtoMessage() {}

func (x *IstioCertificateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_security_v1alpha1_ca_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IstioCertificateResponse.ProtoReflect.Descriptor instead.
func (*IstioCertificateResponse) Descriptor() ([]byte, []int) {
	return file_security_v1alpha1_ca_proto_rawDescGZIP(), []int{1}
}

func (x *IstioCertificateResponse) GetCertChain() []string {
	if x != nil {
		return x.CertChain
	}
	return nil
}

var File_security_v1alpha1_ca_proto protoreflect.FileDescriptor

var file_security_v1alpha1_ca_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x69, 0x73,
	0x74, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x17, 0x49, 0x73,
	0x74, 0x69, 0x6f, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x73, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x63, 0x73, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x39, 0x0a, 0x18, 0x49, 0x73, 0x74,
	0x69, 0x6f, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x63, 0x65, 0x72, 0x74, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x32, 0x81, 0x01, 0x0a, 0x17, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x66, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x49, 0x73,
	0x74, 0x69, 0x6f, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x69, 0x73, 0x74, 0x69,
	0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_security_v1alpha1_ca_proto_rawDescOnce sync.Once
	file_security_v1alpha1_ca_proto_rawDescData = file_security_v1alpha1_ca_proto_rawDesc
)

func file_security_v1alpha1_ca_proto_rawDescGZIP() []byte {
	file_security_v1alpha1_ca_proto_rawDescOnce.Do(func() {
		file_security_v1alpha1_ca_proto_rawDescData = protoimpl.X.CompressGZIP(file_security_v1alpha1_ca_proto_rawDescData)
	})
	return file_security_v1alpha1_ca_proto_rawDescData
}

var file_security_v1alpha1_ca_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_security_v1alpha1_ca_proto_goTypes = []any{
	(*IstioCertificateRequest)(nil),  // 0: istio.v1.auth.IstioCertificateRequest
	(*IstioCertificateResponse)(nil), // 1: istio.v1.auth.IstioCertificateResponse
	(*_struct.Struct)(nil),           // 2: google.protobuf.Struct
}
var file_security_v1alpha1_ca_proto_depIdxs = []int32{
	2, // 0: istio.v1.auth.IstioCertificateRequest.metadata:type_name -> google.protobuf.Struct
	0, // 1: istio.v1.auth.IstioCertificateService.CreateCertificate:input_type -> istio.v1.auth.IstioCertificateRequest
	1, // 2: istio.v1.auth.IstioCertificateService.CreateCertificate:output_type -> istio.v1.auth.IstioCertificateResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_security_v1alpha1_ca_proto_init() }
func file_security_v1alpha1_ca_proto_init() {
	if File_security_v1alpha1_ca_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_security_v1alpha1_ca_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_security_v1alpha1_ca_proto_goTypes,
		DependencyIndexes: file_security_v1alpha1_ca_proto_depIdxs,
		MessageInfos:      file_security_v1alpha1_ca_proto_msgTypes,
	}.Build()
	File_security_v1alpha1_ca_proto = out.File
	file_security_v1alpha1_ca_proto_rawDesc = nil
	file_security_v1alpha1_ca_proto_goTypes = nil
	file_security_v1alpha1_ca_proto_depIdxs = nil
}
