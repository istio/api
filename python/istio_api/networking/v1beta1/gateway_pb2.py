# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: networking/v1beta1/gateway.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import field_behavior_pb2 as google_dot_api_dot_field__behavior__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='networking/v1beta1/gateway.proto',
  package='istio.networking.v1beta1',
  syntax='proto3',
  serialized_options=_b('Z\037istio.io/api/networking/v1beta1'),
  serialized_pb=_b('\n networking/v1beta1/gateway.proto\x12\x18istio.networking.v1beta1\x1a\x1fgoogle/api/field_behavior.proto\"\xba\x01\n\x07Gateway\x12\x36\n\x07servers\x18\x01 \x03(\x0b\x32 .istio.networking.v1beta1.ServerB\x03\xe0\x41\x02\x12\x46\n\x08selector\x18\x02 \x03(\x0b\x32/.istio.networking.v1beta1.Gateway.SelectorEntryB\x03\xe0\x41\x02\x1a/\n\rSelectorEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xb1\x01\n\x06Server\x12\x31\n\x04port\x18\x01 \x01(\x0b\x32\x1e.istio.networking.v1beta1.PortB\x03\xe0\x41\x02\x12\x0c\n\x04\x62ind\x18\x04 \x01(\t\x12\x12\n\x05hosts\x18\x02 \x03(\tB\x03\xe0\x41\x02\x12\x38\n\x03tls\x18\x03 \x01(\x0b\x32+.istio.networking.v1beta1.ServerTLSSettings\x12\x18\n\x10\x64\x65\x66\x61ult_endpoint\x18\x05 \x01(\t\"E\n\x04Port\x12\x13\n\x06number\x18\x01 \x01(\rB\x03\xe0\x41\x02\x12\x15\n\x08protocol\x18\x02 \x01(\tB\x03\xe0\x41\x02\x12\x11\n\x04name\x18\x03 \x01(\tB\x03\xe0\x41\x02\"\xa0\x05\n\x11ServerTLSSettings\x12\x16\n\x0ehttps_redirect\x18\x01 \x01(\x08\x12\x41\n\x04mode\x18\x02 \x01(\x0e\x32\x33.istio.networking.v1beta1.ServerTLSSettings.TLSmode\x12\x1a\n\x12server_certificate\x18\x03 \x01(\t\x12\x13\n\x0bprivate_key\x18\x04 \x01(\t\x12\x17\n\x0f\x63\x61_certificates\x18\x05 \x01(\t\x12\x17\n\x0f\x63redential_name\x18\n \x01(\t\x12\x19\n\x11subject_alt_names\x18\x06 \x03(\t\x12\x1f\n\x17verify_certificate_spki\x18\x0b \x03(\t\x12\x1f\n\x17verify_certificate_hash\x18\x0c \x03(\t\x12U\n\x14min_protocol_version\x18\x07 \x01(\x0e\x32\x37.istio.networking.v1beta1.ServerTLSSettings.TLSProtocol\x12U\n\x14max_protocol_version\x18\x08 \x01(\x0e\x32\x37.istio.networking.v1beta1.ServerTLSSettings.TLSProtocol\x12\x15\n\rcipher_suites\x18\t \x03(\t\"Z\n\x07TLSmode\x12\x0f\n\x0bPASSTHROUGH\x10\x00\x12\n\n\x06SIMPLE\x10\x01\x12\n\n\x06MUTUAL\x10\x02\x12\x14\n\x10\x41UTO_PASSTHROUGH\x10\x03\x12\x10\n\x0cISTIO_MUTUAL\x10\x04\"O\n\x0bTLSProtocol\x12\x0c\n\x08TLS_AUTO\x10\x00\x12\x0b\n\x07TLSV1_0\x10\x01\x12\x0b\n\x07TLSV1_1\x10\x02\x12\x0b\n\x07TLSV1_2\x10\x03\x12\x0b\n\x07TLSV1_3\x10\x04\x42!Z\x1fistio.io/api/networking/v1beta1b\x06proto3')
  ,
  dependencies=[google_dot_api_dot_field__behavior__pb2.DESCRIPTOR,])



