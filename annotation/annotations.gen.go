
// GENERATED FILE -- DO NOT EDIT

package annotation

type FeatureStatus int

const (
	Alpha FeatureStatus = iota
	Beta
	Stable
)

func (s FeatureStatus) String() string {
	switch s {
	case Alpha:
		return "Alpha"
	case Beta:
		return "Beta"
	case Stable:
		return "Stable"
	}
	return "Unknown"
}

type ResourceTypes int

const (
	Unknown ResourceTypes = iota
    Any
    AuthorizationPolicy
    Gateway
    GatewayClass
    Ingress
    Namespace
    Pod
    Service
    ServiceEntry
    WorkloadEntry
)

func (r ResourceTypes) String() string {
	switch r {
	case 1:
		return "Any"
	case 2:
		return "AuthorizationPolicy"
	case 3:
		return "Gateway"
	case 4:
		return "GatewayClass"
	case 5:
		return "Ingress"
	case 6:
		return "Namespace"
	case 7:
		return "Pod"
	case 8:
		return "Service"
	case 9:
		return "ServiceEntry"
	case 10:
		return "WorkloadEntry"
	}
	return "Unknown"
}

// Instance describes a single resource annotation
type Instance struct {
	// The name of the annotation.
	Name string

	// Description of the annotation.
	Description string

	// FeatureStatus of this annotation.
	FeatureStatus FeatureStatus

	// Hide the existence of this annotation when outputting usage information.
	Hidden bool

	// Mark this annotation as deprecated when generating usage information.
	Deprecated bool

	// The types of resources this annotation applies to.
	Resources []ResourceTypes
}

