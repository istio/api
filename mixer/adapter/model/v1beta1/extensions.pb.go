// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/model/v1beta1/extensions.proto

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	mixer/adapter/model/v1beta1/extensions.proto
	mixer/adapter/model/v1beta1/service.proto
	mixer/adapter/model/v1beta1/type.proto

It has these top-level messages:
	CreateSessionRequest
	CreateSessionResponse
	ValidateRequest
	ValidateResponse
	CloseSessionRequest
	CloseSessionResponse
	Value
	IPAddress
	Duration
	TimeStamp
	DNSName
	EmailAddress
	Uri
*/
package v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// The available varieties of templates, controlling the semantics of what an adapter does with each instance.
type TemplateVariety int32

const (
	// Makes the template applicable for Mixer's check calls.
	TEMPLATE_VARIETY_CHECK TemplateVariety = 0
	// Makes the template applicable for Mixer's report calls.
	TEMPLATE_VARIETY_REPORT TemplateVariety = 1
	// Makes the template applicable for Mixer's quota calls.
	TEMPLATE_VARIETY_QUOTA TemplateVariety = 2
	// Makes the template applicable for Mixer's quota calls.
	TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR TemplateVariety = 3
)

var TemplateVariety_name = map[int32]string{
	0: "TEMPLATE_VARIETY_CHECK",
	1: "TEMPLATE_VARIETY_REPORT",
	2: "TEMPLATE_VARIETY_QUOTA",
	3: "TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR",
}
var TemplateVariety_value = map[string]int32{
	"TEMPLATE_VARIETY_CHECK":               0,
	"TEMPLATE_VARIETY_REPORT":              1,
	"TEMPLATE_VARIETY_QUOTA":               2,
	"TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR": 3,
}

func (TemplateVariety) EnumDescriptor() ([]byte, []int) { return fileDescriptorExtensions, []int{0} }

var E_TemplateVariety = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FileOptions)(nil),
	ExtensionType: (*TemplateVariety)(nil),
	Field:         72295727,
	Name:          "istio.mixer.adapter.model.v1beta1.template_variety",
	Tag:           "varint,72295727,opt,name=template_variety,json=templateVariety,enum=istio.mixer.adapter.model.v1beta1.TemplateVariety",
	Filename:      "mixer/adapter/model/v1beta1/extensions.proto",
}

func init() {
	proto.RegisterEnum("istio.mixer.adapter.model.v1beta1.TemplateVariety", TemplateVariety_name, TemplateVariety_value)
	proto.RegisterExtension(E_TemplateVariety)
}
func (x TemplateVariety) String() string {
	s, ok := TemplateVariety_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

func init() {
	proto.RegisterFile("mixer/adapter/model/v1beta1/extensions.proto", fileDescriptorExtensions)
}

var fileDescriptorExtensions = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xc9, 0xcd, 0xac, 0x48,
	0x2d, 0xd2, 0x4f, 0x4c, 0x49, 0x2c, 0x28, 0x49, 0x2d, 0xd2, 0xcf, 0xcd, 0x4f, 0x49, 0xcd, 0xd1,
	0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc,
	0xcf, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0xcc, 0x2c, 0x2e, 0xc9, 0xcc, 0xd7,
	0x03, 0xeb, 0xd1, 0x83, 0xea, 0xd1, 0x03, 0xeb, 0xd1, 0x83, 0xea, 0x91, 0x52, 0x48, 0xcf, 0xcf,
	0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x6b, 0x48, 0x2a, 0x4d, 0xd3, 0x4f, 0x49, 0x2d, 0x4e, 0x2e, 0xca,
	0x2c, 0x28, 0xc9, 0x2f, 0x82, 0x18, 0xa2, 0x35, 0x81, 0x91, 0x8b, 0x3f, 0x24, 0x35, 0xb7, 0x20,
	0x27, 0xb1, 0x24, 0x35, 0x2c, 0xb1, 0x28, 0x33, 0xb5, 0xa4, 0x52, 0x48, 0x8a, 0x4b, 0x2c, 0xc4,
	0xd5, 0x37, 0xc0, 0xc7, 0x31, 0xc4, 0x35, 0x3e, 0xcc, 0x31, 0xc8, 0xd3, 0x35, 0x24, 0x32, 0xde,
	0xd9, 0xc3, 0xd5, 0xd9, 0x5b, 0x80, 0x41, 0x48, 0x9a, 0x4b, 0x1c, 0x43, 0x2e, 0xc8, 0x35, 0xc0,
	0x3f, 0x28, 0x44, 0x80, 0x11, 0xab, 0xc6, 0xc0, 0x50, 0xff, 0x10, 0x47, 0x01, 0x26, 0x21, 0x0d,
	0x2e, 0x15, 0x0c, 0x39, 0xc7, 0x90, 0x90, 0x20, 0x4f, 0xa7, 0xd0, 0x10, 0xd7, 0x78, 0x77, 0x57,
	0x3f, 0xd7, 0x20, 0xc7, 0x10, 0xff, 0x20, 0x01, 0x66, 0xab, 0x3a, 0x2e, 0x81, 0x12, 0xa8, 0x8b,
	0xe2, 0xcb, 0xa0, 0x4e, 0x92, 0xd1, 0x83, 0xf8, 0x44, 0x0f, 0xe6, 0x13, 0x3d, 0xb7, 0xcc, 0x9c,
	0x54, 0xff, 0x82, 0x12, 0x50, 0x78, 0x48, 0xac, 0x3f, 0xb5, 0x47, 0x49, 0x81, 0x51, 0x83, 0xcf,
	0xc8, 0x48, 0x8f, 0x60, 0x98, 0xe8, 0xa1, 0xf9, 0x36, 0x88, 0xbf, 0x04, 0x55, 0xc0, 0x29, 0xec,
	0xc2, 0x43, 0x39, 0x86, 0x1b, 0x0f, 0xe5, 0x18, 0x3e, 0x3c, 0x94, 0x63, 0x6c, 0x78, 0x24, 0xc7,
	0xb8, 0xe2, 0x91, 0x1c, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24,
	0xc7, 0xf8, 0xe2, 0x91, 0x1c, 0xc3, 0x87, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x44, 0x69,
	0x40, 0x6c, 0xcc, 0xcc, 0xd7, 0x4f, 0x2c, 0xc8, 0xd4, 0xc7, 0x13, 0x81, 0x49, 0x6c, 0x60, 0xb7,
	0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xaa, 0xe9, 0x12, 0xe6, 0x01, 0x00, 0x00,
}
