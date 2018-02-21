// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authentication/v1alpha1/policy.proto

/*
Package v1alpha1 is a generated protocol buffer package.

This package defines user-facing authentication policy as well as configs
that the sidecar proxy uses to perform authentication.

It is generated from these files:
	authentication/v1alpha1/policy.proto

It has these top-level messages:
	None
	MutualTls
	Jwt
	SourceAuthenticationMethod
	OriginAuthenticationMethod
	CredentialRule
	Policy
*/
package v1alpha1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import istio_networking_v1alpha3 "istio.io/api/networking/v1alpha3"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Associates authentication with request principal.
type CredentialRule_Binding int32

const (
	// Principal will be set to the identity from source authentication.
	CredentialRule_USE_SOURCE CredentialRule_Binding = 0
	// Principal will be set to the identity from origin authentication.
	CredentialRule_USE_ORIGIN CredentialRule_Binding = 1
)

var CredentialRule_Binding_name = map[int32]string{
	0: "USE_SOURCE",
	1: "USE_ORIGIN",
}
var CredentialRule_Binding_value = map[string]int32{
	"USE_SOURCE": 0,
	"USE_ORIGIN": 1,
}

func (x CredentialRule_Binding) String() string {
	return proto.EnumName(CredentialRule_Binding_name, int32(x))
}
func (CredentialRule_Binding) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

// Placeholder for None authentication params.
type None struct {
}

func (m *None) Reset()                    { *m = None{} }
func (m *None) String() string            { return proto.CompactTextString(m) }
func (*None) ProtoMessage()               {}
func (*None) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Placeholder for mTLS authentication params.
type MutualTls struct {
}

func (m *MutualTls) Reset()                    { *m = MutualTls{} }
func (m *MutualTls) String() string            { return proto.CompactTextString(m) }
func (*MutualTls) ProtoMessage()               {}
func (*MutualTls) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// JSON Web Token (JWT) token format for authentication as defined by
// https://tools.ietf.org/html/rfc7519. See [OAuth
// 2.0](https://tools.ietf.org/html/rfc6749) and [OIDC
// 1.0](http://openid.net/connect) for how this is used in the whole
// authentication flow.
//
// Example,
//
// ```yaml
// issuer: https://example.com
// audiences:
// - bookstore_android.apps.googleusercontent.com
//   bookstore_web.apps.googleusercontent.com
// jwksUri: https://example.com/.well-known/jwks.json
// ```
type Jwt struct {
	// Identifies the issuer that issued the JWT. See
	// [issuer](https://tools.ietf.org/html/rfc7519#section-4.1.1)
	// Usually a URL or an email address.
	//
	// Example: https://securetoken.google.com
	// Example: 1234567-compute@developer.gserviceaccount.com
	Issuer string `protobuf:"bytes,1,opt,name=issuer" json:"issuer,omitempty"`
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
	Audiences []string `protobuf:"bytes,2,rep,name=audiences" json:"audiences,omitempty"`
	// URL of the provider's public key set to validate signature of the
	// JWT. See [OpenID
	// Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	//
	// Optional if the key set document can either (a) be retrieved from
	// [OpenID
	// Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html) of
	// the issuer or (b) inferred from the email domain of the issuer (e.g. a
	// Google service account).
	//
	// Example: https://www.googleapis.com/oauth2/v1/certs
	JwksUri string `protobuf:"bytes,3,opt,name=jwks_uri,json=jwksUri" json:"jwks_uri,omitempty"`
	// JWT is sent in a request header. `header` represents the
	// header name.
	//
	// For example, if `header=x-goog-iap-jwt-assertion`, the header
	// format will be x-goog-iap-jwt-assertion: <JWT>.
	JwtHeaders []string `protobuf:"bytes,6,rep,name=jwt_headers,json=jwtHeaders" json:"jwt_headers,omitempty"`
	// JWT is sent in a query parameter. `query` represents the
	// query parameter name.
	//
	// For example, `query=jwt_token`.
	JwtParams []string `protobuf:"bytes,7,rep,name=jwt_params,json=jwtParams" json:"jwt_params,omitempty"`
}

func (m *Jwt) Reset()                    { *m = Jwt{} }
func (m *Jwt) String() string            { return proto.CompactTextString(m) }
func (*Jwt) ProtoMessage()               {}
func (*Jwt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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

// SourceAuthenticationMethod defines one particular type of authentication, e.g
// mutual TLS, JWT etc, (no authentication is one type by itsefl) that can
// be used for peer authentication.
// The type can be progammatically determine by checking the type of the
// "params" field.
type SourceAuthenticationMethod struct {
	// Types that are valid to be assigned to Params:
	//	*SourceAuthenticationMethod_None
	//	*SourceAuthenticationMethod_Mtls
	//	*SourceAuthenticationMethod_Jwt
	Params isSourceAuthenticationMethod_Params `protobuf_oneof:"params"`
}

func (m *SourceAuthenticationMethod) Reset()                    { *m = SourceAuthenticationMethod{} }
func (m *SourceAuthenticationMethod) String() string            { return proto.CompactTextString(m) }
func (*SourceAuthenticationMethod) ProtoMessage()               {}
func (*SourceAuthenticationMethod) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

<<<<<<< HEAD
<<<<<<< HEAD
type isMechanism_Params interface{ isMechanism_Params() }
=======
type isPeerMechanism_Params interface{ isPeerMechanism_Params() }
>>>>>>> Update authentiation API.
=======
type isSourceAuthenticationMethod_Params interface{ isSourceAuthenticationMethod_Params() }
>>>>>>> Change message and field names. Add more examples.

type SourceAuthenticationMethod_None struct {
	None *None `protobuf:"bytes,1,opt,name=none,oneof"`
}
type SourceAuthenticationMethod_Mtls struct {
	Mtls *MutualTls `protobuf:"bytes,2,opt,name=mtls,oneof"`
}
type SourceAuthenticationMethod_Jwt struct {
	Jwt *Jwt `protobuf:"bytes,3,opt,name=jwt,oneof"`
}

func (*SourceAuthenticationMethod_None) isSourceAuthenticationMethod_Params() {}
func (*SourceAuthenticationMethod_Mtls) isSourceAuthenticationMethod_Params() {}
func (*SourceAuthenticationMethod_Jwt) isSourceAuthenticationMethod_Params()  {}

func (m *SourceAuthenticationMethod) GetParams() isSourceAuthenticationMethod_Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *SourceAuthenticationMethod) GetNone() *None {
	if x, ok := m.GetParams().(*SourceAuthenticationMethod_None); ok {
		return x.None
	}
	return nil
}

func (m *SourceAuthenticationMethod) GetMtls() *MutualTls {
	if x, ok := m.GetParams().(*SourceAuthenticationMethod_Mtls); ok {
		return x.Mtls
	}
	return nil
}

