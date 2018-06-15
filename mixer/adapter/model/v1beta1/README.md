# Protocol Documentation
<a name="top"/>

## Table of Contents

- [mixer/adapter/model/v1beta1/check.proto](#mixer/adapter/model/v1beta1/check.proto)
    - [CheckResult](#istio.mixer.adapter.model.v1beta1.CheckResult)
  
  
  
  

- [mixer/adapter/model/v1beta1/extensions.proto](#mixer/adapter/model/v1beta1/extensions.proto)
  
    - [TemplateVariety](#istio.mixer.adapter.model.v1beta1.TemplateVariety)
  
    - [File-level Extensions](#mixer/adapter/model/v1beta1/extensions.proto-extensions)
    - [File-level Extensions](#mixer/adapter/model/v1beta1/extensions.proto-extensions)
  
  

- [mixer/adapter/model/v1beta1/info.proto](#mixer/adapter/model/v1beta1/info.proto)
    - [Info](#istio.mixer.adapter.model.v1beta1.Info)
  
  
  
  

- [mixer/adapter/model/v1beta1/infrastructure_backend.proto](#mixer/adapter/model/v1beta1/infrastructure_backend.proto)
    - [CloseSessionRequest](#istio.mixer.adapter.model.v1beta1.CloseSessionRequest)
    - [CloseSessionResponse](#istio.mixer.adapter.model.v1beta1.CloseSessionResponse)
    - [CreateSessionRequest](#istio.mixer.adapter.model.v1beta1.CreateSessionRequest)
    - [CreateSessionRequest.InferredTypesEntry](#istio.mixer.adapter.model.v1beta1.CreateSessionRequest.InferredTypesEntry)
    - [CreateSessionResponse](#istio.mixer.adapter.model.v1beta1.CreateSessionResponse)
    - [ValidateRequest](#istio.mixer.adapter.model.v1beta1.ValidateRequest)
    - [ValidateRequest.InferredTypesEntry](#istio.mixer.adapter.model.v1beta1.ValidateRequest.InferredTypesEntry)
    - [ValidateResponse](#istio.mixer.adapter.model.v1beta1.ValidateResponse)
  
  
  
    - [InfrastructureBackend](#istio.mixer.adapter.model.v1beta1.InfrastructureBackend)
  

- [mixer/adapter/model/v1beta1/quota.proto](#mixer/adapter/model/v1beta1/quota.proto)
    - [QuotaRequest](#istio.mixer.adapter.model.v1beta1.QuotaRequest)
    - [QuotaRequest.QuotaParams](#istio.mixer.adapter.model.v1beta1.QuotaRequest.QuotaParams)
    - [QuotaRequest.QuotasEntry](#istio.mixer.adapter.model.v1beta1.QuotaRequest.QuotasEntry)
    - [QuotaResult](#istio.mixer.adapter.model.v1beta1.QuotaResult)
    - [QuotaResult.QuotasEntry](#istio.mixer.adapter.model.v1beta1.QuotaResult.QuotasEntry)
    - [QuotaResult.Result](#istio.mixer.adapter.model.v1beta1.QuotaResult.Result)
  
  
  
  

- [mixer/adapter/model/v1beta1/report.proto](#mixer/adapter/model/v1beta1/report.proto)
    - [ReportResult](#istio.mixer.adapter.model.v1beta1.ReportResult)
  
  
  
  

- [mixer/adapter/model/v1beta1/template.proto](#mixer/adapter/model/v1beta1/template.proto)
    - [Template](#istio.mixer.adapter.model.v1beta1.Template)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="mixer/adapter/model/v1beta1/check.proto"/>
<p align="right"><a href="#top">Top</a></p>

## mixer/adapter/model/v1beta1/check.proto



<a name="istio.mixer.adapter.model.v1beta1.CheckResult"/>

### CheckResult
Expresses the result of a precondition check.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  | A status code of OK indicates preconditions were satisfied. Any other code indicates preconditions were not satisfied and details describe why. |
| valid_duration | [google.protobuf.Duration](#google.protobuf.Duration) |  | The amount of time for which this result can be considered valid. |
| valid_use_count | [int32](#int32) |  | The number of uses for which this result can be considered valid. |





 

 

 

 



<a name="mixer/adapter/model/v1beta1/extensions.proto"/>
<p align="right"><a href="#top">Top</a></p>

## mixer/adapter/model/v1beta1/extensions.proto


 


<a name="istio.mixer.adapter.model.v1beta1.TemplateVariety"/>

### TemplateVariety
The available varieties of templates, controlling the semantics of what an adapter does with each instance.

| Name | Number | Description |
| ---- | ------ | ----------- |
| TEMPLATE_VARIETY_CHECK | 0 | Makes the template applicable for Mixer&#39;s check calls. Instances of such template are created during report calls in Mixer and passed to the handlers based on the rule configurations. |
| TEMPLATE_VARIETY_REPORT | 1 | Makes the template applicable for Mixer&#39;s report calls. Instances of such template are created during check calls in Mixer and passed to the handlers based on the rule configurations. |
| TEMPLATE_VARIETY_QUOTA | 2 | Makes the template applicable for Mixer&#39;s quota calls. Instances of such template are created during quota check calls in Mixer and passed to the handlers based on the rule configurations. |
| TEMPLATE_VARIETY_ATTRIBUTE_GENERATOR | 3 | Makes the template applicable for Mixer&#39;s attribute generation phase. Instances of such template are created during pre-processing attribute generation phase and passed to the handlers based on the rule configurations. |


 


<a name="mixer/adapter/model/v1beta1/extensions.proto-extensions"/>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| template_name | string | .google.protobuf.FileOptions | 72295888 | Optional: option for the template name. If not specified, the last segment of the template proto&#39;s package name is used to derive the template name. |
| template_variety | TemplateVariety | .google.protobuf.FileOptions | 72295727 | Required: option for the TemplateVariety. |

 

 



<a name="mixer/adapter/model/v1beta1/info.proto"/>
<p align="right"><a href="#top">Top</a></p>

## mixer/adapter/model/v1beta1/info.proto



<a name="istio.mixer.adapter.model.v1beta1.Info"/>

### Info
Info describes an adapter or a backend that wants to provide telemetry and policy functionality to Mixer as an
out of process adapter.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the adapter. It must be an RFC 1035 compatible DNS label matching the `^[a-z]([-a-z0-9]*[a-z0-9])?$` regular expression. Name is used in Istio configuration, therefore it should be descriptive but short. example: denier Vendor adapters should use a vendor prefix. example: mycompany-denier |
| description | [string](#string) |  | User-friendly description of the adapter. |
| templates | [string](#string) | repeated | Names of the templates the adapter supports. |
| config | [string](#string) |  | Base64 encoded proto descriptor of the adapter configuration. |
| session_based | [bool](#bool) |  | True if backend has implemented the [InfrastructureBackend](https://github.com/istio/api/blob/master/mixer/adapter/model/v1beta1/infrastructure_backend.proto) service; false otherwise.  If true, during configuration time, Mixer calls the `InfrastructureBackend`&#39; rpcs to validate and pass the handler configuration. And during request-time, Mixer does not pass the handler configuration, and only passes the template-specific Instance payload using the template-specific handle service (Example [HandlerMetricService](https://github.com/istio/istio/blob/master/mixer/template/metric/template_handler_service.proto), [HandlerListEntryService](https://github.com/istio/istio/blob/master/mixer/template/logentry/template_handler_service.proto), [HandleQuotaService](https://github.com/istio/istio/blob/master/mixer/template/quota/template_handler_service.proto) and more). If `session_based` is false, Mixer does not expect backend to implement `InfrastructureBackend` service, and only communicates with the backends during request-time through the template-specific handle service. Without `InfrastructureBackend` service, Mixer passes the handler configuration on each call during request-time. |





 

 

 

 



<a name="mixer/adapter/model/v1beta1/infrastructure_backend.proto"/>
<p align="right"><a href="#top">Top</a></p>

## mixer/adapter/model/v1beta1/infrastructure_backend.proto



<a name="istio.mixer.adapter.model.v1beta1.CloseSessionRequest"/>

### CloseSessionRequest
Request message for `CloseSession` method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_id | [string](#string) |  | Id of the session to be closed. |






<a name="istio.mixer.adapter.model.v1beta1.CloseSessionResponse"/>

### CloseSessionResponse
Response message for `CloseSession` method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  | The success/failure status of close session call. |






<a name="istio.mixer.adapter.model.v1beta1.CreateSessionRequest"/>

### CreateSessionRequest
Request message for `CreateSession` method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| adapter_config | [google.protobuf.Any](#google.protobuf.Any) |  | Adapter specific configuration. |
| inferred_types | [CreateSessionRequest.InferredTypesEntry](#istio.mixer.adapter.model.v1beta1.CreateSessionRequest.InferredTypesEntry) | repeated | Map of instance names to their template-specific inferred type. |






<a name="istio.mixer.adapter.model.v1beta1.CreateSessionRequest.InferredTypesEntry"/>

### CreateSessionRequest.InferredTypesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="istio.mixer.adapter.model.v1beta1.CreateSessionResponse"/>

### CreateSessionResponse
Response message for `CreateSession` method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_id | [string](#string) |  | Id of the created session. |
| status | [google.rpc.Status](#google.rpc.Status) |  | The success/failure status of create session call. |






<a name="istio.mixer.adapter.model.v1beta1.ValidateRequest"/>

### ValidateRequest
Request message for `Validate` method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| adapter_config | [google.protobuf.Any](#google.protobuf.Any) |  | Adapter specific configuration. |
| inferred_types | [ValidateRequest.InferredTypesEntry](#istio.mixer.adapter.model.v1beta1.ValidateRequest.InferredTypesEntry) | repeated | Map of instance names to their template-specific inferred type. |






<a name="istio.mixer.adapter.model.v1beta1.ValidateRequest.InferredTypesEntry"/>

### ValidateRequest.InferredTypesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="istio.mixer.adapter.model.v1beta1.ValidateResponse"/>

### ValidateResponse
Response message for `Validate` method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [google.rpc.Status](#google.rpc.Status) |  | The success/failure status of validation call. |





 

 

 


<a name="istio.mixer.adapter.model.v1beta1.InfrastructureBackend"/>

### InfrastructureBackend
`InfrastructureBackend` is implemented by backends that wants to provide telemetry and policy functionality to Mixer as an
out of process adapter.

`InfrastructureBackend` allows Mixer and the backends to have a session based model. In a session based model, Mixer passes
the relevant configuration state to the backend only once and estabilshes a session using a session identifier.
All future communications between Mixer and the backend uses the session identifier which refers to the previously
passed in configuration data.

For a given `InfrastructureBackend`, Mixer calls the `Validate` function, followed by `CreateSession`. The `session_id`
returned from `CreateSession` is used to make subsequent request-time Handle calls and the eventual `CloseSession` function.
Mixer assumes that, given the `session_id`, the backend can retrieve the necessary context to serve the
Handle calls that contains the request-time payload (template-specific instance protobuf).

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Validate | [ValidateRequest](#istio.mixer.adapter.model.v1beta1.ValidateRequest) | [ValidateResponse](#istio.mixer.adapter.model.v1beta1.ValidateRequest) | Validates the handler configuration along with the template-specific instances that would be routed to that handler. The `CreateSession` for a specific handler configuration is invoked only if its associated `Validate` call has returned success. |
| CreateSession | [CreateSessionRequest](#istio.mixer.adapter.model.v1beta1.CreateSessionRequest) | [CreateSessionResponse](#istio.mixer.adapter.model.v1beta1.CreateSessionRequest) | Creates a session for a given handler configuration and the template-specific instances that would be routed to that handler. For every handler configuration, Mixer creates a separate session by invoking `CreateSession` on the backend.  `CreateSessionRequest` contains the adapter specific handler configuration and the inferred type information about the instances the handler would receive during request processing.  `CreateSession` must return a `session_id` which Mixer uses to invoke template-specific Handle functions during request processing. The `session_id` provides the Handle functions a way to retrieve the necessary configuration associated with the session. Upon Mixer configuration change, Mixer will re-invoke `CreateSession` for all handler configurations whose existing sessions are invalidated or didn&#39;t existed.  Backend is allowed to return the same session id if given the same configuration block. This would happen when multiple instances of Mixer in a deployment all create sessions with the same configuration. Note that given individial instances of Mixer can call `CloseSession`, reusing `session_id` by the backend assumes that the backend is doing reference counting.  If the backend couldn&#39;t create a session for a specific handler configuration and returns non S_OK status, Mixer will not make request-time Handle calls associated with that handler configuration. |
| CloseSession | [CloseSessionRequest](#istio.mixer.adapter.model.v1beta1.CloseSessionRequest) | [CloseSessionResponse](#istio.mixer.adapter.model.v1beta1.CloseSessionRequest) | Closes the session associated with the `session_id`. Mixer closes a session when its associated handler configuration or the instance configuration changes. Backend is supposed to cleanup all the resources associated with the session_id referenced by CloseSessionRequest. |

 



<a name="mixer/adapter/model/v1beta1/quota.proto"/>
<p align="right"><a href="#top">Top</a></p>

## mixer/adapter/model/v1beta1/quota.proto



<a name="istio.mixer.adapter.model.v1beta1.QuotaRequest"/>

### QuotaRequest
Expresses the quota allocation request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| quotas | [QuotaRequest.QuotasEntry](#istio.mixer.adapter.model.v1beta1.QuotaRequest.QuotasEntry) | repeated | The individual quotas to allocate |






<a name="istio.mixer.adapter.model.v1beta1.QuotaRequest.QuotaParams"/>

### QuotaRequest.QuotaParams
parameters for a quota allocation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| amount | [int64](#int64) |  | Amount of quota to allocate |
| best_effort | [bool](#bool) |  | When true, supports returning less quota than what was requested. |






<a name="istio.mixer.adapter.model.v1beta1.QuotaRequest.QuotasEntry"/>

### QuotaRequest.QuotasEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [QuotaRequest.QuotaParams](#istio.mixer.adapter.model.v1beta1.QuotaRequest.QuotaParams) |  |  |






<a name="istio.mixer.adapter.model.v1beta1.QuotaResult"/>

### QuotaResult
Expresses the result of multiple quota allocations.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| quotas | [QuotaResult.QuotasEntry](#istio.mixer.adapter.model.v1beta1.QuotaResult.QuotasEntry) | repeated | The resulting quota, one entry per requested quota. |






<a name="istio.mixer.adapter.model.v1beta1.QuotaResult.QuotasEntry"/>

### QuotaResult.QuotasEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [QuotaResult.Result](#istio.mixer.adapter.model.v1beta1.QuotaResult.Result) |  |  |






<a name="istio.mixer.adapter.model.v1beta1.QuotaResult.Result"/>

### QuotaResult.Result
Expresses the result of a quota allocation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| valid_duration | [google.protobuf.Duration](#google.protobuf.Duration) |  | The amount of time for which this result can be considered valid. |
| granted_amount | [int64](#int64) |  | The amount of granted quota. When `QuotaParams.best_effort` is true, this will be &gt;= 0. If `QuotaParams.best_effort` is false, this will be either 0 or &gt;= `QuotaParams.amount`. |





 

 

 

 



<a name="mixer/adapter/model/v1beta1/report.proto"/>
<p align="right"><a href="#top">Top</a></p>

## mixer/adapter/model/v1beta1/report.proto



<a name="istio.mixer.adapter.model.v1beta1.ReportResult"/>

### ReportResult
Expresses the result of a report call.





 

 

 

 



<a name="mixer/adapter/model/v1beta1/template.proto"/>
<p align="right"><a href="#top">Top</a></p>

## mixer/adapter/model/v1beta1/template.proto



<a name="istio.mixer.adapter.model.v1beta1.Template"/>

### Template
Template provides the details of a mixer template.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| descriptor | [string](#string) |  | Base64 encoded proto descriptor of the template. |





 

 

 

 



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

