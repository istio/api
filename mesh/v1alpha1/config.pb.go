// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mesh/v1alpha1/config.proto

/*
Package v1alpha1 is a generated protocol buffer package.

It is generated from these files:
	mesh/v1alpha1/config.proto

It has these top-level messages:
	ServerAddress
	ProxyConfig
	MeshConfig
*/
package v1alpha1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// AuthenticationPolicy defines authentication policy. It can be set for
// different scopes (mesh, service …), and the most narrow scope with
// non-INHERIT value will be used.
// Mesh policy cannot be INHERIT.
type AuthenticationPolicy int32

const (
	// Do not encrypt Envoy to Envoy traffic.
	AuthenticationPolicy_NONE AuthenticationPolicy = 0
	// Envoy to Envoy traffic is wrapped into mutual TLS connections.
	AuthenticationPolicy_MUTUAL_TLS AuthenticationPolicy = 1
	// Use the policy defined by the parent scope. Should not be used for mesh
	// policy.
	AuthenticationPolicy_INHERIT AuthenticationPolicy = 1000
)

var AuthenticationPolicy_name = map[int32]string{
	0:    "NONE",
	1:    "MUTUAL_TLS",
	1000: "INHERIT",
}
var AuthenticationPolicy_value = map[string]int32{
	"NONE":       0,
	"MUTUAL_TLS": 1,
	"INHERIT":    1000,
}

func (x AuthenticationPolicy) String() string {
	return proto.EnumName(AuthenticationPolicy_name, int32(x))
}
func (AuthenticationPolicy) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type MeshConfig_IngressControllerMode int32

const (
	// Disables Istio ingress controller.
	MeshConfig_OFF MeshConfig_IngressControllerMode = 0
	// Istio ingress controller will act on ingress resources that do not
	// contain any annotation or whose annotations match the value
	// specified in the ingress_class parameter described earlier. Use this
	// mode if Istio ingress controller will be the default ingress
	// controller for the entire kubernetes cluster.
	MeshConfig_DEFAULT MeshConfig_IngressControllerMode = 1
	// Istio ingress controller will only act on ingress resources whose
	// annotations match the value specified in the ingress_class parameter
	// described earlier. Use this mode if Istio ingress controller will be
	// a secondary ingress controller (e.g., in addition to a
	// cloud-provided ingress controller).
	MeshConfig_STRICT MeshConfig_IngressControllerMode = 2
)

var MeshConfig_IngressControllerMode_name = map[int32]string{
	0: "OFF",
	1: "DEFAULT",
	2: "STRICT",
}
var MeshConfig_IngressControllerMode_value = map[string]int32{
	"OFF":     0,
	"DEFAULT": 1,
	"STRICT":  2,
}

func (x MeshConfig_IngressControllerMode) String() string {
	return proto.EnumName(MeshConfig_IngressControllerMode_name, int32(x))
}
func (MeshConfig_IngressControllerMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 0}
}

// TODO AuthPolicy needs to be removed and merged with AuthPolicy defined above
type MeshConfig_AuthPolicy int32

const (
	// Do not encrypt Envoy to Envoy traffic.
	MeshConfig_NONE MeshConfig_AuthPolicy = 0
	// Envoy to Envoy traffic is wrapped into mutual TLS connections.
	MeshConfig_MUTUAL_TLS MeshConfig_AuthPolicy = 1
)

var MeshConfig_AuthPolicy_name = map[int32]string{
	0: "NONE",
	1: "MUTUAL_TLS",
}
var MeshConfig_AuthPolicy_value = map[string]int32{
	"NONE":       0,
	"MUTUAL_TLS": 1,
}

func (x MeshConfig_AuthPolicy) String() string {
	return proto.EnumName(MeshConfig_AuthPolicy_name, int32(x))
}
func (MeshConfig_AuthPolicy) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 1} }

// ServerAddress specifies the address of Istio components like mixer, pilot, etc.
// At least one of the field needs to be specified.
type ServerAddress struct {
	// The address for mTLS server, e.g., (_istio-pilot:15003_)
	MutualTls string `protobuf:"bytes,1,opt,name=mutual_tls,json=mutualTls" json:"mutual_tls,omitempty"`
	// The address for plain text server, e.g., (_istio-pilot:15005_)
	PlainText string `protobuf:"bytes,2,opt,name=plain_text,json=plainText" json:"plain_text,omitempty"`
}

func (m *ServerAddress) Reset()                    { *m = ServerAddress{} }
func (m *ServerAddress) String() string            { return proto.CompactTextString(m) }
func (*ServerAddress) ProtoMessage()               {}
func (*ServerAddress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServerAddress) GetMutualTls() string {
	if m != nil {
		return m.MutualTls
	}
	return ""
}

func (m *ServerAddress) GetPlainText() string {
	if m != nil {
		return m.PlainText
	}
	return ""
}

