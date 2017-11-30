// Code generated by protoc-gen-go. DO NOT EDIT.
// source: routing/v1alpha1/dest_policy.proto

/*
Package istio_routing_v1alpha1 is a generated protocol buffer package.

It is generated from these files:
	routing/v1alpha1/dest_policy.proto
	routing/v1alpha1/egress_rule.proto
	routing/v1alpha1/http_fault.proto
	routing/v1alpha1/ingress_rule.proto
	routing/v1alpha1/l4_fault.proto
	routing/v1alpha1/route_rule.proto

It has these top-level messages:
	DestinationPolicy
	LoadBalancing
	CircuitBreaker
	EgressRule
	HTTPFaultInjection
	IngressRule
	L4FaultInjection
	RouteRule
	IstioService
	MatchCondition
	MatchRequest
	DestinationWeight
	L4MatchAttributes
	HTTPRedirect
	HTTPRewrite
	StringMatch
	HTTPTimeout
	HTTPRetry
	CorsPolicy
	Port
*/
package istio_routing_v1alpha1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Load balancing algorithms supported by Envoy.
type LoadBalancing_SimpleLBPolicy int32

const (
	// Simple round robin policy.
	LoadBalancing_ROUND_ROBIN LoadBalancing_SimpleLBPolicy = 0
	// The least request load balancer uses an O(1) algorithm which selects
	// two random healthy hosts and picks the host which has fewer active
	// requests.
	LoadBalancing_LEAST_CONN LoadBalancing_SimpleLBPolicy = 1
	// The random load balancer selects a random healthy host. The random
	// load balancer generally performs better than round robin if no health
	// checking policy is configured.
	LoadBalancing_RANDOM LoadBalancing_SimpleLBPolicy = 2
)

var LoadBalancing_SimpleLBPolicy_name = map[int32]string{
	0: "ROUND_ROBIN",
	1: "LEAST_CONN",
	2: "RANDOM",
}
var LoadBalancing_SimpleLBPolicy_value = map[string]int32{
	"ROUND_ROBIN": 0,
	"LEAST_CONN":  1,
	"RANDOM":      2,
}

