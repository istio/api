// Copyright 2018 Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: mesh/v1alpha1/network.proto

package v1alpha1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Network provides information about the endpoints in a routable L3
// network. A single routable L3 network can have one or more service
// registries. Note that the network has no relation to the locality of the
// endpoint. The endpoint locality will be obtained from the service
// registry.
type Network struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of endpoints in the network (obtained through the
	// constituent service registries or from CIDR ranges). All endpoints in
	// the network are directly accessible to one another.
	Endpoints []*Network_NetworkEndpoints `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	// Set of gateways associated with the network.
	Gateways []*Network_IstioNetworkGateway `protobuf:"bytes,3,rep,name=gateways,proto3" json:"gateways,omitempty"`
}

func (x *Network) Reset() {
	*x = Network{}
	mi := &file_mesh_v1alpha1_network_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Network) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Network) ProtoMessage() {}

func (x *Network) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_network_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Network.ProtoReflect.Descriptor instead.
func (*Network) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_network_proto_rawDescGZIP(), []int{0}
}

func (x *Network) GetEndpoints() []*Network_NetworkEndpoints {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *Network) GetGateways() []*Network_IstioNetworkGateway {
	if x != nil {
		return x.Gateways
	}
	return nil
}

// MeshNetworks (config map) provides information about the set of networks
// inside a mesh and how to route to endpoints in each network. For example
//
// MeshNetworks(file/config map):
//
// ```yaml
// networks:
//
//	network1:
//	  endpoints:
//	  - fromRegistry: registry1 #must match kubeconfig name in Kubernetes secret
//	  - fromCidr: 192.168.100.0/22 #a VM network for example
//	  gateways:
//	  - registryServiceName: istio-ingressgateway.istio-system.svc.cluster.local
//	    port: 15443
//	    locality: us-east-1a
//	  - address: 192.168.100.1
//	    port: 15443
//	    locality: us-east-1a
//
// ```
type MeshNetworks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The set of networks inside this mesh. Each network should
	// have a unique name and information about how to infer the endpoints in
	// the network as well as the gateways associated with the network.
	Networks map[string]*Network `protobuf:"bytes,1,rep,name=networks,proto3" json:"networks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MeshNetworks) Reset() {
	*x = MeshNetworks{}
	mi := &file_mesh_v1alpha1_network_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MeshNetworks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshNetworks) ProtoMessage() {}

func (x *MeshNetworks) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_network_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshNetworks.ProtoReflect.Descriptor instead.
func (*MeshNetworks) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_network_proto_rawDescGZIP(), []int{1}
}

func (x *MeshNetworks) GetNetworks() map[string]*Network {
	if x != nil {
		return x.Networks
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
//	a. By matching the registry name with one of the "fromRegistry"
//	in the mesh config. A "fromRegistry" can only be assigned to a
//	single network.
//
//	b. By matching the IP against one of the CIDR ranges in a mesh
//	config network. The CIDR ranges must not overlap and be assigned to
//	a single network.
//
// (2) will override (1) if both are present.
type Network_NetworkEndpoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Ne:
	//
	//	*Network_NetworkEndpoints_FromCidr
	//	*Network_NetworkEndpoints_FromRegistry
	Ne isNetwork_NetworkEndpoints_Ne `protobuf_oneof:"ne"`
}

func (x *Network_NetworkEndpoints) Reset() {
	*x = Network_NetworkEndpoints{}
	mi := &file_mesh_v1alpha1_network_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Network_NetworkEndpoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Network_NetworkEndpoints) ProtoMessage() {}

func (x *Network_NetworkEndpoints) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_network_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Network_NetworkEndpoints.ProtoReflect.Descriptor instead.
func (*Network_NetworkEndpoints) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_network_proto_rawDescGZIP(), []int{0, 0}
}

func (m *Network_NetworkEndpoints) GetNe() isNetwork_NetworkEndpoints_Ne {
	if m != nil {
		return m.Ne
	}
	return nil
}

func (x *Network_NetworkEndpoints) GetFromCidr() string {
	if x, ok := x.GetNe().(*Network_NetworkEndpoints_FromCidr); ok {
		return x.FromCidr
	}
	return ""
}

func (x *Network_NetworkEndpoints) GetFromRegistry() string {
	if x, ok := x.GetNe().(*Network_NetworkEndpoints_FromRegistry); ok {
		return x.FromRegistry
	}
	return ""
}

type isNetwork_NetworkEndpoints_Ne interface {
	isNetwork_NetworkEndpoints_Ne()
}

type Network_NetworkEndpoints_FromCidr struct {
	// A CIDR range for the set of endpoints in this network. The CIDR
	// ranges for endpoints from different networks must not overlap.
	FromCidr string `protobuf:"bytes,1,opt,name=from_cidr,json=fromCidr,proto3,oneof"`
}

type Network_NetworkEndpoints_FromRegistry struct {
	// Add all endpoints from the specified registry into this network.
	// The names of the registries should correspond to the kubeconfig file name
	// inside the secret that was used to configure the registry (Kubernetes
	// multicluster) or supplied by MCP server.
	FromRegistry string `protobuf:"bytes,2,opt,name=from_registry,json=fromRegistry,proto3,oneof"`
}

func (*Network_NetworkEndpoints_FromCidr) isNetwork_NetworkEndpoints_Ne() {}

func (*Network_NetworkEndpoints_FromRegistry) isNetwork_NetworkEndpoints_Ne() {}

// The gateway associated with this network. Traffic from remote networks
// will arrive at the specified gateway:port. All incoming traffic must
// use mTLS.
type Network_IstioNetworkGateway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Gw:
	//
	//	*Network_IstioNetworkGateway_RegistryServiceName
	//	*Network_IstioNetworkGateway_Address
	Gw isNetwork_IstioNetworkGateway_Gw `protobuf_oneof:"gw"`
	// The port associated with the gateway.
	Port uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	// The locality associated with an explicitly specified gateway (i.e. ip)
	Locality string `protobuf:"bytes,4,opt,name=locality,proto3" json:"locality,omitempty"`
}

func (x *Network_IstioNetworkGateway) Reset() {
	*x = Network_IstioNetworkGateway{}
	mi := &file_mesh_v1alpha1_network_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Network_IstioNetworkGateway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Network_IstioNetworkGateway) ProtoMessage() {}

func (x *Network_IstioNetworkGateway) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_network_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Network_IstioNetworkGateway.ProtoReflect.Descriptor instead.
func (*Network_IstioNetworkGateway) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_network_proto_rawDescGZIP(), []int{0, 1}
}

func (m *Network_IstioNetworkGateway) GetGw() isNetwork_IstioNetworkGateway_Gw {
	if m != nil {
		return m.Gw
	}
	return nil
}

func (x *Network_IstioNetworkGateway) GetRegistryServiceName() string {
	if x, ok := x.GetGw().(*Network_IstioNetworkGateway_RegistryServiceName); ok {
		return x.RegistryServiceName
	}
	return ""
}

func (x *Network_IstioNetworkGateway) GetAddress() string {
	if x, ok := x.GetGw().(*Network_IstioNetworkGateway_Address); ok {
		return x.Address
	}
	return ""
}

func (x *Network_IstioNetworkGateway) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Network_IstioNetworkGateway) GetLocality() string {
	if x != nil {
		return x.Locality
	}
	return ""
}

type isNetwork_IstioNetworkGateway_Gw interface {
	isNetwork_IstioNetworkGateway_Gw()
}

type Network_IstioNetworkGateway_RegistryServiceName struct {
	// A fully qualified domain name of the gateway service.  istiod will
	// lookup the service from the service registries in the network and
	// obtain the endpoint IPs of the gateway from the service
	// registry. Note that while the service name is a fully qualified
	// domain name, it need not be resolvable outside the orchestration
	// platform for the registry. e.g., this could be
	// istio-ingressgateway.istio-system.svc.cluster.local.
	RegistryServiceName string `protobuf:"bytes,1,opt,name=registry_service_name,json=registryServiceName,proto3,oneof"`
}

type Network_IstioNetworkGateway_Address struct {
	// IP address or externally resolvable DNS address associated with the gateway.
	Address string `protobuf:"bytes,2,opt,name=address,proto3,oneof"`
}

func (*Network_IstioNetworkGateway_RegistryServiceName) isNetwork_IstioNetworkGateway_Gw() {}

func (*Network_IstioNetworkGateway_Address) isNetwork_IstioNetworkGateway_Gw() {}

var File_mesh_v1alpha1_network_proto protoreflect.FileDescriptor

var file_mesh_v1alpha1_network_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x69,
	0x73, 0x74, 0x69, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x03, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x51, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x12, 0x52, 0x0a, 0x08, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x08, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x1a, 0x5e, 0x0a, 0x10, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x09, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x08, 0x66, 0x72, 0x6f, 0x6d, 0x43, 0x69, 0x64, 0x72, 0x12, 0x25, 0x0a, 0x0d, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x42, 0x04, 0x0a, 0x02, 0x6e, 0x65, 0x1a, 0xa3, 0x01, 0x0a, 0x13, 0x49, 0x73, 0x74, 0x69, 0x6f,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x34,
	0x0a, 0x15, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x13, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x04,
	0xe2, 0x41, 0x01, 0x02, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x04, 0x0a, 0x02, 0x67, 0x77, 0x22, 0xbc, 0x01, 0x0a,
	0x0c, 0x4d, 0x65, 0x73, 0x68, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x51, 0x0a,
	0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x1a, 0x59, 0x0a, 0x0d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x1c, 0x5a, 0x1a, 0x69,
	0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_mesh_v1alpha1_network_proto_rawDescOnce sync.Once
	file_mesh_v1alpha1_network_proto_rawDescData = file_mesh_v1alpha1_network_proto_rawDesc
)

func file_mesh_v1alpha1_network_proto_rawDescGZIP() []byte {
	file_mesh_v1alpha1_network_proto_rawDescOnce.Do(func() {
		file_mesh_v1alpha1_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_mesh_v1alpha1_network_proto_rawDescData)
	})
	return file_mesh_v1alpha1_network_proto_rawDescData
}

var file_mesh_v1alpha1_network_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_mesh_v1alpha1_network_proto_goTypes = []any{
	(*Network)(nil),                     // 0: istio.mesh.v1alpha1.Network
	(*MeshNetworks)(nil),                // 1: istio.mesh.v1alpha1.MeshNetworks
	(*Network_NetworkEndpoints)(nil),    // 2: istio.mesh.v1alpha1.Network.NetworkEndpoints
	(*Network_IstioNetworkGateway)(nil), // 3: istio.mesh.v1alpha1.Network.IstioNetworkGateway
	nil,                                 // 4: istio.mesh.v1alpha1.MeshNetworks.NetworksEntry
}
var file_mesh_v1alpha1_network_proto_depIdxs = []int32{
	2, // 0: istio.mesh.v1alpha1.Network.endpoints:type_name -> istio.mesh.v1alpha1.Network.NetworkEndpoints
	3, // 1: istio.mesh.v1alpha1.Network.gateways:type_name -> istio.mesh.v1alpha1.Network.IstioNetworkGateway
	4, // 2: istio.mesh.v1alpha1.MeshNetworks.networks:type_name -> istio.mesh.v1alpha1.MeshNetworks.NetworksEntry
	0, // 3: istio.mesh.v1alpha1.MeshNetworks.NetworksEntry.value:type_name -> istio.mesh.v1alpha1.Network
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_mesh_v1alpha1_network_proto_init() }
func file_mesh_v1alpha1_network_proto_init() {
	if File_mesh_v1alpha1_network_proto != nil {
		return
	}
	file_mesh_v1alpha1_network_proto_msgTypes[2].OneofWrappers = []any{
		(*Network_NetworkEndpoints_FromCidr)(nil),
		(*Network_NetworkEndpoints_FromRegistry)(nil),
	}
	file_mesh_v1alpha1_network_proto_msgTypes[3].OneofWrappers = []any{
		(*Network_IstioNetworkGateway_RegistryServiceName)(nil),
		(*Network_IstioNetworkGateway_Address)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mesh_v1alpha1_network_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mesh_v1alpha1_network_proto_goTypes,
		DependencyIndexes: file_mesh_v1alpha1_network_proto_depIdxs,
		MessageInfos:      file_mesh_v1alpha1_network_proto_msgTypes,
	}.Build()
	File_mesh_v1alpha1_network_proto = out.File
	file_mesh_v1alpha1_network_proto_rawDesc = nil
	file_mesh_v1alpha1_network_proto_goTypes = nil
	file_mesh_v1alpha1_network_proto_depIdxs = nil
}