// ProxyConfig defines variables for individual Envoy instances.
type ProxyConfig struct {
	// Path to the generated configuration file directory.
	// Proxy agent generates the actual configuration and stores it in this directory.
	ConfigPath string `protobuf:"bytes,1,opt,name=config_path,json=configPath" json:"config_path,omitempty"`
	// Path to the proxy binary
	BinaryPath string `protobuf:"bytes,2,opt,name=binary_path,json=binaryPath" json:"binary_path,omitempty"`
	// Path to template that is used to generate proxy config
	ProxyTemplatePath string `protobuf:"bytes,3,opt,name=proxy_template_path,json=proxyTemplatePath" json:"proxy_template_path,omitempty"`
	// Service cluster defines the name for the service_cluster that is
	// shared by all Envoy instances. This setting corresponds to
	// _--service-cluster_ flag in Envoy.  In a typical Envoy deployment, the
	// _service-cluster_ flag is used to identify the caller, for
	// source-based routing scenarios.
	//
	// Since Istio does not assign a local service/service version to each
	// Envoy instance, the name is same for all of them.  However, the
	// source/caller's identity (e.g., IP address) is encoded in the
	// _--service-node_ flag when launching Envoy.  When the RDS service
	// receives API calls from Envoy, it uses the value of the _service-node_
	// flag to compute routes that are relative to the service instances
	// located at that IP address.
	ServiceCluster string `protobuf:"bytes,3,opt,name=service_cluster,json=serviceCluster" json:"service_cluster,omitempty"`
	// The time in seconds that Envoy will drain connections during a hot
	// restart. MUST be >=1s (e.g., _1s/1m/1h_)
	DrainDuration *google_protobuf.Duration `protobuf:"bytes,4,opt,name=drain_duration,json=drainDuration" json:"drain_duration,omitempty"`
	// The time in seconds that Envoy will wait before shutting down the
	// parent process during a hot restart. MUST be >=1s (e.g., _1s/1m/1h_).
	// MUST BE greater than _drain_duration_ parameter.
	ParentShutdownDuration *google_protobuf.Duration `protobuf:"bytes,5,opt,name=parent_shutdown_duration,json=parentShutdownDuration" json:"parent_shutdown_duration,omitempty"`
	// Deprecated, use server_address instead.
	DiscoveryAddress string `protobuf:"bytes,6,opt,name=discovery_address,json=discoveryAddress" json:"discovery_address,omitempty"`
	// Polling interval for service discovery (used by EDS, CDS, LDS, but not RDS). (MUST BE >=1ms)
	DiscoveryRefreshDelay *google_protobuf.Duration `protobuf:"bytes,7,opt,name=discovery_refresh_delay,json=discoveryRefreshDelay" json:"discovery_refresh_delay,omitempty"`
	// Address of the Zipkin service (e.g. _zipkin:9411_).
	ZipkinAddress string `protobuf:"bytes,8,opt,name=zipkin_address,json=zipkinAddress" json:"zipkin_address,omitempty"`
	// Connection timeout used by Envoy for supporting services. (MUST BE >=1ms)
	ConnectTimeout *google_protobuf.Duration `protobuf:"bytes,9,opt,name=connect_timeout,json=connectTimeout" json:"connect_timeout,omitempty"`
	// IP Address and Port of a statsd UDP listener (e.g. _10.75.241.127:9125_).
	StatsdUdpAddress string `protobuf:"bytes,10,opt,name=statsd_udp_address,json=statsdUdpAddress" json:"statsd_udp_address,omitempty"`
	// Port on which Envoy should listen for administrative commands.
	ProxyAdminPort int32 `protobuf:"varint,11,opt,name=proxy_admin_port,json=proxyAdminPort" json:"proxy_admin_port,omitempty"`
	// The availability zone where this Envoy instance is running. When running
	// Envoy as a sidecar in Kubernetes, this flag must be one of the availability
	// zones assigned to a node using failure-domain.beta.kubernetes.io/zone annotation.
	AvailabilityZone string `protobuf:"bytes,12,opt,name=availability_zone,json=availabilityZone" json:"availability_zone,omitempty"`
	// Authentication policy defines the global switch to control authentication
	// for Envoy-to-Envoy communication for istio components Mixer and Pilot.
	ControlPlaneAuthPolicy AuthenticationPolicy `protobuf:"varint,13,opt,name=control_plane_auth_policy,json=controlPlaneAuthPolicy,enum=istio.mesh.v1alpha1.AuthenticationPolicy" json:"control_plane_auth_policy,omitempty"`
	// File path of custom proxy configuration, currently used by proxies
	// in front of Mixer and Pilot.
	CustomConfigFile string `protobuf:"bytes,14,opt,name=custom_config_file,json=customConfigFile" json:"custom_config_file,omitempty"`
	// Maximum length of name field in Envoy's metrics. The length of the name field
	// is determined by the length of a name field in a service and the set of labels that
	// comprise a particular version of the service. The default value is set to 189 characters.
	// Envoy's internal metrics take up 67 characters, for a total of 256 character name per metric.
	// Increase the value of this field if you find that the metrics from Envoys are truncated.
	StatNameLength int32 `protobuf:"varint,15,opt,name=stat_name_length,json=statNameLength" json:"stat_name_length,omitempty"`
	// The number of worker threads to run. Default value is number of cores on the machine.
	Concurrency int32 `protobuf:"varint,16,opt,name=concurrency" json:"concurrency,omitempty"`
	// Deprecated, use discovery_mtls_address/discovery_plain_address instead.
	Pilot *ServerAddress `protobuf:"bytes,17,opt,name=pilot" json:"pilot,omitempty"`
	// Address of the discovery service exposing xDS with mTLS connection.
	DiscoveryMtlsAddress string `protobuf:"bytes,18,opt,name=discovery_mtls_address,json=discoveryMtlsAddress" json:"discovery_mtls_address,omitempty"`
	// Address of the discovery service exposing xDS with plain text connection.
	DiscoveryPlainAddress string `protobuf:"bytes,19,opt,name=discovery_plain_address,json=discoveryPlainAddress" json:"discovery_plain_address,omitempty"`
}

