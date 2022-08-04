// Copyright Istio Authors
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
// source: extensions/v1alpha1/wasm.proto

// $schema: istio.extensions.v1alpha1.WasmPlugin
// $title: Wasm Plugin
// $description: Extend the functionality provided by the Istio proxy through WebAssembly filters.
// $location: https://istio.io/docs/reference/config/proxy_extensions/wasm-plugin.html
// $aliases: [/docs/reference/config/extensions/v1alpha1/wasm-plugin]

// WasmPlugins provides a mechanism to extend the functionality provided by
// the Istio proxy through WebAssembly filters.
//
// Order of execution (as part of Envoy's filter chain) is determined by
// phase and priority settings, allowing the configuration of complex
// interactions between user-supplied WasmPlugins and Istio's internal
// filters.
//
// Examples:
//
// AuthN Filter deployed to ingress-gateway that implements an OpenID flow
// and populates the `Authorization` header with a JWT to be consumed by
// Istio AuthN.
//
// ```yaml
// apiVersion: extensions.istio.io/v1alpha1
// kind: WasmPlugin
// metadata:
//   name: openid-connect
//   namespace: istio-ingress
// spec:
//   selector:
//     matchLabels:
//       istio: ingressgateway
//   url: file:///opt/filters/openid.wasm
//   sha256: 1ef0c9a92b0420cf25f7fe5d481b231464bc88f486ca3b9c83ed5cc21d2f6210
//   phase: AUTHN
//   pluginConfig:
//     openid_server: authn
//     openid_realm: ingress
// ```
//
// This is the same as the last example, but using an OCI image.
//
// ```yaml
// apiVersion: extensions.istio.io/v1alpha1
// kind: WasmPlugin
// metadata:
//   name: openid-connect
//   namespace: istio-ingress
// spec:
//   selector:
//     matchLabels:
//       istio: ingressgateway
//   url: oci://private-registry:5000/openid-connect/openid:latest
//   imagePullPolicy: IfNotPresent
//   imagePullSecret: private-registry-pull-secret
//   phase: AUTHN
//   pluginConfig:
//     openid_server: authn
//     openid_realm: ingress
// ```
//
// This is the same as the last example, but using VmConfig to configure environment variables in the VM.
//
// ```yaml
// apiVersion: extensions.istio.io/v1alpha1
// kind: WasmPlugin
// metadata:
//   name: openid-connect
//   namespace: istio-ingress
// spec:
//   selector:
//     matchLabels:
//       istio: ingressgateway
//   url: oci://private-registry:5000/openid-connect/openid:latest
//   imagePullPolicy: IfNotPresent
//   imagePullSecret: private-registry-pull-secret
//   phase: AUTHN
//   pluginConfig:
//     openid_server: authn
//     openid_realm: ingress
//   vmConfig:
//     env:
//     - name: POD_NAME
//       valueFrom: HOST
//     - name: TRUST_DOMAIN
//       value: "cluster.local"
// ```
//
// This is also the same as the last example, but the Wasm module is pulled via https and updated for each time when this plugin resource is changed.
// ```yaml
// apiVersion: extensions.istio.io/v1alpha1
// kind: WasmPlugin
// metadata:
//   name: openid-connect
//   namespace: istio-ingress
// spec:
//   selector:
//     matchLabels:
//       istio: ingressgateway
//   url: https://private-bucket/filters/openid.wasm
//   imagePullPolicy: Always
//   phase: AUTHN
//   pluginConfig:
//     openid_server: authn
//     openid_realm: ingress
//   vmConfig:
//     env:
//     - name: POD_NAME
//       valueFrom: HOST
//     - name: TRUST_DOMAIN
//       value: "cluster.local"
// ```
//
// And a more complex example that deploys three WasmPlugins and orders them
// using `phase` and `priority`. The (hypothetical) setup is that the
// `openid-connect` filter performs an OpenID Connect flow to authenticate the
// user, writing a signed JWT into the Authorization header of the request,
// which can be verified by the Istio authn plugin. Then, the `acl-check` plugin
// kicks in, passing the JWT to a policy server, which in turn responds with a
// signed token that contains information about which files and functions of the
// system are available to the user that was previously authenticated. The
// `acl-check` filter writes this token to a header. Finally, the `check-header`
// filter verifies the token in that header and makes sure that the token's
// contents (the permitted 'function') matches its plugin configuration.
//
// The resulting filter chain looks like this:
// -> openid-connect -> istio.authn -> acl-check -> check-header -> router
//
// ```yaml
// apiVersion: extensions.istio.io/v1alpha1
// kind: WasmPlugin
// metadata:
//   name: openid-connect
//   namespace: istio-ingress
// spec:
//   selector:
//     matchLabels:
//       istio: ingressgateway
//   url: oci://private-registry:5000/openid-connect/openid:latest
//   imagePullPolicy: IfNotPresent
//   imagePullSecret: private-registry-pull-secret
//   phase: AUTHN
//   pluginConfig:
//     openid_server: authn
//     openid_realm: ingress
// ```
//
// ```yaml
// apiVersion: extensions.istio.io/v1alpha1
// kind: WasmPlugin
// metadata:
//   name: acl-check
//   namespace: istio-ingress
// spec:
//   selector:
//     matchLabels:
//       istio: ingressgateway
//   url: oci://private-registry:5000/acl-check/acl:latest
//   imagePullPolicy: Always
//   imagePullSecret: private-registry-pull-secret
//   phase: AUTHZ
//   priority: 1000
//   pluginConfig:
//     acl_server: some_server
//     set_header: authz_complete
// ```
//
// ```yaml
// apiVersion: extensions.istio.io/v1alpha1
// kind: WasmPlugin
// metadata:
//   name: check-header
//   namespace: istio-ingress
// spec:
//   selector:
//     matchLabels:
//       istio: ingressgateway
//   url: oci://private-registry:5000/check-header:latest
//   imagePullPolicy: IfNotPresent
//   imagePullSecret: private-registry-pull-secret
//   phase: AUTHZ
//   priority: 10
//   pluginConfig:
//     read_header: authz_complete
//     verification_key: a89gAzxvls0JKAKIJSBnnvvvkIO
//     function: read_data
// ```
//

