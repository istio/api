// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rbac/v1alpha1/local_rbac.proto

/*
Package v1alpha1 is a generated protocol buffer package.

It is generated from these files:
	rbac/v1alpha1/local_rbac.proto
	rbac/v1alpha1/rbac.proto

It has these top-level messages:
	LocalRBAC
	ServicePolicy
	Policy
	ServiceRole
	AccessRule
	ServiceRoleBinding
	Subject
	RoleRef
*/
package v1alpha1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// LocalRBAC specifies Istio RBAC policies that apply to the local RBAC engine
// that typically runs on a proxy. There can be multiple services running behind
// a proxy.
//
// Here is an example of local RBAC policy, which contains policies for a single service.
// [
//   {
//     service: products.default.svc.cluster.local
//     policies:
//     - role: service-admin
//       spec:
//         rules:
//         - services: [“*”]
//           methods: [“*”]
//       bindings:
//       - subjects:
//         - user: “cluster.local/ns/default/sa/admin”
//         roleRef:
//           kind: ServiceRole
//           name: “service-admin”
//     - role: product-viewer
//       spec:
//         rules:
//         - services: [“products.default.svc.cluster.local”]
//           paths: [“/products/*”, “*/reviews”]
//           methods: [“GET”, “HEAD”]
//       bindings:
//       - subjects:
//         - user: “alice@yahoo.com”
//         roleRef:
//           kind: ServiceRole
//           name: “product-viewer”
//    },
//  ]
type LocalRBAC struct {
	// A list of ServicePolicy. One per service.
	ServicePolicies []*ServicePolicy `protobuf:"bytes,1,rep,name=service_policies,json=servicePolicies" json:"service_policies,omitempty"`
}

func (m *LocalRBAC) Reset()                    { *m = LocalRBAC{} }
func (m *LocalRBAC) String() string            { return proto.CompactTextString(m) }
func (*LocalRBAC) ProtoMessage()               {}
func (*LocalRBAC) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LocalRBAC) GetServicePolicies() []*ServicePolicy {
	if m != nil {
		return m.ServicePolicies
	}
	return nil
}

// ServicePolicy specifies a service and the Istio RBAC policies that apply to the service.
type ServicePolicy struct {
	// The fully qualified name of the service.
	Service string `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	// The Istio RBAC policies that apply to the service.
	Policies []*Policy `protobuf:"bytes,2,rep,name=policies" json:"policies,omitempty"`
}

func (m *ServicePolicy) Reset()                    { *m = ServicePolicy{} }
func (m *ServicePolicy) String() string            { return proto.CompactTextString(m) }
func (*ServicePolicy) ProtoMessage()               {}
func (*ServicePolicy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServicePolicy) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *ServicePolicy) GetPolicies() []*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

// Policy specifies a ServiceRole and the associated ServiceRoleBindings.
type Policy struct {
	// The name of the ServiceRole.
	Role string `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	// The ServiceRole specification.
	Spec *ServiceRole `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	// The ServiceRoleBindings that are associated with the given ServiceRole.
	Bindings []*ServiceRoleBinding `protobuf:"bytes,3,rep,name=bindings" json:"bindings,omitempty"`
}

func (m *Policy) Reset()                    { *m = Policy{} }
func (m *Policy) String() string            { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()               {}
func (*Policy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Policy) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Policy) GetSpec() *ServiceRole {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Policy) GetBindings() []*ServiceRoleBinding {
	if m != nil {
		return m.Bindings
	}
	return nil
}

func init() {
	proto.RegisterType((*LocalRBAC)(nil), "istio.rbac.v1alpha1.LocalRBAC")
	proto.RegisterType((*ServicePolicy)(nil), "istio.rbac.v1alpha1.ServicePolicy")
	proto.RegisterType((*Policy)(nil), "istio.rbac.v1alpha1.Policy")
}

func init() { proto.RegisterFile("rbac/v1alpha1/local_rbac.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x49, 0x5b, 0x6a, 0x3b, 0x45, 0x94, 0x78, 0x09, 0x55, 0x64, 0xd9, 0x8b, 0x7b, 0xda,
	0xa5, 0x55, 0xf0, 0xec, 0xf6, 0xaa, 0x20, 0xf1, 0xd6, 0x4b, 0xc9, 0xc6, 0xa0, 0x03, 0xa1, 0x13,
	0x36, 0xa5, 0xe0, 0x27, 0xf1, 0xeb, 0xca, 0x66, 0xff, 0xe8, 0xc2, 0x62, 0x6f, 0x99, 0xc7, 0xef,
	0xbd, 0x97, 0x64, 0xe0, 0xb6, 0x2c, 0x94, 0xce, 0x8e, 0x2b, 0x65, 0xdd, 0xa7, 0x5a, 0x65, 0x96,
	0xb4, 0xb2, 0xbb, 0x4a, 0x4b, 0x5d, 0x49, 0x07, 0xe2, 0x57, 0xe8, 0x0f, 0x48, 0x69, 0x50, 0x5a,
	0x6a, 0x29, 0xfa, 0xa6, 0x5f, 0x3c, 0xde, 0xc2, 0xfc, 0xb9, 0x8a, 0x90, 0xf9, 0xd3, 0x86, 0xbf,
	0xc0, 0xa5, 0x37, 0xe5, 0x11, 0xb5, 0xd9, 0x39, 0xb2, 0xa8, 0xd1, 0x78, 0xc1, 0xa2, 0x71, 0xb2,
	0x58, 0xc7, 0xe9, 0x40, 0x6c, 0xfa, 0x56, 0xc3, 0xaf, 0x15, 0xfb, 0x25, 0x2f, 0xfc, 0x9f, 0x11,
	0x8d, 0x8f, 0x0b, 0x38, 0xef, 0x11, 0x5c, 0xc0, 0x59, 0xc3, 0x08, 0x16, 0xb1, 0x64, 0x2e, 0xdb,
	0x91, 0x3f, 0xc2, 0xac, 0x6b, 0x1c, 0x85, 0xc6, 0xeb, 0xc1, 0xc6, 0xa6, 0xaa, 0x83, 0xe3, 0x6f,
	0x06, 0xd3, 0x26, 0x9d, 0xc3, 0xa4, 0x24, 0xdb, 0x46, 0x87, 0x33, 0x7f, 0x80, 0x89, 0x77, 0x46,
	0x8b, 0x51, 0xc4, 0x92, 0xc5, 0x3a, 0xfa, 0xef, 0x15, 0x92, 0xac, 0x91, 0x81, 0xe6, 0x1b, 0x98,
	0x15, 0xb8, 0x7f, 0xc7, 0xfd, 0x87, 0x17, 0xe3, 0x70, 0x9b, 0xbb, 0x53, 0xce, 0xbc, 0xe6, 0x65,
	0x67, 0xcc, 0x6f, 0xb6, 0xcb, 0xda, 0x83, 0x94, 0x29, 0x87, 0x59, 0x6f, 0x05, 0xc5, 0x34, 0x7c,
	0xff, 0xfd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x33, 0x6b, 0xdc, 0xcf, 0x01, 0x00, 0x00,
}