func (m *ProxyConfig) Reset()                    { *m = ProxyConfig{} }
func (m *ProxyConfig) String() string            { return proto.CompactTextString(m) }
func (*ProxyConfig) ProtoMessage()               {}
func (*ProxyConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ProxyConfig) GetConfigPath() string {
	if m != nil {
		return m.ConfigPath
	}
	return ""
}

func (m *ProxyConfig) GetBinaryPath() string {
	if m != nil {
		return m.BinaryPath
	}
	return ""
}

func (m *ProxyConfig) GetProxyTemplatePath() string {
        if m != nil {
                return m.proxyTemplatePath
        }
        return ""
}

func (m *ProxyConfig) GetServiceCluster() string {
	if m != nil {
		return m.ServiceCluster
	}
	return ""
}

func (m *ProxyConfig) GetDrainDuration() *google_protobuf.Duration {
	if m != nil {
		return m.DrainDuration
	}
	return nil
}

func (m *ProxyConfig) GetParentShutdownDuration() *google_protobuf.Duration {
	if m != nil {
		return m.ParentShutdownDuration
	}
	return nil
}

func (m *ProxyConfig) GetDiscoveryAddress() string {
	if m != nil {
		return m.DiscoveryAddress
	}
	return ""
}

func (m *ProxyConfig) GetDiscoveryRefreshDelay() *google_protobuf.Duration {
	if m != nil {
		return m.DiscoveryRefreshDelay
	}
	return nil
}

func (m *ProxyConfig) GetZipkinAddress() string {
	if m != nil {
		return m.ZipkinAddress
	}
	return ""
}

func (m *ProxyConfig) GetConnectTimeout() *google_protobuf.Duration {
	if m != nil {
		return m.ConnectTimeout
	}
	return nil
}

func (m *ProxyConfig) GetStatsdUdpAddress() string {
	if m != nil {
		return m.StatsdUdpAddress
	}
	return ""
}

func (m *ProxyConfig) GetProxyAdminPort() int32 {
	if m != nil {
		return m.ProxyAdminPort
	}
	return 0
}

func (m *ProxyConfig) GetAvailabilityZone() string {
	if m != nil {
		return m.AvailabilityZone
	}
	return ""
}

func (m *ProxyConfig) GetControlPlaneAuthPolicy() AuthenticationPolicy {
	if m != nil {
		return m.ControlPlaneAuthPolicy
	}
	return AuthenticationPolicy_NONE
}

func (m *ProxyConfig) GetCustomConfigFile() string {
	if m != nil {
		return m.CustomConfigFile
	}
	return ""
}

func (m *ProxyConfig) GetStatNameLength() int32 {
	if m != nil {
		return m.StatNameLength
	}
	return 0
}

func (m *ProxyConfig) GetConcurrency() int32 {
	if m != nil {
		return m.Concurrency
	}
	return 0
}

func (m *ProxyConfig) GetPilot() *ServerAddress {
	if m != nil {
		return m.Pilot
	}
	return nil
}

func (m *ProxyConfig) GetDiscoveryMtlsAddress() string {
	if m != nil {
		return m.DiscoveryMtlsAddress
	}
	return ""
}

func (m *ProxyConfig) GetDiscoveryPlainAddress() string {
	if m != nil {
		return m.DiscoveryPlainAddress
	}
	return ""
}

