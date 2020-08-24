// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mesh/v1alpha1/network.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google/api/field_behavior.proto"
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

// Network provides information about the endpoints in a routable L3
// network. A single routable L3 network can have one or more service
// registries. Note that the network has no relation to the locality of the
// endpoint. The endpoint locality will be obtained from the service
// registry.
type Network struct {
	// The list of endpoints in the network (obtained through the
	// constituent service registries or from CIDR ranges). All endpoints in
	// the network are directly accessible to one another.
	Endpoints []*Network_NetworkEndpoints `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	// Set of gateways associated with the network.
	Gateways             []*Network_IstioNetworkGateway `protobuf:"bytes,3,rep,name=gateways,proto3" json:"gateways,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Network) Reset()         { *m = Network{} }
func (m *Network) String() string { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()    {}
func (*Network) Descriptor() ([]byte, []int) {
	return fileDescriptor_a15df2a96e10cd86, []int{0}
}

func (m *Network) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Network.Unmarshal(m, b)
}
func (m *Network) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Network.Marshal(b, m, deterministic)
}
func (m *Network) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Network.Merge(m, src)
}
func (m *Network) XXX_Size() int {
	return xxx_messageInfo_Network.Size(m)
}
func (m *Network) XXX_DiscardUnknown() {
	xxx_messageInfo_Network.DiscardUnknown(m)
}

var xxx_messageInfo_Network proto.InternalMessageInfo

func (m *Network) GetEndpoints() []*Network_NetworkEndpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *Network) GetGateways() []*Network_IstioNetworkGateway {
	if m != nil {
		return m.Gateways
	}
	return nil
}

