# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: envoy/types/percent.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='envoy/types/percent.proto',
  package='istio.envoy.types',
  syntax='proto3',
  serialized_pb=_b('\n\x19\x65nvoy/types/percent.proto\x12\x11istio.envoy.types\x1a\x14gogoproto/gogo.proto\"\x18\n\x07Percent\x12\r\n\x05value\x18\x01 \x01(\x01\x42\x1eZ\x18istio.io/api/envoy/types\xa8\xe2\x1e\x01\x62\x06proto3')
  ,
  dependencies=[gogoproto_dot_gogo__pb2.DESCRIPTOR,])




_PERCENT = _descriptor.Descriptor(
  name='Percent',
  full_name='istio.envoy.types.Percent',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.envoy.types.Percent.value', index=0,
      number=1, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=70,
  serialized_end=94,
)

DESCRIPTOR.message_types_by_name['Percent'] = _PERCENT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Percent = _reflection.GeneratedProtocolMessageType('Percent', (_message.Message,), dict(
  DESCRIPTOR = _PERCENT,
  __module__ = 'envoy.types.percent_pb2'
  # @@protoc_insertion_point(class_scope:istio.envoy.types.Percent)
  ))
_sym_db.RegisterMessage(Percent)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z\030istio.io/api/envoy/types\250\342\036\001'))
# @@protoc_insertion_point(module_scope)
