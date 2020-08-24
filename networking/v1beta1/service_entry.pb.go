// Code generated by protoc-gen-go. DO NOT EDIT.
// source: networking/v1beta1/service_entry.proto

// `ServiceEntry` enables adding additional entries into Istio's
// internal service registry, so that auto-discovered services in the
// mesh can access/route to these manually specified services. A
// service entry describes the properties of a service (DNS name,
// VIPs, ports, protocols, endpoints). These services could be
// external to the mesh (e.g., web APIs) or mesh-internal services
// that are not part of the platform's service registry (e.g., a set
// of VMs talking to services in Kubernetes). In addition, the
// endpoints of a service entry can also be dynamically selected by
// using the `workloadSelector` field. These endpoints can be VM
// workloads declared using the `WorkloadEntry` object or Kubernetes
// pods. The ability to select both pods and VMs under a single
// service allows for migration of services from VMs to Kubernetes
// without having to change the existing DNS names associated with the
// services.
//
// The following example declares a few external APIs accessed by internal
// applications over HTTPS. The sidecar inspects the SNI value in the
// ClientHello message to route to the appropriate external service.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-https
// spec:
//   hosts:
//   - api.dropboxapi.com
//   - www.googleapis.com
//   - api.facebook.com
//   location: MESH_EXTERNAL
//   ports:
//   - number: 443
//     name: https
//     protocol: TLS
//   resolution: DNS
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: external-svc-https
// spec:
//   hosts:
//   - api.dropboxapi.com
//   - www.googleapis.com
//   - api.facebook.com
//   location: MESH_EXTERNAL
//   ports:
//   - number: 443
//     name: https
//     protocol: TLS
//   resolution: DNS
// ```
// {{</tab>}}
// {{</tabset>}}
//
// The following configuration adds a set of MongoDB instances running on
// unmanaged VMs to Istio's registry, so that these services can be treated
// as any other service in the mesh. The associated DestinationRule is used
// to initiate mTLS connections to the database instances.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-mongocluster
// spec:
//   hosts:
//   - mymongodb.somedomain # not used
//   addresses:
//   - 192.192.192.192/24 # VIPs
//   ports:
//   - number: 27018
//     name: mongodb
//     protocol: MONGO
//   location: MESH_INTERNAL
//   resolution: STATIC
//   endpoints:
//   - address: 2.2.2.2
//   - address: 3.3.3.3
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: external-svc-mongocluster
// spec:
//   hosts:
//   - mymongodb.somedomain # not used
//   addresses:
//   - 192.192.192.192/24 # VIPs
//   ports:
//   - number: 27018
//     name: mongodb
//     protocol: MONGO
//   location: MESH_INTERNAL
//   resolution: STATIC
//   endpoints:
//   - address: 2.2.2.2
//   - address: 3.3.3.3
// ```
// {{</tab>}}
// {{</tabset>}}
//
// and the associated DestinationRule
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: DestinationRule
// metadata:
//   name: mtls-mongocluster
// spec:
//   host: mymongodb.somedomain
//   trafficPolicy:
//     tls:
//       mode: MUTUAL
//       clientCertificate: /etc/certs/myclientcert.pem
//       privateKey: /etc/certs/client_private_key.pem
//       caCertificates: /etc/certs/rootcacerts.pem
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: DestinationRule
// metadata:
//   name: mtls-mongocluster
// spec:
//   host: mymongodb.somedomain
//   trafficPolicy:
//     tls:
//       mode: MUTUAL
//       clientCertificate: /etc/certs/myclientcert.pem
//       privateKey: /etc/certs/client_private_key.pem
//       caCertificates: /etc/certs/rootcacerts.pem
// ```
// {{</tab>}}
// {{</tabset>}}
//
// The following example uses a combination of service entry and TLS
// routing in a virtual service to steer traffic based on the SNI value to
// an internal egress firewall.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-redirect
// spec:
//   hosts:
//   - wikipedia.org
//   - "*.wikipedia.org"
//   location: MESH_EXTERNAL
//   ports:
//   - number: 443
//     name: https
//     protocol: TLS
//   resolution: NONE
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: external-svc-redirect
// spec:
//   hosts:
//   - wikipedia.org
//   - "*.wikipedia.org"
//   location: MESH_EXTERNAL
//   ports:
//   - number: 443
//     name: https
//     protocol: TLS
//   resolution: NONE
// ```
// {{</tab>}}
// {{</tabset>}}
//
// And the associated VirtualService to route based on the SNI value.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: VirtualService
// metadata:
//   name: tls-routing
// spec:
//   hosts:
//   - wikipedia.org
//   - "*.wikipedia.org"
//   tls:
//   - match:
//     - sniHosts:
//       - wikipedia.org
//       - "*.wikipedia.org"
//     route:
//     - destination:
//         host: internal-egress-firewall.ns1.svc.cluster.local
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: VirtualService
// metadata:
//   name: tls-routing
// spec:
//   hosts:
//   - wikipedia.org
//   - "*.wikipedia.org"
//   tls:
//   - match:
//     - sniHosts:
//       - wikipedia.org
//       - "*.wikipedia.org"
//     route:
//     - destination:
//         host: internal-egress-firewall.ns1.svc.cluster.local
// ```
// {{</tab>}}
// {{</tabset>}}
//
// The virtual service with TLS match serves to override the default SNI
// match. In the absence of a virtual service, traffic will be forwarded to
// the wikipedia domains.
//
// The following example demonstrates the use of a dedicated egress gateway
// through which all external service traffic is forwarded.
// The 'exportTo' field allows for control over the visibility of a service
// declaration to other namespaces in the mesh. By default, a service is exported
// to all namespaces. The following example restricts the visibility to the
// current namespace, represented by ".", so that it cannot be used by other
// namespaces.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-httpbin
//   namespace : egress
// spec:
//   hosts:
//   - httpbin.com
//   exportTo:
//   - "."
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: DNS
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: external-svc-httpbin
//   namespace : egress
// spec:
//   hosts:
//   - httpbin.com
//   exportTo:
//   - "."
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: DNS
// ```
// {{</tab>}}
// {{</tabset>}}
//
// Define a gateway to handle all egress traffic.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: Gateway
// metadata:
//  name: istio-egressgateway
//  namespace: istio-system
// spec:
//  selector:
//    istio: egressgateway
//  servers:
//  - port:
//      number: 80
//      name: http
//      protocol: HTTP
//    hosts:
//    - "*"
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Gateway
// metadata:
//  name: istio-egressgateway
//  namespace: istio-system
// spec:
//  selector:
//    istio: egressgateway
//  servers:
//  - port:
//      number: 80
//      name: http
//      protocol: HTTP
//    hosts:
//    - "*"
// ```
// {{</tab>}}
// {{</tabset>}}
//
// And the associated `VirtualService` to route from the sidecar to the
// gateway service (`istio-egressgateway.istio-system.svc.cluster.local`), as
// well as route from the gateway to the external service. Note that the
// virtual service is exported to all namespaces enabling them to route traffic
// through the gateway to the external service. Forcing traffic to go through
// a managed middle proxy like this is a common practice.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: VirtualService
// metadata:
//   name: gateway-routing
//   namespace: egress
// spec:
//   hosts:
//   - httpbin.com
//   exportTo:
//   - "*"
//   gateways:
//   - mesh
//   - istio-egressgateway
//   http:
//   - match:
//     - port: 80
//       gateways:
//       - mesh
//     route:
//     - destination:
//         host: istio-egressgateway.istio-system.svc.cluster.local
//   - match:
//     - port: 80
//       gateways:
//       - istio-egressgateway
//     route:
//     - destination:
//         host: httpbin.com
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: VirtualService
// metadata:
//   name: gateway-routing
//   namespace: egress
// spec:
//   hosts:
//   - httpbin.com
//   exportTo:
//   - "*"
//   gateways:
//   - mesh
//   - istio-egressgateway
//   http:
//   - match:
//     - port: 80
//       gateways:
//       - mesh
//     route:
//     - destination:
//         host: istio-egressgateway.istio-system.svc.cluster.local
//   - match:
//     - port: 80
//       gateways:
//       - istio-egressgateway
//     route:
//     - destination:
//         host: httpbin.com
// ```
// {{</tab>}}
// {{</tabset>}}
//
// The following example demonstrates the use of wildcards in the hosts for
// external services. If the connection has to be routed to the IP address
// requested by the application (i.e. application resolves DNS and attempts
// to connect to a specific IP), the discovery mode must be set to `NONE`.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-wildcard-example
// spec:
//   hosts:
//   - "*.bar.com"
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: NONE
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: external-svc-wildcard-example
// spec:
//   hosts:
//   - "*.bar.com"
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: NONE
// ```
// {{</tab>}}
// {{</tabset>}}
//
// The following example demonstrates a service that is available via a
// Unix Domain Socket on the host of the client. The resolution must be
// set to STATIC to use Unix address endpoints.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: unix-domain-socket-example
// spec:
//   hosts:
//   - "example.unix.local"
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: STATIC
//   endpoints:
//   - address: unix:///var/run/example/socket
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: unix-domain-socket-example
// spec:
//   hosts:
//   - "example.unix.local"
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: STATIC
//   endpoints:
//   - address: unix:///var/run/example/socket
// ```
// {{</tab>}}
// {{</tabset>}}
//
// For HTTP-based services, it is possible to create a `VirtualService`
// backed by multiple DNS addressable endpoints. In such a scenario, the
// application can use the `HTTP_PROXY` environment variable to transparently
// reroute API calls for the `VirtualService` to a chosen backend. For
// example, the following configuration creates a non-existent external
// service called foo.bar.com backed by three domains: us.foo.bar.com:8080,
// uk.foo.bar.com:9080, and in.foo.bar.com:7080
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-dns
// spec:
//   hosts:
//   - foo.bar.com
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: DNS
//   endpoints:
//   - address: us.foo.bar.com
//     ports:
//       http: 8080
//   - address: uk.foo.bar.com
//     ports:
//       http: 9080
//   - address: in.foo.bar.com
//     ports:
//       http: 7080
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: external-svc-dns
// spec:
//   hosts:
//   - foo.bar.com
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: DNS
//   endpoints:
//   - address: us.foo.bar.com
//     ports:
//       http: 8080
//   - address: uk.foo.bar.com
//     ports:
//       http: 9080
//   - address: in.foo.bar.com
//     ports:
//       http: 7080
// ```
// {{</tab>}}
// {{</tabset>}}
//
// With `HTTP_PROXY=http://localhost/`, calls from the application to
// `http://foo.bar.com` will be load balanced across the three domains
// specified above. In other words, a call to `http://foo.bar.com/baz` would
// be translated to `http://uk.foo.bar.com/baz`.
//
// The following example illustrates the usage of a `ServiceEntry`
// containing a subject alternate name
// whose format conforms to the [SPIFFE standard](https://github.com/spiffe/spiffe/blob/master/standards/SPIFFE-ID.md):
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: httpbin
//   namespace : httpbin-ns
// spec:
//   hosts:
//   - httpbin.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: STATIC
//   endpoints:
//   - address: 2.2.2.2
//   - address: 3.3.3.3
//   subjectAltNames:
//   - "spiffe://cluster.local/ns/httpbin-ns/sa/httpbin-service-account"
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: httpbin
//   namespace : httpbin-ns
// spec:
//   hosts:
//   - httpbin.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: STATIC
//   endpoints:
//   - address: 2.2.2.2
//   - address: 3.3.3.3
//   subjectAltNames:
//   - "spiffe://cluster.local/ns/httpbin-ns/sa/httpbin-service-account"
// ```
// {{</tab>}}
// {{</tabset>}}
//
// The following example demonstrates the use of `ServiceEntry` with a
// `workloadSelector` to handle the migration of a service
// `details.bookinfo.com` from VMs to Kubernetes. The service has two
// VM-based instances with sidecars as well as a set of Kubernetes
// pods managed by a standard deployment object. Consumers of this
// service in the mesh will be automatically load balanced across the
// VMs and Kubernetes.  VM for the `details.bookinfo.com`
// service. This VM has sidecar installed and bootstrapped using the
// `details-legacy` service account. The sidecar receives HTTP traffic
// on port 80 (wrapped in istio mutual TLS) and forwards it to the
// application on the localhost on the same port.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: WorkloadEntry
// metadata:
//   name: details-vm-1
// spec:
//   serviceAccount: details
//   address: 2.2.2.2
//   labels:
//     app: details
//     instance-id: vm1
// ---
// apiVersion: networking.istio.io/v1alpha3
// kind: WorkloadEntry
// metadata:
//   name: details-vm-2
// spec:
//   serviceAccount: details
//   address: 3.3.3.3
//   labels:
//     app: details
//     instance-id: vm2
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: WorkloadEntry
// metadata:
//   name: details-vm-1
// spec:
//   serviceAccount: details
//   address: 2.2.2.2
//   labels:
//     app: details
//     instance-id: vm1
// ---
// apiVersion: networking.istio.io/v1beta1
// kind: WorkloadEntry
// metadata:
//   name: details-vm-2
// spec:
//   serviceAccount: details
//   address: 3.3.3.3
//   labels:
//     app: details
//     instance-id: vm2
// ```
// {{</tab>}}
// {{</tabset>}}
//
// Assuming there is also a Kubernetes deployment with pod labels
// `app: details` using the same service account `details`, the
// following service entry declares a service spanning both VMs and
// Kubernetes:
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
//   resolution: STATIC
//   workloadSelector:
//     labels:
//       app: details
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
//   resolution: STATIC
//   workloadSelector:
//     labels:
//       app: details
// ```
// {{</tab>}}
// {{</tabset>}}

