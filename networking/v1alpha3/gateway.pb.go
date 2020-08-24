// Code generated by protoc-gen-go. DO NOT EDIT.
// source: networking/v1alpha3/gateway.proto

// `Gateway` describes a load balancer operating at the edge of the mesh
// receiving incoming or outgoing HTTP/TCP connections. The specification
// describes a set of ports that should be exposed, the type of protocol to
// use, SNI configuration for the load balancer, etc.
//
// For example, the following Gateway configuration sets up a proxy to act
// as a load balancer exposing port 80 and 9080 (http), 443 (https),
// 9443(https) and port 2379 (TCP) for ingress.  The gateway will be
// applied to the proxy running on a pod with labels `app:
// my-gateway-controller`. While Istio will configure the proxy to listen
// on these ports, it is the responsibility of the user to ensure that
// external traffic to these ports are allowed into the mesh.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: Gateway
// metadata:
//   name: my-gateway
//   namespace: some-config-namespace
// spec:
//   selector:
//     app: my-gateway-controller
//   servers:
//   - port:
//       number: 80
//       name: http
//       protocol: HTTP
//     hosts:
//     - uk.bookinfo.com
//     - eu.bookinfo.com
//     tls:
//       httpsRedirect: true # sends 301 redirect for http requests
//   - port:
//       number: 443
//       name: https-443
//       protocol: HTTPS
//     hosts:
//     - uk.bookinfo.com
//     - eu.bookinfo.com
//     tls:
//       mode: SIMPLE # enables HTTPS on this port
//       serverCertificate: /etc/certs/servercert.pem
//       privateKey: /etc/certs/privatekey.pem
//   - port:
//       number: 9443
//       name: https-9443
//       protocol: HTTPS
//     hosts:
//     - "bookinfo-namespace/*.bookinfo.com"
//     tls:
//       mode: SIMPLE # enables HTTPS on this port
//       credentialName: bookinfo-secret # fetches certs from Kubernetes secret
//   - port:
//       number: 9080
//       name: http-wildcard
//       protocol: HTTP
//     hosts:
//     - "*"
//   - port:
//       number: 2379 # to expose internal service via external port 2379
//       name: mongo
//       protocol: MONGO
//     hosts:
//     - "*"
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Gateway
// metadata:
//   name: my-gateway
//   namespace: some-config-namespace
// spec:
//   selector:
//     app: my-gateway-controller
//   servers:
//   - port:
//       number: 80
//       name: http
//       protocol: HTTP
//     hosts:
//     - uk.bookinfo.com
//     - eu.bookinfo.com
//     tls:
//       httpsRedirect: true # sends 301 redirect for http requests
//   - port:
//       number: 443
//       name: https-443
//       protocol: HTTPS
//     hosts:
//     - uk.bookinfo.com
//     - eu.bookinfo.com
//     tls:
//       mode: SIMPLE # enables HTTPS on this port
//       serverCertificate: /etc/certs/servercert.pem
//       privateKey: /etc/certs/privatekey.pem
//   - port:
//       number: 9443
//       name: https-9443
//       protocol: HTTPS
//     hosts:
//     - "bookinfo-namespace/*.bookinfo.com"
//     tls:
//       mode: SIMPLE # enables HTTPS on this port
//       credentialName: bookinfo-secret # fetches certs from Kubernetes secret
//   - port:
//       number: 9080
//       name: http-wildcard
//       protocol: HTTP
//     hosts:
//     - "*"
//   - port:
//       number: 2379 # to expose internal service via external port 2379
//       name: mongo
//       protocol: MONGO
//     hosts:
//     - "*"
// ```
// {{</tab>}}
// {{</tabset>}}
//
// The Gateway specification above describes the L4-L6 properties of a load
// balancer. A `VirtualService` can then be bound to a gateway to control
// the forwarding of traffic arriving at a particular host or gateway port.
//
// For example, the following VirtualService splits traffic for
// `https://uk.bookinfo.com/reviews`, `https://eu.bookinfo.com/reviews`,
// `http://uk.bookinfo.com:9080/reviews`,
// `http://eu.bookinfo.com:9080/reviews` into two versions (prod and qa) of
// an internal reviews service on port 9080. In addition, requests
// containing the cookie "user: dev-123" will be sent to special port 7777
// in the qa version. The same rule is also applicable inside the mesh for
// requests to the "reviews.prod.svc.cluster.local" service. This rule is
// applicable across ports 443, 9080. Note that `http://uk.bookinfo.com`
// gets redirected to `https://uk.bookinfo.com` (i.e. 80 redirects to 443).
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: VirtualService
// metadata:
//   name: bookinfo-rule
//   namespace: bookinfo-namespace
// spec:
//   hosts:
//   - reviews.prod.svc.cluster.local
//   - uk.bookinfo.com
//   - eu.bookinfo.com
//   gateways:
//   - some-config-namespace/my-gateway
//   - mesh # applies to all the sidecars in the mesh
//   http:
//   - match:
//     - headers:
//         cookie:
//           exact: "user=dev-123"
//     route:
//     - destination:
//         port:
//           number: 7777
//         host: reviews.qa.svc.cluster.local
//   - match:
//     - uri:
//         prefix: /reviews/
//     route:
//     - destination:
//         port:
//           number: 9080 # can be omitted if it's the only port for reviews
//         host: reviews.prod.svc.cluster.local
//       weight: 80
//     - destination:
//         host: reviews.qa.svc.cluster.local
//       weight: 20
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: VirtualService
// metadata:
//   name: bookinfo-rule
//   namespace: bookinfo-namespace
// spec:
//   hosts:
//   - reviews.prod.svc.cluster.local
//   - uk.bookinfo.com
//   - eu.bookinfo.com
//   gateways:
//   - some-config-namespace/my-gateway
//   - mesh # applies to all the sidecars in the mesh
//   http:
//   - match:
//     - headers:
//         cookie:
//           exact: "user=dev-123"
//     route:
//     - destination:
//         port:
//           number: 7777
//         host: reviews.qa.svc.cluster.local
//   - match:
//     - uri:
//         prefix: /reviews/
//     route:
//     - destination:
//         port:
//           number: 9080 # can be omitted if it's the only port for reviews
//         host: reviews.prod.svc.cluster.local
//       weight: 80
//     - destination:
//         host: reviews.qa.svc.cluster.local
//       weight: 20
// ```
// {{</tab>}}
// {{</tabset>}}
//
// The following VirtualService forwards traffic arriving at (external)
// port 27017 to internal Mongo server on port 5555. This rule is not
// applicable internally in the mesh as the gateway list omits the
// reserved name `mesh`.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: VirtualService
// metadata:
//   name: bookinfo-Mongo
//   namespace: bookinfo-namespace
// spec:
//   hosts:
//   - mongosvr.prod.svc.cluster.local # name of internal Mongo service
//   gateways:
//   - some-config-namespace/my-gateway # can omit the namespace if gateway is in same
//                                        namespace as virtual service.
//   tcp:
//   - match:
//     - port: 27017
//     route:
//     - destination:
//         host: mongo.prod.svc.cluster.local
//         port:
//           number: 5555
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: VirtualService
// metadata:
//   name: bookinfo-Mongo
//   namespace: bookinfo-namespace
// spec:
//   hosts:
//   - mongosvr.prod.svc.cluster.local # name of internal Mongo service
//   gateways:
//   - some-config-namespace/my-gateway # can omit the namespace if gateway is in same
//                                        namespace as virtual service.
//   tcp:
//   - match:
//     - port: 27017
//     route:
//     - destination:
//         host: mongo.prod.svc.cluster.local
//         port:
//           number: 5555
// ```
// {{</tab>}}
// {{</tabset>}}
//
// It is possible to restrict the set of virtual services that can bind to
// a gateway server using the namespace/hostname syntax in the hosts field.
// For example, the following Gateway allows any virtual service in the ns1
// namespace to bind to it, while restricting only the virtual service with
// foo.bar.com host in the ns2 namespace to bind to it.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: Gateway
// metadata:
//   name: my-gateway
//   namespace: some-config-namespace
// spec:
//   selector:
//     app: my-gateway-controller
//   servers:
//   - port:
//       number: 80
//       name: http
//       protocol: HTTP
//     hosts:
//     - "ns1/*"
//     - "ns2/foo.bar.com"
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Gateway
// metadata:
//   name: my-gateway
//   namespace: some-config-namespace
// spec:
//   selector:
//     app: my-gateway-controller
//   servers:
//   - port:
//       number: 80
//       name: http
//       protocol: HTTP
//     hosts:
//     - "ns1/*"
//     - "ns2/foo.bar.com"
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

