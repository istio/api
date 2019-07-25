// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/model/v1beta1/quota.proto

package v1beta1

import (
	fmt "fmt"
	rpc "github.com/gogo/googleapis/google/rpc"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
	time "time"
)

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

// Expresses the quota allocation request.
type QuotaRequest struct {
	// The individual quotas to allocate
	Quotas map[string]QuotaRequest_QuotaParams `protobuf:"bytes,1,rep,name=quotas,proto3" json:"quotas" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *QuotaRequest) Reset()      { *m = QuotaRequest{} }
func (*QuotaRequest) ProtoMessage() {}
func (*QuotaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f07acf62b4429357, []int{0}
}
func (m *QuotaRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuotaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuotaRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuotaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotaRequest.Merge(m, src)
}
func (m *QuotaRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuotaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuotaRequest proto.InternalMessageInfo

// parameters for a quota allocation
type QuotaRequest_QuotaParams struct {
	// Amount of quota to allocate
	Amount int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// When true, supports returning less quota than what was requested.
	BestEffort bool `protobuf:"varint,2,opt,name=best_effort,json=bestEffort,proto3" json:"best_effort,omitempty"`
}

func (m *QuotaRequest_QuotaParams) Reset()      { *m = QuotaRequest_QuotaParams{} }
func (*QuotaRequest_QuotaParams) ProtoMessage() {}
func (*QuotaRequest_QuotaParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_f07acf62b4429357, []int{0, 0}
}
func (m *QuotaRequest_QuotaParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuotaRequest_QuotaParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuotaRequest_QuotaParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuotaRequest_QuotaParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotaRequest_QuotaParams.Merge(m, src)
}
func (m *QuotaRequest_QuotaParams) XXX_Size() int {
	return m.Size()
}
func (m *QuotaRequest_QuotaParams) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotaRequest_QuotaParams.DiscardUnknown(m)
}

var xxx_messageInfo_QuotaRequest_QuotaParams proto.InternalMessageInfo

// Expresses the result of multiple quota allocations.
type QuotaResult struct {
	// The resulting quota, one entry per requested quota.
	Quotas map[string]QuotaResult_Result `protobuf:"bytes,1,rep,name=quotas,proto3" json:"quotas" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *QuotaResult) Reset()      { *m = QuotaResult{} }
func (*QuotaResult) ProtoMessage() {}
func (*QuotaResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_f07acf62b4429357, []int{1}
}
func (m *QuotaResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuotaResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuotaResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuotaResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotaResult.Merge(m, src)
}
func (m *QuotaResult) XXX_Size() int {
	return m.Size()
}
func (m *QuotaResult) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotaResult.DiscardUnknown(m)
}

var xxx_messageInfo_QuotaResult proto.InternalMessageInfo

// Expresses the result of a quota allocation.
type QuotaResult_Result struct {
	// The amount of time for which this result can be considered valid.
	ValidDuration time.Duration `protobuf:"bytes,2,opt,name=valid_duration,json=validDuration,proto3,stdduration" json:"valid_duration"`
	// The amount of granted quota. When `QuotaParams.best_effort` is true, this will be >= 0.
	// If `QuotaParams.best_effort` is false, this will be either 0 or >= `QuotaParams.amount`.
	GrantedAmount int64 `protobuf:"varint,3,opt,name=granted_amount,json=grantedAmount,proto3" json:"granted_amount,omitempty"`
	// A status code of OK indicates quota was fetched successfully. Any other code indicates error in fetching quota.
	Status rpc.Status `protobuf:"bytes,4,opt,name=status,proto3" json:"status"`
}

func (m *QuotaResult_Result) Reset()      { *m = QuotaResult_Result{} }
func (*QuotaResult_Result) ProtoMessage() {}
func (*QuotaResult_Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_f07acf62b4429357, []int{1, 0}
}
func (m *QuotaResult_Result) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuotaResult_Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuotaResult_Result.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuotaResult_Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotaResult_Result.Merge(m, src)
}
func (m *QuotaResult_Result) XXX_Size() int {
	return m.Size()
}
func (m *QuotaResult_Result) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotaResult_Result.DiscardUnknown(m)
}

