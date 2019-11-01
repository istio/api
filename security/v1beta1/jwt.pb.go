// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security/v1beta1/jwt.proto

package v1beta1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// JSON Web Token (JWT) token format for authentication as defined by
// [RFC 7519](https://tools.ietf.org/html/rfc7519). See [OAuth 2.0](https://tools.ietf.org/html/rfc6749) and
// [OIDC 1.0](http://openid.net/connect) for how this is used in the whole
// authentication flow.
//
// Examples:
//
// Spec for a JWT that is issued by `https://example.com`, with the audience claims must be either
// `bookstore_android.apps.googleusercontent.com` or `bookstore_web.apps.googleusercontent.com`.
// The token should be presented at the `Authorization` header (default). The Json web key set (JWKS)
// will be discovered followwing OpenID Connect protocol.
//
// ```yaml
// issuer: https://example.com
// audiences:
// - bookstore_android.apps.googleusercontent.com
//   bookstore_web.apps.googleusercontent.com
// ```
//
// This example specifies token in non-default location (`x-goog-iap-jwt-assertion` header). It also
// defines the URI to fetch JWKS explicitly.
//
// ```yaml
// issuer: https://example.com
// jwksUri: https://example.com/.secret/jwks.json
// jwtHeaders:
// - "x-goog-iap-jwt-assertion"
// ```
type JWT struct {
	// Identifies the issuer that issued the JWT. See
	// [issuer](https://tools.ietf.org/html/rfc7519#section-4.1.1)
	// Usually a URL or an email address.
	//
	// Example: https://foobar.auth0.com
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
	// Note: Only one of jwks_uri and jwks should be used. jwks_uri will be ignored if it does.
	JwksUri string `protobuf:"bytes,3,opt,name=jwks_uri,json=jwksUri,proto3" json:"jwks_uri,omitempty"`
	// JSON Web Key Set of public keys to validate signature of the JWT.
	// See https://auth0.com/docs/jwks.
	//
	// Note: Only one of jwks_uri and jwks should be used. jwks_uri will be ignored if it does.
	Jwks string `protobuf:"bytes,10,opt,name=jwks,proto3" json:"jwks,omitempty"`
	// List of header locations from which JWT is expected. For example, below is the location spec if
	// if JWT is expected to be found in `x-goog-iap-jwt-assertion` header, and have "Bearer " prefix:
	// ```
	//   fromHeaders:
	//   - name: x-goog-iap-jwt-assertion
	//     prefix: "Bearer "
	// ```
	FromHeaders []*JWTHeader `protobuf:"bytes,6,rep,name=from_headers,json=fromHeaders,proto3" json:"from_headers,omitempty"`
	// List of query parameters from which JWT is expected. For example, if JWT is provided via query
	// parameter `my_token` (e.g /path?my_token=<JWT>), the config is:
	// ```
	//   fromParams:
	//   - "my_token"
	// ```
	FromParams           []string `protobuf:"bytes,7,rep,name=from_params,json=fromParams,proto3" json:"from_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JWT) Reset()         { *m = JWT{} }
func (m *JWT) String() string { return proto.CompactTextString(m) }
func (*JWT) ProtoMessage()    {}
func (*JWT) Descriptor() ([]byte, []int) {
	return fileDescriptor_163ab6fd32fb6b15, []int{0}
}
func (m *JWT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JWT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JWT.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JWT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWT.Merge(m, src)
}
func (m *JWT) XXX_Size() int {
	return m.Size()
}
func (m *JWT) XXX_DiscardUnknown() {
	xxx_messageInfo_JWT.DiscardUnknown(m)
}

var xxx_messageInfo_JWT proto.InternalMessageInfo

func (m *JWT) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *JWT) GetAudiences() []string {
	if m != nil {
		return m.Audiences
	}
	return nil
}

func (m *JWT) GetJwksUri() string {
	if m != nil {
		return m.JwksUri
	}
	return ""
}

func (m *JWT) GetJwks() string {
	if m != nil {
		return m.Jwks
	}
	return ""
}

func (m *JWT) GetFromHeaders() []*JWTHeader {
	if m != nil {
		return m.FromHeaders
	}
	return nil
}

func (m *JWT) GetFromParams() []string {
	if m != nil {
		return m.FromParams
	}
	return nil
}

// This message specifies a header location to extract JWT token.
type JWTHeader struct {
	// The HTTP header name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The prefix that should be stripped before decoding the token.
	// For example, for "Authorization: Bearer <token>", prefix="Bearer " with a space at the end.
	// If the header doesn't have this exact prefix, it is considerred invalid.
	Prefix               string   `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JWTHeader) Reset()         { *m = JWTHeader{} }
