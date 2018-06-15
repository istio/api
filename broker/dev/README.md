# Protocol Documentation
<a name="top"/>

## Table of Contents

- [broker/dev/service_class.proto](#broker/dev/service_class.proto)
    - [CatalogEntry](#istio.broker.dev.CatalogEntry)
    - [Deployment](#istio.broker.dev.Deployment)
    - [ServiceClass](#istio.broker.dev.ServiceClass)
  
  
  
  

- [broker/dev/service_plan.proto](#broker/dev/service_plan.proto)
    - [CatalogPlan](#istio.broker.dev.CatalogPlan)
    - [ServicePlan](#istio.broker.dev.ServicePlan)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="broker/dev/service_class.proto"/>
<p align="right"><a href="#top">Top</a></p>

## broker/dev/service_class.proto
Copyright 2017 Istio Authors

  Licensed under the Apache License, Version 2.0 (the &#34;License&#34;);
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an &#34;AS IS&#34; BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.


<a name="istio.broker.dev.CatalogEntry"/>

### CatalogEntry
$hide_from_docs
CatalogEntry defines listing information for this service within the exposed
catalog.  The message is a subset of OSBI service fields defined in
https://github.com/openservicebrokerapi


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Required. Public service name. |
| id | [string](#string) |  | Required. Public unique service guid. |
| description | [string](#string) |  | Required. Public short service description. |






<a name="istio.broker.dev.Deployment"/>

### Deployment
$hide_from_docs
Deployment defines how the service instances are deployed.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| instance | [string](#string) |  | For truely multi-tenant service, the deployed service instance name. |






<a name="istio.broker.dev.ServiceClass"/>

### ServiceClass
$hide_from_docs
ServiceClass defines a service that are exposed to Istio service consumers.
The service is linked into one or more ServicePlan.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deployment | [Deployment](#istio.broker.dev.Deployment) |  | Required. Istio deployment spec for the service class. |
| entry | [CatalogEntry](#istio.broker.dev.CatalogEntry) |  | Required. Listing information for the public catalog. |





 

 

 

 



<a name="broker/dev/service_plan.proto"/>
<p align="right"><a href="#top">Top</a></p>

## broker/dev/service_plan.proto
Copyright 2017 Istio Authors

  Licensed under the Apache License, Version 2.0 (the &#34;License&#34;);
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an &#34;AS IS&#34; BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.


<a name="istio.broker.dev.CatalogPlan"/>

### CatalogPlan
$hide_from_docs
CatalogPlan defines listing information for this service plan within the
exposed catalog.  The message is a subset of OSBI plan fields defined in
https://github.com/openservicebrokerapi


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Required. Public service plan name. |
| id | [string](#string) |  | Required. Public unique service plan guid. |
| description | [string](#string) |  | Required. Public short service plan description. |






<a name="istio.broker.dev.ServicePlan"/>

### ServicePlan
$hide_from_docs
ServicePlan defines the type of services available to Istio service
consumers.  One or more services are included in a plan. The plan is flexible
and subject to change along with business requirements.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| plan | [CatalogPlan](#istio.broker.dev.CatalogPlan) |  | Required. Public plan information. |
| services | [string](#string) | repeated | Required. List of the Keys of serviceclass config instance that are included in the plan. ServiceClass is a type of CRD resource. |





 

 

 

 



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