var xxx_messageInfo_QuotaResult_Result proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QuotaRequest)(nil), "istio.mixer.adapter.model.v1beta1.QuotaRequest")
	proto.RegisterMapType((map[string]QuotaRequest_QuotaParams)(nil), "istio.mixer.adapter.model.v1beta1.QuotaRequest.QuotasEntry")
	proto.RegisterType((*QuotaRequest_QuotaParams)(nil), "istio.mixer.adapter.model.v1beta1.QuotaRequest.QuotaParams")
	proto.RegisterType((*QuotaResult)(nil), "istio.mixer.adapter.model.v1beta1.QuotaResult")
	proto.RegisterMapType((map[string]QuotaResult_Result)(nil), "istio.mixer.adapter.model.v1beta1.QuotaResult.QuotasEntry")
	proto.RegisterType((*QuotaResult_Result)(nil), "istio.mixer.adapter.model.v1beta1.QuotaResult.Result")
}

func init() {
	proto.RegisterFile("mixer/adapter/model/v1beta1/quota.proto", fileDescriptor_f07acf62b4429357)
}

var fileDescriptor_f07acf62b4429357 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x31, 0x6f, 0x13, 0x31,
	0x18, 0xb5, 0x73, 0xe5, 0x54, 0x1c, 0x5a, 0x21, 0x0b, 0x41, 0xb8, 0xc1, 0x09, 0x95, 0x10, 0x99,
	0x6c, 0x0a, 0x42, 0x42, 0x65, 0x22, 0xa2, 0x0c, 0xb0, 0xd0, 0x63, 0x01, 0x96, 0xc8, 0xd7, 0x73,
	0x4e, 0x27, 0x2e, 0xf1, 0xd5, 0xf6, 0x45, 0x74, 0x63, 0x65, 0x63, 0xe4, 0x27, 0xc0, 0xc0, 0xff,
	0xc8, 0x98, 0xb1, 0x13, 0x90, 0xcb, 0xc2, 0xd8, 0x9f, 0x80, 0xce, 0xf6, 0x49, 0x85, 0xa1, 0x05,
	0xa6, 0xd8, 0xdf, 0xf7, 0xbe, 0xf7, 0x3e, 0xbf, 0x97, 0x43, 0x77, 0xa6, 0xf9, 0x3b, 0xa1, 0x18,
	0x4f, 0x79, 0x69, 0x84, 0x62, 0x53, 0x99, 0x8a, 0x82, 0xcd, 0x77, 0x13, 0x61, 0xf8, 0x2e, 0x3b,
	0xaa, 0xa4, 0xe1, 0xb4, 0x54, 0xd2, 0x48, 0x7c, 0x2b, 0xd7, 0x26, 0x97, 0xd4, 0xc2, 0xa9, 0x87,
	0x53, 0x0b, 0xa7, 0x1e, 0x1e, 0x5d, 0xcb, 0x64, 0x26, 0x2d, 0x9a, 0x35, 0x27, 0x37, 0x18, 0x91,
	0x4c, 0xca, 0xac, 0x10, 0xcc, 0xde, 0x92, 0x6a, 0xc2, 0xd2, 0x4a, 0x71, 0x93, 0xcb, 0x99, 0xef,
	0xdf, 0xf0, 0x7d, 0x55, 0x1e, 0x32, 0x6d, 0xb8, 0xa9, 0xb4, 0x6b, 0xec, 0x7c, 0xe9, 0xa0, 0x2b,
	0x07, 0xcd, 0x06, 0xb1, 0x38, 0xaa, 0x84, 0x36, 0xf8, 0x35, 0x0a, 0xed, 0x46, 0xba, 0x07, 0x07,
	0xc1, 0xb0, 0x7b, 0xef, 0x11, 0xbd, 0x70, 0x27, 0x7a, 0x96, 0xc0, 0x5d, 0xf4, 0xfe, 0xcc, 0xa8,
	0xe3, 0xd1, 0xc6, 0xe2, 0x5b, 0x1f, 0xc4, 0x9e, 0x30, 0x7a, 0x8a, 0xba, 0xb6, 0xf9, 0x82, 0x2b,
	0x3e, 0xd5, 0xf8, 0x3a, 0x0a, 0xf9, 0x54, 0x56, 0x33, 0xd3, 0x83, 0x03, 0x38, 0x0c, 0x62, 0x7f,
	0xc3, 0x7d, 0xd4, 0x4d, 0x84, 0x36, 0x63, 0x31, 0x99, 0x48, 0x65, 0x7a, 0x9d, 0x01, 0x1c, 0x6e,
	0xc6, 0xa8, 0x29, 0xed, 0xdb, 0x4a, 0x34, 0xf7, 0x3c, 0x4e, 0x04, 0x5f, 0x45, 0xc1, 0x5b, 0x71,
	0x6c, 0x49, 0x2e, 0xc7, 0xcd, 0x11, 0x1f, 0xa0, 0x4b, 0x73, 0x5e, 0x54, 0xc2, 0xce, 0xfe, 0xef,
	0x13, 0xdc, 0x96, 0xb1, 0x63, 0xda, 0xeb, 0x3c, 0x84, 0x3b, 0x1f, 0x02, 0x2f, 0x1c, 0x0b, 0x5d,
	0x15, 0x06, 0xbf, 0xfa, 0xc3, 0xaa, 0xbd, 0xbf, 0xd7, 0x69, 0xe6, 0xcf, 0x71, 0xea, 0x2b, 0x44,
	0xa1, 0x17, 0x79, 0x86, 0xb6, 0xe7, 0xbc, 0xc8, 0xd3, 0x71, 0x9b, 0xa8, 0x7f, 0xd4, 0x4d, 0xea,
	0x22, 0xa5, 0x6d, 0xe4, 0xf4, 0x89, 0x07, 0x8c, 0x36, 0x1b, 0xae, 0x4f, 0xdf, 0xfb, 0x30, 0xde,
	0xb2, 0xa3, 0x6d, 0x03, 0xdf, 0x46, 0xdb, 0x99, 0xe2, 0x33, 0x23, 0xd2, 0xb1, 0x77, 0x3e, 0xb0,
	0xce, 0x6f, 0xf9, 0xea, 0x63, 0x17, 0xc0, 0x5d, 0x14, 0xba, 0xff, 0x48, 0x6f, 0xc3, 0x4a, 0xe1,
	0x56, 0x4a, 0x95, 0x87, 0xf4, 0xa5, 0xed, 0xb4, 0xfb, 0x3a, 0x5c, 0x54, 0x5e, 0x94, 0xc8, 0xf3,
	0xdf, 0x13, 0x79, 0xf0, 0x8f, 0x4e, 0xb9, 0x9f, 0x33, 0x59, 0x8c, 0x92, 0xc5, 0x8a, 0x80, 0xe5,
	0x8a, 0x80, 0x93, 0x15, 0x01, 0xa7, 0x2b, 0x02, 0xde, 0xd7, 0x04, 0x7e, 0xae, 0x09, 0x58, 0xd4,
	0x04, 0x2e, 0x6b, 0x02, 0x7f, 0xd4, 0x04, 0xfe, 0xac, 0x09, 0x38, 0xad, 0x09, 0xfc, 0xb8, 0x26,
	0x60, 0xb9, 0x26, 0xe0, 0x64, 0x4d, 0xc0, 0x9b, 0xa1, 0x93, 0xce, 0x25, 0xe3, 0x65, 0xce, 0xce,
	0xf9, 0x32, 0x93, 0xd0, 0x5a, 0x7b, 0xff, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xcd, 0xf9,
	0x75, 0xbf, 0x03, 0x00, 0x00,
}

