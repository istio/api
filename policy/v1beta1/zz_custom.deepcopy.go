package v1beta1

import (
	"istio.io/api/util"
)

// DeepCopyInto for AttributeManifest type.
func (in *AttributeManifest) DeepCopyInto(out *AttributeManifest) {
	util.DeepCopyInto(in, out)
}

// DeepCopyInto for Handler type.
func (in *Handler) DeepCopyInto(out *Handler) {
	util.DeepCopyInto(in, out)
}

// DeepCopyInto for Instance type.
func (in *Instance) DeepCopyInto(out *Instance) {
	util.DeepCopyInto(in, out)
}

// DeepCopyInto for Rule type.
func (in *Rule) DeepCopyInto(out *Rule) {
	util.DeepCopyInto(in, out)
}