// TLS modes enforced by the proxy
type ServerTLSSettings_TLSmode int32

const (
	// The SNI string presented by the client will be used as the
	// match criterion in a VirtualService TLS route to determine
	// the destination service from the service registry.
	ServerTLSSettings_PASSTHROUGH ServerTLSSettings_TLSmode = 0
	// Secure connections with standard TLS semantics.
	ServerTLSSettings_SIMPLE ServerTLSSettings_TLSmode = 1
	// Secure connections to the downstream using mutual TLS by
	// presenting server certificates for authentication.
	ServerTLSSettings_MUTUAL ServerTLSSettings_TLSmode = 2
	// Similar to the passthrough mode, except servers with this TLS
	// mode do not require an associated VirtualService to map from
	// the SNI value to service in the registry. The destination
	// details such as the service/subset/port are encoded in the
	// SNI value. The proxy will forward to the upstream (Envoy)
	// cluster (a group of endpoints) specified by the SNI
	// value. This server is typically used to provide connectivity
	// between services in disparate L3 networks that otherwise do
	// not have direct connectivity between their respective
	// endpoints. Use of this mode assumes that both the source and
	// the destination are using Istio mTLS to secure traffic.
	ServerTLSSettings_AUTO_PASSTHROUGH ServerTLSSettings_TLSmode = 3
	// Secure connections from the downstream using mutual TLS by
	// presenting server certificates for authentication.  Compared
	// to Mutual mode, this mode uses certificates, representing
	// gateway workload identity, generated automatically by Istio
	// for mTLS authentication. When this mode is used, all other
	// fields in `TLSOptions` should be empty.
	ServerTLSSettings_ISTIO_MUTUAL ServerTLSSettings_TLSmode = 4
)

