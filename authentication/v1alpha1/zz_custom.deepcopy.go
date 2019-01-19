package v1alpha1

import (
	"istio.io/api/util"
)

// DeepCopyInto for Policy type.
func (in *Policy) DeepCopyInto(out *Policy) {
	util.DeepCopyInto(in, out)
}
