# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: policy/v1beta1/http_response.proto

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
  name='policy/v1beta1/http_response.proto',
  package='istio.policy.v1beta1',
  syntax='proto3',
  serialized_pb=_b('\n\"policy/v1beta1/http_response.proto\x12\x14istio.policy.v1beta1\"V\n\x12\x44irectHttpResponse\x12\x32\n\x04\x63ode\x18\x01 \x01(\x0e\x32$.istio.policy.v1beta1.HttpStatusCode\x12\x0c\n\x04\x62ody\x18\x02 \x01(\t*\xb9\t\n\x0eHttpStatusCode\x12\t\n\x05\x45mpty\x10\x00\x12\x0c\n\x08\x43ontinue\x10\x64\x12\x07\n\x02OK\x10\xc8\x01\x12\x0c\n\x07\x43reated\x10\xc9\x01\x12\r\n\x08\x41\x63\x63\x65pted\x10\xca\x01\x12 \n\x1bNonAuthoritativeInformation\x10\xcb\x01\x12\x0e\n\tNoContent\x10\xcc\x01\x12\x11\n\x0cResetContent\x10\xcd\x01\x12\x13\n\x0ePartialContent\x10\xce\x01\x12\x10\n\x0bMultiStatus\x10\xcf\x01\x12\x14\n\x0f\x41lreadyReported\x10\xd0\x01\x12\x0b\n\x06IMUsed\x10\xe2\x01\x12\x14\n\x0fMultipleChoices\x10\xac\x02\x12\x15\n\x10MovedPermanently\x10\xad\x02\x12\n\n\x05\x46ound\x10\xae\x02\x12\r\n\x08SeeOther\x10\xaf\x02\x12\x10\n\x0bNotModified\x10\xb0\x02\x12\r\n\x08UseProxy\x10\xb1\x02\x12\x16\n\x11TemporaryRedirect\x10\xb3\x02\x12\x16\n\x11PermanentRedirect\x10\xb4\x02\x12\x0f\n\nBadRequest\x10\x90\x03\x12\x11\n\x0cUnauthorized\x10\x91\x03\x12\x14\n\x0fPaymentRequired\x10\x92\x03\x12\x0e\n\tForbidden\x10\x93\x03\x12\r\n\x08NotFound\x10\x94\x03\x12\x15\n\x10MethodNotAllowed\x10\x95\x03\x12\x12\n\rNotAcceptable\x10\x96\x03\x12 \n\x1bProxyAuthenticationRequired\x10\x97\x03\x12\x13\n\x0eRequestTimeout\x10\x98\x03\x12\r\n\x08\x43onflict\x10\x99\x03\x12\t\n\x04Gone\x10\x9a\x03\x12\x13\n\x0eLengthRequired\x10\x9b\x03\x12\x17\n\x12PreconditionFailed\x10\x9c\x03\x12\x14\n\x0fPayloadTooLarge\x10\x9d\x03\x12\x0f\n\nURITooLong\x10\x9e\x03\x12\x19\n\x14UnsupportedMediaType\x10\x9f\x03\x12\x18\n\x13RangeNotSatisfiable\x10\xa0\x03\x12\x16\n\x11\x45xpectationFailed\x10\xa1\x03\x12\x17\n\x12MisdirectedRequest\x10\xa5\x03\x12\x18\n\x13UnprocessableEntity\x10\xa6\x03\x12\x0b\n\x06Locked\x10\xa7\x03\x12\x15\n\x10\x46\x61iledDependency\x10\xa8\x03\x12\x14\n\x0fUpgradeRequired\x10\xaa\x03\x12\x19\n\x14PreconditionRequired\x10\xac\x03\x12\x14\n\x0fTooManyRequests\x10\xad\x03\x12 \n\x1bRequestHeaderFieldsTooLarge\x10\xaf\x03\x12\x18\n\x13InternalServerError\x10\xf4\x03\x12\x13\n\x0eNotImplemented\x10\xf5\x03\x12\x0f\n\nBadGateway\x10\xf6\x03\x12\x17\n\x12ServiceUnavailable\x10\xf7\x03\x12\x13\n\x0eGatewayTimeout\x10\xf8\x03\x12\x1c\n\x17HTTPVersionNotSupported\x10\xf9\x03\x12\x1a\n\x15VariantAlsoNegotiates\x10\xfa\x03\x12\x18\n\x13InsufficientStorage\x10\xfb\x03\x12\x11\n\x0cLoopDetected\x10\xfc\x03\x12\x10\n\x0bNotExtended\x10\xfe\x03\x12\"\n\x1dNetworkAuthenticationRequired\x10\xff\x03\x42\x1dZ\x1bistio.io/api/policy/v1beta1b\x06proto3')
)