// MeshConfig defines mesh-wide variables shared by all Envoy instances in the
// Istio service mesh.
type MeshConfig struct {
	// Deprecated, use mixer_check instead.
	MixerCheckServer string `protobuf:"bytes,1,opt,name=mixer_check_server,json=mixerCheckServer" json:"mixer_check_server,omitempty"`
	// Deprecated, use mixer_report instead.
	MixerReportServer string `protobuf:"bytes,2,opt,name=mixer_report_server,json=mixerReportServer" json:"mixer_report_server,omitempty"`
	// Disable policy checks by the mixer service. Default
	// is false, i.e. mixer policy check is enabled by default.
	DisablePolicyChecks bool `protobuf:"varint,3,opt,name=disable_policy_checks,json=disablePolicyChecks" json:"disable_policy_checks,omitempty"`
	// Port on which Envoy should listen for incoming connections from
	// other services.
	ProxyListenPort int32 `protobuf:"varint,4,opt,name=proxy_listen_port,json=proxyListenPort" json:"proxy_listen_port,omitempty"`
	// Port on which Envoy should listen for HTTP PROXY requests if set.
	ProxyHttpPort int32 `protobuf:"varint,5,opt,name=proxy_http_port,json=proxyHttpPort" json:"proxy_http_port,omitempty"`
	// Connection timeout used by Envoy. (MUST BE >=1ms)
	ConnectTimeout *google_protobuf.Duration `protobuf:"bytes,6,opt,name=connect_timeout,json=connectTimeout" json:"connect_timeout,omitempty"`
	// Class of ingress resources to be processed by Istio ingress
	// controller.  This corresponds to the value of
	// "kubernetes.io/ingress.class" annotation.
	IngressClass string `protobuf:"bytes,7,opt,name=ingress_class,json=ingressClass" json:"ingress_class,omitempty"`
	// Name of the kubernetes service used for the istio ingress controller.
	IngressService string `protobuf:"bytes,8,opt,name=ingress_service,json=ingressService" json:"ingress_service,omitempty"`
	// Defines whether to use Istio ingress controller for annotated or all ingress resources.
	IngressControllerMode MeshConfig_IngressControllerMode `protobuf:"varint,9,opt,name=ingress_controller_mode,json=ingressControllerMode,enum=istio.mesh.v1alpha1.MeshConfig_IngressControllerMode" json:"ingress_controller_mode,omitempty"`
	// Authentication policy defines the global switch to control authentication
	// for Envoy-to-Envoy communication.
	// Use authentication_policy instead.
	AuthPolicy MeshConfig_AuthPolicy `protobuf:"varint,10,opt,name=auth_policy,json=authPolicy,enum=istio.mesh.v1alpha1.MeshConfig_AuthPolicy" json:"auth_policy,omitempty"`
	// Polling interval for RDS (MUST BE >=1ms)
	RdsRefreshDelay *google_protobuf.Duration `protobuf:"bytes,11,opt,name=rds_refresh_delay,json=rdsRefreshDelay" json:"rds_refresh_delay,omitempty"`
	// Flag to control generation of trace spans and request IDs.
	// Requires a trace span collector defined in the proxy configuration.
	EnableTracing bool `protobuf:"varint,12,opt,name=enable_tracing,json=enableTracing" json:"enable_tracing,omitempty"`
	// File address for the proxy access log (e.g. /dev/stdout).
	// Empty value disables access logging.
	AccessLogFile string `protobuf:"bytes,13,opt,name=access_log_file,json=accessLogFile" json:"access_log_file,omitempty"`
	// Default proxy config used by the proxy injection mechanism operating in the mesh
	// (e.g. Kubernetes admission controller)
	// In case of Kubernetes, the proxy config is applied once during the injection process,
	// and remain constant for the duration of the pod. The rest of the mesh config can be changed
	// at runtime and config gets distributed dynamically.
	DefaultConfig *ProxyConfig `protobuf:"bytes,14,opt,name=default_config,json=defaultConfig" json:"default_config,omitempty"`
	// List of remote services for which mTLS authentication should not be expected by Istio .
	// Typically, these are control services (e.g kubernetes API server) that don't have Istio sidecar
	// to transparently terminate mTLS authentication.
	// It has no effect if the authentication policy is already 'NONE'.
	// DO NOT use this setting for services that are managed by Istio (i.e. using Istio sidecar).
	// Instead, use service-level annotations to overwrite the authentication policy.
	MtlsExcludedServices []string `protobuf:"bytes,15,rep,name=mtls_excluded_services,json=mtlsExcludedServices" json:"mtls_excluded_services,omitempty"`
	// DEPRECATED. Mixer address. This option will be removed soon. Please
	// use mixer_check and mixer_report.
	MixerAddress string `protobuf:"bytes,16,opt,name=mixer_address,json=mixerAddress" json:"mixer_address,omitempty"`
	// Address of the server that will be used by the proxies for policy
	// check calls. By using different names for mixerCheck and mixerReport, it
	// is possible to have one set of mixer servers handle policy check calls,
	// while another set of mixer servers handle telemetry calls.
	//
	// NOTE: Omitting mixerCheck while specifying mixerReport is
	// equivalent to setting disablePolicyChecks to true.
	MixerCheck *ServerAddress `protobuf:"bytes,17,opt,name=mixer_check,json=mixerCheck" json:"mixer_check,omitempty"`
	// Address of the server that will be used by the proxies for policy report
	// calls.
	MixerReport *ServerAddress `protobuf:"bytes,18,opt,name=mixer_report,json=mixerReport" json:"mixer_report,omitempty"`
}

