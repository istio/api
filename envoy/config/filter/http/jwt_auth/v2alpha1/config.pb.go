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
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: envoy/config/filter/http/jwt_auth/v2alpha1/config.proto

package v2alpha1

import (
	duration "github.com/golang/protobuf/ptypes/duration"
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

// Copied from @envoy/api/envoy/api/v2/core/http_uri.proto
// Envoy external URI descriptor
type HttpUri struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The HTTP server URI. It should be a full FQDN with protocol, host and path.
	//
	// Example:
	//
	// .. code-block:: yaml
	//
	//	uri: https://www.googleapis.com/oauth2/v1/certs
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Specify how `uri` is to be fetched. Today, this requires an explicit
	// cluster, but in the future we may support dynamic cluster creation or
	// inline DNS resolution. See `issue
	// <https://github.com/envoyproxy/envoy/issues/1606>`_.
	//
	// Types that are assignable to HttpUpstreamType:
	//
	//	*HttpUri_Cluster
	HttpUpstreamType isHttpUri_HttpUpstreamType `protobuf_oneof:"http_upstream_type"`
	// Sets the maximum duration in milliseconds that a response can take to arrive upon request.
	Timeout *duration.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *HttpUri) Reset() {
	*x = HttpUri{}
	mi := &file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HttpUri) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpUri) ProtoMessage() {}

func (x *HttpUri) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpUri.ProtoReflect.Descriptor instead.
func (*HttpUri) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_rawDescGZIP(), []int{0}
}

func (x *HttpUri) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (m *HttpUri) GetHttpUpstreamType() isHttpUri_HttpUpstreamType {
	if m != nil {
		return m.HttpUpstreamType
	}
	return nil
}