_HTTPSTATUSCODE = _descriptor.EnumDescriptor(
  name='HttpStatusCode',
  full_name='istio.policy.v1beta1.HttpStatusCode',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Empty', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Continue', index=1, number=100,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='OK', index=2, number=200,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Created', index=3, number=201,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Accepted', index=4, number=202,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NonAuthoritativeInformation', index=5, number=203,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NoContent', index=6, number=204,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ResetContent', index=7, number=205,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PartialContent', index=8, number=206,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MultiStatus', index=9, number=207,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AlreadyReported', index=10, number=208,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='IMUsed', index=11, number=226,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MultipleChoices', index=12, number=300,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MovedPermanently', index=13, number=301,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Found', index=14, number=302,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SeeOther', index=15, number=303,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NotModified', index=16, number=304,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='UseProxy', index=17, number=305,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TemporaryRedirect', index=18, number=307,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PermanentRedirect', index=19, number=308,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='BadRequest', index=20, number=400,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Unauthorized', index=21, number=401,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PaymentRequired', index=22, number=402,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Forbidden', index=23, number=403,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NotFound', index=24, number=404,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MethodNotAllowed', index=25, number=405,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NotAcceptable', index=26, number=406,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ProxyAuthenticationRequired', index=27, number=407,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='RequestTimeout', index=28, number=408,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Conflict', index=29, number=409,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Gone', index=30, number=410,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='LengthRequired', index=31, number=411,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PreconditionFailed', index=32, number=412,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PayloadTooLarge', index=33, number=413,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='URITooLong', index=34, number=414,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='UnsupportedMediaType', index=35, number=415,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='RangeNotSatisfiable', index=36, number=416,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ExpectationFailed', index=37, number=417,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MisdirectedRequest', index=38, number=421,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='UnprocessableEntity', index=39, number=422,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Locked', index=40, number=423,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FailedDependency', index=41, number=424,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='UpgradeRequired', index=42, number=426,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PreconditionRequired', index=43, number=428,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TooManyRequests', index=44, number=429,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='RequestHeaderFieldsTooLarge', index=45, number=431,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='InternalServerError', index=46, number=500,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NotImplemented', index=47, number=501,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='BadGateway', index=48, number=502,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ServiceUnavailable', index=49, number=503,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='GatewayTimeout', index=50, number=504,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HTTPVersionNotSupported', index=51, number=505,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='VariantAlsoNegotiates', index=52, number=506,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='InsufficientStorage', index=53, number=507,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='LoopDetected', index=54, number=508,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NotExtended', index=55, number=510,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NetworkAuthenticationRequired', index=56, number=511,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=149,
  serialized_end=1358,
)
_sym_db.RegisterEnumDescriptor(_HTTPSTATUSCODE)

HttpStatusCode = enum_type_wrapper.EnumTypeWrapper(_HTTPSTATUSCODE)
Empty = 0
Continue = 100
OK = 200
Created = 201
Accepted = 202
NonAuthoritativeInformation = 203
NoContent = 204
ResetContent = 205
PartialContent = 206
MultiStatus = 207
AlreadyReported = 208
IMUsed = 226
MultipleChoices = 300
MovedPermanently = 301
Found = 302
SeeOther = 303
NotModified = 304
UseProxy = 305
TemporaryRedirect = 307
PermanentRedirect = 308
BadRequest = 400
Unauthorized = 401
PaymentRequired = 402
Forbidden = 403
NotFound = 404
MethodNotAllowed = 405
NotAcceptable = 406
ProxyAuthenticationRequired = 407
RequestTimeout = 408
Conflict = 409
Gone = 410
LengthRequired = 411
PreconditionFailed = 412
PayloadTooLarge = 413
URITooLong = 414
UnsupportedMediaType = 415
RangeNotSatisfiable = 416
ExpectationFailed = 417
MisdirectedRequest = 421
UnprocessableEntity = 422
Locked = 423
FailedDependency = 424
UpgradeRequired = 426
PreconditionRequired = 428
TooManyRequests = 429
RequestHeaderFieldsTooLarge = 431
InternalServerError = 500
NotImplemented = 501
BadGateway = 502
ServiceUnavailable = 503
GatewayTimeout = 504
HTTPVersionNotSupported = 505
VariantAlsoNegotiates = 506
InsufficientStorage = 507
LoopDetected = 508
NotExtended = 510
NetworkAuthenticationRequired = 511



_DIRECTHTTPRESPONSE = _descriptor.Descriptor(
  name='DirectHttpResponse',
  full_name='istio.policy.v1beta1.DirectHttpResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='istio.policy.v1beta1.DirectHttpResponse.code', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='body', full_name='istio.policy.v1beta1.DirectHttpResponse.body', index=1,
      number=2, type=9, cpp_type=9, label=1,
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
  serialized_start=60,
  serialized_end=146,
)

_DIRECTHTTPRESPONSE.fields_by_name['code'].enum_type = _HTTPSTATUSCODE
DESCRIPTOR.message_types_by_name['DirectHttpResponse'] = _DIRECTHTTPRESPONSE
DESCRIPTOR.enum_types_by_name['HttpStatusCode'] = _HTTPSTATUSCODE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

DirectHttpResponse = _reflection.GeneratedProtocolMessageType('DirectHttpResponse', (_message.Message,), dict(
  DESCRIPTOR = _DIRECTHTTPRESPONSE,
  __module__ = 'policy.v1beta1.http_response_pb2'
  # @@protoc_insertion_point(class_scope:istio.policy.v1beta1.DirectHttpResponse)
  ))
_sym_db.RegisterMessage(DirectHttpResponse)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z\033istio.io/api/policy/v1beta1'))
# @@protoc_insertion_point(module_scope)
