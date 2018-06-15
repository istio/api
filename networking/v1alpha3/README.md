# Protocol Documentation
<a name="top"/>

## Table of Contents

- [networking/v1alpha3/destination_rule.proto](#networking/v1alpha3/destination_rule.proto)
    - [ConnectionPoolSettings](#istio.networking.v1alpha3.ConnectionPoolSettings)
    - [ConnectionPoolSettings.HTTPSettings](#istio.networking.v1alpha3.ConnectionPoolSettings.HTTPSettings)
    - [ConnectionPoolSettings.TCPSettings](#istio.networking.v1alpha3.ConnectionPoolSettings.TCPSettings)
    - [DestinationRule](#istio.networking.v1alpha3.DestinationRule)
    - [LoadBalancerSettings](#istio.networking.v1alpha3.LoadBalancerSettings)
    - [LoadBalancerSettings.ConsistentHashLB](#istio.networking.v1alpha3.LoadBalancerSettings.ConsistentHashLB)
    - [OutlierDetection](#istio.networking.v1alpha3.OutlierDetection)
    - [OutlierDetection.HTTPSettings](#istio.networking.v1alpha3.OutlierDetection.HTTPSettings)
    - [Subset](#istio.networking.v1alpha3.Subset)
    - [Subset.LabelsEntry](#istio.networking.v1alpha3.Subset.LabelsEntry)
    - [TLSSettings](#istio.networking.v1alpha3.TLSSettings)
    - [TrafficPolicy](#istio.networking.v1alpha3.TrafficPolicy)
    - [TrafficPolicy.PortTrafficPolicy](#istio.networking.v1alpha3.TrafficPolicy.PortTrafficPolicy)
  
    - [LoadBalancerSettings.SimpleLB](#istio.networking.v1alpha3.LoadBalancerSettings.SimpleLB)
    - [TLSSettings.TLSmode](#istio.networking.v1alpha3.TLSSettings.TLSmode)
  
  
  

- [networking/v1alpha3/envoy_filter.proto](#networking/v1alpha3/envoy_filter.proto)
    - [EnvoyFilters](#istio.networking.v1alpha3.EnvoyFilters)
    - [EnvoyFilters.Filter](#istio.networking.v1alpha3.EnvoyFilters.Filter)
    - [EnvoyFilters.InsertPosition](#istio.networking.v1alpha3.EnvoyFilters.InsertPosition)
    - [EnvoyFilters.ListenerMatch](#istio.networking.v1alpha3.EnvoyFilters.ListenerMatch)
    - [EnvoyFilters.SelectorEntry](#istio.networking.v1alpha3.EnvoyFilters.SelectorEntry)
  
  
  
  

- [networking/v1alpha3/gateway.proto](#networking/v1alpha3/gateway.proto)
    - [Gateway](#istio.networking.v1alpha3.Gateway)
    - [Gateway.SelectorEntry](#istio.networking.v1alpha3.Gateway.SelectorEntry)
    - [Port](#istio.networking.v1alpha3.Port)
    - [Server](#istio.networking.v1alpha3.Server)
    - [Server.TLSOptions](#istio.networking.v1alpha3.Server.TLSOptions)
  
  
  
  

- [networking/v1alpha3/service_entry.proto](#networking/v1alpha3/service_entry.proto)
    - [ServiceEntry](#istio.networking.v1alpha3.ServiceEntry)
    - [ServiceEntry.Endpoint](#istio.networking.v1alpha3.ServiceEntry.Endpoint)
  
    - [ServiceEntry.Location](#istio.networking.v1alpha3.ServiceEntry.Location)
    - [ServiceEntry.Resolution](#istio.networking.v1alpha3.ServiceEntry.Resolution)
  
  
  

- [networking/v1alpha3/virtual_service.proto](#networking/v1alpha3/virtual_service.proto)
    - [CorsPolicy](#istio.networking.v1alpha3.CorsPolicy)
    - [Destination](#istio.networking.v1alpha3.Destination)
    - [DestinationWeight](#istio.networking.v1alpha3.DestinationWeight)
    - [HTTPFaultInjection](#istio.networking.v1alpha3.HTTPFaultInjection)
    - [HTTPFaultInjection.Abort](#istio.networking.v1alpha3.HTTPFaultInjection.Abort)
    - [HTTPFaultInjection.Delay](#istio.networking.v1alpha3.HTTPFaultInjection.Delay)
    - [HTTPMatchRequest](#istio.networking.v1alpha3.HTTPMatchRequest)
    - [HTTPMatchRequest.HeadersEntry](#istio.networking.v1alpha3.HTTPMatchRequest.HeadersEntry)
    - [HTTPMatchRequest.SourceLabelsEntry](#istio.networking.v1alpha3.HTTPMatchRequest.SourceLabelsEntry)
    - [HTTPRedirect](#istio.networking.v1alpha3.HTTPRedirect)
    - [HTTPRetry](#istio.networking.v1alpha3.HTTPRetry)
    - [HTTPRewrite](#istio.networking.v1alpha3.HTTPRewrite)
    - [HTTPRoute](#istio.networking.v1alpha3.HTTPRoute)
    - [HTTPRoute.AppendHeadersEntry](#istio.networking.v1alpha3.HTTPRoute.AppendHeadersEntry)
    - [HTTPRoute.RemoveResponseHeadersEntry](#istio.networking.v1alpha3.HTTPRoute.RemoveResponseHeadersEntry)
    - [L4MatchAttributes](#istio.networking.v1alpha3.L4MatchAttributes)
    - [L4MatchAttributes.SourceLabelsEntry](#istio.networking.v1alpha3.L4MatchAttributes.SourceLabelsEntry)
    - [PortSelector](#istio.networking.v1alpha3.PortSelector)
    - [StringMatch](#istio.networking.v1alpha3.StringMatch)
    - [TCPRoute](#istio.networking.v1alpha3.TCPRoute)
    - [VirtualService](#istio.networking.v1alpha3.VirtualService)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="networking/v1alpha3/destination_rule.proto"/>
<p align="right"><a href="#top">Top</a></p>

## networking/v1alpha3/destination_rule.proto
Copyright 2018 Istio Authors

  Licensed under the Apache License, Version 2.0 (the &#34;License&#34;);
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an &#34;AS IS&#34; BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.


<a name="istio.networking.v1alpha3.ConnectionPoolSettings"/>

### ConnectionPoolSettings
Connection pool settings for an upstream host. The settings apply to
each individual host in the upstream service.  See Envoy&#39;s [circuit
breaker](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/circuit_breaking)
for more details. Connection pool settings can be applied at the TCP
level as well as at HTTP level.

For example, the following rule sets a limit of 100 connections to redis
service called myredissrv with a connect timeout of 30ms

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: bookinfo-redis
spec:
  host: myredissrv.prod.svc.cluster.local
  trafficPolicy:
    connectionPool:
      tcp:
        maxConnections: 100
        connectTimeout: 30ms
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tcp | [ConnectionPoolSettings.TCPSettings](#istio.networking.v1alpha3.ConnectionPoolSettings.TCPSettings) |  | Settings common to both HTTP and TCP upstream connections. |
| http | [ConnectionPoolSettings.HTTPSettings](#istio.networking.v1alpha3.ConnectionPoolSettings.HTTPSettings) |  | HTTP connection pool settings. |






<a name="istio.networking.v1alpha3.ConnectionPoolSettings.HTTPSettings"/>

### ConnectionPoolSettings.HTTPSettings
Settings applicable to HTTP1.1/HTTP2/GRPC connections.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| http1_max_pending_requests | [int32](#int32) |  | Maximum number of pending HTTP requests to a destination. Default 1024. |
| http2_max_requests | [int32](#int32) |  | Maximum number of requests to a backend. Default 1024. |
| max_requests_per_connection | [int32](#int32) |  | Maximum number of requests per connection to a backend. Setting this parameter to 1 disables keep alive. |
| max_retries | [int32](#int32) |  | Maximum number of retries that can be outstanding to all hosts in a cluster at a given time. Defaults to 3. |






<a name="istio.networking.v1alpha3.ConnectionPoolSettings.TCPSettings"/>

### ConnectionPoolSettings.TCPSettings
Settings common to both HTTP and TCP upstream connections.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| max_connections | [int32](#int32) |  | Maximum number of HTTP1 /TCP connections to a destination host. |
| connect_timeout | [google.protobuf.Duration](#google.protobuf.Duration) |  | TCP connection timeout. |






<a name="istio.networking.v1alpha3.DestinationRule"/>

### DestinationRule
`DestinationRule` defines policies that apply to traffic intended for a
service after routing has occurred. These rules specify configuration
for load balancing, connection pool size from the sidecar, and outlier
detection settings to detect and evict unhealthy hosts from the load
balancing pool. For example, a simple load balancing policy for the
ratings service would look as follows:

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: bookinfo-ratings
spec:
  host: ratings.prod.svc.cluster.local
  trafficPolicy:
    loadBalancer:
      simple: LEAST_CONN
```

Version specific policies can be specified by defining a named
`subset` and overriding the settings specified at the service level. The
following rule uses a round robin load balancing policy for all traffic
going to a subset named testversion that is composed of endpoints (e.g.,
pods) with labels (version:v3).

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: bookinfo-ratings
spec:
  host: ratings.prod.svc.cluster.local
  trafficPolicy:
    loadBalancer:
      simple: LEAST_CONN
  subsets:
  - name: testversion
    labels:
      version: v3
    trafficPolicy:
      loadBalancer:
        simple: ROUND_ROBIN
```

**Note:** Policies specified for subsets will not take effect until
a route rule explicitly sends traffic to this subset.

Traffic policies can be customized to specific ports as well. The
following rule uses the least connection load balancing policy for all
traffic to port 80, while uses a round robin load balancing setting for
traffic to the port 9080.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: bookinfo-ratings-port
spec:
  host: ratings.prod.svc.cluster.local
  trafficPolicy: # Apply to all ports
    portLevelSettings:
    - port:
        number: 80
      loadBalancer:
        simple: LEAST_CONN
    - port:
        number: 9080
      loadBalancer:
        simple: ROUND_ROBIN
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| host | [string](#string) |  | REQUIRED. The name of a service from the service registry. Service names are looked up from the platform&#39;s service registry (e.g., Kubernetes services, Consul services, etc.) and from the hosts declared by [ServiceEntries](#ServiceEntry). Rules defined for services that do not exist in the service registry will be ignored.  *Note for Kubernetes users*: When short names are used (e.g. &#34;reviews&#34; instead of &#34;reviews.default.svc.cluster.local&#34;), Istio will interpret the short name based on the namespace of the rule, not the service. A rule in the &#34;default&#34; namespace containing a host &#34;reviews will be interpreted as &#34;reviews.default.svc.cluster.local&#34;, irrespective of the actual namespace associated with the reviews service. _To avoid potential misconfigurations, it is recommended to always use fully qualified domain names over short names._  Note that the host field applies to both HTTP and TCP services. |
| traffic_policy | [TrafficPolicy](#istio.networking.v1alpha3.TrafficPolicy) |  | Traffic policies to apply (load balancing policy, connection pool sizes, outlier detection). |
| subsets | [Subset](#istio.networking.v1alpha3.Subset) | repeated | One or more named sets that represent individual versions of a service. Traffic policies can be overridden at subset level. |






<a name="istio.networking.v1alpha3.LoadBalancerSettings"/>

### LoadBalancerSettings
Load balancing policies to apply for a specific destination. See Envoy&#39;s
load balancing
[documentation](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/load_balancing.html)
for more details.

For example, the following rule uses a round robin load balancing policy
for all traffic going to the ratings service.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: bookinfo-ratings
spec:
  host: ratings.prod.svc.cluster.local
  trafficPolicy:
    loadBalancer:
      simple: ROUND_ROBIN
```

The following example uses the consistent hashing based load balancer
for the same ratings service using the Cookie header as the hash key.

```yaml
 apiVersion: networking.istio.io/v1alpha3
 kind: DestinationRule
 metadata:
   name: bookinfo-ratings
 spec:
   host: ratings.prod.svc.cluster.local
   trafficPolicy:
     loadBalancer:
       consistentHash:
         http_header: Cookie
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| simple | [LoadBalancerSettings.SimpleLB](#istio.networking.v1alpha3.LoadBalancerSettings.SimpleLB) |  |  |
| consistent_hash | [LoadBalancerSettings.ConsistentHashLB](#istio.networking.v1alpha3.LoadBalancerSettings.ConsistentHashLB) |  |  |






<a name="istio.networking.v1alpha3.LoadBalancerSettings.ConsistentHashLB"/>

### LoadBalancerSettings.ConsistentHashLB
Consistent hashing (ketama hash) based load balancer for even load
distribution/redistribution when the connection pool changes. This
load balancing policy is applicable only for HTTP-based
connections. A user specified HTTP header is used as the key with
[xxHash](http://cyan4973.github.io/xxHash) hashing.
$hide_from_docs


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| http_header | [string](#string) |  | REQUIRED. The name of the HTTP request header that will be used to obtain the hash key. If the request header is not present, the load balancer will use a random number as the hash, effectively making the load balancing policy random. |
| minimum_ring_size | [uint32](#uint32) |  | The minimum number of virtual nodes to use for the hash ring. Defaults to 1024. Larger ring sizes result in more granular load distributions. If the number of hosts in the load balancing pool is larger than the ring size, each host will be assigned a single virtual node. |






<a name="istio.networking.v1alpha3.OutlierDetection"/>

### OutlierDetection
A Circuit breaker implementation that tracks the status of each
individual host in the upstream service.  Applicable to both HTTP and
TCP services.  For HTTP services, hosts that continually return 5xx
errors for API calls are ejected from the pool for a pre-defined period
of time. For TCP services, connection timeouts or connection
failures to a given host counts as an error when measuring the
consecutive errors metric. See Envoy&#39;s [outlier
detection](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/outlier)
for more details. 

The following rule sets a connection pool size of 100 connections and
1000 concurrent HTTP2 requests, with no more than 10 req/connection to
&#34;reviews&#34; service. In addition, it configures upstream hosts to be
scanned every 5 mins, such that any host that fails 7 consecutive times
with 5XX error code will be ejected for 15 minutes.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: reviews-cb-policy
spec:
  host: reviews.prod.svc.cluster.local
  trafficPolicy:
    connectionPool:
      tcp:
        maxConnections: 100
      http:
        http2MaxRequests: 1000
        maxRequestsPerConnection: 10
    outlierDetection:
      consecutiveErrors: 7
      interval: 5m
      baseEjectionTime: 15m
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| consecutive_errors | [int32](#int32) |  | Number of errors before a host is ejected from the connection pool. Defaults to 5. When the upstream host is accessed over HTTP, a 5xx return code qualifies as an error. When the upstream host is accessed over an opaque TCP connection, connect timeouts and connection error/failure events qualify as an error. |
| interval | [google.protobuf.Duration](#google.protobuf.Duration) |  | Time interval between ejection sweep analysis. format: 1h/1m/1s/1ms. MUST BE &gt;=1ms. Default is 10s. |
| base_ejection_time | [google.protobuf.Duration](#google.protobuf.Duration) |  | Minimum ejection duration. A host will remain ejected for a period equal to the product of minimum ejection duration and the number of times the host has been ejected. This technique allows the system to automatically increase the ejection period for unhealthy upstream servers. format: 1h/1m/1s/1ms. MUST BE &gt;=1ms. Default is 30s. |
| max_ejection_percent | [int32](#int32) |  | Maximum % of hosts in the load balancing pool for the upstream service that can be ejected. Defaults to 10%. |
| http | [OutlierDetection.HTTPSettings](#istio.networking.v1alpha3.OutlierDetection.HTTPSettings) |  | DEPRECATED. |






<a name="istio.networking.v1alpha3.OutlierDetection.HTTPSettings"/>

### OutlierDetection.HTTPSettings
Outlier detection settings for HTTP1.1/HTTP2/GRPC connections.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| consecutive_errors | [int32](#int32) |  | Number of 5XX errors before a host is ejected from the connection pool. Defaults to 5. |
| interval | [google.protobuf.Duration](#google.protobuf.Duration) |  | Time interval between ejection sweep analysis. format: 1h/1m/1s/1ms. MUST BE &gt;=1ms. Default is 10s. |
| base_ejection_time | [google.protobuf.Duration](#google.protobuf.Duration) |  | Minimum ejection duration. A host will remain ejected for a period equal to the product of minimum ejection duration and the number of times the host has been ejected. This technique allows the system to automatically increase the ejection period for unhealthy upstream servers. format: 1h/1m/1s/1ms. MUST BE &gt;=1ms. Default is 30s. |
| max_ejection_percent | [int32](#int32) |  | Maximum % of hosts in the load balancing pool for the upstream service that can be ejected. Defaults to 10%. |






<a name="istio.networking.v1alpha3.Subset"/>

### Subset
A subset of endpoints of a service. Subsets can be used for scenarios
like A/B testing, or routing to a specific version of a service. Refer
to [VirtualService](#VirtualService) documentation for examples of using
subsets in these scenarios. In addition, traffic policies defined at the
service-level can be overridden at a subset-level. The following rule
uses a round robin load balancing policy for all traffic going to a
subset named testversion that is composed of endpoints (e.g., pods) with
labels (version:v3).

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: bookinfo-ratings
spec:
  host: ratings.prod.svc.cluster.local
  trafficPolicy:
    loadBalancer:
      simple: LEAST_CONN
  subsets:
  - name: testversion
    labels:
      version: v3
    trafficPolicy:
      loadBalancer:
        simple: ROUND_ROBIN
```

**Note:** Policies specified for subsets will not take effect until
a route rule explicitly sends traffic to this subset.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | REQUIRED. Name of the subset. The service name and the subset name can be used for traffic splitting in a route rule. |
| labels | [Subset.LabelsEntry](#istio.networking.v1alpha3.Subset.LabelsEntry) | repeated | REQUIRED. Labels apply a filter over the endpoints of a service in the service registry. See route rules for examples of usage. |
| traffic_policy | [TrafficPolicy](#istio.networking.v1alpha3.TrafficPolicy) |  | Traffic policies that apply to this subset. Subsets inherit the traffic policies specified at the DestinationRule level. Settings specified at the subset level will override the corresponding settings specified at the DestinationRule level. |






<a name="istio.networking.v1alpha3.Subset.LabelsEntry"/>

### Subset.LabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="istio.networking.v1alpha3.TLSSettings"/>

### TLSSettings
SSL/TLS related settings for upstream connections. See Envoy&#39;s [TLS
context](https://www.envoyproxy.io/docs/envoy/latest/api-v1/cluster_manager/cluster_ssl.html#config-cluster-manager-cluster-ssl)
for more details. These settings are common to both HTTP and TCP upstreams.

For example, the following rule configures a client to use mutual TLS
for connections to upstream database cluster.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: db-mtls
spec:
  host: mydbserver.prod.svc.cluster.local
  trafficPolicy:
    tls:
      mode: MUTUAL
      clientCertificate: /etc/certs/myclientcert.pem
      privateKey: /etc/certs/client_private_key.pem
      caCertificates: /etc/certs/rootcacerts.pem
```

The following rule configures a client to use TLS when talking to a
foreign service whose domain matches *.foo.com.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: tls-foo
spec:
  host: &#34;*.foo.com&#34;
  trafficPolicy:
    tls:
      mode: SIMPLE
```

The following rule configures a client to use Istio mutual TLS when talking
to rating services.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: ratings-istio-mtls
spec:
  host: ratings.prod.svc.cluster.local
  trafficPolicy:
    tls:
      mode: ISTIO_MUTUAL
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mode | [TLSSettings.TLSmode](#istio.networking.v1alpha3.TLSSettings.TLSmode) |  | REQUIRED: Indicates whether connections to this port should be secured using TLS. The value of this field determines how TLS is enforced. |
| client_certificate | [string](#string) |  | REQUIRED if mode is `MUTUAL`. The path to the file holding the client-side TLS certificate to use. Should be empty if mode is `ISTIO_MUTUAL`. |
| private_key | [string](#string) |  | REQUIRED if mode is `MUTUAL`. The path to the file holding the client&#39;s private key. Should be empty if mode is `ISTIO_MUTUAL`. |
| ca_certificates | [string](#string) |  | OPTIONAL: The path to the file containing certificate authority certificates to use in verifying a presented server certificate. If omitted, the proxy will not verify the server&#39;s certificate. Should be empty if mode is `ISTIO_MUTUAL`. |
| subject_alt_names | [string](#string) | repeated | A list of alternate names to verify the subject identity in the certificate. If specified, the proxy will verify that the server certificate&#39;s subject alt name matches one of the specified values. Should be empty if mode is `ISTIO_MUTUAL`. |
| sni | [string](#string) |  | SNI string to present to the server during TLS handshake. Should be empty if mode is `ISTIO_MUTUAL`. |






<a name="istio.networking.v1alpha3.TrafficPolicy"/>

### TrafficPolicy
Traffic policies to apply for a specific destination, across all
destination ports. See DestinationRule for examples.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| load_balancer | [LoadBalancerSettings](#istio.networking.v1alpha3.LoadBalancerSettings) |  | Settings controlling the load balancer algorithms. |
| connection_pool | [ConnectionPoolSettings](#istio.networking.v1alpha3.ConnectionPoolSettings) |  | Settings controlling the volume of connections to an upstream service |
| outlier_detection | [OutlierDetection](#istio.networking.v1alpha3.OutlierDetection) |  | Settings controlling eviction of unhealthy hosts from the load balancing pool |
| tls | [TLSSettings](#istio.networking.v1alpha3.TLSSettings) |  | TLS related settings for connections to the upstream service. |
| port_level_settings | [TrafficPolicy.PortTrafficPolicy](#istio.networking.v1alpha3.TrafficPolicy.PortTrafficPolicy) | repeated | Traffic policies specific to individual ports. Note that port level settings will override the destination-level settings. Traffic settings specified at the destination-level will not be inherited when overridden by port-level settings, i.e. default values will be applied to fields omitted in port-level traffic policies. |






<a name="istio.networking.v1alpha3.TrafficPolicy.PortTrafficPolicy"/>

### TrafficPolicy.PortTrafficPolicy
Traffic policies that apply to specific ports of the service


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port | [PortSelector](#istio.networking.v1alpha3.PortSelector) |  | Specifies the port name or number of a port on the destination service on which this policy is being applied.  Names must comply with DNS label syntax (rfc1035) and therefore cannot collide with numbers. If there are multiple ports on a service with the same protocol the names should be of the form &lt;protocol-name&gt;-&lt;DNS label&gt;. |
| load_balancer | [LoadBalancerSettings](#istio.networking.v1alpha3.LoadBalancerSettings) |  | Settings controlling the load balancer algorithms. |
| connection_pool | [ConnectionPoolSettings](#istio.networking.v1alpha3.ConnectionPoolSettings) |  | Settings controlling the volume of connections to an upstream service |
| outlier_detection | [OutlierDetection](#istio.networking.v1alpha3.OutlierDetection) |  | Settings controlling eviction of unhealthy hosts from the load balancing pool |
| tls | [TLSSettings](#istio.networking.v1alpha3.TLSSettings) |  | TLS related settings for connections to the upstream service. |





 


<a name="istio.networking.v1alpha3.LoadBalancerSettings.SimpleLB"/>

### LoadBalancerSettings.SimpleLB
Standard load balancing algorithms that require no tuning.

| Name | Number | Description |
| ---- | ------ | ----------- |
| ROUND_ROBIN | 0 | Round Robin policy. Default |
| LEAST_CONN | 1 | The least request load balancer uses an O(1) algorithm which selects two random healthy hosts and picks the host which has fewer active requests. |
| RANDOM | 2 | The random load balancer selects a random healthy host. The random load balancer generally performs better than round robin if no health checking policy is configured. |
| PASSTHROUGH | 3 | This option will forward the connection to the original IP address requested by the caller without doing any form of load balancing. This option must be used with care. It is meant for advanced use cases. Refer to Original Destination load balancer in Envoy for further details. |



<a name="istio.networking.v1alpha3.TLSSettings.TLSmode"/>

### TLSSettings.TLSmode
TLS connection mode

| Name | Number | Description |
| ---- | ------ | ----------- |
| DISABLE | 0 | Do not setup a TLS connection to the upstream endpoint. |
| SIMPLE | 1 | Originate a TLS connection to the upstream endpoint. |
| MUTUAL | 2 | Secure connections to the upstream using mutual TLS by presenting client certificates for authentication. |
| ISTIO_MUTUAL | 3 | Secure connections to the upstream using mutual TLS by presenting client certificates for authentication. Compared to Mutual mode, this mode uses certificates generated automatically by Istio for mTLS authentication. When this mode is used, all other fields in `TLSSettings` should be empty. |


 

 

 



<a name="networking/v1alpha3/envoy_filter.proto"/>
<p align="right"><a href="#top">Top</a></p>

## networking/v1alpha3/envoy_filter.proto



<a name="istio.networking.v1alpha3.EnvoyFilters"/>

### EnvoyFilters
`EnvoyFilter` describes Envoy proxy-specific filters that can be used to
customize the Envoy proxy configuration generated by Istio networking
subsystem (Pilot). This feature must be used with care, as incorrect
configurations could potentially destabilize the entire mesh.

NOTE: Since this is break glass configuration, there will not be any
backward compatibility across different Istio releases. In other words,
this configuration is subject to change based on internal implementation
of Istio networking subsystem.

The following example for Kubernetes enables Envoy&#39;s Lua filter for all
inbound calls arriving at port 18080 of the reviews service pod with
labels &#34;app: reviews&#34;.

    apiVersion: networking.istio.io/v1alpha3
    kind: EnvoyFilter
    metadata:
      name: reviews-lua
    spec:
      selector:
        app: reviews
      filters:
      - listenerMatch:
          port: 18080
          listenerType: SIDECAR_INBOUND #will match with the listener for the podIP:18080
        filterName: envoy.lua
        filterType: HTTP
        filterConfig:
          inlineCode: |
            ... lua code ...


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selector | [EnvoyFilters.SelectorEntry](#istio.networking.v1alpha3.EnvoyFilters.SelectorEntry) | repeated | One or more labels that indicate a specific set of pods/VMs whose proxies should be configured to use these additional filters. The scope of label search is platform dependent. On Kubernetes, for example, the scope includes pods running in all reachable namespaces. Omitting the selector applies the filter to all proxies in the mesh. |
| filters | [EnvoyFilters.Filter](#istio.networking.v1alpha3.EnvoyFilters.Filter) | repeated | REQUIRED: Envoy network filters/http filters to be added to matching listeners. When adding network filters to http connections, care should be taken to ensure that the filter is added before envoy.http_connection_manager. |






<a name="istio.networking.v1alpha3.EnvoyFilters.Filter"/>

### EnvoyFilters.Filter
Envoy filters to be added to a network or http filter chain.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| listener_match | [EnvoyFilters.ListenerMatch](#istio.networking.v1alpha3.EnvoyFilters.ListenerMatch) |  | Filter will be added to the listner only if the match conditions are true. If not specified, the filters will be applied to all listeners. |
| insert_position | [EnvoyFilters.InsertPosition](#istio.networking.v1alpha3.EnvoyFilters.InsertPosition) |  | Insert position in the filter chain. Defaults to FIRST |
| filter_type | [EnvoyFilters.Filter.FilterType](#istio.networking.v1alpha3.EnvoyFilters.Filter.FilterType) |  | REQUIRED: The type of filter to instantiate. |
| filter_name | [string](#string) |  | REQUIRED: The name of the filter to instantiate. The name must match a supported filter _compiled into_ Envoy. |
| filter_config | [google.protobuf.Struct](#google.protobuf.Struct) |  | REQUIRED: Filter specific configuration which depends on the filter being instantiated. |






<a name="istio.networking.v1alpha3.EnvoyFilters.InsertPosition"/>

### EnvoyFilters.InsertPosition
Indicates the relative index in the filter chain where the filter should be inserted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [EnvoyFilters.InsertPosition.Index](#istio.networking.v1alpha3.EnvoyFilters.InsertPosition.Index) |  | Position of this filter in the filter chain. |
| relative_to | [string](#string) |  | If BEFORE or AFTER position is specified, specify the name of the filter relative to which this filter should be inserted. |






<a name="istio.networking.v1alpha3.EnvoyFilters.ListenerMatch"/>

### EnvoyFilters.ListenerMatch
Select a listener to add the filter to based on the match conditions.
All conditions specified in the ListenerMatch must be met for the filter
to be applied to a listener.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port | [uint32](#uint32) |  | Port associated with the listener. If not specified, matches all listeners. |
| listener_type | [EnvoyFilters.ListenerMatch.ListenerType](#istio.networking.v1alpha3.EnvoyFilters.ListenerMatch.ListenerType) |  | Inbound vs outbound sidecar listener or gateway listener. If not specified, matches all listeners. |
| listener_protocol | [string](#string) |  | Selects a class of listeners for the same protocol. If not specified, applies to listeners on all protocols. Use the protocol selection to select all HTTP listeners (includes HTTP2/gRPC/HTTPS where Envoy terminates TLS) or all TCP listeners (includes HTTPS passthrough using SNI). Acceptable values are TCP/HTTP/MONGO. |
| address | [string](#string) | repeated | One or more IP addresses to which the listener is bound. If specified, should match atleast one address in the list. |






<a name="istio.networking.v1alpha3.EnvoyFilters.SelectorEntry"/>

### EnvoyFilters.SelectorEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 

 

 

 



<a name="networking/v1alpha3/gateway.proto"/>
<p align="right"><a href="#top">Top</a></p>

## networking/v1alpha3/gateway.proto



<a name="istio.networking.v1alpha3.Gateway"/>

### Gateway
`Gateway` describes a load balancer operating at the edge of the mesh
receiving incoming or outgoing HTTP/TCP connections. The specification
describes a set of ports that should be exposed, the type of protocol to
use, SNI configuration for the load balancer, etc.

For example, the following Gateway configuration sets up a proxy to act
as a load balancer exposing port 80 and 9080 (http), 443 (https), and
port 2379 (TCP) for ingress.  The gateway will be applied to the proxy
running on a pod with labels `app: my-gateway-controller`. While Istio
will configure the proxy to listen on these ports, it is the
responsibility of the user to ensure that external traffic to these
ports are allowed into the mesh.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: my-gateway
spec:
  selector:
    app: my-gatweway-controller
  servers:
  - port:
      number: 80
      name: http
      protocol: HTTP
    hosts:
    - uk.bookinfo.com
    - eu.bookinfo.com
    tls:
      httpsRedirect: true # sends 302 redirect for http requests
  - port:
      number: 443
      name: https
      protocol: HTTPS
    hosts:
    - uk.bookinfo.com
    - eu.bookinfo.com
    tls:
      mode: SIMPLE #enables HTTPS on this port
      serverCertificate: /etc/certs/servercert.pem
      privateKey: /etc/certs/privatekey.pem
  - port:
      number: 9080
      name: http-wildcard
      protocol: HTTP
    hosts:
    - &#34;*&#34;
  - port:
      number: 2379 # to expose internal service via external port 2379
      name: mongo
      protocol: MONGO
    hosts:
    - &#34;*&#34;
```
The Gateway specification above describes the L4-L6 properties of a load
balancer. A `VirtualService` can then be bound to a gateway to control
the forwarding of traffic arriving at a particular host or gateway port.

For example, the following VirtualService splits traffic for
&#34;https://uk.bookinfo.com/reviews&#34;, &#34;https://eu.bookinfo.com/reviews&#34;,
&#34;http://uk.bookinfo.com:9080/reviews&#34;,
&#34;http://eu.bookinfo.com:9080/reviews&#34; into two versions (prod and qa) of
an internal reviews service on port 9080. In addition, requests
containing the cookie &#34;user: dev-123&#34; will be sent to special port 7777
in the qa version. The same rule is also applicable inside the mesh for
requests to the &#34;reviews.prod.svc.cluster.local&#34; service. This rule is
applicable across ports 443, 9080. Note that &#34;http://uk.bookinfo.com&#34;
gets redirected to &#34;https://uk.bookinfo.com&#34; (i.e. 80 redirects to 443).

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: bookinfo-rule
spec:
  hosts:
  - reviews.prod.svc.cluster.local
  - uk.bookinfo.com
  - eu.bookinfo.com
  gateways:
  - my-gateway
  - mesh # applies to all the sidecars in the mesh
  http:
  - match:
    - headers:
        cookie:
          user: dev-123
    route:
    - destination:
        port:
          number: 7777
        host: reviews.qa.svc.cluster.local
  - match:
      uri:
        prefix: /reviews/
    route:
    - destination:
        port:
          number: 9080 # can be omitted if its the only port for reviews
        host: reviews.prod.svc.cluster.local
      weight: 80
    - destination:
        host: reviews.qa.svc.cluster.local
      weight: 20
```

The following VirtualService forwards traffic arriving at (external)
port 27017 from &#34;172.17.16.0/24&#34; subnet to internal Mongo server on port
5555. This rule is not applicable internally in the mesh as the gateway
list omits the reserved name `mesh`.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: bookinfo-Mongo
spec:
  hosts:
  - mongosvr.prod.svc.cluster.local #name of internal Mongo service
  gateways:
  - my-gateway
  tcp:
  - match:
    - port:
        number: 27017
      sourceSubnet: &#34;172.17.16.0/24&#34;
    route:
    - destination:
        host: mongo.prod.svc.cluster.local
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| servers | [Server](#istio.networking.v1alpha3.Server) | repeated | REQUIRED: A list of server specifications. |
| selector | [Gateway.SelectorEntry](#istio.networking.v1alpha3.Gateway.SelectorEntry) | repeated | REQUIRED: One or more labels that indicate a specific set of pods/VMs on which this gateway configuration should be applied. The scope of label search is platform dependent. On Kubernetes, for example, the scope includes pods running in all reachable namespaces. |






<a name="istio.networking.v1alpha3.Gateway.SelectorEntry"/>

### Gateway.SelectorEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="istio.networking.v1alpha3.Port"/>

### Port
Port describes the properties of a specific port of a service.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| number | [uint32](#uint32) |  | REQUIRED: A valid non-negative integer port number. |
| protocol | [string](#string) |  | REQUIRED: The protocol exposed on the port. MUST BE one of HTTP|HTTPS|GRPC|HTTP2|MONGO|TCP|TCP-TLS. TCP-TLS is used to indicate secure connections to non HTTP services. |
| name | [string](#string) |  | Label assigned to the port. |






<a name="istio.networking.v1alpha3.Server"/>

### Server
`Server` describes the properties of the proxy on a given load balancer
port. For example,

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: my-ingress
spec:
  selector:
    app: my-ingress-gateway
  servers:
  - port:
      number: 80
      name: http2
      protocol: HTTP2
    hosts:
    - &#34;*&#34;
```

Another example

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: my-tcp-ingress
spec:
  selector:
    app: my-tcp-ingress-gateway
  servers:
  - port:
      number: 27018
      name: mongo
      protocol: MONGO
    hosts:
    - &#34;*&#34;
```

The following is an example of TLS configuration for port 443

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: my-tls-ingress
spec:
  selector:
    app: my-tls-ingress-gateway
  servers:
  - port:
      number: 443
      name: https
      protocol: HTTPS
    hosts:
    - &#34;*&#34;
    tls:
      mode: SIMPLE
      serverCertificate: /etc/certs/server.pem
      privateKey: /etc/certs/privatekey.pem
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port | [Port](#istio.networking.v1alpha3.Port) |  | REQUIRED: The Port on which the proxy should listen for incoming connections |
| hosts | [string](#string) | repeated | REQUIRED. A list of hosts exposed by this gateway. At least one host is required. While typically applicable to HTTP services, it can also be used for TCP services using TLS with SNI. May contain a wildcard prefix for the bottom-level component of a domain name. For example `*.foo.com` matches `bar.foo.com` and `*.com` matches `bar.foo.com`, `example.com`, and so on.  **Note**: A `VirtualService` that is bound to a gateway must have one or more hosts that match the hosts specified in a server. The match could be an exact match or a suffix match with the server&#39;s hosts. For example, if the server&#39;s hosts specifies &#34;*.example.com&#34;, VirtualServices with hosts dev.example.com, prod.example.com will match. However, VirtualServices with hosts example.com or newexample.com will not match. |
| tls | [Server.TLSOptions](#istio.networking.v1alpha3.Server.TLSOptions) |  | Set of TLS related options that govern the server&#39;s behavior. Use these options to control if all http requests should be redirected to https, and the TLS modes to use. |






<a name="istio.networking.v1alpha3.Server.TLSOptions"/>

### Server.TLSOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| https_redirect | [bool](#bool) |  | If set to true, the load balancer will send a 302 redirect for all http connections, asking the clients to use HTTPS. |
| mode | [Server.TLSOptions.TLSmode](#istio.networking.v1alpha3.Server.TLSOptions.TLSmode) |  | Optional: Indicates whether connections to this port should be secured using TLS. The value of this field determines how TLS is enforced. |
| server_certificate | [string](#string) |  | REQUIRED if mode is `SIMPLE` or `MUTUAL`. The path to the file holding the server-side TLS certificate to use. |
| private_key | [string](#string) |  | REQUIRED if mode is `SIMPLE` or `MUTUAL`. The path to the file holding the server&#39;s private key. |
| ca_certificates | [string](#string) |  | REQUIRED if mode is `MUTUAL`. The path to a file containing certificate authority certificates to use in verifying a presented client side certificate. |
| subject_alt_names | [string](#string) | repeated | A list of alternate names to verify the subject identity in the certificate presented by the client. |





 

 

 

 



<a name="networking/v1alpha3/service_entry.proto"/>
<p align="right"><a href="#top">Top</a></p>

## networking/v1alpha3/service_entry.proto



<a name="istio.networking.v1alpha3.ServiceEntry"/>

### ServiceEntry
`ServiceEntry` enables adding additional entries into Istio&#39;s internal
service registry, so that auto-discovered services in the mesh can
access/route to these manually specified services. A service entry
describes the properties of a service (DNS name, VIPs ,ports, protocols,
endpoints). These services could be external to the mesh (e.g., web
APIs) or mesh-internal services that are not part of the platform&#39;s
service registry (e.g., a set of VMs talking to services in Kubernetes).

The following configuration adds a set of MongoDB instances running on
unmanaged VMs to Istio&#39;s registry, so that these services can be treated
as any other service in the mesh. The associated DestinationRule is used
to initiate mTLS connections to the database instances.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: ServiceEntry
metadata:
  name: external-svc-mongocluster
spec:
  hosts:
  - mymongodb.somedomain # not used
  addresses:
  - 192.192.192.192/24 # VIPs
  ports:
  - number: 27018
    name: mongodb
    protocol: MONGO
  location: MESH_INTERNAL
  resolution: STATIC
  endpoints:
  - address: 2.2.2.2
  - address: 3.3.3.3
```

and the associated DestinationRule

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: mtls-mongocluster
spec:
  host: mymongodb.somedomain
  trafficPolicy:
    tls:
      mode: MUTUAL
      clientCertificate: /etc/certs/myclientcert.pem
      privateKey: /etc/certs/client_private_key.pem
      caCertificates: /etc/certs/rootcacerts.pem
```

The following example demonstrates the use of wildcards in the hosts for
external services. If the connection has to be routed to the IP address
requested by the application (i.e. application resolves DNS and attempts
to connect to a specific IP), the discovery mode must be set to `NONE`.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: ServiceEntry
metadata:
  name: external-svc-wildcard-example
spec:
  hosts:
  - &#34;*.bar.com&#34;
  location: MESH_EXTERNAL
  ports:
  - number: 80
    name: http
    protocol: HTTP
  resolution: NONE
```


The following example demonstrates a service that is available via a
Unix Domain Socket on the host of the client. The resolution must be
set to STATIC to use unix address endpoints.

    apiVersion: networking.istio.io/v1alpha3
    kind: ServiceEntry
    metadata:
      name: unix-domain-socket-example
    spec:
      hosts:
      - &#34;example.unix.local&#34;
      location: MESH_EXTERNAL
      ports:
      - number: 80
        name: http
        protocol: HTTP
      resolution: STATIC
      endpoints:
      - address: unix:///var/run/example/socket

For HTTP based services, it is possible to create a VirtualService
backed by multiple DNS addressable endpoints. In such a scenario, the
application can use the HTTP_PROXY environment variable to transparently
reroute API calls for the VirtualService to a chosen backend. For
example, the following configuration creates a non-existent external
service called foo.bar.com backed by three domains: us.foo.bar.com:8443,
uk.foo.bar.com:9443, and in.foo.bar.com:7443

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: ServiceEntry
metadata:
  name: external-svc-dns
spec:
  hosts:
  - foo.bar.com
  location: MESH_EXTERNAL
  ports:
  - number: 443
    name: https
    protocol: HTTP
  resolution: DNS
  endpoints:
  - address: us.foo.bar.com
    ports:
      https: 8443
  - address: uk.foo.bar.com
    ports:
      https: 9443
  - address: in.foo.bar.com
    ports:
      https: 7443
```

and a DestinationRule to initiate TLS connections to the ServiceEntry.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: tls-foobar
spec:
  host: foo.bar.com
  trafficPolicy:
    tls:
      mode: SIMPLE # initiates HTTPS
```

With HTTP_PROXY=http://localhost:443, calls from the application to
http://foo.bar.com will be upgraded to HTTPS and load balanced across
the three domains specified above. In other words, a call to
http://foo.bar.com/baz would be translated to
https://uk.foo.bar.com/baz.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hosts | [string](#string) | repeated | REQUIRED. The hosts associated with the ServiceEntry. Could be a DNS name with wildcard prefix (external services only). DNS names in hosts will be ignored if the application accesses the service over non-HTTP protocols such as mongo/opaque TCP/even HTTPS. In such scenarios, the IP addresses specified in the Addresses field or the port will be used to uniquely identify the destination. |
| addresses | [string](#string) | repeated | The virtual IP addresses associated with the service. Could be CIDR prefix. For HTTP services, the addresses field will be ignored and the destination will be identified based on the HTTP Host/Authority header. For non-HTTP protocols such as mongo/opaque TCP/even HTTPS, the hosts will be ignored. If one or more IP addresses are specified, the incoming traffic will be idenfified as belonging to this service if the destination IP matches the IP/CIDRs specified in the addresses field. If the Addresses field is empty, traffic will be identified solely based on the destination port. In such scenarios, the port on which the service is being accessed must not be shared by any other service in the mesh. In other words, the sidecar will behave as a simple TCP proxy, forwarding incoming traffic on a specified port to the specified destination endpoint IP/host. Unix domain socket addresses are not supported in this field. |
| ports | [Port](#istio.networking.v1alpha3.Port) | repeated | REQUIRED. The ports associated with the external service. If the Endpoints are unix domain socket addresses, there must be exactly one port. |
| location | [ServiceEntry.Location](#istio.networking.v1alpha3.ServiceEntry.Location) |  | Specify whether the service should be considered external to the mesh or part of the mesh. |
| resolution | [ServiceEntry.Resolution](#istio.networking.v1alpha3.ServiceEntry.Resolution) |  | Service discovery mode for the hosts. If not set, Istio will attempt to infer the discovery mode based on the value of hosts and endpoints. |
| endpoints | [ServiceEntry.Endpoint](#istio.networking.v1alpha3.ServiceEntry.Endpoint) | repeated | One or more endpoints associated with the service. |






<a name="istio.networking.v1alpha3.ServiceEntry.Endpoint"/>

### ServiceEntry.Endpoint
Endpoint defines a network address (IP or hostname) associated with
the mesh service.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | REQUIRED: Address associated with the network endpoint without the port. Domain names can be used if and only if the resolution is set to DNS, and must be fully-qualified without wildcards. Use the form unix:///absolute/path/to/socket for unix domain socket endpoints. |
| ports | [ServiceEntry.Endpoint.PortsEntry](#istio.networking.v1alpha3.ServiceEntry.Endpoint.PortsEntry) | repeated | Set of ports associated with the endpoint. The ports must be associated with a port name that was declared as part of the service. Do not use for unix:// addresses. |
| labels | [ServiceEntry.Endpoint.LabelsEntry](#istio.networking.v1alpha3.ServiceEntry.Endpoint.LabelsEntry) | repeated | One or more labels associated with the endpoint. |





 


<a name="istio.networking.v1alpha3.ServiceEntry.Location"/>

### ServiceEntry.Location
Location specifies whether the service is part of Istio mesh or
outside the mesh.  Location determines the behavior of several
features, such as service-to-service mTLS authentication, policy
enforcement, etc. When communicating with services outside the mesh,
Istio&#39;s mTLS authentication is disabled, and policy enforcement is
performed on the client-side as opposed to server-side.

| Name | Number | Description |
| ---- | ------ | ----------- |
| MESH_EXTERNAL | 0 | Signifies that the service is external to the mesh. Typically used to indicate external services consumed through APIs. |
| MESH_INTERNAL | 1 | Signifies that the service is part of the mesh. Typically used to indicate services added explicitly as part of expanding the service mesh to include unmanaged infrastructure (e.g., VMs added to a Kubernetes based service mesh). |



<a name="istio.networking.v1alpha3.ServiceEntry.Resolution"/>

### ServiceEntry.Resolution
Resolution determines how the proxy will resolve the IP addresses of
the network endpoints associated with the service, so that it can
route to one of them. The resolution mode specified here has no impact
on how the application resolves the IP address associated with the
service. The application may still have to use DNS to resolve the
service to an IP so that the outbound traffic can be captured by the
Proxy. Alternatively, for HTTP services, the application could
directly communicate with the proxy (e.g., by setting HTTP_PROXY) to
talk to these services.

| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 | Assume that incoming connections have already been resolved (to a specific destination IP address). Such connections are typically routed via the proxy using mechanisms such as IP table REDIRECT/ eBPF. After performing any routing related transformations, the proxy will forward the connection to the IP address to which the connection was bound. |
| STATIC | 1 | Use the static IP addresses specified in endpoints (see below) as the backing instances associated with the service. |
| DNS | 2 | Attempt to resolve the IP address by querying the ambient DNS, during request processing. If no endpoints are specified, the proxy will resolve the DNS address specified in the hosts field, if wildcards are not used. If endpoints are specified, the DNS addresses specified in the endpoints will be resolved to determine the destination IP address. DNS resolution cannot be used with unix domain socket endpoints. |


 

 

 



<a name="networking/v1alpha3/virtual_service.proto"/>
<p align="right"><a href="#top">Top</a></p>

## networking/v1alpha3/virtual_service.proto



<a name="istio.networking.v1alpha3.CorsPolicy"/>

### CorsPolicy
Describes the Cross-Origin Resource Sharing (CORS) policy, for a given
service. Refer to
https://developer.mozilla.org/en-US/docs/Web/HTTP/Access_control_CORS
for further details about cross origin resource sharing. For example,
the following rule restricts cross origin requests to those originating
from example.com domain using HTTP POST/GET, and sets the
Access-Control-Allow-Credentials header to false. In addition, it only
exposes X-Foo-bar header and sets an expiry period of 1 day.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: ratings-route
spec:
  hosts:
  - ratings.prod.svc.cluster.local
  http:
  - route:
    - destination:
        host: ratings.prod.svc.cluster.local
        subset: v1
    corsPolicy:
      allowOrigin:
      - example.com
      allowMethods:
      - POST
      - GET
      allowCredentials: false
      allowHeaders:
      - X-Foo-Bar
      maxAge: &#34;1d&#34;
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| allow_origin | [string](#string) | repeated | The list of origins that are allowed to perform CORS requests. The content will be serialized into the Access-Control-Allow-Origin header. Wildcard * will allow all origins. |
| allow_methods | [string](#string) | repeated | List of HTTP methods allowed to access the resource. The content will be serialized into the Access-Control-Allow-Methods header. |
| allow_headers | [string](#string) | repeated | List of HTTP headers that can be used when requesting the resource. Serialized to Access-Control-Allow-Methods header. |
| expose_headers | [string](#string) | repeated | A white list of HTTP headers that the browsers are allowed to access. Serialized into Access-Control-Expose-Headers header. |
| max_age | [google.protobuf.Duration](#google.protobuf.Duration) |  | Specifies how long the the results of a preflight request can be cached. Translates to the Access-Control-Max-Age header. |
| allow_credentials | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Indicates whether the caller is allowed to send the actual request (not the preflight) using credentials. Translates to Access-Control-Allow-Credentials header. |






<a name="istio.networking.v1alpha3.Destination"/>

### Destination
Destination indicates the network addressable service to which the
request/connection will be sent after processing a routing rule. The
destination.host should unambiguously refer to a service in the service
registry. Istio&#39;s service registry is composed of all the services found
in the platform&#39;s service registry (e.g., Kubernetes services, Consul
services), as well as services declared through the
[ServiceEntry](#ServiceEntry) resource.

*Note for Kubernetes users*: When short names are used (e.g. &#34;reviews&#34;
instead of &#34;reviews.default.svc.cluster.local&#34;), Istio will interpret
the short name based on the namespace of the rule, not the service. A
rule in the &#34;default&#34; namespace containing a host &#34;reviews will be
interpreted as &#34;reviews.default.svc.cluster.local&#34;, irrespective of the
actual namespace associated with the reviews service. _To avoid potential
misconfigurations, it is recommended to always use fully qualified
domain names over short names._

The following Kubernetes example routes all traffic by default to pods
of the reviews service with label &#34;version: v1&#34; (i.e., subset v1), and
some to subset v2, in a kubernetes environment.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: reviews-route
  namespace: foo
spec:
  hosts:
  - reviews # interpreted as reviews.foo.svc.cluster.local
  http:
  - match:
    - uri:
        prefix: &#34;/wpcatalog&#34;
    - uri:
        prefix: &#34;/consumercatalog&#34;
    rewrite:
      uri: &#34;/newcatalog&#34;
    route:
    - destination:
        host: reviews # interpreted as reviews.foo.svc.cluster.local
        subset: v2
  - route:
    - destination:
        host: reviews # interpreted as reviews.foo.svc.cluster.local
        subset: v1
```

And the associated DestinationRule

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: reviews-destination
  namespace: foo
spec:
  host: reviews # interpreted as reviews.foo.svc.cluster.local
  subsets:
  - name: v1
    labels:
      version: v1
  - name: v2
    labels:
      version: v2
```

The following VirtualService sets a timeout of 5s for all calls to
productpage.prod.svc.cluster.local service in Kubernetes. Notice that
there are no subsets defined in this rule. Istio will fetch all
instances of productpage.prod.svc.cluster.local service from the service
registry and populate the sidecar&#39;s load balancing pool. Also, notice
that this rule is set in the istio-system namespace but uses the fully
qualified domain name of the productpage service,
productpage.prod.svc.cluster.local. Therefore the rule&#39;s namespace does
not have an impact in resolving the name of the productpage service.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: my-productpage-rule
  namespace: istio-system
spec:
  hosts:
  - productpage.prod.svc.cluster.local # ignores rule namespace
  http:
  - timeout: 5s
    route:
    - destination:
        host: productpage.prod.svc.cluster.local
```

To control routing for traffic bound to services outside the mesh, external
services must first be added to Istio&#39;s internal service registry using the
ServiceEntry resource. VirtualServices can then be defined to control traffic
bound to these external services. For example, the following rules define a
Service for wikipedia.org and set a timeout of 5s for http requests.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: ServiceEntry
metadata:
  name: external-svc-wikipedia
spec:
  hosts:
  - wikipedia.org
  location: MESH_EXTERNAL
  ports:
  - number: 80
    name: example-http
    protocol: HTTP
  resolution: DNS

apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: my-wiki-rule
spec:
  hosts:
  - wikipedia.org
  http:
  - timeout: 5s
    route:
    - destination:
        host: wikipedia.org
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| host | [string](#string) |  | REQUIRED. The name of a service from the service registry. Service names are looked up from the platform&#39;s service registry (e.g., Kubernetes services, Consul services, etc.) and from the hosts declared by [ServiceEntry](#ServiceEntry). Traffic forwarded to destinations that are not found in either of the two, will be dropped.  *Note for Kubernetes users*: When short names are used (e.g. &#34;reviews&#34; instead of &#34;reviews.default.svc.cluster.local&#34;), Istio will interpret the short name based on the namespace of the rule, not the service. A rule in the &#34;default&#34; namespace containing a host &#34;reviews will be interpreted as &#34;reviews.default.svc.cluster.local&#34;, irrespective of the actual namespace associated with the reviews service. _To avoid potential misconfigurations, it is recommended to always use fully qualified domain names over short names._ |
| subset | [string](#string) |  | The name of a subset within the service. Applicable only to services within the mesh. The subset must be defined in a corresponding DestinationRule. |
| port | [PortSelector](#istio.networking.v1alpha3.PortSelector) |  | Specifies the port on the host that is being addressed. If a service exposes only a single port it is not required to explicitly select the port. |






<a name="istio.networking.v1alpha3.DestinationWeight"/>

### DestinationWeight
Each routing rule is associated with one or more service versions (see
glossary in beginning of document). Weights associated with the version
determine the proportion of traffic it receives. For example, the
following rule will route 25% of traffic for the &#34;reviews&#34; service to
instances with the &#34;v2&#34; tag and the remaining traffic (i.e., 75%) to
&#34;v1&#34;.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: reviews-route
spec:
  hosts:
  - reviews.prod.svc.cluster.local
  http:
  - route:
    - destination:
        host: reviews.prod.svc.cluster.local
        subset: v2
      weight: 25
    - destination:
        host: reviews.prod.svc.cluster.local
        subset: v1
      weight: 75
```

And the associated DestinationRule

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: reviews-destination
spec:
  host: reviews.prod.svc.cluster.local
  subsets:
  - name: v1
    labels:
      version: v1
  - name: v2
    labels:
      version: v2
```

Traffic can also be split across two entirely different services without
having to define new subsets. For example, the following rule forwards 25% of
traffic to reviews.com to dev.reviews.com

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: reviews-route-two-domains
spec:
  hosts:
  - reviews.com
  http:
  - route:
    - destination:
        host: dev.reviews.com
      weight: 25
    - destination:
        host: reviews.com
      weight: 75
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| destination | [Destination](#istio.networking.v1alpha3.Destination) |  | REQUIRED. Destination uniquely identifies the instances of a service to which the request/connection should be forwarded to. |
| weight | [int32](#int32) |  | REQUIRED. The proportion of traffic to be forwarded to the service version. (0-100). Sum of weights across destinations SHOULD BE == 100. If there is only destination in a rule, the weight value is assumed to be 100. |






<a name="istio.networking.v1alpha3.HTTPFaultInjection"/>

### HTTPFaultInjection
HTTPFaultInjection can be used to specify one or more faults to inject
while forwarding http requests to the destination specified in a route.
Fault specification is part of a VirtualService rule. Faults include
aborting the Http request from downstream service, and/or delaying
proxying of requests. A fault rule MUST HAVE delay or abort or both.

*Note:* Delay and abort faults are independent of one another, even if
both are specified simultaneously.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delay | [HTTPFaultInjection.Delay](#istio.networking.v1alpha3.HTTPFaultInjection.Delay) |  | Delay requests before forwarding, emulating various failures such as network issues, overloaded upstream service, etc. |
| abort | [HTTPFaultInjection.Abort](#istio.networking.v1alpha3.HTTPFaultInjection.Abort) |  | Abort Http request attempts and return error codes back to downstream service, giving the impression that the upstream service is faulty. |






<a name="istio.networking.v1alpha3.HTTPFaultInjection.Abort"/>

### HTTPFaultInjection.Abort
Abort specification is used to prematurely abort a request with a
pre-specified error code. The following example will return an HTTP
400 error code for 10% of the requests to the &#34;ratings&#34; service &#34;v1&#34;.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: ratings-route
spec:
  hosts:
  - ratings.prod.svc.cluster.local
  http:
  - route:
    - destination:
        host: ratings.prod.svc.cluster.local
        subset: v1
    fault:
      abort:
        percent: 10
        httpStatus: 400
```

The _httpStatus_ field is used to indicate the HTTP status code to
return to the caller. The optional _percent_ field, a value between 0
and 100, is used to only abort a certain percentage of requests. If
not specified, all requests are aborted.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| percent | [int32](#int32) |  | Percentage of requests to be aborted with the error code provided (0-100). |
| http_status | [int32](#int32) |  | REQUIRED. HTTP status code to use to abort the Http request. |
| grpc_status | [string](#string) |  | $hide_from_docs |
| http2_error | [string](#string) |  | $hide_from_docs |






<a name="istio.networking.v1alpha3.HTTPFaultInjection.Delay"/>

### HTTPFaultInjection.Delay
Delay specification is used to inject latency into the request
forwarding path. The following example will introduce a 5 second delay
in 10% of the requests to the &#34;v1&#34; version of the &#34;reviews&#34;
service from all pods with label env: prod

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: reviews-route
spec:
  hosts:
  - reviews.prod.svc.cluster.local
  http:
  - match:
    - sourceLabels:
        env: prod
    route:
    - destination:
        host: reviews.prod.svc.cluster.local
        subset: v1
    fault:
      delay:
        percent: 10
        fixedDelay: 5s
```

The _fixedDelay_ field is used to indicate the amount of delay in
seconds. An optional _percent_ field, a value between 0 and 100, can
be used to only delay a certain percentage of requests. If left
unspecified, all request will be delayed.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| percent | [int32](#int32) |  | Percentage of requests on which the delay will be injected (0-100). |
| fixed_delay | [google.protobuf.Duration](#google.protobuf.Duration) |  | REQUIRED. Add a fixed delay before forwarding the request. Format: 1h/1m/1s/1ms. MUST be &gt;=1ms. |
| exponential_delay | [google.protobuf.Duration](#google.protobuf.Duration) |  | $hide_from_docs |






<a name="istio.networking.v1alpha3.HTTPMatchRequest"/>

### HTTPMatchRequest
HttpMatchRequest specifies a set of criterion to be met in order for the
rule to be applied to the HTTP request. For example, the following
restricts the rule to match only requests where the URL path
starts with /ratings/v2/ and the request contains a `cookie` with value
`user=jason`.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: ratings-route
spec:
  hosts:
  - ratings.prod.svc.cluster.local
  http:
  - match:
    - headers:
        cookie:
          regex: &#34;^(.*?;)?(user=jason)(;.*)?&#34;
        uri:
          prefix: &#34;/ratings/v2/&#34;
    route:
    - destination:
        host: ratings.prod.svc.cluster.local
```

HTTPMatchRequest CANNOT be empty.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uri | [StringMatch](#istio.networking.v1alpha3.StringMatch) |  | URI to match values are case-sensitive and formatted as follows:  - `exact: &#34;value&#34;` for exact string match  - `prefix: &#34;value&#34;` for prefix-based match  - `regex: &#34;value&#34;` for ECMAscript style regex-based match |
| scheme | [StringMatch](#istio.networking.v1alpha3.StringMatch) |  | URI Scheme values are case-sensitive and formatted as follows:  - `exact: &#34;value&#34;` for exact string match  - `prefix: &#34;value&#34;` for prefix-based match  - `regex: &#34;value&#34;` for ECMAscript style regex-based match |
| method | [StringMatch](#istio.networking.v1alpha3.StringMatch) |  | HTTP Method values are case-sensitive and formatted as follows:  - `exact: &#34;value&#34;` for exact string match  - `prefix: &#34;value&#34;` for prefix-based match  - `regex: &#34;value&#34;` for ECMAscript style regex-based match |
| authority | [StringMatch](#istio.networking.v1alpha3.StringMatch) |  | HTTP Authority values are case-sensitive and formatted as follows:  - `exact: &#34;value&#34;` for exact string match  - `prefix: &#34;value&#34;` for prefix-based match  - `regex: &#34;value&#34;` for ECMAscript style regex-based match |
| headers | [HTTPMatchRequest.HeadersEntry](#istio.networking.v1alpha3.HTTPMatchRequest.HeadersEntry) | repeated | The header keys must be lowercase and use hyphen as the separator, e.g. _x-request-id_.  Header values are case-sensitive and formatted as follows:  - `exact: &#34;value&#34;` for exact string match  - `prefix: &#34;value&#34;` for prefix-based match  - `regex: &#34;value&#34;` for ECMAscript style regex-based match  **Note:** The keys `uri`, `scheme`, `method`, and `authority` will be ignored. |
| port | [uint32](#uint32) |  | Specifies the ports on the host that is being addressed. Many services only expose a single port or label ports with the protocols they support, in these cases it is not required to explicitly select the port. |
| source_labels | [HTTPMatchRequest.SourceLabelsEntry](#istio.networking.v1alpha3.HTTPMatchRequest.SourceLabelsEntry) | repeated | One or more labels that constrain the applicability of a rule to workloads with the given labels. If the VirtualService has a list of gateways specified at the top, it should include the reserved gateway `mesh` in order for this field to be applicable. |
| gateways | [string](#string) | repeated | Names of gateways where the rule should be applied to. Gateway names at the top of the VirtualService (if any) are overridden. The gateway match is independent of sourceLabels. |






<a name="istio.networking.v1alpha3.HTTPMatchRequest.HeadersEntry"/>

### HTTPMatchRequest.HeadersEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [StringMatch](#istio.networking.v1alpha3.StringMatch) |  |  |






<a name="istio.networking.v1alpha3.HTTPMatchRequest.SourceLabelsEntry"/>

### HTTPMatchRequest.SourceLabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="istio.networking.v1alpha3.HTTPRedirect"/>

### HTTPRedirect
HTTPRedirect can be used to send a 302 redirect response to the caller,
where the Authority/Host and the URI in the response can be swapped with
the specified values. For example, the following rule redirects
requests for /v1/getProductRatings API on the ratings service to
/v1/bookRatings provided by the bookratings service.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: ratings-route
spec:
  hosts:
  - ratings.prod.svc.cluster.local
  http:
  - match:
    - uri:
        exact: /v1/getProductRatings
  redirect:
    uri: /v1/bookRatings
    authority: newratings.default.svc.cluster.local
  ...
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uri | [string](#string) |  | On a redirect, overwrite the Path portion of the URL with this value. Note that the entire path will be replaced, irrespective of the request URI being matched as an exact path or prefix. |
| authority | [string](#string) |  | On a redirect, overwrite the Authority/Host portion of the URL with this value. |






<a name="istio.networking.v1alpha3.HTTPRetry"/>

### HTTPRetry
Describes the retry policy to use when a HTTP request fails. For
example, the following rule sets the maximum number of retries to 3 when
calling ratings:v1 service, with a 2s timeout per retry attempt.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: ratings-route
spec:
  hosts:
  - ratings.prod.svc.cluster.local
  http:
  - route:
    - destination:
        host: ratings.prod.svc.cluster.local
        subset: v1
    retries:
      attempts: 3
      perTryTimeout: 2s
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| attempts | [int32](#int32) |  | REQUIRED. Number of retries for a given request. The interval between retries will be determined automatically (25ms&#43;). Actual number of retries attempted depends on the httpReqTimeout. |
| per_try_timeout | [google.protobuf.Duration](#google.protobuf.Duration) |  | Timeout per retry attempt for a given request. format: 1h/1m/1s/1ms. MUST BE &gt;=1ms. |






<a name="istio.networking.v1alpha3.HTTPRewrite"/>

### HTTPRewrite
HTTPRewrite can be used to rewrite specific parts of a HTTP request
before forwarding the request to the destination. Rewrite primitive can
be used only with the DestinationWeights. The following example
demonstrates how to rewrite the URL prefix for api call (/ratings) to
ratings service before making the actual API call.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: ratings-route
spec:
  hosts:
  - ratings.prod.svc.cluster.local
  http:
  - match:
    - uri:
        prefix: /ratings
    rewrite:
      uri: /v1/bookRatings
    route:
    - destination:
        host: ratings.prod.svc.cluster.local
        subset: v1
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uri | [string](#string) |  | rewrite the path (or the prefix) portion of the URI with this value. If the original URI was matched based on prefix, the value provided in this field will replace the corresponding matched prefix. |
| authority | [string](#string) |  | rewrite the Authority/Host header with this value. |






<a name="istio.networking.v1alpha3.HTTPRoute"/>

### HTTPRoute
Describes match conditions and actions for routing HTTP/1.1, HTTP2, and
gRPC traffic. See VirtualService for usage examples.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| match | [HTTPMatchRequest](#istio.networking.v1alpha3.HTTPMatchRequest) | repeated | Match conditions to be satisfied for the rule to be activated. All conditions inside a single match block have AND semantics, while the list of match blocks have OR semantics. The rule is matched if any one of the match blocks succeed. |
| route | [DestinationWeight](#istio.networking.v1alpha3.DestinationWeight) | repeated | A http rule can either redirect or forward (default) traffic. The forwarding target can be one of several versions of a service (see glossary in beginning of document). Weights associated with the service version determine the proportion of traffic it receives. |
| redirect | [HTTPRedirect](#istio.networking.v1alpha3.HTTPRedirect) |  | A http rule can either redirect or forward (default) traffic. If traffic passthrough option is specified in the rule, route/redirect will be ignored. The redirect primitive can be used to send a HTTP 302 redirect to a different URI or Authority. |
| rewrite | [HTTPRewrite](#istio.networking.v1alpha3.HTTPRewrite) |  | Rewrite HTTP URIs and Authority headers. Rewrite cannot be used with Redirect primitive. Rewrite will be performed before forwarding. |
| websocket_upgrade | [bool](#bool) |  | Indicates that a HTTP/1.1 client connection to this particular route should be allowed (and expected) to upgrade to a WebSocket connection. The default is false. Istio&#39;s reference sidecar implementation (Envoy) expects the first request to this route to contain the WebSocket upgrade headers. Otherwise, the request will be rejected. Note that Websocket allows secondary protocol negotiation which may then be subject to further routing rules based on the protocol selected. |
| timeout | [google.protobuf.Duration](#google.protobuf.Duration) |  | Timeout for HTTP requests. |
| retries | [HTTPRetry](#istio.networking.v1alpha3.HTTPRetry) |  | Retry policy for HTTP requests. |
| fault | [HTTPFaultInjection](#istio.networking.v1alpha3.HTTPFaultInjection) |  | Fault injection policy to apply on HTTP traffic. |
| mirror | [Destination](#istio.networking.v1alpha3.Destination) |  | Mirror HTTP traffic to a another destination in addition to forwarding the requests to the intended destination. Mirrored traffic is on a best effort basis where the sidecar/gateway will not wait for the mirrored cluster to respond before returning the response from the original destination. Statistics will be generated for the mirrored destination. |
| cors_policy | [CorsPolicy](#istio.networking.v1alpha3.CorsPolicy) |  | Cross-Origin Resource Sharing policy (CORS). Refer to https://developer.mozilla.org/en-US/docs/Web/HTTP/Access_control_CORS for further details about cross origin resource sharing. |
| append_headers | [HTTPRoute.AppendHeadersEntry](#istio.networking.v1alpha3.HTTPRoute.AppendHeadersEntry) | repeated | Additional HTTP headers to add before forwarding a request to the destination service. |
| remove_response_headers | [HTTPRoute.RemoveResponseHeadersEntry](#istio.networking.v1alpha3.HTTPRoute.RemoveResponseHeadersEntry) | repeated | Http headers to remove before returning the response to the caller $hide_from_docs |






<a name="istio.networking.v1alpha3.HTTPRoute.AppendHeadersEntry"/>

### HTTPRoute.AppendHeadersEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="istio.networking.v1alpha3.HTTPRoute.RemoveResponseHeadersEntry"/>

### HTTPRoute.RemoveResponseHeadersEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="istio.networking.v1alpha3.L4MatchAttributes"/>

### L4MatchAttributes
L4 connection match attributes. Note that L4 connection matching support
is incomplete.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| destination_subnet | [string](#string) |  | IPv4 or IPv6 ip address of destination with optional subnet. E.g., a.b.c.d/xx form or just a.b.c.d. This is only valid when the destination service has several IPs and the application explicitly specifies a particular IP. |
| port | [uint32](#uint32) |  | Specifies the port on the host that is being addressed. Many services only expose a single port or label ports with the protocols they support, in these cases it is not required to explicitly select the port. |
| source_subnet | [string](#string) |  | IPv4 or IPv6 ip address of source with optional subnet. E.g., a.b.c.d/xx form or just a.b.c.d |
| source_labels | [L4MatchAttributes.SourceLabelsEntry](#istio.networking.v1alpha3.L4MatchAttributes.SourceLabelsEntry) | repeated | One or more labels that constrain the applicability of a rule to workloads with the given labels. If the VirtualService has a list of gateways specified at the top, it should include the reserved gateway `mesh` in order for this field to be applicable. |
| gateways | [string](#string) | repeated | Names of gateways where the rule should be applied to. Gateway names at the top of the VirtualService (if any) are overridden. The gateway match is independent of sourceLabels. |






<a name="istio.networking.v1alpha3.L4MatchAttributes.SourceLabelsEntry"/>

### L4MatchAttributes.SourceLabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="istio.networking.v1alpha3.PortSelector"/>

### PortSelector
PortSelector specifies the number of a port to be used for
matching or selection for final routing.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| number | [uint32](#uint32) |  | Valid port number |
| name | [string](#string) |  | $hide_from_docs |






<a name="istio.networking.v1alpha3.StringMatch"/>

### StringMatch
Describes how to match a given string in HTTP headers. Match is
case-sensitive.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exact | [string](#string) |  | exact string match |
| prefix | [string](#string) |  | prefix-based match |
| regex | [string](#string) |  | ECMAscript style regex-based match |






<a name="istio.networking.v1alpha3.TCPRoute"/>

### TCPRoute
Describes match conditions and actions for routing TCP traffic. The
following routing rule forwards traffic arriving at port 27017 for
mongo.prod.svc.cluster.local from 172.17.16.* subnet to another Mongo
server on port 5555.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: bookinfo-Mongo
spec:
  hosts:
  - mongo.prod.svc.cluster.local
  tcp:
  - match:
    - port: 27017
      sourceSubnet: &#34;172.17.16.0/24&#34;
    route:
    - destination:
        host: mongo.backup.svc.cluster.local
        port:
          number: 5555
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| match | [L4MatchAttributes](#istio.networking.v1alpha3.L4MatchAttributes) | repeated | Match conditions to be satisfied for the rule to be activated. All conditions inside a single match block have AND semantics, while the list of match blocks have OR semantics. The rule is matched if any one of the match blocks succeed. |
| route | [DestinationWeight](#istio.networking.v1alpha3.DestinationWeight) | repeated | The destination to which the connection should be forwarded to. Currently, only one destination is allowed for TCP services. When TCP weighted routing support is introduced in Envoy, multiple destinations with weights can be specified. |






<a name="istio.networking.v1alpha3.VirtualService"/>

### VirtualService
A `VirtualService` defines a set of traffic routing rules to apply when a host is
addressed. Each routing rule defines matching criteria for traffic of a specific
protocol. If the traffic is matched, then it is sent to a named destination service
(or subset/version of it) defined in the registry.

The source of traffic can also be matched in a routing rule. This allows routing
to be customized for specific client contexts.

The following example on Kubernetes, routes all HTTP traffic by default to
pods of the reviews service with label &#34;version: v1&#34;. In addition,
HTTP requests containing /wpcatalog/, /consumercatalog/ url prefixes will
be rewritten to /newcatalog and sent to pods with label &#34;version: v2&#34;.


```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: reviews-route
spec:
  hosts:
  - reviews.prod.svc.cluster.local
  http:
  - match:
    - uri:
        prefix: &#34;/wpcatalog&#34;
    - uri:
        prefix: &#34;/consumercatalog&#34;
    rewrite:
      uri: &#34;/newcatalog&#34;
    route:
    - destination:
        host: reviews.prod.svc.cluster.local
        subset: v2
  - route:
    - destination:
        host: reviews.prod.svc.cluster.local
        subset: v1
```

A subset/version of a route destination is identified with a reference
to a named service subset which must be declared in a corresponding
`DestinationRule`.

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: reviews-destination
spec:
  host: reviews.prod.svc.cluster.local
  subsets:
  - name: v1
    labels:
      version: v1
  - name: v2
    labels:
      version: v2
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hosts | [string](#string) | repeated | REQUIRED. The destination hosts to which traffic is being sent. Could be a DNS name with wildcard prefix or an IP address. Depending on the platform, short-names can also be used instead of a FQDN (i.e. has no dots in the name). In such a scenario, the FQDN of the host would be derived based on the underlying platform.  **A host name can be defined by only one VirtualService**. A single VirtualService can be used to describe traffic properties for multiple HTTP and TCP ports.  *Note for Kubernetes users*: When short names are used (e.g. &#34;reviews&#34; instead of &#34;reviews.default.svc.cluster.local&#34;), Istio will interpret the short name based on the namespace of the rule, not the service. A rule in the &#34;default&#34; namespace containing a host &#34;reviews will be interpreted as &#34;reviews.default.svc.cluster.local&#34;, irrespective of the actual namespace associated with the reviews service. _To avoid potential misconfigurations, it is recommended to always use fully qualified domain names over short names._  The hosts field applies to both HTTP and TCP services. Service inside the mesh, i.e., those found in the service registry, must always be referred to using their alphanumeric names. IP addresses are allowed only for services defined via the Gateway. |
| gateways | [string](#string) | repeated | The names of gateways and sidecars that should apply these routes. A single VirtualService is used for sidecars inside the mesh as well as for one or more gateways. The selection condition imposed by this field can be overridden using the source field in the match conditions of HTTP/TCP routes. The reserved word `mesh` is used to imply all the sidecars in the mesh. When this field is omitted, the default gateway (`mesh`) will be used, which would apply the rule to all sidecars in the mesh. If a list of gateway names is provided, the rules will apply only to the gateways. To apply the rules to both gateways and sidecars, specify `mesh` as one of the gateway names. |
| http | [HTTPRoute](#istio.networking.v1alpha3.HTTPRoute) | repeated | An ordered list of route rules for HTTP traffic. The first rule matching an incoming request is used. |
| tcp | [TCPRoute](#istio.networking.v1alpha3.TCPRoute) | repeated | An ordered list of route rules for TCP traffic. The first rule matching an incoming request is used. |





 

 

 

 



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

