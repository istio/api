// Code generated by protoc-gen-go. DO NOT EDIT.
// source: routing/v1alpha2/gateway.proto

package istio_routing_v1alpha2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf2 "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// TLS modes enforced by the gateway
type Server_TLSOptions_TLSmode int32

const (
	// If set to "passthrough", the gateway will forward the connection to the
	// upstream server as is.
	Server_TLSOptions_PASSTHROUGH Server_TLSOptions_TLSmode = 0
	// If set to "simple", the gateway will secure connections with
	// standard TLS semantics (server certs only).
	Server_TLSOptions_SIMPLE Server_TLSOptions_TLSmode = 1
	// If set to "mutual", the gateway will use standard mTLS authentication.
	Server_TLSOptions_MUTUAL Server_TLSOptions_TLSmode = 2
)

var Server_TLSOptions_TLSmode_name = map[int32]string{
	0: "PASSTHROUGH",
	1: "SIMPLE",
	2: "MUTUAL",
}
var Server_TLSOptions_TLSmode_value = map[string]int32{
	"PASSTHROUGH": 0,
	"SIMPLE":      1,
	"MUTUAL":      2,
}

func (x Server_TLSOptions_TLSmode) String() string {
	return proto.EnumName(Server_TLSOptions_TLSmode_name, int32(x))
}
func (Server_TLSOptions_TLSmode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor2, []int{1, 1, 0}
}

// Gateway describes a load balancer operating at the edge of the mesh
// receiving incoming or outgoing HTTP/TCP connections. The specification
// describes a set of ports that should be exposed, the type of protocol to
// use, SNI configuration for the load balancer, etc.
//
// For example, the following gateway spec sets up a proxy to act as a load
// balancer exposing port 80 and 9080 (http), 443 (https), and port 2379
// (TCP) for ingress.  While Istio will configure the proxy to listen on
// these ports, it is the responsibility of the user to ensure that
// external traffic to these ports are allowed into the mesh.
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: Gateway
//     metadata:
//       name: my-gateway
//     spec:
//       servers:
//       - port:
//           number: 80
//           name: http
//         domains:
//         - uk.bookinfo.com
//         - eu.bookinfo.com
//         tls:
//           httpsRedirect: true # sends 302 redirect for http requests
//       - port:
//           number: 443
//           name: https
//         domains:
//         - uk.bookinfo.com
//         - eu.bookinfo.com
//         tls:
//           mode: simple #enables HTTPS on this port
//           serverCert: server.crt
//           clientCABundle: client.ca-bundle
//       - port:
//           number: 9080
//           name: http-wildcard
//         # no domains implies wildcard match
//       - port:
//           number: 2379 #to expose internal service via external port 2379
//           name: Mongo
//           protocol: MONGO
//
// The gateway specification above describes the L4-L6 properties of a load
// balancer. Routing rules can then be bound to a gateway to control
// the forwarding of traffic arriving at a particular domain or gateway port.
//
// The following sample route rule splits traffic for
// https://uk.bookinfo.com/reviews, https://eu.bookinfo.com/reviews,
// http://uk.bookinfo.com:9080/reviews, http://eu.bookinfo.com:9080/reviews
// into two versions (prod and qa) of an internal reviews service on port
// 9080. In addition, requests containing the cookie user: dev-123 will be
// sent to special port 7777 in the qa version. The same rule is also
// applicable inside the mesh for requests to the reviews.prod
// service. This rule is applicable across ports 443, 9080. Note that
// http://uk.bookinfo.com gets redirected to https://uk.bookinfo.com
// (i.e. 80 redirects to 443).
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: RouteRule
//     metadata:
//       name: bookinfo-rule
//     spec:
//       hosts:
//       - reviews.prod
//       - uk.bookinfo.com
//       - eu.bookinfo.com
//       gateways:
//       - my-gateway
//       - mesh # applies to all the sidecars in the mesh
//       http:
//       - match:
//         - headers:
//             cookie:
//               user: dev-123
//         route:
//         - destination:
//             port:
//               number: 7777
//             name: reviews.qa
//       - match:
//           uri:
//             prefix: /reviews/
//         route:
//         - destination:
//             port:
//               number: 9080 # can be omitted if its the only port for reviews
//             name: reviews.prod
//           weight: 80
//         - destination:
//             name: reviews.qa
//           weight: 20
//
// The following routing rule forwards traffic arriving at (external) port
// 2379 from 172.17.16.0/24 subnet to internal Mongo server on port 5555. This
// rule is not applicable internally in the mesh as the gateway list omits
// the reserved name "mesh".
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: RouteRule
//     metadata:
//       name: bookinfo-Mongo
//     spec:
//       hosts:
//       - Mongosvr #name of Mongo service
//       gateways:
//       - my-gateway
//       tcp:
//       - match:
//         - port:
//             number: 2379
//           sourceSubnet: "172.17.16.0/24"
//         route:
//         - destination:
//             name: mongo.prod
//
type Gateway struct {
	// REQUIRED: A list of server specifications.
	Servers []*Server `protobuf:"bytes,1,rep,name=servers" json:"servers,omitempty"`
}

