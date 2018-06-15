# Protocol Documentation
<a name="top"/>

## Table of Contents

- [mesh/v1alpha1/config.proto](#mesh/v1alpha1/config.proto)
    - [MeshConfig](#istio.mesh.v1alpha1.MeshConfig)
    - [MeshConfig.OutboundTrafficPolicy](#istio.mesh.v1alpha1.MeshConfig.OutboundTrafficPolicy)
    - [ProxyConfig](#istio.mesh.v1alpha1.ProxyConfig)
  
    - [AuthenticationPolicy](#istio.mesh.v1alpha1.AuthenticationPolicy)
    - [MeshConfig.AuthPolicy](#istio.mesh.v1alpha1.MeshConfig.AuthPolicy)
    - [MeshConfig.IngressControllerMode](#istio.mesh.v1alpha1.MeshConfig.IngressControllerMode)
    - [ProxyConfig.InboundInterceptionMode](#istio.mesh.v1alpha1.ProxyConfig.InboundInterceptionMode)
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="mesh/v1alpha1/config.proto"/>
<p align="right"><a href="#top">Top</a></p>

## mesh/v1alpha1/config.proto



<a name="istio.mesh.v1alpha1.MeshConfig"/>

### MeshConfig
MeshConfig defines mesh-wide variables shared by all Envoy instances in the
Istio service mesh.

NOTE: This configuration type should be used for the low-level global
configuration, such as component addresses and port numbers. It should not
be used for the features of the mesh that can be scoped by service or by
namespace. Some of the fields in the mesh config are going to be deprecated
and replaced with several individual configuration types (for example,
tracing configuration).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mixer_check_server | [string](#string) |  | Address of the server that will be used by the proxies for policy check calls. By using different names for mixerCheckServer and mixerReportServer, it is possible to have one set of mixer servers handle policy check calls while another set of mixer servers handle telemetry calls.  NOTE: Omitting mixerCheckServer while specifying mixerReportServer is equivalent to setting disablePolicyChecks to true. |
| mixer_report_server | [string](#string) |  | Address of the server that will be used by the proxies for policy report calls. |
| disable_policy_checks | [bool](#bool) |  | Disable policy checks by the mixer service. Default is false, i.e. mixer policy check is enabled by default. |
| proxy_listen_port | [int32](#int32) |  | Port on which Envoy should listen for incoming connections from other services. |
| proxy_http_port | [int32](#int32) |  | Port on which Envoy should listen for HTTP PROXY requests if set. |
| connect_timeout | [google.protobuf.Duration](#google.protobuf.Duration) |  | Connection timeout used by Envoy. (MUST BE &gt;=1ms) |
| ingress_class | [string](#string) |  | Class of ingress resources to be processed by Istio ingress controller. This corresponds to the value of &#34;kubernetes.io/ingress.class&#34; annotation. |
| ingress_service | [string](#string) |  | Name of the kubernetes service used for the istio ingress controller. |
| ingress_controller_mode | [MeshConfig.IngressControllerMode](#istio.mesh.v1alpha1.MeshConfig.IngressControllerMode) |  | Defines whether to use Istio ingress controller for annotated or all ingress resources. |
| auth_policy | [MeshConfig.AuthPolicy](#istio.mesh.v1alpha1.MeshConfig.AuthPolicy) |  | Authentication policy defines the global switch to control authentication for Envoy-to-Envoy communication. Use authentication_policy instead. |
| rds_refresh_delay | [google.protobuf.Duration](#google.protobuf.Duration) |  | Polling interval for RDS (MUST BE &gt;=1ms) |
| enable_tracing | [bool](#bool) |  | Flag to control generation of trace spans and request IDs. Requires a trace span collector defined in the proxy configuration. |
| access_log_file | [string](#string) |  | File address for the proxy access log (e.g. /dev/stdout). Empty value disables access logging. |
| default_config | [ProxyConfig](#istio.mesh.v1alpha1.ProxyConfig) |  | Default proxy config used by the proxy injection mechanism operating in the mesh (e.g. Kubernetes admission controller) In case of Kubernetes, the proxy config is applied once during the injection process, and remain constant for the duration of the pod. The rest of the mesh config can be changed at runtime and config gets distributed dynamically. |
| mixer_address | [string](#string) |  | DEPRECATED. Mixer address. This option will be removed soon. Please use mixer_check and mixer_report. |
| outbound_traffic_policy | [MeshConfig.OutboundTrafficPolicy](#istio.mesh.v1alpha1.MeshConfig.OutboundTrafficPolicy) |  | Set the default behavior of the sidecar for handling outbound traffic from the application. While the default mode should work out of the box, if your application uses one or more external services that are not known apriori, setting the policy to ALLOW_ANY will cause the sidecars to route traffic to the any requested destination. Users are strongly encouraged to use ServiceEntries to explicitly declare any external dependencies, instead of using allow_any. |






<a name="istio.mesh.v1alpha1.MeshConfig.OutboundTrafficPolicy"/>

### MeshConfig.OutboundTrafficPolicy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mode | [MeshConfig.OutboundTrafficPolicy.Mode](#istio.mesh.v1alpha1.MeshConfig.OutboundTrafficPolicy.Mode) |  |  |






<a name="istio.mesh.v1alpha1.ProxyConfig"/>

### ProxyConfig
ProxyConfig defines variables for individual Envoy instances.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config_path | [string](#string) |  | Path to the generated configuration file directory. Proxy agent generates the actual configuration and stores it in this directory. |
| binary_path | [string](#string) |  | Path to the proxy binary |
| service_cluster | [string](#string) |  | Service cluster defines the name for the service_cluster that is shared by all Envoy instances. This setting corresponds to _--service-cluster_ flag in Envoy. In a typical Envoy deployment, the _service-cluster_ flag is used to identify the caller, for source-based routing scenarios.  Since Istio does not assign a local service/service version to each Envoy instance, the name is same for all of them. However, the source/caller&#39;s identity (e.g., IP address) is encoded in the _--service-node_ flag when launching Envoy. When the RDS service receives API calls from Envoy, it uses the value of the _service-node_ flag to compute routes that are relative to the service instances located at that IP address. |
| drain_duration | [google.protobuf.Duration](#google.protobuf.Duration) |  | The time in seconds that Envoy will drain connections during a hot restart. MUST be &gt;=1s (e.g., _1s/1m/1h_) |
| parent_shutdown_duration | [google.protobuf.Duration](#google.protobuf.Duration) |  | The time in seconds that Envoy will wait before shutting down the parent process during a hot restart. MUST be &gt;=1s (e.g., _1s/1m/1h_). MUST BE greater than _drain_duration_ parameter. |
| discovery_address | [string](#string) |  | Address of the discovery service exposing xDS with mTLS connection. |
| discovery_refresh_delay | [google.protobuf.Duration](#google.protobuf.Duration) |  | Polling interval for service discovery (used by EDS, CDS, LDS, but not RDS). (MUST BE &gt;=1ms) |
| zipkin_address | [string](#string) |  | Address of the Zipkin service (e.g. _zipkin:9411_). |
| connect_timeout | [google.protobuf.Duration](#google.protobuf.Duration) |  | Connection timeout used by Envoy for supporting services. (MUST BE &gt;=1ms) |
| statsd_udp_address | [string](#string) |  | IP Address and Port of a statsd UDP listener (e.g. _10.75.241.127:9125_). |
| proxy_admin_port | [int32](#int32) |  | Port on which Envoy should listen for administrative commands. |
| availability_zone | [string](#string) |  | The availability zone where this Envoy instance is running. When running Envoy as a sidecar in Kubernetes, this flag must be one of the availability zones assigned to a node using failure-domain.beta.kubernetes.io/zone annotation. |
| control_plane_auth_policy | [AuthenticationPolicy](#istio.mesh.v1alpha1.AuthenticationPolicy) |  | Authentication policy defines the global switch to control authentication for Envoy-to-Envoy communication for istio components Mixer and Pilot. |
| custom_config_file | [string](#string) |  | File path of custom proxy configuration, currently used by proxies in front of Mixer and Pilot. |
| stat_name_length | [int32](#int32) |  | Maximum length of name field in Envoy&#39;s metrics. The length of the name field is determined by the length of a name field in a service and the set of labels that comprise a particular version of the service. The default value is set to 189 characters. Envoy&#39;s internal metrics take up 67 characters, for a total of 256 character name per metric. Increase the value of this field if you find that the metrics from Envoys are truncated. |
| concurrency | [int32](#int32) |  | The number of worker threads to run. Default value is number of cores on the machine. |
| proxy_bootstrap_template_path | [string](#string) |  | Path to the proxy bootstrap template file |
| interception_mode | [ProxyConfig.InboundInterceptionMode](#istio.mesh.v1alpha1.ProxyConfig.InboundInterceptionMode) |  | The mode used to redirect inbound traffic to Envoy. |





 


<a name="istio.mesh.v1alpha1.AuthenticationPolicy"/>

### AuthenticationPolicy
AuthenticationPolicy defines authentication policy. It can be set for
different scopes (mesh, service …), and the most narrow scope with
non-INHERIT value will be used.
Mesh policy cannot be INHERIT.

| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 | Do not encrypt Envoy to Envoy traffic. |
| MUTUAL_TLS | 1 | Envoy to Envoy traffic is wrapped into mutual TLS connections. |
| INHERIT | 1000 | Use the policy defined by the parent scope. Should not be used for mesh policy. |



<a name="istio.mesh.v1alpha1.MeshConfig.AuthPolicy"/>

### MeshConfig.AuthPolicy
TODO AuthPolicy needs to be removed and merged with AuthPolicy defined above

| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 | Do not encrypt Envoy to Envoy traffic. |
| MUTUAL_TLS | 1 | Envoy to Envoy traffic is wrapped into mutual TLS connections. |



<a name="istio.mesh.v1alpha1.MeshConfig.IngressControllerMode"/>

### MeshConfig.IngressControllerMode


| Name | Number | Description |
| ---- | ------ | ----------- |
| OFF | 0 | Disables Istio ingress controller. |
| DEFAULT | 1 | Istio ingress controller will act on ingress resources that do not contain any annotation or whose annotations match the value specified in the ingress_class parameter described earlier. Use this mode if Istio ingress controller will be the default ingress controller for the entire kubernetes cluster. |
| STRICT | 2 | Istio ingress controller will only act on ingress resources whose annotations match the value specified in the ingress_class parameter described earlier. Use this mode if Istio ingress controller will be a secondary ingress controller (e.g., in addition to a cloud-provided ingress controller). |



<a name="istio.mesh.v1alpha1.ProxyConfig.InboundInterceptionMode"/>

### ProxyConfig.InboundInterceptionMode
The mode used to redirect inbound traffic to Envoy.
This setting has no effect on outbound traffic: iptables REDIRECT is always used for
outbound connections.

| Name | Number | Description |
| ---- | ------ | ----------- |
| REDIRECT | 0 | The REDIRECT mode uses iptables REDIRECT to NAT and redirect to Envoy. This mode loses source IP addresses during redirection. |
| TPROXY | 1 | The TPROXY mode uses iptables TPROXY to redirect to Envoy. This mode preserves both the source and destination IP addresses and ports, so that they can be used for advanced filtering and manipulation. This mode also configures the sidecar to run with the CAP_NET_ADMIN capability, which is required to use TPROXY. |


 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