func (m *SourceAuthenticationMethod) GetJwt() *Jwt {
	if x, ok := m.GetParams().(*SourceAuthenticationMethod_Jwt); ok {
		return x.Jwt
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SourceAuthenticationMethod) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SourceAuthenticationMethod_OneofMarshaler, _SourceAuthenticationMethod_OneofUnmarshaler, _SourceAuthenticationMethod_OneofSizer, []interface{}{
		(*SourceAuthenticationMethod_None)(nil),
		(*SourceAuthenticationMethod_Mtls)(nil),
		(*SourceAuthenticationMethod_Jwt)(nil),
	}
}

func _SourceAuthenticationMethod_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SourceAuthenticationMethod)
	// params
	switch x := m.Params.(type) {
	case *SourceAuthenticationMethod_None:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.None); err != nil {
			return err
		}
	case *SourceAuthenticationMethod_Mtls:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Mtls); err != nil {
			return err
		}
	case *SourceAuthenticationMethod_Jwt:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Jwt); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SourceAuthenticationMethod.Params has unexpected type %T", x)
	}
	return nil
}

func _SourceAuthenticationMethod_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SourceAuthenticationMethod)
	switch tag {
	case 1: // params.none
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(None)
		err := b.DecodeMessage(msg)
		m.Params = &SourceAuthenticationMethod_None{msg}
		return true, err
	case 2: // params.mtls
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MutualTls)
		err := b.DecodeMessage(msg)
		m.Params = &SourceAuthenticationMethod_Mtls{msg}
		return true, err
	case 3: // params.jwt
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Jwt)
		err := b.DecodeMessage(msg)
		m.Params = &SourceAuthenticationMethod_Jwt{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SourceAuthenticationMethod_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SourceAuthenticationMethod)
	// params
	switch x := m.Params.(type) {
	case *SourceAuthenticationMethod_None:
		s := proto.Size(x.None)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SourceAuthenticationMethod_Mtls:
		s := proto.Size(x.Mtls)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SourceAuthenticationMethod_Jwt:
		s := proto.Size(x.Jwt)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// OriginAuthenticationMethod defines authentication method/params for origin
// authentication. It should have unique name for referring later in credential
// rules. Currently, only JWT is supported.
type OriginAuthenticationMethod struct {
	// Name that can be used to refer to this mechanism.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Jwt params for the method.
	Jwt *Jwt `protobuf:"bytes,2,opt,name=jwt" json:"jwt,omitempty"`
}

func (m *OriginAuthenticationMethod) Reset()                    { *m = OriginAuthenticationMethod{} }
func (m *OriginAuthenticationMethod) String() string            { return proto.CompactTextString(m) }
func (*OriginAuthenticationMethod) ProtoMessage()               {}
func (*OriginAuthenticationMethod) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *OriginAuthenticationMethod) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OriginAuthenticationMethod) GetJwt() *Jwt {
	if m != nil {
		return m.Jwt
	}
	return nil
}

// CredentialRule defines which identity (e.g from peer or end-user
// authentication) will be used as request principal. The rule can be activated
// conditionally, based on matching condition (currently use only peer identity)
type CredentialRule struct {
	// Defines which authentication (source vs origin) will be binded to
	// request principal.
	Binding CredentialRule_Binding `protobuf:"varint,1,opt,name=binding,enum=istio.authentication.v1alpha1.CredentialRule_Binding" json:"binding,omitempty"`
	// If binding is USE_ORIGIN, this list refers the the origin authentication
	// methods should be considered for the rule (leave the field emtpy to
	// include all origin authentication  methods defined in the policy).
	// At run time, each method will be evaluated in order, until the first valid.
	// request.auth.principal will be set to the identity extracted from that
	// valid certificate. If all methods are invalid, authentication should fail.
	// This field should not be set if binding is USE_SOURCE.
	SelectedOriginMethods []string `protobuf:"bytes,2,rep,name=selected_origin_methods,json=selectedOriginMethods" json:"selected_origin_methods,omitempty"`
	// Condition to activate the rule. If not empty, the rule will be activated
	// if request come from one of these sources (identity).
	// Leave blank to activate the rule unconditionally.
	MatchingSources []string `protobuf:"bytes,3,rep,name=matching_sources,json=matchingSources" json:"matching_sources,omitempty"`
}

func (m *CredentialRule) Reset()                    { *m = CredentialRule{} }
func (m *CredentialRule) String() string            { return proto.CompactTextString(m) }
func (*CredentialRule) ProtoMessage()               {}
func (*CredentialRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CredentialRule) GetBinding() CredentialRule_Binding {
	if m != nil {
		return m.Binding
	}
	return CredentialRule_USE_SOURCE
}

func (m *CredentialRule) GetSelectedOriginMethods() []string {
	if m != nil {
		return m.SelectedOriginMethods
	}
	return nil
}

func (m *CredentialRule) GetMatchingSources() []string {
	if m != nil {
		return m.MatchingSources
	}
	return nil
}

