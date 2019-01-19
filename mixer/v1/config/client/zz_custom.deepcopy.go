package client

import (
	"istio.io/api/util"
)

// DeepCopyInto for HTTPAPISpec type.
func (in *HTTPAPISpec) DeepCopyInto(out *HTTPAPISpec) {
	util.DeepCopyInto(in, out)
}

// DeepCopyInto for HTTPAPISpecBinding type.
func (in *HTTPAPISpecBinding) DeepCopyInto(out *HTTPAPISpecBinding) {
	util.DeepCopyInto(in, out)
}

// DeepCopyInto for QuotaSpec type.
func (in *QuotaSpec) DeepCopyInto(out *QuotaSpec) {
	util.DeepCopyInto(in, out)
}

// DeepCopyInto for QuotaSpecBinding type.
func (in *QuotaSpecBinding) DeepCopyInto(out *QuotaSpecBinding) {
	util.DeepCopyInto(in, out)
}
