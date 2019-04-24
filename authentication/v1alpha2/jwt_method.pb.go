// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: authentication/v1alpha2/jwt_method.proto

package v1alpha2

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// $hide_from_docs
// JSON Web Token (JWT) token format for authentication as defined by
// [RFC 7519](`https://tools.ietf.org/html/rfc7519`). See [OAuth
// 2.0](`https://tools.ietf.org/html/rfc6749`) and [OIDC
// 1.0](`http://openid.net/connect`) for how this is used in the whole
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
type Jwt struct {
	// Identifies the issuer that issued the JWT. See
	// [issuer](`https://tools.ietf.org/html/rfc7519#section-4.1.1`)
	// Usually a URL or an email address.
	//
	// Example: `https://securetoken.google.com`
	// Example: `1234567-compute@developer.gserviceaccount.com`
	Issuer string `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// The list of JWT
	// [audiences](`https://tools.ietf.org/html/rfc7519#section-4.1.3`).
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
	// JWT. See [OpenID Discovery](`https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata`).
	//
	// Optional if the key set document can either (a) be retrieved from
	// [OpenID
	// Discovery](`https://openid.net/specs/openid-connect-discovery-1_0.html`) of
	// the issuer or (b) inferred from the email domain of the issuer (e.g. a
	// Google service account).
	//
	// Example: `https://www.googleapis.com/oauth2/v1/certs`
	//
	// Note: Only one of `jwks_uri` and `jwks` should be used.
	JwksUri string `protobuf:"bytes,3,opt,name=jwks_uri,json=jwksUri,proto3" json:"jwks_uri,omitempty"`
	// JSON Web Key Set of public keys to validate signature of the JWT.
	// See `https://auth0.com/docs/jwks`
	//
	// Note: Only one of `jwks_uri` and `jwks` should be used.
	Jwks string `protobuf:"bytes,10,opt,name=jwks,proto3" json:"jwks,omitempty"`
	// JWT is sent in a request header. `header` represents the
	// header name.
	//
	// For example, if `header=x-goog-iap-jwt-assertion`, the header
	// format will be `x-goog-iap-jwt-assertion`: $JWT.
	JwtHeaders []string `protobuf:"bytes,6,rep,name=jwt_headers,json=jwtHeaders,proto3" json:"jwt_headers,omitempty"`
	// JWT is sent in a query parameter. `query` represents the
	// query parameter name.
	//
	// For example, `query=jwt_token`.
	JwtParams            []string `protobuf:"bytes,7,rep,name=jwt_params,json=jwtParams,proto3" json:"jwt_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Jwt) Reset()         { *m = Jwt{} }
func (m *Jwt) String() string { return proto.CompactTextString(m) }
func (*Jwt) ProtoMessage()    {}
func (*Jwt) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed2e14219930498f, []int{0}
}
func (m *Jwt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Jwt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Jwt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Jwt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Jwt.Merge(m, src)
}
func (m *Jwt) XXX_Size() int {
	return m.Size()
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

func init() {
	proto.RegisterType((*Jwt)(nil), "istio.authentication.v1alpha2.Jwt")
}

func init() {
	proto.RegisterFile("authentication/v1alpha2/jwt_method.proto", fileDescriptor_ed2e14219930498f)
}

var fileDescriptor_ed2e14219930498f = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x89, 0x95, 0xae, 0x1d, 0x6f, 0x39, 0x48, 0x04, 0xb7, 0x2e, 0xe2, 0xa1, 0xa7, 0x16,
	0xd7, 0x37, 0xf0, 0x24, 0x9e, 0x64, 0xc1, 0x8b, 0x97, 0x32, 0x6e, 0x03, 0x9d, 0x6a, 0x9b, 0x92,
	0x4c, 0xcd, 0x3b, 0xf9, 0x24, 0x1e, 0x7d, 0x04, 0xe9, 0x93, 0x48, 0x52, 0x45, 0x3c, 0x78, 0xfb,
	0xff, 0x8f, 0x3f, 0x99, 0x99, 0x1f, 0x0a, 0x9c, 0xb8, 0xd5, 0x03, 0xd3, 0x1e, 0x99, 0xcc, 0x50,
	0xbd, 0x5e, 0xe1, 0xcb, 0xd8, 0xe2, 0xb6, 0xea, 0x3c, 0xd7, 0xbd, 0xe6, 0xd6, 0x34, 0xe5, 0x68,
	0x0d, 0x1b, 0xb9, 0x26, 0xc7, 0x64, 0xca, 0xbf, 0xf9, 0xf2, 0x27, 0x7f, 0xf1, 0x26, 0x20, 0xb9,
	0xf3, 0x2c, 0x4f, 0x20, 0x25, 0xe7, 0x26, 0x6d, 0x95, 0xd8, 0x88, 0x22, 0xdb, 0x7d, 0x3b, 0x79,
	0x06, 0x19, 0x4e, 0x0d, 0xe9, 0x61, 0xaf, 0x9d, 0x3a, 0xd8, 0x24, 0x45, 0xb6, 0xfb, 0x05, 0xf2,
	0x14, 0x8e, 0x3a, 0xff, 0xec, 0xea, 0xc9, 0x92, 0x4a, 0xe2, 0xbb, 0x55, 0xf0, 0x0f, 0x96, 0xa4,
	0x84, 0xc3, 0x20, 0x15, 0x44, 0x1c, 0xb5, 0x3c, 0x87, 0xe3, 0xb0, 0x5f, 0xab, 0xb1, 0xd1, 0xd6,
	0xa9, 0x34, 0x7e, 0x07, 0x9d, 0xe7, 0xdb, 0x85, 0xc8, 0x35, 0x04, 0x57, 0x8f, 0x68, 0xb1, 0x77,
	0x6a, 0xb5, 0x8c, 0xeb, 0x3c, 0xdf, 0x47, 0x70, 0xb3, 0x7d, 0x9f, 0x73, 0xf1, 0x31, 0xe7, 0xe2,
	0x73, 0xce, 0xc5, 0xe3, 0xe5, 0x72, 0x19, 0x99, 0x0a, 0x47, 0xaa, 0xfe, 0x29, 0xe4, 0x29, 0x8d,
	0x35, 0x5c, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x57, 0x1b, 0x66, 0x82, 0x32, 0x01, 0x00, 0x00,
}

func (m *Jwt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Jwt) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Issuer) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintJwtMethod(dAtA, i, uint64(len(m.Issuer)))
		i += copy(dAtA[i:], m.Issuer)
	}
	if len(m.Audiences) > 0 {
		for _, s := range m.Audiences {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.JwksUri) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintJwtMethod(dAtA, i, uint64(len(m.JwksUri)))
		i += copy(dAtA[i:], m.JwksUri)
	}
	if len(m.JwtHeaders) > 0 {
		for _, s := range m.JwtHeaders {
			dAtA[i] = 0x32
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.JwtParams) > 0 {
		for _, s := range m.JwtParams {
			dAtA[i] = 0x3a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Jwks) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintJwtMethod(dAtA, i, uint64(len(m.Jwks)))
		i += copy(dAtA[i:], m.Jwks)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintJwtMethod(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Jwt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovJwtMethod(uint64(l))
	}
	if len(m.Audiences) > 0 {
		for _, s := range m.Audiences {
			l = len(s)
			n += 1 + l + sovJwtMethod(uint64(l))
		}
	}
	l = len(m.JwksUri)
	if l > 0 {
		n += 1 + l + sovJwtMethod(uint64(l))
	}
	if len(m.JwtHeaders) > 0 {
		for _, s := range m.JwtHeaders {
			l = len(s)
			n += 1 + l + sovJwtMethod(uint64(l))
		}
	}
	if len(m.JwtParams) > 0 {
		for _, s := range m.JwtParams {
			l = len(s)
			n += 1 + l + sovJwtMethod(uint64(l))
		}
	}
	l = len(m.Jwks)
	if l > 0 {
		n += 1 + l + sovJwtMethod(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovJwtMethod(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozJwtMethod(x uint64) (n int) {
	return sovJwtMethod(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Jwt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJwtMethod
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
			return fmt.Errorf("proto: Jwt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Jwt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwtMethod
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
				return ErrInvalidLengthJwtMethod
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwtMethod
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
					return ErrIntOverflowJwtMethod
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
				return ErrInvalidLengthJwtMethod
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwtMethod
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
					return ErrIntOverflowJwtMethod
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
				return ErrInvalidLengthJwtMethod
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwtMethod
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JwksUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JwtHeaders", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwtMethod
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
				return ErrInvalidLengthJwtMethod
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwtMethod
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JwtHeaders = append(m.JwtHeaders, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JwtParams", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwtMethod
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
				return ErrInvalidLengthJwtMethod
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwtMethod
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JwtParams = append(m.JwtParams, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jwks", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwtMethod
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
				return ErrInvalidLengthJwtMethod
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwtMethod
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Jwks = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJwtMethod(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJwtMethod
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthJwtMethod
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
func skipJwtMethod(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowJwtMethod
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
					return 0, ErrIntOverflowJwtMethod
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
					return 0, ErrIntOverflowJwtMethod
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
				return 0, ErrInvalidLengthJwtMethod
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthJwtMethod
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowJwtMethod
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
				next, err := skipJwtMethod(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthJwtMethod
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
	ErrInvalidLengthJwtMethod = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowJwtMethod   = fmt.Errorf("proto: integer overflow")
)
