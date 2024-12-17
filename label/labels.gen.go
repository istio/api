
// GENERATED FILE -- DO NOT EDIT

package label

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
    Deployment
    Gateway
    GatewayClass
    Namespace
    Node
    Pod
    Service
    ServiceAccount
    ServiceEntry
    WorkloadEntry
)

func (r ResourceTypes) String() string {
	switch r {
	case 1:
		return "Any"
	case 2:
		return "Deployment"
	case 3:
		return "Gateway"
	case 4:
		return "GatewayClass"
	case 5:
		return "Namespace"
	case 6:
		return "Node"
	case 7:
		return "Pod"
	case 8:
		return "Service"
	case 9:
		return "ServiceAccount"
	case 10:
		return "ServiceEntry"
	case 11:
		return "WorkloadEntry"
	}
	return "Unknown"
}

// Instance describes a single resource label
type Instance struct {
	// The name of the label.
	Name string

	// Description of the label.
	Description string

	// FeatureStatus of this label.
	FeatureStatus FeatureStatus

	// Hide the existence of this label when outputting usage information.
	Hidden bool

	// Mark this label as deprecated when generating usage information.
	Deprecated bool

	// The types of resources this label applies to.
	Resources []ResourceTypes
}

var (

	GatewayManaged = Instance {
		Name:          "gateway.istio.io/managed",
		Description:   "Automatically added to all resources [automatically "+
                        "created](/docs/tasks/traffic-management/ingress/gateway-api/#automated-deployment) "+
                        "by Istio Gateway controller, to indicate which controller "+
                        "created the resource. Users should not set this label "+
                        "themselves.",
		FeatureStatus: Stable,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			ServiceAccount,
			Deployment,
			Service,
		},
	}

	IoK8sNetworkingGatewayGatewayName = Instance {
		Name:          "gateway.networking.k8s.io/gateway-name",
		Description:   "Automatically added to all resources [automatically "+
                        "created](/docs/tasks/traffic-management/ingress/gateway-api/#automated-deployment) "+
                        "by Istio Gateway controller to indicate which `Gateway` "+
                        "resulted in the object creation. Users should not set "+
                        "this label themselves.",
		FeatureStatus: Stable,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			ServiceAccount,
			Deployment,
			Service,
		},
	}

	IoIstioDataplaneMode = Instance {
		Name:          "istio.io/dataplane-mode",
		Description:   `When set on a resource, indicates the [data plane mode](/docs/overview/dataplane-modes/) to use.
Possible values: "ambient", "none".
Note: users wishing to use sidecar mode should see the "istio-injection" label; there is no value on this label to configure sidecars.
`,
		FeatureStatus: Stable,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
			Namespace,
		},
	}

	IoIstioRev = Instance {
		Name:          "istio.io/rev",
		Description:   "Istio control plane revision associated with the "+
                        "resource; e.g. `canary`",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Namespace,
		},
	}

	IoIstioTag = Instance {
		Name:          "istio.io/tag",
		Description:   "Istio control plane tag name associated with the "+
                        "resource; e.g. `canary`",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Namespace,
		},
	}

	IoIstioUseWaypoint = Instance {
		Name:          "istio.io/use-waypoint",
		Description:   `When set on a resource, indicates the resource has an associated waypoint with the given name.
The waypoint is assumed to be in the same namespace; for cross-namespace, see "istio.io/use-waypoint-namespace".

When set or a "Pod" or a "Service", this binds that specific resource to the waypoint.
When set on a "Namespace", this applies to all "Pod"/"Service" in the namespace.

Note: the waypoint must allow the type, see "istio.io/waypoint-for".
`,
		FeatureStatus: Stable,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
			WorkloadEntry,
			Service,
			ServiceEntry,
			Namespace,
		},
	}

	IoIstioUseWaypointNamespace = Instance {
		Name:          "istio.io/use-waypoint-namespace",
		Description:   `When set on a resource, indicates the resource has an associated waypoint in the provided namespace.
This must be set in addition to "istio.io/use-waypoint", when a cross-namespace reference is desired.
`,
		FeatureStatus: Beta,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
			WorkloadEntry,
			Service,
			ServiceEntry,
			Namespace,
		},
	}

	IoIstioWaypointFor = Instance {
		Name:          "istio.io/waypoint-for",
		Description:   `When set on a waypoint (either by its specific "Gateway", or for the entire collection on the "GatewayClass"),
indicates the type of traffic this waypoint can handle.

Valid options: "service", "workload", "all", and "none".
`,
		FeatureStatus: Stable,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			GatewayClass,
			Gateway,
		},
	}

	NetworkingEnableAutoallocateIp = Instance {
		Name:          "networking.istio.io/enable-autoallocate-ip",
		Description:   `Configures whether a "ServiceEntry" without any "spec.addresses" set should get an IP address automatically allocated for it.

Valid options: "true", "false"
`,
		FeatureStatus: Beta,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			ServiceEntry,
		},
	}

	NetworkingGatewayPort = Instance {
		Name:          "networking.istio.io/gatewayPort",
		Description:   "IstioGatewayPortLabel overrides the default 15443 value "+
                        "to use for a multi-network gateway's port",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Service,
		},
	}

	OperatorComponent = Instance {
		Name:          "operator.istio.io/component",
		Description:   "Istio operator component name of the resource, e.g. "+
                        "`Pilot`",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Any,
		},
	}

	OperatorManaged = Instance {
		Name:          "operator.istio.io/managed",
		Description:   "Set to `Reconcile` if the Istio operator will reconcile "+
                        "the resource.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Any,
		},
	}

	OperatorVersion = Instance {
		Name:          "operator.istio.io/version",
		Description:   "The Istio operator version that installed the resource, "+
                        "e.g. `1.6.0`",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Any,
		},
	}

	SecurityTlsMode = Instance {
		Name:          "security.istio.io/tlsMode",
		Description:   "Specifies the TLS mode supported by a sidecar proxy. "+
                        "Valid values are 'istio', 'disabled'. When injecting "+
                        "sidecars into Pods, the sidecar injector will set the "+
                        "value of this label to 'istio' indicating that the "+
                        "sidecar is capable of supporting mTLS. Clients injected "+
                        "with sidecar proxies will opportunistically use this "+
                        "label to determine whether or not to secure the traffic "+
                        "to this workload using Istio mutual TLS.",
		FeatureStatus: Alpha,
		Hidden:        true,
		Deprecated:    true,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	ServiceCanonicalName = Instance {
		Name:          "service.istio.io/canonical-name",
		Description:   "The name of the canonical service a workload belongs to",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	ServiceCanonicalRevision = Instance {
		Name:          "service.istio.io/canonical-revision",
		Description:   "The name of a revision within a canonical service that "+
                        "the workload belongs to",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	ServiceWorkloadName = Instance {
		Name:          "service.istio.io/workload-name",
		Description:   `The workload name of the application a workload belongs to. If unset, defaults to the detect parent resource.
For example, a "Pod" resource may default to the "Deployment" name.
`,
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
			WorkloadEntry,
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

	TopologyCluster = Instance {
		Name:          "topology.istio.io/cluster",
		Description:   "This label is applied to a workload internally that "+
                        "identifies the Kubernetes cluster containing the "+
                        "workload. The cluster ID is specified during Istio "+
                        "installation for each cluster via "+
                        "`values.global.multiCluster.clusterName`. It should be "+
                        "noted that this is only used internally within Istio and "+
                        "is not an actual label on workload pods. If a pod "+
                        "contains this label, it will be overridden by Istio "+
                        "internally with the cluster ID specified during Istio "+
                        "installation. This label provides a way to select "+
                        "workloads by cluster when using DestinationRules. For "+
                        "example, a service owner could create a DestinationRule "+
                        "containing a subset per cluster and then use these "+
                        "subsets to control traffic flow to each cluster "+
                        "independently.",
		FeatureStatus: Alpha,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Pod,
		},
	}

	TopologyNetwork = Instance {
		Name:          "topology.istio.io/network",
		Description:   `A label used to identify the network for one or more pods. This is used
internally by Istio to group pods resident in the same L3 domain/network.
Istio assumes that pods in the same network are directly reachable from
one another. When pods are in different networks, an Istio Gateway
(e.g. east-west gateway) is typically used to establish connectivity
(with AUTO_PASSTHROUGH mode). This label can be applied to the following
resources to help automate Istio's multi-network configuration.

* Istio System Namespace: Applying this label to the system namespace
  establishes a default network for pods managed by the control plane.
  This is typically configured during control plane installation using an
  admin-specified value.

* Pod: Applying this label to a pod allows overriding the default network
  on a per-pod basis. This is typically applied to the pod via webhook
  injection, but can also be manually specified on the pod by the service
  owner. The Istio installation in each cluster configures webhook injection
  using an admin-specified value.

* Gateway Service: Applying this label to the Service for an Istio Gateway,
  indicates that Istio should use this service as the gateway for the
  network, when configuring cross-network traffic. Istio will configure
  pods residing outside of the network to access the Gateway service
  via "spec.externalIPs", "status.loadBalancer.ingress[].ip", or in the case
  of a NodePort service, the Node's address. The label is configured when
  installing the gateway (e.g. east-west gateway) and should match either
  the default network for the control plane (as specified by the Istio System
  Namespace label) or the network of the targeted pods.`,
		FeatureStatus: Beta,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Namespace,
			Pod,
			Service,
		},
	}

	TopologySubzone = Instance {
		Name:          "topology.istio.io/subzone",
		Description:   "User-provided node label for identifying the locality "+
                        "subzone of a workload. This allows admins to specify a "+
                        "more granular level of locality than what is offered by "+
                        "default with Kubernetes regions and zones.",
		FeatureStatus: Beta,
		Hidden:        false,
		Deprecated:    false,
		Resources: []ResourceTypes{
			Node,
		},
	}

)

func AllResourceLabels() []*Instance {
	return []*Instance {
		&GatewayManaged,
		&IoK8sNetworkingGatewayGatewayName,
		&IoIstioDataplaneMode,
		&IoIstioRev,
		&IoIstioTag,
		&IoIstioUseWaypoint,
		&IoIstioUseWaypointNamespace,
		&IoIstioWaypointFor,
		&NetworkingEnableAutoallocateIp,
		&NetworkingGatewayPort,
		&OperatorComponent,
		&OperatorManaged,
		&OperatorVersion,
		&SecurityTlsMode,
		&ServiceCanonicalName,
		&ServiceCanonicalRevision,
		&ServiceWorkloadName,
		&SidecarInject,
		&TopologyCluster,
		&TopologyNetwork,
		&TopologySubzone,
	}
}

func AllResourceTypes() []string {
	return []string {
		"Any",
		"Deployment",
		"Gateway",
		"GatewayClass",
		"Namespace",
		"Node",
		"Pod",
		"Service",
		"ServiceAccount",
		"ServiceEntry",
		"WorkloadEntry",
	}
}
