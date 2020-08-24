// Code generated by protoc-gen-go. DO NOT EDIT.
// source: analysis/v1alpha1/message.proto

// Describes the structure of messages generated by Istio analyzers.

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "google/protobuf/struct.proto"
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

// The values here are chosen so that more severe messages get sorted higher,
// as well as leaving space in between to add more later
type AnalysisMessageBase_Level int32

const (
	AnalysisMessageBase_UNKNOWN AnalysisMessageBase_Level = 0
	AnalysisMessageBase_ERROR   AnalysisMessageBase_Level = 3
	AnalysisMessageBase_WARNING AnalysisMessageBase_Level = 8
	AnalysisMessageBase_INFO    AnalysisMessageBase_Level = 12
)

var AnalysisMessageBase_Level_name = map[int32]string{
	0:  "UNKNOWN",
	3:  "ERROR",
	8:  "WARNING",
	12: "INFO",
}

var AnalysisMessageBase_Level_value = map[string]int32{
	"UNKNOWN": 0,
	"ERROR":   3,
	"WARNING": 8,
	"INFO":    12,
}

func (x AnalysisMessageBase_Level) String() string {
	return proto.EnumName(AnalysisMessageBase_Level_name, int32(x))
}

func (AnalysisMessageBase_Level) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4d6fd1414dfe3800, []int{0, 0}
}

// AnalysisMessageBase describes some common information that is needed for all
// messages. All information should be static with respect to the error code.
type AnalysisMessageBase struct {
	Type *AnalysisMessageBase_Type `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Represents how severe a message is. Required.
	Level AnalysisMessageBase_Level `protobuf:"varint,2,opt,name=level,proto3,enum=istio.analysis.v1alpha1.AnalysisMessageBase_Level" json:"level,omitempty"`
	// A url pointing to the Istio documentation for this specific error type.
	// Should be of the form
	// `^http(s)?://(preliminary\.)?istio.io/docs/reference/config/analysis/`
	// Required.
	DocumentationUrl     string   `protobuf:"bytes,3,opt,name=documentation_url,json=documentationUrl,proto3" json:"documentation_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnalysisMessageBase) Reset()         { *m = AnalysisMessageBase{} }
func (m *AnalysisMessageBase) String() string { return proto.CompactTextString(m) }
func (*AnalysisMessageBase) ProtoMessage()    {}
func (*AnalysisMessageBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d6fd1414dfe3800, []int{0}
}

func (m *AnalysisMessageBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalysisMessageBase.Unmarshal(m, b)
}
func (m *AnalysisMessageBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalysisMessageBase.Marshal(b, m, deterministic)
}
func (m *AnalysisMessageBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalysisMessageBase.Merge(m, src)
}
func (m *AnalysisMessageBase) XXX_Size() int {
	return xxx_messageInfo_AnalysisMessageBase.Size(m)
}
func (m *AnalysisMessageBase) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalysisMessageBase.DiscardUnknown(m)
}

var xxx_messageInfo_AnalysisMessageBase proto.InternalMessageInfo

func (m *AnalysisMessageBase) GetType() *AnalysisMessageBase_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *AnalysisMessageBase) GetLevel() AnalysisMessageBase_Level {
	if m != nil {
		return m.Level
	}
	return AnalysisMessageBase_UNKNOWN
}

func (m *AnalysisMessageBase) GetDocumentationUrl() string {
	if m != nil {
		return m.DocumentationUrl
	}
	return ""
}

// A unique identifier for the type of message. Name is intended to be
// human-readable, code is intended to be machine readable. There should be a
// one-to-one mapping between name and code. (i.e. do not re-use names or
// codes between message types.)
type AnalysisMessageBase_Type struct {
	// A human-readable name for the message type. e.g. "InternalError",
	// "PodMissingProxy". This should be the same for all messages of the same type.
	// Required.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A 7 character code matching `^IST[0-9]{4}$` intended to uniquely identify
	// the message type. (e.g. "IST0001" is mapped to the "InternalError" message
	// type.) 0000-0100 are reserved. Required.
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnalysisMessageBase_Type) Reset()         { *m = AnalysisMessageBase_Type{} }
func (m *AnalysisMessageBase_Type) String() string { return proto.CompactTextString(m) }
func (*AnalysisMessageBase_Type) ProtoMessage()    {}
func (*AnalysisMessageBase_Type) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d6fd1414dfe3800, []int{0, 0}
}

func (m *AnalysisMessageBase_Type) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalysisMessageBase_Type.Unmarshal(m, b)
}
func (m *AnalysisMessageBase_Type) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalysisMessageBase_Type.Marshal(b, m, deterministic)
}
func (m *AnalysisMessageBase_Type) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalysisMessageBase_Type.Merge(m, src)
}
func (m *AnalysisMessageBase_Type) XXX_Size() int {
	return xxx_messageInfo_AnalysisMessageBase_Type.Size(m)
}
func (m *AnalysisMessageBase_Type) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalysisMessageBase_Type.DiscardUnknown(m)
}

var xxx_messageInfo_AnalysisMessageBase_Type proto.InternalMessageInfo

func (m *AnalysisMessageBase_Type) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AnalysisMessageBase_Type) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

// AnalysisMessageWeakSchema is the set of information that's needed to define a
// weakly-typed schema. The purpose of this proto is to provide a mechanism for
// validating istio/istio/galley/pkg/config/analysis/msg/messages.yaml to make
// sure that we don't allow committing underspecified types.
type AnalysisMessageWeakSchema struct {
	// Required
	MessageBase *AnalysisMessageBase `protobuf:"bytes,1,opt,name=message_base,json=messageBase,proto3" json:"message_base,omitempty"`
	// A human readable description of what the error means. Required.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// A go-style template string (https://golang.org/pkg/fmt/#hdr-Printing)
	// defining how to combine the args for a  particular message into a log line.
	// Required.
	Template string `protobuf:"bytes,3,opt,name=template,proto3" json:"template,omitempty"`
	// A description of the arguments for a particular message type
	Args                 []*AnalysisMessageWeakSchema_ArgType `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *AnalysisMessageWeakSchema) Reset()         { *m = AnalysisMessageWeakSchema{} }