// NetworkEndpoints describes how the network associated with an endpoint
// should be inferred. An endpoint will be assigned to a network based on
// the following rules:
//
// 1. Implicitly: If the registry explicitly provides information about
// the network to which the endpoint belongs to. In some cases, its
// possible to indicate the network associated with the endpoint by
// adding the `ISTIO_META_NETWORK` environment variable to the sidecar.
//
// 2. Explicitly:
//
//    a. By matching the registry name with one of the "fromRegistry"
//    in the mesh config. A "from_registry" can only be assigned to a
//    single network.
//
//    b. By matching the IP against one of the CIDR ranges in a mesh
//    config network. The CIDR ranges must not overlap and be assigned to
//    a single network.
//
// (2) will override (1) if both are present.
type Network_NetworkEndpoints struct {
	// Types that are valid to be assigned to Ne:
	//	*Network_NetworkEndpoints_FromCidr
	//	*Network_NetworkEndpoints_FromRegistry
	Ne                   isNetwork_NetworkEndpoints_Ne `protobuf_oneof:"ne"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *Network_NetworkEndpoints) Reset()         { *m = Network_NetworkEndpoints{} }
func (m *Network_NetworkEndpoints) String() string { return proto.CompactTextString(m) }
func (*Network_NetworkEndpoints) ProtoMessage()    {}
func (*Network_NetworkEndpoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_a15df2a96e10cd86, []int{0, 0}
}

func (m *Network_NetworkEndpoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Network_NetworkEndpoints.Unmarshal(m, b)
}
func (m *Network_NetworkEndpoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Network_NetworkEndpoints.Marshal(b, m, deterministic)
}
func (m *Network_NetworkEndpoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Network_NetworkEndpoints.Merge(m, src)
}
func (m *Network_NetworkEndpoints) XXX_Size() int {
	return xxx_messageInfo_Network_NetworkEndpoints.Size(m)
}
func (m *Network_NetworkEndpoints) XXX_DiscardUnknown() {
	xxx_messageInfo_Network_NetworkEndpoints.DiscardUnknown(m)
}

var xxx_messageInfo_Network_NetworkEndpoints proto.InternalMessageInfo

type isNetwork_NetworkEndpoints_Ne interface {
	isNetwork_NetworkEndpoints_Ne()
}

type Network_NetworkEndpoints_FromCidr struct {
	FromCidr string `protobuf:"bytes,1,opt,name=from_cidr,json=fromCidr,proto3,oneof"`
}

type Network_NetworkEndpoints_FromRegistry struct {
	FromRegistry string `protobuf:"bytes,2,opt,name=from_registry,json=fromRegistry,proto3,oneof"`
}

func (*Network_NetworkEndpoints_FromCidr) isNetwork_NetworkEndpoints_Ne() {}

func (*Network_NetworkEndpoints_FromRegistry) isNetwork_NetworkEndpoints_Ne() {}

func (m *Network_NetworkEndpoints) GetNe() isNetwork_NetworkEndpoints_Ne {
	if m != nil {
		return m.Ne
	}
	return nil
}

func (m *Network_NetworkEndpoints) GetFromCidr() string {
	if x, ok := m.GetNe().(*Network_NetworkEndpoints_FromCidr); ok {
		return x.FromCidr
	}
	return ""
}

func (m *Network_NetworkEndpoints) GetFromRegistry() string {
	if x, ok := m.GetNe().(*Network_NetworkEndpoints_FromRegistry); ok {
		return x.FromRegistry
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Network_NetworkEndpoints) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Network_NetworkEndpoints_FromCidr)(nil),
		(*Network_NetworkEndpoints_FromRegistry)(nil),
	}
}

// The gateway associated with this network. Traffic from remote networks
// will arrive at the specified gateway:port. All incoming traffic must
// use mTLS.
type Network_IstioNetworkGateway struct {
	// Types that are valid to be assigned to Gw:
	//	*Network_IstioNetworkGateway_RegistryServiceName
	//	*Network_IstioNetworkGateway_Address
	Gw isNetwork_IstioNetworkGateway_Gw `protobuf_oneof:"gw"`
	// The port associated with the gateway.
	Port uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	// The locality associated with an explicitly specified gateway (i.e. ip)
	Locality             string   `protobuf:"bytes,4,opt,name=locality,proto3" json:"locality,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Network_IstioNetworkGateway) Reset()         { *m = Network_IstioNetworkGateway{} }
func (m *Network_IstioNetworkGateway) String() string { return proto.CompactTextString(m) }
func (*Network_IstioNetworkGateway) ProtoMessage()    {}
func (*Network_IstioNetworkGateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_a15df2a96e10cd86, []int{0, 1}
}

func (m *Network_IstioNetworkGateway) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Network_IstioNetworkGateway.Unmarshal(m, b)
}
func (m *Network_IstioNetworkGateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Network_IstioNetworkGateway.Marshal(b, m, deterministic)
}
func (m *Network_IstioNetworkGateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Network_IstioNetworkGateway.Merge(m, src)
}
func (m *Network_IstioNetworkGateway) XXX_Size() int {
	return xxx_messageInfo_Network_IstioNetworkGateway.Size(m)
}
func (m *Network_IstioNetworkGateway) XXX_DiscardUnknown() {
	xxx_messageInfo_Network_IstioNetworkGateway.DiscardUnknown(m)
}

var xxx_messageInfo_Network_IstioNetworkGateway proto.InternalMessageInfo

type isNetwork_IstioNetworkGateway_Gw interface {
	isNetwork_IstioNetworkGateway_Gw()
}

type Network_IstioNetworkGateway_RegistryServiceName struct {
	RegistryServiceName string `protobuf:"bytes,1,opt,name=registry_service_name,json=registryServiceName,proto3,oneof"`
}

type Network_IstioNetworkGateway_Address struct {
	Address string `protobuf:"bytes,2,opt,name=address,proto3,oneof"`
}

func (*Network_IstioNetworkGateway_RegistryServiceName) isNetwork_IstioNetworkGateway_Gw() {}

func (*Network_IstioNetworkGateway_Address) isNetwork_IstioNetworkGateway_Gw() {}

func (m *Network_IstioNetworkGateway) GetGw() isNetwork_IstioNetworkGateway_Gw {
	if m != nil {
		return m.Gw
	}
	return nil
}

func (m *Network_IstioNetworkGateway) GetRegistryServiceName() string {
	if x, ok := m.GetGw().(*Network_IstioNetworkGateway_RegistryServiceName); ok {
		return x.RegistryServiceName
	}
	return ""
}

func (m *Network_IstioNetworkGateway) GetAddress() string {
	if x, ok := m.GetGw().(*Network_IstioNetworkGateway_Address); ok {
		return x.Address
	}
	return ""
}

