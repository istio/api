# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: mixer/adapter/model/v1beta1/info.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='mixer/adapter/model/v1beta1/info.proto',
  package='istio.mixer.adapter.model.v1beta1',
  syntax='proto3',
  serialized_pb=_b('\n&mixer/adapter/model/v1beta1/info.proto\x12!istio.mixer.adapter.model.v1beta1\"L\n\x04Info\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x02 \x01(\t\x12\x11\n\ttemplates\x18\x03 \x03(\t\x12\x0e\n\x06\x63onfig\x18\x04 \x01(\tB*Z(istio.io/api/mixer/adapter/model/v1beta1b\x06proto3')
)




_INFO = _descriptor.Descriptor(
  name='Info',
  full_name='istio.mixer.adapter.model.v1beta1.Info',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='istio.mixer.adapter.model.v1beta1.Info.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='istio.mixer.adapter.model.v1beta1.Info.description', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='templates', full_name='istio.mixer.adapter.model.v1beta1.Info.templates', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='config', full_name='istio.mixer.adapter.model.v1beta1.Info.config', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=77,
  serialized_end=153,
)

DESCRIPTOR.message_types_by_name['Info'] = _INFO
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Info = _reflection.GeneratedProtocolMessageType('Info', (_message.Message,), dict(
  DESCRIPTOR = _INFO,
  __module__ = 'mixer.adapter.model.v1beta1.info_pb2'
  # @@protoc_insertion_point(class_scope:istio.mixer.adapter.model.v1beta1.Info)
  ))
_sym_db.RegisterMessage(Info)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z(istio.io/api/mixer/adapter/model/v1beta1'))
# @@protoc_insertion_point(module_scope)
