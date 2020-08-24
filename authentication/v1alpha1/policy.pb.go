// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authentication/v1alpha1/policy.proto

// This package defines user-facing authentication policy.

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google/api/field_behavior.proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// $hide_from_docs
// Deprecated. When using security/v1beta1/RequestAuthentication, the request principal always
// comes from request authentication (i.e JWT).
// Associates authentication with request principal.
type PrincipalBinding int32

const (
	// Principal will be set to the identity from peer authentication.
	PrincipalBinding_USE_PEER PrincipalBinding = 0
	// Principal will be set to the identity from origin authentication.
	PrincipalBinding_USE_ORIGIN PrincipalBinding = 1
)

var PrincipalBinding_name = map[int32]string{
	0: "USE_PEER",
	1: "USE_ORIGIN",
}

var PrincipalBinding_value = map[string]int32{
	"USE_PEER":   0,
	"USE_ORIGIN": 1,
}

func (x PrincipalBinding) String() string {
	return proto.EnumName(PrincipalBinding_name, int32(x))
}

func (PrincipalBinding) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_30ec3f7cef93301a, []int{0}
}

// $hide_from_docs
// Defines the acceptable connection TLS mode.
type MutualTls_Mode int32

const (
	// Client cert must be presented, connection is in TLS.
	MutualTls_STRICT MutualTls_Mode = 0
	// Connection can be either plaintext or TLS with Client cert.
	MutualTls_PERMISSIVE MutualTls_Mode = 1
)

var MutualTls_Mode_name = map[int32]string{
	0: "STRICT",
	1: "PERMISSIVE",
}

var MutualTls_Mode_value = map[string]int32{
	"STRICT":     0,
	"PERMISSIVE": 1,
}

func (x MutualTls_Mode) String() string {
	return proto.EnumName(MutualTls_Mode_name, int32(x))
}

func (MutualTls_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_30ec3f7cef93301a, []int{1, 0}
}

// $hide_from_docs
// Describes how to match a given string. Match is case-sensitive.
type StringMatch struct {
	// Types that are valid to be assigned to MatchType:
	//	*StringMatch_Exact
	//	*StringMatch_Prefix
	//	*StringMatch_Suffix
	//	*StringMatch_Regex
	MatchType            isStringMatch_MatchType `protobuf_oneof:"match_type"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StringMatch) Reset()         { *m = StringMatch{} }
func (m *StringMatch) String() string { return proto.CompactTextString(m) }
func (*StringMatch) ProtoMessage()    {}
func (*StringMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_30ec3f7cef93301a, []int{0}
}

func (m *StringMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringMatch.Unmarshal(m, b)
}
func (m *StringMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringMatch.Marshal(b, m, deterministic)
}
func (m *StringMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMatch.Merge(m, src)
}
func (m *StringMatch) XXX_Size() int {
	return xxx_messageInfo_StringMatch.Size(m)
}
func (m *StringMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMatch.DiscardUnknown(m)
}

var xxx_messageInfo_StringMatch proto.InternalMessageInfo

type isStringMatch_MatchType interface {
	isStringMatch_MatchType()
}

type StringMatch_Exact struct {
	Exact string `protobuf:"bytes,1,opt,name=exact,proto3,oneof"`
}

type StringMatch_Prefix struct {
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof"`
}

type StringMatch_Suffix struct {
	Suffix string `protobuf:"bytes,3,opt,name=suffix,proto3,oneof"`
}

type StringMatch_Regex struct {
	Regex string `protobuf:"bytes,4,opt,name=regex,proto3,oneof"`
}

func (*StringMatch_Exact) isStringMatch_MatchType() {}

func (*StringMatch_Prefix) isStringMatch_MatchType() {}

func (*StringMatch_Suffix) isStringMatch_MatchType() {}

func (*StringMatch_Regex) isStringMatch_MatchType() {}

func (m *StringMatch) GetMatchType() isStringMatch_MatchType {
	if m != nil {
		return m.MatchType
	}
	return nil
}

func (m *StringMatch) GetExact() string {
	if x, ok := m.GetMatchType().(*StringMatch_Exact); ok {
		return x.Exact
	}
	return ""
}

func (m *StringMatch) GetPrefix() string {
	if x, ok := m.GetMatchType().(*StringMatch_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (m *StringMatch) GetSuffix() string {
	if x, ok := m.GetMatchType().(*StringMatch_Suffix); ok {
		return x.Suffix
	}
	return ""
}

func (m *StringMatch) GetRegex() string {
	if x, ok := m.GetMatchType().(*StringMatch_Regex); ok {
		return x.Regex
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StringMatch) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StringMatch_Exact)(nil),
		(*StringMatch_Prefix)(nil),
		(*StringMatch_Suffix)(nil),
		(*StringMatch_Regex)(nil),
	}
}

