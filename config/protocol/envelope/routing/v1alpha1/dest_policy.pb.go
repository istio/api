// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/protocol/envelope/routing/v1alpha1/dest_policy.proto

package v1alpha1 // import "istio.io/api/config/protocol/envelope/routing/v1alpha1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envelope "istio.io/api/config/protocol/envelope"
import v1alpha1 "istio.io/api/routing/v1alpha1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DestinationPolicy struct {
	Metadata             *envelope.Metadata          `protobuf:"bytes,1,opt,name=Metadata" json:"Metadata,omitempty"`
	Contents             *v1alpha1.DestinationPolicy `protobuf:"bytes,2,opt,name=Contents" json:"Contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *DestinationPolicy) Reset()         { *m = DestinationPolicy{} }
func (m *DestinationPolicy) String() string { return proto.CompactTextString(m) }
func (*DestinationPolicy) ProtoMessage()    {}
func (*DestinationPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_dest_policy_2e8b24a0ad27ac42, []int{0}
}
func (m *DestinationPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestinationPolicy.Unmarshal(m, b)
}
func (m *DestinationPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestinationPolicy.Marshal(b, m, deterministic)
}
func (dst *DestinationPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestinationPolicy.Merge(dst, src)
}
func (m *DestinationPolicy) XXX_Size() int {
	return xxx_messageInfo_DestinationPolicy.Size(m)
}
func (m *DestinationPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_DestinationPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_DestinationPolicy proto.InternalMessageInfo

func (m *DestinationPolicy) GetMetadata() *envelope.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *DestinationPolicy) GetContents() *v1alpha1.DestinationPolicy {
	if m != nil {
		return m.Contents
	}
	return nil
}

func init() {
	proto.RegisterType((*DestinationPolicy)(nil), "istio.config.protocol.envelope.routing.v1alpha1.DestinationPolicy")
}

func init() {
	proto.RegisterFile("config/protocol/envelope/routing/v1alpha1/dest_policy.proto", fileDescriptor_dest_policy_2e8b24a0ad27ac42)
}

var fileDescriptor_dest_policy_2e8b24a0ad27ac42 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd0, 0xb1, 0x4b, 0xc5, 0x30,
	0x10, 0xc7, 0x71, 0xea, 0x20, 0x8f, 0x38, 0xd9, 0xe9, 0xd1, 0x49, 0xba, 0x58, 0x97, 0x3b, 0xaa,
	0x20, 0x82, 0x9b, 0xd6, 0x51, 0x90, 0x8e, 0x2e, 0x12, 0xdb, 0x58, 0x0f, 0x62, 0x2e, 0x34, 0x67,
	0xc1, 0xff, 0xc6, 0x3f, 0x55, 0x48, 0x9b, 0x0e, 0x8a, 0xf4, 0xed, 0xdf, 0xfb, 0xf0, 0x4b, 0xd4,
	0x6d, 0xc7, 0xee, 0x8d, 0x06, 0xf4, 0x23, 0x0b, 0x77, 0x6c, 0xd1, 0xb8, 0xc9, 0x58, 0xf6, 0x06,
	0x47, 0xfe, 0x14, 0x72, 0x03, 0x4e, 0xb5, 0xb6, 0xfe, 0x5d, 0xd7, 0xd8, 0x9b, 0x20, 0x2f, 0x9e,
	0x2d, 0x75, 0x5f, 0x10, 0xf3, 0x1c, 0x29, 0x08, 0x31, 0xcc, 0x04, 0x24, 0x02, 0x12, 0x01, 0x0b,
	0x01, 0x89, 0x28, 0xca, 0x6d, 0xb4, 0x38, 0xff, 0x77, 0xd1, 0x87, 0x11, 0xdd, 0x6b, 0xd1, 0x73,
	0x58, 0x7e, 0x67, 0xea, 0xb4, 0x31, 0x41, 0xc8, 0x69, 0x21, 0x76, 0x4f, 0x11, 0xc9, 0x1b, 0xb5,
	0x7b, 0x5c, 0xba, 0x7d, 0x76, 0x96, 0x55, 0x27, 0x97, 0x15, 0x6c, 0xcc, 0x4c, 0x7d, 0xbb, 0x5e,
	0xe6, 0x0f, 0x6a, 0x77, 0xcf, 0x4e, 0x8c, 0x93, 0xb0, 0x3f, 0x8a, 0xca, 0xc5, 0xa2, 0xfc, 0x7e,
	0x13, 0xfc, 0x99, 0xd0, 0xae, 0xa7, 0x77, 0x37, 0xcf, 0xd7, 0xf3, 0x15, 0x31, 0x6a, 0x4f, 0x78,
	0xf0, 0x67, 0xbf, 0x1e, 0xc7, 0xe6, 0xea, 0x27, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xcd, 0x41, 0x5c,
	0xa0, 0x01, 0x00, 0x00,
}
