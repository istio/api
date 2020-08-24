// Code generated by protoc-gen-go. DO NOT EDIT.
// source: networking/v1alpha3/workload_entry.proto

// `WorkloadEntry` enables operators to describe the properties of a
// single non-Kubernetes workload such as a VM or a bare metal server
// as it is onboarded into the mesh. A `WorkloadEntry` must be
// accompanied by an Istio `ServiceEntry` that selects the workload
// through the appropriate labels and provides the service definition
// for a `MESH_INTERNAL` service (hostnames, port properties, etc.). A
// `ServiceEntry` object can select multiple workload entries as well
// as Kubernetes pods based on the label selector specified in the
// service entry.
//
// When a workload connects to `istiod`, the status field in the
// custom resource will be updated to indicate the health of the
// workload along with other details, similar to how Kubernetes
// updates the status of a pod.
//
// The following example declares a workload entry representing a VM
// for the `details.bookinfo.com` service. This VM has sidecar
// installed and bootstrapped using the `details-legacy` service
// account. The service is exposed on port 80 to applications in the
// mesh. The HTTP traffic to this service is wrapped in Istio mutual
// TLS and sent to sidecars on VMs on target port 8080, that in turn
// forward it to the application on localhost on the same port.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: WorkloadEntry
// metadata:
//   name: details-svc
// spec:
//   # use of the service account indicates that the workload has a
//   # sidecar proxy bootstrapped with this service account. Pods with
//   # sidecars will automatically communicate with the workload using
//   # istio mutual TLS.
//   serviceAccount: details-legacy
//   address: 2.2.2.2
//   labels:
//     app: details-legacy
//     instance-id: vm1
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: WorkloadEntry
// metadata:
//   name: details-svc
// spec:
//   # use of the service account indicates that the workload has a
//   # sidecar proxy bootstrapped with this service account. Pods with
//   # sidecars will automatically communicate with the workload using
//   # istio mutual TLS.
//   serviceAccount: details-legacy
//   address: 2.2.2.2
//   labels:
//     app: details-legacy
//     instance-id: vm1
// ```
// {{</tab>}}
// {{</tabset>}}
//
// and the associated service entry
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: details-svc
// spec:
//   hosts:
//   - details.bookinfo.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//     targetPort: 8080
//   resolution: STATIC
//   workloadSelector:
//     labels:
//       app: details-legacy
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: details-svc
// spec:
//   hosts:
//   - details.bookinfo.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//     targetPort: 8080
//   resolution: STATIC
//   workloadSelector:
//     labels:
//       app: details-legacy
// ```
// {{</tab>}}
// {{</tabset>}}
//
//
// The following example declares the same VM workload using
// its fully qualified DNS name. The service entry's resolution
// mode should be changed to DNS to indicate that the client-side
// sidecars should dynamically resolve the DNS name at runtime before
// forwarding the request.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: WorkloadEntry
// metadata:
//   name: details-svc
// spec:
//   # use of the service account indicates that the workload has a
//   # sidecar proxy bootstrapped with this service account. Pods with
//   # sidecars will automatically communicate with the workload using
//   # istio mutual TLS.
//   serviceAccount: details-legacy
//   address: vm1.vpc01.corp.net
//   labels:
//     app: details-legacy
//     instance-id: vm1
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: WorkloadEntry
// metadata:
//   name: details-svc
// spec:
//   # use of the service account indicates that the workload has a
//   # sidecar proxy bootstrapped with this service account. Pods with
//   # sidecars will automatically communicate with the workload using
//   # istio mutual TLS.
//   serviceAccount: details-legacy
//   address: vm1.vpc01.corp.net
//   labels:
//     app: details-legacy
//     instance-id: vm1
// ```
// {{</tab>}}
// {{</tabset>}}
//
// and the associated service entry
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: details-svc
// spec:
//   hosts:
//   - details.bookinfo.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//     targetPort: 8080
//   resolution: DNS
//   workloadSelector:
//     labels:
//       app: details-legacy
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: details-svc
// spec:
//   hosts:
//   - details.bookinfo.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//     targetPort: 8080
//   resolution: DNS
//   workloadSelector:
//     labels:
//       app: details-legacy
// ```
// {{</tab>}}
// {{</tabset>}}
//

package v1alpha3

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