func (m *Network_IstioNetworkGateway) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Network_IstioNetworkGateway) GetLocality() string {
	if m != nil {
		return m.Locality
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Network_IstioNetworkGateway) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Network_IstioNetworkGateway_RegistryServiceName)(nil),
		(*Network_IstioNetworkGateway_Address)(nil),
	}
}

// MeshNetworks (config map) provides information about the set of networks
// inside a mesh and how to route to endpoints in each network. For example
//
// MeshNetworks(file/config map):
//
// ```yaml
// networks:
//   network1:
//   - endpoints:
//     - fromRegistry: registry1 #must match kubeconfig name in Kubernetes secret
//     - fromCidr: 192.168.100.0/22 #a VM network for example
//     gateways:
//     - registryServiceName: istio-ingressgateway.istio-system.svc.cluster.local
//       port: 15443
//       locality: us-east-1a
//     - address: 192.168.100.1
//       port: 15443
//       locality: us-east-1a
// ```
//
type MeshNetworks struct {
	// The set of networks inside this mesh. Each network should
	// have a unique name and information about how to infer the endpoints in
	// the network as well as the gateways associated with the network.
	Networks             map[string]*Network `protobuf:"bytes,1,rep,name=networks,proto3" json:"networks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MeshNetworks) Reset()         { *m = MeshNetworks{} }
func (m *MeshNetworks) String() string { return proto.CompactTextString(m) }
func (*MeshNetworks) ProtoMessage()    {}
func (*MeshNetworks) Descriptor() ([]byte, []int) {
	return fileDescriptor_a15df2a96e10cd86, []int{1}
}

func (m *MeshNetworks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshNetworks.Unmarshal(m, b)
}
func (m *MeshNetworks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshNetworks.Marshal(b, m, deterministic)
}
func (m *MeshNetworks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshNetworks.Merge(m, src)
}
func (m *MeshNetworks) XXX_Size() int {
	return xxx_messageInfo_MeshNetworks.Size(m)
}
func (m *MeshNetworks) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshNetworks.DiscardUnknown(m)
}

var xxx_messageInfo_MeshNetworks proto.InternalMessageInfo

func (m *MeshNetworks) GetNetworks() map[string]*Network {
	if m != nil {
		return m.Networks
	}
	return nil
}

func init() {
	proto.RegisterType((*Network)(nil), "istio.mesh.v1alpha1.Network")
	proto.RegisterType((*Network_NetworkEndpoints)(nil), "istio.mesh.v1alpha1.Network.NetworkEndpoints")
	proto.RegisterType((*Network_IstioNetworkGateway)(nil), "istio.mesh.v1alpha1.Network.IstioNetworkGateway")
	proto.RegisterType((*MeshNetworks)(nil), "istio.mesh.v1alpha1.MeshNetworks")
	proto.RegisterMapType((map[string]*Network)(nil), "istio.mesh.v1alpha1.MeshNetworks.NetworksEntry")
}

func init() { proto.RegisterFile("mesh/v1alpha1/network.proto", fileDescriptor_a15df2a96e10cd86) }

var fileDescriptor_a15df2a96e10cd86 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcf, 0x6e, 0xd3, 0x30,
	0x18, 0x27, 0x49, 0x61, 0xe9, 0xb7, 0x55, 0x9a, 0x5c, 0x21, 0xa2, 0x30, 0x44, 0x35, 0x09, 0xa9,
	0x17, 0x12, 0x56, 0x38, 0x20, 0x6e, 0x14, 0x4d, 0xc0, 0x81, 0x69, 0x84, 0x13, 0x1c, 0x88, 0xbc,
	0xe6, 0x5b, 0x62, 0x2d, 0x89, 0x23, 0xdb, 0xa4, 0xca, 0xeb, 0xf0, 0x0a, 0x1c, 0x79, 0x19, 0x1e,
	0x05, 0x39, 0x8e, 0x03, 0x9b, 0xaa, 0x9e, 0x5a, 0xff, 0xfe, 0x7d, 0x7f, 0xf2, 0xc1, 0xe3, 0x0a,
	0x65, 0x11, 0xb7, 0x67, 0xb4, 0x6c, 0x0a, 0x7a, 0x16, 0xd7, 0xa8, 0xb6, 0x5c, 0xdc, 0x44, 0x8d,
	0xe0, 0x8a, 0x93, 0x39, 0x93, 0x8a, 0xf1, 0x48, 0x4b, 0x22, 0x2b, 0x09, 0x9f, 0xe6, 0x9c, 0xe7,
	0x25, 0xc6, 0xb4, 0x61, 0xf1, 0x35, 0xc3, 0x32, 0x4b, 0xaf, 0xb0, 0xa0, 0x2d, 0xe3, 0xc2, 0xb8,
	0x4e, 0x7f, 0x79, 0x70, 0x70, 0x61, 0x72, 0xc8, 0x25, 0x4c, 0xb1, 0xce, 0x1a, 0xce, 0x6a, 0x25,
	0x03, 0x77, 0xe1, 0x2d, 0x0f, 0x57, 0xcf, 0xa3, 0x1d, 0xa9, 0xd1, 0x60, 0xb0, 0xbf, 0xe7, 0xd6,
	0xb4, 0xf6, 0xfe, 0xbc, 0x75, 0x93, 0x7f, 0x21, 0xe4, 0x33, 0xf8, 0x39, 0x55, 0xb8, 0xa5, 0x9d,
	0x0c, 0xbc, 0x3e, 0xf0, 0xc5, 0xde, 0xc0, 0x8f, 0x9a, 0x1b, 0x1e, 0xef, 0x8d, 0xd1, 0x64, 0x8e,
	0x31, 0xe1, 0x77, 0x38, 0xbe, 0x5b, 0x96, 0x3c, 0x81, 0xe9, 0xb5, 0xe0, 0x55, 0xba, 0x61, 0x99,
	0x08, 0x9c, 0x85, 0xb3, 0x9c, 0x7e, 0xb8, 0x97, 0xf8, 0x1a, 0x7a, 0xc7, 0x32, 0x41, 0x9e, 0xc1,
	0xac, 0xa7, 0x05, 0xe6, 0x4c, 0x2a, 0xd1, 0x05, 0xee, 0x20, 0x39, 0xd2, 0x70, 0x32, 0xa0, 0xeb,
	0x09, 0xb8, 0x35, 0x86, 0x3f, 0x1d, 0x98, 0xef, 0x68, 0x83, 0xbc, 0x82, 0x87, 0xd6, 0x9f, 0x4a,
	0x14, 0x2d, 0xdb, 0x60, 0x5a, 0xd3, 0x0a, 0xc7, 0x7a, 0x73, 0x4b, 0x7f, 0x31, 0xec, 0x05, 0xad,
	0x90, 0x84, 0x70, 0x40, 0xb3, 0x4c, 0xa0, 0x94, 0x63, 0x51, 0x0b, 0x90, 0x47, 0x30, 0x69, 0xb8,
	0x50, 0x81, 0xb7, 0x70, 0x96, 0x33, 0x33, 0x66, 0x0f, 0x90, 0x10, 0xfc, 0x92, 0x6f, 0x68, 0xc9,
	0x54, 0x17, 0x4c, 0xb4, 0x2b, 0x19, 0xdf, 0xba, 0xc9, 0x7c, 0x7b, 0xfa, 0xdb, 0x81, 0xa3, 0x4f,
	0x28, 0x8b, 0xa1, 0x47, 0x49, 0x2e, 0xc1, 0x1f, 0xae, 0x41, 0x06, 0x4e, 0xbf, 0xe8, 0x78, 0xe7,
	0xa2, 0xff, 0x37, 0xd9, 0xad, 0xcb, 0xf3, 0x5a, 0x8f, 0x6f, 0xf6, 0x6c, 0x53, 0xc2, 0xaf, 0x30,
	0xbb, 0xc5, 0x93, 0x63, 0xf0, 0x6e, 0xb0, 0x33, 0xe3, 0x26, 0xfa, 0x2f, 0x59, 0xc1, 0xfd, 0x96,
	0x96, 0x3f, 0xb0, 0x1f, 0xed, 0x70, 0x75, 0xb2, 0xef, 0xd3, 0x26, 0x46, 0xfa, 0xc6, 0x7d, 0xed,
	0xac, 0x4f, 0xbe, 0x85, 0x46, 0xc9, 0x78, 0x7f, 0x98, 0xb7, 0xae, 0xfa, 0xea, 0x41, 0x7f, 0x98,
	0x2f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x6a, 0x90, 0x3b, 0xed, 0x02, 0x00, 0x00,
}