package v1beta1

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

// Location specifies whether the service is part of Istio mesh or
// outside the mesh.  Location determines the behavior of several
// features, such as service-to-service mTLS authentication, policy
// enforcement, etc. When communicating with services outside the mesh,
// Istio's mTLS authentication is disabled, and policy enforcement is
// performed on the client-side as opposed to server-side.
type ServiceEntry_Location int32

const (
	// Signifies that the service is external to the mesh. Typically used
	// to indicate external services consumed through APIs.
	ServiceEntry_MESH_EXTERNAL ServiceEntry_Location = 0
	// Signifies that the service is part of the mesh. Typically used to
	// indicate services added explicitly as part of expanding the service
	// mesh to include unmanaged infrastructure (e.g., VMs added to a
	// Kubernetes based service mesh).
	ServiceEntry_MESH_INTERNAL ServiceEntry_Location = 1
)

var ServiceEntry_Location_name = map[int32]string{
	0: "MESH_EXTERNAL",
	1: "MESH_INTERNAL",
}

var ServiceEntry_Location_value = map[string]int32{
	"MESH_EXTERNAL": 0,
	"MESH_INTERNAL": 1,
}

func (x ServiceEntry_Location) String() string {
	return proto.EnumName(ServiceEntry_Location_name, int32(x))
}

