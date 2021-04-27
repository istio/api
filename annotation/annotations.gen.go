
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
    Ingress
    Pod
    Service
)

func (r ResourceTypes) String() string {
	switch r {
	case 1:
		return "Any"
	case 2:
		return "AuthorizationPolicy"
	case 3:
		return "Ingress"
	case 4:
		return "Pod"
	case 5:
		return "Service"
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

	AlphaIdentity = Instance {
		Name:          "alpha.istio.io/identity",
		Description:   "Identity for the workload.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Pod,
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

	OperatorInstallChartOwner = Instance {
		Name:          "install.operator.istio.io/chart-owner",
		Description:   "Represents the name of the chart used to create this "+
                        "resource.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Any,
		},
	}

	OperatorInstallOwnerGeneration = Instance {
		Name:          "install.operator.istio.io/owner-generation",
		Description:   "Represents the generation to which the resource was last "+
                        "reconciled.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Any,
		},
	}

	OperatorInstallVersion = Instance {
		Name:          "install.operator.istio.io/version",
		Description:   "Represents the Istio version associated with the resource",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Any,
		},
	}

	IoIstioDryRun = Instance {
		Name:          "istio.io/dry-run",
		Description:   "Specifies whether or not the given resource is in dry-run "+
                        "mode.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			AuthorizationPolicy,
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

	SidecarControlPlaneAuthPolicy = Instance {
		Name:          "sidecar.istio.io/controlPlaneAuthPolicy",
		Description:   "Specifies the auth policy used by the Istio control "+
                        "plane. If NONE, traffic will not be encrypted. If "+
                        "MUTUAL_TLS, traffic between Envoy sidecar will be wrapped "+
                        "into mutual TLS connections.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    true,
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

	SidecarInject = Instance {
		Name:          "sidecar.istio.io/inject",
		Description:   "Specifies whether or not an Envoy sidecar should be "+
                        "automatically injected into the workload.",
		FeatureStatus: Beta,
		Hidden:        false,
		Deprecated:    false,
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
		&AlphaIdentity,
		&AlphaKubernetesServiceAccounts,
		&GalleyAnalyzeSuppress,
		&OperatorInstallChartOwner,
		&OperatorInstallOwnerGeneration,
		&OperatorInstallVersion,
		&IoIstioDryRun,
		&IoKubernetesIngressClass,
		&NetworkingExportTo,
		&PrometheusMergeMetrics,
		&ProxyConfig,
		&SidecarStatusReadinessApplicationPorts,
		&SidecarStatusReadinessFailureThreshold,
		&SidecarStatusReadinessInitialDelaySeconds,
		&SidecarStatusReadinessPeriodSeconds,
		&SidecarAgentLogLevel,
		&SidecarBootstrapOverride,
		&SidecarComponentLogLevel,
		&SidecarControlPlaneAuthPolicy,
		&SidecarDiscoveryAddress,
		&SidecarEnableCoreDump,
		&SidecarInject,
		&SidecarInterceptionMode,
		&SidecarLogLevel,
		&SidecarProxyCPU,
		&SidecarProxyCPULimit,
		&SidecarProxyImage,
		&SidecarProxyMemory,
		&SidecarProxyMemoryLimit,
		&SidecarRewriteAppHTTPProbers,
		&SidecarStatsInclusionPrefixes,
		&SidecarStatsInclusionRegexps,
		&SidecarStatsInclusionSuffixes,
		&SidecarStatus,
		&SidecarUserVolume,
		&SidecarUserVolumeMount,
		&SidecarStatusPort,
		&SidecarTrafficExcludeInboundPorts,
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
		"Ingress",
		"Pod",
		"Service",
	}
}