func (m *Gateway) Reset()                    { *m = Gateway{} }
func (m *Gateway) String() string            { return proto.CompactTextString(m) }
func (*Gateway) ProtoMessage()               {}
func (*Gateway) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Gateway) GetServers() []*Server {
	if m != nil {
		return m.Servers
	}
	return nil
}

// Server describes the properties of the proxy on a given load balancer port.
// For example,
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: Gateway
//     metadata:
//       name: my-ingress
//     spec:
//       servers:
//       - port:
//           number: 80
//           protocol: HTTP2
//
// Another example
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: Gateway
//     metadata:
//       name: my-tcp-ingress
//     spec:
//       servers:
//       - port:
//           number: 27018
//           protocol: MONGO
//
// The following is an example of TLS configuration for port 443
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: Gateway
//     metadata:
//       name: my-ingress
//     spec:
//       servers:
//       - port:
//           number: 443
//           protocol: HTTP
//         tls:
//           mode: simple
//           serverCertificatet: server.crt
//
type Server struct {
	// REQUIRED: The Port on which the proxy should listen for incoming
	// connections
	Port *Server_Port `protobuf:"bytes,1,opt,name=port" json:"port,omitempty"`
	// A list of domains exposed by this gateway. While
	// typically applicable to HTTP services, it can also be used for TCP
	// services using TLS with SNI. Standard DNS wildcard prefix syntax
	// is permitted.
	//
	// RouteRules that are bound to a gateway must having a matching domain
	// in their default destination. Specifically one of the route rule
	// destination domains is a strict suffix of a gateway domain or
	// a gateway domain is a suffix of one of the route rule domains.
	Domains []string `protobuf:"bytes,2,rep,name=domains" json:"domains,omitempty"`
	// Set of TLS related options that govern the server's behavior. Use
	// these options to control if all http requests should be redirected to
	// https, and the TLS modes to use.
	Tls *Server_TLSOptions `protobuf:"bytes,3,opt,name=tls" json:"tls,omitempty"`
	// Extensions contain proxy implementation specific configuration, that
	// cannot be easily expressed through expressed using proxy-independent
	// routing rule or mixer policy configuration. Custom proxy
	// implementations could have enhancements such as caching, dynamic
	// scripting, proprietary protocol support, etc. The configuration for
	// such features could be provided through the extensions mechanism.
	//
	// To extend the proxy configuration for a sidecar, the port and the
	// domains used in the Gateway configuration must match the port and IP
	// on which the sidecar is listening for connections. See Envoy filters
	// configuration for usage examples.
	ProxyExtensions []*google_protobuf2.Any `protobuf:"bytes,4,rep,name=proxy_extensions,json=proxyExtensions" json:"proxy_extensions,omitempty"`
}

func (m *Server) Reset()                    { *m = Server{} }
func (m *Server) String() string            { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()               {}
func (*Server) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Server) GetPort() *Server_Port {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *Server) GetDomains() []string {
	if m != nil {
		return m.Domains
	}
	return nil
}

func (m *Server) GetTls() *Server_TLSOptions {
	if m != nil {
		return m.Tls
	}
	return nil
}

func (m *Server) GetProxyExtensions() []*google_protobuf2.Any {
	if m != nil {
		return m.ProxyExtensions
	}
	return nil
}

// Port describes the properties of a specific port of a service.
type Server_Port struct {
	// REQUIRED: A valid non-negative integer port number.
	Number uint32 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
	// The protocol exposed on the port.
	// MUST BE one of HTTP|HTTPS|GRPC|HTTP2|MONGO|TCP.
	Protocol string `protobuf:"bytes,2,opt,name=protocol" json:"protocol,omitempty"`
	// Label assigned to the port.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *Server_Port) Reset()                    { *m = Server_Port{} }
func (m *Server_Port) String() string            { return proto.CompactTextString(m) }
func (*Server_Port) ProtoMessage()               {}
func (*Server_Port) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0} }