func (x *HttpUri) GetCluster() string {
	if x, ok := x.GetHttpUpstreamType().(*HttpUri_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (x *HttpUri) GetTimeout() *duration.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type isHttpUri_HttpUpstreamType interface {
	isHttpUri_HttpUpstreamType()
}

type HttpUri_Cluster struct {
	// A cluster is created in the Envoy "cluster_manager" config
	// section. This field specifies the cluster name.
	//
	// Example:
	//
	// .. code-block:: yaml
	//
	//	cluster: jwks_cluster
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

func (*HttpUri_Cluster) isHttpUri_HttpUpstreamType() {}

// Copied from @envoy/api/envoy/api/v2/core/base.proto
// Data source consisting of either a file or an inline value.
type DataSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Specifier:
	//
	//	*DataSource_Filename
	//	*DataSource_InlineBytes
	//	*DataSource_InlineString
	Specifier isDataSource_Specifier `protobuf_oneof:"specifier"`
}

func (x *DataSource) Reset() {
	*x = DataSource{}
	mi := &file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSource) ProtoMessage() {}

func (x *DataSource) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSource.ProtoReflect.Descriptor instead.
func (*DataSource) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_rawDescGZIP(), []int{1}
}

func (m *DataSource) GetSpecifier() isDataSource_Specifier {
	if m != nil {
		return m.Specifier
	}
	return nil
}

func (x *DataSource) GetFilename() string {
	if x, ok := x.GetSpecifier().(*DataSource_Filename); ok {
		return x.Filename
	}
	return ""
}

func (x *DataSource) GetInlineBytes() []byte {
	if x, ok := x.GetSpecifier().(*DataSource_InlineBytes); ok {
		return x.InlineBytes
	}
	return nil
}

func (x *DataSource) GetInlineString() string {
	if x, ok := x.GetSpecifier().(*DataSource_InlineString); ok {
		return x.InlineString
	}
	return ""
}

type isDataSource_Specifier interface {
	isDataSource_Specifier()
}

type DataSource_Filename struct {
	// Local filesystem data source.
	Filename string `protobuf:"bytes,1,opt,name=filename,proto3,oneof"`
}

type DataSource_InlineBytes struct {
	// Bytes inlined in the configuration.
	InlineBytes []byte `protobuf:"bytes,2,opt,name=inline_bytes,json=inlineBytes,proto3,oneof"`
}

type DataSource_InlineString struct {
	// String inlined in the configuration.
	InlineString string `protobuf:"bytes,3,opt,name=inline_string,json=inlineString,proto3,oneof"`
}

func (*DataSource_Filename) isDataSource_Specifier() {}

func (*DataSource_InlineBytes) isDataSource_Specifier() {}

func (*DataSource_InlineString) isDataSource_Specifier() {}

// This message specifies how a JSON Web Token (JWT) can be verified. See the [JWT format definition](https://tools.ietf.org/html/rfc7519)
// for details. Please see [OAuth2.0](https://tools.ietf.org/html/rfc6749) and
// [OIDC1.0](http://openid.net/connect) for
// the authentication flow.
//
// Example:
//
// ```yaml
//
//	issuer: https://example.com
//	audiences:
//	- bookstore_android.apps.googleusercontent.com
//	  bookstore_web.apps.googleusercontent.com
//	remote_jwks:
//	- http_uri:
//	  - uri: https://example.com/.well-known/jwks.json
//	    cluster: example_jwks_cluster
//	  cache_duration:
//	  - seconds: 300
//
// ```
type JwtRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifies the principal that issued the JWT. See `here
	//
	//	<https://tools.ietf.org/html/rfc7519#section-4.1.1>`_. Usually a URL or an email address.
	//
	// Example: https://securetoken.google.com
	// Example: 1234567-compute@developer.gserviceaccount.com
	Issuer string `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// The list of JWT `audiences <https://tools.ietf.org/html/rfc7519#section-4.1.3>`_. that are
	// allowed to access. A JWT containing any of these audiences will be accepted. If not specified,
	// will not check audiences in the token.
	//
	// Example:
	//
	// .. code-block:: yaml
	//
	//	audiences:
	//	- bookstore_android.apps.googleusercontent.com
	//	  bookstore_web.apps.googleusercontent.com
	Audiences []string `protobuf:"bytes,2,rep,name=audiences,proto3" json:"audiences,omitempty"`
	// `JSON Web Key Set <https://tools.ietf.org/html/rfc7517#appendix-A>`_ is needed. to validate
	// signature of the JWT. This field specifies where to fetch JWKS.
	//
	// Types that are assignable to JwksSourceSpecifier:
	//
	//	*JwtRule_RemoteJwks
	//	*JwtRule_LocalJwks
	JwksSourceSpecifier isJwtRule_JwksSourceSpecifier `protobuf_oneof:"jwks_source_specifier"`
	// If false, the JWT is removed in the request after a success verification. If true, the JWT is
	// not removed in the request. Default value is false.
	Forward bool `protobuf:"varint,5,opt,name=forward,proto3" json:"forward,omitempty"`
	// Specify the HTTP headers to extract JWT token. For examples, following config:
	//
	// .. code-block:: yaml
	//
	//	from_headers:
	//	- name: x-goog-iap-jwt-assertion
	//
	// can be used to extract token from header::
	//
	//	x-goog-iap-jwt-assertion: <JWT>.
	FromHeaders []*JwtHeader `protobuf:"bytes,6,rep,name=from_headers,json=fromHeaders,proto3" json:"from_headers,omitempty"`
	// JWT is sent in a query parameter. `jwt_params` represents the query parameter names.
	//
	// For example, if config is:
	//
	// .. code-block:: yaml
	//
	//	from_params:
	//	- jwt_token
	//
	// The JWT format in query parameter is::
	//
	//	/path?jwt_token=<JWT>
	FromParams []string `protobuf:"bytes,7,rep,name=from_params,json=fromParams,proto3" json:"from_params,omitempty"`
	// This field specifies the header name to forward a successfully verified JWT payload to the
	// backend. The forwarded data is::
	//
	//	base64_encoded(jwt_payload_in_JSON)
	//
	// If it is not specified, the payload will not be forwarded.
	// Multiple JWTs in a request from different issuers will be supported. Multiple JWTs from the
	// same issuer will not be supported. Each issuer can config this `forward_payload_header`. If
	// multiple JWTs from different issuers want to forward their payloads, their
	// `forward_payload_header` should be different.
	ForwardPayloadHeader string `protobuf:"bytes,8,opt,name=forward_payload_header,json=forwardPayloadHeader,proto3" json:"forward_payload_header,omitempty"`
}