func (m *AnalysisMessageWeakSchema) String() string { return proto.CompactTextString(m) }
func (*AnalysisMessageWeakSchema) ProtoMessage()    {}
func (*AnalysisMessageWeakSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d6fd1414dfe3800, []int{1}
}

func (m *AnalysisMessageWeakSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalysisMessageWeakSchema.Unmarshal(m, b)
}
func (m *AnalysisMessageWeakSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalysisMessageWeakSchema.Marshal(b, m, deterministic)
}
func (m *AnalysisMessageWeakSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalysisMessageWeakSchema.Merge(m, src)
}
func (m *AnalysisMessageWeakSchema) XXX_Size() int {
	return xxx_messageInfo_AnalysisMessageWeakSchema.Size(m)
}
func (m *AnalysisMessageWeakSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalysisMessageWeakSchema.DiscardUnknown(m)
}

var xxx_messageInfo_AnalysisMessageWeakSchema proto.InternalMessageInfo

func (m *AnalysisMessageWeakSchema) GetMessageBase() *AnalysisMessageBase {
	if m != nil {
		return m.MessageBase
	}
	return nil
}

func (m *AnalysisMessageWeakSchema) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AnalysisMessageWeakSchema) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *AnalysisMessageWeakSchema) GetArgs() []*AnalysisMessageWeakSchema_ArgType {
	if m != nil {
		return m.Args
	}
	return nil
}

