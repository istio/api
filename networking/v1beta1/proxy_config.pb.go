// Copyright 2021 Istio Authors
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
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: networking/v1beta1/proxy_config.proto

// $schema: istio.networking.v1beta1.ProxyConfig
// $title: ProxyConfig
// $description: Provides configuration for individual workloads.
// $location: https://istio.io/docs/reference/config/networking/proxy-config.html
// $aliases: [/docs/reference/config/networking/v1beta1/proxy-config]
// $mode: file

// `ProxyConfig` exposes proxy level configuration options. `ProxyConfig` can be configured on a per-workload basis,
// a per-namespace basis, or mesh-wide. `ProxyConfig` is not a required resource; there are default values in place, which are documented
// inline with each field.
//
// **NOTE**: fields in ProxyConfig are not dynamically configured - changes will require restart of workloads to take effect.
//
// For any namespace, including the root configuration namespace, it is only valid
// to have a single workload selector-less `ProxyConfig` resource.
//
// For resources with a workload selector, it is only valid to have one resource selecting
// any given workload.
//
// For mesh level configuration, put the resource in the root configuration namespace for
// your Istio installation *without* a workload selector:
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ProxyConfig
// metadata:
//   name: my-proxyconfig
//   namespace: istio-system
// spec:
//   concurrency: 0
//   image:
//     imageType: distroless
// ```
//
// For namespace level configuration, put the resource in the desired namespace without a workload selector:
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ProxyConfig
// metadata:
//   name: my-ns-proxyconfig
//   namespace: user-namespace
// spec:
//   concurrency: 0
// ```
//
// For workload level configuration, set the `selector` field on the `ProxyConfig` resource:
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ProxyConfig
// metadata:
//   name: per-workload-proxyconfig
//   namespace: example
// spec:
//   selector:
//     matchLabels:
//       app: ratings
//   concurrency: 0
//   image:
//     imageType: debug
// ```
//
// To preserve the header case for HTTP 1.x requests, set the `preserveCase` field on the `ProxyConfig` resource:
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ProxyConfig
// metadata:
//   name: my-ns-proxyconfig
//   namespace: user-namespace
// spec:
//   preserveCase: true
// ```
//
// If a `ProxyConfig` CR is defined that matches a workload it will merge with its `proxy.istio.io/config` annotation if present,
// with the CR taking precedence over the annotation for overlapping fields. Similarly, if a mesh wide `ProxyConfig` CR is defined and
// `meshConfig.DefaultConfig` is set, the two resources will be merged with the CR taking precedence for overlapping fields.
//

package v1beta1

import (
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1beta1 "istio.io/api/type/v1beta1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// `ProxyConfig` exposes proxy level configuration options.
//
// <!-- crd generation tags
// +cue-gen:ProxyConfig:groupName:networking.istio.io
// +cue-gen:ProxyConfig:versions:v1beta1
// +cue-gen:ProxyConfig:storageVersion
// +cue-gen:ProxyConfig:annotations:helm.sh/resource-policy=keep
// +cue-gen:ProxyConfig:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:ProxyConfig:subresource:status
// +cue-gen:ProxyConfig:scope:Namespaced
// +cue-gen:WasmPlugin:releaseChannel:extended
// +cue-gen:ProxyConfig:resource:categories=istio-io,networking-istio-io,plural=proxyconfigs
// +cue-gen:ProxyConfig:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.istio.io/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type ProxyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Selectors specify the set of pods/VMs on which this `ProxyConfig` resource should be applied.
	// If not set, the `ProxyConfig` resource will be applied to all workloads in the namespace where this resource is defined.
	Selector *v1beta1.WorkloadSelector `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// The number of worker threads to run.
	// If unset, this will be automatically determined based on CPU limits.
	// If set to 0, all cores on the machine will be used.
	// +kubebuilder:validation:Minimum=0
	Concurrency *wrappers.Int32Value `protobuf:"bytes,2,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	// Additional environment variables for the proxy.
	// Names starting with `ISTIO_META_` will be included in the generated bootstrap configuration and sent to the XDS server.
	// +protoc-gen-crd:map-value-validation:MaxLength=2048
	EnvironmentVariables map[string]string `protobuf:"bytes,3,rep,name=environment_variables,json=environmentVariables,proto3" json:"environment_variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Specifies the details of the proxy image.
	Image *ProxyImage `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	// When true, the original case of HTTP/1.x headers will be preserved
	// as they pass through the proxy, rather than normalizing them to lowercase.
	// This field is particularly useful for applications that require case-sensitive
	// headers for interoperability with downstream systems or APIs that expect specific
	// casing.
	// The preserve_case option only applies to HTTP/1.x traffic, as HTTP/2 requires all headers
	// to be lowercase per the protocol specification. Envoy will ignore this field for HTTP/2
	// requests and automatically normalize headers to lowercase, ensuring compliance with HTTP/2
	// standards.
	PreserveCase bool `protobuf:"varint,5,opt,name=preserve_case,json=preserveCase,proto3" json:"preserve_case,omitempty"`
}

func (x *ProxyConfig) Reset() {
	*x = ProxyConfig{}
	mi := &file_networking_v1beta1_proxy_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProxyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyConfig) ProtoMessage() {}

