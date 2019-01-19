package v1alpha3

import (
	"istio.io/api/util"
)

// DeepCopyInto for DestinationRule type.
func (in *DestinationRule) DeepCopyInto(out *DestinationRule) {
	util.DeepCopyInto(in, out)
}

// DeepCopyInto for EnvoyFilter type.
func (in *EnvoyFilter) DeepCopyInto(out *EnvoyFilter) {
	util.DeepCopyInto(in, out)
}

// DeepCopyInto for Gateway type.
func (in *Gateway) DeepCopyInto(out *Gateway) {
	util.DeepCopyInto(in, out)
}

// DeepCopyInto for ServiceEntry type.
func (in *ServiceEntry) DeepCopyInto(out *ServiceEntry) {
	util.DeepCopyInto(in, out)
}

// DeepCopyInto for Sidecar type.
func (in *Sidecar) DeepCopyInto(out *Sidecar) {
	util.DeepCopyInto(in, out)
}

// DeepCopyInto for VirtualService type.
func (in *VirtualService) DeepCopyInto(out *VirtualService) {
	util.DeepCopyInto(in, out)
}
