// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mesh/v1alpha1/proxy.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "istio.io/api/networking/v1alpha3"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using Tracing within kubernetes types, where deepcopy-gen is used.
func (in *Tracing) DeepCopyInto(out *Tracing) {
	p := proto.Clone(in).(*Tracing)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tracing. Required by controller-gen.
func (in *Tracing) DeepCopy() *Tracing {
	if in == nil {
		return nil
	}
	out := new(Tracing)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Tracing. Required by controller-gen.
func (in *Tracing) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Tracing_Zipkin within kubernetes types, where deepcopy-gen is used.
func (in *Tracing_Zipkin) DeepCopyInto(out *Tracing_Zipkin) {
	p := proto.Clone(in).(*Tracing_Zipkin)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tracing_Zipkin. Required by controller-gen.
func (in *Tracing_Zipkin) DeepCopy() *Tracing_Zipkin {
	if in == nil {
		return nil
	}
	out := new(Tracing_Zipkin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Tracing_Zipkin. Required by controller-gen.
func (in *Tracing_Zipkin) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Tracing_Lightstep within kubernetes types, where deepcopy-gen is used.
func (in *Tracing_Lightstep) DeepCopyInto(out *Tracing_Lightstep) {
	p := proto.Clone(in).(*Tracing_Lightstep)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tracing_Lightstep. Required by controller-gen.
func (in *Tracing_Lightstep) DeepCopy() *Tracing_Lightstep {
	if in == nil {
		return nil
	}
	out := new(Tracing_Lightstep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Tracing_Lightstep. Required by controller-gen.
func (in *Tracing_Lightstep) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Tracing_Datadog within kubernetes types, where deepcopy-gen is used.
func (in *Tracing_Datadog) DeepCopyInto(out *Tracing_Datadog) {
	p := proto.Clone(in).(*Tracing_Datadog)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tracing_Datadog. Required by controller-gen.
func (in *Tracing_Datadog) DeepCopy() *Tracing_Datadog {
	if in == nil {
		return nil
	}
	out := new(Tracing_Datadog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Tracing_Datadog. Required by controller-gen.
func (in *Tracing_Datadog) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Tracing_Stackdriver within kubernetes types, where deepcopy-gen is used.
func (in *Tracing_Stackdriver) DeepCopyInto(out *Tracing_Stackdriver) {
	p := proto.Clone(in).(*Tracing_Stackdriver)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tracing_Stackdriver. Required by controller-gen.
func (in *Tracing_Stackdriver) DeepCopy() *Tracing_Stackdriver {
	if in == nil {
		return nil
	}
	out := new(Tracing_Stackdriver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Tracing_Stackdriver. Required by controller-gen.
func (in *Tracing_Stackdriver) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Tracing_OpenCensusAgent within kubernetes types, where deepcopy-gen is used.
func (in *Tracing_OpenCensusAgent) DeepCopyInto(out *Tracing_OpenCensusAgent) {
	p := proto.Clone(in).(*Tracing_OpenCensusAgent)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tracing_OpenCensusAgent. Required by controller-gen.
func (in *Tracing_OpenCensusAgent) DeepCopy() *Tracing_OpenCensusAgent {
	if in == nil {
		return nil
	}
	out := new(Tracing_OpenCensusAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Tracing_OpenCensusAgent. Required by controller-gen.
func (in *Tracing_OpenCensusAgent) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Tracing_CustomTag within kubernetes types, where deepcopy-gen is used.
func (in *Tracing_CustomTag) DeepCopyInto(out *Tracing_CustomTag) {
	p := proto.Clone(in).(*Tracing_CustomTag)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tracing_CustomTag. Required by controller-gen.
func (in *Tracing_CustomTag) DeepCopy() *Tracing_CustomTag {
	if in == nil {
		return nil
	}
	out := new(Tracing_CustomTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Tracing_CustomTag. Required by controller-gen.
func (in *Tracing_CustomTag) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Tracing_Literal within kubernetes types, where deepcopy-gen is used.
func (in *Tracing_Literal) DeepCopyInto(out *Tracing_Literal) {
	p := proto.Clone(in).(*Tracing_Literal)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tracing_Literal. Required by controller-gen.
func (in *Tracing_Literal) DeepCopy() *Tracing_Literal {
	if in == nil {
		return nil
	}
	out := new(Tracing_Literal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Tracing_Literal. Required by controller-gen.
func (in *Tracing_Literal) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Tracing_Environment within kubernetes types, where deepcopy-gen is used.
func (in *Tracing_Environment) DeepCopyInto(out *Tracing_Environment) {
	p := proto.Clone(in).(*Tracing_Environment)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tracing_Environment. Required by controller-gen.
func (in *Tracing_Environment) DeepCopy() *Tracing_Environment {
	if in == nil {
		return nil
	}
	out := new(Tracing_Environment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Tracing_Environment. Required by controller-gen.
func (in *Tracing_Environment) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Tracing_RequestHeader within kubernetes types, where deepcopy-gen is used.
func (in *Tracing_RequestHeader) DeepCopyInto(out *Tracing_RequestHeader) {
	p := proto.Clone(in).(*Tracing_RequestHeader)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tracing_RequestHeader. Required by controller-gen.
func (in *Tracing_RequestHeader) DeepCopy() *Tracing_RequestHeader {
	if in == nil {
		return nil
	}
	out := new(Tracing_RequestHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Tracing_RequestHeader. Required by controller-gen.
func (in *Tracing_RequestHeader) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using SDS within kubernetes types, where deepcopy-gen is used.
func (in *SDS) DeepCopyInto(out *SDS) {
	p := proto.Clone(in).(*SDS)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SDS. Required by controller-gen.
func (in *SDS) DeepCopy() *SDS {
	if in == nil {
		return nil
	}
	out := new(SDS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new SDS. Required by controller-gen.
func (in *SDS) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Topology within kubernetes types, where deepcopy-gen is used.
func (in *Topology) DeepCopyInto(out *Topology) {
	p := proto.Clone(in).(*Topology)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Topology. Required by controller-gen.
func (in *Topology) DeepCopy() *Topology {
	if in == nil {
		return nil
	}
	out := new(Topology)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Topology. Required by controller-gen.
func (in *Topology) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ProxyConfig within kubernetes types, where deepcopy-gen is used.
func (in *ProxyConfig) DeepCopyInto(out *ProxyConfig) {
	p := proto.Clone(in).(*ProxyConfig)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyConfig. Required by controller-gen.
func (in *ProxyConfig) DeepCopy() *ProxyConfig {
	if in == nil {
		return nil
	}
	out := new(ProxyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ProxyConfig. Required by controller-gen.
func (in *ProxyConfig) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RemoteService within kubernetes types, where deepcopy-gen is used.
func (in *RemoteService) DeepCopyInto(out *RemoteService) {
	p := proto.Clone(in).(*RemoteService)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteService. Required by controller-gen.
func (in *RemoteService) DeepCopy() *RemoteService {
	if in == nil {
		return nil
	}
	out := new(RemoteService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RemoteService. Required by controller-gen.
func (in *RemoteService) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
