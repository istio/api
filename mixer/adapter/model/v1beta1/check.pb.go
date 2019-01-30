// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/model/v1beta1/check.proto

/*
	Package v1beta1 is a generated protocol buffer package.

	It is generated from these files:
		mixer/adapter/model/v1beta1/check.proto
		mixer/adapter/model/v1beta1/extensions.proto
		mixer/adapter/model/v1beta1/info.proto
		mixer/adapter/model/v1beta1/infrastructure_backend.proto
		mixer/adapter/model/v1beta1/quota.proto
		mixer/adapter/model/v1beta1/report.proto
		mixer/adapter/model/v1beta1/template.proto

	It has these top-level messages:
		CheckResult
		Info
		CreateSessionRequest
		CreateSessionResponse
		ValidateRequest
		ValidateResponse
		CloseSessionRequest
		CloseSessionResponse
		QuotaRequest
		QuotaResult
		ReportResult
		Template
*/
package v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"
import google_rpc "github.com/gogo/googleapis/google/rpc"

import time "time"

import types "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Expresses the result of a precondition check.
type CheckResult struct {
	// A status code of OK indicates preconditions were satisfied. Any other code indicates preconditions were not
	// satisfied and details describe why.
	Status google_rpc.Status `protobuf:"bytes,1,opt,name=status" json:"status"`
	// The amount of time for which this result can be considered valid.
	ValidDuration time.Duration `protobuf:"bytes,2,opt,name=valid_duration,json=validDuration,stdduration" json:"valid_duration"`
	// The number of uses for which this result can be considered valid.
	ValidUseCount int32 `protobuf:"varint,3,opt,name=valid_use_count,json=validUseCount,proto3" json:"valid_use_count,omitempty"`
}

func (m *CheckResult) Reset()                    { *m = CheckResult{} }
func (*CheckResult) ProtoMessage()               {}
func (*CheckResult) Descriptor() ([]byte, []int) { return fileDescriptorCheck, []int{0} }

func init() {
	proto.RegisterType((*CheckResult)(nil), "istio.mixer.adapter.model.v1beta1.CheckResult")
}
func (m *CheckResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckResult) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintCheck(dAtA, i, uint64(m.Status.Size()))
	n1, err := m.Status.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintCheck(dAtA, i, uint64(types.SizeOfStdDuration(m.ValidDuration)))
	n2, err := types.StdDurationMarshalTo(m.ValidDuration, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.ValidUseCount != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintCheck(dAtA, i, uint64(m.ValidUseCount))
	}
	return i, nil
}

func encodeVarintCheck(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CheckResult) Size() (n int) {
	var l int
	_ = l
	l = m.Status.Size()
	n += 1 + l + sovCheck(uint64(l))
	l = types.SizeOfStdDuration(m.ValidDuration)
	n += 1 + l + sovCheck(uint64(l))
	if m.ValidUseCount != 0 {
		n += 1 + sovCheck(uint64(m.ValidUseCount))
	}
	return n
}

