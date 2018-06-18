// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/protocol/envelope/mesh/v1alpha1/config.proto

package v1alpha1 // import "istio.io/api/config/protocol/envelope/mesh/v1alpha1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envelope "istio.io/api/config/protocol/envelope"
import v1alpha1 "istio.io/api/mesh/v1alpha1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProxyConfig struct {
	Metadata             *envelope.Metadata    `protobuf:"bytes,1,opt,name=Metadata" json:"Metadata,omitempty"`
	Contents             *v1alpha1.ProxyConfig `protobuf:"bytes,2,opt,name=Contents" json:"Contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ProxyConfig) Reset()         { *m = ProxyConfig{} }
func (m *ProxyConfig) String() string { return proto.CompactTextString(m) }
func (*ProxyConfig) ProtoMessage()    {}
func (*ProxyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_bd511fa29958f997, []int{0}
}
func (m *ProxyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyConfig.Unmarshal(m, b)
}
func (m *ProxyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyConfig.Marshal(b, m, deterministic)
}
func (dst *ProxyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyConfig.Merge(dst, src)
}
func (m *ProxyConfig) XXX_Size() int {
	return xxx_messageInfo_ProxyConfig.Size(m)
}
func (m *ProxyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyConfig proto.InternalMessageInfo

func (m *ProxyConfig) GetMetadata() *envelope.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ProxyConfig) GetContents() *v1alpha1.ProxyConfig {
	if m != nil {
		return m.Contents
	}
	return nil
}

func init() {
	proto.RegisterType((*ProxyConfig)(nil), "istio.config.protocol.envelope.mesh.v1alpha1.ProxyConfig")
}

func init() {
	proto.RegisterFile("config/protocol/envelope/mesh/v1alpha1/config.proto", fileDescriptor_config_bd511fa29958f997)
}

var fileDescriptor_config_bd511fa29958f997 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0xd7, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x4f, 0xcd, 0x2b, 0x4b,
	0xcd, 0xc9, 0x2f, 0x48, 0xd5, 0xcf, 0x4d, 0x2d, 0xce, 0xd0, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8,
	0x48, 0x34, 0xd4, 0x87, 0x28, 0xd3, 0x03, 0x2b, 0x13, 0xd2, 0xc9, 0x2c, 0x2e, 0xc9, 0xcc, 0xd7,
	0x43, 0x16, 0x4b, 0xce, 0xcf, 0xd1, 0x83, 0x69, 0xd5, 0x03, 0x69, 0xd5, 0x83, 0x69, 0x95, 0x92,
	0xc2, 0x6d, 0x92, 0x94, 0x3a, 0x1e, 0xeb, 0x4b, 0x12, 0x53, 0x12, 0x4b, 0x12, 0x21, 0x0a, 0x95,
	0x26, 0x32, 0x72, 0x71, 0x07, 0x14, 0xe5, 0x57, 0x54, 0x3a, 0x83, 0x35, 0x08, 0xb9, 0x70, 0x71,
	0xf8, 0x42, 0x55, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x69, 0xe8, 0x11, 0x70, 0x15, 0x4c,
	0x7d, 0x10, 0x5c, 0xa7, 0x90, 0x0d, 0x17, 0x87, 0x73, 0x7e, 0x5e, 0x49, 0x6a, 0x5e, 0x49, 0xb1,
	0x04, 0x13, 0xd8, 0x14, 0x05, 0xa8, 0x29, 0x28, 0x5e, 0xd0, 0x43, 0xb2, 0x39, 0x08, 0xae, 0xc3,
	0xc9, 0x34, 0xca, 0x18, 0xa2, 0x38, 0x33, 0x5f, 0x3f, 0xb1, 0x20, 0x53, 0x9f, 0xb8, 0xa0, 0x4c,
	0x62, 0x03, 0x2b, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x47, 0xfa, 0x6f, 0x90, 0x7b, 0x01,
	0x00, 0x00,
}