// Policy defines what authentication methods can be accepted on workload(s), and
// if authenticated, which method/certificate will set the request principal
// (i.e request.auth.principal attribute).
//
// Authentication policy is composed of 2-part authentication:
// - source: verify caller service credentials. This part will set source.user
// (source identity).
// - origin: verify the origin credentials. This part will set request.auth.use
// (origin identity), as well as other attributes like request.auth.presenter,
// request.auth.audiences and raw claims. Note that the identity could be
// end-user, service account, device etc.
//
// request.auth.principal will be assigned follow the credential rules. The
// rule also dictates which origin authentication method(s) should run, based
// on source identity.
//
// Examples:
// Policy to enable mTLS for all services in namespace frod
//
// ```yaml
// apiVersion: authentication.istio.io/v1alpha1
// kind: Policy
// metadata:
//   name: mTLS-enable
//   namespace: frod
// spec:
//   destinations:
//   sources:
//   - mtls: null
// ```
// Policy to enable mTLS, and use JWT for productpage:9000.
//
// ```yaml
// apiVersion: authentication.istio.io/v1alpha1
// kind: Policy
// metadata:
//   name: mTLS-enable
//   namespace: frod
// spec:
//   destinations:
//   - name: productpage
//     port:
//       number: 9000
//   sources:
//   - mtls: null
//   origins:
//   - name: google_jwt
//     jwt:
//       issuer: "https://securetoken.google.com"
//       audiences:
//       - "productpage"
//       jwksUri: "https://www.googleapis.com/oauth2/v1/certs"
//       locations:
//       - header: x-goog-iap-jwt-assertion
//   credentialRules:
//   - useOrigins:
//     - google_jwt
// ```
//
// Policy to enable mTLS, and use JWT for productpage:9000 only when caller is
// frontend.serviceaccount.
//
// ```yaml
// apiVersion: authentication.istio.io/v1alpha1
// kind: Policy
// metadata:
//   name: mTLS-enable
//   namespace: frod
// spec:
//   destinations:
//   - name: productpage
//     port:
//       number: 9000
//   sources:
//   - mtls: null
//   origins:
//   - name: google_jwt
//     jwt:
//       issuer: "https://securetoken.google.com"
//       audiences:
//       - "productpage"
//       jwksUri: "https://www.googleapis.com/oauth2/v1/certs"
//       locations:
//       - header: x-goog-iap-jwt-assertion
//   credentialRules:
//   - useOrigins:
//     - google_jwt
//     matchingSources:
//     - frontend.serviceaccount
// ```
//
// Note that a credential rule that uncondtional-use-source (identity)  is
// implicitly check if no rule match, so the above credentialRules is the same
// as this:
//
// ```
// credentialRules:
// - useOrigins:
//   - google_jwt
//   matchingSources:
//   - productpage.serviceaccount
// - useSource: true
// ```
//
// Policy that enable mTLS, requires google JWT if caller is
// frontend.serviceaccount, no JWT (i.e source authentication only) if caller
// is admin, and istio JWT in all other cases.
//
// ```yaml
// apiVersion: authentication.istio.io/v1alpha1
// kind: Policy
// metadata:
//   name: mTLS-enable
//   namespace: frod
// spec:
//   destinations:
//   - name: productpage
//     port:
//       number: 9000
//   sources:
//   - mtls: null
//   origins:
//   - name: google_jwt
//     jwt:
//       issuer: "https://securetoken.google.com"
//       audiences:
//       - "productpage"
//       jwksUri: "https://www.googleapis.com/oauth2/v1/certs"
//       locations:
//       - header: x-goog-iap-jwt-assertion
//   - name: istio_jwt
//     jwt:
//       issuer: "https://securetoken.istio.io"
//       locations:
//       - header: x-istio-jwt-assertion
//   credentialRules:
//   - useOrigins:
//     - google_jwt
//     matchingSources:
//     - productpage.serviceaccount
//   - useSource: true
//     matchingSource:
//     - admin
//   - useOrigins:
//     - istio_jwt
// ```
type Policy struct {
	// List of destinations (workloads) that the policy should be applied on.
	// If empty, policy will be used on all destinations in the same namespace.
<<<<<<< HEAD
	Destinations []*istio_networking_v1alpha3.Destination `protobuf:"bytes,1,rep,name=destinations" json:"destinations,omitempty"`
	// List of credential that should be checked by peer authentication. They
	// will be validated in sequence, until the first one satisfied. If none of
	// the specified mechanism valid, the whole authentication should fail.
	// On the other hand, the first valid credential will be used to extract
	// peer identity (i.e the source.user attribute in the request to Mixer).
	Peers []*PeerMechanism `protobuf:"bytes,2,rep,name=peers" json:"peers,omitempty"`
=======
	Destinations []*istio_routing_v1alpha2.Destination `protobuf:"bytes,1,rep,name=destinations" json:"destinations,omitempty"`
	// List of authentication methods that can be use for source authentication.
	// They will be evaluated in order, until the first one satisfied; source
	// identity is then extracted from the associated certificate. On the other
	// hand, if none of these methods pass, request should be rejected with
	// authentication faile error (401).
	// Leave the list empty if no source authentication required, or have single
	// entry of method 'None'. The source.user attribute will not be set in that
	// case.
	Sources []*SourceAuthenticationMethod `protobuf:"bytes,2,rep,name=sources" json:"sources,omitempty"`
>>>>>>> Change message and field names. Add more examples.
	// List of supported end-user authentication. Which one(s) are used for a specific
	// request is decided at runtime, based on credential_rules bellow.
	Origins []*OriginAuthenticationMethod `protobuf:"bytes,3,rep,name=origins" json:"origins,omitempty"`
	// Rules to define how request pricipal will be set. Each rule can have
	// conditions that determine if the rule should be applied or not. The rule
	// will be checked for matching conditions at runtime, in order, and stop at
	// the first match. If there are no rule matching condtion, source identity
	// will be used as principal (in the other words, the credential rule with
	// use_source = true with no matching condition is implicitly added the end
	// of the list.)
	CredentialRules []*CredentialRule `protobuf:"bytes,4,rep,name=credential_rules,json=credentialRules" json:"credential_rules,omitempty"`
}