package v1alpha1

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// The phase in the filter chain where the plugin will be injected.
type PluginPhase int32

const (
	// Control plane decides where to insert the plugin. This will generally
	// be at the end of the filter chain, right before the Router.
	// Do not specify `PluginPhase` if the plugin is independent of others.
	PluginPhase_UNSPECIFIED_PHASE PluginPhase = 0
	// Insert plugin before Istio authentication filters.
	PluginPhase_AUTHN PluginPhase = 1
	// Insert plugin before Istio authorization filters and after Istio authentication filters.
	PluginPhase_AUTHZ PluginPhase = 2
	// Insert plugin before Istio stats filters and after Istio authorization filters.
	PluginPhase_STATS PluginPhase = 3
)

// Enum value maps for PluginPhase.
var (
	PluginPhase_name = map[int32]string{
		0: "UNSPECIFIED_PHASE",
		1: "AUTHN",
		2: "AUTHZ",
		3: "STATS",
	}
	PluginPhase_value = map[string]int32{
		"UNSPECIFIED_PHASE": 0,
		"AUTHN":             1,
		"AUTHZ":             2,
		"STATS":             3,
	}
)

func (x PluginPhase) Enum() *PluginPhase {
	p := new(PluginPhase)
	*p = x
	return p
}

func (x PluginPhase) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PluginPhase) Descriptor() protoreflect.EnumDescriptor {
	return file_extensions_v1alpha1_wasm_proto_enumTypes[0].Descriptor()
}

