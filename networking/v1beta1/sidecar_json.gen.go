// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1beta1/sidecar.proto

// `Sidecar` describes the configuration of the sidecar proxy that mediates
// inbound and outbound communication to the workload instance it is attached to. By
// default, Istio will program all sidecar proxies in the mesh with the
// necessary configuration required to reach every workload instance in the mesh, as
// well as accept traffic on all the ports associated with the
// workload. The `Sidecar` configuration provides a way to fine tune the set of
// ports, protocols that the proxy will accept when forwarding traffic to
// and from the workload. In addition, it is possible to restrict the set
// of services that the proxy can reach when forwarding outbound traffic
// from workload instances.
//
// Services and configuration in a mesh are organized into one or more
// namespaces (e.g., a Kubernetes namespace or a CF org/space). A `Sidecar`
// configuration in a namespace will apply to one or more workload instances in the same
// namespace, selected using the `workloadSelector` field. In the absence of a
// `workloadSelector`, it will apply to all workload instances in the same
// namespace. When determining the `Sidecar` configuration to be applied to a
// workload instance, preference will be given to the resource with a
// `workloadSelector` that selects this workload instance, over a `Sidecar` configuration
// without any `workloadSelector`.
//
// **NOTE 1**: *_Each namespace can have only one `Sidecar` configuration without any
// `workloadSelector`_*. The behavior of the system is undefined if more
// than one selector-less `Sidecar` configurations exist in a given namespace. The
// behavior of the system is undefined if two or more `Sidecar` configurations
// with a `workloadSelector` select the same workload instance.
//
// **NOTE 2**: *_A `Sidecar` configuration in the `MeshConfig`
// [root namespace](https://istio.io/docs/reference/config/istio.mesh.v1alpha1/#MeshConfig)
// will be applied by default to all namespaces without a `Sidecar`
// configuration_*. This global default `Sidecar` configuration should not have
// any `workloadSelector`.
//
// The example below declares a global default `Sidecar` configuration
// in the root namespace called `istio-config`, that configures
// sidecars in all namespaces to accept both plaintext and Istio
// mutual TLS traffic, and allow egress traffic only to other
// workloads in the same namespace as well as to services in the
// `istio-system` namespace.
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Sidecar
// metadata:
//   name: default
//   namespace: istio-config
// spec:
//   inboundTrafficPolicy:
//     tls:
//       mode: ISTIO_MUTUAL_AND_PLAINTEXT
//   egress:
//   - hosts:
//     - "./*"
//     - "istio-system/*"
//```
//
// The example below declares a `Sidecar` configuration in the
// `prod-us1` namespace that overrides the global default defined
// above, and configures the sidecars in the namespace to only accept
// Istio mutual TLS traffic and allow egress traffic to public
// services in the `prod-us1`, `prod-apis`, and the `istio-system`
// namespaces.
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Sidecar
// metadata:
//   name: default
//   namespace: prod-us1
// spec:
//   inboundTrafficPolicy:
//     tls:
//       mode: ISTIO_MUTUAL
//   egress:
//   - hosts:
//     - "prod-us1/*"
//     - "prod-apis/*"
//     - "istio-system/*"
// ```
//
// The following example declares a `Sidecar`
// configuration in the `prod-us1` namespace for all pods with labels
// `app: ratings` belonging to the `ratings.prod-us1` service.
// The workload accepts inbound HTTP traffic on port 9080 without any
// authentication, and HTTPS traffic on port 9443 with one-way TLS
// termination using custom certificates. _The TLS settings for these
// two ports override the namespace-wide defaults for Istio mutual TLS
// defined previously. Any other auto-generated listener for this
// workload will still obey the namespace-wide defaults_. The traffic
// is then forwarded to the attached workload instance listening on a
// Unix domain socket. In the egress direction, in addition to the
// `istio-system` namespace, the sidecar proxies only HTTP traffic
// bound for port 9080 for services in the `prod-us1` namespace.
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Sidecar
// metadata:
//   name: ratings
//   namespace: prod-us1
// spec:
//   workloadSelector:
//     labels:
//       app: ratings
//   ingress:
//   - port:
//       number: 9080
//       protocol: HTTP
//       name: somename
//     tls:
//       mode: DISABLED # allow plaintext on 9080
//     defaultEndpoint: unix:///var/run/someuds.sock
//   - port:
//       number: 9443
//       protocol: HTTPS
//       name: httpsport
//     tls:
//       mode: SIMPLE # overrides namespace default
//       serverCertificate: /etc/certs/servercert.pem
//       privateKey: /etc/certs/privatekey.pem
//     defaultEndpoint: unix:///var/run/someuds.sock
//   egress:
//   - port:
//       number: 9080
//       protocol: HTTP
//       name: egresshttp
//     hosts:
//     - "prod-us1/*"
//   - hosts:
//     - "istio-system/*"
// ```
//
// If the workload is deployed without IPTables-based traffic capture, the
// `Sidecar` configuration is the only way to configure the ports on the proxy
// attached to the workload instance. The following example declares a `Sidecar`
// configuration in the `prod-us1` namespace for all pods with labels
// `app: productpage` belonging to the `productpage.prod-us1` service. Assuming
// that these pods are deployed without IPtable rules (i.e. the `istio-init`
// container) and the proxy metadata `ISTIO_META_INTERCEPTION_MODE` is set to
// `NONE`, the specification, below, allows such pods to receive HTTP traffic
// on port 9080 and forward it to the application listening on
// `127.0.0.1:8080`. It also allows the application to communicate with a
// backing MySQL database on `127.0.0.1:3306`, that then gets proxied to the
// externally hosted MySQL service at `mysql.foo.com:3306`.
//
// ** NOTE 1**: Since the ingress listener for port 9080 does not
// override the TLS mode, the default inbound traffic policy defined
// by the namespace-wide sidecar, i.e. ISTIO_MUTUAL, will be
// applied to port 80. Clients connecting to this port are expected to
// be other sidecars in the mesh, initiating Istio mutual TLS connection.
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Sidecar
// metadata:
//   name: no-ip-tables
//   namespace: prod-us1
// spec:
//   workloadSelector:
//     labels:
//       app: productpage
//   ingress:
//   - port:
//       number: 9080 # binds to proxy_instance_ip:9080 (0.0.0.0:9080, if no unicast IP is available for the instance)
//       protocol: HTTP
//       name: somename
//     defaultEndpoint: 127.0.0.1:8080
//     captureMode: NONE # not needed if metadata is set for entire proxy
//   egress:
//   - port:
//       number: 3306
//       protocol: MYSQL
//       name: egressmysql
//     captureMode: NONE # not needed if metadata is set for entire proxy
//     bind: 127.0.0.1
//     hosts:
//     - "*/mysql.foo.com"
// ```
//
// And the associated service entry for routing to `mysql.foo.com:3306`
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: external-svc-mysql
//   namespace: ns1
// spec:
//   hosts:
//   - mysql.foo.com
//   ports:
//   - number: 3306
//     name: mysql
//     protocol: MYSQL
//   location: MESH_EXTERNAL
//   resolution: DNS
// ```
//
// It is also possible to mix and match traffic capture modes in a single
// proxy. For example, consider a setup where internal services are on the
// `192.168.0.0/16` subnet. So, IP tables are setup on the VM to capture all
// outbound traffic on `192.168.0.0/16` subnet. Assume that the VM has an
// additional network interface on `172.16.0.0/16` subnet for inbound
// traffic. The following `Sidecar` configuration allows the VM to expose a
// listener on `172.16.1.32:80` (the VM's IP) for traffic arriving from the
// `172.16.0.0/16` subnet.
//
// **NOTE**: The `ISTIO_META_INTERCEPTION_MODE` metadata on the
// proxy in the VM should contain `REDIRECT` or `TPROXY` as its value,
// implying that IP tables based traffic capture is active.
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Sidecar
// metadata:
//   name: partial-ip-tables
//   namespace: prod-us1
// spec:
//   workloadSelector:
//     labels:
//       app: productpage
//   ingress:
//   - bind: 172.16.1.32
//     port:
//       number: 80 # binds to 172.16.1.32:80
//       protocol: HTTP
//       name: somename
//     defaultEndpoint: 127.0.0.1:8080
//     captureMode: NONE
//     tls:
//       mode: DISABLED # allow plaintext on 80
//   egress:
//     # use the system detected defaults
//     # sets up configuration to handle outbound traffic to services
//     # in 192.168.0.0/16 subnet, based on information provided by the
//     # service registry
//   - captureMode: IPTABLES
//     hosts:
//     - "*/*"
// ```
//