func (m *QuotaRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuotaRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuotaRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Quotas) > 0 {
		for k := range m.Quotas {
			v := m.Quotas[k]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuota(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintQuota(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintQuota(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QuotaRequest_QuotaParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuotaRequest_QuotaParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuotaRequest_QuotaParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BestEffort {
		i--
		if m.BestEffort {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Amount != 0 {
		i = encodeVarintQuota(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QuotaResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuotaResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuotaResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Quotas) > 0 {
		for k := range m.Quotas {
			v := m.Quotas[k]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuota(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintQuota(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintQuota(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QuotaResult_Result) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuotaResult_Result) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuotaResult_Result) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuota(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.GrantedAmount != 0 {
		i = encodeVarintQuota(dAtA, i, uint64(m.GrantedAmount))
		i--
		dAtA[i] = 0x18
	}
	n4, err4 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.ValidDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.ValidDuration):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintQuota(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}

func encodeVarintQuota(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuota(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QuotaRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Quotas) > 0 {
		for k, v := range m.Quotas {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovQuota(uint64(len(k))) + 1 + l + sovQuota(uint64(l))
			n += mapEntrySize + 1 + sovQuota(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *QuotaRequest_QuotaParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Amount != 0 {
		n += 1 + sovQuota(uint64(m.Amount))
	}
	if m.BestEffort {
		n += 2
	}
	return n
}

func (m *QuotaResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Quotas) > 0 {
		for k, v := range m.Quotas {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovQuota(uint64(len(k))) + 1 + l + sovQuota(uint64(l))
			n += mapEntrySize + 1 + sovQuota(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *QuotaResult_Result) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.ValidDuration)
	n += 1 + l + sovQuota(uint64(l))
	if m.GrantedAmount != 0 {
		n += 1 + sovQuota(uint64(m.GrantedAmount))
	}
	l = m.Status.Size()
	n += 1 + l + sovQuota(uint64(l))
	return n
}

func sovQuota(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuota(x uint64) (n int) {
	return sovQuota(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *QuotaRequest) String() string {
	if this == nil {
		return "nil"
	}
	keysForQuotas := make([]string, 0, len(this.Quotas))
	for k, _ := range this.Quotas {
		keysForQuotas = append(keysForQuotas, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForQuotas)
	mapStringForQuotas := "map[string]QuotaRequest_QuotaParams{"
	for _, k := range keysForQuotas {
		mapStringForQuotas += fmt.Sprintf("%v: %v,", k, this.Quotas[k])
	}
	mapStringForQuotas += "}"
	s := strings.Join([]string{`&QuotaRequest{`,
		`Quotas:` + mapStringForQuotas + `,`,
		`}`,
	}, "")
	return s
}
func (this *QuotaRequest_QuotaParams) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QuotaRequest_QuotaParams{`,
		`Amount:` + fmt.Sprintf("%v", this.Amount) + `,`,
		`BestEffort:` + fmt.Sprintf("%v", this.BestEffort) + `,`,
		`}`,
	}, "")
	return s
}
func (this *QuotaResult) String() string {
	if this == nil {
		return "nil"
	}
	keysForQuotas := make([]string, 0, len(this.Quotas))
	for k, _ := range this.Quotas {
		keysForQuotas = append(keysForQuotas, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForQuotas)
	mapStringForQuotas := "map[string]QuotaResult_Result{"
	for _, k := range keysForQuotas {
		mapStringForQuotas += fmt.Sprintf("%v: %v,", k, this.Quotas[k])
	}
	mapStringForQuotas += "}"
	s := strings.Join([]string{`&QuotaResult{`,
		`Quotas:` + mapStringForQuotas + `,`,
		`}`,
	}, "")
	return s
}
func (this *QuotaResult_Result) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QuotaResult_Result{`,
		`ValidDuration:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ValidDuration), "Duration", "types.Duration", 1), `&`, ``, 1) + `,`,
		`GrantedAmount:` + fmt.Sprintf("%v", this.GrantedAmount) + `,`,
		`Status:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Status), "Status", "rpc.Status", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringQuota(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *QuotaRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuota
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
			return fmt.Errorf("proto: QuotaRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuotaRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quotas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
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
				return ErrInvalidLengthQuota
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuota
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Quotas == nil {
				m.Quotas = make(map[string]QuotaRequest_QuotaParams)
			}
			var mapkey string
			mapvalue := &QuotaRequest_QuotaParams{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuota
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuota
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthQuota
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthQuota
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuota
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthQuota
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthQuota
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &QuotaRequest_QuotaParams{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipQuota(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthQuota
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Quotas[mapkey] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuota(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuota
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuota
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
func (m *QuotaRequest_QuotaParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuota
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
			return fmt.Errorf("proto: QuotaParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuotaParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BestEffort", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.BestEffort = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipQuota(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuota
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuota
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
func (m *QuotaResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuota
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
			return fmt.Errorf("proto: QuotaResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuotaResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quotas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
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
				return ErrInvalidLengthQuota
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuota
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Quotas == nil {
				m.Quotas = make(map[string]QuotaResult_Result)
			}
			var mapkey string
			mapvalue := &QuotaResult_Result{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuota
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuota
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthQuota
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthQuota
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuota
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthQuota
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthQuota
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &QuotaResult_Result{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipQuota(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthQuota
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Quotas[mapkey] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuota(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuota
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuota
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
func (m *QuotaResult_Result) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuota
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
			return fmt.Errorf("proto: Result: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Result: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
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
				return ErrInvalidLengthQuota
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuota
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.ValidDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrantedAmount", wireType)
			}
			m.GrantedAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GrantedAmount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
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
				return ErrInvalidLengthQuota
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuota
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuota(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuota
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuota
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
func skipQuota(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuota
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
					return 0, ErrIntOverflowQuota
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
					return 0, ErrIntOverflowQuota
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
				return 0, ErrInvalidLengthQuota
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthQuota
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowQuota
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
				next, err := skipQuota(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthQuota
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
	ErrInvalidLengthQuota = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuota   = fmt.Errorf("proto: integer overflow")
)