func (PluginPhase) Type() protoreflect.EnumType {
	return &file_extensions_v1alpha1_wasm_proto_enumTypes[0]
}

func (x PluginPhase) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PluginPhase.Descriptor instead.
func (PluginPhase) EnumDescriptor() ([]byte, []int) {
	return file_extensions_v1alpha1_wasm_proto_rawDescGZIP(), []int{0}
}

// The pull behaviour to be applied when fetching a Wam module,
// mirroring K8s behaviour.
//
// <!--
// buf:lint:ignore ENUM_VALUE_UPPER_SNAKE_CASE
// -->
type PullPolicy int32

const (
	// Defaults to IfNotPresent, except for OCI images with tag `latest`, for which
	// the default will be Always.
	PullPolicy_UNSPECIFIED_POLICY PullPolicy = 0
	// If an existing version of the image has been pulled before, that
	// will be used. If no version of the image is present locally, we
	// will pull the latest version.
	PullPolicy_IfNotPresent PullPolicy = 1
	// We will always pull the latest version of an image when changing
	// this plugin. Note that the change includes `metadata` field as well.
	PullPolicy_Always PullPolicy = 2
)

// Enum value maps for PullPolicy.
var (
	PullPolicy_name = map[int32]string{
		0: "UNSPECIFIED_POLICY",
		1: "IfNotPresent",
		2: "Always",
	}
	PullPolicy_value = map[string]int32{
		"UNSPECIFIED_POLICY": 0,
		"IfNotPresent":       1,
		"Always":             2,
	}
)

func (x PullPolicy) Enum() *PullPolicy {
	p := new(PullPolicy)
	*p = x
	return p
}

func (x PullPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PullPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_extensions_v1alpha1_wasm_proto_enumTypes[1].Descriptor()
}

func (PullPolicy) Type() protoreflect.EnumType {
	return &file_extensions_v1alpha1_wasm_proto_enumTypes[1]
}