package v1beta1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for Sidecar
func (this *Sidecar) MarshalJSON() ([]byte, error) {
	str, err := SidecarMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Sidecar
func (this *Sidecar) UnmarshalJSON(b []byte) error {
	return SidecarUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for IstioIngressListener
func (this *IstioIngressListener) MarshalJSON() ([]byte, error) {
	str, err := SidecarMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for IstioIngressListener
func (this *IstioIngressListener) UnmarshalJSON(b []byte) error {
	return SidecarUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for IstioEgressListener
func (this *IstioEgressListener) MarshalJSON() ([]byte, error) {
	str, err := SidecarMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for IstioEgressListener
func (this *IstioEgressListener) UnmarshalJSON(b []byte) error {
	return SidecarUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for WorkloadSelector
func (this *WorkloadSelector) MarshalJSON() ([]byte, error) {
	str, err := SidecarMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WorkloadSelector
func (this *WorkloadSelector) UnmarshalJSON(b []byte) error {
	return SidecarUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for OutboundTrafficPolicy
func (this *OutboundTrafficPolicy) MarshalJSON() ([]byte, error) {
	str, err := SidecarMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for OutboundTrafficPolicy
func (this *OutboundTrafficPolicy) UnmarshalJSON(b []byte) error {
	return SidecarUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for InboundTrafficPolicy
func (this *InboundTrafficPolicy) MarshalJSON() ([]byte, error) {
	str, err := SidecarMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for InboundTrafficPolicy
func (this *InboundTrafficPolicy) UnmarshalJSON(b []byte) error {
	return SidecarUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	SidecarMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	SidecarUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
