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
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: envoy/config/filter/http/authn/v2alpha1/config.proto

// $title: Internal API for authentication implementation on Envoy.

package v2alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1alpha1 "istio.io/api/authentication/v1alpha1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// FilterConfig is the config for Istio-specific filter that is used to enforce
// authentication policy on Envoy.
type FilterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Policy is the original copy of the policy.
	Policy *v1alpha1.Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	// Map from issuer to location of the payload that is emitted by Jwt filter.
	// This information is added by pilot when construct and add Jwt and
	// authN filters.
	JwtOutputPayloadLocations map[string]string `protobuf:"bytes,2,rep,name=jwt_output_payload_locations,json=jwtOutputPayloadLocations,proto3" json:"jwt_output_payload_locations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Skips validating the peer's trust domain.
	// By default, the istio authn filter will reject the request if the peer and
	// the local service is not in the same trust domain.
	// Set this field to true to skip the validation and allows peers from any
	// trust domains.
	// Note, the istio authn filter only validates the trust domain when mTLS is
	// used, In other words, this field has no effect for plaintext traffic.
	// TODO(incfly): deprecate this after allowed_trust_domains is shipped.
	SkipValidateTrustDomain bool `protobuf:"varint,3,opt,name=skip_validate_trust_domain,json=skipValidateTrustDomain,proto3" json:"skip_validate_trust_domain,omitempty"`
	// allowed_trust_domains contains a list of trust domains the authn
	// filter should validate against. When configured, only requests with a
	// peer from one of the allowed trust domain will be admitted.
	// An empty list means all trust domains are allowed.
	// When this field is set, the skip_validate_trust_domain field is ignored.
	// This field has no effect for plaintext traffic.
	AllowedTrustDomains []string `protobuf:"bytes,4,rep,name=allowed_trust_domains,json=allowedTrustDomains,proto3" json:"allowed_trust_domains,omitempty"`
	// By default the authn filter will clear the route cache so that the validated
	// JWT token claims can be used in routing.
	// Advanced users can set this to true to disable the behavior if they do not
	// want the authn filter to clear the route cache for any reasons.
	// Warning: setting this to true will break the JWT claim based routing.
	DisableClearRouteCache bool `protobuf:"varint,5,opt,name=disable_clear_route_cache,json=disableClearRouteCache,proto3" json:"disable_clear_route_cache,omitempty"`
}

func (x *FilterConfig) Reset() {
	*x = FilterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_http_authn_v2alpha1_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterConfig) ProtoMessage() {}

func (x *FilterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_authn_v2alpha1_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_envoy_config_filter_http_authn_v2alpha1_config_proto_rawDescGZIP(), []int{0}
}

func (x *FilterConfig) GetPolicy() *v1alpha1.Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

func (x *FilterConfig) GetJwtOutputPayloadLocations() map[string]string {
	if x != nil {
		return x.JwtOutputPayloadLocations
	}
	return nil
}

func (x *FilterConfig) GetSkipValidateTrustDomain() bool {
	if x != nil {
		return x.SkipValidateTrustDomain
	}
	return false
}

func (x *FilterConfig) GetAllowedTrustDomains() []string {
	if x != nil {
		return x.AllowedTrustDomains
	}
	return nil
}

func (x *FilterConfig) GetDisableClearRouteCache() bool {
	if x != nil {
		return x.DisableClearRouteCache
	}
	return false
}

var File_envoy_config_filter_http_authn_v2alpha1_config_proto protoreflect.FileDescriptor

var file_envoy_config_filter_http_authn_v2alpha1_config_proto_rawDesc = []byte{
	0x0a, 0x34, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6e,
	0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x24, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x03, 0x0a, 0x0c,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3d, 0x0a, 0x06,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69,
	0x73, 0x74, 0x69, 0x6f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x9b, 0x01, 0x0a, 0x1c,
	0x6a, 0x77, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x5a, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x4a, 0x77, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x19,
	0x6a, 0x77, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3b, 0x0a, 0x1a, 0x73, 0x6b, 0x69,
	0x70, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x72, 0x75, 0x73, 0x74,
	0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x73,
	0x6b, 0x69, 0x70, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x75, 0x73, 0x74,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x32, 0x0a, 0x15, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x5f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x54, 0x72,
	0x75, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x39, 0x0a, 0x19, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x1a, 0x4c, 0x0a, 0x1e, 0x4a, 0x77, 0x74, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x36, 0x5a, 0x34, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x6e, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_filter_http_authn_v2alpha1_config_proto_rawDescOnce sync.Once
	file_envoy_config_filter_http_authn_v2alpha1_config_proto_rawDescData = file_envoy_config_filter_http_authn_v2alpha1_config_proto_rawDesc
)

func file_envoy_config_filter_http_authn_v2alpha1_config_proto_rawDescGZIP() []byte {
	file_envoy_config_filter_http_authn_v2alpha1_config_proto_rawDescOnce.Do(func() {
		file_envoy_config_filter_http_authn_v2alpha1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_filter_http_authn_v2alpha1_config_proto_rawDescData)
	})
	return file_envoy_config_filter_http_authn_v2alpha1_config_proto_rawDescData
}

var file_envoy_config_filter_http_authn_v2alpha1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_config_filter_http_authn_v2alpha1_config_proto_goTypes = []interface{}{
	(*FilterConfig)(nil),    // 0: istio.envoy.config.filter.http.authn.v2alpha1.FilterConfig
	nil,                     // 1: istio.envoy.config.filter.http.authn.v2alpha1.FilterConfig.JwtOutputPayloadLocationsEntry
	(*v1alpha1.Policy)(nil), // 2: istio.authentication.v1alpha1.Policy
}
var file_envoy_config_filter_http_authn_v2alpha1_config_proto_depIdxs = []int32{
	2, // 0: istio.envoy.config.filter.http.authn.v2alpha1.FilterConfig.policy:type_name -> istio.authentication.v1alpha1.Policy
	1, // 1: istio.envoy.config.filter.http.authn.v2alpha1.FilterConfig.jwt_output_payload_locations:type_name -> istio.envoy.config.filter.http.authn.v2alpha1.FilterConfig.JwtOutputPayloadLocationsEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_config_filter_http_authn_v2alpha1_config_proto_init() }
func file_envoy_config_filter_http_authn_v2alpha1_config_proto_init() {
	if File_envoy_config_filter_http_authn_v2alpha1_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_filter_http_authn_v2alpha1_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterConfig); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_filter_http_authn_v2alpha1_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_filter_http_authn_v2alpha1_config_proto_goTypes,
		DependencyIndexes: file_envoy_config_filter_http_authn_v2alpha1_config_proto_depIdxs,
		MessageInfos:      file_envoy_config_filter_http_authn_v2alpha1_config_proto_msgTypes,
	}.Build()
	File_envoy_config_filter_http_authn_v2alpha1_config_proto = out.File
	file_envoy_config_filter_http_authn_v2alpha1_config_proto_rawDesc = nil
	file_envoy_config_filter_http_authn_v2alpha1_config_proto_goTypes = nil
	file_envoy_config_filter_http_authn_v2alpha1_config_proto_depIdxs = nil
}