// $hide_from_docs
// Deprecated. Please use security/v1beta1/PeerAuthentication instead.
// TLS authentication params.
type MutualTls struct {
	// Deprecated. Please use mode = PERMISSIVE instead.
	// If set, will translate to `TLS_PERMISSIVE` mode.
	// Set this flag to true to allow regular TLS (i.e without client x509
	// certificate). If request carries client certificate, identity will be
	// extracted and used (set to peer identity). Otherwise, peer identity will
	// be left unset.
	// When the flag is false (default), request must have client certificate.
	AllowTls bool `protobuf:"varint,1,opt,name=allow_tls,json=allowTls,proto3" json:"allow_tls,omitempty"` // Deprecated: Do not use.
	// Defines the mode of mTLS authentication.
	Mode                 MutualTls_Mode `protobuf:"varint,2,opt,name=mode,proto3,enum=istio.authentication.v1alpha1.MutualTls_Mode" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MutualTls) Reset()         { *m = MutualTls{} }
func (m *MutualTls) String() string { return proto.CompactTextString(m) }
func (*MutualTls) ProtoMessage()    {}
func (*MutualTls) Descriptor() ([]byte, []int) {
	return fileDescriptor_30ec3f7cef93301a, []int{1}
}

func (m *MutualTls) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutualTls.Unmarshal(m, b)
}
func (m *MutualTls) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutualTls.Marshal(b, m, deterministic)
}
func (m *MutualTls) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutualTls.Merge(m, src)
}
func (m *MutualTls) XXX_Size() int {
	return xxx_messageInfo_MutualTls.Size(m)
}
func (m *MutualTls) XXX_DiscardUnknown() {
	xxx_messageInfo_MutualTls.DiscardUnknown(m)
}

var xxx_messageInfo_MutualTls proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *MutualTls) GetAllowTls() bool {
	if m != nil {
		return m.AllowTls
	}
	return false
}

func (m *MutualTls) GetMode() MutualTls_Mode {
	if m != nil {
		return m.Mode
	}
	return MutualTls_STRICT
}

// $hide_from_docs
// Deprecated. Please use security/v1beta1/RequestAuthentication instead.
// JSON Web Token (JWT) token format for authentication as defined by
// [RFC 7519](https://tools.ietf.org/html/rfc7519). See [OAuth 2.0](https://tools.ietf.org/html/rfc6749) and
// [OIDC 1.0](http://openid.net/connect) for how this is used in the whole
// authentication flow.
//
// For example:
//
// A JWT for any requests:
//
// ```yaml
// issuer: https://example.com
// audiences:
// - bookstore_android.apps.googleusercontent.com
//   bookstore_web.apps.googleusercontent.com
// jwksUri: https://example.com/.well-known/jwks.json
// ```
//
// A JWT for all requests except request at path `/health_check` and path with
// prefix `/status/`. This is useful to expose some paths for public access but
// keep others JWT validated.
//
// ```yaml
// issuer: https://example.com
// jwksUri: https://example.com/.well-known/jwks.json
// triggerRules:
// - excludedPaths:
//   - exact: /health_check
//   - prefix: /status/
// ```
//
// A JWT only for requests at path `/admin`. This is useful to only require JWT
// validation on a specific set of paths but keep others public accessible.
//
// ```yaml
// issuer: https://example.com
// jwksUri: https://example.com/.well-known/jwks.json
// triggerRules:
// - includedPaths:
//   - prefix: /admin
// ```
//
// A JWT only for requests at path of prefix `/status/` but except the path of
// `/status/version`. This means for any request path with prefix `/status/` except
// `/status/version` will require a valid JWT to proceed.
//
// ```yaml
// issuer: https://example.com
// jwksUri: https://example.com/.well-known/jwks.json
// triggerRules:
// - excludedPaths:
//   - exact: /status/version
//   includedPaths:
//   - prefix: /status/
// ```
type Jwt struct {
	// Identifies the issuer that issued the JWT. See
	// [issuer](https://tools.ietf.org/html/rfc7519#section-4.1.1)
	// Usually a URL or an email address.
	//
	// Example: https://securetoken.google.com
	// Example: 1234567-compute@developer.gserviceaccount.com
	Issuer string `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// The list of JWT
	// [audiences](https://tools.ietf.org/html/rfc7519#section-4.1.3).
	// that are allowed to access. A JWT containing any of these
	// audiences will be accepted.
	//
	// The service name will be accepted if audiences is empty.
	//
	// Example:
	//
	// ```yaml
	// audiences:
	// - bookstore_android.apps.googleusercontent.com
	//   bookstore_web.apps.googleusercontent.com
	// ```
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
	// Note: Only one of jwks_uri and jwks should be used.
	JwksUri string `protobuf:"bytes,3,opt,name=jwks_uri,json=jwksUri,proto3" json:"jwks_uri,omitempty"`
	// JSON Web Key Set of public keys to validate signature of the JWT.
	// See https://auth0.com/docs/jwks.
	//
	// Note: Only one of jwks_uri and jwks should be used.
	Jwks string `protobuf:"bytes,10,opt,name=jwks,proto3" json:"jwks,omitempty"`
	// JWT is sent in a request header. `header` represents the
	// header name.
	//
	// For example, if `header=x-goog-iap-jwt-assertion`, the header
	// format will be `x-goog-iap-jwt-assertion: <JWT>`.
	JwtHeaders []string `protobuf:"bytes,6,rep,name=jwt_headers,json=jwtHeaders,proto3" json:"jwt_headers,omitempty"`
	// JWT is sent in a query parameter. `query` represents the
	// query parameter name.
	//
	// For example, `query=jwt_token`.
	JwtParams []string `protobuf:"bytes,7,rep,name=jwt_params,json=jwtParams,proto3" json:"jwt_params,omitempty"`
	// List of trigger rules to decide if this JWT should be used to validate the
	// request. The JWT validation happens if any one of the rules matched.
	// If the list is not empty and none of the rules matched, authentication will
	// skip the JWT validation.
	// Leave this empty to always trigger the JWT validation.
	TriggerRules         []*Jwt_TriggerRule `protobuf:"bytes,9,rep,name=trigger_rules,json=triggerRules,proto3" json:"trigger_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Jwt) Reset()         { *m = Jwt{} }
func (m *Jwt) String() string { return proto.CompactTextString(m) }
func (*Jwt) ProtoMessage()    {}
func (*Jwt) Descriptor() ([]byte, []int) {
	return fileDescriptor_30ec3f7cef93301a, []int{2}
}

func (m *Jwt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Jwt.Unmarshal(m, b)
}
func (m *Jwt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Jwt.Marshal(b, m, deterministic)
}
func (m *Jwt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Jwt.Merge(m, src)
}
func (m *Jwt) XXX_Size() int {
	return xxx_messageInfo_Jwt.Size(m)
}
func (m *Jwt) XXX_DiscardUnknown() {
	xxx_messageInfo_Jwt.DiscardUnknown(m)
}

var xxx_messageInfo_Jwt proto.InternalMessageInfo

func (m *Jwt) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *Jwt) GetAudiences() []string {
	if m != nil {
		return m.Audiences
	}
	return nil
}

func (m *Jwt) GetJwksUri() string {
	if m != nil {
		return m.JwksUri
	}
	return ""
}

func (m *Jwt) GetJwks() string {
	if m != nil {
		return m.Jwks
	}
	return ""
}

func (m *Jwt) GetJwtHeaders() []string {
	if m != nil {
		return m.JwtHeaders
	}
	return nil
}

func (m *Jwt) GetJwtParams() []string {
	if m != nil {
		return m.JwtParams
	}
	return nil
}

func (m *Jwt) GetTriggerRules() []*Jwt_TriggerRule {
	if m != nil {
		return m.TriggerRules
	}
	return nil
}

// $hide_from_docs
// Trigger rule to match against a request. The trigger rule is satisfied if
// and only if both rules, excluded_paths and include_paths are satisfied.
type Jwt_TriggerRule struct {
	// List of paths to be excluded from the request. The rule is satisfied if
	// request path does not match to any of the path in this list.
	ExcludedPaths []*StringMatch `protobuf:"bytes,1,rep,name=excluded_paths,json=excludedPaths,proto3" json:"excluded_paths,omitempty"`
	// List of paths that the request must include. If the list is not empty, the
	// rule is satisfied if request path matches at least one of the path in the list.
	// If the list is empty, the rule is ignored, in other words the rule is always satisfied.
	IncludedPaths        []*StringMatch `protobuf:"bytes,2,rep,name=included_paths,json=includedPaths,proto3" json:"included_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Jwt_TriggerRule) Reset()         { *m = Jwt_TriggerRule{} }
func (m *Jwt_TriggerRule) String() string { return proto.CompactTextString(m) }
func (*Jwt_TriggerRule) ProtoMessage()    {}
func (*Jwt_TriggerRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_30ec3f7cef93301a, []int{2, 0}
}

func (m *Jwt_TriggerRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Jwt_TriggerRule.Unmarshal(m, b)
}
func (m *Jwt_TriggerRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Jwt_TriggerRule.Marshal(b, m, deterministic)
}
func (m *Jwt_TriggerRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Jwt_TriggerRule.Merge(m, src)
}
func (m *Jwt_TriggerRule) XXX_Size() int {
	return xxx_messageInfo_Jwt_TriggerRule.Size(m)
}
func (m *Jwt_TriggerRule) XXX_DiscardUnknown() {
	xxx_messageInfo_Jwt_TriggerRule.DiscardUnknown(m)
}

var xxx_messageInfo_Jwt_TriggerRule proto.InternalMessageInfo

func (m *Jwt_TriggerRule) GetExcludedPaths() []*StringMatch {
	if m != nil {
		return m.ExcludedPaths
	}
	return nil
}

func (m *Jwt_TriggerRule) GetIncludedPaths() []*StringMatch {
	if m != nil {
		return m.IncludedPaths
	}
	return nil
}

// $hide_from_docs
// Deprecated. Please use security/v1beta1/PeerAuthentication instead.
// PeerAuthenticationMethod defines one particular type of authentication. Only mTLS is supported
// at the moment.
// The type can be progammatically determine by checking the type of the
// "params" field.
type PeerAuthenticationMethod struct {
	// $hide_from_docs
	//
	// Types that are valid to be assigned to Params:
	//	*PeerAuthenticationMethod_Mtls
	//	*PeerAuthenticationMethod_Jwt
	Params               isPeerAuthenticationMethod_Params `protobuf_oneof:"params"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *PeerAuthenticationMethod) Reset()         { *m = PeerAuthenticationMethod{} }
func (m *PeerAuthenticationMethod) String() string { return proto.CompactTextString(m) }
func (*PeerAuthenticationMethod) ProtoMessage()    {}
func (*PeerAuthenticationMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_30ec3f7cef93301a, []int{3}
}

func (m *PeerAuthenticationMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerAuthenticationMethod.Unmarshal(m, b)
}
func (m *PeerAuthenticationMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerAuthenticationMethod.Marshal(b, m, deterministic)
}
func (m *PeerAuthenticationMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerAuthenticationMethod.Merge(m, src)
}
func (m *PeerAuthenticationMethod) XXX_Size() int {
	return xxx_messageInfo_PeerAuthenticationMethod.Size(m)
}
func (m *PeerAuthenticationMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerAuthenticationMethod.DiscardUnknown(m)
}

var xxx_messageInfo_PeerAuthenticationMethod proto.InternalMessageInfo

type isPeerAuthenticationMethod_Params interface {
	isPeerAuthenticationMethod_Params()
}

type PeerAuthenticationMethod_Mtls struct {
	Mtls *MutualTls `protobuf:"bytes,1,opt,name=mtls,proto3,oneof"`
}

type PeerAuthenticationMethod_Jwt struct {
	Jwt *Jwt `protobuf:"bytes,2,opt,name=jwt,proto3,oneof"`
}

func (*PeerAuthenticationMethod_Mtls) isPeerAuthenticationMethod_Params() {}

func (*PeerAuthenticationMethod_Jwt) isPeerAuthenticationMethod_Params() {}

func (m *PeerAuthenticationMethod) GetParams() isPeerAuthenticationMethod_Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *PeerAuthenticationMethod) GetMtls() *MutualTls {
	if x, ok := m.GetParams().(*PeerAuthenticationMethod_Mtls); ok {
		return x.Mtls
	}
	return nil
}

// Deprecated: Do not use.
func (m *PeerAuthenticationMethod) GetJwt() *Jwt {
	if x, ok := m.GetParams().(*PeerAuthenticationMethod_Jwt); ok {
		return x.Jwt
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PeerAuthenticationMethod) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PeerAuthenticationMethod_Mtls)(nil),
		(*PeerAuthenticationMethod_Jwt)(nil),
	}
}

// $hide_from_docs
// Deprecated. Please use security/v1beta1/RequestAuthentication instead.
// OriginAuthenticationMethod defines authentication method/params for origin
// authentication. Origin could be end-user, device, delegate service etc.
// Currently, only JWT is supported for origin authentication.
type OriginAuthenticationMethod struct {
	// Jwt params for the method.
	Jwt                  *Jwt     `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OriginAuthenticationMethod) Reset()         { *m = OriginAuthenticationMethod{} }
func (m *OriginAuthenticationMethod) String() string { return proto.CompactTextString(m) }
func (*OriginAuthenticationMethod) ProtoMessage()    {}
func (*OriginAuthenticationMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_30ec3f7cef93301a, []int{4}
}

func (m *OriginAuthenticationMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OriginAuthenticationMethod.Unmarshal(m, b)
}
func (m *OriginAuthenticationMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OriginAuthenticationMethod.Marshal(b, m, deterministic)
}
func (m *OriginAuthenticationMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OriginAuthenticationMethod.Merge(m, src)
}
func (m *OriginAuthenticationMethod) XXX_Size() int {
	return xxx_messageInfo_OriginAuthenticationMethod.Size(m)
}
func (m *OriginAuthenticationMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_OriginAuthenticationMethod.DiscardUnknown(m)
}

var xxx_messageInfo_OriginAuthenticationMethod proto.InternalMessageInfo

func (m *OriginAuthenticationMethod) GetJwt() *Jwt {
	if m != nil {
		return m.Jwt
	}
	return nil
}

// $hide_from_docs
// Policy defines what authentication methods can be accepted on workload(s),
// and if authenticated, which method/certificate will set the request principal
// (i.e request.auth.principal attribute).
//
// Authentication policy is composed of 2-part authentication:
// - peer: verify caller service credentials. This part will set source.user
// (peer identity).
// - origin: verify the origin credentials. This part will set request.auth.user
// (origin identity), as well as other attributes like request.auth.presenter,
// request.auth.audiences and raw claims. Note that the identity could be
// end-user, service account, device etc.
//
// Last but not least, the principal binding rule defines which identity (peer
// or origin) should be used as principal. By default, it uses peer.
//
// Examples:
//
// Policy to enable mTLS for all services in namespace frod. The policy name must be
// `default`, and it contains no rule for `targets`.
//
// ```yaml
// apiVersion: authentication.istio.io/v1alpha1
// kind: Policy
// metadata:
//   name: default
//   namespace: frod
// spec:
//   peers:
//   - mtls:
// ```
// Policy to disable mTLS for "productpage" service
//
// ```yaml
// apiVersion: authentication.istio.io/v1alpha1
// kind: Policy
// metadata:
//   name: productpage-mTLS-disable
//   namespace: frod
// spec:
//   targets:
//   - name: productpage
// ```
// Policy to require mTLS for peer authentication, and JWT for origin authentication
// for productpage:9000 except the path '/health_check' . Principal is set from origin identity.
//
// ```yaml
// apiVersion: authentication.istio.io/v1alpha1
// kind: Policy
// metadata:
//   name: productpage-mTLS-with-JWT
//   namespace: frod
// spec:
//   targets:
//   - name: productpage
//     ports:
//     - number: 9000
//   peers:
//   - mtls:
//   origins:
//   - jwt:
//       issuer: "https://securetoken.google.com"
//       audiences:
//       - "productpage"
//       jwksUri: "https://www.googleapis.com/oauth2/v1/certs"
//       jwtHeaders:
//       - "x-goog-iap-jwt-assertion"
//       triggerRules:
//       - excludedPaths:
//         - exact: /health_check
//   principalBinding: USE_ORIGIN
// ```
type Policy struct {
	// Deprecated. Only mesh-level and namespace-level policies are supported.
	// List rules to select workloads that the policy should be applied on.
	// If empty, policy will be used on all workloads in the same namespace.
	Targets []*TargetSelector `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty"` // Deprecated: Do not use.
	// $hide_from_docs
	// Deprecated. Please use security/v1beta1/PeerAuthentication instead.
	// List of authentication methods that can be used for peer authentication.
	// They will be evaluated in order; the first validate one will be used to
	// set peer identity (source.user) and other peer attributes. If none of
	// these methods pass, request will be rejected with authentication failed error (401).
	// Leave the list empty if peer authentication is not required
	Peers []*PeerAuthenticationMethod `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
	// Deprecated. Should set mTLS to PERMISSIVE instead.
	// Set this flag to true to accept request (for peer authentication perspective),
	// even when none of the peer authentication methods defined above satisfied.
	// Typically, this is used to delay the rejection decision to next layer (e.g
	// authorization).
	// This flag is ignored if no authentication defined for peer (peers field is empty).
	PeerIsOptional bool `protobuf:"varint,3,opt,name=peer_is_optional,json=peerIsOptional,proto3" json:"peer_is_optional,omitempty"` // Deprecated: Do not use.
	// Deprecated. Please use security/v1beta1/RequestAuthentication instead.
	// List of authentication methods that can be used for origin authentication.
	// Similar to peers, these will be evaluated in order; the first validate one
	// will be used to set origin identity and attributes (i.e request.auth.user,
	// request.auth.issuer etc). If none of these methods pass, request will be
	// rejected with authentication failed error (401).
	// A method may be skipped, depends on its trigger rule. If all of these methods
	// are skipped, origin authentication will be ignored, as if it is not defined.
	// Leave the list empty if origin authentication is not required.
	Origins []*OriginAuthenticationMethod `protobuf:"bytes,4,rep,name=origins,proto3" json:"origins,omitempty"` // Deprecated: Do not use.
	// Deprecated. Please use security/v1beta1/RequestAuthentication instead.
	// Set this flag to true to accept request (for origin authentication perspective),
	// even when none of the origin authentication methods defined above satisfied.
	// Typically, this is used to delay the rejection decision to next layer (e.g
	// authorization).
	// This flag is ignored if no authentication defined for origin (origins field is empty).
	OriginIsOptional bool `protobuf:"varint,5,opt,name=origin_is_optional,json=originIsOptional,proto3" json:"origin_is_optional,omitempty"` // Deprecated: Do not use.
	// Deprecated. Source principal is always from peer, and request principal is always from
	// RequestAuthentication.
	// Define whether peer or origin identity should be use for principal. Default
	// value is USE_PEER.
	// If peer (or origin) identity is not available, either because of peer/origin
	// authentication is not defined, or failed, principal will be left unset.
	// In other words, binding rule does not affect the decision to accept or
	// reject request.
	PrincipalBinding     PrincipalBinding `protobuf:"varint,6,opt,name=principal_binding,json=principalBinding,proto3,enum=istio.authentication.v1alpha1.PrincipalBinding" json:"principal_binding,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_30ec3f7cef93301a, []int{5}
}

func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *Policy) GetTargets() []*TargetSelector {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *Policy) GetPeers() []*PeerAuthenticationMethod {
	if m != nil {
		return m.Peers
	}
	return nil
}

// Deprecated: Do not use.
func (m *Policy) GetPeerIsOptional() bool {
	if m != nil {
		return m.PeerIsOptional
	}
	return false
}

// Deprecated: Do not use.
func (m *Policy) GetOrigins() []*OriginAuthenticationMethod {
	if m != nil {
		return m.Origins
	}
	return nil
}

// Deprecated: Do not use.
func (m *Policy) GetOriginIsOptional() bool {
	if m != nil {
		return m.OriginIsOptional
	}
	return false
}

// Deprecated: Do not use.
func (m *Policy) GetPrincipalBinding() PrincipalBinding {
	if m != nil {
		return m.PrincipalBinding
	}
	return PrincipalBinding_USE_PEER
}

// $hide_from_docs
// Deprecated. Only support mesh and namespace level policy in the future.
// TargetSelector defines a matching rule to a workload. A workload is selected
// if it is associated with the service name and service port(s) specified in the selector rule.
type TargetSelector struct {
	// The name must be a short name from the service registry. The
	// fully qualified domain name will be resolved in a platform specific manner.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Specifies the ports. Note that this is the port(s) exposed by the service, not workload instance ports.
	// For example, if a service is defined as below, then `8000` should be used, not `9000`.
	// ```yaml
	// kind: Service
	// metadata:
	//   ...
	// spec:
	//   ports:
	//   - name: http
	//     port: 8000
	//     targetPort: 9000
	//   selector:
	//     app: backend
	// ```
	//Leave empty to match all ports that are exposed.
	Ports                []*PortSelector `protobuf:"bytes,2,rep,name=ports,proto3" json:"ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TargetSelector) Reset()         { *m = TargetSelector{} }
func (m *TargetSelector) String() string { return proto.CompactTextString(m) }
func (*TargetSelector) ProtoMessage()    {}
func (*TargetSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_30ec3f7cef93301a, []int{6}
}

func (m *TargetSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetSelector.Unmarshal(m, b)
}
func (m *TargetSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetSelector.Marshal(b, m, deterministic)
}
func (m *TargetSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetSelector.Merge(m, src)
}
func (m *TargetSelector) XXX_Size() int {
	return xxx_messageInfo_TargetSelector.Size(m)
}
func (m *TargetSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetSelector.DiscardUnknown(m)
}

var xxx_messageInfo_TargetSelector proto.InternalMessageInfo

func (m *TargetSelector) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TargetSelector) GetPorts() []*PortSelector {
	if m != nil {
		return m.Ports
	}
	return nil
}

// $hide_from_docs
// Deprecated. Only support mesh and namespace level policy in the future.
// PortSelector specifies the name or number of a port to be used for
// matching targets for authentication policy. This is copied from
// networking API to avoid dependency.
type PortSelector struct {
	// Types that are valid to be assigned to Port:
	//	*PortSelector_Number
	//	*PortSelector_Name
	Port                 isPortSelector_Port `protobuf_oneof:"port"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PortSelector) Reset()         { *m = PortSelector{} }
func (m *PortSelector) String() string { return proto.CompactTextString(m) }
func (*PortSelector) ProtoMessage()    {}
func (*PortSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_30ec3f7cef93301a, []int{7}
}

func (m *PortSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortSelector.Unmarshal(m, b)
}
func (m *PortSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortSelector.Marshal(b, m, deterministic)
}
func (m *PortSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortSelector.Merge(m, src)
}
func (m *PortSelector) XXX_Size() int {
	return xxx_messageInfo_PortSelector.Size(m)
}
func (m *PortSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_PortSelector.DiscardUnknown(m)
}

var xxx_messageInfo_PortSelector proto.InternalMessageInfo

type isPortSelector_Port interface {
	isPortSelector_Port()
}

type PortSelector_Number struct {
	Number uint32 `protobuf:"varint,1,opt,name=number,proto3,oneof"`
}

type PortSelector_Name struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*PortSelector_Number) isPortSelector_Port() {}

