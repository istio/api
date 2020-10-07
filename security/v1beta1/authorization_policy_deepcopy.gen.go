// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security/v1beta1/authorization_policy.proto

// Istio Authorization Policy enables access control on workloads in the mesh.
//
// Authorization policy supports both allow and deny policies. When allow and
// deny policies are used for a workload at the same time, the deny policies are
// evaluated first. The evaluation is determined by the following rules:
//
// 1. If there are any DENY policies that match the request, deny the request.
// 2. If there are no ALLOW policies for the workload, allow the request.
// 3. If any of the ALLOW policies match the request, allow the request.
// 4. Deny the request.
//
// Istio Authorization Policy also supports the AUDIT action to decide whether to log requests.
// AUDIT policies do not affect whether requests are allowed or denied to the workload.
// Requests will be allowed or denied based solely on ALLOW and DENY policies.
//
// A request will be internally marked that it should be audited if there is an AUDIT policy on the workload that matches the request.
// A separate plugin must be configured and enabled to actually fulfill the audit decision and complete the audit behavior.
// The request will not be audited if there are no such supporting plugins enabled.
// Currently, the only supported plugin is the [Stackdriver](https://istio.io/latest/docs/reference/config/proxy_extensions/stackdriver/) plugin.
//
// Here is an example of Istio Authorization Policy:
//
// It sets the `action` to "ALLOW" to create an allow policy. The default action is "ALLOW"
// but it is useful to be explicit in the policy.
//
// It allows requests from:
//
// - service account "cluster.local/ns/default/sa/sleep" or
// - namespace "test"
//
// to access the workload with:
//
// - "GET" method at paths of prefix "/info" or,
// - "POST" method at path "/data".
//
// when the request has a valid JWT token issued by "https://accounts.google.com".
//
// Any other requests will be denied.
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: httpbin
//  namespace: foo
// spec:
//  action: ALLOW
//  rules:
//  - from:
//    - source:
//        principals: ["cluster.local/ns/default/sa/sleep"]
//    - source:
//        namespaces: ["test"]
//    to:
//    - operation:
//        methods: ["GET"]
//        paths: ["/info*"]
//    - operation:
//        methods: ["POST"]
//        paths: ["/data"]
//    when:
//    - key: request.auth.claims[iss]
//      values: ["https://accounts.google.com"]
// ```
//
// The following is another example that sets `action` to "DENY" to create a deny policy.
// It denies requests from the "dev" namespace to the "POST" method on all workloads
// in the "foo" namespace.
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: httpbin
//  namespace: foo
// spec:
//  action: DENY
//  rules:
//  - from:
//    - source:
//        namespaces: ["dev"]
//    to:
//    - operation:
//        methods: ["POST"]
// ```
//
// The following authorization policy sets the `action` to "AUDIT". It will audit any GET requests to the path with the
// prefix "/user/profile".
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//   namespace: ns1
//   name: anyname
// spec:
//   selector:
//     matchLabels:
//       app: myapi
//   action: audit
//   rules:
//   - to:
//     - operation:
//         methods: ["GET"]
//         paths: ["/user/profile/*"]
// ```
//
// Authorization Policy scope (target) is determined by "metadata/namespace" and
// an optional "selector".
//
// - "metadata/namespace" tells which namespace the policy applies. If set to root
// namespace, the policy applies to all namespaces in a mesh.
// - workload "selector" can be used to further restrict where a policy applies.
//
// For example,
//
// The following authorization policy applies to workloads containing label
// "app: httpbin" in namespace bar.
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: policy
//  namespace: bar
// spec:
//  selector:
//    matchLabels:
//      app: httpbin
// ```
//
// The following authorization policy applies to all workloads in namespace foo.
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: policy
//  namespace: foo
// spec:
//   {}
// ```
//
// The following authorization policy applies to workloads containing label
// "version: v1" in all namespaces in the mesh. (Assuming the root namespace is
// configured to "istio-config").
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: policy
//  namespace: istio-config
// spec:
//  selector:
//    matchLabels:
//      version: v1
// ```
//
// The following authorization policy applies to ingress gateway to enable the external authorization for all HTTP and
// TCP requests.
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: ext-auth
//  namespace: istio-system
// spec:
//  selector:
//    matchLabels:
//      app: istio-ingressgateway
//  action: EXTERNAL
//  external:
//    http:
//      server: "grpc://ext-authz.foo.svc.cluster.local:9000"
//    tcp:
//      server: "grpc://ext-authz.foo.svc.cluster.local:9000"
//  rules:
//  # Empty rules to always trigger the check request.
//  - {}
// ```
//
// The following authorization policy applies to ingress gateway to enable the external authorization for HTTP requests
// if the request path has prefix "/data/".
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//  name: ext-auth
//  namespace: istio-system
// spec:
//  selector:
//    matchLabels:
//      app: istio-ingressgateway
//  action: EXTERNAL
//  external:
//    http:
//      server: "http://ext-authz.foo.svc.cluster.local:8000"
//  rules:
//  # Specify rules to conditionally trigger the check request only if the path has prefix "/data/".
//  - to:
//    - operation:
//        paths: ["/data/*"]
// ```
//
// The EXTERNAL action currently uses a default ext_authz configuration and you can customize it with EnvoyFilter.
// For example, the following EnvoyFilter patches the default configuration to change timeout to 2s and also include the
// extra "X-foo" header in the check request. We will further improve the user experience here by allowing to directly
// customize the configurations right in the authorization policy so that you don't need the EnvoyFilter in
// most use cases).
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: EnvoyFilter
// metadata:
//   name: ext-authz
//   namespace: istio-system
// spec:
//   workloadSelector:
//     labels:
//       app: istio-ingressgateway
//   configPatches:
//   - applyTo: HTTP_FILTER
//     match:
//       context: ANY
//       listener:
//         filterChain:
//           filter:
//             name: envoy.http_connection_manager
//             subFilter:
//               name: envoy.filters.http.ext_authz
//     patch:
//       operation: MERGE
//       value:
//         name: envoy.filters.http.ext_authz
//         typed_config:
//           '@type': type.googleapis.com/envoy.extensions.filters.http.ext_authz.v3.ExtAuthz
//           http_service:
//             server_uri:
//               timeout: 2s
//           authorization_request:
//             allowed_headers:
//               patterns:
//               - exact: X-foo
// ```