func (m *JWTHeader) String() string { return proto.CompactTextString(m) }
func (*JWTHeader) ProtoMessage()    {}
func (*JWTHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_163ab6fd32fb6b15, []int{1}
}
func (m *JWTHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JWTHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JWTHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JWTHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWTHeader.Merge(m, src)
}
func (m *JWTHeader) XXX_Size() int {
	return m.Size()
}
func (m *JWTHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_JWTHeader.DiscardUnknown(m)
}

var xxx_messageInfo_JWTHeader proto.InternalMessageInfo

func (m *JWTHeader) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JWTHeader) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func init() {
	proto.RegisterType((*JWT)(nil), "istio.security.v1beta1.JWT")
	proto.RegisterType((*JWTHeader)(nil), "istio.security.v1beta1.JWTHeader")
}

func init() { proto.RegisterFile("security/v1beta1/jwt.proto", fileDescriptor_163ab6fd32fb6b15) }

var fileDescriptor_163ab6fd32fb6b15 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0x33, 0x94, 0x94, 0xaf, 0x97, 0x6f, 0x35, 0x0b, 0x1c, 0x51, 0x01, 0x59, 0x91, 0x98,
	0xb4, 0x41, 0xb7, 0x6e, 0x24, 0x2e, 0x0c, 0x2b, 0xd3, 0x60, 0x48, 0xdc, 0x34, 0x03, 0x5c, 0xe0,
	0x22, 0x30, 0xcd, 0x4c, 0x0b, 0xfa, 0x86, 0x2e, 0x5d, 0xb9, 0x26, 0x7d, 0x12, 0xd3, 0x69, 0xfd,
	0x13, 0xe3, 0xee, 0x9c, 0x5f, 0xce, 0x49, 0xce, 0xdc, 0x81, 0xa6, 0xc1, 0x69, 0xaa, 0x29, 0x79,
	0x09, 0x76, 0xfd, 0x09, 0x26, 0xb2, 0x1f, 0xac, 0xf6, 0x89, 0x1f, 0x6b, 0x95, 0x28, 0xde, 0x20,
	0x93, 0x90, 0xf2, 0x3f, 0x13, 0x7e, 0x99, 0x68, 0xb6, 0x17, 0x4a, 0x2d, 0xd6, 0x18, 0xc8, 0x98,
	0x82, 0x39, 0xe1, 0x7a, 0x16, 0x4d, 0x70, 0x29, 0x77, 0xa4, 0x74, 0x51, 0xec, 0xbe, 0x33, 0x70,
	0x86, 0xe3, 0x11, 0x3f, 0x01, 0x97, 0x8c, 0x49, 0x51, 0x0b, 0xd6, 0x61, 0x3d, 0x6f, 0xe0, 0x1c,
	0x6e, 0x2a, 0x61, 0x89, 0xf8, 0x29, 0x78, 0x32, 0x9d, 0x11, 0x6e, 0xa7, 0x68, 0x44, 0xa5, 0xe3,
	0xf4, 0xbc, 0xf0, 0x1b, 0xf0, 0x63, 0xf8, 0xb7, 0xda, 0x3f, 0x99, 0x28, 0xd5, 0x24, 0x9c, 0xbc,
	0x1c, 0xd6, 0x72, 0xff, 0xa0, 0x89, 0x73, 0xa8, 0xe6, 0x52, 0x80, 0xc5, 0x56, 0xf3, 0x5b, 0xf8,
	0x3f, 0xd7, 0x6a, 0x13, 0x2d, 0x51, 0xce, 0x50, 0x1b, 0xe1, 0x76, 0x9c, 0x5e, 0xfd, 0xf2, 0xdc,
	0xff, 0xfb, 0x05, 0xfe, 0x70, 0x3c, 0xba, 0xb3, 0xc9, 0xb0, 0x9e, 0xd7, 0x0a, 0x6d, 0x78, 0x1b,
	0xac, 0x8d, 0x62, 0xa9, 0xe5, 0xc6, 0x88, 0x9a, 0x1d, 0x05, 0x39, 0xba, 0xb7, 0xa4, 0x7b, 0x0d,
	0xde, 0x57, 0x95, 0x1f, 0x41, 0x75, 0x2b, 0x37, 0xf8, 0xf3, 0x6d, 0x16, 0xf0, 0x06, 0xb8, 0xb1,
	0xc6, 0x39, 0x3d, 0x8b, 0x8a, 0x9d, 0x58, 0xba, 0xc1, 0xc5, 0x6b, 0xd6, 0x62, 0x6f, 0x59, 0x8b,
	0x1d, 0xb2, 0x16, 0x7b, 0x3c, 0x2b, 0xb6, 0x91, 0xb2, 0x77, 0xfc, 0xfd, 0x0d, 0x13, 0xd7, 0x9e,
	0xf2, 0xea, 0x23, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x0d, 0x0c, 0x86, 0xa1, 0x01, 0x00, 0x00,
}