func (*PortSelector_Name) isPortSelector_Port() {}

func (m *PortSelector) GetPort() isPortSelector_Port {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *PortSelector) GetNumber() uint32 {
	if x, ok := m.GetPort().(*PortSelector_Number); ok {
		return x.Number
	}
	return 0
}

func (m *PortSelector) GetName() string {
	if x, ok := m.GetPort().(*PortSelector_Name); ok {
		return x.Name
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PortSelector) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PortSelector_Number)(nil),
		(*PortSelector_Name)(nil),
	}
}

func init() {
	proto.RegisterEnum("istio.authentication.v1alpha1.PrincipalBinding", PrincipalBinding_name, PrincipalBinding_value)
	proto.RegisterEnum("istio.authentication.v1alpha1.MutualTls_Mode", MutualTls_Mode_name, MutualTls_Mode_value)
	proto.RegisterType((*StringMatch)(nil), "istio.authentication.v1alpha1.StringMatch")
	proto.RegisterType((*MutualTls)(nil), "istio.authentication.v1alpha1.MutualTls")
	proto.RegisterType((*Jwt)(nil), "istio.authentication.v1alpha1.Jwt")
	proto.RegisterType((*Jwt_TriggerRule)(nil), "istio.authentication.v1alpha1.Jwt.TriggerRule")
	proto.RegisterType((*PeerAuthenticationMethod)(nil), "istio.authentication.v1alpha1.PeerAuthenticationMethod")
	proto.RegisterType((*OriginAuthenticationMethod)(nil), "istio.authentication.v1alpha1.OriginAuthenticationMethod")
	proto.RegisterType((*Policy)(nil), "istio.authentication.v1alpha1.Policy")
	proto.RegisterType((*TargetSelector)(nil), "istio.authentication.v1alpha1.TargetSelector")
	proto.RegisterType((*PortSelector)(nil), "istio.authentication.v1alpha1.PortSelector")
}