var ServerTLSSettings_TLSmode_name = map[int32]string{
	0: "PASSTHROUGH",
	1: "SIMPLE",
	2: "MUTUAL",
	3: "AUTO_PASSTHROUGH",
	4: "ISTIO_MUTUAL",
}

var ServerTLSSettings_TLSmode_value = map[string]int32{
	"PASSTHROUGH":      0,
	"SIMPLE":           1,
	"MUTUAL":           2,
	"AUTO_PASSTHROUGH": 3,
	"ISTIO_MUTUAL":     4,
}

func (x ServerTLSSettings_TLSmode) String() string {
	return proto.EnumName(ServerTLSSettings_TLSmode_name, int32(x))
}

func (ServerTLSSettings_TLSmode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_067d98d02f84cc0b, []int{3, 0}
}

// TLS protocol versions.
type ServerTLSSettings_TLSProtocol int32

const (
	// Automatically choose the optimal TLS version.
	ServerTLSSettings_TLS_AUTO ServerTLSSettings_TLSProtocol = 0
	// TLS version 1.0
	ServerTLSSettings_TLSV1_0 ServerTLSSettings_TLSProtocol = 1
	// TLS version 1.1
	ServerTLSSettings_TLSV1_1 ServerTLSSettings_TLSProtocol = 2
	// TLS version 1.2
	ServerTLSSettings_TLSV1_2 ServerTLSSettings_TLSProtocol = 3
	// TLS version 1.3
	ServerTLSSettings_TLSV1_3 ServerTLSSettings_TLSProtocol = 4
)

var ServerTLSSettings_TLSProtocol_name = map[int32]string{
	0: "TLS_AUTO",
	1: "TLSV1_0",
	2: "TLSV1_1",
	3: "TLSV1_2",
	4: "TLSV1_3",
}

var ServerTLSSettings_TLSProtocol_value = map[string]int32{
	"TLS_AUTO": 0,
	"TLSV1_0":  1,
	"TLSV1_1":  2,
	"TLSV1_2":  3,
	"TLSV1_3":  4,
}

func (x ServerTLSSettings_TLSProtocol) String() string {
	return proto.EnumName(ServerTLSSettings_TLSProtocol_name, int32(x))
}

func (ServerTLSSettings_TLSProtocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_067d98d02f84cc0b, []int{3, 1}
}

// Gateway describes a load balancer operating at the edge of the mesh
// receiving incoming or outgoing HTTP/TCP connections.
//
// <!-- crd generation tags
// +cue-gen:Gateway:groupName:networking.istio.io
// +cue-gen:Gateway:version:v1alpha3
// +cue-gen:Gateway:storageVersion
// +cue-gen:Gateway:annotations:helm.sh/resource-policy=keep
// +cue-gen:Gateway:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:Gateway:subresource:status
// +cue-gen:Gateway:scope:Namespaced
// +cue-gen:Gateway:resource:categories=istio-io,networking-istio-io,shortNames=gw
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.istio.io/v1alpha3
// +genclient
// +k8s:deepcopy-gen=true
// -->
type Gateway struct {
	// A list of server specifications.
	Servers []*Server `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
	// One or more labels that indicate a specific set of pods/VMs
	// on which this gateway configuration should be applied.
	// By default workloads are searched across all namespaces based on label selectors.
	// This implies that a gateway resource in the namespace "foo" can select pods in
	// the namespace "bar" based on labels.
	// This behavior can be controlled via the `PILOT_SCOPE_GATEWAY_TO_NAMESPACE`
	// environment variable in istiod. If this variable is set
	// to true, the scope of label search is restricted to the configuration
	// namespace in which the the resource is present. In other words, the Gateway
	// resource must reside in the same namespace as the gateway workload
	// instance.
	// If selector is nil, the Gateway will be applied to all workloads.
	Selector             map[string]string `protobuf:"bytes,2,rep,name=selector,proto3" json:"selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Gateway) Reset()         { *m = Gateway{} }
func (m *Gateway) String() string { return proto.CompactTextString(m) }
func (*Gateway) ProtoMessage()    {}
func (*Gateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_067d98d02f84cc0b, []int{0}
}

func (m *Gateway) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gateway.Unmarshal(m, b)
}
func (m *Gateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gateway.Marshal(b, m, deterministic)
}
func (m *Gateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gateway.Merge(m, src)
}
func (m *Gateway) XXX_Size() int {
	return xxx_messageInfo_Gateway.Size(m)
}
func (m *Gateway) XXX_DiscardUnknown() {
	xxx_messageInfo_Gateway.DiscardUnknown(m)
}

var xxx_messageInfo_Gateway proto.InternalMessageInfo

func (m *Gateway) GetServers() []*Server {
	if m != nil {
		return m.Servers
	}
	return nil
}

func (m *Gateway) GetSelector() map[string]string {
	if m != nil {
		return m.Selector
	}
	return nil
}

