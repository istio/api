# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: mesh/v1alpha1/proxy.proto

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


from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='mesh/v1alpha1/proxy.proto',
  package='istio.mesh.v1alpha1',
  syntax='proto3',
  serialized_pb=_b('\n\x19mesh/v1alpha1/proxy.proto\x12\x13istio.mesh.v1alpha1\x1a\x1egoogle/protobuf/duration.proto\"\xb6\x06\n\x07Tracing\x12\x35\n\x06zipkin\x18\x01 \x01(\x0b\x32#.istio.mesh.v1alpha1.Tracing.ZipkinH\x00\x12;\n\tlightstep\x18\x02 \x01(\x0b\x32&.istio.mesh.v1alpha1.Tracing.LightstepH\x00\x12\x37\n\x07\x64\x61tadog\x18\x03 \x01(\x0b\x32$.istio.mesh.v1alpha1.Tracing.DatadogH\x00\x12=\n\nopencensus\x18\x04 \x01(\x0b\x32\'.istio.mesh.v1alpha1.Tracing.OpenCensusH\x00\x1a\x19\n\x06Zipkin\x12\x0f\n\x07\x61\x64\x64ress\x18\x01 \x01(\t\x1aW\n\tLightstep\x12\x0f\n\x07\x61\x64\x64ress\x18\x01 \x01(\t\x12\x14\n\x0c\x61\x63\x63\x65ss_token\x18\x02 \x01(\t\x12\x0e\n\x06secure\x18\x03 \x01(\x08\x12\x13\n\x0b\x63\x61\x63\x65rt_path\x18\x04 \x01(\t\x1a\x1a\n\x07\x44\x61tadog\x12\x0f\n\x07\x61\x64\x64ress\x18\x01 \x01(\t\x1a\xa4\x03\n\nOpenCensus\x12\x1f\n\x17stdout_exporter_enabled\x18\x01 \x01(\x08\x12$\n\x1cstackdriver_exporter_enabled\x18\x02 \x01(\x08\x12\x1e\n\x16stackdriver_project_id\x18\x03 \x01(\t\x12\x1f\n\x17zipkin_exporter_enabled\x18\x04 \x01(\x08\x12\x12\n\nzipkin_url\x18\x05 \x01(\t\x12T\n\x16incoming_trace_context\x18\x06 \x03(\x0e\x32\x34.istio.mesh.v1alpha1.Tracing.OpenCensus.TraceContext\x12T\n\x16outgoing_trace_context\x18\x07 \x03(\x0e\x32\x34.istio.mesh.v1alpha1.Tracing.OpenCensus.TraceContext\"N\n\x0cTraceContext\x12\x11\n\rTRACE_CONTEXT\x10\x00\x12\x12\n\x0eGRPC_TRACE_BIN\x10\x01\x12\x17\n\x13\x43LOUD_TRACE_CONTEXT\x10\x02\x42\x08\n\x06tracer\"/\n\x03SDS\x12\x0f\n\x07\x65nabled\x18\x01 \x01(\x08\x12\x17\n\x0fk8s_sa_jwt_path\x18\x02 \x01(\t\"\xb1\x07\n\x0bProxyConfig\x12\x13\n\x0b\x63onfig_path\x18\x01 \x01(\t\x12\x13\n\x0b\x62inary_path\x18\x02 \x01(\t\x12\x17\n\x0fservice_cluster\x18\x03 \x01(\t\x12\x31\n\x0e\x64rain_duration\x18\x04 \x01(\x0b\x32\x19.google.protobuf.Duration\x12;\n\x18parent_shutdown_duration\x18\x05 \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x19\n\x11\x64iscovery_address\x18\x06 \x01(\t\x12>\n\x17\x64iscovery_refresh_delay\x18\x07 \x01(\x0b\x32\x19.google.protobuf.DurationB\x02\x18\x01\x12\x1a\n\x0ezipkin_address\x18\x08 \x01(\tB\x02\x18\x01\x12\x32\n\x0f\x63onnect_timeout\x18\t \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x1a\n\x12statsd_udp_address\x18\n \x01(\t\x12%\n\x1d\x65nvoy_metrics_service_address\x18\x14 \x01(\t\x12\x18\n\x10proxy_admin_port\x18\x0b \x01(\x05\x12\x1d\n\x11\x61vailability_zone\x18\x0c \x01(\tB\x02\x18\x01\x12L\n\x19\x63ontrol_plane_auth_policy\x18\r \x01(\x0e\x32).istio.mesh.v1alpha1.AuthenticationPolicy\x12\x1a\n\x12\x63ustom_config_file\x18\x0e \x01(\t\x12\x18\n\x10stat_name_length\x18\x0f \x01(\x05\x12\x13\n\x0b\x63oncurrency\x18\x10 \x01(\x05\x12%\n\x1dproxy_bootstrap_template_path\x18\x11 \x01(\t\x12S\n\x11interception_mode\x18\x12 \x01(\x0e\x32\x38.istio.mesh.v1alpha1.ProxyConfig.InboundInterceptionMode\x12-\n\x07tracing\x18\x13 \x01(\x0b\x32\x1c.istio.mesh.v1alpha1.Tracing\x12%\n\x03sds\x18\x15 \x01(\x0b\x32\x18.istio.mesh.v1alpha1.SDS\x12(\n envoy_access_log_service_address\x18\x16 \x01(\t\"3\n\x17InboundInterceptionMode\x12\x0c\n\x08REDIRECT\x10\x00\x12\n\n\x06TPROXY\x10\x01*>\n\x14\x41uthenticationPolicy\x12\x08\n\x04NONE\x10\x00\x12\x0e\n\nMUTUAL_TLS\x10\x01\x12\x0c\n\x07INHERIT\x10\xe8\x07\x42\x1cZ\x1aistio.io/api/mesh/v1alpha1b\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_duration__pb2.DESCRIPTOR,])

_AUTHENTICATIONPOLICY = _descriptor.EnumDescriptor(
  name='AuthenticationPolicy',
  full_name='istio.mesh.v1alpha1.AuthenticationPolicy',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MUTUAL_TLS', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='INHERIT', index=2, number=1000,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1904,
  serialized_end=1966,
)
_sym_db.RegisterEnumDescriptor(_AUTHENTICATIONPOLICY)

AuthenticationPolicy = enum_type_wrapper.EnumTypeWrapper(_AUTHENTICATIONPOLICY)
NONE = 0
MUTUAL_TLS = 1
INHERIT = 1000


_TRACING_OPENCENSUS_TRACECONTEXT = _descriptor.EnumDescriptor(
  name='TraceContext',
  full_name='istio.mesh.v1alpha1.Tracing.OpenCensus.TraceContext',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='TRACE_CONTEXT', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='GRPC_TRACE_BIN', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CLOUD_TRACE_CONTEXT', index=2, number=2,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=817,
  serialized_end=895,
)
_sym_db.RegisterEnumDescriptor(_TRACING_OPENCENSUS_TRACECONTEXT)

_PROXYCONFIG_INBOUNDINTERCEPTIONMODE = _descriptor.EnumDescriptor(
  name='InboundInterceptionMode',
  full_name='istio.mesh.v1alpha1.ProxyConfig.InboundInterceptionMode',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='REDIRECT', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TPROXY', index=1, number=1,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1851,
  serialized_end=1902,
)
_sym_db.RegisterEnumDescriptor(_PROXYCONFIG_INBOUNDINTERCEPTIONMODE)


_TRACING_ZIPKIN = _descriptor.Descriptor(
  name='Zipkin',
  full_name='istio.mesh.v1alpha1.Tracing.Zipkin',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='address', full_name='istio.mesh.v1alpha1.Tracing.Zipkin.address', index=0,
      number=1, type=9, cpp_type=9, label=1,
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
  serialized_start=330,
  serialized_end=355,
)

_TRACING_LIGHTSTEP = _descriptor.Descriptor(
  name='Lightstep',
  full_name='istio.mesh.v1alpha1.Tracing.Lightstep',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='address', full_name='istio.mesh.v1alpha1.Tracing.Lightstep.address', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='access_token', full_name='istio.mesh.v1alpha1.Tracing.Lightstep.access_token', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='secure', full_name='istio.mesh.v1alpha1.Tracing.Lightstep.secure', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='cacert_path', full_name='istio.mesh.v1alpha1.Tracing.Lightstep.cacert_path', index=3,
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
  serialized_start=357,
  serialized_end=444,
)

_TRACING_DATADOG = _descriptor.Descriptor(
  name='Datadog',
  full_name='istio.mesh.v1alpha1.Tracing.Datadog',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='address', full_name='istio.mesh.v1alpha1.Tracing.Datadog.address', index=0,
      number=1, type=9, cpp_type=9, label=1,
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
  serialized_start=446,
  serialized_end=472,
)

_TRACING_OPENCENSUS = _descriptor.Descriptor(
  name='OpenCensus',
  full_name='istio.mesh.v1alpha1.Tracing.OpenCensus',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='stdout_exporter_enabled', full_name='istio.mesh.v1alpha1.Tracing.OpenCensus.stdout_exporter_enabled', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='stackdriver_exporter_enabled', full_name='istio.mesh.v1alpha1.Tracing.OpenCensus.stackdriver_exporter_enabled', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='stackdriver_project_id', full_name='istio.mesh.v1alpha1.Tracing.OpenCensus.stackdriver_project_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='zipkin_exporter_enabled', full_name='istio.mesh.v1alpha1.Tracing.OpenCensus.zipkin_exporter_enabled', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='zipkin_url', full_name='istio.mesh.v1alpha1.Tracing.OpenCensus.zipkin_url', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='incoming_trace_context', full_name='istio.mesh.v1alpha1.Tracing.OpenCensus.incoming_trace_context', index=5,
      number=6, type=14, cpp_type=8, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='outgoing_trace_context', full_name='istio.mesh.v1alpha1.Tracing.OpenCensus.outgoing_trace_context', index=6,
      number=7, type=14, cpp_type=8, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _TRACING_OPENCENSUS_TRACECONTEXT,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=475,
  serialized_end=895,
)

_TRACING = _descriptor.Descriptor(
  name='Tracing',
  full_name='istio.mesh.v1alpha1.Tracing',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='zipkin', full_name='istio.mesh.v1alpha1.Tracing.zipkin', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='lightstep', full_name='istio.mesh.v1alpha1.Tracing.lightstep', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='datadog', full_name='istio.mesh.v1alpha1.Tracing.datadog', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='opencensus', full_name='istio.mesh.v1alpha1.Tracing.opencensus', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_TRACING_ZIPKIN, _TRACING_LIGHTSTEP, _TRACING_DATADOG, _TRACING_OPENCENSUS, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='tracer', full_name='istio.mesh.v1alpha1.Tracing.tracer',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=83,
  serialized_end=905,
)


_SDS = _descriptor.Descriptor(
  name='SDS',
  full_name='istio.mesh.v1alpha1.SDS',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='enabled', full_name='istio.mesh.v1alpha1.SDS.enabled', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='k8s_sa_jwt_path', full_name='istio.mesh.v1alpha1.SDS.k8s_sa_jwt_path', index=1,
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
  serialized_start=907,
  serialized_end=954,
)


_PROXYCONFIG = _descriptor.Descriptor(
  name='ProxyConfig',
  full_name='istio.mesh.v1alpha1.ProxyConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='config_path', full_name='istio.mesh.v1alpha1.ProxyConfig.config_path', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='binary_path', full_name='istio.mesh.v1alpha1.ProxyConfig.binary_path', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='service_cluster', full_name='istio.mesh.v1alpha1.ProxyConfig.service_cluster', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='drain_duration', full_name='istio.mesh.v1alpha1.ProxyConfig.drain_duration', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='parent_shutdown_duration', full_name='istio.mesh.v1alpha1.ProxyConfig.parent_shutdown_duration', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='discovery_address', full_name='istio.mesh.v1alpha1.ProxyConfig.discovery_address', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='discovery_refresh_delay', full_name='istio.mesh.v1alpha1.ProxyConfig.discovery_refresh_delay', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\030\001')), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='zipkin_address', full_name='istio.mesh.v1alpha1.ProxyConfig.zipkin_address', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\030\001')), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='connect_timeout', full_name='istio.mesh.v1alpha1.ProxyConfig.connect_timeout', index=8,
      number=9, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='statsd_udp_address', full_name='istio.mesh.v1alpha1.ProxyConfig.statsd_udp_address', index=9,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='envoy_metrics_service_address', full_name='istio.mesh.v1alpha1.ProxyConfig.envoy_metrics_service_address', index=10,
      number=20, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='proxy_admin_port', full_name='istio.mesh.v1alpha1.ProxyConfig.proxy_admin_port', index=11,
      number=11, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='availability_zone', full_name='istio.mesh.v1alpha1.ProxyConfig.availability_zone', index=12,
      number=12, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\030\001')), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='control_plane_auth_policy', full_name='istio.mesh.v1alpha1.ProxyConfig.control_plane_auth_policy', index=13,
      number=13, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='custom_config_file', full_name='istio.mesh.v1alpha1.ProxyConfig.custom_config_file', index=14,
      number=14, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='stat_name_length', full_name='istio.mesh.v1alpha1.ProxyConfig.stat_name_length', index=15,
      number=15, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='concurrency', full_name='istio.mesh.v1alpha1.ProxyConfig.concurrency', index=16,
      number=16, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='proxy_bootstrap_template_path', full_name='istio.mesh.v1alpha1.ProxyConfig.proxy_bootstrap_template_path', index=17,
      number=17, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='interception_mode', full_name='istio.mesh.v1alpha1.ProxyConfig.interception_mode', index=18,
      number=18, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tracing', full_name='istio.mesh.v1alpha1.ProxyConfig.tracing', index=19,
      number=19, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sds', full_name='istio.mesh.v1alpha1.ProxyConfig.sds', index=20,
      number=21, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='envoy_access_log_service_address', full_name='istio.mesh.v1alpha1.ProxyConfig.envoy_access_log_service_address', index=21,
      number=22, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _PROXYCONFIG_INBOUNDINTERCEPTIONMODE,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=957,
  serialized_end=1902,
)

