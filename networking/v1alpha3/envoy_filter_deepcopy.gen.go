// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/envoy_filter.proto

// `EnvoyFilter` provides a mechanism to customize the Envoy
// configuration generated by Istio Pilot. Use EnvoyFilter to modify
// values for certain fields, add specific filters, or even add
// entirely new listeners, clusters, etc. This feature must be used
// with care, as incorrect configurations could potentially
// destabilize the entire mesh. Unlike other Istio networking objects,
// EnvoyFilters are additively applied. Any number of EnvoyFilters can
// exist for a given workload in a specific namespace. The order of
// application of these EnvoyFilters is as follows: all EnvoyFilters
// in the config [root
// namespace](https://istio.io/docs/reference/config/istio.mesh.v1alpha1/#MeshConfig),
// followed by all matching EnvoyFilters in the workload's namespace.
//
// **NOTE 1**: Some aspects of this API is deeply tied to the internal
// implementation in Istio networking subsystem as well as Envoy's XDS
// API. While the EnvoyFilter API by itself will maintain backward
// compatibility, any envoy configuration provided through this
// mechanism should be carefully monitored across Istio proxy version
// upgrades, to ensure that deprecated fields are removed and replaced
// appropriately.
//
// **NOTE 2**: When multiple EnvoyFilters are bound to the same
// workload in a given namespace, all patches will be processed
// sequentially in order of creation time.  The behavior is undefined
// if multiple EnvoyFilter configurations conflict with each other.
//
// **NOTE 3**: To apply an EnvoyFilter resource to all workloads
// (sidecars and gateways) in the system, define the resource in the
// config [root
// namespace](https://istio.io/docs/reference/config/istio.mesh.v1alpha1/#MeshConfig),
// without a workloadSelector.
//
// The example below declares a global default EnvoyFilter resource in
// the root namespace called `istio-config`, that adds a custom
// protocol filter on all sidecars in the system, for outbound port
// 9307. The filter should be added before the terminating tcp_proxy
// filter to take effect. In addition, it sets a 30s idle timeout for
// all HTTP connections in both gateways and sidecars.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: EnvoyFilter
// metadata:
//   name: custom-protocol
//   namespace: istio-config # as defined in meshConfig resource.
// spec:
//   configPatches:
//   - applyTo: NETWORK_FILTER
//     match:
//       context: SIDECAR_OUTBOUND # will match outbound listeners in all sidecars
//       listener:
//         portNumber: 9307
//         filterChain:
//           filter:
//             name: "envoy.tcp_proxy"
//     patch:
//       operation: INSERT_BEFORE
//       value:
//         # This is the full filter config including the name and config or typed_config section.
//         name: "envoy.config.filter.network.custom_protocol"
//         config:
//          ...
//   - applyTo: NETWORK_FILTER # http connection manager is a filter in Envoy
//     match:
//       # context omitted so that this applies to both sidecars and gateways
//       listener:
//         filterChain:
//           filter:
//             name: "envoy.http_connection_manager"
//     patch:
//       operation: MERGE
//       value:
//         name: "envoy.http_connection_manager"
//         typed_config:
//           "@type": "type.googleapis.com/envoy.config.filter.network.http_connection_manager.v2.HttpConnectionManager"
//           common_http_protocol_options:
//             idle_timeout: 30s
//```
//
// The following example enables Envoy's Lua filter for all inbound
// HTTP calls arriving at service port 8080 of the reviews service pod
// with labels "app: reviews", in the bookinfo namespace. The lua
// filter calls out to an external service internal.org.net:8888 that
// requires a special cluster definition in envoy. The cluster is also
// added to the sidecar as part of this configuration.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: EnvoyFilter
// metadata:
//   name: reviews-lua
//   namespace: bookinfo
// spec:
//   workloadSelector:
//     labels:
//       app: reviews
//   configPatches:
//     # The first patch adds the lua filter to the listener/http connection manager
//   - applyTo: HTTP_FILTER
//     match:
//       context: SIDECAR_INBOUND
//       listener:
//         portNumber: 8080
//         filterChain:
//           filter:
//             name: "envoy.http_connection_manager"
//             subFilter:
//               name: "envoy.router"
//     patch:
//       operation: INSERT_BEFORE
//       value: # lua filter specification
//         name: envoy.lua
//         typed_config:
//           "@type": "type.googleapis.com/envoy.config.filter.http.lua.v2.Lua"
//           inlineCode: |
//            function envoy_on_request(request_handle)
//              -- Make an HTTP call to an upstream host with the following headers, body, and timeout.
//              local headers, body = request_handle:httpCall(
//               "lua_cluster",
//               {
//                [":method"] = "POST",
//                [":path"] = "/acl",
//                [":authority"] = "internal.org.net"
//               },
//              "authorize call",
//              5000)
//            end
//   # The second patch adds the cluster that is referenced by the lua code
//   # cds match is omitted as a new cluster is being added
//   - applyTo: CLUSTER
//     match:
//       context: SIDECAR_OUTBOUND
//     patch:
//       operation: ADD
//       value: # cluster specification
//         name: "lua_cluster"
//         type: STRICT_DNS
//         connect_timeout: 0.5s
//         lb_policy: ROUND_ROBIN
//         hosts:
//         - socket_address:
//             protocol: TCP
//             address: "internal.org.net"
//             port_value: 8888
//
// ```
//
// The following example overwrites certain fields (HTTP idle timeout
// and X-Forward-For trusted hops) in the HTTP connection manager in a
// listener on the ingress gateway in istio-system namespace for the
// SNI host app.example.com:
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: EnvoyFilter
// metadata:
//   name: hcm-tweaks
//   namespace: istio-system
// spec:
//   workloadSelector:
//     labels:
//       istio: ingressgateway
//   configPatches:
//   - applyTo: NETWORK_FILTER # http connection manager is a filter in Envoy
//     match:
//       context: GATEWAY
//       listener:
//         filterChain:
//           sni: app.example.com
//           filter:
//             name: "envoy.http_connection_manager"
//     patch:
//       operation: MERGE
//       value:
//         common_http_protocol_options:
//           idle_timeout: 30s
//         xff_num_trusted_hops: 5
//```
//
// The following example inserts an optional filter for OpenApi processing
// in the config root namespace and only instantiates the filter for
// a specific workload by providing workload bound configuration.
// Note that you do not need to specify order when overriding configuration.
// `istio.openapi` filter calculates `request.operation` which is used by the
// `istio.stats` filter therefore the order is important.
//
// Create a placeholder for istio.openapi before istio.stats.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: EnvoyFilter
// metadata:
//   name: openapi
//   namespace: istio-system
// spec:
//   configPatches:
//   - applyTo: HTTP_FILTER
//     match:
//       context: SIDECAR_INBOUND
//         filterChain:
//           filter:
//             name: "envoy.http_connection_manager" # can be elided
//             subFilter:
//               name: "istio.stats"
//     patch:
//       operation: INSERT_BEFORE
//       use_if_overriden: true # allows for incomplete specification.
//       value: # partial filter specification, only name is required.
//         name: istio.openapi
// ```
//
// Provide full configuration for istio.openapi in the workload namespace.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: EnvoyFilter
// metadata:
//   name: mysvc-openapi
//   namespace: myns
// spec:
//   workloadSelector:
//     labels:
//       app: mysvc
//   configPatches:
//   - applyTo: HTTP_FILTER
//     match:
//       context: SIDECAR_INBOUND
//         filterChain:
//           filter:
//             name: "envoy.http_connection_manager" # can be elided
//             subFilter:
//               name: "istio.openapi" # replacing previous placeholder by name.
//     patch:
//       operation: REPLACE
//       value:
//         name: istio.openapi
//         # Full configuration must be specified here.
// ```
//