func sovCheck(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCheck(x uint64) (n int) {
	return sovCheck(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CheckResult) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CheckResult{`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "Status", "google_rpc.Status", 1), `&`, ``, 1) + `,`,
		`ValidDuration:` + strings.Replace(strings.Replace(this.ValidDuration.String(), "Duration", "google_protobuf1.Duration", 1), `&`, ``, 1) + `,`,
		`ValidUseCount:` + fmt.Sprintf("%v", this.ValidUseCount) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCheck(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CheckResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCheck
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CheckResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCheck
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCheck
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdDurationUnmarshal(&m.ValidDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidUseCount", wireType)
			}
			m.ValidUseCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidUseCount |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCheck(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCheck
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCheck(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCheck
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
					return 0, ErrIntOverflowCheck
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
					return 0, ErrIntOverflowCheck
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCheck
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCheck
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
				next, err := skipCheck(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthCheck = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCheck   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("mixer/adapter/model/v1beta1/check.proto", fileDescriptorCheck) }

var fileDescriptorCheck = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x31, 0x4e, 0xf3, 0x30,
	0x18, 0x86, 0xed, 0xff, 0x87, 0x0a, 0xa5, 0x02, 0xa4, 0x2c, 0x94, 0x0e, 0x5f, 0x0b, 0x03, 0x74,
	0xc1, 0xa6, 0x70, 0x83, 0x96, 0x89, 0xb1, 0x88, 0xa5, 0x4b, 0xe5, 0x24, 0x26, 0xb5, 0x48, 0xeb,
	0x28, 0xb1, 0x2b, 0x46, 0x8e, 0xc0, 0xc8, 0x11, 0xd8, 0xb9, 0x44, 0xc7, 0x8e, 0x4c, 0x40, 0xcc,
	0xc2, 0xd8, 0x23, 0xa0, 0xda, 0x8e, 0xd8, 0xd8, 0xf2, 0xe5, 0x79, 0xde, 0xcf, 0xaf, 0xec, 0xe0,
	0x74, 0x26, 0x1e, 0x78, 0x41, 0x59, 0xc2, 0x72, 0xc5, 0x0b, 0x3a, 0x93, 0x09, 0xcf, 0xe8, 0xa2,
	0x1f, 0x71, 0xc5, 0xfa, 0x34, 0x9e, 0xf2, 0xf8, 0x9e, 0xe4, 0x85, 0x54, 0x32, 0x3c, 0x12, 0xa5,
	0x12, 0x92, 0x58, 0x9d, 0x78, 0x9d, 0x58, 0x9d, 0x78, 0xbd, 0x7d, 0x96, 0x0a, 0x35, 0xd5, 0x11,
	0x89, 0xe5, 0x8c, 0xa6, 0x32, 0x95, 0xd4, 0x26, 0x23, 0x7d, 0x67, 0x27, 0x3b, 0xd8, 0x2f, 0xb7,
	0xb1, 0x0d, 0xa9, 0x94, 0x69, 0xc6, 0x7f, 0xad, 0x44, 0x17, 0x4c, 0x09, 0x39, 0xf7, 0xfc, 0xc0,
	0xf3, 0x22, 0x8f, 0x69, 0xa9, 0x98, 0xd2, 0xa5, 0x03, 0xc7, 0xaf, 0x38, 0x68, 0x0e, 0x37, 0xd5,
	0x46, 0xbc, 0xd4, 0x99, 0x0a, 0xcf, 0x83, 0x86, 0xe3, 0x2d, 0xdc, 0xc5, 0xbd, 0xe6, 0x45, 0x48,
	0x5c, 0x92, 0x14, 0x79, 0x4c, 0x6e, 0x2c, 0x19, 0x6c, 0x2d, 0xdf, 0x3b, 0x68, 0xe4, 0xbd, 0xf0,
	0x3a, 0xd8, 0x5b, 0xb0, 0x4c, 0x24, 0x93, 0xfa, 0xc8, 0xd6, 0x3f, 0x9b, 0x3c, 0xac, 0x93, 0x75,
	0x27, 0x72, 0xe5, 0x85, 0xc1, 0xce, 0x66, 0xc1, 0xf3, 0x47, 0x07, 0x8f, 0x76, 0x6d, 0xb4, 0x06,
	0xe1, 0x49, 0xb0, 0xef, 0x76, 0xe9, 0x92, 0x4f, 0x62, 0xa9, 0xe7, 0xaa, 0xf5, 0xbf, 0x8b, 0x7b,
	0xdb, 0xde, 0xbb, 0x2d, 0xf9, 0x70, 0xf3, 0x73, 0x30, 0x5e, 0x56, 0x80, 0x56, 0x15, 0xa0, 0xb7,
	0x0a, 0xd0, 0xba, 0x02, 0xf4, 0x68, 0x00, 0xbf, 0x18, 0x40, 0x4b, 0x03, 0x78, 0x65, 0x00, 0x7f,
	0x1a, 0xc0, 0xdf, 0x06, 0xd0, 0xda, 0x00, 0x7e, 0xfa, 0x02, 0x34, 0xee, 0xb9, 0x2b, 0x17, 0x92,
	0xb2, 0x5c, 0xd0, 0x3f, 0x1e, 0x2a, 0x6a, 0xd8, 0xbe, 0x97, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x80, 0xaa, 0x73, 0x47, 0xce, 0x01, 0x00, 0x00,
}