// WorkloadEntry enables specifying the properties of a single non-Kubernetes workload such a VM or a bare metal services that can be referred to by service entries.
//
// <!-- crd generation tags
// +cue-gen:WorkloadEntry:groupName:networking.istio.io
// +cue-gen:WorkloadEntry:version:v1alpha3
// +cue-gen:WorkloadEntry:storageVersion
// +cue-gen:WorkloadEntry:annotations:helm.sh/resource-policy=keep
// +cue-gen:WorkloadEntry:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:WorkloadEntry:subresource:status
// +cue-gen:WorkloadEntry:scope:Namespaced
// +cue-gen:WorkloadEntry:resource:categories=istio-io,networking-istio-io,shortNames=we,plural=workloadentries
// +cue-gen:WorkloadEntry:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// +cue-gen:WorkloadEntry:printerColumn:name=Address,type=string,JSONPath=.spec.address,description="Address associated with the network endpoint."
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.istio.io/v1alpha3
// +genclient
// +k8s:deepcopy-gen=true
// -->
type WorkloadEntry struct {
	// Address associated with the network endpoint without the
	// port.  Domain names can be used if and only if the resolution is set
	// to DNS, and must be fully-qualified without wildcards. Use the form
	// unix:///absolute/path/to/socket for Unix domain socket endpoints.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Set of ports associated with the endpoint. If the port map is
	// specified, it must be a map of servicePortName to this endpoint's
	// port, such that traffic to the service port will be forwarded to
	// the endpoint port that maps to the service's portName. If
	// omitted, and the targetPort is specified as part of the service's
	// port specification, traffic to the service port will be forwarded
	// to one of the endpoints on the specified `targetPort`. If both
	// the targetPort and endpoint's port map are not specified, traffic
	// to a service port will be forwarded to one of the endpoints on
	// the same port.
	//
	// **NOTE 1:** Do not use for `unix://` addresses.
	//
	// **NOTE 2:** endpoint port map takes precedence over targetPort.
	Ports map[string]uint32 `protobuf:"bytes,2,rep,name=ports,proto3" json:"ports,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// One or more labels associated with the endpoint.
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Network enables Istio to group endpoints resident in the same L3
	// domain/network. All endpoints in the same network are assumed to be
	// directly reachable from one another. When endpoints in different
	// networks cannot reach each other directly, an Istio Gateway can be
	// used to establish connectivity (usually using the
	// `AUTO_PASSTHROUGH` mode in a Gateway Server). This is
	// an advanced configuration used typically for spanning an Istio mesh
	// over multiple clusters.
	Network string `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
	// The locality associated with the endpoint. A locality corresponds
	// to a failure domain (e.g., country/region/zone). Arbitrary failure
	// domain hierarchies can be represented by separating each
	// encapsulating failure domain by /. For example, the locality of an
	// an endpoint in US, in US-East-1 region, within availability zone
	// az-1, in data center rack r11 can be represented as
	// us/us-east-1/az-1/r11. Istio will configure the sidecar to route to
	// endpoints within the same locality as the sidecar. If none of the
	// endpoints in the locality are available, endpoints parent locality
	// (but within the same network ID) will be chosen. For example, if
	// there are two endpoints in same network (networkID "n1"), say e1
	// with locality us/us-east-1/az-1/r11 and e2 with locality
	// us/us-east-1/az-2/r12, a sidecar from us/us-east-1/az-1/r11 locality
	// will prefer e1 from the same locality over e2 from a different
	// locality. Endpoint e2 could be the IP associated with a gateway
	// (that bridges networks n1 and n2), or the IP associated with a
	// standard service endpoint.
	Locality string `protobuf:"bytes,5,opt,name=locality,proto3" json:"locality,omitempty"`
	// The load balancing weight associated with the endpoint. Endpoints
	// with higher weights will receive proportionally higher traffic.
	Weight uint32 `protobuf:"varint,6,opt,name=weight,proto3" json:"weight,omitempty"`
	// The service account associated with the workload if a sidecar
	// is present in the workload. The service account must be present
	// in the same namespace as the configuration ( WorkloadEntry or a
	// ServiceEntry)
	ServiceAccount       string   `protobuf:"bytes,7,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkloadEntry) Reset()         { *m = WorkloadEntry{} }
func (m *WorkloadEntry) String() string { return proto.CompactTextString(m) }
func (*WorkloadEntry) ProtoMessage()    {}
func (*WorkloadEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d7791a8c20f14f2, []int{0}
}

func (m *WorkloadEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkloadEntry.Unmarshal(m, b)
}
func (m *WorkloadEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkloadEntry.Marshal(b, m, deterministic)
}
func (m *WorkloadEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadEntry.Merge(m, src)
}
func (m *WorkloadEntry) XXX_Size() int {
	return xxx_messageInfo_WorkloadEntry.Size(m)
}
func (m *WorkloadEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadEntry.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadEntry proto.InternalMessageInfo

func (m *WorkloadEntry) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *WorkloadEntry) GetPorts() map[string]uint32 {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *WorkloadEntry) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *WorkloadEntry) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *WorkloadEntry) GetLocality() string {
	if m != nil {
		return m.Locality
	}
	return ""
}

func (m *WorkloadEntry) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *WorkloadEntry) GetServiceAccount() string {
	if m != nil {
		return m.ServiceAccount
	}
	return ""
}

func init() {
	proto.RegisterType((*WorkloadEntry)(nil), "istio.networking.v1alpha3.WorkloadEntry")
	proto.RegisterMapType((map[string]string)(nil), "istio.networking.v1alpha3.WorkloadEntry.LabelsEntry")
	proto.RegisterMapType((map[string]uint32)(nil), "istio.networking.v1alpha3.WorkloadEntry.PortsEntry")
}

func init() {
	proto.RegisterFile("networking/v1alpha3/workload_entry.proto", fileDescriptor_6d7791a8c20f14f2)
}

var fileDescriptor_6d7791a8c20f14f2 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0x49, 0xf2, 0x6f, 0xfa, 0xef, 0x94, 0xaa, 0x2c, 0x22, 0x6b, 0x40, 0x0c, 0xbd, 0x98,
	0x53, 0x82, 0xd6, 0x43, 0xf5, 0xd6, 0x82, 0x07, 0xa1, 0x07, 0xc9, 0x45, 0xf0, 0x52, 0xb6, 0xc9,
	0x9a, 0x2e, 0x5d, 0xb2, 0x61, 0x77, 0x9b, 0xd2, 0xe7, 0xf2, 0x85, 0x7c, 0x14, 0xc9, 0x6e, 0x6a,
	0x2b, 0x28, 0x7a, 0x9b, 0xef, 0xcb, 0x7c, 0xbf, 0x64, 0x66, 0x02, 0x51, 0x49, 0xf5, 0x46, 0xc8,
	0x15, 0x2b, 0x8b, 0xa4, 0xbe, 0x26, 0xbc, 0x5a, 0x92, 0x51, 0xd2, 0x18, 0x5c, 0x90, 0x7c, 0x4e,
	0x4b, 0x2d, 0xb7, 0x71, 0x25, 0x85, 0x16, 0xe8, 0x9c, 0x29, 0xcd, 0x44, 0xbc, 0xef, 0x8f, 0x77,
	0xfd, 0xc1, 0x65, 0x21, 0x44, 0xc1, 0x69, 0x42, 0x2a, 0x96, 0xbc, 0x32, 0xca, 0xf3, 0xf9, 0x82,
	0x2e, 0x49, 0xcd, 0x84, 0xb4, 0xd9, 0xe1, 0x9b, 0x07, 0x83, 0xe7, 0x16, 0xfa, 0xd0, 0x30, 0xd1,
	0x05, 0x74, 0x49, 0x9e, 0x4b, 0xaa, 0x14, 0x76, 0x42, 0x27, 0xea, 0x4d, 0xbd, 0xf7, 0x89, 0x9b,
	0xee, 0x3c, 0xf4, 0x08, 0x9d, 0x4a, 0x48, 0xad, 0xb0, 0x1b, 0x7a, 0x51, 0xff, 0x66, 0x14, 0xff,
	0xf8, 0xf2, 0xf8, 0x0b, 0x37, 0x7e, 0x6a, 0x52, 0xa6, 0x4c, 0x2d, 0x01, 0xcd, 0xc0, 0xe7, 0x64,
	0x41, 0xb9, 0xc2, 0x9e, 0x61, 0xdd, 0xfe, 0x99, 0x35, 0x33, 0x31, 0x0b, 0x6b, 0x19, 0x08, 0x43,
	0xb7, 0x0d, 0xe2, 0x7f, 0xcd, 0x77, 0xa7, 0x3b, 0x89, 0x02, 0xf8, 0xcf, 0x45, 0x46, 0x38, 0xd3,
	0x5b, 0xdc, 0x31, 0x8f, 0x3e, 0x35, 0x3a, 0x03, 0x7f, 0x43, 0x59, 0xb1, 0xd4, 0xd8, 0x0f, 0x9d,
	0x68, 0x90, 0xb6, 0x0a, 0x5d, 0xc1, 0xb1, 0xa2, 0xb2, 0x66, 0x19, 0x9d, 0x93, 0x2c, 0x13, 0xeb,
	0x52, 0xe3, 0xae, 0x89, 0x1e, 0xb5, 0xf6, 0xc4, 0xba, 0xc1, 0x18, 0x60, 0x3f, 0x19, 0x3a, 0x01,
	0x6f, 0x45, 0xb7, 0x76, 0x71, 0x69, 0x53, 0xa2, 0x53, 0xe8, 0xd4, 0x84, 0xaf, 0x29, 0x76, 0x0d,
	0xdf, 0x8a, 0x7b, 0x77, 0xec, 0x04, 0x77, 0xd0, 0x3f, 0x98, 0xe3, 0xb7, 0x68, 0xef, 0x20, 0x3a,
	0x1d, 0xbe, 0x84, 0x76, 0x55, 0x4c, 0x98, 0xd3, 0x7e, 0xf3, 0xab, 0x2c, 0x7c, 0x73, 0xe0, 0xd1,
	0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0x04, 0x2e, 0x50, 0x48, 0x02, 0x00, 0x00,
}
