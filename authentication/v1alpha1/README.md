# Protocol Documentation
<a name="top"/>

## Table of Contents

- [authentication/v1alpha1/policy.proto](#authentication/v1alpha1/policy.proto)
    - [Jwt](#istio.authentication.v1alpha1.Jwt)
    - [MutualTls](#istio.authentication.v1alpha1.MutualTls)
    - [OriginAuthenticationMethod](#istio.authentication.v1alpha1.OriginAuthenticationMethod)
    - [PeerAuthenticationMethod](#istio.authentication.v1alpha1.PeerAuthenticationMethod)
    - [Policy](#istio.authentication.v1alpha1.Policy)
    - [PortSelector](#istio.authentication.v1alpha1.PortSelector)
    - [TargetSelector](#istio.authentication.v1alpha1.TargetSelector)
  
    - [MutualTls.Mode](#istio.authentication.v1alpha1.MutualTls.Mode)
    - [PrincipalBinding](#istio.authentication.v1alpha1.PrincipalBinding)
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="authentication/v1alpha1/policy.proto"/>
<p align="right"><a href="#top">Top</a></p>

## authentication/v1alpha1/policy.proto



<a name="istio.authentication.v1alpha1.Jwt"/>

### Jwt
JSON Web Token (JWT) token format for authentication as defined by
https://tools.ietf.org/html/rfc7519. See [OAuth
2.0](https://tools.ietf.org/html/rfc6749) and [OIDC
1.0](http://openid.net/connect) for how this is used in the whole
authentication flow.

Example,

```yaml
issuer: https://example.com
audiences:
- bookstore_android.apps.googleusercontent.com
  bookstore_web.apps.googleusercontent.com
jwksUri: https://example.com/.well-known/jwks.json
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| issuer | [string](#string) |  | Identifies the issuer that issued the JWT. See [issuer](https://tools.ietf.org/html/rfc7519#section-4.1.1) Usually a URL or an email address.  Example: https://securetoken.google.com Example: 1234567-compute@developer.gserviceaccount.com |
| audiences | [string](#string) | repeated | The list of JWT [audiences](https://tools.ietf.org/html/rfc7519#section-4.1.3). that are allowed to access. A JWT containing any of these audiences will be accepted.  The service name will be accepted if audiences is empty.  Example:  ```yaml audiences: - bookstore_android.apps.googleusercontent.com bookstore_web.apps.googleusercontent.com ``` |
| jwks_uri | [string](#string) |  | URL of the provider&#39;s public key set to validate signature of the JWT. See [OpenID Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  Optional if the key set document can either (a) be retrieved from [OpenID Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html) of the issuer or (b) inferred from the email domain of the issuer (e.g. a Google service account).  Example: https://www.googleapis.com/oauth2/v1/certs |
| jwt_headers | [string](#string) | repeated | JWT is sent in a request header. `header` represents the header name.  For example, if `header=x-goog-iap-jwt-assertion`, the header format will be x-goog-iap-jwt-assertion: &lt;JWT&gt;. |
| jwt_params | [string](#string) | repeated | JWT is sent in a query parameter. `query` represents the query parameter name.  For example, `query=jwt_token`. |






<a name="istio.authentication.v1alpha1.MutualTls"/>

### MutualTls
TLS authentication params.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| allow_tls | [bool](#bool) |  | WILL BE DEPRECATED, if set, will translates to `TLS_PERMISSIVE` mode. Set this flag to true to allow regular TLS (i.e without client x509 certificate). If request carries client certificate, identity will be extracted and used (set to peer identity). Otherwise, peer identity will be left unset. When the flag is false (default), request must have client certificate. |
| mode | [MutualTls.Mode](#istio.authentication.v1alpha1.MutualTls.Mode) |  | Defines the mode of mTLS authentication. |






<a name="istio.authentication.v1alpha1.OriginAuthenticationMethod"/>

### OriginAuthenticationMethod
OriginAuthenticationMethod defines authentication method/params for origin
authentication. Origin could be end-user, device, delegate service etc.
Currently, only JWT is supported for origin authentication.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| jwt | [Jwt](#istio.authentication.v1alpha1.Jwt) |  | Jwt params for the method. |






<a name="istio.authentication.v1alpha1.PeerAuthenticationMethod"/>

### PeerAuthenticationMethod
PeerAuthenticationMethod defines one particular type of authentication, e.g
mutual TLS, JWT etc, (no authentication is one type by itself) that can
be used for peer authentication.
The type can be progammatically determine by checking the type of the
&#34;params&#34; field.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mtls | [MutualTls](#istio.authentication.v1alpha1.MutualTls) |  | Set if mTLS is used. |
| jwt | [Jwt](#istio.authentication.v1alpha1.Jwt) |  | $hide_from_docs Set if JWT is used. This option is not yet available. |






<a name="istio.authentication.v1alpha1.Policy"/>

### Policy
Policy defines what authentication methods can be accepted on workload(s),
and if authenticated, which method/certificate will set the request principal
(i.e request.auth.principal attribute).

Authentication policy is composed of 2-part authentication:
- peer: verify caller service credentials. This part will set source.user
(peer identity).
- origin: verify the origin credentials. This part will set request.auth.user
(origin identity), as well as other attributes like request.auth.presenter,
request.auth.audiences and raw claims. Note that the identity could be
end-user, service account, device etc.

Last but not least, the principal binding rule defines which identity (peer
or origin) should be used as principal. By default, it uses peer.

Examples:

Policy to enable mTLS for all services in namespace frod

```yaml
apiVersion: authentication.istio.io/v1alpha1
kind: Policy
metadata:
  name: mTLS_enable
  namespace: frod
spec:
  peers:
  - mtls:
```
Policy to disable mTLS for &#34;productpage&#34; service

```yaml
apiVersion: authentication.istio.io/v1alpha1
kind: Policy
metadata:
  name: mTLS_disable
  namespace: frod
spec:
  targets:
  - name: productpage
```
Policy to require mTLS for peer authentication, and JWT for origin authenticationn
for productpage:9000. Principal is set from origin identity.

```yaml
apiVersion: authentication.istio.io/v1alpha1
kind: Policy
metadata:
  name: mTLS_enable
  namespace: frod
spec:
  target:
  - name: productpage
    ports:
    - number: 9000
  peers:
  - mtls:
  origins:
  - jwt:
      issuer: &#34;https://securetoken.google.com&#34;
      audiences:
      - &#34;productpage&#34;
      jwksUri: &#34;https://www.googleapis.com/oauth2/v1/certs&#34;
      jwt_headers:
      - &#34;x-goog-iap-jwt-assertion&#34;
  principaBinding: USE_ORIGIN
```

Policy to require mTLS for peer authentication, and JWT for origin authenticationn
for productpage:9000, but allow origin authentication failed. Principal is set
from origin identity.
Note: this example can be used for use cases when we want to allow request from
certain peers, given it comes with an approperiate authorization poicy to check
and reject request accoridingly.

```yaml
apiVersion: authentication.istio.io/v1alpha1
kind: Policy
metadata:
  name: mTLS_enable
  namespace: frod
spec:
  target:
  - name: productpage
    ports:
    - number: 9000
  peers:
  - mtls:
  origins:
  - jwt:
      issuer: &#34;https://securetoken.google.com&#34;
      audiences:
      - &#34;productpage&#34;
      jwksUri: &#34;https://www.googleapis.com/oauth2/v1/certs&#34;
      jwt_headers:
      - &#34;x-goog-iap-jwt-assertion&#34;
  originIsOptional: true
  principalBinding: USE_ORIGIN
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| targets | [TargetSelector](#istio.authentication.v1alpha1.TargetSelector) | repeated | List rules to select destinations that the policy should be applied on. If empty, policy will be used on all destinations in the same namespace. |
| peers | [PeerAuthenticationMethod](#istio.authentication.v1alpha1.PeerAuthenticationMethod) | repeated | List of authentication methods that can be used for peer authentication. They will be evaluated in order; the first validate one will be used to set peer identity (source.user) and other peer attributes. If none of these methods pass, and peer_is_optional flag is false (see below), request will be rejected with authentication failed error (401). Leave the list empty if peer authentication is not required |
| peer_is_optional | [bool](#bool) |  | Set this flag to true to accept request (for peer authentication perspective), even when none of the peer authentication methods defined above satisfied. Typically, this is used to delay the rejection decision to next layer (e.g authorization). This flag is ignored if no authentication defined for peer (peers field is empty). |
| origins | [OriginAuthenticationMethod](#istio.authentication.v1alpha1.OriginAuthenticationMethod) | repeated | List of authentication methods that can be used for origin authentication. Similar to peers, these will be evaluated in order; the first validate one will be used to set origin identity and attributes (i.e request.auth.user, request.auth.issuer etc). If none of these methods pass, and origin_is_optional is false (see below), request will be rejected with authentication failed error (401). Leave the list empty if origin authentication is not required. |
| origin_is_optional | [bool](#bool) |  | Set this flag to true to accept request (for origin authentication perspective), even when none of the origin authentication methods defined above satisfied. Typically, this is used to delay the rejection decision to next layer (e.g authorization). This flag is ignored if no authentication defined for origin (origins field is empty). |
| principal_binding | [PrincipalBinding](#istio.authentication.v1alpha1.PrincipalBinding) |  | Define whether peer or origin identity should be use for principal. Default value is USE_PEER. If peer (or orgin) identity is not available, either because of peer/origin authentication is not defined, or failed, principal will be left unset. In other words, binding rule does not affect the decision to accept or reject request. |






<a name="istio.authentication.v1alpha1.PortSelector"/>

### PortSelector
PortSelector specifies the name or number of a port to be used for
matching targets for authenticationn policy. This is copied from
networking API to avoid dependency.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| number | [uint32](#uint32) |  | Valid port number |
| name | [string](#string) |  | Port name |






<a name="istio.authentication.v1alpha1.TargetSelector"/>

### TargetSelector
TargetSelector defines a matching rule to a service/destination.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | REQUIRED. The name must be a short name from the service registry. The fully qualified domain name will be resolved in a platform specific manner. |
| ports | [PortSelector](#istio.authentication.v1alpha1.PortSelector) | repeated | Specifies the ports on the destination. Leave empty to match all ports that are exposed. |





 


<a name="istio.authentication.v1alpha1.MutualTls.Mode"/>

### MutualTls.Mode
Defines the acceptable connection TLS mode.

| Name | Number | Description |
| ---- | ------ | ----------- |
| STRICT | 0 | Client cert must be presented, connection is in TLS. |
| PERMISSIVE | 1 | Connection can be either plaintext or TLS, and client cert can be omitted. |



<a name="istio.authentication.v1alpha1.PrincipalBinding"/>

### PrincipalBinding
Associates authentication with request principal.

| Name | Number | Description |
| ---- | ------ | ----------- |
| USE_PEER | 0 | Principal will be set to the identity from peer authentication. |
| USE_ORIGIN | 1 | Principal will be set to the identity from origin authentication. |


 

 

 



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