var (

	AlphaCanonicalServiceAccounts = Instance {
		Name:          "alpha.istio.io/canonical-serviceaccounts",
		Description:   "Specifies the non-Kubernetes service accounts that are "+
                        "allowed to run this service.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Service,
		},
	}

	AlphaKubernetesServiceAccounts = Instance {
		Name:          "alpha.istio.io/kubernetes-serviceaccounts",
		Description:   "Specifies the Kubernetes service accounts that are "+
                        "allowed to run this service on the VMs.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Service,
		},
	}

	AmbientBypassInboundCapture = Instance {
		Name:          "ambient.istio.io/bypass-inbound-capture",
		Description:   `When specified on a "Pod" enrolled in ambient mesh, only outbound traffic will be captured.
This is intended to be used when enrolling a workload that only receives traffic from out-of-the-mesh clients, such as third party ingress controllers.
`,
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	AmbientRedirection = Instance {
		Name:          "ambient.istio.io/redirection",
		Description:   `Automatically configured by Istio to indicate a Pod was successfully enrolled in ambient mode.
This shows the actual state; to specify intent that a workload should be in ambient mode, see "istio.io/dataplane-mode".
User should not manually modify this annotation.`,
		FeatureStatus: Beta,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	AmbientWaypointInboundBinding = Instance {
		Name:          "ambient.istio.io/waypoint-inbound-binding",
		Description:   `When set on a waypoint (either by its specific "Gateway", or for the entire collection on the "GatewayClass"),
indicates how traffic should be sent to the waypoint. If unset, traffic will be sent to the waypoint as HBONE directly.

This takes the format: "<protocol>" or "<protocol>/<port>".
`,
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			GatewayClass,
			Gateway,
		},
	}

	GalleyAnalyzeSuppress = Instance {
		Name:          "galley.istio.io/analyze-suppress",
		Description:   "A comma separated list of configuration analysis message "+
                        "codes to suppress when Istio analyzers are run. For "+
                        "example, to suppress reporting of IST0103 "+
                        "(PodMissingProxy) and IST0108 (UnknownAnnotation) on a "+
                        "resource, apply the annotation "+
                        "'galley.istio.io/analyze-suppress=IST0108,IST0103'. If "+
                        "the value is '*', then all configuration analysis "+
                        "messages are suppressed.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Any,
		},
	}

	GatewayControllerVersion = Instance {
		Name:          "gateway.istio.io/controller-version",
		Description:   "A version added to the Gateway by the controller "+
                        "specifying the `controller version`.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Any,
		},
	}

	GatewayNameOverride = Instance {
		Name:          "gateway.istio.io/name-override",
		Description:   `Overrides the name of the generated "Deployment" and "Service" resource when using [Gateway auto-deployment](/docs/tasks/traffic-management/ingress/gateway-api/#automated-deployment)
`,
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Gateway,
		},
	}

	GatewayServiceAccount = Instance {
		Name:          "gateway.istio.io/service-account",
		Description:   `Overrides the name of the generated "ServiceAccount" resource when using [Gateway auto-deployment](/docs/tasks/traffic-management/ingress/gateway-api/#automated-deployment)
`,
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Gateway,
		},
	}

	InjectTemplates = Instance {
		Name:          "inject.istio.io/templates",
		Description:   "The name of the inject template(s) to use, as a comma "+
                        "separate list. See "+
                        "https://istio.io/latest/docs/setup/additional-setup/sidecar-injection/#custom-templates-experimental "+
                        "for more information.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	IoIstioAutoRegistrationGroup = Instance {
		Name:          "istio.io/autoRegistrationGroup",
		Description:   "On a WorkloadEntry stores the associated WorkloadGroup.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			WorkloadEntry,
		},
	}

	IoIstioConnectedAt = Instance {
		Name:          "istio.io/connectedAt",
		Description:   "On a WorkloadEntry stores the time in nanoseconds when "+
                        "the associated workload connected to a Pilot instance.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			WorkloadEntry,
		},
	}

	IoIstioDisconnectedAt = Instance {
		Name:          "istio.io/disconnectedAt",
		Description:   "On a WorkloadEntry stores the time in nanoseconds when "+
                        "the associated workload disconnected from a Pilot "+
                        "instance.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			WorkloadEntry,
		},
	}

	IoIstioDryRun = Instance {
		Name:          "istio.io/dry-run",
		Description:   "Specifies whether or not the given resource is in dry-run "+
                        "mode. See "+
                        "https://istio.io/latest/docs/tasks/security/authorization/authz-dry-run/ "+
                        "for more information.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			AuthorizationPolicy,
		},
	}

	IoIstioRev = Instance {
		Name:          "istio.io/rev",
		Description:   "Specifies a control plane revision to which a given proxy "+
                        "is connected. This annotation is added automatically, not "+
                        "set by a user. In contrary to the label istio.io/rev, it "+
                        "represents the actual revision, not the requested "+
                        "revision.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	IoIstioWorkloadController = Instance {
		Name:          "istio.io/workloadController",
		Description:   "On a WorkloadEntry should store the current/last pilot "+
                        "instance connected to the workload for XDS.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			WorkloadEntry,
		},
	}

	IoKubernetesIngressClass = Instance {
		Name:          "kubernetes.io/ingress.class",
		Description:   "Annotation on an Ingress resources denoting the class of "+
                        "controllers responsible for it.",
		FeatureStatus: Stable,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Ingress,
		},
	}

	NetworkingExportTo = Instance {
		Name:          "networking.istio.io/exportTo",
		Description:   "Specifies the namespaces to which this service should be "+
                        "exported to. A value of '*' indicates it is reachable "+
                        "within the mesh '.' indicates it is reachable within its "+
                        "namespace.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Service,
		},
	}

	NetworkingServiceType = Instance {
		Name:          "networking.istio.io/service-type",
		Description:   `Overrides the type of the generated "Service" resource when using [Gateway auto-deployment](/docs/tasks/traffic-management/ingress/gateway-api/#automated-deployment)
`,
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Gateway,
		},
	}

	NetworkingTrafficDistribution = Instance {
		Name:          "networking.istio.io/traffic-distribution",
		Description:   `Controls how traffic is distributed across the set of available endpoints.

At this time, this annotation only impacts routing done by Ztunnel.

Accepted values:
* "PreferClose": endpoints will be categorized by how "close" they are, consider network, region, zone, and subzone.
  Traffic will be prioritized to the closest healthy endpoints.
  For example, if I have a Service with "PreferClose" set, with endpoints in zones "us-west,us-west,us-east". When 
  sending traffic from a client in zone "us-west", all traffic will go to the two "us-west" backends.
  If one those backends become unhealthy, all traffic will go to the remaining endpoint in "us-west".
  If that backend becomes unhealthy, traffic will sent to "us-east".
`,
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Service,
			ServiceEntry,
		},
	}

	PrometheusMergeMetrics = Instance {
		Name:          "prometheus.istio.io/merge-metrics",
		Description:   "Specifies if application Prometheus metric will be merged "+
                        "with Envoy metrics for this workload.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	ProxyConfig = Instance {
		Name:          "proxy.istio.io/config",
		Description:   "Overrides for the proxy configuration for this specific "+
                        "proxy. Available options can be found at "+
                        "https://istio.io/docs/reference/config/istio.mesh.v1alpha1/#ProxyConfig.",
		FeatureStatus: Beta,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	ProxyOverrides = Instance {
		Name:          "proxy.istio.io/overrides",
		Description:   "Used internally to indicate user-specified overrides in "+
                        "the proxy container of the pod during injection.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarStatusReadinessApplicationPorts = Instance {
		Name:          "readiness.status.sidecar.istio.io/applicationPorts",
		Description:   "Specifies the list of ports exposed by the application "+
                        "container. Used by the Envoy sidecar readiness probe to "+
                        "determine that Envoy is configured and ready to receive "+
                        "traffic.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarStatusReadinessFailureThreshold = Instance {
		Name:          "readiness.status.sidecar.istio.io/failureThreshold",
		Description:   "Specifies the failure threshold for the Envoy sidecar "+
                        "readiness probe.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarStatusReadinessInitialDelaySeconds = Instance {
		Name:          "readiness.status.sidecar.istio.io/initialDelaySeconds",
		Description:   "Specifies the initial delay (in seconds) for the Envoy "+
                        "sidecar readiness probe.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarStatusReadinessPeriodSeconds = Instance {
		Name:          "readiness.status.sidecar.istio.io/periodSeconds",
		Description:   "Specifies the period (in seconds) for the Envoy sidecar "+
                        "readiness probe.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarAgentLogLevel = Instance {
		Name:          "sidecar.istio.io/agentLogLevel",
		Description:   "Specifies the log output level for pilot-agent.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarBootstrapOverride = Instance {
		Name:          "sidecar.istio.io/bootstrapOverride",
		Description:   "Specifies an alternative Envoy bootstrap configuration "+
                        "file.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarComponentLogLevel = Instance {
		Name:          "sidecar.istio.io/componentLogLevel",
		Description:   "Specifies the component log level for Envoy.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarDiscoveryAddress = Instance {
		Name:          "sidecar.istio.io/discoveryAddress",
		Description:   "Specifies the XDS discovery address to be used by the "+
                        "Envoy sidecar.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarEnableCoreDump = Instance {
		Name:          "sidecar.istio.io/enableCoreDump",
		Description:   "Specifies whether or not an Envoy sidecar should enable "+
                        "core dump.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarExtraStatTags = Instance {
		Name:          "sidecar.istio.io/extraStatTags",
		Description:   "An additional list of tags to extract from the in-proxy "+
                        "Istio Wasm telemetry. Each additional tag needs to be "+
                        "present in this list.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarInject = Instance {
		Name:          "sidecar.istio.io/inject",
		Description:   "Specifies whether or not an Envoy sidecar should be "+
                        "automatically injected into the workload. Deprecated in "+
                        "favor of `sidecar.istio.io/inject` label.",
		FeatureStatus: Beta,
		Hidden:        false,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarInterceptionMode = Instance {
		Name:          "sidecar.istio.io/interceptionMode",
		Description:   "Specifies the mode used to redirect inbound connections "+
                        "to Envoy (REDIRECT or TPROXY).",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarLogLevel = Instance {
		Name:          "sidecar.istio.io/logLevel",
		Description:   "Specifies the log level for Envoy.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarNativeSidecar = Instance {
		Name:          "sidecar.istio.io/nativeSidecar",
		Description:   "Specifies if the istio-proxy sidecar should be injected "+
                        "as a native sidecar or not. Takes precedence over the "+
                        "ENABLE_NATIVE_SIDECARS environment variable.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarProxyCPU = Instance {
		Name:          "sidecar.istio.io/proxyCPU",
		Description:   "Specifies the requested CPU setting for the Envoy "+
                        "sidecar.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarProxyCPULimit = Instance {
		Name:          "sidecar.istio.io/proxyCPULimit",
		Description:   "Specifies the CPU limit for the Envoy sidecar.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarProxyImage = Instance {
		Name:          "sidecar.istio.io/proxyImage",
		Description:   "Specifies the Docker image to be used by the Envoy "+
                        "sidecar.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarProxyImageType = Instance {
		Name:          "sidecar.istio.io/proxyImageType",
		Description:   "Specifies the Docker image type to be used by the Envoy "+
                        "sidecar. Istio publishes debug and distroless image types "+
                        "for every release tag.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarProxyMemory = Instance {
		Name:          "sidecar.istio.io/proxyMemory",
		Description:   "Specifies the requested memory setting for the Envoy "+
                        "sidecar.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarProxyMemoryLimit = Instance {
		Name:          "sidecar.istio.io/proxyMemoryLimit",
		Description:   "Specifies the memory limit for the Envoy sidecar.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarRewriteAppHTTPProbers = Instance {
		Name:          "sidecar.istio.io/rewriteAppHTTPProbers",
		Description:   "Rewrite HTTP readiness and liveness probes to be "+
                        "redirected to the Envoy sidecar.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarStatsHistogramBuckets = Instance {
		Name:          "sidecar.istio.io/statsHistogramBuckets",
		Description:   "Specifies the custom histogram buckets with a prefix "+
                        "matcher to separate the Istio mesh metrics from the Envoy "+
                        "stats, e.g. "+
                        "`{`istiocustom`:[1,5,10,50,100,500,1000,5000,10000],`cluster.xds-grpc`:[1,5,10,25,50,100,250,500,1000,2500,5000,10000]}`. "+
                        "Default buckets are "+
                        "`[0.5,1,5,10,25,50,100,250,500,1000,2500,5000,10000,30000,60000,300000,600000,1800000,3600000]`.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarStatsInclusionPrefixes = Instance {
		Name:          "sidecar.istio.io/statsInclusionPrefixes",
		Description:   "Specifies the comma separated list of prefixes of the "+
                        "stats to be emitted by Envoy.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarStatsInclusionRegexps = Instance {
		Name:          "sidecar.istio.io/statsInclusionRegexps",
		Description:   "Specifies the comma separated list of regexes the stats "+
                        "should match to be emitted by Envoy.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarStatsInclusionSuffixes = Instance {
		Name:          "sidecar.istio.io/statsInclusionSuffixes",
		Description:   "Specifies the comma separated list of suffixes of the "+
                        "stats to be emitted by Envoy.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarStatus = Instance {
		Name:          "sidecar.istio.io/status",
		Description:   "Generated by Envoy sidecar injection that indicates the "+
                        "status of the operation. Includes a version hash of the "+
                        "executed template, as well as names of injected "+
                        "resources.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarUserVolume = Instance {
		Name:          "sidecar.istio.io/userVolume",
		Description:   "Specifies one or more user volumes (as a JSON array) to "+
                        "be added to the Envoy sidecar.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarUserVolumeMount = Instance {
		Name:          "sidecar.istio.io/userVolumeMount",
		Description:   "Specifies one or more user volume mounts (as a JSON "+
                        "array) to be added to the Envoy sidecar.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarStatusPort = Instance {
		Name:          "status.sidecar.istio.io/port",
		Description:   "Specifies the HTTP status Port for the Envoy sidecar. If "+
                        "zero, the sidecar will not provide status.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	TopologyControlPlaneClusters = Instance {
		Name:          "topology.istio.io/controlPlaneClusters",
		Description:   "A comma-separated list of clusters (or * for any) running "+
                        "istiod that should attempt leader election for a remote "+
                        "cluster thats system namespace includes this annotation. "+
                        "Istiod will not attempt to lead unannotated remote "+
                        "clusters.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Namespace,
		},
	}

	TrafficNodeSelector = Instance {
		Name:          "traffic.istio.io/nodeSelector",
		Description:   "This annotation is a set of node-labels "+
                        "(key1=value,key2=value). If the annotated Service is of "+
                        "type NodePort and is a multi-network gateway (see "+
                        "topology.istio.io/network), the addresses for selected "+
                        "nodes will be used for cross-network communication.",
		FeatureStatus: Stable,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Service,
		},
	}

	SidecarTrafficExcludeInboundPorts = Instance {
		Name:          "traffic.sidecar.istio.io/excludeInboundPorts",
		Description:   "A comma separated list of inbound ports to be excluded "+
                        "from redirection to Envoy. Only applies when all inbound "+
                        "traffic (i.e. '*') is being redirected.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarTrafficExcludeInterfaces = Instance {
		Name:          "traffic.sidecar.istio.io/excludeInterfaces",
		Description:   "A comma separated list of interfaces to be excluded from "+
                        "Istio traffic capture",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarTrafficExcludeOutboundIPRanges = Instance {
		Name:          "traffic.sidecar.istio.io/excludeOutboundIPRanges",
		Description:   "A comma separated list of IP ranges in CIDR form to be "+
                        "excluded from redirection. Only applies when all outbound "+
                        "traffic (i.e. '*') is being redirected.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarTrafficExcludeOutboundPorts = Instance {
		Name:          "traffic.sidecar.istio.io/excludeOutboundPorts",
		Description:   "A comma separated list of outbound ports to be excluded "+
                        "from redirection to Envoy.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarTrafficIncludeInboundPorts = Instance {
		Name:          "traffic.sidecar.istio.io/includeInboundPorts",
		Description:   "A comma separated list of inbound ports for which traffic "+
                        "is to be redirected to Envoy. The wildcard character '*' "+
                        "can be used to configure redirection for all ports. An "+
                        "empty list will disable all inbound redirection.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarTrafficIncludeOutboundIPRanges = Instance {
		Name:          "traffic.sidecar.istio.io/includeOutboundIPRanges",
		Description:   "A comma separated list of IP ranges in CIDR form to "+
                        "redirect to Envoy (optional). The wildcard character '*' "+
                        "can be used to redirect all outbound traffic. An empty "+
                        "list will disable all outbound redirection.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarTrafficIncludeOutboundPorts = Instance {
		Name:          "traffic.sidecar.istio.io/includeOutboundPorts",
		Description:   "A comma separated list of outbound ports for which "+
                        "traffic is to be redirected to Envoy, regardless of the "+
                        "destination IP.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	SidecarTrafficKubevirtInterfaces = Instance {
		Name:          "traffic.sidecar.istio.io/kubevirtInterfaces",
		Description:   "A comma separated list of virtual interfaces whose "+
                        "inbound traffic (from VM) will be treated as outbound.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

)

func AllResourceAnnotations() []*Instance {
	return []*Instance {
		&AlphaCanonicalServiceAccounts,
		&AlphaKubernetesServiceAccounts,
		&AmbientBypassInboundCapture,
		&AmbientRedirection,
		&AmbientWaypointInboundBinding,
		&GalleyAnalyzeSuppress,
		&GatewayControllerVersion,
		&GatewayNameOverride,
		&GatewayServiceAccount,
		&InjectTemplates,
		&IoIstioAutoRegistrationGroup,
		&IoIstioConnectedAt,
		&IoIstioDisconnectedAt,
		&IoIstioDryRun,
		&IoIstioRev,
		&IoIstioWorkloadController,
		&IoKubernetesIngressClass,
		&NetworkingExportTo,
		&NetworkingServiceType,
		&NetworkingTrafficDistribution,
		&PrometheusMergeMetrics,
		&ProxyConfig,
		&ProxyOverrides,
		&SidecarStatusReadinessApplicationPorts,
		&SidecarStatusReadinessFailureThreshold,
		&SidecarStatusReadinessInitialDelaySeconds,
		&SidecarStatusReadinessPeriodSeconds,
		&SidecarAgentLogLevel,
		&SidecarBootstrapOverride,
		&SidecarComponentLogLevel,
		&SidecarDiscoveryAddress,
		&SidecarEnableCoreDump,
		&SidecarExtraStatTags,
		&SidecarInject,
		&SidecarInterceptionMode,
		&SidecarLogLevel,
		&SidecarNativeSidecar,
		&SidecarProxyCPU,
		&SidecarProxyCPULimit,
		&SidecarProxyImage,
		&SidecarProxyImageType,
		&SidecarProxyMemory,
		&SidecarProxyMemoryLimit,
		&SidecarRewriteAppHTTPProbers,
		&SidecarStatsHistogramBuckets,
		&SidecarStatsInclusionPrefixes,
		&SidecarStatsInclusionRegexps,
		&SidecarStatsInclusionSuffixes,
		&SidecarStatus,
		&SidecarUserVolume,
		&SidecarUserVolumeMount,
		&SidecarStatusPort,
		&TopologyControlPlaneClusters,
		&TrafficNodeSelector,
		&SidecarTrafficExcludeInboundPorts,
		&SidecarTrafficExcludeInterfaces,
		&SidecarTrafficExcludeOutboundIPRanges,
		&SidecarTrafficExcludeOutboundPorts,
		&SidecarTrafficIncludeInboundPorts,
		&SidecarTrafficIncludeOutboundIPRanges,
		&SidecarTrafficIncludeOutboundPorts,
		&SidecarTrafficKubevirtInterfaces,
	}
}

func AllResourceTypes() []string {
	return []string {
		"Any",
		"AuthorizationPolicy",
		"Gateway",
		"GatewayClass",
		"Ingress",
		"Namespace",
		"Pod",
		"Service",
		"ServiceEntry",
		"WorkloadEntry",
	}
}