func (ServiceEntry_Location) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_be685d9fa1e5ca12, []int{0, 0}
}

// Resolution determines how the proxy will resolve the IP addresses of
// the network endpoints associated with the service, so that it can
// route to one of them. The resolution mode specified here has no impact
// on how the application resolves the IP address associated with the
// service. The application may still have to use DNS to resolve the
// service to an IP so that the outbound traffic can be captured by the
// Proxy. Alternatively, for HTTP services, the application could
// directly communicate with the proxy (e.g., by setting HTTP_PROXY) to
// talk to these services.
type ServiceEntry_Resolution int32

const (
	// Assume that incoming connections have already been resolved (to a
	// specific destination IP address). Such connections are typically
	// routed via the proxy using mechanisms such as IP table REDIRECT/
	// eBPF. After performing any routing related transformations, the
	// proxy will forward the connection to the IP address to which the
	// connection was bound.
	ServiceEntry_NONE ServiceEntry_Resolution = 0
	// Use the static IP addresses specified in endpoints (see below) as the
	// backing instances associated with the service.
	ServiceEntry_STATIC ServiceEntry_Resolution = 1
	// Attempt to resolve the IP address by querying the ambient DNS,
	// during request processing. If no endpoints are specified, the proxy
	// will resolve the DNS address specified in the hosts field, if
	// wildcards are not used. If endpoints are specified, the DNS
	// addresses specified in the endpoints will be resolved to determine
	// the destination IP address.  DNS resolution cannot be used with Unix
	// domain socket endpoints.
	ServiceEntry_DNS ServiceEntry_Resolution = 2
)