_SERVERTLSSETTINGS_TLSMODE = _descriptor.EnumDescriptor(
  name='TLSmode',
  full_name='istio.networking.v1beta1.ServerTLSSettings.TLSmode',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='PASSTHROUGH', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SIMPLE', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MUTUAL', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AUTO_PASSTHROUGH', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ISTIO_MUTUAL', index=4, number=4,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1037,
  serialized_end=1127,
)
_sym_db.RegisterEnumDescriptor(_SERVERTLSSETTINGS_TLSMODE)

_SERVERTLSSETTINGS_TLSPROTOCOL = _descriptor.EnumDescriptor(
  name='TLSProtocol',
  full_name='istio.networking.v1beta1.ServerTLSSettings.TLSProtocol',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='TLS_AUTO', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TLSV1_0', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TLSV1_1', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TLSV1_2', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TLSV1_3', index=4, number=4,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1129,
  serialized_end=1208,
)
_sym_db.RegisterEnumDescriptor(_SERVERTLSSETTINGS_TLSPROTOCOL)


_GATEWAY_SELECTORENTRY = _descriptor.Descriptor(
  name='SelectorEntry',
  full_name='istio.networking.v1beta1.Gateway.SelectorEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.networking.v1beta1.Gateway.SelectorEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.networking.v1beta1.Gateway.SelectorEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=235,
  serialized_end=282,
)

_GATEWAY = _descriptor.Descriptor(
  name='Gateway',
  full_name='istio.networking.v1beta1.Gateway',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='servers', full_name='istio.networking.v1beta1.Gateway.servers', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\340A\002'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='selector', full_name='istio.networking.v1beta1.Gateway.selector', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\340A\002'), file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_GATEWAY_SELECTORENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=96,
  serialized_end=282,
)


_SERVER = _descriptor.Descriptor(
  name='Server',
  full_name='istio.networking.v1beta1.Server',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='port', full_name='istio.networking.v1beta1.Server.port', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\340A\002'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='bind', full_name='istio.networking.v1beta1.Server.bind', index=1,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='hosts', full_name='istio.networking.v1beta1.Server.hosts', index=2,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\340A\002'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tls', full_name='istio.networking.v1beta1.Server.tls', index=3,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='default_endpoint', full_name='istio.networking.v1beta1.Server.default_endpoint', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=285,
  serialized_end=462,
)


_PORT = _descriptor.Descriptor(
  name='Port',
  full_name='istio.networking.v1beta1.Port',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='number', full_name='istio.networking.v1beta1.Port.number', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\340A\002'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='protocol', full_name='istio.networking.v1beta1.Port.protocol', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\340A\002'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='istio.networking.v1beta1.Port.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\340A\002'), file=DESCRIPTOR),
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
  serialized_start=464,
  serialized_end=533,
)


_SERVERTLSSETTINGS = _descriptor.Descriptor(
  name='ServerTLSSettings',
  full_name='istio.networking.v1beta1.ServerTLSSettings',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='https_redirect', full_name='istio.networking.v1beta1.ServerTLSSettings.https_redirect', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='mode', full_name='istio.networking.v1beta1.ServerTLSSettings.mode', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='server_certificate', full_name='istio.networking.v1beta1.ServerTLSSettings.server_certificate', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='private_key', full_name='istio.networking.v1beta1.ServerTLSSettings.private_key', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ca_certificates', full_name='istio.networking.v1beta1.ServerTLSSettings.ca_certificates', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='credential_name', full_name='istio.networking.v1beta1.ServerTLSSettings.credential_name', index=5,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='subject_alt_names', full_name='istio.networking.v1beta1.ServerTLSSettings.subject_alt_names', index=6,
      number=6, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='verify_certificate_spki', full_name='istio.networking.v1beta1.ServerTLSSettings.verify_certificate_spki', index=7,
      number=11, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='verify_certificate_hash', full_name='istio.networking.v1beta1.ServerTLSSettings.verify_certificate_hash', index=8,
      number=12, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='min_protocol_version', full_name='istio.networking.v1beta1.ServerTLSSettings.min_protocol_version', index=9,
      number=7, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='max_protocol_version', full_name='istio.networking.v1beta1.ServerTLSSettings.max_protocol_version', index=10,
      number=8, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='cipher_suites', full_name='istio.networking.v1beta1.ServerTLSSettings.cipher_suites', index=11,
      number=9, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _SERVERTLSSETTINGS_TLSMODE,
    _SERVERTLSSETTINGS_TLSPROTOCOL,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=536,
  serialized_end=1208,
)