package v1alpha3

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using EnvoyFilter within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter) DeepCopyInto(out *EnvoyFilter) {
	p := proto.Clone(in).(*EnvoyFilter)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter. Required by controller-gen.
func (in *EnvoyFilter) DeepCopy() *EnvoyFilter {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_ProxyMatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_ProxyMatch) DeepCopyInto(out *EnvoyFilter_ProxyMatch) {
	p := proto.Clone(in).(*EnvoyFilter_ProxyMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_ProxyMatch. Required by controller-gen.
func (in *EnvoyFilter_ProxyMatch) DeepCopy() *EnvoyFilter_ProxyMatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_ProxyMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_ClusterMatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_ClusterMatch) DeepCopyInto(out *EnvoyFilter_ClusterMatch) {
	p := proto.Clone(in).(*EnvoyFilter_ClusterMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_ClusterMatch. Required by controller-gen.
func (in *EnvoyFilter_ClusterMatch) DeepCopy() *EnvoyFilter_ClusterMatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_ClusterMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_RouteConfigurationMatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_RouteConfigurationMatch) DeepCopyInto(out *EnvoyFilter_RouteConfigurationMatch) {
	p := proto.Clone(in).(*EnvoyFilter_RouteConfigurationMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_RouteConfigurationMatch. Required by controller-gen.
func (in *EnvoyFilter_RouteConfigurationMatch) DeepCopy() *EnvoyFilter_RouteConfigurationMatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_RouteConfigurationMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_RouteConfigurationMatch_RouteMatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_RouteConfigurationMatch_RouteMatch) DeepCopyInto(out *EnvoyFilter_RouteConfigurationMatch_RouteMatch) {
	p := proto.Clone(in).(*EnvoyFilter_RouteConfigurationMatch_RouteMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_RouteConfigurationMatch_RouteMatch. Required by controller-gen.
func (in *EnvoyFilter_RouteConfigurationMatch_RouteMatch) DeepCopy() *EnvoyFilter_RouteConfigurationMatch_RouteMatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_RouteConfigurationMatch_RouteMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_RouteConfigurationMatch_VirtualHostMatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_RouteConfigurationMatch_VirtualHostMatch) DeepCopyInto(out *EnvoyFilter_RouteConfigurationMatch_VirtualHostMatch) {
	p := proto.Clone(in).(*EnvoyFilter_RouteConfigurationMatch_VirtualHostMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_RouteConfigurationMatch_VirtualHostMatch. Required by controller-gen.
func (in *EnvoyFilter_RouteConfigurationMatch_VirtualHostMatch) DeepCopy() *EnvoyFilter_RouteConfigurationMatch_VirtualHostMatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_RouteConfigurationMatch_VirtualHostMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_ListenerMatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_ListenerMatch) DeepCopyInto(out *EnvoyFilter_ListenerMatch) {
	p := proto.Clone(in).(*EnvoyFilter_ListenerMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_ListenerMatch. Required by controller-gen.
func (in *EnvoyFilter_ListenerMatch) DeepCopy() *EnvoyFilter_ListenerMatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_ListenerMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_ListenerMatch_FilterChainMatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_ListenerMatch_FilterChainMatch) DeepCopyInto(out *EnvoyFilter_ListenerMatch_FilterChainMatch) {
	p := proto.Clone(in).(*EnvoyFilter_ListenerMatch_FilterChainMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_ListenerMatch_FilterChainMatch. Required by controller-gen.
func (in *EnvoyFilter_ListenerMatch_FilterChainMatch) DeepCopy() *EnvoyFilter_ListenerMatch_FilterChainMatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_ListenerMatch_FilterChainMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_ListenerMatch_FilterMatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_ListenerMatch_FilterMatch) DeepCopyInto(out *EnvoyFilter_ListenerMatch_FilterMatch) {
	p := proto.Clone(in).(*EnvoyFilter_ListenerMatch_FilterMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_ListenerMatch_FilterMatch. Required by controller-gen.
func (in *EnvoyFilter_ListenerMatch_FilterMatch) DeepCopy() *EnvoyFilter_ListenerMatch_FilterMatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_ListenerMatch_FilterMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_ListenerMatch_SubFilterMatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_ListenerMatch_SubFilterMatch) DeepCopyInto(out *EnvoyFilter_ListenerMatch_SubFilterMatch) {
	p := proto.Clone(in).(*EnvoyFilter_ListenerMatch_SubFilterMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_ListenerMatch_SubFilterMatch. Required by controller-gen.
func (in *EnvoyFilter_ListenerMatch_SubFilterMatch) DeepCopy() *EnvoyFilter_ListenerMatch_SubFilterMatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_ListenerMatch_SubFilterMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_Patch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_Patch) DeepCopyInto(out *EnvoyFilter_Patch) {
	p := proto.Clone(in).(*EnvoyFilter_Patch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_Patch. Required by controller-gen.
func (in *EnvoyFilter_Patch) DeepCopy() *EnvoyFilter_Patch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_Patch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_EnvoyConfigObjectMatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_EnvoyConfigObjectMatch) DeepCopyInto(out *EnvoyFilter_EnvoyConfigObjectMatch) {
	p := proto.Clone(in).(*EnvoyFilter_EnvoyConfigObjectMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_EnvoyConfigObjectMatch. Required by controller-gen.
func (in *EnvoyFilter_EnvoyConfigObjectMatch) DeepCopy() *EnvoyFilter_EnvoyConfigObjectMatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_EnvoyConfigObjectMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using EnvoyFilter_EnvoyConfigObjectPatch within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyFilter_EnvoyConfigObjectPatch) DeepCopyInto(out *EnvoyFilter_EnvoyConfigObjectPatch) {
	p := proto.Clone(in).(*EnvoyFilter_EnvoyConfigObjectPatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyFilter_EnvoyConfigObjectPatch. Required by controller-gen.
func (in *EnvoyFilter_EnvoyConfigObjectPatch) DeepCopy() *EnvoyFilter_EnvoyConfigObjectPatch {
	if in == nil {
		return nil
	}
	out := new(EnvoyFilter_EnvoyConfigObjectPatch)
	in.DeepCopyInto(out)
	return out
}