func (x *JwtRule) Reset() {
	*x = JwtRule{}
	mi := &file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JwtRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JwtRule) ProtoMessage() {}

func (x *JwtRule) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JwtRule.ProtoReflect.Descriptor instead.
func (*JwtRule) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_rawDescGZIP(), []int{2}
}

func (x *JwtRule) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *JwtRule) GetAudiences() []string {
	if x != nil {
		return x.Audiences
	}
	return nil
}

func (m *JwtRule) GetJwksSourceSpecifier() isJwtRule_JwksSourceSpecifier {
	if m != nil {
		return m.JwksSourceSpecifier
	}
	return nil
}

func (x *JwtRule) GetRemoteJwks() *RemoteJwks {
	if x, ok := x.GetJwksSourceSpecifier().(*JwtRule_RemoteJwks); ok {
		return x.RemoteJwks
	}
	return nil
}

func (x *JwtRule) GetLocalJwks() *DataSource {
	if x, ok := x.GetJwksSourceSpecifier().(*JwtRule_LocalJwks); ok {
		return x.LocalJwks
	}
	return nil
}

func (x *JwtRule) GetForward() bool {
	if x != nil {
		return x.Forward
	}
	return false
}

func (x *JwtRule) GetFromHeaders() []*JwtHeader {
	if x != nil {
		return x.FromHeaders
	}
	return nil
}

func (x *JwtRule) GetFromParams() []string {
	if x != nil {
		return x.FromParams
	}
	return nil
}

func (x *JwtRule) GetForwardPayloadHeader() string {
	if x != nil {
		return x.ForwardPayloadHeader
	}
	return ""
}

type isJwtRule_JwksSourceSpecifier interface {
	isJwtRule_JwksSourceSpecifier()
}

type JwtRule_RemoteJwks struct {
	// JWKS can be fetched from remote server via HTTP/HTTPS. This field specifies the remote HTTP
	// URI and how the fetched JWKS should be cached.
	//
	// Example:
	//
	// .. code-block:: yaml
	//
	//	remote_jwks:
	//	- http_uri:
	//	  - uri: https://www.googleapis.com/oauth2/v1/certs
	//	    cluster: jwt.www.googleapis.com|443
	//	  cache_duration:
	//	  - seconds: 300
	RemoteJwks *RemoteJwks `protobuf:"bytes,3,opt,name=remote_jwks,json=remoteJwks,proto3,oneof"`
}

type JwtRule_LocalJwks struct {
	// JWKS is in local data source. It could be either in a local file or embedded in the
	// inline_string.
	//
	// Example: local file
	//
	// .. code-block:: yaml
	//
	//	local_jwks:
	//	- filename: /etc/envoy/jwks/jwks1.txt
	//
	// Example: inline_string
	//
	// .. code-block:: yaml
	//
	//	local_jwks:
	//	- inline_string: "ACADADADADA"
	LocalJwks *DataSource `protobuf:"bytes,4,opt,name=local_jwks,json=localJwks,proto3,oneof"`
}

func (*JwtRule_RemoteJwks) isJwtRule_JwksSourceSpecifier() {}

func (*JwtRule_LocalJwks) isJwtRule_JwksSourceSpecifier() {}

// This message specifies how to fetch JWKS from remote and how to cache it.
type RemoteJwks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The HTTP URI to fetch the JWKS. For example:
	//
	// .. code-block:: yaml
	//
	//	http_uri:
	//	- uri: https://www.googleapis.com/oauth2/v1/certs
	//	  cluster: jwt.www.googleapis.com|443
	HttpUri *HttpUri `protobuf:"bytes,1,opt,name=http_uri,json=httpUri,proto3" json:"http_uri,omitempty"`
	// Duration after which the cached JWKS should be expired. If not specified, default cache
	// duration is 5 minutes.
	CacheDuration *duration.Duration `protobuf:"bytes,2,opt,name=cache_duration,json=cacheDuration,proto3" json:"cache_duration,omitempty"`
}