// `Server` describes the properties of the proxy on a given load balancer
// port. For example,
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: Gateway
// metadata:
//   name: my-ingress
// spec:
//   selector:
//     app: my-ingressgateway
//   servers:
//   - port:
//       number: 80
//       name: http2
//       protocol: HTTP2
//     hosts:
//     - "*"
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Gateway
// metadata:
//   name: my-ingress
// spec:
//   selector:
//     app: my-ingressgateway
//   servers:
//   - port:
//       number: 80
//       name: http2
//       protocol: HTTP2
//     hosts:
//     - "*"
// ```
// {{</tab>}}
// {{</tabset>}}
//
// Another example
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: Gateway
// metadata:
//   name: my-tcp-ingress
// spec:
//   selector:
//     app: my-tcp-ingressgateway
//   servers:
//   - port:
//       number: 27018
//       name: mongo
//       protocol: MONGO
//     hosts:
//     - "*"
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Gateway
// metadata:
//   name: my-tcp-ingress
// spec:
//   selector:
//     app: my-tcp-ingressgateway
//   servers:
//   - port:
//       number: 27018
//       name: mongo
//       protocol: MONGO
//     hosts:
//     - "*"
// ```
// {{</tab>}}
// {{</tabset>}}
//
// The following is an example of TLS configuration for port 443
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: Gateway
// metadata:
//   name: my-tls-ingress
// spec:
//   selector:
//     app: my-tls-ingressgateway
//   servers:
//   - port:
//       number: 443
//       name: https
//       protocol: HTTPS
//     hosts:
//     - "*"
//     tls:
//       mode: SIMPLE
//       serverCertificate: /etc/certs/server.pem
//       privateKey: /etc/certs/privatekey.pem
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Gateway
// metadata:
//   name: my-tls-ingress
// spec:
//   selector:
//     app: my-tls-ingressgateway
//   servers:
//   - port:
//       number: 443
//       name: https
//       protocol: HTTPS
//     hosts:
//     - "*"
//     tls:
//       mode: SIMPLE
//       serverCertificate: /etc/certs/server.pem
//       privateKey: /etc/certs/privatekey.pem
// ```
// {{</tab>}}
// {{</tabset>}}
//
type Server struct {
	// The Port on which the proxy should listen for incoming
	// connections.
	Port *Port `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	// $hide_from_docs
	// The ip or the Unix domain socket to which the listener should be bound
	// to. Format: `x.x.x.x` or `unix:///path/to/uds` or `unix://@foobar`
	// (Linux abstract namespace). When using Unix domain sockets, the port
	// number should be 0.
	Bind string `protobuf:"bytes,4,opt,name=bind,proto3" json:"bind,omitempty"`
	// One or more hosts exposed by this gateway.
	// While typically applicable to
	// HTTP services, it can also be used for TCP services using TLS with SNI.
	// A host is specified as a `dnsName` with an optional `namespace/` prefix.
	// The `dnsName` should be specified using FQDN format, optionally including
	// a wildcard character in the left-most component (e.g., `prod/*.example.com`).
	// Set the `dnsName` to `*` to select all `VirtualService` hosts from the
	// specified namespace (e.g.,`prod/*`).
	//
	// The `namespace` can be set to `*` or `.`, representing any or the current
	// namespace, respectively. For example, `*/foo.example.com` selects the
	// service from any available namespace while `./foo.example.com` only selects
	// the service from the namespace of the sidecar. The default, if no `namespace/`
	// is specified, is `*/`, that is, select services from any namespace.
	// Any associated `DestinationRule` in the selected namespace will also be used.
	//
	// A `VirtualService` must be bound to the gateway and must have one or
	// more hosts that match the hosts specified in a server. The match
	// could be an exact match or a suffix match with the server's hosts. For
	// example, if the server's hosts specifies `*.example.com`, a
	// `VirtualService` with hosts `dev.example.com` or `prod.example.com` will
	// match. However, a `VirtualService` with host `example.com` or
	// `newexample.com` will not match.
	//
	// NOTE: Only virtual services exported to the gateway's namespace
	// (e.g., `exportTo` value of `*`) can be referenced.
	// Private configurations (e.g., `exportTo` set to `.`) will not be
	// available. Refer to the `exportTo` setting in `VirtualService`,
	// `DestinationRule`, and `ServiceEntry` configurations for details.
	Hosts []string `protobuf:"bytes,2,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// Set of TLS related options that govern the server's behavior. Use
	// these options to control if all http requests should be redirected to
	// https, and the TLS modes to use.
	Tls *ServerTLSSettings `protobuf:"bytes,3,opt,name=tls,proto3" json:"tls,omitempty"`
	// The loopback IP endpoint or Unix domain socket to which traffic should
	// be forwarded to by default. Format should be `127.0.0.1:PORT` or
	// `unix:///path/to/socket` or `unix://@foobar` (Linux abstract namespace).
	// NOT IMPLEMENTED.
	// $hide_from_docs
	DefaultEndpoint string `protobuf:"bytes,5,opt,name=default_endpoint,json=defaultEndpoint,proto3" json:"default_endpoint,omitempty"`
	// An optional name of the server, when set must be unique across all servers.
	// This will be used for variety of purposes like prefixing stats generated with
	// this name etc.
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_067d98d02f84cc0b, []int{1}
}