func (x LoadBalancing_SimpleLBPolicy) String() string {
	return proto.EnumName(LoadBalancing_SimpleLBPolicy_name, int32(x))
}
func (LoadBalancing_SimpleLBPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

// DestinationPolicy defines client/caller-side policies that determine how
// to handle traffic bound to a particular destination service. The policy
// specifies configuration for load balancing and circuit breakers. For
// example, a simple load balancing policy for the ratings service would
// look as follows:
//
//     metadata:
//       name: ratings-lb-policy
//       namespace: default # optional (default is "default")
//     spec:
//       destination:
//         name: ratings
//       loadBalancing:
//         name: ROUND_ROBIN
//
// The FQDN of the destination service is composed from the destination name and meta namespace fields, along with
// a platform-specific domain suffix
// (e.g. on Kubernetes, "reviews" + "default" + "svc.cluster.local" -> "reviews.default.svc.cluster.local").
//
// A destination policy can be restricted to a particular version of a
// service or applied to all versions. It can also be restricted to calls from
// a particular source. For example, the following load balancing policy
// applies to version v1 of the ratings service running in the prod
// environment but only when called from version v2 of the reviews service:
//
//     metadata:
//       name: ratings-lb-policy
//       namespace: default
//     spec:
//       source:
//         name: reviews
//         labels:
//           version: v2
//       destination:
//         name: ratings
//         labels:
//           env: prod
//           version: v1
//       loadBalancing:
//         name: ROUND_ROBIN
//
// *Note:* Destination policies will be applied only if the corresponding
// tagged instances are explicitly routed to. In other words, for every
// destination policy defined, at least one route rule must refer to the
// service version indicated in the destination policy.
type DestinationPolicy struct {
	// Optional: Destination uniquely identifies the destination service associated
	// with this policy.
	Destination *IstioService `protobuf:"bytes,1,opt,name=destination" json:"destination,omitempty"`
	// Optional: Source uniquely identifies the source service associated
	// with this policy.
	Source *IstioService `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	// Load balancing policy.
	LoadBalancing *LoadBalancing `protobuf:"bytes,3,opt,name=load_balancing,json=loadBalancing" json:"load_balancing,omitempty"`
	// Circuit breaker policy.
	CircuitBreaker *CircuitBreaker `protobuf:"bytes,4,opt,name=circuit_breaker,json=circuitBreaker" json:"circuit_breaker,omitempty"`
	// (-- Other custom policy implementations --)
	Custom *google_protobuf.Any `protobuf:"bytes,100,opt,name=custom" json:"custom,omitempty"`
}

func (m *DestinationPolicy) Reset()                    { *m = DestinationPolicy{} }
func (m *DestinationPolicy) String() string            { return proto.CompactTextString(m) }
func (*DestinationPolicy) ProtoMessage()               {}
func (*DestinationPolicy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DestinationPolicy) GetDestination() *IstioService {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *DestinationPolicy) GetSource() *IstioService {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *DestinationPolicy) GetLoadBalancing() *LoadBalancing {
	if m != nil {
		return m.LoadBalancing
	}
	return nil
}

func (m *DestinationPolicy) GetCircuitBreaker() *CircuitBreaker {
	if m != nil {
		return m.CircuitBreaker
	}
	return nil
}

func (m *DestinationPolicy) GetCustom() *google_protobuf.Any {
	if m != nil {
		return m.Custom
	}
	return nil
}

// Load balancing policy to use when forwarding traffic. These policies
// directly correlate to [load balancer
// types](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/load_balancing)
// supported by Envoy. Example,
//
//     metadata:
//       name: reviews-lb-policy
//       namespace: default
//     spec:
//       destination:
//         name: reviews
//       loadBalancing:
//         name: RANDOM
//
type LoadBalancing struct {
	// Types that are valid to be assigned to LbPolicy:
	//	*LoadBalancing_Name
	//	*LoadBalancing_Custom
	LbPolicy isLoadBalancing_LbPolicy `protobuf_oneof:"lb_policy"`
}

func (m *LoadBalancing) Reset()                    { *m = LoadBalancing{} }
func (m *LoadBalancing) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancing) ProtoMessage()               {}
func (*LoadBalancing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isLoadBalancing_LbPolicy interface {
	isLoadBalancing_LbPolicy()
}

type LoadBalancing_Name struct {
	Name LoadBalancing_SimpleLBPolicy `protobuf:"varint,1,opt,name=name,enum=istio.routing.v1alpha1.LoadBalancing_SimpleLBPolicy,oneof"`
}
type LoadBalancing_Custom struct {
	Custom *google_protobuf.Any `protobuf:"bytes,2,opt,name=custom,oneof"`
}

func (*LoadBalancing_Name) isLoadBalancing_LbPolicy()   {}
func (*LoadBalancing_Custom) isLoadBalancing_LbPolicy() {}

func (m *LoadBalancing) GetLbPolicy() isLoadBalancing_LbPolicy {
	if m != nil {
		return m.LbPolicy
	}
	return nil
}

func (m *LoadBalancing) GetName() LoadBalancing_SimpleLBPolicy {
	if x, ok := m.GetLbPolicy().(*LoadBalancing_Name); ok {
		return x.Name
	}
	return LoadBalancing_ROUND_ROBIN
}

func (m *LoadBalancing) GetCustom() *google_protobuf.Any {
	if x, ok := m.GetLbPolicy().(*LoadBalancing_Custom); ok {
		return x.Custom
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LoadBalancing) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LoadBalancing_OneofMarshaler, _LoadBalancing_OneofUnmarshaler, _LoadBalancing_OneofSizer, []interface{}{
		(*LoadBalancing_Name)(nil),
		(*LoadBalancing_Custom)(nil),
	}
}

func _LoadBalancing_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LoadBalancing)
	// lb_policy
	switch x := m.LbPolicy.(type) {
	case *LoadBalancing_Name:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Name))
	case *LoadBalancing_Custom:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Custom); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LoadBalancing.LbPolicy has unexpected type %T", x)
	}
	return nil
}

func _LoadBalancing_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LoadBalancing)
	switch tag {
	case 1: // lb_policy.name
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.LbPolicy = &LoadBalancing_Name{LoadBalancing_SimpleLBPolicy(x)}
		return true, err
	case 2: // lb_policy.custom
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf.Any)
		err := b.DecodeMessage(msg)
		m.LbPolicy = &LoadBalancing_Custom{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LoadBalancing_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LoadBalancing)
	// lb_policy
	switch x := m.LbPolicy.(type) {
	case *LoadBalancing_Name:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Name))
	case *LoadBalancing_Custom:
		s := proto.Size(x.Custom)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Circuit breaker configuration for Envoy. The circuit breaker
// implementation is fine-grained in that it tracks the success/failure
// rates of individual hosts in the load balancing pool. Hosts that
// continually return errors for API calls are ejected from the pool for a
// pre-defined period of time.
// See Envoy's
// [circuit breaker](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/circuit_breaking)
// and [outlier detection](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/outlier)
// for more details.
type CircuitBreaker struct {
	// Types that are valid to be assigned to CbPolicy:
	//	*CircuitBreaker_SimpleCb
	//	*CircuitBreaker_Custom
	CbPolicy isCircuitBreaker_CbPolicy `protobuf_oneof:"cb_policy"`
}

func (m *CircuitBreaker) Reset()                    { *m = CircuitBreaker{} }
func (m *CircuitBreaker) String() string            { return proto.CompactTextString(m) }
func (*CircuitBreaker) ProtoMessage()               {}
func (*CircuitBreaker) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isCircuitBreaker_CbPolicy interface {
	isCircuitBreaker_CbPolicy()
}

type CircuitBreaker_SimpleCb struct {
	SimpleCb *CircuitBreaker_SimpleCircuitBreakerPolicy `protobuf:"bytes,1,opt,name=simple_cb,json=simpleCb,oneof"`
}
type CircuitBreaker_Custom struct {
	Custom *google_protobuf.Any `protobuf:"bytes,2,opt,name=custom,oneof"`
}

func (*CircuitBreaker_SimpleCb) isCircuitBreaker_CbPolicy() {}
func (*CircuitBreaker_Custom) isCircuitBreaker_CbPolicy()   {}

func (m *CircuitBreaker) GetCbPolicy() isCircuitBreaker_CbPolicy {
	if m != nil {
		return m.CbPolicy
	}
	return nil
}

func (m *CircuitBreaker) GetSimpleCb() *CircuitBreaker_SimpleCircuitBreakerPolicy {
	if x, ok := m.GetCbPolicy().(*CircuitBreaker_SimpleCb); ok {
		return x.SimpleCb
	}
	return nil
}

func (m *CircuitBreaker) GetCustom() *google_protobuf.Any {
	if x, ok := m.GetCbPolicy().(*CircuitBreaker_Custom); ok {
		return x.Custom
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CircuitBreaker) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CircuitBreaker_OneofMarshaler, _CircuitBreaker_OneofUnmarshaler, _CircuitBreaker_OneofSizer, []interface{}{
		(*CircuitBreaker_SimpleCb)(nil),
		(*CircuitBreaker_Custom)(nil),
	}
}

func _CircuitBreaker_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CircuitBreaker)
	// cb_policy
	switch x := m.CbPolicy.(type) {
	case *CircuitBreaker_SimpleCb:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SimpleCb); err != nil {
			return err
		}
	case *CircuitBreaker_Custom:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Custom); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CircuitBreaker.CbPolicy has unexpected type %T", x)
	}
	return nil
}

func _CircuitBreaker_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CircuitBreaker)
	switch tag {
	case 1: // cb_policy.simple_cb
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CircuitBreaker_SimpleCircuitBreakerPolicy)
		err := b.DecodeMessage(msg)
		m.CbPolicy = &CircuitBreaker_SimpleCb{msg}
		return true, err
	case 2: // cb_policy.custom
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf.Any)
		err := b.DecodeMessage(msg)
		m.CbPolicy = &CircuitBreaker_Custom{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CircuitBreaker_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CircuitBreaker)
	// cb_policy
	switch x := m.CbPolicy.(type) {
	case *CircuitBreaker_SimpleCb:
		s := proto.Size(x.SimpleCb)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CircuitBreaker_Custom:
		s := proto.Size(x.Custom)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A simple circuit breaker can be set based on a number of criteria such as
// connection and request limits. For example, the following destination
// policy sets a limit of 100 connections to "reviews" service version
// "v1" backends.
//
//     metadata:
//       name: reviews-cb-policy
//       namespace: default
//     spec:
//       destination:
//         name: reviews
//         labels:
//           version: v1
//       circuitBreaker:
//         simpleCb:
//           maxConnections: 100
//
// The following destination policy sets a limit of 100 connections and
// 1000 concurrent requests, with no more than 10 req/connection to
// "reviews" service version "v1" backends. In addition, it configures
// hosts to be scanned every 5 mins, such that any host that fails 7
// consecutive times with 5XX error code will be ejected for 15 minutes.
//
//     metadata:
//       name: reviews-cb-policy
//       namespace: default
//     spec:
//       destination:
//         name: reviews
//         labels:
//           version: v1
//       circuitBreaker:
//         simpleCb:
//           maxConnections: 100
//           httpMaxRequests: 1000
//           httpMaxRequestsPerConnection: 10
//           httpConsecutiveErrors: 7
//           sleepWindow: 15m
//           httpDetectionInterval: 5m
//
type CircuitBreaker_SimpleCircuitBreakerPolicy struct {
	// Maximum number of connections to a backend.
	MaxConnections int32 `protobuf:"varint,1,opt,name=max_connections,json=maxConnections" json:"max_connections,omitempty"`
	// Maximum number of pending requests to a backend. Default 1024
	HttpMaxPendingRequests int32 `protobuf:"varint,2,opt,name=http_max_pending_requests,json=httpMaxPendingRequests" json:"http_max_pending_requests,omitempty"`
	// Maximum number of requests to a backend. Default 1024
	HttpMaxRequests int32 `protobuf:"varint,3,opt,name=http_max_requests,json=httpMaxRequests" json:"http_max_requests,omitempty"`
	// Minimum time the circuit will be closed. format: 1h/1m/1s/1ms. MUST
	// BE >=1ms. Default is 30s.
	SleepWindow *google_protobuf1.Duration `protobuf:"bytes,4,opt,name=sleep_window,json=sleepWindow" json:"sleep_window,omitempty"`
	// Number of 5XX errors before circuit is opened. Defaults to 5.
	HttpConsecutiveErrors int32 `protobuf:"varint,5,opt,name=http_consecutive_errors,json=httpConsecutiveErrors" json:"http_consecutive_errors,omitempty"`
	// Time interval between ejection sweep analysis. format:
	// 1h/1m/1s/1ms. MUST BE >=1ms. Default is 10s.
	HttpDetectionInterval *google_protobuf1.Duration `protobuf:"bytes,6,opt,name=http_detection_interval,json=httpDetectionInterval" json:"http_detection_interval,omitempty"`
	// Maximum number of requests per connection to a backend. Setting this
	// parameter to 1 disables keep alive.
	HttpMaxRequestsPerConnection int32 `protobuf:"varint,7,opt,name=http_max_requests_per_connection,json=httpMaxRequestsPerConnection" json:"http_max_requests_per_connection,omitempty"`
	// Maximum % of hosts in the load balancing pool for the destination
	// service that can be ejected by the circuit breaker. Defaults to
	// 10%.
	HttpMaxEjectionPercent int32 `protobuf:"varint,8,opt,name=http_max_ejection_percent,json=httpMaxEjectionPercent" json:"http_max_ejection_percent,omitempty"`
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) Reset() {
	*m = CircuitBreaker_SimpleCircuitBreakerPolicy{}
}
func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) String() string { return proto.CompactTextString(m) }
func (*CircuitBreaker_SimpleCircuitBreakerPolicy) ProtoMessage()    {}
func (*CircuitBreaker_SimpleCircuitBreakerPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 0}
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetMaxConnections() int32 {
	if m != nil {
		return m.MaxConnections
	}
	return 0
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetHttpMaxPendingRequests() int32 {
	if m != nil {
		return m.HttpMaxPendingRequests
	}
	return 0
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetHttpMaxRequests() int32 {
	if m != nil {
		return m.HttpMaxRequests
	}
	return 0
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetSleepWindow() *google_protobuf1.Duration {
	if m != nil {
		return m.SleepWindow
	}
	return nil
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetHttpConsecutiveErrors() int32 {
	if m != nil {
		return m.HttpConsecutiveErrors
	}
	return 0
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetHttpDetectionInterval() *google_protobuf1.Duration {
	if m != nil {
		return m.HttpDetectionInterval
	}
	return nil
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetHttpMaxRequestsPerConnection() int32 {
	if m != nil {
		return m.HttpMaxRequestsPerConnection
	}
	return 0
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetHttpMaxEjectionPercent() int32 {
	if m != nil {
		return m.HttpMaxEjectionPercent
	}
	return 0
}

func init() {
	proto.RegisterType((*DestinationPolicy)(nil), "istio.routing.v1alpha1.DestinationPolicy")
	proto.RegisterType((*LoadBalancing)(nil), "istio.routing.v1alpha1.LoadBalancing")
	proto.RegisterType((*CircuitBreaker)(nil), "istio.routing.v1alpha1.CircuitBreaker")
	proto.RegisterType((*CircuitBreaker_SimpleCircuitBreakerPolicy)(nil), "istio.routing.v1alpha1.CircuitBreaker.SimpleCircuitBreakerPolicy")
	proto.RegisterEnum("istio.routing.v1alpha1.LoadBalancing_SimpleLBPolicy", LoadBalancing_SimpleLBPolicy_name, LoadBalancing_SimpleLBPolicy_value)
}

func init() { proto.RegisterFile("routing/v1alpha1/dest_policy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xed, 0x4e, 0xdb, 0x3a,
	0x18, 0xa6, 0x85, 0xf6, 0x80, 0x7b, 0x68, 0xc1, 0x3a, 0x87, 0x13, 0xaa, 0xa3, 0x89, 0x55, 0xfb,
	0xd2, 0x34, 0xa5, 0x82, 0x4d, 0x93, 0x26, 0xb1, 0x1f, 0xfd, 0x00, 0xc1, 0x54, 0xda, 0x2e, 0x6c,
	0xda, 0x4f, 0xcf, 0x71, 0xde, 0x15, 0x6f, 0xa9, 0x9d, 0x39, 0x4e, 0x81, 0x5b, 0xda, 0x0d, 0xed,
	0x0e, 0x76, 0x0d, 0xfb, 0x39, 0xc5, 0x71, 0x5b, 0x02, 0xab, 0x06, 0x3f, 0xe3, 0xe7, 0xc3, 0xcf,
	0xfb, 0xf8, 0x0d, 0x6a, 0x28, 0x99, 0x68, 0x2e, 0x46, 0xcd, 0xc9, 0x2e, 0x0d, 0xa3, 0x33, 0xba,
	0xdb, 0x0c, 0x20, 0xd6, 0x24, 0x92, 0x21, 0x67, 0x97, 0x6e, 0xa4, 0xa4, 0x96, 0x78, 0x8b, 0xc7,
	0x9a, 0x4b, 0xd7, 0x32, 0xdd, 0x29, 0xb3, 0xbe, 0x3d, 0x92, 0x72, 0x14, 0x42, 0xd3, 0xb0, 0xfc,
	0xe4, 0x53, 0x93, 0x0a, 0x2b, 0xa9, 0xdf, 0xbb, 0x0e, 0x05, 0x89, 0xa2, 0x9a, 0x4b, 0x61, 0xf1,
	0xfb, 0x37, 0xae, 0x4d, 0x0f, 0x80, 0xa8, 0x24, 0x84, 0x8c, 0xd2, 0xf8, 0x51, 0x44, 0x9b, 0x5d,
	0x88, 0x35, 0x17, 0x46, 0x38, 0x34, 0x89, 0xf0, 0x21, 0xaa, 0x04, 0xf3, 0x43, 0xa7, 0xb0, 0x53,
	0x78, 0x52, 0xd9, 0x7b, 0xe0, 0xfe, 0x3e, 0xa1, 0x7b, 0x9c, 0x1e, 0x9f, 0x82, 0x9a, 0x70, 0x06,
	0xde, 0x55, 0x21, 0xde, 0x47, 0xe5, 0x58, 0x26, 0x8a, 0x81, 0x53, 0xbc, 0x83, 0x85, 0xd5, 0xe0,
	0x1e, 0xaa, 0x86, 0x92, 0x06, 0xc4, 0xa7, 0x21, 0x15, 0x8c, 0x8b, 0x91, 0xb3, 0x6c, 0x5c, 0x1e,
	0x2e, 0x72, 0xe9, 0x49, 0x1a, 0xb4, 0xa7, 0x64, 0x6f, 0x3d, 0xbc, 0xfa, 0x89, 0x07, 0xa8, 0xc6,
	0xb8, 0x62, 0x09, 0xd7, 0xc4, 0x57, 0x40, 0xbf, 0x80, 0x72, 0x56, 0x8c, 0xdd, 0xa3, 0x45, 0x76,
	0x9d, 0x8c, 0xde, 0xce, 0xd8, 0x5e, 0x95, 0xe5, 0xbe, 0xf1, 0x33, 0x54, 0x66, 0x49, 0xac, 0xe5,
	0xd8, 0x09, 0x8c, 0xcf, 0x3f, 0x6e, 0xf6, 0x1c, 0xee, 0xf4, 0x39, 0xdc, 0x96, 0xb8, 0xf4, 0x2c,
	0xa7, 0xf1, 0xbd, 0x80, 0xd6, 0x73, 0xf9, 0xf0, 0x1b, 0xb4, 0x22, 0xe8, 0x18, 0x4c, 0xbb, 0xd5,
	0xbd, 0x17, 0xb7, 0x1a, 0xca, 0x3d, 0xe5, 0xe3, 0x28, 0x84, 0x5e, 0x3b, 0x7b, 0xa8, 0xa3, 0x25,
	0xcf, 0x78, 0x60, 0x77, 0x96, 0xa5, 0xb8, 0x38, 0xcb, 0xd1, 0xd2, 0x2c, 0xcd, 0x6b, 0x54, 0xcd,
	0x3b, 0xe1, 0x1a, 0xaa, 0x78, 0x83, 0xf7, 0xfd, 0x2e, 0xf1, 0x06, 0xed, 0xe3, 0xfe, 0xc6, 0x12,
	0xae, 0x22, 0xd4, 0x3b, 0x68, 0x9d, 0xbe, 0x23, 0x9d, 0x41, 0xbf, 0xbf, 0x51, 0xc0, 0x08, 0x95,
	0xbd, 0x56, 0xbf, 0x3b, 0x38, 0xd9, 0x28, 0xb6, 0x2b, 0x68, 0x2d, 0xf4, 0xed, 0xfa, 0x36, 0xbe,
	0x95, 0x50, 0x35, 0x5f, 0x15, 0xfe, 0x88, 0xd6, 0x62, 0x63, 0x4f, 0x98, 0x6f, 0xb7, 0xa7, 0x75,
	0xbb, 0x96, 0xed, 0x80, 0xf9, 0xc3, 0xd9, 0xb0, 0xab, 0x99, 0x6b, 0xc7, 0xbf, 0xeb, 0xc0, 0xf5,
	0x9f, 0xcb, 0xa8, 0xbe, 0xd8, 0x1a, 0x3f, 0x46, 0xb5, 0x31, 0xbd, 0x20, 0x4c, 0x0a, 0x01, 0x2c,
	0x5d, 0xdd, 0xd8, 0xc4, 0x2e, 0x79, 0xd5, 0x31, 0xbd, 0xe8, 0xcc, 0x4f, 0xf1, 0x2b, 0xb4, 0x7d,
	0xa6, 0x75, 0x44, 0x52, 0x76, 0x04, 0x22, 0xe0, 0x62, 0x44, 0x14, 0x7c, 0x4d, 0x20, 0xd6, 0xb1,
	0x89, 0x52, 0xf2, 0xb6, 0x52, 0xc2, 0x09, 0xbd, 0x18, 0x66, 0xb0, 0x67, 0x51, 0xfc, 0x14, 0x6d,
	0xce, 0xa4, 0x33, 0xc9, 0xb2, 0x91, 0xd4, 0xac, 0x64, 0xc6, 0xdd, 0x47, 0x7f, 0xc7, 0x21, 0x40,
	0x44, 0xce, 0xb9, 0x08, 0xe4, 0xb9, 0xdd, 0xd4, 0xed, 0x1b, 0x43, 0x76, 0xed, 0x0f, 0xef, 0x55,
	0x0c, 0xfd, 0x83, 0x61, 0xe3, 0x97, 0xe8, 0x3f, 0x73, 0x13, 0x93, 0x22, 0x06, 0x96, 0x68, 0x3e,
	0x01, 0x02, 0x4a, 0x49, 0x15, 0x3b, 0x25, 0x73, 0xdf, 0xbf, 0x29, 0xdc, 0x99, 0xa3, 0x07, 0x06,
	0xc4, 0x6f, 0xad, 0x2e, 0x00, 0x9d, 0xcd, 0x4b, 0xb8, 0xd0, 0xa0, 0x26, 0x34, 0x74, 0xca, 0x7f,
	0x0a, 0x60, 0x2c, 0xbb, 0x53, 0xe1, 0xb1, 0xd5, 0xe1, 0x43, 0xb4, 0x73, 0x63, 0x68, 0x12, 0x81,
	0xba, 0x52, 0xb5, 0xf3, 0x97, 0xc9, 0xf4, 0xff, 0xb5, 0x0e, 0x86, 0xa0, 0xe6, 0xc5, 0xe7, 0x7a,
	0x87, 0xcf, 0x36, 0x5d, 0x04, 0x8a, 0x81, 0xd0, 0xce, 0x6a, 0xae, 0xf7, 0x03, 0x0b, 0x0f, 0x33,
	0x34, 0x5d, 0x56, 0x36, 0x5d, 0x56, 0xbf, 0x6c, 0x92, 0x3f, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff,
	0x45, 0x80, 0xbc, 0xa9, 0x92, 0x05, 0x00, 0x00,
}