type AnalysisMessageWeakSchema_ArgType struct {
	// Required
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. Should be a golang type, used in code generation.
	// Ideally this will change to a less language-pinned type before this gets
	// out of alpha, but for compatibility with current istio/istio code it's
	// go_type for now.
	GoType               string   `protobuf:"bytes,2,opt,name=go_type,json=goType,proto3" json:"go_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnalysisMessageWeakSchema_ArgType) Reset()         { *m = AnalysisMessageWeakSchema_ArgType{} }
func (m *AnalysisMessageWeakSchema_ArgType) String() string { return proto.CompactTextString(m) }
func (*AnalysisMessageWeakSchema_ArgType) ProtoMessage()    {}
func (*AnalysisMessageWeakSchema_ArgType) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d6fd1414dfe3800, []int{1, 0}
}

func (m *AnalysisMessageWeakSchema_ArgType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalysisMessageWeakSchema_ArgType.Unmarshal(m, b)
}
func (m *AnalysisMessageWeakSchema_ArgType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalysisMessageWeakSchema_ArgType.Marshal(b, m, deterministic)
}
func (m *AnalysisMessageWeakSchema_ArgType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalysisMessageWeakSchema_ArgType.Merge(m, src)
}
func (m *AnalysisMessageWeakSchema_ArgType) XXX_Size() int {
	return xxx_messageInfo_AnalysisMessageWeakSchema_ArgType.Size(m)
}
func (m *AnalysisMessageWeakSchema_ArgType) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalysisMessageWeakSchema_ArgType.DiscardUnknown(m)
}

var xxx_messageInfo_AnalysisMessageWeakSchema_ArgType proto.InternalMessageInfo

func (m *AnalysisMessageWeakSchema_ArgType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AnalysisMessageWeakSchema_ArgType) GetGoType() string {
	if m != nil {
		return m.GoType
	}
	return ""
}

// GenericAnalysisMessage is an instance of an AnalysisMessage defined by a
// schema, whose metaschema is AnalysisMessageWeakSchema. (Names are hard.) Code
// should be able to perform validation of arguments as needed by using the
// message type information to look at the AnalysisMessageWeakSchema and examine the
// list of args at runtime. Developers can also create stronger-typed versions
// of GenericAnalysisMessage for well-known and stable message types.
type GenericAnalysisMessage struct {
	// Required
	MessageBase *AnalysisMessageBase `protobuf:"bytes,1,opt,name=message_base,json=messageBase,proto3" json:"message_base,omitempty"`
	// Any message-type specific arguments that need to get codified. Optional.
	Args *_struct.Struct `protobuf:"bytes,2,opt,name=args,proto3" json:"args,omitempty"`
	// A list of strings specifying the resource identifiers that were the cause
	// of message generation. A "path" here is a (NAMESPACE\/)?RESOURCETYPE/NAME
	// tuple that uniquely identifies a particular resource. There doesn't seem to
	// be a single concept for this, but this is intuitively taken from
	// https://kubernetes.io/docs/reference/using-api/api-concepts/#standard-api-terminology
	// At least one is required.
	ResourcePaths        []string `protobuf:"bytes,3,rep,name=resource_paths,json=resourcePaths,proto3" json:"resource_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenericAnalysisMessage) Reset()         { *m = GenericAnalysisMessage{} }
func (m *GenericAnalysisMessage) String() string { return proto.CompactTextString(m) }
func (*GenericAnalysisMessage) ProtoMessage()    {}
func (*GenericAnalysisMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d6fd1414dfe3800, []int{2}
}

func (m *GenericAnalysisMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericAnalysisMessage.Unmarshal(m, b)
}
func (m *GenericAnalysisMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericAnalysisMessage.Marshal(b, m, deterministic)
}
func (m *GenericAnalysisMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericAnalysisMessage.Merge(m, src)
}
func (m *GenericAnalysisMessage) XXX_Size() int {
	return xxx_messageInfo_GenericAnalysisMessage.Size(m)
}
func (m *GenericAnalysisMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericAnalysisMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GenericAnalysisMessage proto.InternalMessageInfo

func (m *GenericAnalysisMessage) GetMessageBase() *AnalysisMessageBase {
	if m != nil {
		return m.MessageBase
	}
	return nil
}

func (m *GenericAnalysisMessage) GetArgs() *_struct.Struct {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *GenericAnalysisMessage) GetResourcePaths() []string {
	if m != nil {
		return m.ResourcePaths
	}
	return nil
}

// InternalErrorAnalysisMessage is a strongly-typed message representing some
// error in Istio code that prevented us from performing analysis at all.
type InternalErrorAnalysisMessage struct {
	// Required
	MessageBase *AnalysisMessageBase `protobuf:"bytes,1,opt,name=message_base,json=messageBase,proto3" json:"message_base,omitempty"`
	// Any detail regarding specifics of the error. Should be human-readable.
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalErrorAnalysisMessage) Reset()         { *m = InternalErrorAnalysisMessage{} }
func (m *InternalErrorAnalysisMessage) String() string { return proto.CompactTextString(m) }
func (*InternalErrorAnalysisMessage) ProtoMessage()    {}
func (*InternalErrorAnalysisMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d6fd1414dfe3800, []int{3}
}

func (m *InternalErrorAnalysisMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalErrorAnalysisMessage.Unmarshal(m, b)
}
func (m *InternalErrorAnalysisMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalErrorAnalysisMessage.Marshal(b, m, deterministic)
}
func (m *InternalErrorAnalysisMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalErrorAnalysisMessage.Merge(m, src)
}
func (m *InternalErrorAnalysisMessage) XXX_Size() int {
	return xxx_messageInfo_InternalErrorAnalysisMessage.Size(m)
}
func (m *InternalErrorAnalysisMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalErrorAnalysisMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InternalErrorAnalysisMessage proto.InternalMessageInfo

func (m *InternalErrorAnalysisMessage) GetMessageBase() *AnalysisMessageBase {
	if m != nil {
		return m.MessageBase
	}
	return nil
}

func (m *InternalErrorAnalysisMessage) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func init() {
	proto.RegisterEnum("istio.analysis.v1alpha1.AnalysisMessageBase_Level", AnalysisMessageBase_Level_name, AnalysisMessageBase_Level_value)
	proto.RegisterType((*AnalysisMessageBase)(nil), "istio.analysis.v1alpha1.AnalysisMessageBase")
	proto.RegisterType((*AnalysisMessageBase_Type)(nil), "istio.analysis.v1alpha1.AnalysisMessageBase.Type")
	proto.RegisterType((*AnalysisMessageWeakSchema)(nil), "istio.analysis.v1alpha1.AnalysisMessageWeakSchema")
	proto.RegisterType((*AnalysisMessageWeakSchema_ArgType)(nil), "istio.analysis.v1alpha1.AnalysisMessageWeakSchema.ArgType")
	proto.RegisterType((*GenericAnalysisMessage)(nil), "istio.analysis.v1alpha1.GenericAnalysisMessage")
	proto.RegisterType((*InternalErrorAnalysisMessage)(nil), "istio.analysis.v1alpha1.InternalErrorAnalysisMessage")
}

func init() { proto.RegisterFile("analysis/v1alpha1/message.proto", fileDescriptor_4d6fd1414dfe3800) }

var fileDescriptor_4d6fd1414dfe3800 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x35, 0x6d, 0xfa, 0x75, 0xbb, 0x2e, 0x75, 0x84, 0x6d, 0x2d, 0x8b, 0x86, 0x82, 0x50, 0x58,
	0x99, 0xd0, 0x0a, 0xfb, 0xe0, 0x5b, 0x17, 0xea, 0x5a, 0xd4, 0x54, 0x66, 0x5d, 0x0a, 0xbe, 0x94,
	0x69, 0x7a, 0x4d, 0x83, 0x49, 0x26, 0xcc, 0x4c, 0x16, 0xfa, 0x0b, 0xfc, 0x3f, 0xbe, 0x8a, 0xff,
	0x4d, 0x92, 0x49, 0x74, 0x5d, 0x57, 0x70, 0x41, 0xdf, 0x66, 0x4e, 0xef, 0x3d, 0x73, 0xce, 0xb9,
	0xbd, 0x81, 0x27, 0x3c, 0xe1, 0xd1, 0x5e, 0x85, 0xca, 0xbd, 0x9a, 0xf0, 0x28, 0xdd, 0xf1, 0x89,
	0x1b, 0xa3, 0x52, 0x3c, 0x40, 0x9a, 0x4a, 0xa1, 0x05, 0xe9, 0x87, 0x4a, 0x87, 0x82, 0x56, 0x65,
	0xb4, 0x2a, 0x1b, 0x1e, 0x07, 0x42, 0x04, 0x11, 0xba, 0x45, 0xd9, 0x26, 0xfb, 0xe8, 0x2a, 0x2d,
	0x33, 0x5f, 0x9b, 0xb6, 0xd1, 0xd7, 0x1a, 0x3c, 0x9c, 0x95, 0x3d, 0x6f, 0x0d, 0xe1, 0x19, 0x57,
	0x48, 0xe6, 0x60, 0xeb, 0x7d, 0x8a, 0x03, 0xcb, 0xb1, 0xc6, 0xdd, 0xe9, 0x84, 0xfe, 0x81, 0x9d,
	0xde, 0xd2, 0x4b, 0xdf, 0xef, 0x53, 0x64, 0x45, 0x3b, 0x79, 0x05, 0x8d, 0x08, 0xaf, 0x30, 0x1a,
	0xd4, 0x1c, 0x6b, 0x7c, 0x38, 0x9d, 0xde, 0x89, 0xe7, 0x4d, 0xde, 0xc9, 0x0c, 0x01, 0x39, 0x81,
	0x07, 0x5b, 0xe1, 0x67, 0x31, 0x26, 0x9a, 0xeb, 0x50, 0x24, 0xeb, 0x4c, 0x46, 0x83, 0xba, 0x63,
	0x8d, 0x3b, 0xac, 0xf7, 0xcb, 0x0f, 0x97, 0x32, 0x1a, 0x52, 0xb0, 0x73, 0x11, 0x84, 0x80, 0x9d,
	0xf0, 0xd8, 0xb8, 0xe8, 0xb0, 0xe2, 0x9c, 0x63, 0xbe, 0xd8, 0x62, 0xa1, 0xa8, 0xc3, 0x8a, 0xf3,
	0xe8, 0x14, 0x1a, 0xc5, 0x63, 0xa4, 0x0b, 0xad, 0x4b, 0xef, 0xb5, 0xb7, 0x5c, 0x79, 0xbd, 0x7b,
	0xa4, 0x03, 0x8d, 0x39, 0x63, 0x4b, 0xd6, 0xab, 0xe7, 0xf8, 0x6a, 0xc6, 0xbc, 0x85, 0x77, 0xde,
	0x6b, 0x93, 0x36, 0xd8, 0x0b, 0xef, 0xe5, 0xb2, 0x77, 0x30, 0xfa, 0x52, 0x83, 0x47, 0x37, 0x94,
	0xaf, 0x90, 0x7f, 0xba, 0xf0, 0x77, 0x18, 0x73, 0xb2, 0x84, 0x83, 0x72, 0x46, 0xeb, 0x0d, 0x57,
	0x55, 0x96, 0xcf, 0xee, 0x92, 0x01, 0xeb, 0xc6, 0xd7, 0x86, 0xe2, 0x40, 0x77, 0x8b, 0xca, 0x97,
	0x61, 0x9a, 0x1b, 0x2d, 0x1d, 0x5c, 0x87, 0xc8, 0x10, 0xda, 0x1a, 0xe3, 0x34, 0xe2, 0x1a, 0xcb,
	0x70, 0x7e, 0xdc, 0x89, 0x07, 0x36, 0x97, 0x81, 0x1a, 0xd8, 0x4e, 0x7d, 0xdc, 0x9d, 0xbe, 0xf8,
	0x5b, 0x19, 0x3f, 0x0d, 0xd1, 0x99, 0x0c, 0xcc, 0x6c, 0x73, 0x9e, 0xe1, 0x29, 0xb4, 0x4a, 0xe0,
	0xd6, 0x9c, 0xfb, 0xd0, 0x0a, 0xc4, 0xba, 0xf8, 0x13, 0x19, 0xa1, 0xcd, 0x40, 0xe4, 0xc5, 0xa3,
	0x6f, 0x16, 0x1c, 0x9d, 0x63, 0x82, 0x32, 0xf4, 0x6f, 0x3c, 0xf5, 0xef, 0x13, 0x3b, 0x29, 0x3d,
	0xd7, 0x0a, 0xa2, 0x3e, 0x35, 0xbb, 0x40, 0xab, 0x5d, 0xa0, 0x17, 0xc5, 0x2e, 0x18, 0x43, 0xe4,
	0x29, 0x1c, 0x4a, 0x54, 0x22, 0x93, 0x3e, 0xae, 0x53, 0xae, 0x77, 0x6a, 0x50, 0x77, 0xea, 0xe3,
	0x0e, 0xbb, 0x5f, 0xa1, 0xef, 0x72, 0x70, 0xf4, 0xd9, 0x82, 0xe3, 0x45, 0xa2, 0x51, 0x26, 0x3c,
	0x9a, 0x4b, 0x29, 0xe4, 0x7f, 0x77, 0x71, 0x04, 0xcd, 0x2d, 0x6a, 0x1e, 0x46, 0x55, 0x92, 0xe6,
	0x76, 0xe6, 0x7c, 0x78, 0x6c, 0x38, 0x43, 0xe1, 0xf2, 0x34, 0x74, 0x7f, 0xfb, 0x46, 0x6c, 0x9a,
	0x85, 0xd3, 0xe7, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x35, 0x65, 0xf6, 0xbe, 0x3f, 0x04, 0x00,
	0x00,
}