func (m *Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server.Unmarshal(m, b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server.Marshal(b, m, deterministic)
}
func (m *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(m, src)
}
func (m *Server) XXX_Size() int {
	return xxx_messageInfo_Server.Size(m)
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetPort() *Port {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *Server) GetBind() string {
	if m != nil {
		return m.Bind
	}
	return ""
}

func (m *Server) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *Server) GetTls() *ServerTLSSettings {
	if m != nil {
		return m.Tls
	}
	return nil
}

func (m *Server) GetDefaultEndpoint() string {
	if m != nil {
		return m.DefaultEndpoint
	}
	return ""
}

func (m *Server) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Port describes the properties of a specific port of a service.
type Port struct {
	// A valid non-negative integer port number.
	Number uint32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	// The protocol exposed on the port.
	// MUST BE one of HTTP|HTTPS|GRPC|HTTP2|MONGO|TCP|TLS.
	// TLS implies the connection will be routed based on the SNI header to
	// the destination without terminating the TLS connection.
	Protocol string `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// Label assigned to the port.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The port number on the endpoint where the traffic will be
	// received. Applicable only when used with ServiceEntries.
	TargetPort           uint32   `protobuf:"varint,4,opt,name=target_port,json=targetPort,proto3" json:"target_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_067d98d02f84cc0b, []int{2}
}

func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (m *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(m, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Port) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Port) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Port) GetTargetPort() uint32 {
	if m != nil {
		return m.TargetPort
	}
	return 0
}

