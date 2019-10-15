# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: envoy/config/filter/http/alpn/v2alpha1/config.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='envoy/config/filter/http/alpn/v2alpha1/config.proto',
  package='istio.envoy.config.filter.http.alpn.v2alpha1',
  syntax='proto3',
  serialized_options=_b('Z3istio.io/api/envoy/config/filter/http/alpn/v2alpha1'),
  serialized_pb=_b('\n3envoy/config/filter/http/alpn/v2alpha1/config.proto\x12,istio.envoy.config.filter.http.alpn.v2alpha1\"\xce\x02\n\x0c\x46ilterConfig\x12\x19\n\ralpn_override\x18\x01 \x03(\tB\x02\x18\x01\x12k\n\x1a\x61lpn_override_per_protocol\x18\x02 \x03(\x0b\x32G.istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.AlpnOverride\x1a\x86\x01\n\x0c\x41lpnOverride\x12^\n\x11upstream_protocol\x18\x01 \x01(\x0e\x32\x43.istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.Protocol\x12\x16\n\x0e\x61lpn_overrides\x18\x02 \x03(\t\"-\n\x08Protocol\x12\n\n\x06HTTP10\x10\x00\x12\n\n\x06HTTP11\x10\x01\x12\t\n\x05HTTP2\x10\x02\x42\x35Z3istio.io/api/envoy/config/filter/http/alpn/v2alpha1b\x06proto3')
)



_FILTERCONFIG_PROTOCOL = _descriptor.EnumDescriptor(
  name='Protocol',
  full_name='istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.Protocol',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='HTTP10', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HTTP11', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HTTP2', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=391,
  serialized_end=436,
)
_sym_db.RegisterEnumDescriptor(_FILTERCONFIG_PROTOCOL)


_FILTERCONFIG_ALPNOVERRIDE = _descriptor.Descriptor(
  name='AlpnOverride',
  full_name='istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.AlpnOverride',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='upstream_protocol', full_name='istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.AlpnOverride.upstream_protocol', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='alpn_overrides', full_name='istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.AlpnOverride.alpn_overrides', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=255,
  serialized_end=389,
)

_FILTERCONFIG = _descriptor.Descriptor(
  name='FilterConfig',
  full_name='istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='alpn_override', full_name='istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.alpn_override', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\030\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='alpn_override_per_protocol', full_name='istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.alpn_override_per_protocol', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_FILTERCONFIG_ALPNOVERRIDE, ],
  enum_types=[
    _FILTERCONFIG_PROTOCOL,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=102,
  serialized_end=436,
)

_FILTERCONFIG_ALPNOVERRIDE.fields_by_name['upstream_protocol'].enum_type = _FILTERCONFIG_PROTOCOL
_FILTERCONFIG_ALPNOVERRIDE.containing_type = _FILTERCONFIG
_FILTERCONFIG.fields_by_name['alpn_override_per_protocol'].message_type = _FILTERCONFIG_ALPNOVERRIDE
_FILTERCONFIG_PROTOCOL.containing_type = _FILTERCONFIG
DESCRIPTOR.message_types_by_name['FilterConfig'] = _FILTERCONFIG
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

FilterConfig = _reflection.GeneratedProtocolMessageType('FilterConfig', (_message.Message,), {

  'AlpnOverride' : _reflection.GeneratedProtocolMessageType('AlpnOverride', (_message.Message,), {
    'DESCRIPTOR' : _FILTERCONFIG_ALPNOVERRIDE,
    '__module__' : 'envoy.config.filter.http.alpn.v2alpha1.config_pb2'
    # @@protoc_insertion_point(class_scope:istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig.AlpnOverride)
    })
  ,
  'DESCRIPTOR' : _FILTERCONFIG,
  '__module__' : 'envoy.config.filter.http.alpn.v2alpha1.config_pb2'
  # @@protoc_insertion_point(class_scope:istio.envoy.config.filter.http.alpn.v2alpha1.FilterConfig)
  })
_sym_db.RegisterMessage(FilterConfig)
_sym_db.RegisterMessage(FilterConfig.AlpnOverride)


DESCRIPTOR._options = None
_FILTERCONFIG.fields_by_name['alpn_override']._options = None
# @@protoc_insertion_point(module_scope)
