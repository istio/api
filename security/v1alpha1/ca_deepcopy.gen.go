// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package v1alpha1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using IstioCertificateRequest within kubernetes types, where deepcopy-gen is used.
func (in *IstioCertificateRequest) DeepCopyInto(out *IstioCertificateRequest) {
	p := proto.Clone(in).(*IstioCertificateRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioCertificateRequest. Required by controller-gen.
func (in *IstioCertificateRequest) DeepCopy() *IstioCertificateRequest {
	if in == nil {
		return nil
	}
	out := new(IstioCertificateRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new IstioCertificateRequest. Required by controller-gen.
func (in *IstioCertificateRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using IstioCertificateResponse within kubernetes types, where deepcopy-gen is used.
func (in *IstioCertificateResponse) DeepCopyInto(out *IstioCertificateResponse) {
	p := proto.Clone(in).(*IstioCertificateResponse)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioCertificateResponse. Required by controller-gen.
func (in *IstioCertificateResponse) DeepCopy() *IstioCertificateResponse {
	if in == nil {
		return nil
	}
	out := new(IstioCertificateResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new IstioCertificateResponse. Required by controller-gen.
func (in *IstioCertificateResponse) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
