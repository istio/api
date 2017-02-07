// Code generated by protoc-gen-go.
// source: mixer/v1/config/descriptor/quota_descriptor.proto
// DO NOT EDIT!

package istio_mixer_v1_config_descriptor

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Configuration state for a particular quota.
//
// Quotas are similar to metrics, except that they are mutated through method
// calls and there are limits on the allowed values.
// The descriptor below lets you define a quota and indicate the maximum
// amount values of this quota are allowed to hold.
//
// A given quota is described by a set of attributes. These attributes represent
// the different dimensions to associate with the quota. A given quota holds a
// unique value for potentially any combination of these attributes.
//
// The quota kind controls the general behavior of the quota. An allocation
// quota is only adjusted through explicit method calls. A rate limit quota's
// values are reset to 0 automatically at a fixed interval.
type QuotaDescriptor struct {
	// The name of this descriptor.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A optional concise name for the quota, which can be displayed in user interfaces.
	// Use sentence case without an ending period, for example "Request count".
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	// An optional description of the quota, which can be used in documentation.
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// The name of the attribute that supplies the amount for a given
	// quota allocation or release operation.
	AmountAttribute string `protobuf:"bytes,4,opt,name=amount_attribute,json=amountAttribute" json:"amount_attribute,omitempty"`
	// The set of attributes that are necessary to describe a specific value cell
	// for a quota of this type.
	Attributes []string `protobuf:"bytes,5,rep,name=attributes" json:"attributes,omitempty"`
	// The default imposed maximum amount for values of this quota.
	MaxAmount int64 `protobuf:"varint,6,opt,name=max_amount,json=maxAmount" json:"max_amount,omitempty"`
	// The amount of time allocated quota remains valid before it is
	// automatically released. If this is 0, then allocated quota is
	// not automatically released.
	ExpirationSeconds int32 `protobuf:"varint,7,opt,name=expiration_seconds,json=expirationSeconds" json:"expiration_seconds,omitempty"`
}

func (m *QuotaDescriptor) Reset()                    { *m = QuotaDescriptor{} }
func (m *QuotaDescriptor) String() string            { return proto.CompactTextString(m) }
func (*QuotaDescriptor) ProtoMessage()               {}
func (*QuotaDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *QuotaDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *QuotaDescriptor) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *QuotaDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *QuotaDescriptor) GetAmountAttribute() string {
	if m != nil {
		return m.AmountAttribute
	}
	return ""
}

func (m *QuotaDescriptor) GetAttributes() []string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *QuotaDescriptor) GetMaxAmount() int64 {
	if m != nil {
		return m.MaxAmount
	}
	return 0
}

func (m *QuotaDescriptor) GetExpirationSeconds() int32 {
	if m != nil {
		return m.ExpirationSeconds
	}
	return 0
}

func init() {
	proto.RegisterType((*QuotaDescriptor)(nil), "istio.mixer.v1.config.descriptor.QuotaDescriptor")
}

func init() { proto.RegisterFile("mixer/v1/config/descriptor/quota_descriptor.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xe9, 0xba, 0x4d, 0xfa, 0x4d, 0x98, 0xe6, 0x94, 0x8b, 0x12, 0x3d, 0xd5, 0x83, 0x2d,
	0xc5, 0x5f, 0x30, 0xf0, 0x2c, 0x58, 0x7f, 0x40, 0xc8, 0xda, 0x28, 0x1f, 0x98, 0x7c, 0x35, 0x49,
	0x47, 0x3d, 0xfb, 0xc7, 0xa5, 0xe9, 0xb6, 0xee, 0x16, 0x9e, 0xf7, 0x79, 0x12, 0x08, 0x54, 0x06,
	0x07, 0xed, 0xca, 0x43, 0x55, 0x36, 0x64, 0x3f, 0xf1, 0xab, 0x6c, 0xb5, 0x6f, 0x1c, 0x76, 0x81,
	0x5c, 0xf9, 0xd3, 0x53, 0x50, 0x72, 0x06, 0x45, 0xe7, 0x28, 0x10, 0x13, 0xe8, 0x03, 0x52, 0x11,
	0xc3, 0xe2, 0x50, 0x15, 0x53, 0x58, 0xcc, 0xde, 0xe3, 0xdf, 0x02, 0xb6, 0xef, 0x63, 0xfc, 0x7a,
	0x66, 0x8c, 0xc1, 0xd2, 0x2a, 0xa3, 0x79, 0x22, 0x92, 0x3c, 0xab, 0xe3, 0x99, 0x3d, 0xc0, 0x75,
	0x8b, 0xbe, 0xfb, 0x56, 0xbf, 0x32, 0x6e, 0x8b, 0xb8, 0x6d, 0x8e, 0xec, 0x6d, 0x54, 0x04, 0x6c,
	0x4e, 0x17, 0x23, 0x59, 0x9e, 0x1e, 0x8d, 0x19, 0xb1, 0x27, 0xb8, 0x51, 0x86, 0x7a, 0x1b, 0xa4,
	0x0a, 0xc1, 0xe1, 0xbe, 0x0f, 0x9a, 0x2f, 0xa3, 0xb6, 0x9d, 0xf8, 0xee, 0x84, 0xd9, 0x3d, 0xc0,
	0xd9, 0xf1, 0x7c, 0x25, 0xd2, 0x3c, 0xab, 0x2f, 0x08, 0xbb, 0x03, 0x30, 0x6a, 0x90, 0x53, 0xc6,
	0xd7, 0x22, 0xc9, 0xd3, 0x3a, 0x33, 0x6a, 0xd8, 0x45, 0xc0, 0x9e, 0x81, 0xe9, 0xa1, 0x43, 0xa7,
	0xc6, 0x77, 0xa5, 0xd7, 0x0d, 0xd9, 0xd6, 0xf3, 0x2b, 0x91, 0xe4, 0xab, 0xfa, 0x76, 0x5e, 0x3e,
	0xa6, 0x61, 0xbf, 0x8e, 0xdf, 0xf5, 0xf2, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xaa, 0x29, 0xb7,
	0x63, 0x01, 0x00, 0x00,
}
