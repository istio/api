# Protocol Documentation
<a name="top"/>

## Table of Contents

- [mixer/v1/config/client/api_spec.proto](#mixer/v1/config/client/api_spec.proto)
    - [APIKey](#istio.mixer.v1.config.client.APIKey)
    - [HTTPAPISpec](#istio.mixer.v1.config.client.HTTPAPISpec)
    - [HTTPAPISpecBinding](#istio.mixer.v1.config.client.HTTPAPISpecBinding)
    - [HTTPAPISpecPattern](#istio.mixer.v1.config.client.HTTPAPISpecPattern)
    - [HTTPAPISpecReference](#istio.mixer.v1.config.client.HTTPAPISpecReference)
  
  
  
  

- [mixer/v1/config/client/client_config.proto](#mixer/v1/config/client/client_config.proto)
    - [HttpClientConfig](#istio.mixer.v1.config.client.HttpClientConfig)
    - [HttpClientConfig.ServiceConfigsEntry](#istio.mixer.v1.config.client.HttpClientConfig.ServiceConfigsEntry)
    - [NetworkFailPolicy](#istio.mixer.v1.config.client.NetworkFailPolicy)
    - [ServiceConfig](#istio.mixer.v1.config.client.ServiceConfig)
    - [TcpClientConfig](#istio.mixer.v1.config.client.TcpClientConfig)
    - [TransportConfig](#istio.mixer.v1.config.client.TransportConfig)
  
    - [NetworkFailPolicy.FailPolicy](#istio.mixer.v1.config.client.NetworkFailPolicy.FailPolicy)
  
  
  

- [mixer/v1/config/client/quota.proto](#mixer/v1/config/client/quota.proto)
    - [AttributeMatch](#istio.mixer.v1.config.client.AttributeMatch)
    - [AttributeMatch.ClauseEntry](#istio.mixer.v1.config.client.AttributeMatch.ClauseEntry)
    - [Quota](#istio.mixer.v1.config.client.Quota)
    - [QuotaRule](#istio.mixer.v1.config.client.QuotaRule)
    - [QuotaSpec](#istio.mixer.v1.config.client.QuotaSpec)
    - [QuotaSpecBinding](#istio.mixer.v1.config.client.QuotaSpecBinding)
    - [QuotaSpecBinding.QuotaSpecReference](#istio.mixer.v1.config.client.QuotaSpecBinding.QuotaSpecReference)
    - [StringMatch](#istio.mixer.v1.config.client.StringMatch)
  
  
  
  

- [mixer/v1/config/client/service.proto](#mixer/v1/config/client/service.proto)
    - [IstioService](#istio.mixer.v1.config.client.IstioService)
    - [IstioService.LabelsEntry](#istio.mixer.v1.config.client.IstioService.LabelsEntry)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="mixer/v1/config/client/api_spec.proto"/>
<p align="right"><a href="#top">Top</a></p>

## mixer/v1/config/client/api_spec.proto



<a name="istio.mixer.v1.config.client.APIKey"/>

### APIKey
APIKey defines the explicit configuration for generating the
`request.api_key` attribute from HTTP requests.

See https://swagger.io/docs/specification/authentication/api-keys
for a general overview of API keys as defined by OpenAPI.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [string](#string) |  | API Key is sent as a query parameter. `query` represents the query string parameter name.  For example, `query=api_key` should be used with the following request:   GET /something?api_key=abcdef12345 |
| header | [string](#string) |  | API key is sent in a request header. `header` represents the header name.  For example, `header=X-API-KEY` should be used with the following request:   GET /something HTTP/1.1 X-API-Key: abcdef12345 |
| cookie | [string](#string) |  | API key is sent in a [cookie](https://swagger.io/docs/specification/authentication/cookie-authentication),  For example, `cookie=X-API-KEY` should be used for the following request:   GET /something HTTP/1.1 Cookie: X-API-KEY=abcdef12345 |






<a name="istio.mixer.v1.config.client.HTTPAPISpec"/>

### HTTPAPISpec
HTTPAPISpec defines the canonical configuration for generating
API-related attributes from HTTP requests based on the method and
uri templated path matches. It is sufficient for defining the API
surface of a service for the purposes of API attribute
generation. It is not intended to represent auth, quota,
documentation, or other information commonly found in other API
specifications, e.g. OpenAPI.

Existing standards that define operations (or methods) in terms of
HTTP methods and paths can be normalized to this format for use in
Istio. For example, a simple petstore API described by OpenAPIv2
[here](https://github.com/googleapis/gnostic/blob/master/examples/v2.0/yaml/petstore-simple.yaml)
can be represented with the following HTTPAPISpec.

```yaml
apiVersion: config.istio.io/v1alpha2
kind: HTTPAPISpec
metadata:
  name: petstore
  namespace: default
spec:
  attributes:
    api.service: petstore.swagger.io
    api.version: 1.0.0
  patterns:
  - attributes:
      api.operation: findPets
    httpMethod: GET
    uriTemplate: /api/pets
  - attributes:
      api.operation: addPet
    httpMethod: POST
    uriTemplate: /api/pets
  - attributes:
      api.operation: findPetById
    httpMethod: GET
    uriTemplate: /api/pets/{id}
  - attributes:
      api.operation: deletePet
    httpMethod: DELETE
    uriTemplate: /api/pets/{id}
  api_keys:
  - query: api-key
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| attributes | [istio.mixer.v1.Attributes](#istio.mixer.v1.Attributes) |  | List of attributes that are generated when *any* of the HTTP patterns match. This list typically includes the &#34;api.service&#34; and &#34;api.version&#34; attributes. |
| patterns | [HTTPAPISpecPattern](#istio.mixer.v1.config.client.HTTPAPISpecPattern) | repeated | List of HTTP patterns to match. |
| api_keys | [APIKey](#istio.mixer.v1.config.client.APIKey) | repeated | List of APIKey that describes how to extract an API-KEY from an HTTP request. The first API-Key match found in the list is used, i.e. &#39;OR&#39; semantics.  The following default policies are used to generate the `request.api_key` attribute if no explicit APIKey is defined.   `query: key, `query: api_key`, and then `header: x-api-key` |






<a name="istio.mixer.v1.config.client.HTTPAPISpecBinding"/>

### HTTPAPISpecBinding
HTTPAPISpecBinding defines the binding between HTTPAPISpecs and one or more
IstioService. For example, the following establishes a binding
between the HTTPAPISpec `petstore` and service `foo` in namespace `bar`.

```yaml
apiVersion: config.istio.io/v1alpha2
kind: HTTPAPISpecBinding
metadata:
  name: my-binding
  namespace: default
spec:
  services:
  - name: foo
    namespace: bar
  api_specs:
  - name: petstore
    namespace: default
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| services | [IstioService](#istio.mixer.v1.config.client.IstioService) | repeated | REQUIRED. One or more services to map the listed HTTPAPISpec onto. |
| api_specs | [HTTPAPISpecReference](#istio.mixer.v1.config.client.HTTPAPISpecReference) | repeated | REQUIRED. One or more HTTPAPISpec references that should be mapped to the specified service(s). The aggregate collection of match conditions defined in the HTTPAPISpecs should not overlap. |






<a name="istio.mixer.v1.config.client.HTTPAPISpecPattern"/>

### HTTPAPISpecPattern
HTTPAPISpecPattern defines a single pattern to match against
incoming HTTP requests. The per-pattern list of attributes is
generated if both the http_method and uri_template match. In
addition, the top-level list of attributes in the HTTPAPISpec is also
generated.

```yaml
pattern:
- attributes
    api.operation: doFooBar
  httpMethod: GET
  uriTemplate: /foo/bar
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| attributes | [istio.mixer.v1.Attributes](#istio.mixer.v1.Attributes) |  | List of attributes that are generated if the HTTP request matches the specified http_method and uri_template. This typically includes the &#34;api.operation&#34; attribute. |
| http_method | [string](#string) |  | HTTP request method to match against as defined by [rfc7231](https://tools.ietf.org/html/rfc7231#page-21). For example: GET, HEAD, POST, PUT, DELETE. |
| uri_template | [string](#string) |  | URI template to match against as defined by [rfc6570](https://tools.ietf.org/html/rfc6570). For example, the following are valid URI templates:   /pets /pets/{id} /dictionary/{term:1}/{term} /search{?q*,lang} |
| regex | [string](#string) |  | EXPERIMENTAL:  ecmascript style regex-based match as defined by [EDCA-262](http://en.cppreference.com/w/cpp/regex/ecmascript). For example,   &#34;^/pets/(.*?)?&#34; |






<a name="istio.mixer.v1.config.client.HTTPAPISpecReference"/>

### HTTPAPISpecReference
HTTPAPISpecReference defines a reference to an HTTPAPISpec. This is
typically used for establishing bindings between an HTTPAPISpec and an
IstioService. For example, the following defines an
HTTPAPISpecReference for service `foo` in namespace `bar`.

```yaml
- name: foo
  namespace: bar
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | REQUIRED. The short name of the HTTPAPISpec. This is the resource name defined by the metadata name field. |
| namespace | [string](#string) |  | Optional namespace of the HTTPAPISpec. Defaults to the encompassing HTTPAPISpecBinding&#39;s metadata namespace field. |





 

 

 

 



<a name="mixer/v1/config/client/client_config.proto"/>
<p align="right"><a href="#top">Top</a></p>

## mixer/v1/config/client/client_config.proto



<a name="istio.mixer.v1.config.client.HttpClientConfig"/>

### HttpClientConfig
Defines the client config for HTTP.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transport | [TransportConfig](#istio.mixer.v1.config.client.TransportConfig) |  | The transport config. |
| service_configs | [HttpClientConfig.ServiceConfigsEntry](#istio.mixer.v1.config.client.HttpClientConfig.ServiceConfigsEntry) | repeated | Map of control configuration indexed by destination.service. This is used to support per-service configuration for cases where a mixerclient serves multiple services. |
| default_destination_service | [string](#string) |  | Default destination service name if none was specified in the client request. |
| mixer_attributes | [istio.mixer.v1.Attributes](#istio.mixer.v1.Attributes) |  | Default attributes to send to Mixer in both Check and Report. This typically includes &#34;destination.ip&#34; and &#34;destination.uid&#34; attributes. |
| forward_attributes | [istio.mixer.v1.Attributes](#istio.mixer.v1.Attributes) |  | Default attributes to forward to upstream. This typically includes the &#34;source.ip&#34; and &#34;source.uid&#34; attributes. |






<a name="istio.mixer.v1.config.client.HttpClientConfig.ServiceConfigsEntry"/>

### HttpClientConfig.ServiceConfigsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [ServiceConfig](#istio.mixer.v1.config.client.ServiceConfig) |  |  |






<a name="istio.mixer.v1.config.client.NetworkFailPolicy"/>

### NetworkFailPolicy
Specifies the behavior when the client is unable to connect to Mixer.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| policy | [NetworkFailPolicy.FailPolicy](#istio.mixer.v1.config.client.NetworkFailPolicy.FailPolicy) |  | Specifies the behavior when the client is unable to connect to Mixer. |






<a name="istio.mixer.v1.config.client.ServiceConfig"/>

### ServiceConfig
Defines the per-service client configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| disable_check_calls | [bool](#bool) |  | If true, do not call Mixer Check. |
| disable_report_calls | [bool](#bool) |  | If true, do not call Mixer Report. |
| mixer_attributes | [istio.mixer.v1.Attributes](#istio.mixer.v1.Attributes) |  | Send these attributes to Mixer in both Check and Report. This typically includes the &#34;destination.service&#34; attribute. In case of a per-route override, per-route attributes take precedence over the attributes supplied in the client configuration. |
| http_api_spec | [HTTPAPISpec](#istio.mixer.v1.config.client.HTTPAPISpec) | repeated | HTTP API specifications to generate API attributes. |
| quota_spec | [QuotaSpec](#istio.mixer.v1.config.client.QuotaSpec) | repeated | Quota specifications to generate quota requirements. |
| network_fail_policy | [NetworkFailPolicy](#istio.mixer.v1.config.client.NetworkFailPolicy) |  | Specifies the behavior when the client is unable to connect to Mixer. This is the service-level policy. It overrides [mesh-level policy][istio.mixer.v1.config.client.TransportConfig.network_fail_policy]. |
| forward_attributes | [istio.mixer.v1.Attributes](#istio.mixer.v1.Attributes) |  | Default attributes to forward to upstream. This typically includes the &#34;source.ip&#34; and &#34;source.uid&#34; attributes. In case of a per-route override, per-route attributes take precedence over the attributes supplied in the client configuration.  Forwarded attributes take precedence over the static Mixer attributes. The full order of application is as follows: 1. static Mixer attributes from the filter config; 2. static Mixer attributes from the route config; 3. forwarded attributes from the source filter config (if any); 4. forwarded attributes from the source route config (if any); 5. derived attributes from the request metadata. |






<a name="istio.mixer.v1.config.client.TcpClientConfig"/>

### TcpClientConfig
Defines the client config for TCP.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transport | [TransportConfig](#istio.mixer.v1.config.client.TransportConfig) |  | The transport config. |
| mixer_attributes | [istio.mixer.v1.Attributes](#istio.mixer.v1.Attributes) |  | Default attributes to send to Mixer in both Check and Report. This typically includes &#34;destination.ip&#34; and &#34;destination.uid&#34; attributes. |
| disable_check_calls | [bool](#bool) |  | If set to true, disables mixer check calls. |
| disable_report_calls | [bool](#bool) |  | If set to true, disables mixer check calls. |
| connection_quota_spec | [QuotaSpec](#istio.mixer.v1.config.client.QuotaSpec) |  | Quota specifications to generate quota requirements. It applies on the new TCP connections. |
| report_interval | [google.protobuf.Duration](#google.protobuf.Duration) |  | Specify report interval to send periodical reports for long TCP connections. If not specified, the interval is 10 seconds. This interval should not be less than 1 second, otherwise it will be reset to 1 second. |






<a name="istio.mixer.v1.config.client.TransportConfig"/>

### TransportConfig
Defines the transport config on how to call Mixer.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| disable_check_cache | [bool](#bool) |  | The flag to disable check cache. |
| disable_quota_cache | [bool](#bool) |  | The flag to disable quota cache. |
| disable_report_batch | [bool](#bool) |  | The flag to disable report batch. |
| network_fail_policy | [NetworkFailPolicy](#istio.mixer.v1.config.client.NetworkFailPolicy) |  | Specifies the behavior when the client is unable to connect to Mixer. This is the mesh level policy. The default value for policy is FAIL_OPEN. |
| stats_update_interval | [google.protobuf.Duration](#google.protobuf.Duration) |  | Specify refresh interval to write mixer client statistics to Envoy share memory. If not specified, the interval is 10 seconds. |
| check_cluster | [string](#string) |  | Name of the cluster that will forward check calls to a pool of mixer servers. Defaults to &#34;mixer_server&#34;. By using different names for checkCluster and reportCluster, it is possible to have one set of mixer servers handle check calls, while another set of mixer servers handle report calls.  NOTE: Any value other than the default &#34;mixer_server&#34; will require the Istio Grafana dashboards to be reconfigured to use the new name. |
| report_cluster | [string](#string) |  | Name of the cluster that will forward report calls to a pool of mixer servers. Defaults to &#34;mixer_server&#34;. By using different names for checkCluster and reportCluster, it is possible to have one set of mixer servers handle check calls, while another set of mixer servers handle report calls.  NOTE: Any value other than the default &#34;mixer_server&#34; will require the Istio Grafana dashboards to be reconfigured to use the new name. |
| attributes_for_mixer_proxy | [istio.mixer.v1.Attributes](#istio.mixer.v1.Attributes) |  | Default attributes to forward to mixer upstream. This typically includes the &#34;source.ip&#34; and &#34;source.uid&#34; attributes. These attributes are consumed by the proxy in front of mixer. |





 


<a name="istio.mixer.v1.config.client.NetworkFailPolicy.FailPolicy"/>

### NetworkFailPolicy.FailPolicy
Describes the policy.

| Name | Number | Description |
| ---- | ------ | ----------- |
| FAIL_OPEN | 0 | If network connection fails, request is allowed and delivered to the service. |
| FAIL_CLOSE | 1 | If network connection fails, request is rejected. |


 

 

 



<a name="mixer/v1/config/client/quota.proto"/>
<p align="right"><a href="#top">Top</a></p>

## mixer/v1/config/client/quota.proto



<a name="istio.mixer.v1.config.client.AttributeMatch"/>

### AttributeMatch
Specifies a match clause to match Istio attributes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clause | [AttributeMatch.ClauseEntry](#istio.mixer.v1.config.client.AttributeMatch.ClauseEntry) | repeated | Map of attribute names to StringMatch type. Each map element specifies one condition to match.  Example:   clause: source.uid: exact: SOURCE_UID request.http_method: exact: POST |






<a name="istio.mixer.v1.config.client.AttributeMatch.ClauseEntry"/>

### AttributeMatch.ClauseEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [StringMatch](#istio.mixer.v1.config.client.StringMatch) |  |  |






<a name="istio.mixer.v1.config.client.Quota"/>

### Quota
Specifies a quota to use with quota name and amount.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| quota | [string](#string) |  | The quota name to charge |
| charge | [int64](#int64) |  | The quota amount to charge |






<a name="istio.mixer.v1.config.client.QuotaRule"/>

### QuotaRule
Specifies a rule with list of matches and list of quotas.
If any clause matched, the list of quotas will be used.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| match | [AttributeMatch](#istio.mixer.v1.config.client.AttributeMatch) | repeated | If empty, match all request. If any of match is true, it is matched. |
| quotas | [Quota](#istio.mixer.v1.config.client.Quota) | repeated | The list of quotas to charge. |






<a name="istio.mixer.v1.config.client.QuotaSpec"/>

### QuotaSpec
Determines the quotas used for individual requests.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rules | [QuotaRule](#istio.mixer.v1.config.client.QuotaRule) | repeated | A list of Quota rules. |






<a name="istio.mixer.v1.config.client.QuotaSpecBinding"/>

### QuotaSpecBinding
QuotaSpecBinding defines the binding between QuotaSpecs and one or more
IstioService.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| services | [IstioService](#istio.mixer.v1.config.client.IstioService) | repeated | REQUIRED. One or more services to map the listed QuotaSpec onto. |
| quota_specs | [QuotaSpecBinding.QuotaSpecReference](#istio.mixer.v1.config.client.QuotaSpecBinding.QuotaSpecReference) | repeated | REQUIRED. One or more QuotaSpec references that should be mapped to the specified service(s). The aggregate collection of match conditions defined in the QuotaSpecs should not overlap. |






<a name="istio.mixer.v1.config.client.QuotaSpecBinding.QuotaSpecReference"/>

### QuotaSpecBinding.QuotaSpecReference
QuotaSpecReference uniquely identifies the QuotaSpec used in the
Binding.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | REQUIRED. The short name of the QuotaSpec. This is the resource name defined by the metadata name field. |
| namespace | [string](#string) |  | Optional namespace of the QuotaSpec. Defaults to the value of the metadata namespace field. |






<a name="istio.mixer.v1.config.client.StringMatch"/>

### StringMatch
Describes how to match a given string in HTTP headers. Match is
case-sensitive.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exact | [string](#string) |  | exact string match |
| prefix | [string](#string) |  | prefix-based match |
| regex | [string](#string) |  | ECMAscript style regex-based match |





 

 

 

 



<a name="mixer/v1/config/client/service.proto"/>
<p align="right"><a href="#top">Top</a></p>

## mixer/v1/config/client/service.proto



<a name="istio.mixer.v1.config.client.IstioService"/>

### IstioService
IstioService identifies a service and optionally service version.
The FQDN of the service is composed from the name, namespace, and implementation-specific domain suffix
(e.g. on Kubernetes, &#34;reviews&#34; &#43; &#34;default&#34; &#43; &#34;svc.cluster.local&#34; -&gt; &#34;reviews.default.svc.cluster.local&#34;).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The short name of the service such as &#34;foo&#34;. |
| namespace | [string](#string) |  | Optional namespace of the service. Defaults to value of metadata namespace field. |
| domain | [string](#string) |  | Domain suffix used to construct the service FQDN in implementations that support such specification. |
| service | [string](#string) |  | The service FQDN. |
| labels | [IstioService.LabelsEntry](#istio.mixer.v1.config.client.IstioService.LabelsEntry) | repeated | Optional one or more labels that uniquely identify the service version.  *Note:* When used for a RouteRule destination, labels MUST be empty. |






<a name="istio.mixer.v1.config.client.IstioService.LabelsEntry"/>

### IstioService.LabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 

 

 

 



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