func (x PullPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PullPolicy.Descriptor instead.
func (PullPolicy) EnumDescriptor() ([]byte, []int) {
	return file_extensions_v1alpha1_wasm_proto_rawDescGZIP(), []int{1}
}

type EnvValueSource int32

const (
	// Explicitly given key-value pairs to be injected to this VM
	EnvValueSource_INLINE EnvValueSource = 0
	// *Istio-proxy's* environment variables exposed to this VM.
	EnvValueSource_HOST EnvValueSource = 1
)

// Enum value maps for EnvValueSource.
var (
	EnvValueSource_name = map[int32]string{
		0: "INLINE",
		1: "HOST",
	}
	EnvValueSource_value = map[string]int32{
		"INLINE": 0,
		"HOST":   1,
	}
)

func (x EnvValueSource) Enum() *EnvValueSource {
	p := new(EnvValueSource)
	*p = x
	return p
}

func (x EnvValueSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnvValueSource) Descriptor() protoreflect.EnumDescriptor {
	return file_extensions_v1alpha1_wasm_proto_enumTypes[2].Descriptor()
}

func (EnvValueSource) Type() protoreflect.EnumType {
	return &file_extensions_v1alpha1_wasm_proto_enumTypes[2]
}

func (x EnvValueSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnvValueSource.Descriptor instead.
func (EnvValueSource) EnumDescriptor() ([]byte, []int) {
	return file_extensions_v1alpha1_wasm_proto_rawDescGZIP(), []int{2}
}

// WasmPlugins provides a mechanism to extend the functionality provided by
// the Istio proxy through WebAssembly filters.
//
// <!-- crd generation tags
// +cue-gen:WasmPlugin:groupName:extensions.istio.io
// +cue-gen:WasmPlugin:version:v1alpha1
// +cue-gen:WasmPlugin:storageVersion
// +cue-gen:WasmPlugin:annotations:helm.sh/resource-policy=keep
// +cue-gen:WasmPlugin:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:WasmPlugin:subresource:status
// +cue-gen:WasmPlugin:scope:Namespaced
// +cue-gen:WasmPlugin:resource:categories=istio-io,extensions-istio-io
// +cue-gen:WasmPlugin:preserveUnknownFields:pluginConfig
// +cue-gen:WasmPlugin:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=extensions.istio.io/v1alpha1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type WasmPlugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Criteria used to select the specific set of pods/VMs on which
	// this plugin configuration should be applied. If omitted, this
	// configuration will be applied to all workload instances in the same
	// namespace. If the `WasmPlugin` is present in the config root
	// namespace, it will be applied to all applicable workloads in any
	// namespace.
	Selector *v1beta1.WorkloadSelector `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// URL of a Wasm module or OCI container. If no scheme is present,
	// defaults to `oci://`, referencing an OCI image. Other valid schemes
	// are `file://` for referencing .wasm module files present locally
	// within the proxy container, and `http[s]://` for .wasm module files
	// hosted remotely.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// SHA256 checksum that will be used to verify Wasm module or OCI container.
	// If the `url` field already references a SHA256 (using the `@sha256:`
	// notation), it must match the value of this field. If an OCI image is
	// referenced by tag and this field is set, its checksum will be verified
	// against the contents of this field after pulling.
	Sha256 string `protobuf:"bytes,3,opt,name=sha256,proto3" json:"sha256,omitempty"`
	// The pull behaviour to be applied when fetching Wasm module by either
	// OCI image or http/https. Only relevant when referencing Wasm module without
	// any digest, including the digest in OCI image URL or sha256 field in `vm_config`.
	// Defaults to IfNotPresent, except when an OCI image is referenced in the `url`
	// and the `latest` tag is used, in which case `Always` is the default,
	// mirroring K8s behaviour.
	ImagePullPolicy PullPolicy `protobuf:"varint,4,opt,name=image_pull_policy,json=imagePullPolicy,proto3,enum=istio.extensions.v1alpha1.PullPolicy" json:"image_pull_policy,omitempty"`
	// Credentials to use for OCI image pulling.
	// Name of a K8s Secret in the same namespace as the `WasmPlugin` that
	// contains a docker pull secret which is to be used to authenticate
	// against the registry when pulling the image.
	ImagePullSecret string `protobuf:"bytes,5,opt,name=image_pull_secret,json=imagePullSecret,proto3" json:"image_pull_secret,omitempty"`
	// $hide_from_docs
	// Public key that will be used to verify signatures of signed OCI images
	// or Wasm modules.
	//
	// At this moment, various ways for signing/verifying are emerging and being proposed.
	// We can observe two major streams for signing OCI images: Cosign from Sigstore and Notary,
	// which is used in Docker Content Trust.
	// In case of Wasm module, multiple approaches are still in discussion.
	//   - https://github.com/WebAssembly/design/issues/1413
	//   - https://github.com/wasm-signatures/design (various signing tools are enumerated)
	//
	// In addition, for each method for signing&verifying, we may need to consider to provide
	// additional data or configuration (e.g., key rolling, KMS, root certs, ...) as well.
	//
	// To deal with this situation, we need to elaborate more generic way to describe
	// how to sign and verify the image or wasm binary, and how to specify relevant data,
	// including this `verification_key`.
	//
	// Therefore, this field will not be implemented until the detailed design is established.
	// For the future use, just keep this field in proto and hide from documentation.
	VerificationKey string `protobuf:"bytes,6,opt,name=verification_key,json=verificationKey,proto3" json:"verification_key,omitempty"`
	// The configuration that will be passed on to the plugin.
	PluginConfig *_struct.Struct `protobuf:"bytes,7,opt,name=plugin_config,json=pluginConfig,proto3" json:"plugin_config,omitempty"`
	// The plugin name to be used in the Envoy configuration (used to be called
	// `rootID`). Some .wasm modules might require this value to select the Wasm
	// plugin to execute.
	PluginName string `protobuf:"bytes,8,opt,name=plugin_name,json=pluginName,proto3" json:"plugin_name,omitempty"`
	// Determines where in the filter chain this `WasmPlugin` is to be injected.
	Phase PluginPhase `protobuf:"varint,9,opt,name=phase,proto3,enum=istio.extensions.v1alpha1.PluginPhase" json:"phase,omitempty"`
	// Determines ordering of `WasmPlugins` in the same `phase`.
	// When multiple `WasmPlugins` are applied to the same workload in the
	// same `phase`, they will be applied by priority, in descending order.
	// If `priority` is not set, or two `WasmPlugins` exist with the same
	// value, the ordering will be deterministically derived from name and
	// namespace of the `WasmPlugins`. Defaults to `0`.
	Priority *wrappers.Int64Value `protobuf:"bytes,10,opt,name=priority,proto3" json:"priority,omitempty"`
	// Configuration for a Wasm VM.
	// more details can be found [here](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/wasm/v3/wasm.proto#extensions-wasm-v3-vmconfig).
	VmConfig *VmConfig `protobuf:"bytes,11,opt,name=vm_config,json=vmConfig,proto3" json:"vm_config,omitempty"`
}

func (x *WasmPlugin) Reset() {
	*x = WasmPlugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extensions_v1alpha1_wasm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WasmPlugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WasmPlugin) ProtoMessage() {}

func (x *WasmPlugin) ProtoReflect() protoreflect.Message {
	mi := &file_extensions_v1alpha1_wasm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WasmPlugin.ProtoReflect.Descriptor instead.
func (*WasmPlugin) Descriptor() ([]byte, []int) {
	return file_extensions_v1alpha1_wasm_proto_rawDescGZIP(), []int{0}
}

func (x *WasmPlugin) GetSelector() *v1beta1.WorkloadSelector {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *WasmPlugin) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *WasmPlugin) GetSha256() string {
	if x != nil {
		return x.Sha256
	}
	return ""
}

func (x *WasmPlugin) GetImagePullPolicy() PullPolicy {
	if x != nil {
		return x.ImagePullPolicy
	}
	return PullPolicy_UNSPECIFIED_POLICY
}

func (x *WasmPlugin) GetImagePullSecret() string {
	if x != nil {
		return x.ImagePullSecret
	}
	return ""
}

func (x *WasmPlugin) GetVerificationKey() string {
	if x != nil {
		return x.VerificationKey
	}
	return ""
}

func (x *WasmPlugin) GetPluginConfig() *_struct.Struct {
	if x != nil {
		return x.PluginConfig
	}
	return nil
}

func (x *WasmPlugin) GetPluginName() string {
	if x != nil {
		return x.PluginName
	}
	return ""
}

func (x *WasmPlugin) GetPhase() PluginPhase {
	if x != nil {
		return x.Phase
	}
	return PluginPhase_UNSPECIFIED_PHASE
}

func (x *WasmPlugin) GetPriority() *wrappers.Int64Value {
	if x != nil {
		return x.Priority
	}
	return nil
}

func (x *WasmPlugin) GetVmConfig() *VmConfig {
	if x != nil {
		return x.VmConfig
	}
	return nil
}

// Configuration for a Wasm VM.
// more details can be found [here](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/wasm/v3/wasm.proto#extensions-wasm-v3-vmconfig).
type VmConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies environment variables to be injected to this VM.
	// Note that if a key does not exist, it will be ignored.
	Env []*EnvVar `protobuf:"bytes,1,rep,name=env,proto3" json:"env,omitempty"`
}

