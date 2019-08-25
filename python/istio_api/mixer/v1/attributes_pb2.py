# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: mixer/v1/attributes.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='mixer/v1/attributes.proto',
  package='istio.mixer.v1',
  syntax='proto3',
  serialized_options=_b('Z\025istio.io/api/mixer/v1\370\001\001\310\341\036\000\250\342\036\000\360\341\036\000\330\342\036\001'),
  serialized_pb=_b('\n\x19mixer/v1/attributes.proto\x12\x0eistio.mixer.v1\x1a\x14gogoproto/gogo.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xe9\x04\n\nAttributes\x12>\n\nattributes\x18\x01 \x03(\x0b\x32*.istio.mixer.v1.Attributes.AttributesEntry\x1a\\\n\x0f\x41ttributesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x38\n\x05value\x18\x02 \x01(\x0b\x32).istio.mixer.v1.Attributes.AttributeValue:\x02\x38\x01\x1a\xbb\x02\n\x0e\x41ttributeValue\x12\x16\n\x0cstring_value\x18\x02 \x01(\tH\x00\x12\x15\n\x0bint64_value\x18\x03 \x01(\x03H\x00\x12\x16\n\x0c\x64ouble_value\x18\x04 \x01(\x01H\x00\x12\x14\n\nbool_value\x18\x05 \x01(\x08H\x00\x12\x15\n\x0b\x62ytes_value\x18\x06 \x01(\x0cH\x00\x12\x35\n\x0ftimestamp_value\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.TimestampH\x00\x12\x33\n\x0e\x64uration_value\x18\x08 \x01(\x0b\x32\x19.google.protobuf.DurationH\x00\x12@\n\x10string_map_value\x18\t \x01(\x0b\x32$.istio.mixer.v1.Attributes.StringMapH\x00\x42\x07\n\x05value\x1a\x7f\n\tStringMap\x12\x42\n\x07\x65ntries\x18\x01 \x03(\x0b\x32\x31.istio.mixer.v1.Attributes.StringMap.EntriesEntry\x1a.\n\x0c\x45ntriesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xbb\x08\n\x14\x43ompressedAttributes\x12\r\n\x05words\x18\x01 \x03(\t\x12\x42\n\x07strings\x18\x02 \x03(\x0b\x32\x31.istio.mixer.v1.CompressedAttributes.StringsEntry\x12@\n\x06int64s\x18\x03 \x03(\x0b\x32\x30.istio.mixer.v1.CompressedAttributes.Int64sEntry\x12\x42\n\x07\x64oubles\x18\x04 \x03(\x0b\x32\x31.istio.mixer.v1.CompressedAttributes.DoublesEntry\x12>\n\x05\x62ools\x18\x05 \x03(\x0b\x32/.istio.mixer.v1.CompressedAttributes.BoolsEntry\x12R\n\ntimestamps\x18\x06 \x03(\x0b\x32\x34.istio.mixer.v1.CompressedAttributes.TimestampsEntryB\x08\xc8\xde\x1f\x00\x90\xdf\x1f\x01\x12P\n\tdurations\x18\x07 \x03(\x0b\x32\x33.istio.mixer.v1.CompressedAttributes.DurationsEntryB\x08\xc8\xde\x1f\x00\x98\xdf\x1f\x01\x12>\n\x05\x62ytes\x18\x08 \x03(\x0b\x32/.istio.mixer.v1.CompressedAttributes.BytesEntry\x12O\n\x0bstring_maps\x18\t \x03(\x0b\x32\x34.istio.mixer.v1.CompressedAttributes.StringMapsEntryB\x04\xc8\xde\x1f\x00\x1a.\n\x0cStringsEntry\x12\x0b\n\x03key\x18\x01 \x01(\x11\x12\r\n\x05value\x18\x02 \x01(\x11:\x02\x38\x01\x1a-\n\x0bInt64sEntry\x12\x0b\n\x03key\x18\x01 \x01(\x11\x12\r\n\x05value\x18\x02 \x01(\x03:\x02\x38\x01\x1a.\n\x0c\x44oublesEntry\x12\x0b\n\x03key\x18\x01 \x01(\x11\x12\r\n\x05value\x18\x02 \x01(\x01:\x02\x38\x01\x1a,\n\nBoolsEntry\x12\x0b\n\x03key\x18\x01 \x01(\x11\x12\r\n\x05value\x18\x02 \x01(\x08:\x02\x38\x01\x1aM\n\x0fTimestampsEntry\x12\x0b\n\x03key\x18\x01 \x01(\x11\x12)\n\x05value\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp:\x02\x38\x01\x1aK\n\x0e\x44urationsEntry\x12\x0b\n\x03key\x18\x01 \x01(\x11\x12(\n\x05value\x18\x02 \x01(\x0b\x32\x19.google.protobuf.Duration:\x02\x38\x01\x1a,\n\nBytesEntry\x12\x0b\n\x03key\x18\x01 \x01(\x11\x12\r\n\x05value\x18\x02 \x01(\x0c:\x02\x38\x01\x1aL\n\x0fStringMapsEntry\x12\x0b\n\x03key\x18\x01 \x01(\x11\x12(\n\x05value\x18\x02 \x01(\x0b\x32\x19.istio.mixer.v1.StringMap:\x02\x38\x01\"t\n\tStringMap\x12\x37\n\x07\x65ntries\x18\x01 \x03(\x0b\x32&.istio.mixer.v1.StringMap.EntriesEntry\x1a.\n\x0c\x45ntriesEntry\x12\x0b\n\x03key\x18\x01 \x01(\x11\x12\r\n\x05value\x18\x02 \x01(\x11:\x02\x38\x01\x42*Z\x15istio.io/api/mixer/v1\xf8\x01\x01\xc8\xe1\x1e\x00\xa8\xe2\x1e\x00\xf0\xe1\x1e\x00\xd8\xe2\x1e\x01\x62\x06proto3')
  ,
  dependencies=[gogoproto_dot_gogo__pb2.DESCRIPTOR,google_dot_protobuf_dot_duration__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])




