# Protocol Documentation
<a name="top"/>

## Table of Contents

- [rbac/v1alpha1/rbac.proto](#rbac/v1alpha1/rbac.proto)
    - [AccessRule](#istio.rbac.v1alpha1.AccessRule)
    - [AccessRule.Constraint](#istio.rbac.v1alpha1.AccessRule.Constraint)
    - [RbacConfig](#istio.rbac.v1alpha1.RbacConfig)
    - [RbacConfig.Target](#istio.rbac.v1alpha1.RbacConfig.Target)
    - [RoleRef](#istio.rbac.v1alpha1.RoleRef)
    - [ServiceRole](#istio.rbac.v1alpha1.ServiceRole)
    - [ServiceRoleBinding](#istio.rbac.v1alpha1.ServiceRoleBinding)
    - [Subject](#istio.rbac.v1alpha1.Subject)
    - [Subject.PropertiesEntry](#istio.rbac.v1alpha1.Subject.PropertiesEntry)
  
    - [EnforcementMode](#istio.rbac.v1alpha1.EnforcementMode)
    - [RbacConfig.Mode](#istio.rbac.v1alpha1.RbacConfig.Mode)
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="rbac/v1alpha1/rbac.proto"/>
<p align="right"><a href="#top">Top</a></p>

## rbac/v1alpha1/rbac.proto



<a name="istio.rbac.v1alpha1.AccessRule"/>

### AccessRule
AccessRule defines a permission to access a list of services.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| services | [string](#string) | repeated | Required. A list of service names. Exact match, prefix match, and suffix match are supported for service names. For example, the service name &#34;bookstore.mtv.cluster.local&#34; matches &#34;bookstore.mtv.cluster.local&#34; (exact match), or &#34;bookstore*&#34; (prefix match), or &#34;*.mtv.cluster.local&#34; (suffix match). If set to [&#34;*&#34;], it refers to all services in the namespace. |
| paths | [string](#string) | repeated | Optional. A list of HTTP paths or gRPC methods. gRPC methods must be presented as fully-qualified name in the form of packageName.serviceName/methodName. Exact match, prefix match, and suffix match are supported for paths. For example, the path &#34;/books/review&#34; matches &#34;/books/review&#34; (exact match), or &#34;/books/*&#34; (prefix match), or &#34;*/review&#34; (suffix match). If not specified, it applies to any path. |
| methods | [string](#string) | repeated | Optional. A list of HTTP methods (e.g., &#34;GET&#34;, &#34;POST&#34;). It is ignored in gRPC case because the value is always &#34;POST&#34;. If set to [&#34;*&#34;] or not specified, it applies to any method. |
| constraints | [AccessRule.Constraint](#istio.rbac.v1alpha1.AccessRule.Constraint) | repeated | Optional. Extra constraints in the ServiceRole specification. The above ServiceRole examples shows an example of constraint &#34;version&#34;. |






<a name="istio.rbac.v1alpha1.AccessRule.Constraint"/>

### AccessRule.Constraint
Definition of a custom constraint. The key of a custom constraint must match
one of the &#34;properties&#34; in the &#34;action&#34; part of the &#34;authorization&#34; template
(https://github.com/istio/istio/blob/master/mixer/template/authorization/template.proto).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  | Key of the constraint. |
| values | [string](#string) | repeated | List of valid values for the constraint. Exact match, prefix match, and suffix match are supported for constraint values. For example, the value &#34;v1alpha2&#34; matches &#34;v1alpha2&#34; (exact match), or &#34;v1*&#34; (prefix match), or &#34;*alpha2&#34; (suffix match). |






<a name="istio.rbac.v1alpha1.RbacConfig"/>

### RbacConfig
RbacConfig defines the global config to control Istio RBAC behavior.
This Custom Resource is a singleton where only one Custom Resource should be created globally in
the mesh and the namespace should be the same to other Istio components, which usually is istio-system.
Note: This is enforced in both istioctl and server side, new Custom Resource will be rejected if found any
existing one, the user should either delete the existing one or change the existing one directly.

Below is an example of RbacConfig object &#34;istio-rbac-config&#34; which enables Istio RBAC for all
services in the default namespace.

```yaml
apiVersion: &#34;config.istio.io/v1alpha1&#34;
kind: RbacConfig
metadata:
  name: istio-rbac-config
  namespace: istio-system
spec:
  mode: ON_WITH_INCLUSION
  inclusion:
    namespaces: [ &#34;default&#34; ]
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mode | [RbacConfig.Mode](#istio.rbac.v1alpha1.RbacConfig.Mode) |  | Istio RBAC mode. |
| inclusion | [RbacConfig.Target](#istio.rbac.v1alpha1.RbacConfig.Target) |  | A list of services or namespaces that should be enforced by Istio RBAC policies. Note: This field have effect only when mode is ON_WITH_INCLUSION and will be ignored for any other modes. |
| exclusion | [RbacConfig.Target](#istio.rbac.v1alpha1.RbacConfig.Target) |  | A list of services or namespaces that should not be enforced by Istio RBAC policies. Note: This field have effect only when mode is ON_WITH_EXCLUSION and will be ignored for any other modes. |






<a name="istio.rbac.v1alpha1.RbacConfig.Target"/>

### RbacConfig.Target
Target defines a list of services or namespaces.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| services | [string](#string) | repeated | A list of services. |
| namespaces | [string](#string) | repeated | A list of namespaces. |






<a name="istio.rbac.v1alpha1.RoleRef"/>

### RoleRef
RoleRef refers to a role object.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind | [string](#string) |  | Required. The type of the role being referenced. Currently, &#34;ServiceRole&#34; is the only supported value for &#34;kind&#34;. |
| name | [string](#string) |  | Required. The name of the ServiceRole object being referenced. The ServiceRole object must be in the same namespace as the ServiceRoleBinding object. |






<a name="istio.rbac.v1alpha1.ServiceRole"/>

### ServiceRole
ServiceRole specification contains a list of access rules (permissions).
This represent the &#34;Spec&#34; part of the ServiceRole object. The name and namespace
of the ServiceRole is specified in &#34;metadata&#34; section of the ServiceRole object.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rules | [AccessRule](#istio.rbac.v1alpha1.AccessRule) | repeated | Required. The set of access rules (permissions) that the role has. |
| mode | [EnforcementMode](#istio.rbac.v1alpha1.EnforcementMode) |  | $hide_from_docs Indicates enforcement mode of the ServiceRole config. If ServiceRole is in PERMISSIVE mode, all ServiceRoleBindings that refers to it will be treated as PERMISSIVE mode. If ServiceRole is in ENFORCED mode, ServiceRoleBindings that refers to it will decide their own enforcement modes. |






<a name="istio.rbac.v1alpha1.ServiceRoleBinding"/>

### ServiceRoleBinding
ServiceRoleBinding assigns a ServiceRole to a list of subjects.
This represents the &#34;Spec&#34; part of the ServiceRoleBinding object. The name and namespace
of the ServiceRoleBinding is specified in &#34;metadata&#34; section of the ServiceRoleBinding
object.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subjects | [Subject](#istio.rbac.v1alpha1.Subject) | repeated | Required. List of subjects that are assigned the ServiceRole object. |
| roleRef | [RoleRef](#istio.rbac.v1alpha1.RoleRef) |  | Required. Reference to the ServiceRole object. |
| mode | [EnforcementMode](#istio.rbac.v1alpha1.EnforcementMode) |  | $hide_from_docs Indicates enforcement mode of the ServiceRoleBinding config. If RoleRef it refers to is in PERMISSIVE mode, the ServiceRoleBinding will be treated as PERMISSIVE mode regardless this mode field. If RoleRef it refers to is in ENFORCED mode, the ServiceRoleBinding will decide its own enforcement mode from this mode field. |






<a name="istio.rbac.v1alpha1.Subject"/>

### Subject
Subject defines an identity or a group of identities. The identity is either a user or
a group or identified by a set of &#34;properties&#34;. The name of the &#34;properties&#34; must match
the &#34;properties&#34; in the &#34;subject&#34; part of the &#34;authorization&#34; template
(https://github.com/istio/istio/blob/master/mixer/template/authorization/template.proto).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [string](#string) |  | Optional. The user name/ID that the subject represents. |
| group | [string](#string) |  | Optional. The group that the subject belongs to. |
| properties | [Subject.PropertiesEntry](#istio.rbac.v1alpha1.Subject.PropertiesEntry) | repeated | Optional. The set of properties that identify the subject. In the above ServiceRoleBinding example, the second subject has two properties: service: &#34;reviews&#34; namespace: &#34;abc&#34; |






<a name="istio.rbac.v1alpha1.Subject.PropertiesEntry"/>

### Subject.PropertiesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 


<a name="istio.rbac.v1alpha1.EnforcementMode"/>

### EnforcementMode
$hide_from_docs
RBAC config enforcement mode, used to verify new ServiceRole/ServiceBinding configs
work as expected before rolling to prod. RBAC engine only logs results from
ServiceRole/ServiceBinding configs that are in permissive mode, and discards result
before returning to the user.

| Name | Number | Description |
| ---- | ------ | ----------- |
| ENFORCED | 0 | Policy in ENFORCED mode has impact on user experience. Policy is in ENFORCED mode by default. |
| PERMISSIVE | 1 | Policy in PERMISSIVE doesn&#39;t impact on user experience. RBAC engine run policies in PERMISSIVE and logs decision. |



<a name="istio.rbac.v1alpha1.RbacConfig.Mode"/>

### RbacConfig.Mode


| Name | Number | Description |
| ---- | ------ | ----------- |
| OFF | 0 | Disable Istio RBAC completely, any other config in RbacConfig will be ignored and Istio RBAC policies will not be enforced. |
| ON | 1 | Enable Istio RBAC for all services and namespaces. |
| ON_WITH_INCLUSION | 2 | Enable Istio RBAC only for services and namespaces specified in the inclusion field. Any other services and namespaces not in the inclusion field will not be enforced by Istio RBAC policies. |
| ON_WITH_EXCLUSION | 3 | Enable Istio RBAC for all services and namespaces except those specified in the exclusion field. Any other services and namespaces not in the exclusion field will be enforced by Istio RBAC policies. |


 

 

 



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