var ServiceEntry_Resolution_name = map[int32]string{
	0: "NONE",
	1: "STATIC",
	2: "DNS",
}

var ServiceEntry_Resolution_value = map[string]int32{
	"NONE":   0,
	"STATIC": 1,
	"DNS":    2,
}

func (x ServiceEntry_Resolution) String() string {
	return proto.EnumName(ServiceEntry_Resolution_name, int32(x))
}

func (ServiceEntry_Resolution) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_be685d9fa1e5ca12, []int{0, 1}
}

// ServiceEntry enables adding additional entries into Istio's internal
// service registry.
//
// <!-- crd generation tags
// +cue-gen:ServiceEntry:groupName:networking.istio.io
// +cue-gen:ServiceEntry:version:v1beta1
// +cue-gen:ServiceEntry:annotations:helm.sh/resource-policy=keep
// +cue-gen:ServiceEntry:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:ServiceEntry:subresource:status
// +cue-gen:ServiceEntry:scope:Namespaced
// +cue-gen:ServiceEntry:resource:categories=istio-io,networking-istio-io,shortNames=se,plural=serviceentries
// +cue-gen:ServiceEntry:printerColumn:name=Hosts,type=string,JSONPath=.spec.hosts,description="The hosts associated with the ServiceEntry"
// +cue-gen:ServiceEntry:printerColumn:name=Location,type=string,JSONPath=.spec.location,description="Whether the service is external to the
// mesh or part of the mesh (MESH_EXTERNAL or MESH_INTERNAL)"
// +cue-gen:ServiceEntry:printerColumn:name=Resolution,type=string,JSONPath=.spec.resolution,description="Service discovery mode for the hosts
// (NONE, STATIC, or DNS)"
// +cue-gen:ServiceEntry:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.istio.io/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type ServiceEntry struct {
	// The hosts associated with the ServiceEntry. Could be a DNS
	// name with wildcard prefix.
	//
	// 1. The hosts field is used to select matching hosts in VirtualServices and DestinationRules.
	// 2. For HTTP traffic the HTTP Host/Authority header will be matched against the hosts field.
	// 3. For HTTPs or TLS traffic containing Server Name Indication (SNI), the SNI value
	// will be matched against the hosts field.
	//
	// **NOTE 1:** When resolution is set to type DNS and no endpoints
	// are specified, the host field will be used as the DNS name of the
	// endpoint to route traffic to.
	//
	// **NOTE 2:** If the hostname matches with the name of a service
	// from another service registry such as Kubernetes that also
	// supplies its own set of endpoints, the ServiceEntry will be
	// treated as a decorator of the existing Kubernetes
	// service. Properties in the service entry will be added to the
	// Kubernetes service if applicable. Currently, the only the
	// following additional properties will be considered by `istiod`:
	//
	// 1. subjectAltNames: In addition to verifying the SANs of the
	//    service accounts associated with the pods of the service, the
	//    SANs specified here will also be verified.
	//
	Hosts []string `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// The virtual IP addresses associated with the service. Could be CIDR
	// prefix. For HTTP traffic, generated route configurations will include http route
	// domains for both the `addresses` and `hosts` field values and the destination will
	// be identified based on the HTTP Host/Authority header.
	// If one or more IP addresses are specified,
	// the incoming traffic will be identified as belonging to this service
	// if the destination IP matches the IP/CIDRs specified in the addresses
	// field. If the Addresses field is empty, traffic will be identified
	// solely based on the destination port. In such scenarios, the port on
	// which the service is being accessed must not be shared by any other
	// service in the mesh. In other words, the sidecar will behave as a
	// simple TCP proxy, forwarding incoming traffic on a specified port to
	// the specified destination endpoint IP/host. Unix domain socket
	// addresses are not supported in this field.
	Addresses []string `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
	// The ports associated with the external service. If the
	// Endpoints are Unix domain socket addresses, there must be exactly one
	// port.
	Ports []*Port `protobuf:"bytes,3,rep,name=ports,proto3" json:"ports,omitempty"`
	// Specify whether the service should be considered external to the mesh
	// or part of the mesh.
	Location ServiceEntry_Location `protobuf:"varint,4,opt,name=location,proto3,enum=istio.networking.v1beta1.ServiceEntry_Location" json:"location,omitempty"`
	// Service discovery mode for the hosts. Care must be taken
	// when setting the resolution mode to NONE for a TCP port without
	// accompanying IP addresses. In such cases, traffic to any IP on
	// said port will be allowed (i.e. `0.0.0.0:<port>`).
	Resolution ServiceEntry_Resolution `protobuf:"varint,5,opt,name=resolution,proto3,enum=istio.networking.v1beta1.ServiceEntry_Resolution" json:"resolution,omitempty"`
	// One or more endpoints associated with the service. Only one of
	// `endpoints` or `workloadSelector` can be specified.
	Endpoints []*WorkloadEntry `protobuf:"bytes,6,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	// Applicable only for MESH_INTERNAL services. Only one of
	// `endpoints` or `workloadSelector` can be specified. Selects one
	// or more Kubernetes pods or VM workloads (specified using
	// `WorkloadEntry`) based on their labels. The `WorkloadEntry` object
	// representing the VMs should be defined in the same namespace as
	// the ServiceEntry.
	WorkloadSelector *WorkloadSelector `protobuf:"bytes,9,opt,name=workload_selector,json=workloadSelector,proto3" json:"workload_selector,omitempty"`
	// A list of namespaces to which this service is exported. Exporting a service
	// allows it to be used by sidecars, gateways and virtual services defined in
	// other namespaces. This feature provides a mechanism for service owners
	// and mesh administrators to control the visibility of services across
	// namespace boundaries.
	//
	// If no namespaces are specified then the service is exported to all
	// namespaces by default.
	//
	// The value "." is reserved and defines an export to the same namespace that
	// the service is declared in. Similarly the value "*" is reserved and
	// defines an export to all namespaces.
	//
	// For a Kubernetes Service, the equivalent effect can be achieved by setting
	// the annotation "networking.istio.io/exportTo" to a comma-separated list
	// of namespace names.
	//
	// NOTE: in the current release, the `exportTo` value is restricted to
	// "." or "*" (i.e., the current namespace or all namespaces).
	ExportTo []string `protobuf:"bytes,7,rep,name=export_to,json=exportTo,proto3" json:"export_to,omitempty"`
	// If specified, the proxy will verify that the server certificate's
	// subject alternate name matches one of the specified values.
	//
	// NOTE: When using the workloadEntry with workloadSelectors, the
	// service account specified in the workloadEntry will also be used
	// to derive the additional subject alternate names that should be
	// verified.
	SubjectAltNames      []string `protobuf:"bytes,8,rep,name=subject_alt_names,json=subjectAltNames,proto3" json:"subject_alt_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceEntry) Reset()         { *m = ServiceEntry{} }
