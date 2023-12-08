// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package v1beta1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using ParentReference within kubernetes types, where deepcopy-gen is used.
func (in *ParentReference) DeepCopyInto(out *ParentReference) {
	p := proto.Clone(in).(*ParentReference)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParentReference. Required by controller-gen.
func (in *ParentReference) DeepCopy() *ParentReference {
	if in == nil {
		return nil
	}
	out := new(ParentReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ParentReference. Required by controller-gen.
func (in *ParentReference) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using PolicyAncestorStatus within kubernetes types, where deepcopy-gen is used.
func (in *PolicyAncestorStatus) DeepCopyInto(out *PolicyAncestorStatus) {
	p := proto.Clone(in).(*PolicyAncestorStatus)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyAncestorStatus. Required by controller-gen.
func (in *PolicyAncestorStatus) DeepCopy() *PolicyAncestorStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyAncestorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new PolicyAncestorStatus. Required by controller-gen.
func (in *PolicyAncestorStatus) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Condition within kubernetes types, where deepcopy-gen is used.
func (in *Condition) DeepCopyInto(out *Condition) {
	p := proto.Clone(in).(*Condition)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition. Required by controller-gen.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Condition. Required by controller-gen.
func (in *Condition) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
