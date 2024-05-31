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
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: security/v1beta1/request_authentication.proto

// $schema: istio.security.v1beta1.RequestAuthentication
// $title: RequestAuthentication
// $description: Request authentication configuration for workloads.
// $location: https://istio.io/docs/reference/config/security/request_authentication.html
// $aliases: [/docs/reference/config/security/v1beta1/request_authentication, /docs/reference/config/security/v1beta1/jwt, /docs/reference/config/security/v1beta1/jwt.html]

package v1beta1

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// RequestAuthentication defines what request authentication methods are supported by a workload.
// It will reject a request if the request contains invalid authentication information, based on the
// configured authentication rules. A request that does not contain any authentication credentials
// will be accepted but will not have any authenticated identity. To restrict access to authenticated
// requests only, this should be accompanied by an authorization rule.
// Examples:
//
// - Require JWT for all request for workloads that have label `app:httpbin`
//
// ```yaml
// apiVersion: security.istio.io/v1
// kind: RequestAuthentication
// metadata:
//
//	name: httpbin
//	namespace: foo
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: httpbin
//	jwtRules:
//	- issuer: "issuer-foo"
//	  jwksUri: https://example.com/.well-known/jwks.json
//
// ---
// apiVersion: security.istio.io/v1
// kind: AuthorizationPolicy
// metadata:
//
//	name: httpbin
//	namespace: foo
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: httpbin
//	rules:
//	- from:
//	  - source:
//	      requestPrincipals: ["*"]
//
// ```
//
// - A policy in the root namespace ("istio-system" by default) applies to workloads in all namespaces
// in a mesh. The following policy makes all workloads only accept requests that contain a
// valid JWT token.
//
// ```yaml
// apiVersion: security.istio.io/v1
// kind: RequestAuthentication
// metadata:
//
//	name: req-authn-for-all
//	namespace: istio-system
//
// spec:
//
//	jwtRules:
//	- issuer: "issuer-foo"
//	  jwksUri: https://example.com/.well-known/jwks.json
//
// ---
// apiVersion: security.istio.io/v1
// kind: AuthorizationPolicy
// metadata:
//
//	name: require-jwt-for-all
//	namespace: istio-system
//
// spec:
//
//	rules:
//	- from:
//	  - source:
//	      requestPrincipals: ["*"]
//
// ```
//
// - The next example shows how to set a different JWT requirement for a different `host`. The `RequestAuthentication`
// declares it can accept JWTs issued by either `issuer-foo` or `issuer-bar` (the public key set is implicitly
// set from the OpenID Connect spec).
//
// ```yaml
// apiVersion: security.istio.io/v1
// kind: RequestAuthentication
// metadata:
//
//	name: httpbin
//	namespace: foo
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: httpbin
//	jwtRules:
//	- issuer: "issuer-foo"
//	- issuer: "issuer-bar"
//
// ---
// apiVersion: security.istio.io/v1
// kind: AuthorizationPolicy
// metadata:
//
//	name: httpbin
//	namespace: foo
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: httpbin
//	rules:
//	- from:
//	  - source:
//	      requestPrincipals: ["issuer-foo/*"]
//	  to:
//	  - operation:
//	      hosts: ["example.com"]
//	- from:
//	  - source:
//	      requestPrincipals: ["issuer-bar/*"]
//	  to:
//	  - operation:
//	      hosts: ["another-host.com"]
//
// ```
//
// - You can fine tune the authorization policy to set different requirement per path. For example,
// to require JWT on all paths, except /healthz, the same `RequestAuthentication` can be used, but the
// authorization policy could be:
//
// ```yaml
// apiVersion: security.istio.io/v1
// kind: AuthorizationPolicy
// metadata:
//
//	name: httpbin
//	namespace: foo
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: httpbin
//	rules:
//	- from:
//	  - source:
//	      requestPrincipals: ["*"]
//	- to:
//	  - operation:
//	      paths: ["/healthz"]
//
// ```
//
// [Experimental] Routing based on derived [metadata](https://istio.io/latest/docs/reference/config/security/conditions/)
// is now supported. A prefix '@' is used to denote a match against internal metadata instead of the headers in the request.
// Currently this feature is only supported for the following metadata:
//
// - `request.auth.claims.{claim-name}[.{nested-claim}]*` which are extracted from validated JWT tokens.
// Use the `.` or `[]` as a separator for nested claim names.
// Examples: `request.auth.claims.sub`, `request.auth.claims.name.givenName` and `request.auth.claims[foo.com/name]`.
// For more information, see [JWT claim based routing](https://istio.io/latest/docs/tasks/security/authentication/jwt-route/).
//
// The use of matches against JWT claim metadata is only supported in Gateways. The following example shows:
//
// - RequestAuthentication to decode and validate a JWT. This also makes the `@request.auth.claims` available for use in the VirtualService.
// - AuthorizationPolicy to check for valid principals in the request. This makes the JWT required for the request.
// - VirtualService to route the request based on the "sub" claim.
//
// ```yaml
// apiVersion: security.istio.io/v1
// kind: RequestAuthentication
// metadata:
//
//	name: jwt-on-ingress
//	namespace: istio-system
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: istio-ingressgateway
//	jwtRules:
//	- issuer: "example.com"
//	  jwksUri: https://example.com/.well-known/jwks.json
//
// ---
// apiVersion: security.istio.io/v1
// kind: AuthorizationPolicy
// metadata:
//
//	name: require-jwt
//	namespace: istio-system
//
// spec:
//
//	selector:
//	  matchLabels:
//	    app: istio-ingressgateway
//	rules:
//	- from:
//	  - source:
//	      requestPrincipals: ["*"]
//
// ---
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: route-jwt
//
// spec:
//
//	hosts:
//	- foo.prod.svc.cluster.local
//	gateways:
//	- istio-ingressgateway
//	http:
//	- name: "v2"
//	  match:
//	  - headers:
//	      "@request.auth.claims.sub":
//	        exact: "dev"
//	  route:
//	  - destination:
//	      host: foo.prod.svc.cluster.local
//	      subset: v2
//	- name: "default"
//	  route:
//	  - destination:
//	      host: foo.prod.svc.cluster.local
//	      subset: v1
//
// ```
//
// <!-- crd generation tags
// +cue-gen:RequestAuthentication:groupName:security.istio.io
// +cue-gen:RequestAuthentication:versions:v1beta1,v1
// +cue-gen:RequestAuthentication:storageVersion
// +cue-gen:RequestAuthentication:annotations:helm.sh/resource-policy=keep
// +cue-gen:RequestAuthentication:labels:app=istio-pilot,chart=istio,istio=security,heritage=Tiller,release=istio
// +cue-gen:RequestAuthentication:subresource:status
// +cue-gen:RequestAuthentication:scope:Namespaced
// +cue-gen:RequestAuthentication:resource:categories=istio-io,security-istio-io,shortNames=ra
// +cue-gen:RequestAuthentication:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=security.istio.io/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
// +kubebuilder:validation:XValidation:message="only one of targetRefs or workloadSelector can be set",rule="(has(self.selector)?1:0)+(has(self.targetRef)?1:0)+(has(self.targetRefs)?1:0)<=1"
type RequestAuthentication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The selector decides where to apply the request authentication policy. The selector will match with workloads
	// in the same namespace as the request authentication policy. If the request authentication policy is in the root namespace,
	// the selector will additionally match with workloads in all namespaces.
	//
	// If not set, the selector will match all workloads.
	//
	// At most one of `selector` or `targetRefs` can be set for a given policy.
	Selector *v1beta1.WorkloadSelector `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// $hide_from_docs
	TargetRef *v1beta1.PolicyTargetReference `protobuf:"bytes,3,opt,name=targetRef,proto3" json:"targetRef,omitempty"`
	// Optional. The targetRefs specifies a list of resources the policy should be
	// applied to. The targeted resources specified will determine which workloads
	// the policy applies to.
	//
	// Currently, the following resource attachment types are supported:
	// * `kind: Gateway` with `group: gateway.networking.k8s.io` in the same namespace.
	// * `kind: Service` with `""` in the same namespace. This type is only supported for waypoints.
	//
	// If not set, the policy is applied as defined by the selector.
	// At most one of the selector and targetRefs can be set.
	//
	// NOTE: If you are using the `targetRefs` field in a multi-revision environment with Istio versions prior to 1.22,
	// it is highly recommended that you pin the policy to a revision running 1.22+ via the `istio.io/rev` label.
	// This is to prevent proxies connected to older control planes (that don't know about the `targetRefs` field)
	// from misinterpreting the policy as namespace-wide during the upgrade process.
	//
	// NOTE: Waypoint proxies are required to use this field for policies to apply; `selector` policies will be ignored.
	TargetRefs []*v1beta1.PolicyTargetReference `protobuf:"bytes,4,rep,name=targetRefs,proto3" json:"targetRefs,omitempty"`
	// Define the list of JWTs that can be validated at the selected workloads' proxy. A valid token
	// will be used to extract the authenticated identity.
	// Each rule will be activated only when a token is presented at the location recognized by the
	// rule. The token will be validated based on the JWT rule config. If validation fails, the request will
	// be rejected.
	// Note: Requests with multiple tokens (at different locations) are not supported, the output principal of
	// such requests is undefined.
	// +kubebuilder:validation:MaxItems=4096
	JwtRules []*JWTRule `protobuf:"bytes,2,rep,name=jwt_rules,json=jwtRules,proto3" json:"jwt_rules,omitempty"`
}

func (x *RequestAuthentication) Reset() {
	*x = RequestAuthentication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_v1beta1_request_authentication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAuthentication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAuthentication) ProtoMessage() {}

func (x *RequestAuthentication) ProtoReflect() protoreflect.Message {
	mi := &file_security_v1beta1_request_authentication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAuthentication.ProtoReflect.Descriptor instead.
func (*RequestAuthentication) Descriptor() ([]byte, []int) {
	return file_security_v1beta1_request_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *RequestAuthentication) GetSelector() *v1beta1.WorkloadSelector {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *RequestAuthentication) GetTargetRef() *v1beta1.PolicyTargetReference {
	if x != nil {
		return x.TargetRef
	}
	return nil
}

func (x *RequestAuthentication) GetTargetRefs() []*v1beta1.PolicyTargetReference {
	if x != nil {
		return x.TargetRefs
	}
	return nil
}

func (x *RequestAuthentication) GetJwtRules() []*JWTRule {
	if x != nil {
		return x.JwtRules
	}
	return nil
}

// JSON Web Token (JWT) token format for authentication as defined by
// [RFC 7519](https://tools.ietf.org/html/rfc7519). See [OAuth 2.0](https://tools.ietf.org/html/rfc6749) and
// [OIDC 1.0](http://openid.net/connect) for how this is used in the whole
// authentication flow.
//
// Examples:
//
// Spec for a JWT that is issued by `https://example.com`, with the audience claims must be either
// `bookstore_android.apps.example.com` or `bookstore_web.apps.example.com`.
// The token should be presented at the `Authorization` header (default). The JSON Web Key Set (JWKS)
// will be discovered following OpenID Connect protocol.
//
// ```yaml
// issuer: https://example.com
// audiences:
//   - bookstore_android.apps.example.com
//     bookstore_web.apps.example.com
//
// ```
//
// This example specifies a token in a non-default location (`x-goog-iap-jwt-assertion` header). It also
// defines the URI to fetch JWKS explicitly.
//
// ```yaml
// issuer: https://example.com
// jwksUri: https://example.com/.secret/jwks.json
// fromHeaders:
// - "x-goog-iap-jwt-assertion"
// ```
// +kubebuilder:validation:XValidation:message="only one of jwks or jwksUri can be set",rule="(has(self.jwksUri)?1:0)+(has(self.jwks_uri)?1:0)+(has(self.jwks)?1:0)<=1"
type JWTRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifies the issuer that issued the JWT. See
	// [issuer](https://tools.ietf.org/html/rfc7519#section-4.1.1)
	// A JWT with different `iss` claim will be rejected.
	//
	// Example: `https://foobar.auth0.com`
	// Example: `1234567-compute@developer.gserviceaccount.com`
	// +kubebuilder:validation:MinLength=1
	Issuer string `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// The list of JWT
	// [audiences](https://tools.ietf.org/html/rfc7519#section-4.1.3)
	// that are allowed to access. A JWT containing any of these
	// audiences will be accepted.
	//
	// The service name will be accepted if audiences is empty.
	//
	// Example:
	//
	// ```yaml
	// audiences:
	//   - bookstore_android.apps.example.com
	//     bookstore_web.apps.example.com
	//
	// ```
	// +kubebuilder:list-value-validation:MinLength=1
	Audiences []string `protobuf:"bytes,2,rep,name=audiences,proto3" json:"audiences,omitempty"`
	// URL of the provider's public key set to validate signature of the
	// JWT. See [OpenID Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	//
	// Optional if the key set document can either (a) be retrieved from
	// [OpenID
	// Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html) of
	// the issuer or (b) inferred from the email domain of the issuer (e.g. a
	// Google service account).
	//
	// Example: `https://www.googleapis.com/oauth2/v1/certs`
	//
	// Note: Only one of `jwksUri` and `jwks` should be used.
	// +kubebuilder:altName=jwks_uri
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=2048
	// +kubebuilder:validation:XValidation:message="url must have scheme http:// or https://",rule="url(self).getScheme() in ['http', 'https']"
	JwksUri string `protobuf:"bytes,3,opt,name=jwks_uri,json=jwksUri,proto3" json:"jwks_uri,omitempty"`
	// JSON Web Key Set of public keys to validate signature of the JWT.
	// See https://auth0.com/docs/jwks.
	//
	// Note: Only one of `jwksUri` and `jwks` should be used.
	Jwks string `protobuf:"bytes,10,opt,name=jwks,proto3" json:"jwks,omitempty"`
	// List of header locations from which JWT is expected. For example, below is the location spec
	// if JWT is expected to be found in `x-jwt-assertion` header, and have `Bearer` prefix:
	//
	// ```yaml
	//
	//	fromHeaders:
	//	- name: x-jwt-assertion
	//	  prefix: "Bearer "
	//
	// ```
	//
	// Note: Requests with multiple tokens (at different locations) are not supported, the output principal of
	// such requests is undefined.
	FromHeaders []*JWTHeader `protobuf:"bytes,6,rep,name=from_headers,json=fromHeaders,proto3" json:"from_headers,omitempty"`
	// List of query parameters from which JWT is expected. For example, if JWT is provided via query
	// parameter `my_token` (e.g `/path?my_token=<JWT>`), the config is:
	//
	// ```yaml
	//
	//	fromParams:
	//	- "my_token"
	//
	// ```
	//
	// Note: Requests with multiple tokens (at different locations) are not supported, the output principal of
	// such requests is undefined.
	// +kubebuilder:list-value-validation:MinLength=1
	FromParams []string `protobuf:"bytes,7,rep,name=from_params,json=fromParams,proto3" json:"from_params,omitempty"`
	// This field specifies the header name to output a successfully verified JWT payload to the
	// backend. The forwarded data is `base64_encoded(jwt_payload_in_JSON)`. If it is not specified,
	// the payload will not be emitted.
	OutputPayloadToHeader string `protobuf:"bytes,8,opt,name=output_payload_to_header,json=outputPayloadToHeader,proto3" json:"output_payload_to_header,omitempty"`
	// List of cookie names from which JWT is expected.	//
	// For example, if config is:
	//
	// ``` yaml
	//
	//	from_cookies:
	//	- auth-token
	//
	// ```
	// Then JWT will be extracted from “auth-token“ cookie in the request.
	//
	// Note: Requests with multiple tokens (at different locations) are not supported, the output principal of
	// such requests is undefined.
	// +kubebuilder:list-value-validation:MinLength=1
	FromCookies []string `protobuf:"bytes,12,rep,name=from_cookies,json=fromCookies,proto3" json:"from_cookies,omitempty"`
	// If set to true, the original token will be kept for the upstream request. Default is false.
	ForwardOriginalToken bool `protobuf:"varint,9,opt,name=forward_original_token,json=forwardOriginalToken,proto3" json:"forward_original_token,omitempty"`
	// This field specifies a list of operations to copy the claim to HTTP headers on a successfully verified token.
	// This differs from the `output_payload_to_header` by allowing outputting individual claims instead of the whole payload.
	// The header specified in each operation in the list must be unique. Nested claims of type string/int/bool is supported as well.
	// ```
	//
	//	outputClaimToHeaders:
	//	- header: x-my-company-jwt-group
	//	  claim: my-group
	//	- header: x-test-environment-flag
	//	  claim: test-flag
	//	- header: x-jwt-claim-group
	//	  claim: nested.key.group
	//
	// ```
	// [Experimental] This feature is a experimental feature.
	OutputClaimToHeaders []*ClaimToHeader `protobuf:"bytes,11,rep,name=output_claim_to_headers,json=outputClaimToHeaders,proto3" json:"output_claim_to_headers,omitempty"` // [TODO:Update the status whenever this feature is promoted.]
	// The maximum amount of time that the resolver, determined by the PILOT_JWT_ENABLE_REMOTE_JWKS environment variable,
	// will spend waiting for the JWKS to be fetched. Default is 5s.
	Timeout *duration.Duration `protobuf:"bytes,13,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *JWTRule) Reset() {
	*x = JWTRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_v1beta1_request_authentication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTRule) ProtoMessage() {}

func (x *JWTRule) ProtoReflect() protoreflect.Message {
	mi := &file_security_v1beta1_request_authentication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTRule.ProtoReflect.Descriptor instead.
func (*JWTRule) Descriptor() ([]byte, []int) {
	return file_security_v1beta1_request_authentication_proto_rawDescGZIP(), []int{1}
}

func (x *JWTRule) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *JWTRule) GetAudiences() []string {
	if x != nil {
		return x.Audiences
	}
	return nil
}

func (x *JWTRule) GetJwksUri() string {
	if x != nil {
		return x.JwksUri
	}
	return ""
}

func (x *JWTRule) GetJwks() string {
	if x != nil {
		return x.Jwks
	}
	return ""
}

func (x *JWTRule) GetFromHeaders() []*JWTHeader {
	if x != nil {
		return x.FromHeaders
	}
	return nil
}

func (x *JWTRule) GetFromParams() []string {
	if x != nil {
		return x.FromParams
	}
	return nil
}

func (x *JWTRule) GetOutputPayloadToHeader() string {
	if x != nil {
		return x.OutputPayloadToHeader
	}
	return ""
}

func (x *JWTRule) GetFromCookies() []string {
	if x != nil {
		return x.FromCookies
	}
	return nil
}

func (x *JWTRule) GetForwardOriginalToken() bool {
	if x != nil {
		return x.ForwardOriginalToken
	}
	return false
}

func (x *JWTRule) GetOutputClaimToHeaders() []*ClaimToHeader {
	if x != nil {
		return x.OutputClaimToHeaders
	}
	return nil
}

func (x *JWTRule) GetTimeout() *duration.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

// This message specifies a header location to extract JWT token.
type JWTHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The HTTP header name.
	// +kubebuilder:validation:MinLength=1
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The prefix that should be stripped before decoding the token.
	// For example, for `Authorization: Bearer <token>`, prefix=`Bearer` with a space at the end.
	// If the header doesn't have this exact prefix, it is considered invalid.
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *JWTHeader) Reset() {
	*x = JWTHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_v1beta1_request_authentication_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTHeader) ProtoMessage() {}

func (x *JWTHeader) ProtoReflect() protoreflect.Message {
	mi := &file_security_v1beta1_request_authentication_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTHeader.ProtoReflect.Descriptor instead.
func (*JWTHeader) Descriptor() ([]byte, []int) {
	return file_security_v1beta1_request_authentication_proto_rawDescGZIP(), []int{2}
}

func (x *JWTHeader) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JWTHeader) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

// This message specifies the detail for copying claim to header.
type ClaimToHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the header to be created. The header will be overridden if it already exists in the request.
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern=^[-_A-Za-z0-9]+$
	Header string `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The name of the claim to be copied from. Only claim of type string/int/bool is supported.
	// The header will not be there if the claim does not exist or the type of the claim is not supported.
	// +kubebuilder:validation:MinLength=1
	Claim string `protobuf:"bytes,2,opt,name=claim,proto3" json:"claim,omitempty"`
}

