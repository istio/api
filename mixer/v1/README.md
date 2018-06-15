# Protocol Documentation
<a name="top"/>

## Table of Contents

- [mixer/v1/attributes.proto](#mixer/v1/attributes.proto)
    - [Attributes](#istio.mixer.v1.Attributes)
    - [Attributes.AttributeValue](#istio.mixer.v1.Attributes.AttributeValue)
    - [Attributes.AttributesEntry](#istio.mixer.v1.Attributes.AttributesEntry)
    - [Attributes.StringMap](#istio.mixer.v1.Attributes.StringMap)
    - [CompressedAttributes](#istio.mixer.v1.CompressedAttributes)
    - [CompressedAttributes.BoolsEntry](#istio.mixer.v1.CompressedAttributes.BoolsEntry)
    - [CompressedAttributes.BytesEntry](#istio.mixer.v1.CompressedAttributes.BytesEntry)
    - [CompressedAttributes.DoublesEntry](#istio.mixer.v1.CompressedAttributes.DoublesEntry)
    - [CompressedAttributes.DurationsEntry](#istio.mixer.v1.CompressedAttributes.DurationsEntry)
    - [CompressedAttributes.Int64sEntry](#istio.mixer.v1.CompressedAttributes.Int64sEntry)
    - [CompressedAttributes.StringMapsEntry](#istio.mixer.v1.CompressedAttributes.StringMapsEntry)
    - [CompressedAttributes.StringsEntry](#istio.mixer.v1.CompressedAttributes.StringsEntry)
    - [CompressedAttributes.TimestampsEntry](#istio.mixer.v1.CompressedAttributes.TimestampsEntry)
    - [StringMap](#istio.mixer.v1.StringMap)
    - [StringMap.EntriesEntry](#istio.mixer.v1.StringMap.EntriesEntry)
  
  
  
  

- [mixer/v1/check.proto](#mixer/v1/check.proto)
    - [CheckRequest](#istio.mixer.v1.CheckRequest)
    - [CheckRequest.QuotaParams](#istio.mixer.v1.CheckRequest.QuotaParams)
    - [CheckRequest.QuotasEntry](#istio.mixer.v1.CheckRequest.QuotasEntry)
    - [CheckResponse](#istio.mixer.v1.CheckResponse)
    - [CheckResponse.PreconditionResult](#istio.mixer.v1.CheckResponse.PreconditionResult)
    - [CheckResponse.QuotaResult](#istio.mixer.v1.CheckResponse.QuotaResult)
    - [CheckResponse.QuotasEntry](#istio.mixer.v1.CheckResponse.QuotasEntry)
    - [HeaderOperation](#istio.mixer.v1.HeaderOperation)
    - [ReferencedAttributes](#istio.mixer.v1.ReferencedAttributes)
    - [ReferencedAttributes.AttributeMatch](#istio.mixer.v1.ReferencedAttributes.AttributeMatch)
    - [RouteDirective](#istio.mixer.v1.RouteDirective)
  
    - [HeaderOperation.Operation](#istio.mixer.v1.HeaderOperation.Operation)
    - [ReferencedAttributes.Condition](#istio.mixer.v1.ReferencedAttributes.Condition)
  
  
  

- [mixer/v1/report.proto](#mixer/v1/report.proto)
    - [ReportRequest](#istio.mixer.v1.ReportRequest)
    - [ReportResponse](#istio.mixer.v1.ReportResponse)
  
  
  
  

- [mixer/v1/service.proto](#mixer/v1/service.proto)
  
  
  
    - [Mixer](#istio.mixer.v1.Mixer)
  

- [Scalar Value Types](#scalar-value-types)



<a name="mixer/v1/attributes.proto"/>
<p align="right"><a href="#top">Top</a></p>

## mixer/v1/attributes.proto



<a name="istio.mixer.v1.Attributes"/>

### Attributes
Attributes represents a set of typed name/value pairs. Many of Mixer&#39;s
API either consume and/or return attributes.

Istio uses attributes to control the runtime behavior of services running in the service mesh.
Attributes are named and typed pieces of metadata describing ingress and egress traffic and the
environment this traffic occurs in. An Istio attribute carries a specific piece
of information such as the error code of an API request, the latency of an API request, or the
original IP address of a TCP connection. For example:

```
request.path: xyz/abc
request.size: 234
request.time: 12:34:56.789 04/17/2017
source.ip: 192.168.0.1
target.service: example
```

A given Istio deployment has a fixed vocabulary of attributes that it understands.
The specific vocabulary is determined by the set of attribute producers being used
in the deployment. The primary attribute producer in Istio is Envoy, although
specialized Mixer adapters and services can also generate attributes.

The common baseline set of attributes available in most Istio deployments is defined
[here](https://istio.io/docs/reference/config/mixer/attribute-vocabulary.html).

Attributes are strongly typed. The supported attribute types are defined by
[ValueType](https://github.com/istio/api/blob/master/policy/v1beta1/value_type.proto).
Each type of value is encoded into one of the so-called transport types present
in this message.

Defines a map of attributes in uncompressed format.
Following places may use this message:
1) Configure Istio/Proxy with static per-proxy attributes, such as source.uid.
2) Service IDL definition to extract api attributes for active requests.
3) Forward attributes from client proxy to server proxy for HTTP requests.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| attributes | [Attributes.AttributesEntry](#istio.mixer.v1.Attributes.AttributesEntry) | repeated | A map of attribute name to its value. |






<a name="istio.mixer.v1.Attributes.AttributeValue"/>

### Attributes.AttributeValue
Specifies one attribute value with different type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| string_value | [string](#string) |  | Used for values of type STRING, DNS_NAME, EMAIL_ADDRESS, and URI |
| int64_value | [int64](#int64) |  | Used for values of type INT64 |
| double_value | [double](#double) |  | Used for values of type DOUBLE |
| bool_value | [bool](#bool) |  | Used for values of type BOOL |
| bytes_value | [bytes](#bytes) |  | Used for values of type BYTES |
| timestamp_value | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Used for values of type TIMESTAMP |
| duration_value | [google.protobuf.Duration](#google.protobuf.Duration) |  | Used for values of type DURATION |
| string_map_value | [Attributes.StringMap](#istio.mixer.v1.Attributes.StringMap) |  | Used for values of type STRING_MAP |






<a name="istio.mixer.v1.Attributes.AttributesEntry"/>

### Attributes.AttributesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [Attributes.AttributeValue](#istio.mixer.v1.Attributes.AttributeValue) |  |  |






<a name="istio.mixer.v1.Attributes.StringMap"/>

### Attributes.StringMap
Defines a string map.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entries | [Attributes.StringMap.EntriesEntry](#istio.mixer.v1.Attributes.StringMap.EntriesEntry) | repeated | Holds a set of name/value pairs. |






<a name="istio.mixer.v1.CompressedAttributes"/>

### CompressedAttributes
Defines a list of attributes in compressed format optimized for transport.
Within this message, strings are referenced using integer indices into
one of two string dictionaries. Positive integers index into the global
deployment-wide dictionary, whereas negative integers index into the message-level
dictionary instead. The message-level dictionary is carried by the
`words` field of this message, the deployment-wide dictionary is determined via
configuration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| words | [string](#string) | repeated | The message-level dictionary. |
| strings | [CompressedAttributes.StringsEntry](#istio.mixer.v1.CompressedAttributes.StringsEntry) | repeated | Holds attributes of type STRING, DNS_NAME, EMAIL_ADDRESS, URI |
| int64s | [CompressedAttributes.Int64sEntry](#istio.mixer.v1.CompressedAttributes.Int64sEntry) | repeated | Holds attributes of type INT64 |
| doubles | [CompressedAttributes.DoublesEntry](#istio.mixer.v1.CompressedAttributes.DoublesEntry) | repeated | Holds attributes of type DOUBLE |
| bools | [CompressedAttributes.BoolsEntry](#istio.mixer.v1.CompressedAttributes.BoolsEntry) | repeated | Holds attributes of type BOOL |
| timestamps | [CompressedAttributes.TimestampsEntry](#istio.mixer.v1.CompressedAttributes.TimestampsEntry) | repeated | Holds attributes of type TIMESTAMP |
| durations | [CompressedAttributes.DurationsEntry](#istio.mixer.v1.CompressedAttributes.DurationsEntry) | repeated | Holds attributes of type DURATION |
| bytes | [CompressedAttributes.BytesEntry](#istio.mixer.v1.CompressedAttributes.BytesEntry) | repeated | Holds attributes of type BYTES |
| string_maps | [CompressedAttributes.StringMapsEntry](#istio.mixer.v1.CompressedAttributes.StringMapsEntry) | repeated | Holds attributes of type STRING_MAP |






<a name="istio.mixer.v1.CompressedAttributes.BoolsEntry"/>

### CompressedAttributes.BoolsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [sint32](#sint32) |  |  |
| value | [bool](#bool) |  |  |






<a name="istio.mixer.v1.CompressedAttributes.BytesEntry"/>

### CompressedAttributes.BytesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [sint32](#sint32) |  |  |
| value | [bytes](#bytes) |  |  |






<a name="istio.mixer.v1.CompressedAttributes.DoublesEntry"/>

### CompressedAttributes.DoublesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [sint32](#sint32) |  |  |
| value | [double](#double) |  |  |






<a name="istio.mixer.v1.CompressedAttributes.DurationsEntry"/>

### CompressedAttributes.DurationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [sint32](#sint32) |  |  |
| value | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="istio.mixer.v1.CompressedAttributes.Int64sEntry"/>

### CompressedAttributes.Int64sEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [sint32](#sint32) |  |  |
| value | [int64](#int64) |  |  |






<a name="istio.mixer.v1.CompressedAttributes.StringMapsEntry"/>

### CompressedAttributes.StringMapsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [sint32](#sint32) |  |  |
| value | [StringMap](#istio.mixer.v1.StringMap) |  |  |






<a name="istio.mixer.v1.CompressedAttributes.StringsEntry"/>

### CompressedAttributes.StringsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [sint32](#sint32) |  |  |
| value | [sint32](#sint32) |  |  |






<a name="istio.mixer.v1.CompressedAttributes.TimestampsEntry"/>

### CompressedAttributes.TimestampsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [sint32](#sint32) |  |  |
| value | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="istio.mixer.v1.StringMap"/>

### StringMap
A map of string to string. The keys and values in this map are dictionary
indices (see the [Attributes][istio.mixer.v1.CompressedAttributes] message for an explanation)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entries | [StringMap.EntriesEntry](#istio.mixer.v1.StringMap.EntriesEntry) | repeated | Holds a set of name/value pairs. |






<a name="istio.mixer.v1.StringMap.EntriesEntry"/>

### StringMap.EntriesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [sint32](#sint32) |  |  |
| value | [sint32](#sint32) |  |  |





 

 

 

 



<a name="mixer/v1/check.proto"/>
<p align="right"><a href="#top">Top</a></p>

## mixer/v1/check.proto



<a name="istio.mixer.v1.CheckRequest"/>

### CheckRequest
Used to get a thumbs-up/thumbs-down before performing an action.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| attributes | [CompressedAttributes](#istio.mixer.v1.CompressedAttributes) |  | The attributes to use for this request.  Mixer&#39;s configuration determines how these attributes are used to establish the result returned in the response. |
| global_word_count | [uint32](#uint32) |  | The number of words in the global dictionary, used with to populate the attributes. This value is used as a quick way to determine whether the client is using a dictionary that the server understands. |
| deduplication_id | [string](#string) |  | Used for deduplicating `Check` calls in the case of failed RPCs and retries. This should be a UUID per call, where the same UUID is used for retries of the same call. |
| quotas | [CheckRequest.QuotasEntry](#istio.mixer.v1.CheckRequest.QuotasEntry) | repeated | The individual quotas to allocate |






<a name="istio.mixer.v1.CheckRequest.QuotaParams"/>

### CheckRequest.QuotaParams
parameters for a quota allocation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| amount | [int64](#int64) |  | Amount of quota to allocate |
| best_effort | [bool](#bool) |  | When true, supports returning less quota than what was requested. |






<a name="istio.mixer.v1.CheckRequest.QuotasEntry"/>

### CheckRequest.QuotasEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [CheckRequest.QuotaParams](#istio.mixer.v1.CheckRequest.QuotaParams) |  |  |






<a name="istio.mixer.v1.CheckResponse"/>

### CheckResponse
The response generated by the Check method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| precondition | [CheckResponse.PreconditionResult](#istio.mixer.v1.CheckResponse.PreconditionResult) |  | The precondition check results. |
| quotas | [CheckResponse.QuotasEntry](#istio.mixer.v1.CheckResponse.QuotasEntry) | repeated | The resulting quota, one entry per requested quota. |






<a name="istio.mixer.v1.CheckResponse.PreconditionResult"/>

### CheckResponse.PreconditionResult
Expresses the result of a precondition check.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  | A status code of OK indicates all preconditions were satisfied. Any other code indicates not all preconditions were satisfied and details describe why. |
| valid_duration | [google.protobuf.Duration](#google.protobuf.Duration) |  | The amount of time for which this result can be considered valid. |
| valid_use_count | [int32](#int32) |  | The number of uses for which this result can be considered valid. |
| referenced_attributes | [ReferencedAttributes](#istio.mixer.v1.ReferencedAttributes) |  | The total set of attributes that were used in producing the result along with matching conditions. |
| route_directive | [RouteDirective](#istio.mixer.v1.RouteDirective) |  | An optional routing directive, used to manipulate the traffic metadata whenever all preconditions are satisfied. |






<a name="istio.mixer.v1.CheckResponse.QuotaResult"/>

### CheckResponse.QuotaResult
Expresses the result of a quota allocation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| valid_duration | [google.protobuf.Duration](#google.protobuf.Duration) |  | The amount of time for which this result can be considered valid. |
| granted_amount | [int64](#int64) |  | The amount of granted quota. When `QuotaParams.best_effort` is true, this will be &gt;= 0. If `QuotaParams.best_effort` is false, this will be either 0 or &gt;= `QuotaParams.amount`. |
| referenced_attributes | [ReferencedAttributes](#istio.mixer.v1.ReferencedAttributes) |  | The total set of attributes that were used in producing the result along with matching conditions. |






<a name="istio.mixer.v1.CheckResponse.QuotasEntry"/>

### CheckResponse.QuotasEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [CheckResponse.QuotaResult](#istio.mixer.v1.CheckResponse.QuotaResult) |  |  |






<a name="istio.mixer.v1.HeaderOperation"/>

### HeaderOperation
Operation on HTTP headers to replace, append, or remove a header. Header
names are normalized to lower-case with dashes, e.g.  &#34;x-request-id&#34;.
Pseudo-headers &#34;:path&#34;, &#34;:authority&#34;, and &#34;:method&#34; are supported to modify
the request headers.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Header name. |
| value | [string](#string) |  | Header value. |
| operation | [HeaderOperation.Operation](#istio.mixer.v1.HeaderOperation.Operation) |  | Header operation. |






<a name="istio.mixer.v1.ReferencedAttributes"/>

### ReferencedAttributes
Describes the attributes that were used to determine the response.
This can be used to construct a response cache.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| words | [string](#string) | repeated | The message-level dictionary. Refer to [CompressedAttributes][istio.mixer.v1.CompressedAttributes] for information on using dictionaries. |
| attribute_matches | [ReferencedAttributes.AttributeMatch](#istio.mixer.v1.ReferencedAttributes.AttributeMatch) | repeated | Describes a set of attributes. |






<a name="istio.mixer.v1.ReferencedAttributes.AttributeMatch"/>

### ReferencedAttributes.AttributeMatch
Describes a single attribute match.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [sint32](#sint32) |  | The name of the attribute. This is a dictionary index encoded in a manner identical to all strings in the [CompressedAttributes][istio.mixer.v1.CompressedAttributes] message. |
| condition | [ReferencedAttributes.Condition](#istio.mixer.v1.ReferencedAttributes.Condition) |  | The kind of match against the attribute value. |
| regex | [string](#string) |  | If a REGEX condition is provided for a STRING_MAP attribute, clients should use the regex value to match against map keys. |
| map_key | [sint32](#sint32) |  | A key in a STRING_MAP. When multiple keys from a STRING_MAP attribute were referenced, there will be multiple AttributeMatch messages with different map_key values. Values for map_key SHOULD be ignored for attributes that are not STRING_MAP.  Indices for the keys are used (taken either from the message dictionary from the `words` field or the global dictionary).  If no map_key value is provided for a STRING_MAP attribute, the entire STRING_MAP will be used. |






<a name="istio.mixer.v1.RouteDirective"/>

### RouteDirective
Expresses the routing manipulation actions to be performed on behalf of
Mixer in response to a successful precondition check.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_header_operations | [HeaderOperation](#istio.mixer.v1.HeaderOperation) | repeated | Operations on the request headers. |
| response_header_operations | [HeaderOperation](#istio.mixer.v1.HeaderOperation) | repeated | Operations on the response headers. |
| direct_response_code | [uint32](#uint32) |  | If set, enables a direct response without proxying the request to the routing destination. Required to be a value in the 2xx or 3xx range. |
| direct_response_body | [string](#string) |  | Supplies the response body for the direct response. If this setting is omitted, no body is included in the generated response. |





 


<a name="istio.mixer.v1.HeaderOperation.Operation"/>

### HeaderOperation.Operation
Operation type.

| Name | Number | Description |
| ---- | ------ | ----------- |
| REPLACE | 0 | replaces the header with the given name |
| REMOVE | 1 | removes the header with the given name (the value is ignored) |
| APPEND | 2 | appends the value to the header value, or sets it if not present |



<a name="istio.mixer.v1.ReferencedAttributes.Condition"/>

### ReferencedAttributes.Condition
How an attribute&#39;s value was matched

| Name | Number | Description |
| ---- | ------ | ----------- |
| CONDITION_UNSPECIFIED | 0 | should not occur |
| ABSENCE | 1 | match when attribute doesn&#39;t exist |
| EXACT | 2 | match when attribute value is an exact byte-for-byte match |
| REGEX | 3 | match when attribute value matches the included regex |


 

 

 



<a name="mixer/v1/report.proto"/>
<p align="right"><a href="#top">Top</a></p>

## mixer/v1/report.proto



<a name="istio.mixer.v1.ReportRequest"/>

### ReportRequest
Used to report telemetry after performing one or more actions.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| attributes | [CompressedAttributes](#istio.mixer.v1.CompressedAttributes) | repeated | The attributes to use for this request.  Each `Attributes` element represents the state of a single action. Multiple actions can be provided in a single message in order to improve communication efficiency. The client can accumulate a set of actions and send them all in one single message.  Although each `Attributes` message is semantically treated as an independent stand-alone entity unrelated to the other attributes within the message, this message format leverages delta-encoding between attribute messages in order to substantially reduce the request size and improve end-to-end efficiency. Each individual set of attributes is used to modify the previous set. This eliminates the need to redundantly send the same attributes multiple times over within a single request.  If a client is not sophisticated and doesn&#39;t want to use delta-encoding, a degenerate case is to include all attributes in every individual message. |
| default_words | [string](#string) | repeated | The default message-level dictionary for all the attributes. Individual attribute messages can have their own dictionaries, but if they don&#39;t then this set of words, if it is provided, is used instead.  This makes it possible to share the same dictionary for all attributes in this request, which can substantially reduce the overall request size. |
| global_word_count | [uint32](#uint32) |  | The number of words in the global dictionary. To detect global dictionary out of sync between client and server. |






<a name="istio.mixer.v1.ReportResponse"/>

### ReportResponse
Used to carry responses to telemetry reports





 

 

 

 



<a name="mixer/v1/service.proto"/>
<p align="right"><a href="#top">Top</a></p>

## mixer/v1/service.proto


 

 

 


<a name="istio.mixer.v1.Mixer"/>

### Mixer
Mixer provides three core features:

- *Precondition Checking*. Enables callers to verify a number of preconditions
before responding to an incoming request from a service consumer.
Preconditions can include whether the service consumer is properly
authenticated, is on the service’s whitelist, passes ACL checks, and more.

- *Quota Management*. Enables services to allocate and free quota on a number
of dimensions, Quotas are used as a relatively simple resource management tool
to provide some fairness between service consumers when contending for limited
resources. Rate limits are examples of quotas.

- *Telemetry Reporting*. Enables services to report logging and monitoring.
In the future, it will also enable tracing and billing streams intended for
both the service operator as well as for service consumers.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Check | [CheckRequest](#istio.mixer.v1.CheckRequest) | [CheckResponse](#istio.mixer.v1.CheckRequest) | Checks preconditions and allocate quota before performing an operation. The preconditions enforced depend on the set of supplied attributes and the active configuration. |
| Report | [ReportRequest](#istio.mixer.v1.ReportRequest) | [ReportResponse](#istio.mixer.v1.ReportRequest) | Reports telemetry, such as logs and metrics. The reported information depends on the set of supplied attributes and the active configuration. |

 



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