func (m *Policy) Reset()                    { *m = Policy{} }
func (m *Policy) String() string            { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()               {}
func (*Policy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Policy) GetDestinations() []*istio_networking_v1alpha3.Destination {
	if m != nil {
		return m.Destinations
	}
	return nil
}

func (m *Policy) GetSources() []*SourceAuthenticationMethod {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *Policy) GetOrigins() []*OriginAuthenticationMethod {
	if m != nil {
		return m.Origins
	}
	return nil
}

func (m *Policy) GetCredentialRules() []*CredentialRule {
	if m != nil {
		return m.CredentialRules
	}
	return nil
}

func init() {
	proto.RegisterType((*None)(nil), "istio.authentication.v1alpha1.None")
	proto.RegisterType((*MutualTls)(nil), "istio.authentication.v1alpha1.MutualTls")
	proto.RegisterType((*Jwt)(nil), "istio.authentication.v1alpha1.Jwt")
	proto.RegisterType((*SourceAuthenticationMethod)(nil), "istio.authentication.v1alpha1.SourceAuthenticationMethod")
	proto.RegisterType((*OriginAuthenticationMethod)(nil), "istio.authentication.v1alpha1.OriginAuthenticationMethod")
	proto.RegisterType((*CredentialRule)(nil), "istio.authentication.v1alpha1.CredentialRule")
	proto.RegisterType((*Policy)(nil), "istio.authentication.v1alpha1.Policy")
	proto.RegisterEnum("istio.authentication.v1alpha1.CredentialRule_Binding", CredentialRule_Binding_name, CredentialRule_Binding_value)
}

func init() { proto.RegisterFile("authentication/v1alpha1/policy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x86, 0xa3, 0xc8, 0x55, 0xac, 0x51, 0x4f, 0x7b, 0x28, 0x6a, 0x69, 0xa8, 0x51, 0x43, 0x70,
	0x2f, 0x12, 0xb1, 0xa1, 0xd0, 0x4b, 0x0e, 0xa1, 0x05, 0x63, 0x48, 0x09, 0xa2, 0xb9, 0xf4, 0x22,
	0xb6, 0xd2, 0x50, 0x8f, 0x22, 0xef, 0x8a, 0xdd, 0x95, 0x45, 0x5f, 0xa4, 0xcf, 0xd4, 0x77, 0xe8,
	0xcb, 0x94, 0x5d, 0x45, 0x35, 0x3e, 0xb4, 0x26, 0xc7, 0x99, 0x9d, 0xef, 0x63, 0xe7, 0xdf, 0x85,
	0x0b, 0xde, 0x99, 0x0d, 0x0a, 0x43, 0x25, 0x37, 0x24, 0x45, 0xb6, 0xbb, 0xe2, 0x4d, 0xbb, 0xe1,
	0x57, 0x59, 0x2b, 0x1b, 0x2a, 0x7f, 0xa4, 0xad, 0x92, 0x46, 0xb2, 0x73, 0xd2, 0x86, 0x64, 0x7a,
	0x38, 0x9b, 0x8e, 0xb3, 0xaf, 0xde, 0x09, 0x34, 0xbd, 0x54, 0x0f, 0x24, 0xbe, 0x8f, 0x82, 0x65,
	0xb6, 0x23, 0x65, 0x3a, 0xde, 0x14, 0x1a, 0xd5, 0x8e, 0x4a, 0x1c, 0x4c, 0x49, 0x00, 0x93, 0xcf,
	0x52, 0x60, 0x12, 0x41, 0x78, 0xdb, 0xd9, 0xf3, 0x2f, 0x8d, 0x4e, 0x7e, 0x7a, 0xe0, 0xaf, 0x7b,
	0xc3, 0x5e, 0x40, 0x40, 0x5a, 0x77, 0xa8, 0x62, 0x6f, 0xe6, 0xcd, 0xc3, 0xfc, 0xb1, 0x62, 0xaf,
	0x21, 0xe4, 0x5d, 0x45, 0x28, 0x4a, 0xd4, 0xf1, 0xe9, 0xcc, 0x9f, 0x87, 0xf9, 0xbe, 0xc1, 0x5e,
	0xc2, 0xb4, 0xee, 0x1f, 0x74, 0xd1, 0x29, 0x8a, 0x7d, 0xc7, 0x9d, 0xd9, 0xfa, 0x5e, 0x11, 0x7b,
	0x03, 0x51, 0xdd, 0x9b, 0x62, 0x83, 0xbc, 0x42, 0xa5, 0xe3, 0xc0, 0xa1, 0x50, 0xf7, 0x66, 0x35,
	0x74, 0xd8, 0x39, 0xd8, 0xaa, 0x68, 0xb9, 0xe2, 0x5b, 0x1d, 0x9f, 0x0d, 0xea, 0xba, 0x37, 0x77,
	0xae, 0x91, 0xfc, 0xf2, 0x20, 0xbc, 0xc5, 0x72, 0xc3, 0x05, 0xe9, 0x2d, 0xfb, 0x00, 0x13, 0x21,
	0x05, 0xba, 0xcb, 0x45, 0x8b, 0xb7, 0xe9, 0x7f, 0x43, 0x49, 0xed, 0x9a, 0xab, 0x93, 0xdc, 0x21,
	0xec, 0x1a, 0x26, 0x5b, 0xd3, 0xd8, 0xcb, 0x5b, 0x74, 0x7e, 0x04, 0xfd, 0x9b, 0x8c, 0xe5, 0x2d,
	0xc7, 0xde, 0x83, 0x5f, 0xf7, 0xc6, 0xad, 0x17, 0x2d, 0x92, 0x23, 0xf8, 0xba, 0x37, 0xab, 0x93,
	0xdc, 0x02, 0x37, 0x53, 0x08, 0x86, 0xdd, 0x92, 0xdf, 0x1e, 0x04, 0x77, 0xee, 0x4d, 0xd9, 0x1a,
	0x9e, 0x57, 0xa8, 0x0d, 0x09, 0xc7, 0xe9, 0xd8, 0x9b, 0xf9, 0xf3, 0x68, 0x71, 0xf9, 0x68, 0xdd,
	0xbf, 0xe5, 0x68, 0x5c, 0xa6, 0x1f, 0xf7, 0xe3, 0xf9, 0x01, 0xcb, 0xae, 0xe1, 0x59, 0x8b, 0x36,
	0xdb, 0x53, 0x27, 0x39, 0xba, 0xd9, 0x18, 0x66, 0x3e, 0x60, 0xec, 0x13, 0x84, 0x28, 0xaa, 0xa2,
	0xd3, 0xd6, 0xe1, 0x3f, 0xd1, 0x31, 0x45, 0x51, 0xdd, 0x5b, 0xf2, 0xe6, 0xf2, 0xeb, 0xc5, 0x00,
	0x91, 0xcc, 0x78, 0x4b, 0xd9, 0x3f, 0x7e, 0xf5, 0xb7, 0xc0, 0xfd, 0xc2, 0xe5, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x83, 0xd1, 0x9f, 0x7b, 0xf7, 0x02, 0x00, 0x00,
=======
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0xeb, 0x24, 0x9f, 0x9b, 0x4c, 0xbe, 0x96, 0x6a, 0x2f, 0x90, 0x39, 0x54, 0x14, 0x53,
	0x50, 0x2e, 0xc0, 0x56, 0x83, 0x84, 0xc4, 0x0d, 0x17, 0x01, 0x44, 0x54, 0x51, 0x14, 0x59, 0xad,
	0x84, 0xb8, 0xb1, 0x16, 0x7b, 0xd4, 0xac, 0xb1, 0xd7, 0xd6, 0x1e, 0xb0, 0x78, 0x11, 0x5e, 0x80,
	0xe7, 0xe0, 0x0d, 0x78, 0x28, 0xb4, 0xeb, 0x9c, 0x8c, 0x04, 0x16, 0xe2, 0x2e, 0x33, 0x9a, 0xdf,
	0x3f, 0xf3, 0x9f, 0x99, 0x35, 0x9c, 0x52, 0xad, 0x96, 0xc8, 0x15, 0x4b, 0xa8, 0x62, 0x25, 0x0f,
	0x3f, 0x9f, 0xd1, 0xbc, 0x5a, 0xd2, 0xb3, 0xb0, 0x2a, 0x73, 0x96, 0x7c, 0x09, 0x2a, 0x51, 0xaa,
	0x92, 0x1c, 0x33, 0xa9, 0x58, 0x19, 0xb4, 0x6b, 0x83, 0x75, 0xed, 0xed, 0xfb, 0xa2, 0xd4, 0x8a,
	0xf1, 0xeb, 0x35, 0x3d, 0x0d, 0x4d, 0x02, 0x63, 0xa1, 0x73, 0x6c, 0x14, 0x7c, 0x17, 0x06, 0xef,
	0x4a, 0x8e, 0xfe, 0x18, 0x46, 0x17, 0x5a, 0x69, 0x9a, 0x5f, 0xe6, 0xd2, 0xff, 0xea, 0x40, 0xff,
	0xbc, 0x56, 0xe4, 0x26, 0xb8, 0x4c, 0x4a, 0x8d, 0xc2, 0x73, 0x4e, 0x9c, 0xc9, 0x28, 0x5a, 0x45,
	0xe4, 0x2e, 0x8c, 0xa8, 0x4e, 0x19, 0xf2, 0x04, 0xa5, 0xd7, 0x3b, 0xe9, 0x4f, 0x46, 0xd1, 0x36,
	0x41, 0x6e, 0xc1, 0x30, 0xab, 0x3f, 0xc9, 0x58, 0x0b, 0xe6, 0xf5, 0x2d, 0xb7, 0x6f, 0xe2, 0x2b,
	0xc1, 0xc8, 0x3d, 0x18, 0x67, 0xb5, 0x8a, 0x97, 0x48, 0x53, 0x14, 0xd2, 0x73, 0x2d, 0x0a, 0x59,
	0xad, 0xe6, 0x4d, 0x86, 0x1c, 0x83, 0x89, 0xe2, 0x8a, 0x0a, 0x5a, 0x48, 0x6f, 0xbf, 0x91, 0xce,
	0x6a, 0xb5, 0xb0, 0x09, 0xff, 0x87, 0x03, 0x07, 0x0b, 0x44, 0x71, 0x81, 0xc9, 0x92, 0x72, 0x26,
	0x0b, 0xf2, 0x1c, 0x06, 0xbc, 0xe4, 0x68, 0x1b, 0x1c, 0x4f, 0x1f, 0x04, 0x7f, 0x1c, 0x48, 0x60,
	0xac, 0xce, 0xf7, 0x22, 0x8b, 0x90, 0x17, 0x30, 0x28, 0x54, 0x6e, 0x0c, 0x18, 0x74, 0xd2, 0x81,
	0x6e, 0xa6, 0x63, 0x78, 0xc3, 0x91, 0x67, 0xd0, 0xcf, 0x6a, 0x65, 0x2d, 0x8e, 0xa7, 0x7e, 0x07,
	0x7e, 0x5e, 0xab, 0xf9, 0x5e, 0x64, 0x80, 0xd9, 0x10, 0xdc, 0xc6, 0x9f, 0xff, 0xcd, 0x81, 0xa3,
	0xd7, 0x3c, 0xbd, 0x92, 0xbb, 0x8e, 0x08, 0x0c, 0x38, 0x2d, 0x70, 0x35, 0x72, 0xfb, 0x7b, 0xe3,
	0xb2, 0xf7, 0xf7, 0x2e, 0xff, 0xbd, 0xcb, 0x4b, 0x38, 0x7c, 0x29, 0x30, 0x35, 0xf5, 0x34, 0x8f,
	0x74, 0x8e, 0xe4, 0x0e, 0x8c, 0x90, 0xa7, 0xb1, 0x96, 0x66, 0x89, 0x8e, 0x5d, 0xd2, 0x10, 0x1b,
	0x1f, 0x92, 0x3c, 0x84, 0xc3, 0x82, 0xaa, 0x64, 0xc9, 0xf8, 0x75, 0x5c, 0xa1, 0xa9, 0x68, 0x2e,
	0xe4, 0x60, 0x9d, 0x35, 0x0b, 0x94, 0xfe, 0xf7, 0x1e, 0xb8, 0x0b, 0x7b, 0xcb, 0xe4, 0x0d, 0xfc,
	0x9f, 0xa2, 0x54, 0x8c, 0xdb, 0x6e, 0x1a, 0xc5, 0xad, 0xcb, 0xd5, 0x0d, 0xaf, 0x9b, 0x9c, 0x06,
	0xaf, 0xb6, 0xb5, 0x51, 0x0b, 0x24, 0x33, 0xf8, 0x6f, 0xfb, 0x8f, 0xe3, 0xe9, 0xe3, 0x0e, 0xb7,
	0xad, 0x4b, 0x8a, 0x1a, 0x94, 0xbc, 0xdd, 0xf5, 0xd6, 0xb7, 0x3a, 0x61, 0x87, 0xce, 0xaf, 0x2b,
	0xdc, 0x19, 0xc6, 0x7b, 0x38, 0x4a, 0x36, 0xb3, 0xb3, 0xef, 0x4e, 0x7a, 0x03, 0x2b, 0xfa, 0xa4,
	0x43, 0xb4, 0x3d, 0xf2, 0xe8, 0x46, 0xd2, 0x8a, 0xe5, 0xec, 0xd1, 0x87, 0xd3, 0x46, 0x80, 0x95,
	0x21, 0xad, 0x58, 0xf8, 0x9b, 0xef, 0xc5, 0x47, 0xd7, 0xbe, 0xf3, 0xa7, 0x3f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x8e, 0xff, 0x77, 0xe1, 0x51, 0x04, 0x00, 0x00,
>>>>>>> Update authentiation API.
=======
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdf, 0x6b, 0x13, 0x41,
	0x10, 0x6e, 0x72, 0xe1, 0x92, 0x4c, 0xc4, 0x96, 0x7d, 0x90, 0x33, 0x58, 0x8c, 0x67, 0x91, 0xf8,
	0xe0, 0x85, 0x46, 0x11, 0xfa, 0x22, 0x58, 0x05, 0x43, 0xa1, 0x5a, 0xb6, 0x0a, 0xe2, 0x4b, 0x58,
	0x2f, 0x6b, 0x6e, 0xcf, 0xcb, 0xee, 0xb1, 0x3f, 0x3c, 0xfa, 0x8f, 0xf8, 0x9f, 0xf9, 0xe0, 0x7f,
	0x23, 0xbb, 0x9b, 0x6b, 0x3d, 0x30, 0x5e, 0xdf, 0x32, 0x93, 0xf9, 0xbe, 0x99, 0xef, 0x9b, 0xd9,
	0x83, 0x23, 0x62, 0x74, 0x46, 0xb9, 0x66, 0x29, 0xd1, 0x4c, 0xf0, 0xd9, 0x8f, 0x63, 0x52, 0x94,
	0x19, 0x39, 0x9e, 0x95, 0xa2, 0x60, 0xe9, 0x55, 0x52, 0x4a, 0xa1, 0x05, 0x3a, 0x64, 0x4a, 0x33,
	0x91, 0x34, 0x6b, 0x93, 0xba, 0x76, 0xfc, 0x48, 0x0a, 0xa3, 0x19, 0x5f, 0xd7, 0xe8, 0xf9, 0xcc,
	0x26, 0xe8, 0x52, 0x9a, 0x82, 0x7a, 0x86, 0x38, 0x84, 0xde, 0x7b, 0xc1, 0x69, 0x3c, 0x82, 0xe1,
	0xb9, 0xd1, 0x86, 0x14, 0x1f, 0x0b, 0x15, 0xff, 0xec, 0x40, 0x70, 0x56, 0x69, 0x74, 0x0f, 0x42,
	0xa6, 0x94, 0xa1, 0x32, 0xea, 0x4c, 0x3a, 0xd3, 0x21, 0xde, 0x46, 0xe8, 0x01, 0x0c, 0x89, 0x59,
	0x31, 0xca, 0x53, 0xaa, 0xa2, 0xee, 0x24, 0x98, 0x0e, 0xf1, 0x4d, 0x02, 0xdd, 0x87, 0x41, 0x5e,
	0x7d, 0x57, 0x4b, 0x23, 0x59, 0x14, 0x38, 0x5c, 0xdf, 0xc6, 0x9f, 0x24, 0x43, 0x0f, 0x61, 0x94,
	0x57, 0x7a, 0x99, 0x51, 0xb2, 0xa2, 0x52, 0x45, 0xa1, 0x83, 0x42, 0x5e, 0xe9, 0x85, 0xcf, 0xa0,
	0x43, 0xb0, 0xd1, 0xb2, 0x24, 0x92, 0x6c, 0x54, 0xd4, 0xf7, 0xd4, 0x79, 0xa5, 0x2f, 0x5c, 0x22,
	0xfe, 0xdd, 0x81, 0xf1, 0xa5, 0x30, 0x32, 0xa5, 0xaf, 0x1b, 0x92, 0xcf, 0xa9, 0xce, 0xc4, 0x0a,
	0x9d, 0x40, 0x8f, 0x0b, 0x4e, 0xdd, 0xb4, 0xa3, 0xf9, 0xe3, 0xe4, 0xbf, 0xee, 0x24, 0x56, 0xf7,
	0x62, 0x0f, 0x3b, 0x08, 0x7a, 0x05, 0xbd, 0x8d, 0x2e, 0xac, 0x1a, 0x0b, 0x9d, 0xb6, 0x40, 0xaf,
	0xad, 0xb2, 0x78, 0x8b, 0x43, 0x2f, 0x21, 0xc8, 0x2b, 0xed, 0xf4, 0x8e, 0xe6, 0x71, 0x0b, 0xfc,
	0xac, 0xd2, 0x8b, 0x3d, 0x6c, 0x01, 0xa7, 0x03, 0x08, 0xbd, 0xd8, 0xf8, 0x1b, 0x8c, 0x3f, 0x48,
	0xb6, 0x66, 0xfc, 0x9f, 0xd2, 0x10, 0xf4, 0x38, 0xd9, 0xd0, 0xed, 0x22, 0xdc, 0x6f, 0xf4, 0xc2,
	0xf7, 0xec, 0xde, 0xb6, 0xa7, 0xeb, 0x18, 0x5f, 0xc1, 0xdd, 0x37, 0x92, 0xae, 0x6c, 0x0d, 0x29,
	0xb0, 0x29, 0xa8, 0x35, 0xdd, 0x28, 0xba, 0x54, 0xce, 0x58, 0xd7, 0x61, 0x80, 0x87, 0x46, 0x51,
	0xef, 0x74, 0xfd, 0xb7, 0x70, 0xc3, 0xd5, 0xeb, 0x36, 0x8a, 0xfa, 0x69, 0xd1, 0x53, 0x38, 0xd8,
	0x10, 0x9d, 0x66, 0x8c, 0xaf, 0xb7, 0x14, 0x2a, 0x0a, 0x5c, 0xd1, 0x7e, 0x9d, 0xf7, 0x44, 0x2a,
	0xfe, 0xd5, 0x85, 0xf0, 0xc2, 0xdd, 0x2f, 0x7a, 0x07, 0x77, 0x56, 0x54, 0x69, 0xc6, 0xdd, 0x98,
	0x2a, 0xea, 0x4c, 0x82, 0xbf, 0x56, 0xb6, 0xbd, 0xdb, 0x7a, 0xfa, 0x79, 0xf2, 0xf6, 0xa6, 0x16,
	0x37, 0x80, 0xe8, 0x12, 0xfa, 0x75, 0xd7, 0xae, 0xe3, 0x38, 0x69, 0x31, 0x62, 0xf7, 0xfd, 0xe0,
	0x9a, 0xc9, 0x92, 0x7a, 0xb9, 0x5e, 0x4a, 0x3b, 0xe9, 0xee, 0xcd, 0xe1, 0x9a, 0x09, 0x7d, 0x86,
	0x83, 0xf4, 0xda, 0x78, 0xf7, 0x06, 0x55, 0xd4, 0x73, 0xec, 0xcf, 0x5a, 0xd8, 0x9b, 0xfb, 0xc2,
	0xfb, 0x69, 0x23, 0x56, 0xa7, 0x4f, 0xbe, 0x1c, 0x79, 0x02, 0x26, 0x66, 0xa4, 0x64, 0xb3, 0x1d,
	0xdf, 0x8e, 0xaf, 0xa1, 0x7b, 0xf3, 0xcf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x24, 0x7f, 0xc0,
	0x2a, 0x5d, 0x04, 0x00, 0x00,
>>>>>>> Change message and field names. Add more examples.
=======
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdf, 0x8b, 0xd3, 0x40,
	0x10, 0xbe, 0x36, 0x25, 0x6d, 0xa7, 0xe2, 0x1d, 0xfb, 0x20, 0xb1, 0x78, 0x58, 0xe3, 0x21, 0xf5,
	0xc1, 0x84, 0xab, 0x22, 0xdc, 0x8b, 0xe0, 0x29, 0x58, 0x0e, 0x4e, 0x8f, 0x3d, 0x05, 0xf1, 0x25,
	0xac, 0xc9, 0xda, 0x6c, 0x4c, 0x77, 0xc3, 0xfe, 0x30, 0x88, 0xff, 0x87, 0xff, 0x99, 0x0f, 0xfe,
	0x37, 0xb2, 0x9b, 0xe6, 0xce, 0x80, 0xb5, 0xf7, 0xd6, 0x99, 0xce, 0xf7, 0xcd, 0x7c, 0xdf, 0xcc,
	0x06, 0x8e, 0x88, 0xd1, 0x39, 0xe5, 0x9a, 0xa5, 0x44, 0x33, 0xc1, 0xe3, 0x6f, 0xc7, 0xa4, 0xac,
	0x72, 0x72, 0x1c, 0x57, 0xa2, 0x64, 0xe9, 0xf7, 0xa8, 0x92, 0x42, 0x0b, 0x74, 0xc8, 0x94, 0x66,
	0x22, 0xea, 0xd6, 0x46, 0x6d, 0xed, 0xf4, 0x81, 0x14, 0x46, 0x33, 0xbe, 0x6a, 0xd1, 0x8b, 0xd8,
	0x26, 0x68, 0x22, 0x4d, 0x49, 0x1b, 0x86, 0xd0, 0x87, 0xc1, 0x5b, 0xc1, 0x69, 0x38, 0x81, 0xf1,
	0xb9, 0xd1, 0x86, 0x94, 0xef, 0x4b, 0x15, 0xfe, 0xec, 0x81, 0x77, 0x56, 0x6b, 0x74, 0x07, 0x7c,
	0xa6, 0x94, 0xa1, 0x32, 0xe8, 0xcd, 0x7a, 0xf3, 0x31, 0xde, 0x44, 0xe8, 0x1e, 0x8c, 0x89, 0xc9,
	0x18, 0xe5, 0x29, 0x55, 0x41, 0x7f, 0xe6, 0xcd, 0xc7, 0xf8, 0x3a, 0x81, 0xee, 0xc2, 0xa8, 0xa8,
	0xbf, 0xaa, 0xc4, 0x48, 0x16, 0x78, 0x0e, 0x37, 0xb4, 0xf1, 0x07, 0xc9, 0xd0, 0x7d, 0x98, 0x14,
	0xb5, 0x4e, 0x72, 0x4a, 0x32, 0x2a, 0x55, 0xe0, 0x3b, 0x28, 0x14, 0xb5, 0x5e, 0x36, 0x19, 0x74,
	0x08, 0x36, 0x4a, 0x2a, 0x22, 0xc9, 0x5a, 0x05, 0xc3, 0x86, 0xba, 0xa8, 0xf5, 0x85, 0x4b, 0x84,
	0xbf, 0x7b, 0x30, 0xbd, 0x14, 0x46, 0xa6, 0xf4, 0x65, 0x47, 0xf2, 0x39, 0xd5, 0xb9, 0xc8, 0xd0,
	0x09, 0x0c, 0xb8, 0xe0, 0xd4, 0x4d, 0x3b, 0x59, 0x3c, 0x8c, 0xfe, 0xeb, 0x4e, 0x64, 0x75, 0x2f,
	0xf7, 0xb0, 0x83, 0xa0, 0x17, 0x30, 0x58, 0xeb, 0xd2, 0xaa, 0xb1, 0xd0, 0xf9, 0x0e, 0xe8, 0x95,
	0x55, 0x16, 0x6f, 0x71, 0xe8, 0x39, 0x78, 0x45, 0xad, 0x9d, 0xde, 0xc9, 0x22, 0xdc, 0x01, 0x3f,
	0xab, 0xf5, 0x72, 0x0f, 0x5b, 0xc0, 0xe9, 0x08, 0xfc, 0x46, 0x6c, 0xf8, 0x05, 0xa6, 0xef, 0x24,
	0x5b, 0x31, 0xfe, 0x4f, 0x69, 0x08, 0x06, 0x9c, 0xac, 0xe9, 0x66, 0x11, 0xee, 0x37, 0x7a, 0xd6,
	0xf4, 0xec, 0xdf, 0xb4, 0xa7, 0xeb, 0x18, 0xfe, 0x80, 0xdb, 0xaf, 0x24, 0xcd, 0x6c, 0x0d, 0x29,
	0xb1, 0x29, 0xa9, 0x35, 0xdd, 0x28, 0x9a, 0x28, 0x67, 0xac, 0xeb, 0x30, 0xc2, 0x63, 0xa3, 0x68,
	0xe3, 0xb4, 0x5d, 0x9a, 0xfd, 0x5b, 0xb8, 0xe1, 0xda, 0x7d, 0x5b, 0x44, 0x33, 0xae, 0x42, 0x8f,
	0xe1, 0x60, 0x4d, 0x74, 0x9a, 0x33, 0xbe, 0xda, 0x90, 0xa8, 0xc0, 0x73, 0x55, 0xfb, 0x6d, 0xbe,
	0xa1, 0x52, 0xe1, 0xaf, 0x3e, 0xf8, 0x17, 0xee, 0x82, 0xd1, 0x1b, 0xb8, 0x95, 0x51, 0xa5, 0x19,
	0x77, 0x83, 0xaa, 0xa0, 0x37, 0xf3, 0xfe, 0x5a, 0xda, 0xe6, 0x72, 0xdb, 0xf9, 0x17, 0xd1, 0xeb,
	0xeb, 0x5a, 0xdc, 0x01, 0xa2, 0x4b, 0x18, 0xb6, 0x5d, 0xfb, 0x8e, 0xe3, 0x64, 0x87, 0x15, 0xdb,
	0x2f, 0x08, 0xb7, 0x4c, 0x96, 0xb4, 0x15, 0xec, 0xdd, 0x88, 0x74, 0xfb, 0xee, 0x70, 0xcb, 0x84,
	0x3e, 0xc2, 0x41, 0x7a, 0x65, 0xbd, 0x7b, 0x85, 0x2a, 0x18, 0x38, 0xf6, 0x27, 0x3b, 0xd8, 0xbb,
	0x1b, 0xc3, 0xfb, 0x69, 0x27, 0x56, 0xa7, 0x8f, 0x3e, 0x1d, 0x35, 0x04, 0x4c, 0xc4, 0xa4, 0x62,
	0xf1, 0x96, 0xaf, 0xc7, 0x67, 0xdf, 0xbd, 0xfa, 0xa7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x86,
	0xb4, 0xfa, 0x1b, 0x5f, 0x04, 0x00, 0x00,
>>>>>>> Change use_origin to plural.
=======
	// 572 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0x97, 0xa6, 0x4a, 0xd7, 0x2b, 0xda, 0x2a, 0x4b, 0x40, 0xa8, 0x98, 0x28, 0x61, 0x42,
	0xdd, 0x03, 0xa9, 0x56, 0x60, 0xd2, 0x5e, 0x90, 0xe8, 0x98, 0xd6, 0x4d, 0xda, 0x3a, 0xb9, 0x54,
	0x42, 0xbc, 0x44, 0x5e, 0x62, 0x5a, 0x97, 0xd4, 0xae, 0x62, 0x87, 0x88, 0x2f, 0xc2, 0x37, 0xe3,
	0x81, 0x4f, 0xc1, 0x57, 0x40, 0xb1, 0xeb, 0x8d, 0x4a, 0x94, 0xc2, 0x5b, 0xef, 0xea, 0xdf, 0xff,
	0xf2, 0xbf, 0x3b, 0x1b, 0xf6, 0x49, 0xae, 0xa6, 0x94, 0x2b, 0x16, 0x13, 0xc5, 0x04, 0xef, 0x7e,
	0x39, 0x24, 0xe9, 0x62, 0x4a, 0x0e, 0xbb, 0x0b, 0x91, 0xb2, 0xf8, 0x6b, 0xb8, 0xc8, 0x84, 0x12,
	0x68, 0x8f, 0x49, 0xc5, 0x44, 0xb8, 0x7a, 0x36, 0xb4, 0x67, 0x5b, 0x4f, 0x33, 0x91, 0x2b, 0xc6,
	0x27, 0x96, 0xee, 0x75, 0xcb, 0x04, 0x8d, 0xb2, 0x3c, 0xa5, 0x46, 0x21, 0xf0, 0xa0, 0x7a, 0x25,
	0x38, 0x0d, 0x1a, 0x50, 0xbf, 0xcc, 0x55, 0x4e, 0xd2, 0xf7, 0xa9, 0x0c, 0xbe, 0x39, 0xe0, 0x5e,
	0x14, 0x0a, 0x3d, 0x00, 0x8f, 0x49, 0x99, 0xd3, 0xcc, 0x77, 0xda, 0x4e, 0xa7, 0x8e, 0x97, 0x11,
	0x7a, 0x0c, 0x75, 0x92, 0x27, 0x8c, 0xf2, 0x98, 0x4a, 0xbf, 0xd2, 0x76, 0x3b, 0x75, 0x7c, 0x97,
	0x40, 0x8f, 0x60, 0x7b, 0x56, 0x7c, 0x96, 0x51, 0x9e, 0x31, 0xdf, 0xd5, 0x5c, 0xad, 0x8c, 0xc7,
	0x19, 0x43, 0x4f, 0xa0, 0x31, 0x2b, 0x54, 0x34, 0xa5, 0x24, 0xa1, 0x99, 0xf4, 0x3d, 0x8d, 0xc2,
	0xac, 0x50, 0x03, 0x93, 0x41, 0x7b, 0x50, 0x46, 0xd1, 0x82, 0x64, 0x64, 0x2e, 0xfd, 0x9a, 0x91,
	0x9e, 0x15, 0xea, 0x5a, 0x27, 0x82, 0x1f, 0x0e, 0xb4, 0x46, 0x22, 0xcf, 0x62, 0xfa, 0x76, 0xc5,
	0xf2, 0x25, 0x55, 0x53, 0x91, 0xa0, 0x63, 0xa8, 0x72, 0xc1, 0xa9, 0xfe, 0xda, 0x46, 0xef, 0x59,
	0xf8, 0xd7, 0xee, 0x84, 0xa5, 0xef, 0xc1, 0x16, 0xd6, 0x08, 0x7a, 0x03, 0xd5, 0xb9, 0x4a, 0x4b,
	0x37, 0x25, 0xda, 0xd9, 0x80, 0xde, 0xb6, 0xaa, 0xe4, 0x4b, 0x0e, 0x1d, 0x81, 0x3b, 0x2b, 0x94,
	0xf6, 0xdb, 0xe8, 0x05, 0x1b, 0xf0, 0x8b, 0x42, 0x0d, 0xb6, 0x70, 0x09, 0xf4, 0xb7, 0xc1, 0x33,
	0x66, 0x83, 0x4f, 0xd0, 0x1a, 0x66, 0x6c, 0xc2, 0xf8, 0x1f, 0xad, 0x21, 0xa8, 0x72, 0x32, 0xa7,
	0xcb, 0x41, 0xe8, 0xdf, 0xe8, 0x95, 0xa9, 0x59, 0xf9, 0xd7, 0x9a, 0xba, 0x62, 0xf0, 0xd3, 0x81,
	0x9d, 0x93, 0x8c, 0x26, 0xe5, 0x21, 0x92, 0xe2, 0x3c, 0xa5, 0x68, 0x08, 0xb5, 0x1b, 0xc6, 0x13,
	0xc6, 0x27, 0x5a, 0x7f, 0xa7, 0xf7, 0x7a, 0x83, 0xd8, 0x2a, 0x1f, 0xf6, 0x0d, 0x8c, 0xad, 0x0a,
	0x3a, 0x82, 0x87, 0x92, 0xa6, 0x34, 0x56, 0x34, 0x89, 0x84, 0x36, 0x15, 0xcd, 0xb5, 0x0f, 0xbb,
	0x2e, 0xf7, 0xed, 0xdf, 0xc6, 0xb2, 0x31, 0x29, 0xd1, 0x01, 0x34, 0xe7, 0x44, 0xc5, 0x53, 0xc6,
	0x27, 0x91, 0xd4, 0x73, 0x96, 0xbe, 0xab, 0x81, 0x5d, 0x9b, 0x37, 0xe3, 0x97, 0xc1, 0x01, 0xd4,
	0x96, 0x65, 0xd1, 0x0e, 0xc0, 0x78, 0x74, 0x1a, 0x8d, 0x86, 0x63, 0x7c, 0x72, 0xda, 0xdc, 0xb2,
	0xf1, 0x10, 0x9f, 0x9f, 0x9d, 0x5f, 0x35, 0x9d, 0xe0, 0x7b, 0x05, 0xbc, 0x6b, 0x7d, 0x6d, 0xd0,
	0x19, 0xdc, 0x4b, 0xa8, 0x54, 0x8c, 0x6b, 0x43, 0xd2, 0x77, 0xda, 0xee, 0x6f, 0x9b, 0xb2, 0xbc,
	0x2e, 0xd6, 0x67, 0x2f, 0x7c, 0x77, 0x77, 0x16, 0xaf, 0x80, 0x68, 0x04, 0x35, 0xfb, 0x81, 0x15,
	0xad, 0x71, 0xbc, 0xa1, 0x65, 0xeb, 0xd7, 0x16, 0x5b, 0xa5, 0x52, 0xd4, 0x74, 0xcb, 0xb8, 0xde,
	0x2c, 0xba, 0x7e, 0x61, 0xb0, 0x55, 0x42, 0x1f, 0xa0, 0x19, 0xdf, 0x8e, 0x4b, 0x5f, 0x7d, 0xe9,
	0x57, 0xb5, 0xfa, 0x8b, 0xff, 0x9a, 0x32, 0xde, 0x8d, 0x57, 0x62, 0xd9, 0x7f, 0xfe, 0x71, 0xdf,
	0x08, 0x30, 0xd1, 0x25, 0x0b, 0xd6, 0x5d, 0xf3, 0x64, 0xdd, 0x78, 0xfa, 0xa9, 0x79, 0xf9, 0x2b,
	0x00, 0x00, 0xff, 0xff, 0x2a, 0x4c, 0xd0, 0x18, 0xd4, 0x04, 0x00, 0x00,
>>>>>>> Change credential rules to use enum to define principal binding.
}