func (x *ProxyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_networking_v1beta1_proxy_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyConfig.ProtoReflect.Descriptor instead.
func (*ProxyConfig) Descriptor() ([]byte, []int) {
	return file_networking_v1beta1_proxy_config_proto_rawDescGZIP(), []int{0}
}

func (x *ProxyConfig) GetSelector() *v1beta1.WorkloadSelector {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *ProxyConfig) GetConcurrency() *wrappers.Int32Value {
	if x != nil {
		return x.Concurrency
	}
	return nil
}

func (x *ProxyConfig) GetEnvironmentVariables() map[string]string {
	if x != nil {
		return x.EnvironmentVariables
	}
	return nil
}

func (x *ProxyConfig) GetImage() *ProxyImage {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *ProxyConfig) GetPreserveCase() bool {
	if x != nil {
		return x.PreserveCase
	}
	return false
}

// The following values are used to construct proxy image url.
// format: `${hub}/${image_name}/${tag}-${image_type}`,
// example: `docker.io/istio/proxyv2:1.11.1` or `docker.io/istio/proxyv2:1.11.1-distroless`.
// This information was previously part of the Values API.
type ProxyImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The image type of the image.
	// Istio publishes default, debug, and distroless images.
	// Other values are allowed if those image types (example: centos) are published to the specified hub.
	// supported values: default, debug, distroless.
	ImageType string `protobuf:"bytes,1,opt,name=image_type,json=imageType,proto3" json:"image_type,omitempty"`
}

func (x *ProxyImage) Reset() {
	*x = ProxyImage{}
	mi := &file_networking_v1beta1_proxy_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProxyImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyImage) ProtoMessage() {}

func (x *ProxyImage) ProtoReflect() protoreflect.Message {
	mi := &file_networking_v1beta1_proxy_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyImage.ProtoReflect.Descriptor instead.
func (*ProxyImage) Descriptor() ([]byte, []int) {
	return file_networking_v1beta1_proxy_config_proto_rawDescGZIP(), []int{1}
}

func (x *ProxyImage) GetImageType() string {
	if x != nil {
		return x.ImageType
	}
	return ""
}

var File_networking_v1beta1_proxy_config_proto protoreflect.FileDescriptor

var file_networking_v1beta1_proxy_config_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae,
	0x03, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x40,
	0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x3d, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x74, 0x0a, 0x15, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f,
	0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x14, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x5f, 0x63, 0x61,
	0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x43, 0x61, 0x73, 0x65, 0x1a, 0x47, 0x0a, 0x19, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x2b, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x21, 0x5a, 0x1f,
	0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_networking_v1beta1_proxy_config_proto_rawDescOnce sync.Once
	file_networking_v1beta1_proxy_config_proto_rawDescData = file_networking_v1beta1_proxy_config_proto_rawDesc
)

func file_networking_v1beta1_proxy_config_proto_rawDescGZIP() []byte {
	file_networking_v1beta1_proxy_config_proto_rawDescOnce.Do(func() {
		file_networking_v1beta1_proxy_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_networking_v1beta1_proxy_config_proto_rawDescData)
	})
	return file_networking_v1beta1_proxy_config_proto_rawDescData
}

var file_networking_v1beta1_proxy_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_networking_v1beta1_proxy_config_proto_goTypes = []any{
	(*ProxyConfig)(nil),              // 0: istio.networking.v1beta1.ProxyConfig
	(*ProxyImage)(nil),               // 1: istio.networking.v1beta1.ProxyImage
	nil,                              // 2: istio.networking.v1beta1.ProxyConfig.EnvironmentVariablesEntry
	(*v1beta1.WorkloadSelector)(nil), // 3: istio.type.v1beta1.WorkloadSelector
	(*wrappers.Int32Value)(nil),      // 4: google.protobuf.Int32Value
}
var file_networking_v1beta1_proxy_config_proto_depIdxs = []int32{
	3, // 0: istio.networking.v1beta1.ProxyConfig.selector:type_name -> istio.type.v1beta1.WorkloadSelector
	4, // 1: istio.networking.v1beta1.ProxyConfig.concurrency:type_name -> google.protobuf.Int32Value
	2, // 2: istio.networking.v1beta1.ProxyConfig.environment_variables:type_name -> istio.networking.v1beta1.ProxyConfig.EnvironmentVariablesEntry
	1, // 3: istio.networking.v1beta1.ProxyConfig.image:type_name -> istio.networking.v1beta1.ProxyImage
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_networking_v1beta1_proxy_config_proto_init() }
func file_networking_v1beta1_proxy_config_proto_init() {
	if File_networking_v1beta1_proxy_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_networking_v1beta1_proxy_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_networking_v1beta1_proxy_config_proto_goTypes,
		DependencyIndexes: file_networking_v1beta1_proxy_config_proto_depIdxs,
		MessageInfos:      file_networking_v1beta1_proxy_config_proto_msgTypes,
	}.Build()
	File_networking_v1beta1_proxy_config_proto = out.File
	file_networking_v1beta1_proxy_config_proto_rawDesc = nil
	file_networking_v1beta1_proxy_config_proto_goTypes = nil
	file_networking_v1beta1_proxy_config_proto_depIdxs = nil
}