func (m *Server_Port) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Server_Port) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Server_Port) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Server_TLSOptions struct {
	// If set to true, the load balancer will send a 302 redirect for all
	// http connections, asking the clients to use HTTPS.
	HttpsRedirect bool `protobuf:"varint,1,opt,name=https_redirect,json=httpsRedirect" json:"https_redirect,omitempty"`
	// Optional: Indicates whether connections to this port should be
	// secured using TLS.  The value of this field determines how TLS is
	// enforced.
	Mode Server_TLSOptions_TLSmode `protobuf:"varint,2,opt,name=mode,enum=istio.routing.v1alpha2.Server_TLSOptions_TLSmode" json:"mode,omitempty"`
	// REQUIRED if mode == SIMPLE/MUTUAL. The name of the file holding the
	// server-side TLS certificate to use.  It is the responsibility of the
	// underlying platform to mount the certificate as a file under
	// /etc/istio/ingress-certs with the same name as the specified in this
	// field.
	ServerCertificate string `protobuf:"bytes,3,opt,name=server_certificate,json=serverCertificate" json:"server_certificate,omitempty"`
	// REQUIRED if mode == MUTUAL. To use mutual TLS for external clients,
	// specify the name of the file holding the CA certificate to validate
	// the client's certificate. It is the responsibility of the underlying
	// platform to mount the certificate as a file under
	// /etc/istio/ingress-certs with the same name as specified in this
	// field.
	ClientCaBundle string `protobuf:"bytes,4,opt,name=client_ca_bundle,json=clientCaBundle" json:"client_ca_bundle,omitempty"`
}

func (m *Server_TLSOptions) Reset()                    { *m = Server_TLSOptions{} }
func (m *Server_TLSOptions) String() string            { return proto.CompactTextString(m) }
func (*Server_TLSOptions) ProtoMessage()               {}
func (*Server_TLSOptions) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 1} }

func (m *Server_TLSOptions) GetHttpsRedirect() bool {
	if m != nil {
		return m.HttpsRedirect
	}
	return false
}

func (m *Server_TLSOptions) GetMode() Server_TLSOptions_TLSmode {
	if m != nil {
		return m.Mode
	}
	return Server_TLSOptions_PASSTHROUGH
}

func (m *Server_TLSOptions) GetServerCertificate() string {
	if m != nil {
		return m.ServerCertificate
	}
	return ""
}

func (m *Server_TLSOptions) GetClientCaBundle() string {
	if m != nil {
		return m.ClientCaBundle
	}
	return ""
}

func init() {
	proto.RegisterType((*Gateway)(nil), "istio.routing.v1alpha2.Gateway")
	proto.RegisterType((*Server)(nil), "istio.routing.v1alpha2.Server")
	proto.RegisterType((*Server_Port)(nil), "istio.routing.v1alpha2.Server.Port")
	proto.RegisterType((*Server_TLSOptions)(nil), "istio.routing.v1alpha2.Server.TLSOptions")
	proto.RegisterEnum("istio.routing.v1alpha2.Server_TLSOptions_TLSmode", Server_TLSOptions_TLSmode_name, Server_TLSOptions_TLSmode_value)
}