func (x *RemoteJwks) Reset() {
	*x = RemoteJwks{}
	mi := &file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoteJwks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteJwks) ProtoMessage() {}

func (x *RemoteJwks) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteJwks.ProtoReflect.Descriptor instead.
func (*RemoteJwks) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_rawDescGZIP(), []int{3}
}

func (x *RemoteJwks) GetHttpUri() *HttpUri {
	if x != nil {
		return x.HttpUri
	}
	return nil
}

func (x *RemoteJwks) GetCacheDuration() *duration.Duration {
	if x != nil {
		return x.CacheDuration
	}
	return nil
}

// This message specifies a header location to extract JWT token.
type JwtHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The HTTP header name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The value prefix. The value format is "value_prefix<token>"
	// For example, for "Authorization: Bearer <token>", value_prefix="Bearer " with a space at the
	// end.
	ValuePrefix string `protobuf:"bytes,2,opt,name=value_prefix,json=valuePrefix,proto3" json:"value_prefix,omitempty"`
}

func (x *JwtHeader) Reset() {
	*x = JwtHeader{}
	mi := &file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JwtHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JwtHeader) ProtoMessage() {}

func (x *JwtHeader) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JwtHeader.ProtoReflect.Descriptor instead.
func (*JwtHeader) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_rawDescGZIP(), []int{4}
}

func (x *JwtHeader) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JwtHeader) GetValuePrefix() string {
	if x != nil {
		return x.ValuePrefix
	}
	return ""
}

// This is the Envoy HTTP filter config for JWT authentication.
// [#not-implemented-hide:]
type JwtAuthentication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of JWT rules to valide.
	Rules []*JwtRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	// If true, the request is allowed if JWT is missing or JWT verification fails.
	// Default is false, a request without JWT or failed JWT verification is not allowed.
	AllowMissingOrFailed bool `protobuf:"varint,2,opt,name=allow_missing_or_failed,json=allowMissingOrFailed,proto3" json:"allow_missing_or_failed,omitempty"`
}

func (x *JwtAuthentication) Reset() {
	*x = JwtAuthentication{}
	mi := &file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JwtAuthentication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JwtAuthentication) ProtoMessage() {}

func (x *JwtAuthentication) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JwtAuthentication.ProtoReflect.Descriptor instead.
func (*JwtAuthentication) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_rawDescGZIP(), []int{5}
}

func (x *JwtAuthentication) GetRules() []*JwtRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *JwtAuthentication) GetAllowMissingOrFailed() bool {
	if x != nil {
		return x.AllowMissingOrFailed
	}
	return false
}

var File_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto protoreflect.FileDescriptor

var file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_rawDesc = []byte{
	0x0a, 0x37, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x6a, 0x77, 0x74, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x30, 0x69, 0x73, 0x74, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6a, 0x77, 0x74, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x07,
	0x48, 0x74, 0x74, 0x70, 0x55, 0x72, 0x69, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x1a, 0x0a, 0x07, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0x14, 0x0a, 0x12, 0x68, 0x74,
	0x74, 0x70, 0x5f, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x83, 0x01, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x1c, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x0c, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0b, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x12, 0x25, 0x0a, 0x0d, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x0b, 0x0a, 0x09, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0xe9, 0x03, 0x0a, 0x07, 0x4a, 0x77, 0x74, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75,
	0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x5f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x5f, 0x6a, 0x77, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e,
	0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6a,
	0x77, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4a, 0x77, 0x6b, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4a, 0x77, 0x6b, 0x73, 0x12, 0x5d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x5f, 0x6a, 0x77, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e,
	0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6a,
	0x77, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x09, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x4a, 0x77, 0x6b, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x12, 0x5e, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6a, 0x77, 0x74, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4a, 0x77, 0x74, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x17, 0x0a, 0x15, 0x6a, 0x77, 0x6b,
	0x73, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x22, 0xa4, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4a, 0x77, 0x6b,
	0x73, 0x12, 0x54, 0x0a, 0x08, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x6a, 0x77, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x55, 0x72, 0x69, 0x52, 0x07,
	0x68, 0x74, 0x74, 0x70, 0x55, 0x72, 0x69, 0x12, 0x40, 0x0a, 0x0e, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x42, 0x0a, 0x09, 0x4a, 0x77, 0x74,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x9b, 0x01,
	0x0a, 0x11, 0x4a, 0x77, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x39, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x6a, 0x77, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4a, 0x77, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72,
	0x75, 0x6c, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x17, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x72, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4d, 0x69, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x4f, 0x72, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x42, 0x39, 0x5a, 0x37, 0x69,
	0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x2f, 0x6a, 0x77, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_rawDescOnce sync.Once
	file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_rawDescData = file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_rawDesc
)