func (m *JWT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JWT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JWT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Jwks) > 0 {
		i -= len(m.Jwks)
		copy(dAtA[i:], m.Jwks)
		i = encodeVarintJwt(dAtA, i, uint64(len(m.Jwks)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.FromParams) > 0 {
		for iNdEx := len(m.FromParams) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.FromParams[iNdEx])
			copy(dAtA[i:], m.FromParams[iNdEx])
			i = encodeVarintJwt(dAtA, i, uint64(len(m.FromParams[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.FromHeaders) > 0 {
		for iNdEx := len(m.FromHeaders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FromHeaders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintJwt(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.JwksUri) > 0 {
		i -= len(m.JwksUri)
		copy(dAtA[i:], m.JwksUri)
		i = encodeVarintJwt(dAtA, i, uint64(len(m.JwksUri)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Audiences) > 0 {
		for iNdEx := len(m.Audiences) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Audiences[iNdEx])
			copy(dAtA[i:], m.Audiences[iNdEx])
			i = encodeVarintJwt(dAtA, i, uint64(len(m.Audiences[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintJwt(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *JWTHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JWTHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JWTHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Prefix) > 0 {
		i -= len(m.Prefix)
		copy(dAtA[i:], m.Prefix)
		i = encodeVarintJwt(dAtA, i, uint64(len(m.Prefix)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintJwt(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintJwt(dAtA []byte, offset int, v uint64) int {
	offset -= sovJwt(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *JWT) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovJwt(uint64(l))
	}
	if len(m.Audiences) > 0 {
		for _, s := range m.Audiences {
			l = len(s)
			n += 1 + l + sovJwt(uint64(l))
		}
	}
	l = len(m.JwksUri)
	if l > 0 {
		n += 1 + l + sovJwt(uint64(l))
	}
	if len(m.FromHeaders) > 0 {
		for _, e := range m.FromHeaders {
			l = e.Size()
			n += 1 + l + sovJwt(uint64(l))
		}
	}
	if len(m.FromParams) > 0 {
		for _, s := range m.FromParams {
			l = len(s)
			n += 1 + l + sovJwt(uint64(l))
		}
	}
	l = len(m.Jwks)
	if l > 0 {
		n += 1 + l + sovJwt(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *JWTHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovJwt(uint64(l))
	}
	l = len(m.Prefix)
	if l > 0 {
		n += 1 + l + sovJwt(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovJwt(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozJwt(x uint64) (n int) {
	return sovJwt(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *JWT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJwt
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: JWT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JWT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Audiences", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Audiences = append(m.Audiences, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JwksUri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JwksUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromHeaders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthJwt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromHeaders = append(m.FromHeaders, &JWTHeader{})
			if err := m.FromHeaders[len(m.FromHeaders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromParams", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromParams = append(m.FromParams, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jwks", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Jwks = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJwt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJwt
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthJwt
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *JWTHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJwt
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: JWTHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JWTHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJwt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJwt
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthJwt
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipJwt(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowJwt
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthJwt
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthJwt
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowJwt
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipJwt(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthJwt
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthJwt = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowJwt   = fmt.Errorf("proto: integer overflow")
)