func init() { proto.RegisterFile("routing/v1alpha2/gateway.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0xad, 0x6c, 0x61, 0xc7, 0x63, 0xe2, 0xa8, 0x43, 0x09, 0xaa, 0x0f, 0xc1, 0xb8, 0x14, 0xd4,
	0x43, 0xd7, 0x44, 0x3d, 0xb4, 0xd0, 0x43, 0x71, 0x8d, 0x49, 0x0a, 0x4e, 0x63, 0x56, 0xf6, 0x59,
	0xac, 0xe5, 0x8d, 0xb3, 0x20, 0xef, 0x8a, 0xdd, 0x75, 0x1a, 0x7f, 0x41, 0x3f, 0xa8, 0x3f, 0x58,
	0xb4, 0x92, 0x92, 0x4b, 0x69, 0x7b, 0x9b, 0x79, 0xf3, 0xde, 0xdb, 0x99, 0xb7, 0x70, 0xa1, 0xd5,
	0xc1, 0x0a, 0xb9, 0x9b, 0x3c, 0x5c, 0xb2, 0xbc, 0xb8, 0x67, 0xf1, 0x64, 0xc7, 0x2c, 0xff, 0xc1,
	0x8e, 0xa4, 0xd0, 0xca, 0x2a, 0x3c, 0x17, 0xc6, 0x0a, 0x45, 0x6a, 0x16, 0x69, 0x58, 0xc3, 0xd7,
	0x3b, 0xa5, 0x76, 0x39, 0x9f, 0x38, 0xd6, 0xe6, 0x70, 0x37, 0x61, 0xb2, 0x96, 0x8c, 0x67, 0xd0,
	0xbd, 0xaa, 0x3c, 0xf0, 0x13, 0x74, 0x0d, 0xd7, 0x0f, 0x5c, 0x9b, 0xd0, 0x1b, 0xb5, 0xa3, 0x7e,
	0x7c, 0x41, 0xfe, 0xec, 0x47, 0x12, 0x47, 0xa3, 0x0d, 0x7d, 0xfc, 0xcb, 0x87, 0x4e, 0x85, 0xe1,
	0x47, 0xf0, 0x0b, 0xa5, 0x6d, 0xe8, 0x8d, 0xbc, 0xa8, 0x1f, 0xbf, 0xf9, 0xbb, 0x03, 0x59, 0x2a,
	0x6d, 0xa9, 0x13, 0x60, 0x08, 0xdd, 0xad, 0xda, 0x33, 0x21, 0x4d, 0xd8, 0x1a, 0xb5, 0xa3, 0x1e,
	0x6d, 0x5a, 0xfc, 0x0c, 0x6d, 0x9b, 0x9b, 0xb0, 0xed, 0x1c, 0xdf, 0xfd, 0xc3, 0x71, 0xb5, 0x48,
	0x6e, 0x0b, 0x2b, 0x94, 0x34, 0xb4, 0x54, 0xe1, 0x17, 0x08, 0x0a, 0xad, 0x1e, 0x8f, 0x29, 0x7f,
	0xb4, 0x5c, 0x9a, 0x72, 0x10, 0xfa, 0xee, 0xba, 0x57, 0xa4, 0x4a, 0x85, 0x34, 0xa9, 0x90, 0xa9,
	0x3c, 0xd2, 0x33, 0xc7, 0x9e, 0x3f, 0x91, 0x87, 0xdf, 0xc1, 0x2f, 0xb7, 0xc4, 0x73, 0xe8, 0xc8,
	0xc3, 0x7e, 0xc3, 0xb5, 0x3b, 0xed, 0x94, 0xd6, 0x1d, 0x0e, 0xe1, 0xc4, 0x19, 0x64, 0x2a, 0x0f,
	0x5b, 0x23, 0x2f, 0xea, 0xd1, 0xa7, 0x1e, 0x11, 0x7c, 0xc9, 0xf6, 0xdc, 0xad, 0xde, 0xa3, 0xae,
	0x1e, 0xfe, 0x6c, 0x01, 0x3c, 0x2f, 0x89, 0x6f, 0x61, 0x70, 0x6f, 0x6d, 0x61, 0x52, 0xcd, 0xb7,
	0x42, 0xf3, 0xac, 0x4a, 0xee, 0x84, 0x9e, 0x3a, 0x94, 0xd6, 0x20, 0xce, 0xc1, 0xdf, 0xab, 0x2d,
	0x77, 0x2f, 0x0c, 0xe2, 0xcb, 0xff, 0x0e, 0xa1, 0x2c, 0x4b, 0x21, 0x75, 0x72, 0x7c, 0x0f, 0x58,
	0xfd, 0x59, 0x9a, 0x71, 0x6d, 0xc5, 0x9d, 0xc8, 0x98, 0x6d, 0xd6, 0x7b, 0x59, 0x4d, 0x66, 0xcf,
	0x03, 0x8c, 0x20, 0xc8, 0x72, 0xc1, 0xa5, 0x4d, 0x33, 0x96, 0x6e, 0x0e, 0x72, 0x9b, 0xf3, 0xd0,
	0x77, 0xe4, 0x41, 0x85, 0xcf, 0xd8, 0x57, 0x87, 0x8e, 0x63, 0xe8, 0xd6, 0x2f, 0xe1, 0x19, 0xf4,
	0x97, 0xd3, 0x24, 0x59, 0x5d, 0xd3, 0xdb, 0xf5, 0xd5, 0x75, 0xf0, 0x02, 0x01, 0x3a, 0xc9, 0xb7,
	0x9b, 0xe5, 0x62, 0x1e, 0x78, 0x65, 0x7d, 0xb3, 0x5e, 0xad, 0xa7, 0x8b, 0xa0, 0xb5, 0xe9, 0xb8,
	0x9c, 0x3e, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x73, 0x93, 0xb4, 0x0d, 0xd6, 0x02, 0x00, 0x00,
}