_TRACING_ZIPKIN.containing_type = _TRACING
_TRACING_LIGHTSTEP.containing_type = _TRACING
_TRACING_DATADOG.containing_type = _TRACING
_TRACING_OPENCENSUS.fields_by_name['incoming_trace_context'].enum_type = _TRACING_OPENCENSUS_TRACECONTEXT
_TRACING_OPENCENSUS.fields_by_name['outgoing_trace_context'].enum_type = _TRACING_OPENCENSUS_TRACECONTEXT
_TRACING_OPENCENSUS.containing_type = _TRACING
_TRACING_OPENCENSUS_TRACECONTEXT.containing_type = _TRACING_OPENCENSUS
_TRACING.fields_by_name['zipkin'].message_type = _TRACING_ZIPKIN
_TRACING.fields_by_name['lightstep'].message_type = _TRACING_LIGHTSTEP
_TRACING.fields_by_name['datadog'].message_type = _TRACING_DATADOG
_TRACING.fields_by_name['opencensus'].message_type = _TRACING_OPENCENSUS
_TRACING.oneofs_by_name['tracer'].fields.append(
  _TRACING.fields_by_name['zipkin'])
_TRACING.fields_by_name['zipkin'].containing_oneof = _TRACING.oneofs_by_name['tracer']
_TRACING.oneofs_by_name['tracer'].fields.append(
  _TRACING.fields_by_name['lightstep'])