func file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_rawDescGZIP() []byte {
	file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_rawDescOnce.Do(func() {
		file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_rawDescData)
	})
	return file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_rawDescData
}

var file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_goTypes = []any{
	(*HttpUri)(nil),           // 0: istio.envoy.config.filter.http.jwt_auth.v2alpha1.HttpUri
	(*DataSource)(nil),        // 1: istio.envoy.config.filter.http.jwt_auth.v2alpha1.DataSource
	(*JwtRule)(nil),           // 2: istio.envoy.config.filter.http.jwt_auth.v2alpha1.JwtRule
	(*RemoteJwks)(nil),        // 3: istio.envoy.config.filter.http.jwt_auth.v2alpha1.RemoteJwks
	(*JwtHeader)(nil),         // 4: istio.envoy.config.filter.http.jwt_auth.v2alpha1.JwtHeader
	(*JwtAuthentication)(nil), // 5: istio.envoy.config.filter.http.jwt_auth.v2alpha1.JwtAuthentication
	(*duration.Duration)(nil), // 6: google.protobuf.Duration
}
var file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_depIdxs = []int32{
	6, // 0: istio.envoy.config.filter.http.jwt_auth.v2alpha1.HttpUri.timeout:type_name -> google.protobuf.Duration
	3, // 1: istio.envoy.config.filter.http.jwt_auth.v2alpha1.JwtRule.remote_jwks:type_name -> istio.envoy.config.filter.http.jwt_auth.v2alpha1.RemoteJwks
	1, // 2: istio.envoy.config.filter.http.jwt_auth.v2alpha1.JwtRule.local_jwks:type_name -> istio.envoy.config.filter.http.jwt_auth.v2alpha1.DataSource
	4, // 3: istio.envoy.config.filter.http.jwt_auth.v2alpha1.JwtRule.from_headers:type_name -> istio.envoy.config.filter.http.jwt_auth.v2alpha1.JwtHeader
	0, // 4: istio.envoy.config.filter.http.jwt_auth.v2alpha1.RemoteJwks.http_uri:type_name -> istio.envoy.config.filter.http.jwt_auth.v2alpha1.HttpUri
	6, // 5: istio.envoy.config.filter.http.jwt_auth.v2alpha1.RemoteJwks.cache_duration:type_name -> google.protobuf.Duration
	2, // 6: istio.envoy.config.filter.http.jwt_auth.v2alpha1.JwtAuthentication.rules:type_name -> istio.envoy.config.filter.http.jwt_auth.v2alpha1.JwtRule
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_init() }
func file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_init() {
	if File_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto != nil {
		return
	}
	file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_msgTypes[0].OneofWrappers = []any{
		(*HttpUri_Cluster)(nil),
	}
	file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_msgTypes[1].OneofWrappers = []any{
		(*DataSource_Filename)(nil),
		(*DataSource_InlineBytes)(nil),
		(*DataSource_InlineString)(nil),
	}
	file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_msgTypes[2].OneofWrappers = []any{
		(*JwtRule_RemoteJwks)(nil),
		(*JwtRule_LocalJwks)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_goTypes,
		DependencyIndexes: file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_depIdxs,
		MessageInfos:      file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_msgTypes,
	}.Build()
	File_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto = out.File
	file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_rawDesc = nil
	file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_goTypes = nil
	file_envoy_config_filter_http_jwt_auth_v2alpha1_config_proto_depIdxs = nil
}