func (x *VmConfig) Reset() {
	*x = VmConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extensions_v1alpha1_wasm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VmConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VmConfig) ProtoMessage() {}

func (x *VmConfig) ProtoReflect() protoreflect.Message {
	mi := &file_extensions_v1alpha1_wasm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VmConfig.ProtoReflect.Descriptor instead.
func (*VmConfig) Descriptor() ([]byte, []int) {
	return file_extensions_v1alpha1_wasm_proto_rawDescGZIP(), []int{1}
}

func (x *VmConfig) GetEnv() []*EnvVar {
	if x != nil {
		return x.Env
	}
	return nil
}

type EnvVar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required
	// Name of the environment variable. Must be a C_IDENTIFIER.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required
	// Source for the environment variable's value.
	ValueFrom EnvValueSource `protobuf:"varint,3,opt,name=value_from,json=valueFrom,proto3,enum=istio.extensions.v1alpha1.EnvValueSource" json:"value_from,omitempty"`
	// Value for the environment variable.
	// Note that if `value_from` is `HOST`, it will be ignored.
	// Defaults to "".
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EnvVar) Reset() {
	*x = EnvVar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extensions_v1alpha1_wasm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvVar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvVar) ProtoMessage() {}

func (x *EnvVar) ProtoReflect() protoreflect.Message {
	mi := &file_extensions_v1alpha1_wasm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvVar.ProtoReflect.Descriptor instead.
func (*EnvVar) Descriptor() ([]byte, []int) {
	return file_extensions_v1alpha1_wasm_proto_rawDescGZIP(), []int{2}
}

func (x *EnvVar) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EnvVar) GetValueFrom() EnvValueSource {
	if x != nil {
		return x.ValueFrom
	}
	return EnvValueSource_INLINE
}