_TRACING.fields_by_name['lightstep'].containing_oneof = _TRACING.oneofs_by_name['tracer']
_TRACING.oneofs_by_name['tracer'].fields.append(
  _TRACING.fields_by_name['datadog'])
_TRACING.fields_by_name['datadog'].containing_oneof = _TRACING.oneofs_by_name['tracer']
_TRACING.oneofs_by_name['tracer'].fields.append(
  _TRACING.fields_by_name['opencensus'])
_TRACING.fields_by_name['opencensus'].containing_oneof = _TRACING.oneofs_by_name['tracer']
_PROXYCONFIG.fields_by_name['drain_duration'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_PROXYCONFIG.fields_by_name['parent_shutdown_duration'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_PROXYCONFIG.fields_by_name['discovery_refresh_delay'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_PROXYCONFIG.fields_by_name['connect_timeout'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_PROXYCONFIG.fields_by_name['control_plane_auth_policy'].enum_type = _AUTHENTICATIONPOLICY
_PROXYCONFIG.fields_by_name['interception_mode'].enum_type = _PROXYCONFIG_INBOUNDINTERCEPTIONMODE
_PROXYCONFIG.fields_by_name['tracing'].message_type = _TRACING
_PROXYCONFIG.fields_by_name['sds'].message_type = _SDS
_PROXYCONFIG_INBOUNDINTERCEPTIONMODE.containing_type = _PROXYCONFIG
DESCRIPTOR.message_types_by_name['Tracing'] = _TRACING
DESCRIPTOR.message_types_by_name['SDS'] = _SDS
DESCRIPTOR.message_types_by_name['ProxyConfig'] = _PROXYCONFIG
DESCRIPTOR.enum_types_by_name['AuthenticationPolicy'] = _AUTHENTICATIONPOLICY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Tracing = _reflection.GeneratedProtocolMessageType('Tracing', (_message.Message,), dict(

  Zipkin = _reflection.GeneratedProtocolMessageType('Zipkin', (_message.Message,), dict(
    DESCRIPTOR = _TRACING_ZIPKIN,
    __module__ = 'mesh.v1alpha1.proxy_pb2'
    # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.Tracing.Zipkin)
    ))
  ,

  Lightstep = _reflection.GeneratedProtocolMessageType('Lightstep', (_message.Message,), dict(
    DESCRIPTOR = _TRACING_LIGHTSTEP,
    __module__ = 'mesh.v1alpha1.proxy_pb2'
    # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.Tracing.Lightstep)
    ))
  ,

  Datadog = _reflection.GeneratedProtocolMessageType('Datadog', (_message.Message,), dict(
    DESCRIPTOR = _TRACING_DATADOG,
    __module__ = 'mesh.v1alpha1.proxy_pb2'
    # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.Tracing.Datadog)
    ))
  ,

  OpenCensus = _reflection.GeneratedProtocolMessageType('OpenCensus', (_message.Message,), dict(
    DESCRIPTOR = _TRACING_OPENCENSUS,
    __module__ = 'mesh.v1alpha1.proxy_pb2'
    # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.Tracing.OpenCensus)
    ))
  ,
  DESCRIPTOR = _TRACING,
  __module__ = 'mesh.v1alpha1.proxy_pb2'
  # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.Tracing)
  ))
_sym_db.RegisterMessage(Tracing)
_sym_db.RegisterMessage(Tracing.Zipkin)
_sym_db.RegisterMessage(Tracing.Lightstep)
_sym_db.RegisterMessage(Tracing.Datadog)
_sym_db.RegisterMessage(Tracing.OpenCensus)

SDS = _reflection.GeneratedProtocolMessageType('SDS', (_message.Message,), dict(
  DESCRIPTOR = _SDS,
  __module__ = 'mesh.v1alpha1.proxy_pb2'
  # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.SDS)
  ))
_sym_db.RegisterMessage(SDS)

ProxyConfig = _reflection.GeneratedProtocolMessageType('ProxyConfig', (_message.Message,), dict(
  DESCRIPTOR = _PROXYCONFIG,
  __module__ = 'mesh.v1alpha1.proxy_pb2'
  # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.ProxyConfig)
  ))
_sym_db.RegisterMessage(ProxyConfig)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z\032istio.io/api/mesh/v1alpha1'))
_PROXYCONFIG.fields_by_name['discovery_refresh_delay'].has_options = True
_PROXYCONFIG.fields_by_name['discovery_refresh_delay']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\030\001'))
_PROXYCONFIG.fields_by_name['zipkin_address'].has_options = True
_PROXYCONFIG.fields_by_name['zipkin_address']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\030\001'))
_PROXYCONFIG.fields_by_name['availability_zone'].has_options = True
_PROXYCONFIG.fields_by_name['availability_zone']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\030\001'))
# @@protoc_insertion_point(module_scope)
