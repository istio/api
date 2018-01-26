// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/v1/config/cfg.proto

/*
Package config is a generated protocol buffer package.

It is generated from these files:
	mixer/v1/config/cfg.proto

It has these top-level messages:
	AttributeManifest
	Rule
	Action
	Instance
	Handler
*/
package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/struct"
import istio_mixer_v1_config_descriptor "istio.io/api/mixer/v1/config/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// AttributeManifest describes a set of Attributes produced by some component
// of an Istio deployment.
type AttributeManifest struct {
	// Optional. The revision of this document. Assigned by server.
	Revision string `protobuf:"bytes,1,opt,name=revision,proto3" json:"revision,omitempty"`
	// Required. Name of the component producing these attributes. This can be
	// the proxy (with the canonical name "istio-proxy") or the name of an
	// `attributes` kind adapter in Mixer.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The set of attributes this Istio component will be responsible for producing at runtime.
	// We map from attribute name to the attribute's specification. The name of an attribute,
	// which is how attributes are referred to in aspect configuration, must conform to:
	//
	//     Name = IDENT { SEPARATOR IDENT };
	//
	// Where `IDENT` must match the regular expression `[a-z][a-z0-9]+` and `SEPARATOR` must
	// match the regular expression `[\.-]`.
	//
	// Attribute names must be unique within a single Istio deployment. The set of canonical
	// attributes are described at https://istio.io/docs/reference/attribute-vocabulary.html.
	// Attributes not in that list should be named with a component-specific suffix such as
	// request.count-my.component
	Attributes map[string]*AttributeManifest_AttributeInfo `protobuf:"bytes,3,rep,name=attributes" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *AttributeManifest) Reset()                    { *m = AttributeManifest{} }
func (m *AttributeManifest) String() string            { return proto.CompactTextString(m) }
func (*AttributeManifest) ProtoMessage()               {}
func (*AttributeManifest) Descriptor() ([]byte, []int) { return fileDescriptorCfg, []int{0} }

func (m *AttributeManifest) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *AttributeManifest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AttributeManifest) GetAttributes() map[string]*AttributeManifest_AttributeInfo {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// AttributeInfo describes the schema of an Istio `Attribute`.
//
// ## Istio Attributes
//
// Istio uses `attributes` to describe runtime activities of Istio services.
// An Istio attribute carries a specific piece of information about an activity,
// such as the error code of an API request, the latency of an API request, or the
// original IP address of a TCP connection. The attributes are often generated
// and consumed by different services. For example, a frontend service can
// generate an authenticated user attribute and pass it to a backend service for
// access control purpose.
//
// To simplify the system and improve developer experience, Istio uses
// shared attribute definitions across all components. For example, the same
// authenticated user attribute will be used for logging, monitoring, analytics,
// billing, access control, auditing. Many Istio components provide their
// functionality by collecting, generating, and operating on attributes.
// For example, the proxy collects the error code attribute, and the logging
// stores it into a log.
//
// ## Design
//
// Each Istio attribute must conform to an `AttributeInfo` in an
// `AttributeManifest` in the current Istio deployment at runtime. An
// [`AttributeInfo`][istio.mixer.v1.config] is used to define an attribute's
// metadata: the type of its value and a detailed description that explains
// the semantics of the attribute type. Each attribute's name is globally unique;
// in other words an attribute name can only appear once across all manifests.
//
// The runtime presentation of an attribute is intentionally left out of this
// specification, because passing attribute using JSON, XML, or Protocol Buffers
// does not change the semantics of the attribute. Different implementations
// can choose different representations based on their needs.
//
// ## HTTP Mapping
//
// Because many systems already have REST APIs, it makes sense to define a
// standard HTTP mapping for Istio attributes that are compatible with typical
// REST APIs. The design is to map one attribute to one HTTP header, the
// attribute name and value becomes the HTTP header name and value. The actual
// encoding scheme will be decided later.
type AttributeManifest_AttributeInfo struct {
	// Optional. A human-readable description of the attribute's purpose.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// Required. The type of data carried by this attribute.
	ValueType istio_mixer_v1_config_descriptor.ValueType `protobuf:"varint,2,opt,name=value_type,json=valueType,proto3,enum=istio.mixer.v1.config.descriptor.ValueType" json:"value_type,omitempty"`
}

func (m *AttributeManifest_AttributeInfo) Reset()         { *m = AttributeManifest_AttributeInfo{} }
func (m *AttributeManifest_AttributeInfo) String() string { return proto.CompactTextString(m) }
func (*AttributeManifest_AttributeInfo) ProtoMessage()    {}
func (*AttributeManifest_AttributeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptorCfg, []int{0, 0}
}

func (m *AttributeManifest_AttributeInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AttributeManifest_AttributeInfo) GetValueType() istio_mixer_v1_config_descriptor.ValueType {
	if m != nil {
		return m.ValueType
	}
	return istio_mixer_v1_config_descriptor.VALUE_TYPE_UNSPECIFIED
}

// A Rule is a selector and a set of intentions to be executed when the
// selector is `true`
//
// The following example instructs Mixer to invoke 'prometheus-handler' handler for all services and pass it the
// instance constructed using the 'RequestCountByService' instance.
//
// ```yaml
// - match: destination.service == "*"
//   actions:
//   - handler: prometheus-handler
//     instances:
//     - RequestCountByService
// ```
type Rule struct {
	// Required. Match is an attribute based predicate. When Mixer receives a
	// request it evaluates the match expression and executes all the associated `actions`
	// if the match evaluates to true.
	//
	// A few example match:
	//
	// * an empty match evaluates to `true`
	// * `true`, a boolean literal; a rule with this match will always be executed
	// * `destination.service == ratings*` selects any request targeting a service whose
	// name starts with "ratings"
	// * `attr1 == "20" && attr2 == "30"` logical AND, OR, and NOT are also available
	Match string `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	// Optional. The actions that will be executed when match evaluates to `true`.
	Actions []*Action `protobuf:"bytes,2,rep,name=actions" json:"actions,omitempty"`
}

func (m *Rule) Reset()                    { *m = Rule{} }
func (m *Rule) String() string            { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()               {}
func (*Rule) Descriptor() ([]byte, []int) { return fileDescriptorCfg, []int{1} }

func (m *Rule) GetMatch() string {
	if m != nil {
		return m.Match
	}
	return ""
}

func (m *Rule) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

// Action describes which [Handler][istio.mixer.v1.config.Handler] to invoke and what data to pass to it for processing.
//
// The following example instructs Mixer to invoke 'prometheus-handler' handler and pass it the object
// constructed using the instance 'RequestCountByService'
//
// ```yaml
//   handler: prometheus-handler
//   instances:
//   - RequestCountByService
// ```
type Action struct {
	// Required. Fully qualified name of the handler to invoke.
	// Must match the `name` of a [Handler][istio.mixer.v1.config.Handler.name].
	Handler string `protobuf:"bytes,2,opt,name=handler,proto3" json:"handler,omitempty"`
	// Required. Each value must match the fully qualified name of the
	// [Instance][istio.mixer.v1.config.Instance.name]s.
	// Referenced instances are evaluated by resolving the attributes/literals for all the fields.
	// The constructed objects are then passed to the `handler` referenced within this action.
	Instances []string `protobuf:"bytes,3,rep,name=instances" json:"instances,omitempty"`
}

func (m *Action) Reset()                    { *m = Action{} }
func (m *Action) String() string            { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()               {}
func (*Action) Descriptor() ([]byte, []int) { return fileDescriptorCfg, []int{2} }

func (m *Action) GetHandler() string {
	if m != nil {
		return m.Handler
	}
	return ""
}

func (m *Action) GetInstances() []string {
	if m != nil {
		return m.Instances
	}
	return nil
}

// An Instance tells Mixer how to create instances for particular template.
//
// Instance is defined by the operator. Instance is defined relative to a known
// template. Their purpose is to tell Mixer how to use attributes or literals to produce
// instances of the specified template at runtime.
//
// The following example instructs Mixer to construct an instance associated with template
// 'istio.mixer.adapter.metric.Metric'. It provides a mapping from the template's fields to expressions.
// Instances produced with this instance can be referenced by [Actions][istio.mixer.v1.config.Action] using name
// 'RequestCountByService'.
//
// ```yaml
// - name: RequestCountByService
//   template: istio.mixer.adapter.metric.Metric
//   params:
//     value: 1
//     dimensions:
//       source: source.service
//       destination_ip: destination.ip
// ```
type Instance struct {
	// Required. The name of this instance
	//
	// Must be unique amongst other Instances in scope. Used by [Action][istio.mixer.v1.config.Action] to refer
	// to an instance produced by this instance.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The name of the template this instance creates instances for.
	// The value must match the name of the available template in scope.
	Template string `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	// Required. Depends on referenced template. Struct representation of a
	// proto defined by the template; this varies depending on the value of field `template`.
	Params interface{} `protobuf:"bytes,3,opt,name=params" json:"params,omitempty"`
}

func (m *Instance) Reset()                    { *m = Instance{} }
func (m *Instance) String() string            { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()               {}
func (*Instance) Descriptor() ([]byte, []int) { return fileDescriptorCfg, []int{3} }

func (m *Instance) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Instance) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *Instance) GetParams() interface{} {
	if m != nil {
		return m.Params
	}
	return nil
}

// Handler allows the operator to configure a specific adapter implementation.
// Each adapter implementation defines its own `params` proto.
//
// In the following example we define a `metrics` handler using the Mixer's prepackaged
// prometheus adapter. This handler doesn't require any parameters.
//
// ```yaml
// name: prometheus-handler
// adapter: prometheus
// params:
// ```
type Handler struct {
	// Required. Must be unique in the entire mixer configuration. Used by [Actions][istio.mixer.v1.config.Action.handler]
	// to refer to this handler.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The name of a specific adapter implementation. An adapter's
	// implementation name is typically a constant in its code.
	Adapter string `protobuf:"bytes,2,opt,name=adapter,proto3" json:"adapter,omitempty"`
	// Optional. Depends on adapter implementation. Struct representation of a
	// proto defined by the adapter implementation; this varies depending on the value of field `adapter`.
	Params interface{} `protobuf:"bytes,3,opt,name=params" json:"params,omitempty"`
}

func (m *Handler) Reset()                    { *m = Handler{} }
func (m *Handler) String() string            { return proto.CompactTextString(m) }
func (*Handler) ProtoMessage()               {}
func (*Handler) Descriptor() ([]byte, []int) { return fileDescriptorCfg, []int{4} }

func (m *Handler) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Handler) GetAdapter() string {
	if m != nil {
		return m.Adapter
	}
	return ""
}

func (m *Handler) GetParams() interface{} {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*AttributeManifest)(nil), "istio.mixer.v1.config.AttributeManifest")
	proto.RegisterType((*AttributeManifest_AttributeInfo)(nil), "istio.mixer.v1.config.AttributeManifest.AttributeInfo")
	proto.RegisterType((*Rule)(nil), "istio.mixer.v1.config.Rule")
	proto.RegisterType((*Action)(nil), "istio.mixer.v1.config.Action")
	proto.RegisterType((*Instance)(nil), "istio.mixer.v1.config.Instance")
	proto.RegisterType((*Handler)(nil), "istio.mixer.v1.config.Handler")
}

func init() { proto.RegisterFile("mixer/v1/config/cfg.proto", fileDescriptorCfg) }

var fileDescriptorCfg = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x8b, 0xdb, 0x30,
	0x10, 0x25, 0xc9, 0x6e, 0x3e, 0x26, 0xf4, 0x4b, 0xb4, 0xd4, 0x35, 0x69, 0x09, 0x3e, 0x2d, 0x2c,
	0x48, 0x6c, 0x0a, 0xed, 0xd2, 0x53, 0x5b, 0x28, 0x74, 0x4b, 0x7b, 0x71, 0x3f, 0x28, 0xbd, 0x14,
	0xc5, 0x91, 0x13, 0xb1, 0xb6, 0x64, 0xa4, 0xb1, 0x69, 0x0e, 0xfd, 0x15, 0xfd, 0xc3, 0xc5, 0xb2,
	0x1c, 0x87, 0xec, 0xee, 0x21, 0xb7, 0x99, 0xf1, 0x9b, 0x79, 0xf3, 0x46, 0xcf, 0xf0, 0x2c, 0x97,
	0x7f, 0x84, 0x61, 0xd5, 0x05, 0x4b, 0xb4, 0x4a, 0xe5, 0x9a, 0x25, 0xe9, 0x9a, 0x16, 0x46, 0xa3,
	0x26, 0x4f, 0xa4, 0x45, 0xa9, 0xa9, 0x03, 0xd0, 0xea, 0x82, 0x36, 0x80, 0x70, 0xb6, 0xd6, 0x7a,
	0x9d, 0x09, 0xe6, 0x40, 0xcb, 0x32, 0x65, 0x16, 0x4d, 0x99, 0x60, 0xd3, 0x14, 0x9e, 0x1f, 0xce,
	0x5b, 0x09, 0x9b, 0x18, 0x59, 0xa0, 0x36, 0xac, 0xe2, 0x59, 0x29, 0x7e, 0xe3, 0xb6, 0x10, 0x0d,
	0x38, 0xfa, 0x37, 0x80, 0x47, 0xef, 0x10, 0x8d, 0x5c, 0x96, 0x28, 0xbe, 0x70, 0x25, 0x53, 0x61,
	0x91, 0x84, 0x30, 0x36, 0xa2, 0x92, 0x56, 0x6a, 0x15, 0xf4, 0xe6, 0xbd, 0xb3, 0x49, 0xbc, 0xcb,
	0x09, 0x81, 0x13, 0xc5, 0x73, 0x11, 0xf4, 0x5d, 0xdd, 0xc5, 0xe4, 0x27, 0x00, 0x6f, 0x87, 0xd8,
	0x60, 0x30, 0x1f, 0x9c, 0x4d, 0x17, 0x97, 0xf4, 0xd6, 0xe5, 0xe9, 0x0d, 0xb6, 0xae, 0x62, 0x3f,
	0x28, 0x34, 0xdb, 0x78, 0x6f, 0x56, 0xf8, 0x17, 0xee, 0xed, 0x3e, 0x5f, 0xa9, 0x54, 0x93, 0x39,
	0x4c, 0x5b, 0x3d, 0xdd, 0x76, 0xfb, 0x25, 0xf2, 0x09, 0xa0, 0x93, 0xe9, 0xd6, 0xbc, 0xbf, 0x38,
	0xbf, 0x63, 0x99, 0xee, 0x34, 0xf4, 0x47, 0xdd, 0xf3, 0x6d, 0x5b, 0x88, 0x78, 0x52, 0xb5, 0x61,
	0x58, 0xc2, 0x83, 0x83, 0xed, 0xc8, 0x43, 0x18, 0x5c, 0x8b, 0xad, 0x27, 0xae, 0x43, 0xf2, 0x19,
	0x4e, 0x5d, 0x87, 0xe3, 0x9a, 0x2e, 0x5e, 0x1d, 0x2f, 0xbc, 0x56, 0x16, 0x37, 0x43, 0xde, 0xf4,
	0x2f, 0x7b, 0xd1, 0x77, 0x38, 0x89, 0xcb, 0x4c, 0x90, 0xc7, 0x70, 0x9a, 0x73, 0x4c, 0x36, 0x9e,
	0xad, 0x49, 0xc8, 0x6b, 0x18, 0xf1, 0xa4, 0x96, 0x6a, 0x83, 0xbe, 0x3b, 0xf5, 0xf3, 0xbb, 0x18,
	0x1d, 0x2a, 0x6e, 0xd1, 0xd1, 0x5b, 0x18, 0x36, 0x25, 0x12, 0xc0, 0x68, 0xc3, 0xd5, 0x2a, 0x13,
	0xc6, 0xbf, 0x63, 0x9b, 0x92, 0x19, 0x4c, 0xa4, 0xb2, 0xc8, 0x55, 0xe2, 0x5f, 0x72, 0x12, 0x77,
	0x85, 0xe8, 0x1a, 0xc6, 0x57, 0x3e, 0xd9, 0x19, 0xa1, 0xb7, 0x67, 0x84, 0x10, 0xc6, 0x28, 0xf2,
	0x22, 0xe3, 0xd8, 0x1a, 0x64, 0x97, 0x13, 0x06, 0xc3, 0x82, 0x1b, 0x9e, 0xd7, 0x63, 0xeb, 0x3b,
	0x3d, 0xa5, 0x8d, 0x8d, 0x69, 0x6b, 0x63, 0xfa, 0xd5, 0xd9, 0x38, 0xf6, 0xb0, 0x68, 0x03, 0xa3,
	0x8f, 0x7e, 0xab, 0xdb, 0xb8, 0x02, 0x18, 0xf1, 0x15, 0x2f, 0xb0, 0xd3, 0xe0, 0xd3, 0xa3, 0x99,
	0xde, 0xbf, 0xf8, 0x35, 0x6b, 0x2e, 0x28, 0x35, 0xe3, 0x85, 0x64, 0x07, 0x7f, 0xd0, 0x72, 0xe8,
	0x1a, 0x5f, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xc3, 0x04, 0x7d, 0xab, 0x03, 0x00, 0x00,
}