type ServerTLSSettings struct {
	// If set to true, the load balancer will send a 301 redirect for
	// all http connections, asking the clients to use HTTPS.
	HttpsRedirect bool `protobuf:"varint,1,opt,name=https_redirect,json=httpsRedirect,proto3" json:"https_redirect,omitempty"`
	// Optional: Indicates whether connections to this port should be
	// secured using TLS. The value of this field determines how TLS is
	// enforced.
	Mode ServerTLSSettings_TLSmode `protobuf:"varint,2,opt,name=mode,proto3,enum=istio.networking.v1alpha3.ServerTLSSettings_TLSmode" json:"mode,omitempty"`
	// REQUIRED if mode is `SIMPLE` or `MUTUAL`. The path to the file
	// holding the server-side TLS certificate to use.
	ServerCertificate string `protobuf:"bytes,3,opt,name=server_certificate,json=serverCertificate,proto3" json:"server_certificate,omitempty"`
	// REQUIRED if mode is `SIMPLE` or `MUTUAL`. The path to the file
	// holding the server's private key.
	PrivateKey string `protobuf:"bytes,4,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// REQUIRED if mode is `MUTUAL`. The path to a file containing
	// certificate authority certificates to use in verifying a presented
	// client side certificate.
	CaCertificates string `protobuf:"bytes,5,opt,name=ca_certificates,json=caCertificates,proto3" json:"ca_certificates,omitempty"`
	// For gateways running on Kubernetes, the name of the secret that
	// holds the TLS certs including the CA certificates. Applicable
	// only on Kubernetes. The secret (of type `generic`) should
	// contain the following keys and values: `key:
	// <privateKey>` and `cert: <serverCert>`. For mutual TLS,
	// `cacert: <CACertificate>` can be provided in the same secret or
	// a separate secret named `<secret>-cacert`.
	// Secret of type tls for server certificates along with
	// ca.crt key for CA certificates is also supported.
	// Only one of server certificates and CA certificate
	// or credentialName can be specified.
	CredentialName string `protobuf:"bytes,10,opt,name=credential_name,json=credentialName,proto3" json:"credential_name,omitempty"`
	// A list of alternate names to verify the subject identity in the
	// certificate presented by the client.
	SubjectAltNames []string `protobuf:"bytes,6,rep,name=subject_alt_names,json=subjectAltNames,proto3" json:"subject_alt_names,omitempty"`
	// An optional list of base64-encoded SHA-256 hashes of the SKPIs of
	// authorized client certificates.
	// Note: When both verify_certificate_hash and verify_certificate_spki
	// are specified, a hash matching either value will result in the
	// certificate being accepted.
	VerifyCertificateSpki []string `protobuf:"bytes,11,rep,name=verify_certificate_spki,json=verifyCertificateSpki,proto3" json:"verify_certificate_spki,omitempty"`
	// An optional list of hex-encoded SHA-256 hashes of the
	// authorized client certificates. Both simple and colon separated
	// formats are acceptable.
	// Note: When both verify_certificate_hash and verify_certificate_spki
	// are specified, a hash matching either value will result in the
	// certificate being accepted.
	VerifyCertificateHash []string `protobuf:"bytes,12,rep,name=verify_certificate_hash,json=verifyCertificateHash,proto3" json:"verify_certificate_hash,omitempty"`
	// Optional: Minimum TLS protocol version.
	MinProtocolVersion ServerTLSSettings_TLSProtocol `protobuf:"varint,7,opt,name=min_protocol_version,json=minProtocolVersion,proto3,enum=istio.networking.v1alpha3.ServerTLSSettings_TLSProtocol" json:"min_protocol_version,omitempty"`
	// Optional: Maximum TLS protocol version.
	MaxProtocolVersion ServerTLSSettings_TLSProtocol `protobuf:"varint,8,opt,name=max_protocol_version,json=maxProtocolVersion,proto3,enum=istio.networking.v1alpha3.ServerTLSSettings_TLSProtocol" json:"max_protocol_version,omitempty"`
	// Optional: If specified, only support the specified cipher list.
	// Otherwise default to the default cipher list supported by Envoy.
	CipherSuites         []string `protobuf:"bytes,9,rep,name=cipher_suites,json=cipherSuites,proto3" json:"cipher_suites,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerTLSSettings) Reset()         { *m = ServerTLSSettings{} }
func (m *ServerTLSSettings) String() string { return proto.CompactTextString(m) }
func (*ServerTLSSettings) ProtoMessage()    {}
func (*ServerTLSSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_067d98d02f84cc0b, []int{3}
}

func (m *ServerTLSSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerTLSSettings.Unmarshal(m, b)
}
func (m *ServerTLSSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerTLSSettings.Marshal(b, m, deterministic)
}
func (m *ServerTLSSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerTLSSettings.Merge(m, src)
}
func (m *ServerTLSSettings) XXX_Size() int {
	return xxx_messageInfo_ServerTLSSettings.Size(m)
}
func (m *ServerTLSSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerTLSSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ServerTLSSettings proto.InternalMessageInfo

func (m *ServerTLSSettings) GetHttpsRedirect() bool {
	if m != nil {
		return m.HttpsRedirect
	}
	return false
}

func (m *ServerTLSSettings) GetMode() ServerTLSSettings_TLSmode {
	if m != nil {
		return m.Mode
	}
	return ServerTLSSettings_PASSTHROUGH
}

func (m *ServerTLSSettings) GetServerCertificate() string {
	if m != nil {
		return m.ServerCertificate
	}
	return ""
}

func (m *ServerTLSSettings) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *ServerTLSSettings) GetCaCertificates() string {
	if m != nil {
		return m.CaCertificates
	}
	return ""
}

func (m *ServerTLSSettings) GetCredentialName() string {
	if m != nil {
		return m.CredentialName
	}
	return ""
}

func (m *ServerTLSSettings) GetSubjectAltNames() []string {
	if m != nil {
		return m.SubjectAltNames
	}
	return nil
}

func (m *ServerTLSSettings) GetVerifyCertificateSpki() []string {
	if m != nil {
		return m.VerifyCertificateSpki
	}
	return nil
}

func (m *ServerTLSSettings) GetVerifyCertificateHash() []string {
	if m != nil {
		return m.VerifyCertificateHash
	}
	return nil
}

func (m *ServerTLSSettings) GetMinProtocolVersion() ServerTLSSettings_TLSProtocol {
	if m != nil {
		return m.MinProtocolVersion
	}
	return ServerTLSSettings_TLS_AUTO
}

func (m *ServerTLSSettings) GetMaxProtocolVersion() ServerTLSSettings_TLSProtocol {
	if m != nil {
		return m.MaxProtocolVersion
	}
	return ServerTLSSettings_TLS_AUTO
}

func (m *ServerTLSSettings) GetCipherSuites() []string {
	if m != nil {
		return m.CipherSuites
	}
	return nil
}

func init() {
	proto.RegisterEnum("istio.networking.v1alpha3.ServerTLSSettings_TLSmode", ServerTLSSettings_TLSmode_name, ServerTLSSettings_TLSmode_value)
	proto.RegisterEnum("istio.networking.v1alpha3.ServerTLSSettings_TLSProtocol", ServerTLSSettings_TLSProtocol_name, ServerTLSSettings_TLSProtocol_value)
	proto.RegisterType((*Gateway)(nil), "istio.networking.v1alpha3.Gateway")
	proto.RegisterMapType((map[string]string)(nil), "istio.networking.v1alpha3.Gateway.SelectorEntry")
	proto.RegisterType((*Server)(nil), "istio.networking.v1alpha3.Server")
	proto.RegisterType((*Port)(nil), "istio.networking.v1alpha3.Port")
	proto.RegisterType((*ServerTLSSettings)(nil), "istio.networking.v1alpha3.ServerTLSSettings")
}

func init() { proto.RegisterFile("networking/v1alpha3/gateway.proto", fileDescriptor_067d98d02f84cc0b) }

var fileDescriptor_067d98d02f84cc0b = []byte{
	// 772 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x5d, 0x6f, 0x1a, 0x47,
	0x14, 0xcd, 0xb2, 0x18, 0xe3, 0x8b, 0xb1, 0xd7, 0x23, 0x57, 0xd9, 0xa4, 0x0f, 0x38, 0x54, 0x55,
	0xd3, 0xaa, 0x85, 0x18, 0x57, 0x95, 0xd5, 0x4a, 0x55, 0x49, 0x65, 0x05, 0xab, 0xa4, 0xd0, 0x5d,
	0xc8, 0x43, 0x5e, 0x46, 0xc3, 0x32, 0xb0, 0x13, 0x96, 0x9d, 0xd5, 0xcc, 0x40, 0xc2, 0x4b, 0xff,
	0x68, 0x5f, 0xfa, 0x03, 0xfa, 0xd0, 0x9f, 0x50, 0xcd, 0x07, 0x81, 0x36, 0x75, 0xaa, 0xa8, 0x6f,
	0x33, 0xe7, 0xde, 0x73, 0xee, 0xbd, 0xe7, 0x0e, 0x2c, 0x3c, 0xca, 0xa9, 0x7a, 0xcd, 0xc5, 0x82,
	0xe5, 0xf3, 0xf6, 0xfa, 0x92, 0x64, 0x45, 0x4a, 0xae, 0xda, 0x73, 0xa2, 0xe8, 0x6b, 0xb2, 0x69,
	0x15, 0x82, 0x2b, 0x8e, 0x1e, 0x30, 0xa9, 0x18, 0x6f, 0xed, 0x12, 0x5b, 0xdb, 0xc4, 0x87, 0x8d,
	0x39, 0xe7, 0xf3, 0x8c, 0xb6, 0x49, 0xc1, 0xda, 0x33, 0x46, 0xb3, 0x29, 0x9e, 0xd0, 0x94, 0xac,
	0x19, 0x17, 0x96, 0xdb, 0xfc, 0xcd, 0x83, 0xc3, 0x67, 0x56, 0x0d, 0xfd, 0x00, 0x87, 0x92, 0x8a,
	0x35, 0x15, 0x32, 0xf4, 0x2e, 0xfc, 0xc7, 0xb5, 0xce, 0xa3, 0xd6, 0x9d, 0xca, 0xad, 0xd8, 0x64,
	0x3e, 0xf5, 0x7f, 0xef, 0x96, 0xa2, 0x2d, 0x0d, 0xfd, 0x02, 0x55, 0x49, 0x33, 0x9a, 0x28, 0x2e,
	0xc2, 0x92, 0x91, 0x78, 0xf2, 0x1e, 0x09, 0x57, 0xb7, 0x15, 0x3b, 0xca, 0x4d, 0xae, 0xc4, 0xc6,
	0x2a, 0xbe, 0x95, 0x79, 0xf8, 0x1d, 0xd4, 0xff, 0x16, 0x47, 0x01, 0xf8, 0x0b, 0xba, 0x09, 0xbd,
	0x0b, 0xef, 0xf1, 0x51, 0xa4, 0x8f, 0xe8, 0x1c, 0x0e, 0xd6, 0x24, 0x5b, 0xd1, 0xb0, 0x64, 0x30,
	0x7b, 0xf9, 0xb6, 0x74, 0xed, 0x35, 0xff, 0xf4, 0xa0, 0x62, 0x1b, 0x45, 0xd7, 0x50, 0x2e, 0xb8,
	0x50, 0x86, 0x57, 0xeb, 0x34, 0xde, 0xd3, 0xd6, 0x90, 0x0b, 0x65, 0xbb, 0x30, 0x0c, 0x84, 0xa0,
	0x3c, 0x61, 0xf9, 0x34, 0x2c, 0x1b, 0x75, 0x73, 0x46, 0x0f, 0xe0, 0x20, 0xe5, 0x52, 0x49, 0x33,
	0xe5, 0x91, 0xcd, 0xb6, 0x08, 0xfa, 0x1e, 0x7c, 0x95, 0xc9, 0xd0, 0x37, 0x75, 0xbe, 0xfc, 0x4f,
	0x07, 0x47, 0xfd, 0x38, 0xa6, 0x4a, 0xb1, 0x7c, 0x2e, 0x23, 0x4d, 0x44, 0x9f, 0x43, 0x30, 0xa5,
	0x33, 0xb2, 0xca, 0x14, 0xa6, 0xf9, 0xb4, 0xe0, 0x2c, 0x57, 0xe1, 0x81, 0x29, 0x7d, 0xea, 0xf0,
	0x1b, 0x07, 0xeb, 0xce, 0x72, 0xb2, 0xa4, 0x61, 0xc5, 0x76, 0xa6, 0xcf, 0xcd, 0x5f, 0xa1, 0xac,
	0x07, 0x40, 0x1f, 0x43, 0x25, 0x5f, 0x2d, 0x27, 0x54, 0x98, 0x89, 0xeb, 0xb6, 0x45, 0x07, 0xa1,
	0x06, 0x54, 0xcd, 0xfa, 0x13, 0x9e, 0x59, 0xd3, 0x9c, 0xeb, 0x5b, 0x10, 0xdd, 0x77, 0xca, 0xfe,
	0x2e, 0x68, 0x00, 0xd4, 0x80, 0x9a, 0x22, 0x62, 0x4e, 0x15, 0x36, 0x6e, 0x6a, 0x4f, 0xea, 0x11,
	0x58, 0x48, 0xd7, 0x6d, 0xfe, 0x51, 0x81, 0xb3, 0x77, 0x26, 0x43, 0x9f, 0xc2, 0x49, 0xaa, 0x54,
	0x21, 0xb1, 0xa0, 0x53, 0x26, 0x68, 0x62, 0xf7, 0x50, 0x8d, 0xea, 0x06, 0x8d, 0x1c, 0x88, 0x7a,
	0x50, 0x5e, 0xf2, 0xa9, 0x5d, 0xe4, 0x49, 0xe7, 0xeb, 0x0f, 0x31, 0xaf, 0x35, 0xea, 0xc7, 0x9a,
	0x1b, 0x19, 0x05, 0xf4, 0x15, 0x20, 0xfb, 0x28, 0x71, 0x42, 0x85, 0x62, 0x33, 0x96, 0x10, 0xe5,
	0xc6, 0x89, 0xce, 0x6c, 0xe4, 0xc7, 0x5d, 0x40, 0x8f, 0x55, 0x08, 0xb6, 0x26, 0x8a, 0x62, 0xfd,
	0xb8, 0xec, 0xaa, 0xc1, 0x41, 0x3f, 0xd1, 0x0d, 0xfa, 0x0c, 0x4e, 0x13, 0xb2, 0xaf, 0x25, 0xdd,
	0x52, 0x4e, 0x12, 0xb2, 0x27, 0x24, 0x4d, 0xa2, 0xa0, 0x53, 0x9a, 0x2b, 0x46, 0x32, 0x6c, 0x4c,
	0x04, 0x97, 0xf8, 0x16, 0xfe, 0x59, 0x3b, 0xf9, 0x05, 0x9c, 0xc9, 0xd5, 0xe4, 0x15, 0x4d, 0x14,
	0x26, 0x99, 0x32, 0x99, 0x32, 0xac, 0xe8, 0xe7, 0x14, 0x9d, 0xba, 0x40, 0x37, 0x53, 0x3a, 0x55,
	0xa2, 0x6f, 0xe0, 0xfe, 0x9a, 0x0a, 0x36, 0xdb, 0xec, 0x77, 0x80, 0x65, 0xb1, 0x60, 0x61, 0xcd,
	0x30, 0x3e, 0xb2, 0xe1, 0xbd, 0x4e, 0xe2, 0x62, 0xc1, 0xee, 0xe0, 0xa5, 0x44, 0xa6, 0xe1, 0xf1,
	0x1d, 0xbc, 0x1e, 0x91, 0x29, 0x7a, 0x05, 0xe7, 0x4b, 0x96, 0xe3, 0xed, 0x73, 0xc0, 0xfa, 0xc7,
	0xcd, 0x78, 0x1e, 0x1e, 0x9a, 0xbd, 0x5c, 0x7f, 0xe8, 0x5e, 0x86, 0x4e, 0x27, 0x42, 0x4b, 0x96,
	0x6f, 0x2f, 0x2f, 0xac, 0xa6, 0xa9, 0x45, 0xde, 0xbc, 0x5b, 0xab, 0xfa, 0xbf, 0x6b, 0x91, 0x37,
	0xff, 0xac, 0xf5, 0x09, 0xd4, 0x13, 0x56, 0xa4, 0x54, 0x60, 0xb9, 0x62, 0x7a, 0x87, 0x47, 0xc6,
	0x85, 0x63, 0x0b, 0xc6, 0x06, 0x6b, 0xbe, 0x84, 0x43, 0xf7, 0x96, 0xd0, 0x29, 0xd4, 0x86, 0xdd,
	0x38, 0x1e, 0xf5, 0xa2, 0xc1, 0xf8, 0x59, 0x2f, 0xb8, 0x87, 0x00, 0x2a, 0xf1, 0xed, 0xf3, 0x61,
	0xff, 0x26, 0xf0, 0xf4, 0xf9, 0xf9, 0x78, 0x34, 0xee, 0xf6, 0x83, 0x12, 0x3a, 0x87, 0xa0, 0x3b,
	0x1e, 0x0d, 0xf0, 0x7e, 0xb6, 0x8f, 0x02, 0x38, 0xbe, 0x8d, 0x47, 0xb7, 0x03, 0xec, 0xf2, 0xca,
	0xcd, 0x01, 0xd4, 0xf6, 0x7a, 0x44, 0xc7, 0x50, 0x1d, 0xf5, 0x63, 0xac, 0xa9, 0xc1, 0x3d, 0x54,
	0x33, 0x85, 0x5f, 0x5c, 0xe2, 0x27, 0x81, 0xb7, 0xbb, 0x5c, 0x06, 0xa5, 0xdd, 0xa5, 0x13, 0xf8,
	0xbb, 0xcb, 0x55, 0x50, 0x7e, 0xda, 0x7c, 0x79, 0x61, 0x0d, 0x62, 0xdc, 0xfc, 0xc9, 0xff, 0xcb,
	0xd7, 0x62, 0x52, 0x31, 0xde, 0x5e, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x43, 0x03, 0x1c,
	0x4b, 0x06, 0x00, 0x00,
}
