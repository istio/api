# Protocol Documentation
<a name="top"/>

## Table of Contents

- [policy/v1beta1/cfg.proto](#policy/v1beta1/cfg.proto)
    - [Action](#istio.policy.v1beta1.Action)
    - [AttributeManifest](#istio.policy.v1beta1.AttributeManifest)
    - [AttributeManifest.AttributeInfo](#istio.policy.v1beta1.AttributeManifest.AttributeInfo)
    - [AttributeManifest.AttributesEntry](#istio.policy.v1beta1.AttributeManifest.AttributesEntry)
    - [Connection](#istio.policy.v1beta1.Connection)
    - [Handler](#istio.policy.v1beta1.Handler)
    - [Instance](#istio.policy.v1beta1.Instance)
    - [Rule](#istio.policy.v1beta1.Rule)
  
  
  
  

- [policy/v1beta1/type.proto](#policy/v1beta1/type.proto)
    - [DNSName](#istio.policy.v1beta1.DNSName)
    - [Duration](#istio.policy.v1beta1.Duration)
    - [EmailAddress](#istio.policy.v1beta1.EmailAddress)
    - [IPAddress](#istio.policy.v1beta1.IPAddress)
    - [TimeStamp](#istio.policy.v1beta1.TimeStamp)
    - [Uri](#istio.policy.v1beta1.Uri)
    - [Value](#istio.policy.v1beta1.Value)
  
  
  
  

- [policy/v1beta1/value_type.proto](#policy/v1beta1/value_type.proto)
  
    - [ValueType](#istio.policy.v1beta1.ValueType)
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="policy/v1beta1/cfg.proto"/>
<p align="right"><a href="#top">Top</a></p>

## policy/v1beta1/cfg.proto



<a name="istio.policy.v1beta1.Action"/>

### Action
Action describes which [Handler][istio.policy.v1beta1.Handler] to invoke and what data to pass to it for processing.

The following example instructs Mixer to invoke &#39;prometheus-handler&#39; handler and pass it the object
constructed using the instance &#39;RequestCountByService&#39;.

```yaml
  handler: prometheus-handler
  instances:
  - RequestCountByService
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| handler | [string](#string) |  | Required. Fully qualified name of the handler to invoke. Must match the `name` of a [Handler][istio.policy.v1beta1.Handler.name]. |
| instances | [string](#string) | repeated | Required. Each value must match the fully qualified name of the [Instance][istio.policy.v1beta1.Instance.name]s. Referenced instances are evaluated by resolving the attributes/literals for all the fields. The constructed objects are then passed to the `handler` referenced within this action. |






<a name="istio.policy.v1beta1.AttributeManifest"/>

### AttributeManifest
AttributeManifest describes a set of Attributes produced by some component
of an Istio deployment.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| revision | [string](#string) |  | Optional. The revision of this document. Assigned by server. |
| name | [string](#string) |  | Required. Name of the component producing these attributes. This can be the proxy (with the canonical name &#34;istio-proxy&#34;) or the name of an `attributes` kind adapter in Mixer. |
| attributes | [AttributeManifest.AttributesEntry](#istio.policy.v1beta1.AttributeManifest.AttributesEntry) | repeated | The set of attributes this Istio component will be responsible for producing at runtime. We map from attribute name to the attribute&#39;s specification. The name of an attribute, which is how attributes are referred to in aspect configuration, must conform to:   Name = IDENT { SEPARATOR IDENT };  Where `IDENT` must match the regular expression `[a-z][a-z0-9]&#43;` and `SEPARATOR` must match the regular expression `[\.-]`.  Attribute names must be unique within a single Istio deployment. The set of canonical attributes are described at https://istio.io/docs/reference/attribute-vocabulary.html. Attributes not in that list should be named with a component-specific suffix such as request.count-my.component. |






<a name="istio.policy.v1beta1.AttributeManifest.AttributeInfo"/>

### AttributeManifest.AttributeInfo
AttributeInfo describes the schema of an Istio `Attribute`.

# Istio Attributes

Istio uses `attributes` to describe runtime activities of Istio services.
An Istio attribute carries a specific piece of information about an activity,
such as the error code of an API request, the latency of an API request, or the
original IP address of a TCP connection. The attributes are often generated
and consumed by different services. For example, a frontend service can
generate an authenticated user attribute and pass it to a backend service for
access control purpose.

To simplify the system and improve developer experience, Istio uses
shared attribute definitions across all components. For example, the same
authenticated user attribute will be used for logging, monitoring, analytics,
billing, access control, auditing. Many Istio components provide their
functionality by collecting, generating, and operating on attributes.
For example, the proxy collects the error code attribute, and the logging
stores it into a log.

# Design

Each Istio attribute must conform to an `AttributeInfo` in an
`AttributeManifest` in the current Istio deployment at runtime. An
[`AttributeInfo`][istio.policy.v1beta1] is used to define an attribute&#39;s
metadata: the type of its value and a detailed description that explains
the semantics of the attribute type. Each attribute&#39;s name is globally unique;
in other words an attribute name can only appear once across all manifests.

The runtime presentation of an attribute is intentionally left out of this
specification, because passing attribute using JSON, XML, or Protocol Buffers
does not change the semantics of the attribute. Different implementations
can choose different representations based on their needs.

# HTTP Mapping

Because many systems already have REST APIs, it makes sense to define a
standard HTTP mapping for Istio attributes that are compatible with typical
REST APIs. The design is to map one attribute to one HTTP header, the
attribute name and value becomes the HTTP header name and value. The actual
encoding scheme will be decided later.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| description | [string](#string) |  | Optional. A human-readable description of the attribute&#39;s purpose. |
| value_type | [ValueType](#istio.policy.v1beta1.ValueType) |  | Required. The type of data carried by this attribute. |






<a name="istio.policy.v1beta1.AttributeManifest.AttributesEntry"/>

### AttributeManifest.AttributesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [AttributeManifest.AttributeInfo](#istio.policy.v1beta1.AttributeManifest.AttributeInfo) |  |  |






<a name="istio.policy.v1beta1.Connection"/>

### Connection
Connection allows the operator to specify the endpoint for out-of-process infrastructure backend.
Connection is part of the handler custom resource and is specified alongside adapter specific configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | The address of the backend. |






<a name="istio.policy.v1beta1.Handler"/>

### Handler
Handler allows the operator to configure a specific adapter implementation.
Each adapter implementation defines its own `params` proto.

In the following example we define a `metrics` handler for the `prometheus` adapter.
The example is in the form of a kubernetes resource:
* The `metadata.name` is the name of the handler
* The `kind` refers to the adapter name
* The `spec` block represents adapter-specific configuration as well as the connection information

```yaml
# Sample-1: No connection specified (for compiled in adapters)
# Note: if connection information is not specified, the adapter configuration is directly inside
# `spec` block. This is going to be DEPRECATED in favor of Sample-2
apiVersion: &#34;config.istio.io/v1alpha2&#34;
kind: prometheus
metadata:
  name: handler
  namespace: istio-system
spec:
  metrics:
  - name: request_count
    instance_name: requestcount.metric.istio-system
    kind: COUNTER
    label_names:
    - source_service
    - source_version
    - destination_service
    - destination_version
---
# Sample-2: With connection information (for out-of-process adapters)
# Note: Unlike sample-1, the adapter configuration is parallel to `connection` and is nested inside `param` block.
apiVersion: &#34;config.istio.io/v1alpha2&#34;
kind: prometheus
metadata:
  name: handler
  namespace: istio-system
spec:
  param:
    metrics:
    - name: request_count
      instance_name: requestcount.metric.istio-system
      kind: COUNTER
      label_names:
      - source_service
      - source_version
      - destination_service
      - destination_version
  connection:
    address: localhost:8090
---
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Required. Must be unique in the entire mixer configuration. Used by [Actions][istio.policy.v1beta1.Action.handler] to refer to this handler. |
| compiled_adapter | [string](#string) |  | Required. The name of the compiled in adapter this handler instantiates. For referencing non compiled-in adapters, use the `adapter` field instead.  The value must match the name of the available adapter Mixer is built with. An adapter&#39;s name is typically a constant in its code. |
| adapter | [string](#string) |  | Required. The name of a specific adapter implementation. For referencing compiled-in adapters, use the `compiled_adapter` field instead.  An adapter&#39;s implementation name is typically a constant in its code. |
| params | [google.protobuf.Struct](#google.protobuf.Struct) |  | Optional. Depends on adapter implementation. Struct representation of a proto defined by the adapter implementation; this varies depending on the value of field `adapter`. |
| connection | [Connection](#istio.policy.v1beta1.Connection) |  | Optional. Information on how to connect to the out-of-process adapter. This is used if the adapter is not compiled into Mixer binary and is running as a separate process. |






<a name="istio.policy.v1beta1.Instance"/>

### Instance
An Instance tells Mixer how to create instances for particular template.

Instance is defined by the operator. Instance is defined relative to a known
template. Their purpose is to tell Mixer how to use attributes or literals to produce
instances of the specified template at runtime.

The following example instructs Mixer to construct an instance associated with template
&#39;istio.mixer.adapter.metric.Metric&#39;. It provides a mapping from the template&#39;s fields to expressions.
Instances produced with this instance can be referenced by [Actions][istio.policy.v1beta1.Action] using name
&#39;RequestCountByService&#39;

```yaml
- name: RequestCountByService
  template: istio.mixer.adapter.metric.Metric
  params:
    value: 1
    dimensions:
      source: source.service
      destination_ip: destination.ip
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Required. The name of this instance  Must be unique amongst other Instances in scope. Used by [Action][istio.policy.v1beta1.Action] to refer to an instance produced by this instance. |
| compiled_template | [string](#string) |  | Required. The name of the compiled in template this instance creates instances for. For referencing non compiled-in templates, use the `template` field instead.  The value must match the name of the available template Mixer is built with. |
| template | [string](#string) |  | Required. The name of the template this instance creates instances for. For referencing compiled-in templates, use the `compiled_template` field instead.  The value must match the name of the available template in scope. |
| params | [google.protobuf.Struct](#google.protobuf.Struct) |  | Required. Depends on referenced template. Struct representation of a proto defined by the template; this varies depending on the value of field `template`. |






<a name="istio.policy.v1beta1.Rule"/>

### Rule
A Rule is a selector and a set of intentions to be executed when the
selector is `true`

The following example instructs Mixer to invoke &#39;prometheus-handler&#39; handler for all services and pass it the
instance constructed using the &#39;RequestCountByService&#39; instance.

```yaml
- match: destination.service == &#34;*&#34;
  actions:
  - handler: prometheus-handler
    instances:
    - RequestCountByService
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| match | [string](#string) |  | Required. Match is an attribute based predicate. When Mixer receives a request it evaluates the match expression and executes all the associated `actions` if the match evaluates to true.  A few example match:  * an empty match evaluates to `true` * `true`, a boolean literal; a rule with this match will always be executed * `destination.service == ratings*` selects any request targeting a service whose name starts with &#34;ratings&#34; * `attr1 == &#34;20&#34; &amp;&amp; attr2 == &#34;30&#34;` logical AND, OR, and NOT are also available |
| actions | [Action](#istio.policy.v1beta1.Action) | repeated | Optional. The actions that will be executed when match evaluates to `true`. |





 

 

 

 



<a name="policy/v1beta1/type.proto"/>
<p align="right"><a href="#top">Top</a></p>

## policy/v1beta1/type.proto



<a name="istio.policy.v1beta1.DNSName"/>

### DNSName
An instance field of type DNSName denotes that the expression for the field must evalaute to
[ValueType.DNS_NAME][istio.policy.v1beta1.ValueType.DNS_NAME]

Objects of type DNSName are also passed to the adapters during request-time for the instance fields of
type DNSName


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  | DNSName encoded as string. |






<a name="istio.policy.v1beta1.Duration"/>

### Duration
An instance field of type Duration denotes that the expression for the field must evalaute to
[ValueType.DURATION][istio.policy.v1beta1.ValueType.DURATION]

Objects of type Duration are also passed to the adapters during request-time for the instance fields of
type Duration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [google.protobuf.Duration](#google.protobuf.Duration) |  | Duration encoded as google.protobuf.Duration. |






<a name="istio.policy.v1beta1.EmailAddress"/>

### EmailAddress
DO NOT USE !! Under Development
An instance field of type EmailAddress denotes that the expression for the field must evalaute to
[ValueType.EMAIL_ADDRESS][istio.policy.v1beta1.ValueType.EMAIL_ADDRESS]

Objects of type EmailAddress are also passed to the adapters during request-time for the instance fields of
type EmailAddress


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  | EmailAddress encoded as string. |






<a name="istio.policy.v1beta1.IPAddress"/>

### IPAddress
An instance field of type IPAddress denotes that the expression for the field must evalaute to
[ValueType.IP_ADDRESS][istio.policy.v1beta1.ValueType.IP_ADDRESS]

Objects of type IPAddress are also passed to the adapters during request-time for the instance fields of
type IPAddress


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [bytes](#bytes) |  | IPAddress encoded as bytes. |






<a name="istio.policy.v1beta1.TimeStamp"/>

### TimeStamp
An instance field of type TimeStamp denotes that the expression for the field must evalaute to
[ValueType.TIMESTAMP][istio.policy.v1beta1.ValueType.TIMESTAMP]

Objects of type TimeStamp are also passed to the adapters during request-time for the instance fields of
type TimeStamp


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | TimeStamp encoded as google.protobuf.Timestamp. |






<a name="istio.policy.v1beta1.Uri"/>

### Uri
DO NOT USE !! Under Development
An instance field of type Uri denotes that the expression for the field must evalaute to
[ValueType.URI][istio.policy.v1beta1.ValueType.URI]

Objects of type Uri are also passed to the adapters during request-time for the instance fields of
type Uri


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  | Uri encoded as string. |






<a name="istio.policy.v1beta1.Value"/>

### Value
An instance field of type Value denotes that the expression for the field is of dynamic type and can evalaute to any
[ValueType][istio.policy.v1beta1.ValueType] enum values. For example, when
authoring an instance configuration for a template that has a field `data` of type `istio.policy.v1beta1.Value`,
both of the following expressions are valid `data: source.ip | ip(&#34;0.0.0.0&#34;)`, `data: request.id | &#34;&#34;`;
the resulting type is either ValueType.IP_ADDRESS or ValueType.STRING for the two cases respectively.

Objects of type Value are also passed to the adapters during request-time. There is a 1:1 mapping between
oneof fields in `Value` and enum values inside `ValueType`. Depending on the expression&#39;s evaluated `ValueType`,
the equivalent oneof field in `Value` is populated by Mixer and passed to the adapters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| string_value | [string](#string) |  | Used for values of type STRING |
| int64_value | [int64](#int64) |  | Used for values of type INT64 |
| double_value | [double](#double) |  | Used for values of type DOUBLE |
| bool_value | [bool](#bool) |  | Used for values of type BOOL |
| ip_address_value | [IPAddress](#istio.policy.v1beta1.IPAddress) |  | Used for values of type IPAddress |
| timestamp_value | [TimeStamp](#istio.policy.v1beta1.TimeStamp) |  | Used for values of type TIMESTAMP |
| duration_value | [Duration](#istio.policy.v1beta1.Duration) |  | Used for values of type DURATION |
| email_address_value | [EmailAddress](#istio.policy.v1beta1.EmailAddress) |  | Used for values of type EmailAddress |
| dns_name_value | [DNSName](#istio.policy.v1beta1.DNSName) |  | Used for values of type DNSName |
| uri_value | [Uri](#istio.policy.v1beta1.Uri) |  | Used for values of type Uri |





 

 

 

 



<a name="policy/v1beta1/value_type.proto"/>
<p align="right"><a href="#top">Top</a></p>

## policy/v1beta1/value_type.proto


 


<a name="istio.policy.v1beta1.ValueType"/>

### ValueType
ValueType describes the types that values in the Istio system can take. These
are used to describe the type of Attributes at run time, describe the type of
the result of evaluating an expression, and to describe the runtime type of
fields of other descriptors.

| Name | Number | Description |
| ---- | ------ | ----------- |
| VALUE_TYPE_UNSPECIFIED | 0 | Invalid, default value. |
| STRING | 1 | An undiscriminated variable-length string. |
| INT64 | 2 | An undiscriminated 64-bit signed integer. |
| DOUBLE | 3 | An undiscriminated 64-bit floating-point value. |
| BOOL | 4 | An undiscriminated boolean value. |
| TIMESTAMP | 5 | A point in time. |
| IP_ADDRESS | 6 | An IP address. |
| EMAIL_ADDRESS | 7 | An email address. |
| URI | 8 | A URI. |
| DNS_NAME | 9 | A DNS name. |
| DURATION | 10 | A span between two points in time. |
| STRING_MAP | 11 | A map string -&gt; string, typically used by headers. |


 

 

 



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