func (m *ServiceEntry) String() string { return proto.CompactTextString(m) }
func (*ServiceEntry) ProtoMessage()    {}
func (*ServiceEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_be685d9fa1e5ca12, []int{0}
}

func (m *ServiceEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceEntry.Unmarshal(m, b)
}
func (m *ServiceEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceEntry.Marshal(b, m, deterministic)
}
func (m *ServiceEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceEntry.Merge(m, src)
}
func (m *ServiceEntry) XXX_Size() int {
	return xxx_messageInfo_ServiceEntry.Size(m)
}
func (m *ServiceEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceEntry proto.InternalMessageInfo

func (m *ServiceEntry) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *ServiceEntry) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ServiceEntry) GetPorts() []*Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *ServiceEntry) GetLocation() ServiceEntry_Location {
	if m != nil {
		return m.Location
	}
	return ServiceEntry_MESH_EXTERNAL
}

func (m *ServiceEntry) GetResolution() ServiceEntry_Resolution {
	if m != nil {
		return m.Resolution
	}
	return ServiceEntry_NONE
}

func (m *ServiceEntry) GetEndpoints() []*WorkloadEntry {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ServiceEntry) GetWorkloadSelector() *WorkloadSelector {
	if m != nil {
		return m.WorkloadSelector
	}
	return nil
}