func (x *EnvVar) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_extensions_v1alpha1_wasm_proto protoreflect.FileDescriptor

var file_extensions_v1alpha1_wasm_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x04, 0x0a, 0x0a, 0x57, 0x61, 0x73, 0x6d, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x40, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61,
	0x32, 0x35, 0x36, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35,
	0x36, 0x12, 0x51, 0x0a, 0x11, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x75, 0x6c, 0x6c, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x69,
	0x73, 0x74, 0x69, 0x6f, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x0f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x75, 0x6c, 0x6c, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x75,
	0x6c, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x75, 0x6c, 0x6c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x12, 0x29, 0x0a, 0x10, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x3c, 0x0a, 0x0d, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0c, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x70, 0x68,
	0x61, 0x73, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x69, 0x73, 0x74, 0x69,
	0x6f, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x68, 0x61, 0x73,
	0x65, 0x52, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x40, 0x0a, 0x09, 0x76, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x56, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x76, 0x6d, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x3f, 0x0a, 0x08, 0x56, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x33, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69,
	0x73, 0x74, 0x69, 0x6f, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x52,
	0x03, 0x65, 0x6e, 0x76, 0x22, 0x7c, 0x0a, 0x06, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x48, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x2a, 0x45, 0x0a, 0x0b, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x68, 0x61, 0x73,
	0x65, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x5f, 0x50, 0x48, 0x41, 0x53, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x55, 0x54, 0x48,
	0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x55, 0x54, 0x48, 0x5a, 0x10, 0x02, 0x12, 0x09,
	0x0a, 0x05, 0x53, 0x54, 0x41, 0x54, 0x53, 0x10, 0x03, 0x2a, 0x42, 0x0a, 0x0a, 0x50, 0x75, 0x6c,
	0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x10, 0x00, 0x12,
	0x10, 0x0a, 0x0c, 0x49, 0x66, 0x4e, 0x6f, 0x74, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x6c, 0x77, 0x61, 0x79, 0x73, 0x10, 0x02, 0x2a, 0x26, 0x0a,
	0x0e, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x0a, 0x0a, 0x06, 0x49, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48,
	0x4f, 0x53, 0x54, 0x10, 0x01, 0x42, 0x22, 0x5a, 0x20, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_extensions_v1alpha1_wasm_proto_rawDescOnce sync.Once
	file_extensions_v1alpha1_wasm_proto_rawDescData = file_extensions_v1alpha1_wasm_proto_rawDesc
)