_GATEWAY_SELECTORENTRY.containing_type = _GATEWAY
_GATEWAY.fields_by_name['servers'].message_type = _SERVER
_GATEWAY.fields_by_name['selector'].message_type = _GATEWAY_SELECTORENTRY
_SERVER.fields_by_name['port'].message_type = _PORT
_SERVER.fields_by_name['tls'].message_type = _SERVERTLSSETTINGS
_SERVERTLSSETTINGS.fields_by_name['mode'].enum_type = _SERVERTLSSETTINGS_TLSMODE
_SERVERTLSSETTINGS.fields_by_name['min_protocol_version'].enum_type = _SERVERTLSSETTINGS_TLSPROTOCOL
_SERVERTLSSETTINGS.fields_by_name['max_protocol_version'].enum_type = _SERVERTLSSETTINGS_TLSPROTOCOL
_SERVERTLSSETTINGS_TLSMODE.containing_type = _SERVERTLSSETTINGS
_SERVERTLSSETTINGS_TLSPROTOCOL.containing_type = _SERVERTLSSETTINGS
DESCRIPTOR.message_types_by_name['Gateway'] = _GATEWAY
DESCRIPTOR.message_types_by_name['Server'] = _SERVER
DESCRIPTOR.message_types_by_name['Port'] = _PORT
DESCRIPTOR.message_types_by_name['ServerTLSSettings'] = _SERVERTLSSETTINGS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Gateway = _reflection.GeneratedProtocolMessageType('Gateway', (_message.Message,), {

  'SelectorEntry' : _reflection.GeneratedProtocolMessageType('SelectorEntry', (_message.Message,), {
    'DESCRIPTOR' : _GATEWAY_SELECTORENTRY,
    '__module__' : 'networking.v1beta1.gateway_pb2'
    # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.Gateway.SelectorEntry)
    })
  ,
  'DESCRIPTOR' : _GATEWAY,
  '__module__' : 'networking.v1beta1.gateway_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.Gateway)
  })
_sym_db.RegisterMessage(Gateway)
_sym_db.RegisterMessage(Gateway.SelectorEntry)

Server = _reflection.GeneratedProtocolMessageType('Server', (_message.Message,), {
  'DESCRIPTOR' : _SERVER,
  '__module__' : 'networking.v1beta1.gateway_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.Server)
  })
_sym_db.RegisterMessage(Server)

Port = _reflection.GeneratedProtocolMessageType('Port', (_message.Message,), {
  'DESCRIPTOR' : _PORT,
  '__module__' : 'networking.v1beta1.gateway_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.Port)
  })
_sym_db.RegisterMessage(Port)

ServerTLSSettings = _reflection.GeneratedProtocolMessageType('ServerTLSSettings', (_message.Message,), {
  'DESCRIPTOR' : _SERVERTLSSETTINGS,
  '__module__' : 'networking.v1beta1.gateway_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.ServerTLSSettings)
  })
_sym_db.RegisterMessage(ServerTLSSettings)


DESCRIPTOR._options = None
_GATEWAY_SELECTORENTRY._options = None
_GATEWAY.fields_by_name['servers']._options = None
_GATEWAY.fields_by_name['selector']._options = None
_SERVER.fields_by_name['port']._options = None
_SERVER.fields_by_name['hosts']._options = None
_PORT.fields_by_name['number']._options = None
_PORT.fields_by_name['protocol']._options = None
_PORT.fields_by_name['name']._options = None
# @@protoc_insertion_point(module_scope)