func (m *ServiceEntry) GetExportTo() []string {
	if m != nil {
		return m.ExportTo
	}
	return nil
}

func (m *ServiceEntry) GetSubjectAltNames() []string {
	if m != nil {
		return m.SubjectAltNames
	}
	return nil
}

func init() {
	proto.RegisterEnum("istio.networking.v1beta1.ServiceEntry_Location", ServiceEntry_Location_name, ServiceEntry_Location_value)
	proto.RegisterEnum("istio.networking.v1beta1.ServiceEntry_Resolution", ServiceEntry_Resolution_name, ServiceEntry_Resolution_value)
	proto.RegisterType((*ServiceEntry)(nil), "istio.networking.v1beta1.ServiceEntry")
}

func init() {
	proto.RegisterFile("networking/v1beta1/service_entry.proto", fileDescriptor_be685d9fa1e5ca12)
}

var fileDescriptor_be685d9fa1e5ca12 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0x9b, 0x38, 0x49, 0xed, 0xdb, 0xef, 0x03, 0x67, 0x56, 0xa6, 0x20, 0x6a, 0xb2, 0xa0,
	0x51, 0x91, 0x6c, 0x12, 0x56, 0x2c, 0x53, 0xb0, 0x44, 0x45, 0x31, 0xc8, 0x8e, 0x28, 0x62, 0x63,
	0x4d, 0xec, 0x4b, 0x3a, 0x60, 0x7c, 0xa3, 0x99, 0x69, 0x42, 0xdf, 0x96, 0x37, 0xe0, 0x15, 0x90,
	0xff, 0x34, 0x89, 0xa0, 0x01, 0x96, 0x3e, 0x73, 0xce, 0xcf, 0x77, 0xce, 0xcc, 0xc0, 0xe3, 0x02,
	0xf5, 0x8a, 0xe4, 0x17, 0x51, 0xcc, 0xfd, 0xe5, 0x68, 0x86, 0x9a, 0x8f, 0x7c, 0x85, 0x72, 0x29,
	0x52, 0x4c, 0xb0, 0xd0, 0xf2, 0xda, 0x5b, 0x48, 0xd2, 0xc4, 0x1c, 0xa1, 0xb4, 0x20, 0x6f, 0xe3,
	0xf6, 0x1a, 0xf7, 0xe1, 0xd1, 0x9c, 0x68, 0x9e, 0xa3, 0xcf, 0x17, 0xc2, 0xff, 0x24, 0x30, 0xcf,
	0x92, 0x19, 0x5e, 0xf2, 0xa5, 0x20, 0x59, 0x47, 0x0f, 0xdd, 0x5b, 0x7e, 0x31, 0xe7, 0x1a, 0x57,
	0xfc, 0xfa, 0x0f, 0x0e, 0x25, 0x32, 0x4c, 0xf9, 0x0d, 0xe3, 0xf8, 0x16, 0x47, 0xf9, 0x9d, 0x13,
	0xcf, 0xb6, 0xe7, 0x1c, 0xfc, 0xe8, 0xc0, 0x7f, 0x71, 0x3d, 0x7f, 0x50, 0xca, 0xec, 0x1e, 0x74,
	0x2f, 0x49, 0x69, 0xe5, 0xb4, 0x5c, 0x63, 0x68, 0x9d, 0x1a, 0xdf, 0x27, 0xed, 0xa8, 0x56, 0xd8,
	0x03, 0xb0, 0x78, 0x96, 0x49, 0x54, 0x0a, 0x95, 0xd3, 0x2e, 0x97, 0xa3, 0x8d, 0xc0, 0x9e, 0x43,
	0x77, 0x41, 0x52, 0x2b, 0xc7, 0x70, 0x8d, 0xe1, 0xc1, 0xf8, 0xa1, 0xb7, 0xab, 0x01, 0xef, 0x1d,
	0x49, 0xdd, 0x80, 0xab, 0x04, 0x7b, 0x0d, 0x66, 0x4e, 0x29, 0xd7, 0x82, 0x0a, 0xa7, 0xe3, 0xb6,
	0x86, 0x77, 0xc6, 0xfe, 0xee, 0xf4, 0xf6, 0xb4, 0xde, 0x79, 0x13, 0x8b, 0xd6, 0x00, 0xf6, 0x1e,
	0x40, 0xa2, 0xa2, 0xfc, 0xaa, 0xc2, 0x75, 0x2b, 0xdc, 0xe8, 0x1f, 0x71, 0xd1, 0x3a, 0x58, 0xcf,
	0xb7, 0x45, 0x62, 0x01, 0x58, 0x58, 0x64, 0x0b, 0x12, 0x85, 0x56, 0x4e, 0xaf, 0xda, 0xe3, 0xf1,
	0x6e, 0xec, 0x45, 0x53, 0x76, 0xc5, 0x8d, 0x36, 0x49, 0x76, 0x01, 0xfd, 0xf5, 0x41, 0x28, 0xcc,
	0x31, 0xd5, 0x24, 0x1d, 0xcb, 0x6d, 0x0d, 0x0f, 0xc6, 0x27, 0x7f, 0xc7, 0xc5, 0x4d, 0x22, 0xb2,
	0x57, 0xbf, 0x28, 0xec, 0x3e, 0x58, 0xf8, 0xad, 0xec, 0x33, 0xd1, 0xe4, 0xec, 0x57, 0xa7, 0x63,
	0xd6, 0xc2, 0x94, 0xd8, 0x09, 0xf4, 0xd5, 0xd5, 0xec, 0x33, 0xa6, 0x3a, 0xe1, 0xb9, 0x4e, 0x0a,
	0xfe, 0x15, 0x95, 0x63, 0x56, 0xa6, 0xbb, 0xcd, 0xc2, 0x24, 0xd7, 0x61, 0x29, 0x0f, 0x9e, 0x82,
	0x79, 0x53, 0x2b, 0xeb, 0xc3, 0xff, 0x6f, 0x82, 0xf8, 0x55, 0x12, 0x7c, 0x98, 0x06, 0x51, 0x38,
	0x39, 0xb7, 0xf7, 0xd6, 0xd2, 0x59, 0xd8, 0x48, 0xad, 0xc1, 0x13, 0x80, 0x4d, 0x73, 0xcc, 0x84,
	0x4e, 0xf8, 0x36, 0x0c, 0xec, 0x3d, 0x06, 0xd0, 0x8b, 0xa7, 0x93, 0xe9, 0xd9, 0x0b, 0xbb, 0xc5,
	0xf6, 0xc1, 0x78, 0x19, 0xc6, 0x76, 0xfb, 0xf4, 0xd1, 0xc7, 0xa3, 0x7a, 0x9b, 0x82, 0xaa, 0x37,
	0xf0, 0xfb, 0x4d, 0x9d, 0xf5, 0xaa, 0xbb, 0xf9, 0xec, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3,
	0x41, 0x09, 0xb6, 0x6d, 0x03, 0x00, 0x00,
}