func (x *ClaimToHeader) Reset() {
	*x = ClaimToHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_v1beta1_request_authentication_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimToHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimToHeader) ProtoMessage() {}

func (x *ClaimToHeader) ProtoReflect() protoreflect.Message {
	mi := &file_security_v1beta1_request_authentication_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimToHeader.ProtoReflect.Descriptor instead.
func (*ClaimToHeader) Descriptor() ([]byte, []int) {
	return file_security_v1beta1_request_authentication_proto_rawDescGZIP(), []int{3}
}

func (x *ClaimToHeader) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *ClaimToHeader) GetClaim() string {
	if x != nil {
		return x.Claim
	}
	return ""
}

var File_security_v1beta1_request_authentication_proto protoreflect.FileDescriptor

var file_security_v1beta1_request_authentication_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x02, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x40, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x47, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x66, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x66, 0x12, 0x49, 0x0a, 0x0a, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x66, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x52, 0x65, 0x66, 0x73, 0x12, 0x3c, 0x0a, 0x09, 0x6a, 0x77, 0x74, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f,
	0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x4a, 0x57, 0x54, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x08, 0x6a, 0x77, 0x74, 0x52, 0x75,
	0x6c, 0x65, 0x73, 0x22, 0x80, 0x04, 0x0a, 0x07, 0x4a, 0x57, 0x54, 0x52, 0x75, 0x6c, 0x65, 0x12,
	0x1c, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6a,
	0x77, 0x6b, 0x73, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a,
	0x77, 0x6b, 0x73, 0x55, 0x72, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x77, 0x6b, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x77, 0x6b, 0x73, 0x12, 0x44, 0x0a, 0x0c, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4a, 0x57, 0x54, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x37, 0x0a, 0x18, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x15, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x54, 0x6f, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x73, 0x12, 0x34, 0x0a,
	0x16, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x5c, 0x0a, 0x17, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6c,
	0x61, 0x69, 0x6d, 0x5f, 0x74, 0x6f, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6c,
	0x61, 0x69, 0x6d, 0x54, 0x6f, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x14, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x54, 0x6f, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x3d, 0x0a, 0x09, 0x4a, 0x57, 0x54, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x49, 0x0a, 0x0d, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x54, 0x6f,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x42, 0x1f, 0x5a, 0x1d, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_security_v1beta1_request_authentication_proto_rawDescOnce sync.Once
	file_security_v1beta1_request_authentication_proto_rawDescData = file_security_v1beta1_request_authentication_proto_rawDesc
)

