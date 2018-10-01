# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: audit/v1alpha1/audit.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='audit/v1alpha1/audit.proto',
  package='istio.audit.v1alpha1',
  syntax='proto3',
  serialized_pb=_b('\n\x1a\x61udit/v1alpha1/audit.proto\x12\x14istio.audit.v1alpha1\"A\n\x0e\x41uditDirective\x12/\n\x05level\x18\x01 \x01(\x0e\x32 .istio.audit.v1alpha1.AuditLevel*5\n\nAuditLevel\x12\x08\n\x04NONE\x10\x00\x12\x0c\n\x08METADATA\x10\x01\x12\x0f\n\x0b\x41PP_DEFINED\x10\x02\x42\x1dZ\x1bistio.io/api/audit/v1alpha1b\x06proto3')
)

_AUDITLEVEL = _descriptor.EnumDescriptor(
  name='AuditLevel',
  full_name='istio.audit.v1alpha1.AuditLevel',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='METADATA', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='APP_DEFINED', index=2, number=2,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=119,
  serialized_end=172,
)
_sym_db.RegisterEnumDescriptor(_AUDITLEVEL)

AuditLevel = enum_type_wrapper.EnumTypeWrapper(_AUDITLEVEL)
NONE = 0
METADATA = 1
APP_DEFINED = 2



_AUDITDIRECTIVE = _descriptor.Descriptor(
  name='AuditDirective',
  full_name='istio.audit.v1alpha1.AuditDirective',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='level', full_name='istio.audit.v1alpha1.AuditDirective.level', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=52,
  serialized_end=117,
)

_AUDITDIRECTIVE.fields_by_name['level'].enum_type = _AUDITLEVEL
DESCRIPTOR.message_types_by_name['AuditDirective'] = _AUDITDIRECTIVE
DESCRIPTOR.enum_types_by_name['AuditLevel'] = _AUDITLEVEL
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

AuditDirective = _reflection.GeneratedProtocolMessageType('AuditDirective', (_message.Message,), dict(
  DESCRIPTOR = _AUDITDIRECTIVE,
  __module__ = 'audit.v1alpha1.audit_pb2'
  # @@protoc_insertion_point(class_scope:istio.audit.v1alpha1.AuditDirective)
  ))
_sym_db.RegisterMessage(AuditDirective)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z\033istio.io/api/audit/v1alpha1'))
# @@protoc_insertion_point(module_scope)