func init() {
	proto.RegisterFile("authentication/v1alpha1/policy.proto", fileDescriptor_30ec3f7cef93301a)
}

var fileDescriptor_30ec3f7cef93301a = []byte{
	// 822 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xef, 0x6e, 0x23, 0x35,
	0x10, 0xef, 0x66, 0xb7, 0xdb, 0x64, 0xd2, 0x46, 0xc1, 0x42, 0xc7, 0x52, 0x71, 0x6a, 0xb5, 0x3a,
	0xa1, 0xea, 0x80, 0xe4, 0xae, 0x20, 0x21, 0xf8, 0x80, 0xd4, 0xa0, 0x40, 0x53, 0x14, 0x1a, 0x9c,
	0x1c, 0x48, 0x7c, 0x59, 0x9c, 0xac, 0xbb, 0x71, 0x70, 0xd6, 0x2b, 0xdb, 0x7b, 0xe9, 0xf1, 0x85,
	0x87, 0xe0, 0x05, 0x78, 0x07, 0xc4, 0x67, 0x5e, 0x85, 0x47, 0x41, 0xb6, 0xb3, 0xbd, 0xe4, 0x44,
	0x49, 0xf9, 0xe6, 0xdf, 0xfc, 0xf9, 0xcd, 0x78, 0x66, 0x3c, 0x86, 0x27, 0xa4, 0xd4, 0x73, 0x9a,
	0x6b, 0x36, 0x23, 0x9a, 0x89, 0xbc, 0xfb, 0xf2, 0x39, 0xe1, 0xc5, 0x9c, 0x3c, 0xef, 0x16, 0x82,
	0xb3, 0xd9, 0xab, 0x4e, 0x21, 0x85, 0x16, 0xe8, 0x31, 0x53, 0x9a, 0x89, 0xce, 0xb6, 0x6d, 0xa7,
	0xb2, 0x3d, 0x3e, 0xc9, 0x84, 0xc8, 0x38, 0xed, 0x92, 0x82, 0x75, 0x6f, 0x18, 0xe5, 0x69, 0x32,
	0xa5, 0x73, 0xf2, 0x92, 0x09, 0xe9, 0xfc, 0xe3, 0x5f, 0xa1, 0x39, 0xd6, 0x92, 0xe5, 0xd9, 0x90,
	0xe8, 0xd9, 0x1c, 0x3d, 0x82, 0x7d, 0x7a, 0x4b, 0x66, 0x3a, 0xf2, 0x4e, 0xbd, 0xb3, 0xc6, 0xe5,
	0x1e, 0x76, 0x10, 0x45, 0x10, 0x16, 0x92, 0xde, 0xb0, 0xdb, 0xa8, 0xb6, 0x56, 0xac, 0xb1, 0xd1,
	0xa8, 0xf2, 0xc6, 0x68, 0xfc, 0x4a, 0xe3, 0xb0, 0xe1, 0x92, 0x34, 0xa3, 0xb7, 0x51, 0x50, 0x71,
	0x59, 0xd8, 0x3b, 0x04, 0x58, 0x9a, 0x60, 0x89, 0x7e, 0x55, 0xd0, 0xf8, 0x37, 0x0f, 0x1a, 0xc3,
	0x52, 0x97, 0x84, 0x4f, 0xb8, 0x42, 0x27, 0xd0, 0x20, 0x9c, 0x8b, 0x55, 0xa2, 0xb9, 0xb2, 0x39,
	0xd4, 0x7b, 0xb5, 0xc8, 0xc3, 0x75, 0x2b, 0x34, 0x06, 0x17, 0x10, 0x2c, 0x45, 0x4a, 0x6d, 0x1a,
	0xad, 0xf3, 0x8f, 0x3a, 0xff, 0x79, 0xfd, 0xce, 0x1d, 0x71, 0x67, 0x28, 0x52, 0x8a, 0xad, 0x6b,
	0x1c, 0x43, 0x60, 0x10, 0x02, 0x08, 0xc7, 0x13, 0x3c, 0xf8, 0x72, 0xd2, 0xde, 0x43, 0x2d, 0x80,
	0x51, 0x1f, 0x0f, 0x07, 0xe3, 0xf1, 0xe0, 0xfb, 0x7e, 0xdb, 0x8b, 0xff, 0xf4, 0xc1, 0xbf, 0x5a,
	0x69, 0xf4, 0x08, 0x42, 0xa6, 0x54, 0x49, 0xa5, 0x2b, 0x08, 0x5e, 0x23, 0xf4, 0x1e, 0x34, 0x48,
	0x99, 0x32, 0x9a, 0xcf, 0xa8, 0x8a, 0x6a, 0xa7, 0xfe, 0x59, 0x03, 0xbf, 0x16, 0xa0, 0x77, 0xa1,
	0xbe, 0x58, 0xfd, 0xac, 0x92, 0x52, 0x32, 0x57, 0x15, 0x7c, 0x60, 0xf0, 0x0b, 0xc9, 0x10, 0x82,
	0xc0, 0x1c, 0x23, 0xb0, 0x62, 0x7b, 0x46, 0x27, 0xd0, 0x5c, 0xac, 0x74, 0x32, 0xa7, 0x24, 0xa5,
	0x52, 0x45, 0xa1, 0xa5, 0x83, 0xc5, 0x4a, 0x5f, 0x3a, 0x09, 0x7a, 0x0c, 0x06, 0x25, 0x05, 0x91,
	0x64, 0xa9, 0xa2, 0x03, 0x17, 0x6e, 0xb1, 0xd2, 0x23, 0x2b, 0x40, 0x63, 0x38, 0xd2, 0x92, 0x65,
	0x19, 0x95, 0x89, 0x2c, 0x39, 0x55, 0x51, 0xe3, 0xd4, 0x3f, 0x6b, 0x9e, 0x77, 0x76, 0x14, 0xe7,
	0x6a, 0xa5, 0x3b, 0x13, 0xe7, 0x87, 0x4b, 0x4e, 0xf1, 0xa1, 0x7e, 0x0d, 0xd4, 0xf1, 0x1f, 0x1e,
	0x34, 0x37, 0xb4, 0xe8, 0x3b, 0x68, 0xd1, 0xdb, 0x19, 0x2f, 0x53, 0x9a, 0x26, 0x05, 0xd1, 0x73,
	0xd3, 0x1e, 0x13, 0xe5, 0xe9, 0x8e, 0x28, 0x1b, 0xd3, 0x85, 0x8f, 0x2a, 0x86, 0x91, 0x21, 0x30,
	0x94, 0x2c, 0xdf, 0xa2, 0xac, 0xfd, 0x7f, 0xca, 0x8a, 0xc1, 0x52, 0xc6, 0xbf, 0x7b, 0x10, 0x8d,
	0x28, 0x95, 0x17, 0x5b, 0xae, 0x43, 0xaa, 0xe7, 0x22, 0x45, 0x5f, 0x40, 0xb0, 0xac, 0xe6, 0xaa,
	0x79, 0x7e, 0xf6, 0xd0, 0xd9, 0xb9, 0xdc, 0xc3, 0xd6, 0x0f, 0x7d, 0x0e, 0xfe, 0x62, 0xa5, 0xed,
	0xe8, 0x35, 0xcf, 0xe3, 0xdd, 0xd5, 0x35, 0xa3, 0x7b, 0xb9, 0x87, 0x8d, 0x53, 0xaf, 0x0e, 0xa1,
	0x6b, 0x5f, 0x8c, 0xe1, 0xf8, 0x5a, 0xb2, 0x8c, 0xe5, 0xff, 0x9a, 0xe3, 0x27, 0x2e, 0x86, 0xf7,
	0xd0, 0x18, 0x96, 0x3d, 0xfe, 0xcb, 0x87, 0x70, 0x64, 0xd7, 0x02, 0xfa, 0x06, 0x0e, 0x34, 0x91,
	0x19, 0xd5, 0x55, 0x83, 0x76, 0xbd, 0x91, 0x89, 0xb5, 0x1e, 0x53, 0x4e, 0x67, 0x5a, 0x48, 0xfb,
	0xdc, 0x2a, 0x06, 0x34, 0x84, 0xfd, 0x82, 0x9a, 0x99, 0x74, 0x8d, 0xf9, 0x74, 0x07, 0xd5, 0x7d,
	0x95, 0xc7, 0x8e, 0x05, 0x7d, 0x08, 0x6d, 0x73, 0x48, 0x98, 0x4a, 0x44, 0x61, 0xd4, 0x84, 0xdb,
	0xf7, 0xe1, 0x1e, 0x79, 0xcb, 0xe8, 0x06, 0xea, 0x7a, 0xad, 0x41, 0x3f, 0xc0, 0x81, 0xb0, 0x85,
	0x52, 0x51, 0x60, 0xc3, 0x7f, 0xb6, 0x23, 0xfc, 0xfd, 0x65, 0x75, 0xb7, 0x5a, 0xb3, 0xa1, 0x67,
	0x80, 0xdc, 0x71, 0x2b, 0x91, 0xfd, 0xbb, 0x44, 0xda, 0x4e, 0xbb, 0x91, 0xca, 0x4f, 0xf0, 0x56,
	0x21, 0x59, 0x3e, 0x63, 0x05, 0xe1, 0xc9, 0x94, 0xe5, 0x29, 0xcb, 0xb3, 0x28, 0xb4, 0x2b, 0xa8,
	0xbb, 0xab, 0x26, 0x95, 0x5f, 0xcf, 0xb9, 0xb9, 0x08, 0xc5, 0x1b, 0xd2, 0xf8, 0x17, 0x68, 0x6d,
	0x37, 0x02, 0xbd, 0x03, 0x41, 0x4e, 0x96, 0xd4, 0x2d, 0x9e, 0x9e, 0xff, 0xf7, 0x45, 0x0d, 0x5b,
	0x01, 0xba, 0x80, 0xfd, 0x42, 0x48, 0x5d, 0x35, 0xe5, 0x83, 0x5d, 0x09, 0x08, 0x79, 0x47, 0x8a,
	0x9d, 0xe7, 0x55, 0x50, 0xf7, 0xdb, 0x01, 0x0e, 0x39, 0x99, 0x52, 0xae, 0xe2, 0xaf, 0xe0, 0x70,
	0xd3, 0xc8, 0xac, 0xf4, 0xbc, 0x5c, 0x4e, 0xd7, 0x4b, 0xef, 0xc8, 0xac, 0x74, 0x87, 0xd1, 0xdb,
	0xeb, 0x9c, 0xaa, 0x4f, 0xc0, 0xa2, 0x5e, 0x08, 0x81, 0xa1, 0x7d, 0xfa, 0x0c, 0xda, 0x6f, 0xde,
	0x16, 0x1d, 0x42, 0xfd, 0xc5, 0xb8, 0x9f, 0x8c, 0xfa, 0x7d, 0xec, 0xd6, 0xac, 0x41, 0xd7, 0x78,
	0xf0, 0xf5, 0xe0, 0xdb, 0xb6, 0xd7, 0x7b, 0xff, 0xc7, 0x27, 0x2e, 0x79, 0x26, 0xec, 0x17, 0x75,
	0xcf, 0x97, 0x37, 0x0d, 0xed, 0x67, 0xf5, 0xf1, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x67,
	0xf9, 0xbf, 0x14, 0x07, 0x00, 0x00,
}