func file_extensions_v1alpha1_wasm_proto_rawDescGZIP() []byte {
	file_extensions_v1alpha1_wasm_proto_rawDescOnce.Do(func() {
		file_extensions_v1alpha1_wasm_proto_rawDescData = protoimpl.X.CompressGZIP(file_extensions_v1alpha1_wasm_proto_rawDescData)
	})
	return file_extensions_v1alpha1_wasm_proto_rawDescData
}

var file_extensions_v1alpha1_wasm_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_extensions_v1alpha1_wasm_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_extensions_v1alpha1_wasm_proto_goTypes = []interface{}{
	(PluginPhase)(0),                 // 0: istio.extensions.v1alpha1.PluginPhase
	(PullPolicy)(0),                  // 1: istio.extensions.v1alpha1.PullPolicy
	(EnvValueSource)(0),              // 2: istio.extensions.v1alpha1.EnvValueSource
	(*WasmPlugin)(nil),               // 3: istio.extensions.v1alpha1.WasmPlugin
	(*VmConfig)(nil),                 // 4: istio.extensions.v1alpha1.VmConfig
	(*EnvVar)(nil),                   // 5: istio.extensions.v1alpha1.EnvVar
	(*v1beta1.WorkloadSelector)(nil), // 6: istio.type.v1beta1.WorkloadSelector
	(*_struct.Struct)(nil),           // 7: google.protobuf.Struct
	(*wrappers.Int64Value)(nil),      // 8: google.protobuf.Int64Value
}
var file_extensions_v1alpha1_wasm_proto_depIdxs = []int32{
	6, // 0: istio.extensions.v1alpha1.WasmPlugin.selector:type_name -> istio.type.v1beta1.WorkloadSelector
	1, // 1: istio.extensions.v1alpha1.WasmPlugin.image_pull_policy:type_name -> istio.extensions.v1alpha1.PullPolicy
	7, // 2: istio.extensions.v1alpha1.WasmPlugin.plugin_config:type_name -> google.protobuf.Struct
	0, // 3: istio.extensions.v1alpha1.WasmPlugin.phase:type_name -> istio.extensions.v1alpha1.PluginPhase
	8, // 4: istio.extensions.v1alpha1.WasmPlugin.priority:type_name -> google.protobuf.Int64Value
	4, // 5: istio.extensions.v1alpha1.WasmPlugin.vm_config:type_name -> istio.extensions.v1alpha1.VmConfig
	5, // 6: istio.extensions.v1alpha1.VmConfig.env:type_name -> istio.extensions.v1alpha1.EnvVar
	2, // 7: istio.extensions.v1alpha1.EnvVar.value_from:type_name -> istio.extensions.v1alpha1.EnvValueSource
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_extensions_v1alpha1_wasm_proto_init() }
func file_extensions_v1alpha1_wasm_proto_init() {
	if File_extensions_v1alpha1_wasm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_extensions_v1alpha1_wasm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WasmPlugin); i {
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
		file_extensions_v1alpha1_wasm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VmConfig); i {
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
		file_extensions_v1alpha1_wasm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvVar); i {
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
			RawDescriptor: file_extensions_v1alpha1_wasm_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_extensions_v1alpha1_wasm_proto_goTypes,
		DependencyIndexes: file_extensions_v1alpha1_wasm_proto_depIdxs,
		EnumInfos:         file_extensions_v1alpha1_wasm_proto_enumTypes,
		MessageInfos:      file_extensions_v1alpha1_wasm_proto_msgTypes,
	}.Build()
	File_extensions_v1alpha1_wasm_proto = out.File
	file_extensions_v1alpha1_wasm_proto_rawDesc = nil
	file_extensions_v1alpha1_wasm_proto_goTypes = nil
	file_extensions_v1alpha1_wasm_proto_depIdxs = nil
}
