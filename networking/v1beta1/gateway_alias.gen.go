// Code generated by protoc-gen-alias. DO NOT EDIT.
package v1beta1

import "istio.io/api/networking/v1alpha3"

// Gateway describes a load balancer operating at the edge of the mesh
// receiving incoming or outgoing HTTP/TCP connections.
//
// <!-- crd generation tags
// +cue-gen:Gateway:groupName:networking.istio.io
// +cue-gen:Gateway:versions:v1beta1,v1alpha3,v1
// +cue-gen:Gateway:annotations:helm.sh/resource-policy=keep
// +cue-gen:Gateway:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:Gateway:subresource:status
// +cue-gen:Gateway:scope:Namespaced
// +cue-gen:Gateway:resource:categories=istio-io,networking-istio-io,shortNames=gw
// +cue-gen:Gateway:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.istio.io/v1alpha3
// +genclient
// +k8s:deepcopy-gen=true
// -->
type Gateway = v1alpha3.Gateway

// `Server` describes the properties of the proxy on a given load balancer
// port. For example,
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: Gateway
// metadata:
//
//	name: my-ingress
//
// spec:
//
//	selector:
//	  app: my-ingressgateway
//	servers:
//	- port:
//	    number: 80
//	    name: http2
//	    protocol: HTTP2
//	  hosts:
//	  - "*"
//
// ```
//
// # Another example
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: Gateway
// metadata:
//
//	name: my-tcp-ingress
//
// spec:
//
//	selector:
//	  app: my-tcp-ingressgateway
//	servers:
//	- port:
//	    number: 27018
//	    name: mongo
//	    protocol: MONGO
//	  hosts:
//	  - "*"
//
// ```
//
// # The following is an example of TLS configuration for port 443
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: Gateway
// metadata:
//
//	name: my-tls-ingress
//
// spec:
//
//	selector:
//	  app: my-tls-ingressgateway
//	servers:
//	- port:
//	    number: 443
//	    name: https
//	    protocol: HTTPS
//	  hosts:
//	  - "*"
//	  tls:
//	    mode: SIMPLE
//	    credentialName: tls-cert
//
// ```
type Server = v1alpha3.Server

// Port describes the properties of a specific port of a service.
type Port = v1alpha3.Port

// +kubebuilder:validation:XValidation:message="only one of credentialNames or tlsCertificates can be set",rule="oneof(self.tlsCertificates, self.credentialNames)"
// +kubebuilder:validation:XValidation:message="only one of credentialName or credentialNames can be set",rule="oneof(self.credentialName, self.credentialNames)"
// +kubebuilder:validation:XValidation:message="only one of credentialName or tlsCertificates can be set",rule="oneof(self.credentialNames, self.tlsCertificates)"
type ServerTLSSettings = v1alpha3.ServerTLSSettings

// TLSCertificate describes the server's TLS certificate.
type ServerTLSSettings_TLSCertificate = v1alpha3.ServerTLSSettings_TLSCertificate

// TLS modes enforced by the proxy
type ServerTLSSettings_TLSmode = v1alpha3.ServerTLSSettings_TLSmode

// The SNI string presented by the client will be used as the
// match criterion in a VirtualService TLS route to determine
// the destination service from the service registry.
const ServerTLSSettings_PASSTHROUGH ServerTLSSettings_TLSmode = v1alpha3.ServerTLSSettings_PASSTHROUGH

// Secure connections with standard TLS semantics. In this mode
// client certificate is not requested during handshake.
const ServerTLSSettings_SIMPLE ServerTLSSettings_TLSmode = v1alpha3.ServerTLSSettings_SIMPLE

// Secure connections to the downstream using mutual TLS by
// presenting server certificates for authentication.
// A client certificate will also be requested during the handshake and
// at least one valid certificate is required to be sent by the client.
const ServerTLSSettings_MUTUAL ServerTLSSettings_TLSmode = v1alpha3.ServerTLSSettings_MUTUAL

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
const ServerTLSSettings_AUTO_PASSTHROUGH ServerTLSSettings_TLSmode = v1alpha3.ServerTLSSettings_AUTO_PASSTHROUGH

// Secure connections from the downstream using mutual TLS by
// presenting server certificates for authentication.  Compared
// to Mutual mode, this mode uses certificates, representing
// gateway workload identity, generated automatically by Istio
// for mTLS authentication. When this mode is used, all other
// fields in `TLSOptions` should be empty.
const ServerTLSSettings_ISTIO_MUTUAL ServerTLSSettings_TLSmode = v1alpha3.ServerTLSSettings_ISTIO_MUTUAL

// Similar to MUTUAL mode, except that the client certificate
// is optional. Unlike SIMPLE mode, A client certificate will
// still be explicitly requested during handshake, but the client
// is not required to send a certificate. If a client certificate
// is presented, it will be validated. ca_certificates should
// be specified for validating client certificates.
const ServerTLSSettings_OPTIONAL_MUTUAL ServerTLSSettings_TLSmode = v1alpha3.ServerTLSSettings_OPTIONAL_MUTUAL

// TLS protocol versions.
type ServerTLSSettings_TLSProtocol = v1alpha3.ServerTLSSettings_TLSProtocol

// Automatically choose the optimal TLS version.
const ServerTLSSettings_TLS_AUTO ServerTLSSettings_TLSProtocol = v1alpha3.ServerTLSSettings_TLS_AUTO

// TLS version 1.0
const ServerTLSSettings_TLSV1_0 ServerTLSSettings_TLSProtocol = v1alpha3.ServerTLSSettings_TLSV1_0

// TLS version 1.1
const ServerTLSSettings_TLSV1_1 ServerTLSSettings_TLSProtocol = v1alpha3.ServerTLSSettings_TLSV1_1

// TLS version 1.2
const ServerTLSSettings_TLSV1_2 ServerTLSSettings_TLSProtocol = v1alpha3.ServerTLSSettings_TLSV1_2

// TLS version 1.3
const ServerTLSSettings_TLSV1_3 ServerTLSSettings_TLSProtocol = v1alpha3.ServerTLSSettings_TLSV1_3