_ATTRIBUTES_ATTRIBUTESENTRY = _descriptor.Descriptor(
  name='AttributesEntry',
  full_name='istio.mixer.v1.Attributes.AttributesEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.mixer.v1.Attributes.AttributesEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.mixer.v1.Attributes.AttributesEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=211,
  serialized_end=303,
)

_ATTRIBUTES_ATTRIBUTEVALUE = _descriptor.Descriptor(
  name='AttributeValue',
  full_name='istio.mixer.v1.Attributes.AttributeValue',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='string_value', full_name='istio.mixer.v1.Attributes.AttributeValue.string_value', index=0,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='int64_value', full_name='istio.mixer.v1.Attributes.AttributeValue.int64_value', index=1,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='double_value', full_name='istio.mixer.v1.Attributes.AttributeValue.double_value', index=2,
      number=4, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='bool_value', full_name='istio.mixer.v1.Attributes.AttributeValue.bool_value', index=3,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='bytes_value', full_name='istio.mixer.v1.Attributes.AttributeValue.bytes_value', index=4,
      number=6, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timestamp_value', full_name='istio.mixer.v1.Attributes.AttributeValue.timestamp_value', index=5,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='duration_value', full_name='istio.mixer.v1.Attributes.AttributeValue.duration_value', index=6,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='string_map_value', full_name='istio.mixer.v1.Attributes.AttributeValue.string_map_value', index=7,
      number=9, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
    _descriptor.OneofDescriptor(
      name='value', full_name='istio.mixer.v1.Attributes.AttributeValue.value',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=306,
  serialized_end=621,
)

_ATTRIBUTES_STRINGMAP_ENTRIESENTRY = _descriptor.Descriptor(
  name='EntriesEntry',
  full_name='istio.mixer.v1.Attributes.StringMap.EntriesEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.mixer.v1.Attributes.StringMap.EntriesEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.mixer.v1.Attributes.StringMap.EntriesEntry.value', index=1,
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
  serialized_start=704,
  serialized_end=750,
)

_ATTRIBUTES_STRINGMAP = _descriptor.Descriptor(
  name='StringMap',
  full_name='istio.mixer.v1.Attributes.StringMap',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='entries', full_name='istio.mixer.v1.Attributes.StringMap.entries', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_ATTRIBUTES_STRINGMAP_ENTRIESENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=623,
  serialized_end=750,
)

_ATTRIBUTES = _descriptor.Descriptor(
  name='Attributes',
  full_name='istio.mixer.v1.Attributes',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='attributes', full_name='istio.mixer.v1.Attributes.attributes', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_ATTRIBUTES_ATTRIBUTESENTRY, _ATTRIBUTES_ATTRIBUTEVALUE, _ATTRIBUTES_STRINGMAP, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=133,
  serialized_end=750,
)


_COMPRESSEDATTRIBUTES_STRINGSENTRY = _descriptor.Descriptor(
  name='StringsEntry',
  full_name='istio.mixer.v1.CompressedAttributes.StringsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.mixer.v1.CompressedAttributes.StringsEntry.key', index=0,
      number=1, type=17, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.mixer.v1.CompressedAttributes.StringsEntry.value', index=1,
      number=2, type=17, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=1369,
  serialized_end=1415,
)

_COMPRESSEDATTRIBUTES_INT64SENTRY = _descriptor.Descriptor(
  name='Int64sEntry',
  full_name='istio.mixer.v1.CompressedAttributes.Int64sEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.mixer.v1.CompressedAttributes.Int64sEntry.key', index=0,
      number=1, type=17, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.mixer.v1.CompressedAttributes.Int64sEntry.value', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=1417,
  serialized_end=1462,
)

_COMPRESSEDATTRIBUTES_DOUBLESENTRY = _descriptor.Descriptor(
  name='DoublesEntry',
  full_name='istio.mixer.v1.CompressedAttributes.DoublesEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.mixer.v1.CompressedAttributes.DoublesEntry.key', index=0,
      number=1, type=17, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.mixer.v1.CompressedAttributes.DoublesEntry.value', index=1,
      number=2, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
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
  serialized_start=1464,
  serialized_end=1510,
)

_COMPRESSEDATTRIBUTES_BOOLSENTRY = _descriptor.Descriptor(
  name='BoolsEntry',
  full_name='istio.mixer.v1.CompressedAttributes.BoolsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.mixer.v1.CompressedAttributes.BoolsEntry.key', index=0,
      number=1, type=17, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.mixer.v1.CompressedAttributes.BoolsEntry.value', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=1512,
  serialized_end=1556,
)

_COMPRESSEDATTRIBUTES_TIMESTAMPSENTRY = _descriptor.Descriptor(
  name='TimestampsEntry',
  full_name='istio.mixer.v1.CompressedAttributes.TimestampsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.mixer.v1.CompressedAttributes.TimestampsEntry.key', index=0,
      number=1, type=17, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.mixer.v1.CompressedAttributes.TimestampsEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=1558,
  serialized_end=1635,
)

_COMPRESSEDATTRIBUTES_DURATIONSENTRY = _descriptor.Descriptor(
  name='DurationsEntry',
  full_name='istio.mixer.v1.CompressedAttributes.DurationsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.mixer.v1.CompressedAttributes.DurationsEntry.key', index=0,
      number=1, type=17, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.mixer.v1.CompressedAttributes.DurationsEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=1637,
  serialized_end=1712,
)

_COMPRESSEDATTRIBUTES_BYTESENTRY = _descriptor.Descriptor(
  name='BytesEntry',
  full_name='istio.mixer.v1.CompressedAttributes.BytesEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.mixer.v1.CompressedAttributes.BytesEntry.key', index=0,
      number=1, type=17, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.mixer.v1.CompressedAttributes.BytesEntry.value', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
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
  serialized_start=1714,
  serialized_end=1758,
)

_COMPRESSEDATTRIBUTES_STRINGMAPSENTRY = _descriptor.Descriptor(
  name='StringMapsEntry',
  full_name='istio.mixer.v1.CompressedAttributes.StringMapsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.mixer.v1.CompressedAttributes.StringMapsEntry.key', index=0,
      number=1, type=17, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.mixer.v1.CompressedAttributes.StringMapsEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=1760,
  serialized_end=1836,
)

_COMPRESSEDATTRIBUTES = _descriptor.Descriptor(
  name='CompressedAttributes',
  full_name='istio.mixer.v1.CompressedAttributes',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='words', full_name='istio.mixer.v1.CompressedAttributes.words', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='strings', full_name='istio.mixer.v1.CompressedAttributes.strings', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='int64s', full_name='istio.mixer.v1.CompressedAttributes.int64s', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='doubles', full_name='istio.mixer.v1.CompressedAttributes.doubles', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='bools', full_name='istio.mixer.v1.CompressedAttributes.bools', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timestamps', full_name='istio.mixer.v1.CompressedAttributes.timestamps', index=5,
      number=6, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\310\336\037\000\220\337\037\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='durations', full_name='istio.mixer.v1.CompressedAttributes.durations', index=6,
      number=7, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\310\336\037\000\230\337\037\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='bytes', full_name='istio.mixer.v1.CompressedAttributes.bytes', index=7,
      number=8, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='string_maps', full_name='istio.mixer.v1.CompressedAttributes.string_maps', index=8,
      number=9, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\310\336\037\000'), file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_COMPRESSEDATTRIBUTES_STRINGSENTRY, _COMPRESSEDATTRIBUTES_INT64SENTRY, _COMPRESSEDATTRIBUTES_DOUBLESENTRY, _COMPRESSEDATTRIBUTES_BOOLSENTRY, _COMPRESSEDATTRIBUTES_TIMESTAMPSENTRY, _COMPRESSEDATTRIBUTES_DURATIONSENTRY, _COMPRESSEDATTRIBUTES_BYTESENTRY, _COMPRESSEDATTRIBUTES_STRINGMAPSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=753,
  serialized_end=1836,
)


_STRINGMAP_ENTRIESENTRY = _descriptor.Descriptor(
  name='EntriesEntry',
  full_name='istio.mixer.v1.StringMap.EntriesEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.mixer.v1.StringMap.EntriesEntry.key', index=0,
      number=1, type=17, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.mixer.v1.StringMap.EntriesEntry.value', index=1,
      number=2, type=17, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=1908,
  serialized_end=1954,
)

_STRINGMAP = _descriptor.Descriptor(
  name='StringMap',
  full_name='istio.mixer.v1.StringMap',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='entries', full_name='istio.mixer.v1.StringMap.entries', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_STRINGMAP_ENTRIESENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1838,
  serialized_end=1954,
)

_ATTRIBUTES_ATTRIBUTESENTRY.fields_by_name['value'].message_type = _ATTRIBUTES_ATTRIBUTEVALUE
_ATTRIBUTES_ATTRIBUTESENTRY.containing_type = _ATTRIBUTES
_ATTRIBUTES_ATTRIBUTEVALUE.fields_by_name['timestamp_value'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_ATTRIBUTES_ATTRIBUTEVALUE.fields_by_name['duration_value'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_ATTRIBUTES_ATTRIBUTEVALUE.fields_by_name['string_map_value'].message_type = _ATTRIBUTES_STRINGMAP
_ATTRIBUTES_ATTRIBUTEVALUE.containing_type = _ATTRIBUTES
_ATTRIBUTES_ATTRIBUTEVALUE.oneofs_by_name['value'].fields.append(
  _ATTRIBUTES_ATTRIBUTEVALUE.fields_by_name['string_value'])
_ATTRIBUTES_ATTRIBUTEVALUE.fields_by_name['string_value'].containing_oneof = _ATTRIBUTES_ATTRIBUTEVALUE.oneofs_by_name['value']
_ATTRIBUTES_ATTRIBUTEVALUE.oneofs_by_name['value'].fields.append(
  _ATTRIBUTES_ATTRIBUTEVALUE.fields_by_name['int64_value'])
_ATTRIBUTES_ATTRIBUTEVALUE.fields_by_name['int64_value'].containing_oneof = _ATTRIBUTES_ATTRIBUTEVALUE.oneofs_by_name['value']
_ATTRIBUTES_ATTRIBUTEVALUE.oneofs_by_name['value'].fields.append(
  _ATTRIBUTES_ATTRIBUTEVALUE.fields_by_name['double_value'])
_ATTRIBUTES_ATTRIBUTEVALUE.fields_by_name['double_value'].containing_oneof = _ATTRIBUTES_ATTRIBUTEVALUE.oneofs_by_name['value']
_ATTRIBUTES_ATTRIBUTEVALUE.oneofs_by_name['value'].fields.append(
  _ATTRIBUTES_ATTRIBUTEVALUE.fields_by_name['bool_value'])
_ATTRIBUTES_ATTRIBUTEVALUE.fields_by_name['bool_value'].containing_oneof = _ATTRIBUTES_ATTRIBUTEVALUE.oneofs_by_name['value']
_ATTRIBUTES_ATTRIBUTEVALUE.oneofs_by_name['value'].fields.append(
  _ATTRIBUTES_ATTRIBUTEVALUE.fields_by_name['bytes_value'])
_ATTRIBUTES_ATTRIBUTEVALUE.fields_by_name['bytes_value'].containing_oneof = _ATTRIBUTES_ATTRIBUTEVALUE.oneofs_by_name['value']
_ATTRIBUTES_ATTRIBUTEVALUE.oneofs_by_name['value'].fields.append(
  _ATTRIBUTES_ATTRIBUTEVALUE.fields_by_name['timestamp_value'])
_ATTRIBUTES_ATTRIBUTEVALUE.fields_by_name['timestamp_value'].containing_oneof = _ATTRIBUTES_ATTRIBUTEVALUE.oneofs_by_name['value']
_ATTRIBUTES_ATTRIBUTEVALUE.oneofs_by_name['value'].fields.append(
  _ATTRIBUTES_ATTRIBUTEVALUE.fields_by_name['duration_value'])
_ATTRIBUTES_ATTRIBUTEVALUE.fields_by_name['duration_value'].containing_oneof = _ATTRIBUTES_ATTRIBUTEVALUE.oneofs_by_name['value']
_ATTRIBUTES_ATTRIBUTEVALUE.oneofs_by_name['value'].fields.append(
  _ATTRIBUTES_ATTRIBUTEVALUE.fields_by_name['string_map_value'])
_ATTRIBUTES_ATTRIBUTEVALUE.fields_by_name['string_map_value'].containing_oneof = _ATTRIBUTES_ATTRIBUTEVALUE.oneofs_by_name['value']
_ATTRIBUTES_STRINGMAP_ENTRIESENTRY.containing_type = _ATTRIBUTES_STRINGMAP
_ATTRIBUTES_STRINGMAP.fields_by_name['entries'].message_type = _ATTRIBUTES_STRINGMAP_ENTRIESENTRY
_ATTRIBUTES_STRINGMAP.containing_type = _ATTRIBUTES
_ATTRIBUTES.fields_by_name['attributes'].message_type = _ATTRIBUTES_ATTRIBUTESENTRY
_COMPRESSEDATTRIBUTES_STRINGSENTRY.containing_type = _COMPRESSEDATTRIBUTES
_COMPRESSEDATTRIBUTES_INT64SENTRY.containing_type = _COMPRESSEDATTRIBUTES
_COMPRESSEDATTRIBUTES_DOUBLESENTRY.containing_type = _COMPRESSEDATTRIBUTES
_COMPRESSEDATTRIBUTES_BOOLSENTRY.containing_type = _COMPRESSEDATTRIBUTES
_COMPRESSEDATTRIBUTES_TIMESTAMPSENTRY.fields_by_name['value'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_COMPRESSEDATTRIBUTES_TIMESTAMPSENTRY.containing_type = _COMPRESSEDATTRIBUTES
_COMPRESSEDATTRIBUTES_DURATIONSENTRY.fields_by_name['value'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_COMPRESSEDATTRIBUTES_DURATIONSENTRY.containing_type = _COMPRESSEDATTRIBUTES
_COMPRESSEDATTRIBUTES_BYTESENTRY.containing_type = _COMPRESSEDATTRIBUTES
_COMPRESSEDATTRIBUTES_STRINGMAPSENTRY.fields_by_name['value'].message_type = _STRINGMAP
_COMPRESSEDATTRIBUTES_STRINGMAPSENTRY.containing_type = _COMPRESSEDATTRIBUTES
_COMPRESSEDATTRIBUTES.fields_by_name['strings'].message_type = _COMPRESSEDATTRIBUTES_STRINGSENTRY
_COMPRESSEDATTRIBUTES.fields_by_name['int64s'].message_type = _COMPRESSEDATTRIBUTES_INT64SENTRY
_COMPRESSEDATTRIBUTES.fields_by_name['doubles'].message_type = _COMPRESSEDATTRIBUTES_DOUBLESENTRY
_COMPRESSEDATTRIBUTES.fields_by_name['bools'].message_type = _COMPRESSEDATTRIBUTES_BOOLSENTRY
_COMPRESSEDATTRIBUTES.fields_by_name['timestamps'].message_type = _COMPRESSEDATTRIBUTES_TIMESTAMPSENTRY
_COMPRESSEDATTRIBUTES.fields_by_name['durations'].message_type = _COMPRESSEDATTRIBUTES_DURATIONSENTRY
_COMPRESSEDATTRIBUTES.fields_by_name['bytes'].message_type = _COMPRESSEDATTRIBUTES_BYTESENTRY
_COMPRESSEDATTRIBUTES.fields_by_name['string_maps'].message_type = _COMPRESSEDATTRIBUTES_STRINGMAPSENTRY
_STRINGMAP_ENTRIESENTRY.containing_type = _STRINGMAP
_STRINGMAP.fields_by_name['entries'].message_type = _STRINGMAP_ENTRIESENTRY
DESCRIPTOR.message_types_by_name['Attributes'] = _ATTRIBUTES
DESCRIPTOR.message_types_by_name['CompressedAttributes'] = _COMPRESSEDATTRIBUTES
DESCRIPTOR.message_types_by_name['StringMap'] = _STRINGMAP
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Attributes = _reflection.GeneratedProtocolMessageType('Attributes', (_message.Message,), dict(

  AttributesEntry = _reflection.GeneratedProtocolMessageType('AttributesEntry', (_message.Message,), dict(
    DESCRIPTOR = _ATTRIBUTES_ATTRIBUTESENTRY,
    __module__ = 'mixer.v1.attributes_pb2'
    # @@protoc_insertion_point(class_scope:istio.mixer.v1.Attributes.AttributesEntry)
    ))
  ,

  AttributeValue = _reflection.GeneratedProtocolMessageType('AttributeValue', (_message.Message,), dict(
    DESCRIPTOR = _ATTRIBUTES_ATTRIBUTEVALUE,
    __module__ = 'mixer.v1.attributes_pb2'
    # @@protoc_insertion_point(class_scope:istio.mixer.v1.Attributes.AttributeValue)
    ))
  ,

  StringMap = _reflection.GeneratedProtocolMessageType('StringMap', (_message.Message,), dict(

    EntriesEntry = _reflection.GeneratedProtocolMessageType('EntriesEntry', (_message.Message,), dict(
      DESCRIPTOR = _ATTRIBUTES_STRINGMAP_ENTRIESENTRY,
      __module__ = 'mixer.v1.attributes_pb2'
      # @@protoc_insertion_point(class_scope:istio.mixer.v1.Attributes.StringMap.EntriesEntry)
      ))
    ,
    DESCRIPTOR = _ATTRIBUTES_STRINGMAP,
    __module__ = 'mixer.v1.attributes_pb2'
    # @@protoc_insertion_point(class_scope:istio.mixer.v1.Attributes.StringMap)
    ))
  ,
  DESCRIPTOR = _ATTRIBUTES,
  __module__ = 'mixer.v1.attributes_pb2'
  # @@protoc_insertion_point(class_scope:istio.mixer.v1.Attributes)
  ))
_sym_db.RegisterMessage(Attributes)
_sym_db.RegisterMessage(Attributes.AttributesEntry)
_sym_db.RegisterMessage(Attributes.AttributeValue)
_sym_db.RegisterMessage(Attributes.StringMap)
_sym_db.RegisterMessage(Attributes.StringMap.EntriesEntry)

CompressedAttributes = _reflection.GeneratedProtocolMessageType('CompressedAttributes', (_message.Message,), dict(

  StringsEntry = _reflection.GeneratedProtocolMessageType('StringsEntry', (_message.Message,), dict(
    DESCRIPTOR = _COMPRESSEDATTRIBUTES_STRINGSENTRY,
    __module__ = 'mixer.v1.attributes_pb2'
    # @@protoc_insertion_point(class_scope:istio.mixer.v1.CompressedAttributes.StringsEntry)
    ))
  ,

  Int64sEntry = _reflection.GeneratedProtocolMessageType('Int64sEntry', (_message.Message,), dict(
    DESCRIPTOR = _COMPRESSEDATTRIBUTES_INT64SENTRY,
    __module__ = 'mixer.v1.attributes_pb2'
    # @@protoc_insertion_point(class_scope:istio.mixer.v1.CompressedAttributes.Int64sEntry)
    ))
  ,

  DoublesEntry = _reflection.GeneratedProtocolMessageType('DoublesEntry', (_message.Message,), dict(
    DESCRIPTOR = _COMPRESSEDATTRIBUTES_DOUBLESENTRY,
    __module__ = 'mixer.v1.attributes_pb2'
    # @@protoc_insertion_point(class_scope:istio.mixer.v1.CompressedAttributes.DoublesEntry)
    ))
  ,

  BoolsEntry = _reflection.GeneratedProtocolMessageType('BoolsEntry', (_message.Message,), dict(
    DESCRIPTOR = _COMPRESSEDATTRIBUTES_BOOLSENTRY,
    __module__ = 'mixer.v1.attributes_pb2'
    # @@protoc_insertion_point(class_scope:istio.mixer.v1.CompressedAttributes.BoolsEntry)
    ))
  ,

  TimestampsEntry = _reflection.GeneratedProtocolMessageType('TimestampsEntry', (_message.Message,), dict(
    DESCRIPTOR = _COMPRESSEDATTRIBUTES_TIMESTAMPSENTRY,
    __module__ = 'mixer.v1.attributes_pb2'
    # @@protoc_insertion_point(class_scope:istio.mixer.v1.CompressedAttributes.TimestampsEntry)
    ))
  ,

  DurationsEntry = _reflection.GeneratedProtocolMessageType('DurationsEntry', (_message.Message,), dict(
    DESCRIPTOR = _COMPRESSEDATTRIBUTES_DURATIONSENTRY,
    __module__ = 'mixer.v1.attributes_pb2'
    # @@protoc_insertion_point(class_scope:istio.mixer.v1.CompressedAttributes.DurationsEntry)
    ))
  ,

  BytesEntry = _reflection.GeneratedProtocolMessageType('BytesEntry', (_message.Message,), dict(
    DESCRIPTOR = _COMPRESSEDATTRIBUTES_BYTESENTRY,
    __module__ = 'mixer.v1.attributes_pb2'
    # @@protoc_insertion_point(class_scope:istio.mixer.v1.CompressedAttributes.BytesEntry)
    ))
  ,

  StringMapsEntry = _reflection.GeneratedProtocolMessageType('StringMapsEntry', (_message.Message,), dict(
    DESCRIPTOR = _COMPRESSEDATTRIBUTES_STRINGMAPSENTRY,
    __module__ = 'mixer.v1.attributes_pb2'
    # @@protoc_insertion_point(class_scope:istio.mixer.v1.CompressedAttributes.StringMapsEntry)
    ))
  ,
  DESCRIPTOR = _COMPRESSEDATTRIBUTES,
  __module__ = 'mixer.v1.attributes_pb2'
  # @@protoc_insertion_point(class_scope:istio.mixer.v1.CompressedAttributes)
  ))
_sym_db.RegisterMessage(CompressedAttributes)
_sym_db.RegisterMessage(CompressedAttributes.StringsEntry)
_sym_db.RegisterMessage(CompressedAttributes.Int64sEntry)
_sym_db.RegisterMessage(CompressedAttributes.DoublesEntry)
_sym_db.RegisterMessage(CompressedAttributes.BoolsEntry)
_sym_db.RegisterMessage(CompressedAttributes.TimestampsEntry)
_sym_db.RegisterMessage(CompressedAttributes.DurationsEntry)
_sym_db.RegisterMessage(CompressedAttributes.BytesEntry)
_sym_db.RegisterMessage(CompressedAttributes.StringMapsEntry)

StringMap = _reflection.GeneratedProtocolMessageType('StringMap', (_message.Message,), dict(

  EntriesEntry = _reflection.GeneratedProtocolMessageType('EntriesEntry', (_message.Message,), dict(
    DESCRIPTOR = _STRINGMAP_ENTRIESENTRY,
    __module__ = 'mixer.v1.attributes_pb2'
    # @@protoc_insertion_point(class_scope:istio.mixer.v1.StringMap.EntriesEntry)
    ))
  ,
  DESCRIPTOR = _STRINGMAP,
  __module__ = 'mixer.v1.attributes_pb2'
  # @@protoc_insertion_point(class_scope:istio.mixer.v1.StringMap)
  ))
_sym_db.RegisterMessage(StringMap)
_sym_db.RegisterMessage(StringMap.EntriesEntry)


DESCRIPTOR._options = None
_ATTRIBUTES_ATTRIBUTESENTRY._options = None
_ATTRIBUTES_STRINGMAP_ENTRIESENTRY._options = None
_COMPRESSEDATTRIBUTES_STRINGSENTRY._options = None
_COMPRESSEDATTRIBUTES_INT64SENTRY._options = None
_COMPRESSEDATTRIBUTES_DOUBLESENTRY._options = None
_COMPRESSEDATTRIBUTES_BOOLSENTRY._options = None
_COMPRESSEDATTRIBUTES_TIMESTAMPSENTRY._options = None
_COMPRESSEDATTRIBUTES_DURATIONSENTRY._options = None
_COMPRESSEDATTRIBUTES_BYTESENTRY._options = None
_COMPRESSEDATTRIBUTES_STRINGMAPSENTRY._options = None
_COMPRESSEDATTRIBUTES.fields_by_name['timestamps']._options = None
_COMPRESSEDATTRIBUTES.fields_by_name['durations']._options = None
_COMPRESSEDATTRIBUTES.fields_by_name['string_maps']._options = None
_STRINGMAP_ENTRIESENTRY._options = None
# @@protoc_insertion_point(module_scope)
