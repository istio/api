// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/destination_rule.proto

// `DestinationRule` defines policies that apply to traffic intended for a
// service after routing has occurred. These rules specify configuration
// for load balancing, connection pool size from the sidecar, and outlier
// detection settings to detect and evict unhealthy hosts from the load
// balancing pool. For example, a simple load balancing policy for the
// ratings service would look as follows:
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: DestinationRule
// metadata:
//   name: bookinfo-ratings
// spec:
//   host: ratings.prod.svc.cluster.local
//   trafficPolicy:
//     loadBalancer:
//       simple: LEAST_REQUEST
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: DestinationRule
// metadata:
//   name: bookinfo-ratings
// spec:
//   host: ratings.prod.svc.cluster.local
//   trafficPolicy:
//     loadBalancer:
//       simple: LEAST_REQUEST
// ```
// {{</tab>}}
// {{</tabset>}}
//
// Version specific policies can be specified by defining a named
// `subset` and overriding the settings specified at the service level. The
// following rule uses a round robin load balancing policy for all traffic
// going to a subset named testversion that is composed of endpoints (e.g.,
// pods) with labels (version:v3).
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: DestinationRule
// metadata:
//   name: bookinfo-ratings
// spec:
//   host: ratings.prod.svc.cluster.local
//   trafficPolicy:
//     loadBalancer:
//       simple: LEAST_REQUEST
//   subsets:
//   - name: testversion
//     labels:
//       version: v3
//     trafficPolicy:
//       loadBalancer:
//         simple: ROUND_ROBIN
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: DestinationRule
// metadata:
//   name: bookinfo-ratings
// spec:
//   host: ratings.prod.svc.cluster.local
//   trafficPolicy:
//     loadBalancer:
//       simple: LEAST_REQUEST
//   subsets:
//   - name: testversion
//     labels:
//       version: v3
//     trafficPolicy:
//       loadBalancer:
//         simple: ROUND_ROBIN
// ```
// {{</tab>}}
// {{</tabset>}}
//
// **Note:** Policies specified for subsets will not take effect until
// a route rule explicitly sends traffic to this subset.
//
// Traffic policies can be customized to specific ports as well. The
// following rule uses the least connection load balancing policy for all
// traffic to port 80, while uses a round robin load balancing setting for
// traffic to the port 9080.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: DestinationRule
// metadata:
//   name: bookinfo-ratings-port
// spec:
//   host: ratings.prod.svc.cluster.local
//   trafficPolicy: # Apply to all ports
//     portLevelSettings:
//     - port:
//         number: 80
//       loadBalancer:
//         simple: LEAST_REQUEST
//     - port:
//         number: 9080
//       loadBalancer:
//         simple: ROUND_ROBIN
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: DestinationRule
// metadata:
//   name: bookinfo-ratings-port
// spec:
//   host: ratings.prod.svc.cluster.local
//   trafficPolicy: # Apply to all ports
//     portLevelSettings:
//     - port:
//         number: 80
//       loadBalancer:
//         simple: LEAST_REQUEST
//     - port:
//         number: 9080
//       loadBalancer:
//         simple: ROUND_ROBIN
// ```
// {{</tab>}}
//
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: DestinationRule
// metadata:
//   name: configure-client-mtls-dr
//   workloadSelector:
//     labels:
//       app: ratings
//   spec:
//     trafficPolicy:
//       loadBalancer:
//         simple: ROUND_ROBIN
//       portLevelSettings:
//         - port:
//             number: 31443
//           tls:
//             credentialName: client-credential
//             mode: MUTUAL
// ```
// {{</tab>}}
// {{</tabset>}}
//

package v1alpha3

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for DestinationRule
func (this *DestinationRule) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DestinationRule
func (this *DestinationRule) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicy
func (this *TrafficPolicy) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicy
func (this *TrafficPolicy) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicy_PortTrafficPolicy
func (this *TrafficPolicy_PortTrafficPolicy) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicy_PortTrafficPolicy
func (this *TrafficPolicy_PortTrafficPolicy) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Subset
func (this *Subset) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Subset
func (this *Subset) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for LoadBalancerSettings
func (this *LoadBalancerSettings) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LoadBalancerSettings
func (this *LoadBalancerSettings) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for LoadBalancerSettings_ConsistentHashLB
func (this *LoadBalancerSettings_ConsistentHashLB) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LoadBalancerSettings_ConsistentHashLB
func (this *LoadBalancerSettings_ConsistentHashLB) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for LoadBalancerSettings_ConsistentHashLB_HTTPCookie
func (this *LoadBalancerSettings_ConsistentHashLB_HTTPCookie) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LoadBalancerSettings_ConsistentHashLB_HTTPCookie
func (this *LoadBalancerSettings_ConsistentHashLB_HTTPCookie) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ConnectionPoolSettings
func (this *ConnectionPoolSettings) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ConnectionPoolSettings
func (this *ConnectionPoolSettings) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ConnectionPoolSettings_TCPSettings
func (this *ConnectionPoolSettings_TCPSettings) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ConnectionPoolSettings_TCPSettings
func (this *ConnectionPoolSettings_TCPSettings) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ConnectionPoolSettings_TCPSettings_TcpKeepalive
func (this *ConnectionPoolSettings_TCPSettings_TcpKeepalive) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ConnectionPoolSettings_TCPSettings_TcpKeepalive
func (this *ConnectionPoolSettings_TCPSettings_TcpKeepalive) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ConnectionPoolSettings_HTTPSettings
func (this *ConnectionPoolSettings_HTTPSettings) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ConnectionPoolSettings_HTTPSettings
func (this *ConnectionPoolSettings_HTTPSettings) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for OutlierDetection
func (this *OutlierDetection) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for OutlierDetection
func (this *OutlierDetection) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ClientTLSSettings
func (this *ClientTLSSettings) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ClientTLSSettings
func (this *ClientTLSSettings) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for LocalityLoadBalancerSetting
func (this *LocalityLoadBalancerSetting) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LocalityLoadBalancerSetting
func (this *LocalityLoadBalancerSetting) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for LocalityLoadBalancerSetting_Distribute
func (this *LocalityLoadBalancerSetting_Distribute) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LocalityLoadBalancerSetting_Distribute
func (this *LocalityLoadBalancerSetting_Distribute) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for LocalityLoadBalancerSetting_Failover
func (this *LocalityLoadBalancerSetting_Failover) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LocalityLoadBalancerSetting_Failover
func (this *LocalityLoadBalancerSetting_Failover) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	DestinationRuleMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	DestinationRuleUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{AllowUnknownFields: true}
)