func (m *MeshConfig) Reset()                    { *m = MeshConfig{} }
func (m *MeshConfig) String() string            { return proto.CompactTextString(m) }
func (*MeshConfig) ProtoMessage()               {}
func (*MeshConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MeshConfig) GetMixerCheckServer() string {
	if m != nil {
		return m.MixerCheckServer
	}
	return ""
}

func (m *MeshConfig) GetMixerReportServer() string {
	if m != nil {
		return m.MixerReportServer
	}
	return ""
}

func (m *MeshConfig) GetDisablePolicyChecks() bool {
	if m != nil {
		return m.DisablePolicyChecks
	}
	return false
}

func (m *MeshConfig) GetProxyListenPort() int32 {
	if m != nil {
		return m.ProxyListenPort
	}
	return 0
}

func (m *MeshConfig) GetProxyHttpPort() int32 {
	if m != nil {
		return m.ProxyHttpPort
	}
	return 0
}

func (m *MeshConfig) GetConnectTimeout() *google_protobuf.Duration {
	if m != nil {
		return m.ConnectTimeout
	}
	return nil
}

func (m *MeshConfig) GetIngressClass() string {
	if m != nil {
		return m.IngressClass
	}
	return ""
}

func (m *MeshConfig) GetIngressService() string {
	if m != nil {
		return m.IngressService
	}
	return ""
}

func (m *MeshConfig) GetIngressControllerMode() MeshConfig_IngressControllerMode {
	if m != nil {
		return m.IngressControllerMode
	}
	return MeshConfig_OFF
}

func (m *MeshConfig) GetAuthPolicy() MeshConfig_AuthPolicy {
	if m != nil {
		return m.AuthPolicy
	}
	return MeshConfig_NONE
}

func (m *MeshConfig) GetRdsRefreshDelay() *google_protobuf.Duration {
	if m != nil {
		return m.RdsRefreshDelay
	}
	return nil
}

func (m *MeshConfig) GetEnableTracing() bool {
	if m != nil {
		return m.EnableTracing
	}
	return false
}

func (m *MeshConfig) GetAccessLogFile() string {
	if m != nil {
		return m.AccessLogFile
	}
	return ""
}

func (m *MeshConfig) GetDefaultConfig() *ProxyConfig {
	if m != nil {
		return m.DefaultConfig
	}
	return nil
}

func (m *MeshConfig) GetMtlsExcludedServices() []string {
	if m != nil {
		return m.MtlsExcludedServices
	}
	return nil
}

func (m *MeshConfig) GetMixerAddress() string {
	if m != nil {
		return m.MixerAddress
	}
	return ""
}

func (m *MeshConfig) GetMixerCheck() *ServerAddress {
	if m != nil {
		return m.MixerCheck
	}
	return nil
}

func (m *MeshConfig) GetMixerReport() *ServerAddress {
	if m != nil {
		return m.MixerReport
	}
	return nil
}

func init() {
	proto.RegisterType((*ServerAddress)(nil), "istio.mesh.v1alpha1.ServerAddress")
	proto.RegisterType((*ProxyConfig)(nil), "istio.mesh.v1alpha1.ProxyConfig")
	proto.RegisterType((*MeshConfig)(nil), "istio.mesh.v1alpha1.MeshConfig")
	proto.RegisterEnum("istio.mesh.v1alpha1.AuthenticationPolicy", AuthenticationPolicy_name, AuthenticationPolicy_value)
	proto.RegisterEnum("istio.mesh.v1alpha1.MeshConfig_IngressControllerMode", MeshConfig_IngressControllerMode_name, MeshConfig_IngressControllerMode_value)
	proto.RegisterEnum("istio.mesh.v1alpha1.MeshConfig_AuthPolicy", MeshConfig_AuthPolicy_name, MeshConfig_AuthPolicy_value)
}