func file_security_v1beta1_request_authentication_proto_rawDescGZIP() []byte {
	file_security_v1beta1_request_authentication_proto_rawDescOnce.Do(func() {
		file_security_v1beta1_request_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(file_security_v1beta1_request_authentication_proto_rawDescData)
	})
	return file_security_v1beta1_request_authentication_proto_rawDescData
}

var file_security_v1beta1_request_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_security_v1beta1_request_authentication_proto_goTypes = []interface{}{
	(*RequestAuthentication)(nil),         // 0: istio.security.v1beta1.RequestAuthentication
	(*JWTRule)(nil),                       // 1: istio.security.v1beta1.JWTRule
	(*JWTHeader)(nil),                     // 2: istio.security.v1beta1.JWTHeader
	(*ClaimToHeader)(nil),                 // 3: istio.security.v1beta1.ClaimToHeader
	(*v1beta1.WorkloadSelector)(nil),      // 4: istio.type.v1beta1.WorkloadSelector
	(*v1beta1.PolicyTargetReference)(nil), // 5: istio.type.v1beta1.PolicyTargetReference
	(*duration.Duration)(nil),             // 6: google.protobuf.Duration
}
var file_security_v1beta1_request_authentication_proto_depIdxs = []int32{
	4, // 0: istio.security.v1beta1.RequestAuthentication.selector:type_name -> istio.type.v1beta1.WorkloadSelector
	5, // 1: istio.security.v1beta1.RequestAuthentication.targetRef:type_name -> istio.type.v1beta1.PolicyTargetReference
	5, // 2: istio.security.v1beta1.RequestAuthentication.targetRefs:type_name -> istio.type.v1beta1.PolicyTargetReference
	1, // 3: istio.security.v1beta1.RequestAuthentication.jwt_rules:type_name -> istio.security.v1beta1.JWTRule
	2, // 4: istio.security.v1beta1.JWTRule.from_headers:type_name -> istio.security.v1beta1.JWTHeader
	3, // 5: istio.security.v1beta1.JWTRule.output_claim_to_headers:type_name -> istio.security.v1beta1.ClaimToHeader
	6, // 6: istio.security.v1beta1.JWTRule.timeout:type_name -> google.protobuf.Duration
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_security_v1beta1_request_authentication_proto_init() }
func file_security_v1beta1_request_authentication_proto_init() {
	if File_security_v1beta1_request_authentication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_security_v1beta1_request_authentication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAuthentication); i {
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
		file_security_v1beta1_request_authentication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWTRule); i {
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
		file_security_v1beta1_request_authentication_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWTHeader); i {
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
		file_security_v1beta1_request_authentication_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimToHeader); i {
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
			RawDescriptor: file_security_v1beta1_request_authentication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_security_v1beta1_request_authentication_proto_goTypes,
		DependencyIndexes: file_security_v1beta1_request_authentication_proto_depIdxs,
		MessageInfos:      file_security_v1beta1_request_authentication_proto_msgTypes,
	}.Build()
	File_security_v1beta1_request_authentication_proto = out.File
	file_security_v1beta1_request_authentication_proto_rawDesc = nil
	file_security_v1beta1_request_authentication_proto_goTypes = nil
	file_security_v1beta1_request_authentication_proto_depIdxs = nil
}