package v1beta1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "istio.io/api/type/v1beta1"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using AuthorizationPolicy within kubernetes types, where deepcopy-gen is used.
func (in *AuthorizationPolicy) DeepCopyInto(out *AuthorizationPolicy) {
	p := proto.Clone(in).(*AuthorizationPolicy)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationPolicy. Required by controller-gen.
func (in *AuthorizationPolicy) DeepCopy() *AuthorizationPolicy {
	if in == nil {
		return nil
	}
	out := new(AuthorizationPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationPolicy. Required by controller-gen.
func (in *AuthorizationPolicy) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using External within kubernetes types, where deepcopy-gen is used.
func (in *External) DeepCopyInto(out *External) {
	p := proto.Clone(in).(*External)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new External. Required by controller-gen.
func (in *External) DeepCopy() *External {
	if in == nil {
		return nil
	}
	out := new(External)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new External. Required by controller-gen.
func (in *External) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using External_Config within kubernetes types, where deepcopy-gen is used.
func (in *External_Config) DeepCopyInto(out *External_Config) {
	p := proto.Clone(in).(*External_Config)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new External_Config. Required by controller-gen.
func (in *External_Config) DeepCopy() *External_Config {
	if in == nil {
		return nil
	}
	out := new(External_Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new External_Config. Required by controller-gen.
func (in *External_Config) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Rule within kubernetes types, where deepcopy-gen is used.
func (in *Rule) DeepCopyInto(out *Rule) {
	p := proto.Clone(in).(*Rule)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule. Required by controller-gen.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Rule. Required by controller-gen.
func (in *Rule) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Rule_From within kubernetes types, where deepcopy-gen is used.
func (in *Rule_From) DeepCopyInto(out *Rule_From) {
	p := proto.Clone(in).(*Rule_From)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule_From. Required by controller-gen.
func (in *Rule_From) DeepCopy() *Rule_From {
	if in == nil {
		return nil
	}
	out := new(Rule_From)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Rule_From. Required by controller-gen.
func (in *Rule_From) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Rule_To within kubernetes types, where deepcopy-gen is used.
func (in *Rule_To) DeepCopyInto(out *Rule_To) {
	p := proto.Clone(in).(*Rule_To)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule_To. Required by controller-gen.
func (in *Rule_To) DeepCopy() *Rule_To {
	if in == nil {
		return nil
	}
	out := new(Rule_To)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Rule_To. Required by controller-gen.
func (in *Rule_To) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Source within kubernetes types, where deepcopy-gen is used.
func (in *Source) DeepCopyInto(out *Source) {
	p := proto.Clone(in).(*Source)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source. Required by controller-gen.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Source. Required by controller-gen.
func (in *Source) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Operation within kubernetes types, where deepcopy-gen is used.
func (in *Operation) DeepCopyInto(out *Operation) {
	p := proto.Clone(in).(*Operation)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Operation. Required by controller-gen.
func (in *Operation) DeepCopy() *Operation {
	if in == nil {
		return nil
	}
	out := new(Operation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Operation. Required by controller-gen.
func (in *Operation) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Condition within kubernetes types, where deepcopy-gen is used.
func (in *Condition) DeepCopyInto(out *Condition) {
	p := proto.Clone(in).(*Condition)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition. Required by controller-gen.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Condition. Required by controller-gen.
func (in *Condition) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