func init() { proto.RegisterFile("mesh/v1alpha1/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1078 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0x5b, 0x4f, 0x1b, 0x47,
	0x14, 0x8e, 0xb9, 0x9a, 0x63, 0x7c, 0x61, 0x08, 0x64, 0x83, 0x7a, 0xb1, 0x5c, 0x35, 0xa5, 0x34,
	0xb2, 0x15, 0x9a, 0x56, 0xed, 0x4b, 0x55, 0x30, 0xa6, 0x41, 0x32, 0xe0, 0xae, 0xcd, 0x4b, 0x5e,
	0x46, 0xc3, 0xee, 0xe0, 0x1d, 0x65, 0xbc, 0xb3, 0x9a, 0x99, 0xa5, 0x38, 0x7f, 0xa0, 0x7f, 0xb5,
	0x2f, 0xfd, 0x0f, 0xd5, 0x5c, 0xd6, 0x76, 0x22, 0x57, 0x24, 0x8f, 0xfe, 0xce, 0x77, 0xce, 0x99,
	0x73, 0xf9, 0x8e, 0x17, 0x0e, 0x26, 0x54, 0x25, 0x9d, 0xfb, 0x57, 0x84, 0x67, 0x09, 0x79, 0xd5,
	0x89, 0x44, 0x7a, 0xc7, 0xc6, 0xed, 0x4c, 0x0a, 0x2d, 0xd0, 0x2e, 0x53, 0x9a, 0x89, 0xb6, 0x61,
	0xb4, 0x0b, 0xc6, 0xc1, 0x57, 0x63, 0x21, 0xc6, 0x9c, 0x76, 0x2c, 0xe5, 0x36, 0xbf, 0xeb, 0xc4,
	0xb9, 0x24, 0x9a, 0x89, 0xd4, 0x39, 0xb5, 0x2e, 0xa1, 0x3a, 0xa4, 0xf2, 0x9e, 0xca, 0x93, 0x38,
	0x96, 0x54, 0x29, 0xf4, 0x25, 0xc0, 0x24, 0xd7, 0x39, 0xe1, 0x58, 0x73, 0x15, 0x94, 0x9a, 0xa5,
	0xc3, 0xad, 0x70, 0xcb, 0x21, 0x23, 0x6e, 0xcd, 0x19, 0x27, 0x2c, 0xc5, 0x9a, 0x3e, 0xe8, 0x60,
	0xc5, 0x99, 0x2d, 0x32, 0xa2, 0x0f, 0xba, 0xf5, 0x77, 0x19, 0x2a, 0x03, 0x29, 0x1e, 0xa6, 0x5d,
	0xfb, 0x32, 0xf4, 0x35, 0x54, 0xdc, 0x1b, 0x71, 0x46, 0x74, 0xe2, 0xc3, 0x81, 0x83, 0x06, 0x44,
	0x27, 0x86, 0x70, 0xcb, 0x52, 0x22, 0xa7, 0x8e, 0xe0, 0x02, 0x82, 0x83, 0x2c, 0xe1, 0x3b, 0xa8,
	0x2b, 0x2a, 0xef, 0x59, 0x44, 0x71, 0xc4, 0x73, 0xa5, 0xa9, 0x0c, 0x56, 0x2d, 0xa9, 0xe6, 0xe1,
	0xae, 0x43, 0xd1, 0xef, 0x50, 0x8b, 0xa5, 0x79, 0x59, 0x51, 0x61, 0xb0, 0xd6, 0x2c, 0x1d, 0x56,
	0x8e, 0x9f, 0xb7, 0x5d, 0x0b, 0xda, 0x45, 0x0b, 0xda, 0x67, 0x9e, 0x10, 0x56, 0xad, 0x43, 0xf1,
	0x13, 0x0d, 0x21, 0xc8, 0x88, 0xa4, 0xa9, 0xc6, 0x2a, 0xc9, 0x75, 0x2c, 0xfe, 0x5a, 0x88, 0xb5,
	0xfe, 0x58, 0xac, 0x7d, 0xe7, 0x3a, 0xf4, 0x9e, 0xb3, 0xa0, 0x3f, 0xc0, 0x4e, 0xcc, 0x54, 0x24,
	0xee, 0xa9, 0x9c, 0x62, 0xe2, 0x9a, 0x1c, 0x6c, 0xd8, 0x0a, 0x1a, 0x33, 0x43, 0xd1, 0xfc, 0x3f,
	0xe1, 0xd9, 0x9c, 0x2c, 0xe9, 0x9d, 0xa4, 0x2a, 0xc1, 0x31, 0xe5, 0x64, 0x1a, 0x6c, 0x3e, 0xf6,
	0x80, 0xbd, 0x99, 0x67, 0xe8, 0x1c, 0xcf, 0x8c, 0x1f, 0xfa, 0x16, 0x6a, 0xef, 0x59, 0xf6, 0x8e,
	0xa5, 0xb3, 0xe4, 0x65, 0x9b, 0xbc, 0xea, 0xd0, 0x22, 0xf3, 0x29, 0xd4, 0x23, 0x91, 0xa6, 0x34,
	0xd2, 0x58, 0xb3, 0x09, 0x15, 0xb9, 0x0e, 0xb6, 0x1e, 0xcb, 0x58, 0xf3, 0x1e, 0x23, 0xe7, 0x80,
	0x5e, 0x02, 0x52, 0x9a, 0x68, 0x15, 0xe3, 0x3c, 0xce, 0x66, 0xe9, 0xc0, 0xd5, 0xea, 0x2c, 0x37,
	0x71, 0x56, 0x64, 0x3c, 0x84, 0x46, 0x66, 0x36, 0x05, 0x93, 0x78, 0xc2, 0x52, 0x9c, 0x09, 0xa9,
	0x83, 0x4a, 0xb3, 0x74, 0xb8, 0x1e, 0xd6, 0x2c, 0x7e, 0x62, 0xe0, 0x81, 0x90, 0xda, 0xb4, 0x90,
	0xdc, 0x13, 0xc6, 0xc9, 0x2d, 0xe3, 0x4c, 0x4f, 0xf1, 0x7b, 0x91, 0xd2, 0x60, 0xdb, 0x85, 0x5d,
	0x34, 0xbc, 0x15, 0x29, 0x45, 0x31, 0x3c, 0x8f, 0x44, 0xaa, 0xa5, 0xe0, 0x38, 0xe3, 0x24, 0xa5,
	0x98, 0xe4, 0x3a, 0xc1, 0x99, 0xe0, 0x2c, 0x9a, 0x06, 0xd5, 0x66, 0xe9, 0xb0, 0x76, 0xfc, 0x7d,
	0x7b, 0x89, 0x52, 0xda, 0x27, 0xb9, 0x4e, 0x68, 0xaa, 0x59, 0x64, 0x8b, 0x1b, 0x58, 0x87, 0x70,
	0xdf, 0xc7, 0x1a, 0x98, 0x50, 0x86, 0xe1, 0x70, 0x53, 0x6a, 0x94, 0x2b, 0x2d, 0x26, 0xd8, 0xaf,
	0xf7, 0x1d, 0xe3, 0x34, 0xa8, 0xb9, 0x37, 0x39, 0x8b, 0x53, 0xc0, 0x39, 0xe3, 0xd4, 0x94, 0x6a,
	0xca, 0xc7, 0x29, 0x99, 0x50, 0xcc, 0x69, 0x3a, 0xd6, 0x49, 0x50, 0x77, 0xa5, 0x1a, 0xfc, 0x8a,
	0x4c, 0x68, 0xdf, 0xa2, 0xa8, 0x69, 0xf5, 0x12, 0xe5, 0x52, 0xd2, 0x34, 0x9a, 0x06, 0x0d, 0x4b,
	0x5a, 0x84, 0xd0, 0x2f, 0xb0, 0x9e, 0x31, 0x2e, 0x74, 0xb0, 0x63, 0xc7, 0xd3, 0x5a, 0x5a, 0xcb,
	0x07, 0x92, 0x0e, 0x9d, 0x03, 0x7a, 0x0d, 0xfb, 0xf3, 0xe5, 0x9a, 0x68, 0xae, 0x66, 0x23, 0x42,
	0xf6, 0xdd, 0x4f, 0x67, 0xd6, 0x4b, 0xcd, 0x55, 0x31, 0xa6, 0x9f, 0x17, 0x57, 0xd2, 0x49, 0xbf,
	0x70, 0xdb, 0xb5, 0x6e, 0xf3, 0xbd, 0x1b, 0x18, 0xab, 0xf7, 0x6b, 0xfd, 0x5b, 0x06, 0xb8, 0xa4,
	0x2a, 0xf1, 0x87, 0xe0, 0x25, 0xa0, 0x09, 0x7b, 0xa0, 0x12, 0x47, 0x09, 0x8d, 0xde, 0x61, 0x65,
	0x1f, 0xe8, 0xef, 0x41, 0xc3, 0x5a, 0xba, 0xc6, 0xe0, 0x1e, 0x8e, 0xda, 0xb0, 0xeb, 0xd8, 0x92,
	0x9a, 0xbd, 0x28, 0xe8, 0xee, 0x3a, 0xec, 0x58, 0x53, 0x68, 0x2d, 0x9e, 0x7f, 0x0c, 0xe6, 0x15,
	0xe4, 0x96, 0x53, 0x3f, 0x69, 0x97, 0x46, 0xd9, 0x53, 0x51, 0x0e, 0x77, 0xbd, 0xd1, 0x0d, 0xcf,
	0x26, 0x52, 0xe8, 0x08, 0x76, 0xdc, 0xfe, 0x71, 0xa6, 0x34, 0xf5, 0x0b, 0xb8, 0x66, 0x1b, 0x5e,
	0xb7, 0x86, 0xbe, 0xc5, 0xed, 0x06, 0xbe, 0x00, 0x07, 0xe1, 0x44, 0xeb, 0xcc, 0x31, 0xd7, 0x2d,
	0xb3, 0x6a, 0xe1, 0x37, 0x5a, 0x67, 0x96, 0xb7, 0x44, 0x45, 0x1b, 0x9f, 0xab, 0xa2, 0x6f, 0xa0,
	0xca, 0xd2, 0xb1, 0xe9, 0x21, 0x8e, 0x38, 0x51, 0xca, 0x2a, 0x7f, 0x2b, 0xdc, 0xf6, 0x60, 0xd7,
	0x60, 0xe6, 0x2a, 0x16, 0x24, 0x7f, 0x06, 0xbd, 0xac, 0x6b, 0x1e, 0x1e, 0x3a, 0x14, 0x4d, 0xe0,
	0xd9, 0x2c, 0x9a, 0x5b, 0x65, 0x4e, 0x25, 0x9e, 0x88, 0x98, 0x5a, 0x7d, 0xd7, 0x8e, 0x7f, 0x5a,
	0xba, 0x40, 0xf3, 0xc9, 0xb5, 0x2f, 0x7c, 0xde, 0x99, 0xf7, 0xa5, 0x88, 0x69, 0xb8, 0xc7, 0x96,
	0xc1, 0xe8, 0x1a, 0x2a, 0x8b, 0x7a, 0x03, 0x9b, 0xe2, 0xe8, 0xb1, 0x14, 0x73, 0x61, 0x9d, 0xae,
	0x04, 0xa5, 0x10, 0xc8, 0x5c, 0x68, 0x3d, 0xd8, 0x91, 0xb1, 0xfa, 0xe8, 0x16, 0x56, 0x1e, 0xeb,
	0x69, 0x5d, 0xc6, 0xea, 0xe3, 0x2b, 0x48, 0x53, 0xbb, 0x1f, 0x5a, 0x92, 0x88, 0xa5, 0x63, 0x7b,
	0x3f, 0xca, 0x61, 0xd5, 0xa1, 0x23, 0x07, 0x9a, 0x39, 0x93, 0x28, 0x32, 0xcd, 0xe2, 0xc2, 0x6b,
	0xba, 0xea, 0xae, 0xa5, 0x83, 0xfb, 0xc2, 0x09, 0xfa, 0x0f, 0xa8, 0xc5, 0xf4, 0x8e, 0xe4, 0x5c,
	0x7b, 0xfd, 0x5b, 0xe9, 0x57, 0x8e, 0x9b, 0x4b, 0x2b, 0x5d, 0xf8, 0x43, 0x0c, 0xab, 0xde, 0xcf,
	0xcb, 0xe2, 0x35, 0xec, 0x5b, 0x25, 0xd2, 0x87, 0x88, 0xe7, 0x31, 0x8d, 0x8b, 0x69, 0xaa, 0xa0,
	0xde, 0x5c, 0x35, 0x9a, 0x34, 0xd6, 0x9e, 0x37, 0xfa, 0x99, 0x2a, 0xb3, 0x22, 0x4e, 0x1e, 0x85,
	0x12, 0x1b, 0x6e, 0x45, 0x2c, 0x58, 0x08, 0xb7, 0x0b, 0x95, 0x05, 0xc5, 0x7d, 0xc6, 0xb9, 0x80,
	0xb9, 0x1c, 0x51, 0x0f, 0xb6, 0x17, 0x85, 0x68, 0x2f, 0xc5, 0xa7, 0x45, 0xa9, 0x2c, 0xa8, 0xb4,
	0xf5, 0x2b, 0xec, 0x2d, 0x5d, 0x23, 0xb4, 0x09, 0xab, 0xd7, 0xe7, 0xe7, 0x8d, 0x27, 0xa8, 0x02,
	0x9b, 0x67, 0xbd, 0xf3, 0x93, 0x9b, 0xfe, 0xa8, 0x51, 0x42, 0x00, 0x1b, 0xc3, 0x51, 0x78, 0xd1,
	0x1d, 0x35, 0x56, 0x5a, 0x2f, 0x00, 0x16, 0xee, 0x6e, 0x19, 0xd6, 0xae, 0xae, 0xaf, 0x7a, 0x8d,
	0x27, 0xa8, 0x06, 0x70, 0x79, 0x33, 0xba, 0x39, 0xe9, 0xe3, 0x51, 0x7f, 0xd8, 0x28, 0x1d, 0xfd,
	0x06, 0x4f, 0x97, 0x5d, 0xf0, 0xff, 0xf7, 0x40, 0xdb, 0xb0, 0x79, 0x71, 0xf5, 0xa6, 0x17, 0x5e,
	0x8c, 0x1a, 0xff, 0x6c, 0x9e, 0x7e, 0xf1, 0xf6, 0xc0, 0x15, 0xc5, 0x44, 0x87, 0x64, 0xac, 0xf3,
	0xc1, 0x87, 0xd6, 0xed, 0x86, 0xdd, 0xb1, 0x1f, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x56, 0x82,
	0x0e, 0xa4, 0x80, 0x09, 0x00, 0x00,
}
