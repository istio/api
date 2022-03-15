# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: networking/v1beta1/destination_rule.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import field_behavior_pb2 as google_dot_api_dot_field__behavior__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2
from networking.v1beta1 import virtual_service_pb2 as networking_dot_v1beta1_dot_virtual__service__pb2
from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n)networking/v1beta1/destination_rule.proto\x12\x18istio.networking.v1beta1\x1a\x1fgoogle/api/field_behavior.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a(networking/v1beta1/virtual_service.proto\x1a\x14gogoproto/gogo.proto\"\xd4\x01\n\x0f\x44\x65stinationRule\x12\x18\n\x04host\x18\x01 \x01(\tB\x04\xe2\x41\x01\x02R\x04host\x12N\n\x0etraffic_policy\x18\x02 \x01(\x0b\x32\'.istio.networking.v1beta1.TrafficPolicyR\rtrafficPolicy\x12:\n\x07subsets\x18\x03 \x03(\x0b\x32 .istio.networking.v1beta1.SubsetR\x07subsets\x12\x1b\n\texport_to\x18\x04 \x03(\tR\x08\x65xportTo\"\xdc\x06\n\rTrafficPolicy\x12S\n\rload_balancer\x18\x01 \x01(\x0b\x32..istio.networking.v1beta1.LoadBalancerSettingsR\x0cloadBalancer\x12Y\n\x0f\x63onnection_pool\x18\x02 \x01(\x0b\x32\x30.istio.networking.v1beta1.ConnectionPoolSettingsR\x0e\x63onnectionPool\x12W\n\x11outlier_detection\x18\x03 \x01(\x0b\x32*.istio.networking.v1beta1.OutlierDetectionR\x10outlierDetection\x12=\n\x03tls\x18\x04 \x01(\x0b\x32+.istio.networking.v1beta1.ClientTLSSettingsR\x03tls\x12i\n\x13port_level_settings\x18\x05 \x03(\x0b\x32\x39.istio.networking.v1beta1.TrafficPolicy.PortTrafficPolicyR\x11portLevelSettings\x1a\x97\x03\n\x11PortTrafficPolicy\x12:\n\x04port\x18\x01 \x01(\x0b\x32&.istio.networking.v1beta1.PortSelectorR\x04port\x12S\n\rload_balancer\x18\x02 \x01(\x0b\x32..istio.networking.v1beta1.LoadBalancerSettingsR\x0cloadBalancer\x12Y\n\x0f\x63onnection_pool\x18\x03 \x01(\x0b\x32\x30.istio.networking.v1beta1.ConnectionPoolSettingsR\x0e\x63onnectionPool\x12W\n\x11outlier_detection\x18\x04 \x01(\x0b\x32*.istio.networking.v1beta1.OutlierDetectionR\x10outlierDetection\x12=\n\x03tls\x18\x05 \x01(\x0b\x32+.istio.networking.v1beta1.ClientTLSSettingsR\x03tls\"\xf3\x01\n\x06Subset\x12\x18\n\x04name\x18\x01 \x01(\tB\x04\xe2\x41\x01\x02R\x04name\x12\x44\n\x06labels\x18\x02 \x03(\x0b\x32,.istio.networking.v1beta1.Subset.LabelsEntryR\x06labels\x12N\n\x0etraffic_policy\x18\x03 \x01(\x0b\x32\'.istio.networking.v1beta1.TrafficPolicyR\rtrafficPolicy\x1a\x39\n\x0bLabelsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\"\xc2\x07\n\x14LoadBalancerSettings\x12Q\n\x06simple\x18\x01 \x01(\x0e\x32\x37.istio.networking.v1beta1.LoadBalancerSettings.SimpleLBH\x00R\x06simple\x12j\n\x0f\x63onsistent_hash\x18\x02 \x01(\x0b\x32?.istio.networking.v1beta1.LoadBalancerSettings.ConsistentHashLBH\x00R\x0e\x63onsistentHash\x12\x65\n\x13locality_lb_setting\x18\x03 \x01(\x0b\x32\x35.istio.networking.v1beta1.LocalityLoadBalancerSettingR\x11localityLbSetting\x12K\n\x14warmup_duration_secs\x18\x04 \x01(\x0b\x32\x19.google.protobuf.DurationR\x12warmupDurationSecs\x1a\xb7\x03\n\x10\x43onsistentHashLB\x12*\n\x10http_header_name\x18\x01 \x01(\tH\x00R\x0ehttpHeaderName\x12m\n\x0bhttp_cookie\x18\x02 \x01(\x0b\x32J.istio.networking.v1beta1.LoadBalancerSettings.ConsistentHashLB.HTTPCookieH\x00R\nhttpCookie\x12$\n\ruse_source_ip\x18\x03 \x01(\x08H\x00R\x0buseSourceIp\x12;\n\x19http_query_parameter_name\x18\x05 \x01(\tH\x00R\x16httpQueryParameterName\x12*\n\x11minimum_ring_size\x18\x04 \x01(\x04R\x0fminimumRingSize\x1am\n\nHTTPCookie\x12\x18\n\x04name\x18\x01 \x01(\tB\x04\xe2\x41\x01\x02R\x04name\x12\x12\n\x04path\x18\x02 \x01(\tR\x04path\x12\x31\n\x03ttl\x18\x03 \x01(\x0b\x32\x19.google.protobuf.DurationB\x04\xe2\x41\x01\x02R\x03ttlB\n\n\x08hash_key\"p\n\x08SimpleLB\x12\x0f\n\x0bUNSPECIFIED\x10\x00\x12\x12\n\nLEAST_CONN\x10\x01\x1a\x02\x08\x01\x12\n\n\x06RANDOM\x10\x02\x12\x0f\n\x0bPASSTHROUGH\x10\x03\x12\x0f\n\x0bROUND_ROBIN\x10\x04\x12\x11\n\rLEAST_REQUEST\x10\x05\x42\x0b\n\tlb_policy\"\xbd\x08\n\x16\x43onnectionPoolSettings\x12N\n\x03tcp\x18\x01 \x01(\x0b\x32<.istio.networking.v1beta1.ConnectionPoolSettings.TCPSettingsR\x03tcp\x12Q\n\x04http\x18\x02 \x01(\x0b\x32=.istio.networking.v1beta1.ConnectionPoolSettings.HTTPSettingsR\x04http\x1a\xf9\x02\n\x0bTCPSettings\x12\'\n\x0fmax_connections\x18\x01 \x01(\x05R\x0emaxConnections\x12\x42\n\x0f\x63onnect_timeout\x18\x02 \x01(\x0b\x32\x19.google.protobuf.DurationR\x0e\x63onnectTimeout\x12n\n\rtcp_keepalive\x18\x03 \x01(\x0b\x32I.istio.networking.v1beta1.ConnectionPoolSettings.TCPSettings.TcpKeepaliveR\x0ctcpKeepalive\x1a\x8c\x01\n\x0cTcpKeepalive\x12\x16\n\x06probes\x18\x01 \x01(\rR\x06probes\x12-\n\x04time\x18\x02 \x01(\x0b\x32\x19.google.protobuf.DurationR\x04time\x12\x35\n\x08interval\x18\x03 \x01(\x0b\x32\x19.google.protobuf.DurationR\x08interval\x1a\x83\x04\n\x0cHTTPSettings\x12;\n\x1ahttp1_max_pending_requests\x18\x01 \x01(\x05R\x17http1MaxPendingRequests\x12,\n\x12http2_max_requests\x18\x02 \x01(\x05R\x10http2MaxRequests\x12=\n\x1bmax_requests_per_connection\x18\x03 \x01(\x05R\x18maxRequestsPerConnection\x12\x1f\n\x0bmax_retries\x18\x04 \x01(\x05R\nmaxRetries\x12<\n\x0cidle_timeout\x18\x05 \x01(\x0b\x32\x19.google.protobuf.DurationR\x0bidleTimeout\x12y\n\x11h2_upgrade_policy\x18\x06 \x01(\x0e\x32M.istio.networking.v1beta1.ConnectionPoolSettings.HTTPSettings.H2UpgradePolicyR\x0fh2UpgradePolicy\x12.\n\x13use_client_protocol\x18\x07 \x01(\x08R\x11useClientProtocol\"?\n\x0fH2UpgradePolicy\x12\x0b\n\x07\x44\x45\x46\x41ULT\x10\x00\x12\x12\n\x0e\x44O_NOT_UPGRADE\x10\x01\x12\x0b\n\x07UPGRADE\x10\x02\"\x8a\x05\n\x10OutlierDetection\x12\x31\n\x12\x63onsecutive_errors\x18\x01 \x01(\x05\x42\x02\x18\x01R\x11\x63onsecutiveErrors\x12J\n\"split_external_local_origin_errors\x18\x08 \x01(\x08R\x1esplitExternalLocalOriginErrors\x12g\n!consecutive_local_origin_failures\x18\t \x01(\x0b\x32\x1c.google.protobuf.UInt32ValueR\x1e\x63onsecutiveLocalOriginFailures\x12Z\n\x1a\x63onsecutive_gateway_errors\x18\x06 \x01(\x0b\x32\x1c.google.protobuf.UInt32ValueR\x18\x63onsecutiveGatewayErrors\x12R\n\x16\x63onsecutive_5xx_errors\x18\x07 \x01(\x0b\x32\x1c.google.protobuf.UInt32ValueR\x14\x63onsecutive5xxErrors\x12\x35\n\x08interval\x18\x02 \x01(\x0b\x32\x19.google.protobuf.DurationR\x08interval\x12G\n\x12\x62\x61se_ejection_time\x18\x03 \x01(\x0b\x32\x19.google.protobuf.DurationR\x10\x62\x61seEjectionTime\x12\x30\n\x14max_ejection_percent\x18\x04 \x01(\x05R\x12maxEjectionPercent\x12,\n\x12min_health_percent\x18\x05 \x01(\x05R\x10minHealthPercent\"\xd2\x03\n\x11\x43lientTLSSettings\x12M\n\x04mode\x18\x01 \x01(\x0e\x32\x33.istio.networking.v1beta1.ClientTLSSettings.TLSmodeB\x04\xe2\x41\x01\x02R\x04mode\x12-\n\x12\x63lient_certificate\x18\x02 \x01(\tR\x11\x63lientCertificate\x12\x1f\n\x0bprivate_key\x18\x03 \x01(\tR\nprivateKey\x12\'\n\x0f\x63\x61_certificates\x18\x04 \x01(\tR\x0e\x63\x61\x43\x65rtificates\x12\'\n\x0f\x63redential_name\x18\x07 \x01(\tR\x0e\x63redentialName\x12*\n\x11subject_alt_names\x18\x05 \x03(\tR\x0fsubjectAltNames\x12\x10\n\x03sni\x18\x06 \x01(\tR\x03sni\x12L\n\x14insecure_skip_verify\x18\x08 \x01(\x0b\x32\x1a.google.protobuf.BoolValueR\x12insecureSkipVerify\"@\n\x07TLSmode\x12\x0b\n\x07\x44ISABLE\x10\x00\x12\n\n\x06SIMPLE\x10\x01\x12\n\n\x06MUTUAL\x10\x02\x12\x10\n\x0cISTIO_MUTUAL\x10\x03\"\xa2\x04\n\x1bLocalityLoadBalancerSetting\x12`\n\ndistribute\x18\x01 \x03(\x0b\x32@.istio.networking.v1beta1.LocalityLoadBalancerSetting.DistributeR\ndistribute\x12Z\n\x08\x66\x61ilover\x18\x02 \x03(\x0b\x32>.istio.networking.v1beta1.LocalityLoadBalancerSetting.FailoverR\x08\x66\x61ilover\x12+\n\x11\x66\x61ilover_priority\x18\x04 \x03(\tR\x10\x66\x61iloverPriority\x12\x34\n\x07\x65nabled\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.BoolValueR\x07\x65nabled\x1a\xb1\x01\n\nDistribute\x12\x12\n\x04\x66rom\x18\x01 \x01(\tR\x04\x66rom\x12X\n\x02to\x18\x02 \x03(\x0b\x32H.istio.networking.v1beta1.LocalityLoadBalancerSetting.Distribute.ToEntryR\x02to\x1a\x35\n\x07ToEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\rR\x05value:\x02\x38\x01\x1a.\n\x08\x46\x61ilover\x12\x12\n\x04\x66rom\x18\x01 \x01(\tR\x04\x66rom\x12\x0e\n\x02to\x18\x02 \x01(\tR\x02toB!Z\x1fistio.io/api/networking/v1beta1b\x06proto3')



_DESTINATIONRULE = DESCRIPTOR.message_types_by_name['DestinationRule']
_TRAFFICPOLICY = DESCRIPTOR.message_types_by_name['TrafficPolicy']
_TRAFFICPOLICY_PORTTRAFFICPOLICY = _TRAFFICPOLICY.nested_types_by_name['PortTrafficPolicy']
_SUBSET = DESCRIPTOR.message_types_by_name['Subset']
_SUBSET_LABELSENTRY = _SUBSET.nested_types_by_name['LabelsEntry']
_LOADBALANCERSETTINGS = DESCRIPTOR.message_types_by_name['LoadBalancerSettings']
_LOADBALANCERSETTINGS_CONSISTENTHASHLB = _LOADBALANCERSETTINGS.nested_types_by_name['ConsistentHashLB']
_LOADBALANCERSETTINGS_CONSISTENTHASHLB_HTTPCOOKIE = _LOADBALANCERSETTINGS_CONSISTENTHASHLB.nested_types_by_name['HTTPCookie']
_CONNECTIONPOOLSETTINGS = DESCRIPTOR.message_types_by_name['ConnectionPoolSettings']
_CONNECTIONPOOLSETTINGS_TCPSETTINGS = _CONNECTIONPOOLSETTINGS.nested_types_by_name['TCPSettings']
_CONNECTIONPOOLSETTINGS_TCPSETTINGS_TCPKEEPALIVE = _CONNECTIONPOOLSETTINGS_TCPSETTINGS.nested_types_by_name['TcpKeepalive']
_CONNECTIONPOOLSETTINGS_HTTPSETTINGS = _CONNECTIONPOOLSETTINGS.nested_types_by_name['HTTPSettings']
_OUTLIERDETECTION = DESCRIPTOR.message_types_by_name['OutlierDetection']
_CLIENTTLSSETTINGS = DESCRIPTOR.message_types_by_name['ClientTLSSettings']
_LOCALITYLOADBALANCERSETTING = DESCRIPTOR.message_types_by_name['LocalityLoadBalancerSetting']
_LOCALITYLOADBALANCERSETTING_DISTRIBUTE = _LOCALITYLOADBALANCERSETTING.nested_types_by_name['Distribute']
_LOCALITYLOADBALANCERSETTING_DISTRIBUTE_TOENTRY = _LOCALITYLOADBALANCERSETTING_DISTRIBUTE.nested_types_by_name['ToEntry']
_LOCALITYLOADBALANCERSETTING_FAILOVER = _LOCALITYLOADBALANCERSETTING.nested_types_by_name['Failover']
_LOADBALANCERSETTINGS_SIMPLELB = _LOADBALANCERSETTINGS.enum_types_by_name['SimpleLB']
_CONNECTIONPOOLSETTINGS_HTTPSETTINGS_H2UPGRADEPOLICY = _CONNECTIONPOOLSETTINGS_HTTPSETTINGS.enum_types_by_name['H2UpgradePolicy']
_CLIENTTLSSETTINGS_TLSMODE = _CLIENTTLSSETTINGS.enum_types_by_name['TLSmode']
DestinationRule = _reflection.GeneratedProtocolMessageType('DestinationRule', (_message.Message,), {
  'DESCRIPTOR' : _DESTINATIONRULE,
  '__module__' : 'networking.v1beta1.destination_rule_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.DestinationRule)
  })
_sym_db.RegisterMessage(DestinationRule)

TrafficPolicy = _reflection.GeneratedProtocolMessageType('TrafficPolicy', (_message.Message,), {

  'PortTrafficPolicy' : _reflection.GeneratedProtocolMessageType('PortTrafficPolicy', (_message.Message,), {
    'DESCRIPTOR' : _TRAFFICPOLICY_PORTTRAFFICPOLICY,
    '__module__' : 'networking.v1beta1.destination_rule_pb2'
    # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.TrafficPolicy.PortTrafficPolicy)
    })
  ,
  'DESCRIPTOR' : _TRAFFICPOLICY,
  '__module__' : 'networking.v1beta1.destination_rule_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.TrafficPolicy)
  })
_sym_db.RegisterMessage(TrafficPolicy)
_sym_db.RegisterMessage(TrafficPolicy.PortTrafficPolicy)

Subset = _reflection.GeneratedProtocolMessageType('Subset', (_message.Message,), {

  'LabelsEntry' : _reflection.GeneratedProtocolMessageType('LabelsEntry', (_message.Message,), {
    'DESCRIPTOR' : _SUBSET_LABELSENTRY,
    '__module__' : 'networking.v1beta1.destination_rule_pb2'
    # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.Subset.LabelsEntry)
    })
  ,
  'DESCRIPTOR' : _SUBSET,
  '__module__' : 'networking.v1beta1.destination_rule_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.Subset)
  })
_sym_db.RegisterMessage(Subset)
_sym_db.RegisterMessage(Subset.LabelsEntry)

LoadBalancerSettings = _reflection.GeneratedProtocolMessageType('LoadBalancerSettings', (_message.Message,), {

  'ConsistentHashLB' : _reflection.GeneratedProtocolMessageType('ConsistentHashLB', (_message.Message,), {

    'HTTPCookie' : _reflection.GeneratedProtocolMessageType('HTTPCookie', (_message.Message,), {
      'DESCRIPTOR' : _LOADBALANCERSETTINGS_CONSISTENTHASHLB_HTTPCOOKIE,
      '__module__' : 'networking.v1beta1.destination_rule_pb2'
      # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.LoadBalancerSettings.ConsistentHashLB.HTTPCookie)
      })
    ,
    'DESCRIPTOR' : _LOADBALANCERSETTINGS_CONSISTENTHASHLB,
    '__module__' : 'networking.v1beta1.destination_rule_pb2'
    # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.LoadBalancerSettings.ConsistentHashLB)
    })
  ,
  'DESCRIPTOR' : _LOADBALANCERSETTINGS,
  '__module__' : 'networking.v1beta1.destination_rule_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.LoadBalancerSettings)
  })
_sym_db.RegisterMessage(LoadBalancerSettings)
_sym_db.RegisterMessage(LoadBalancerSettings.ConsistentHashLB)
_sym_db.RegisterMessage(LoadBalancerSettings.ConsistentHashLB.HTTPCookie)

ConnectionPoolSettings = _reflection.GeneratedProtocolMessageType('ConnectionPoolSettings', (_message.Message,), {

  'TCPSettings' : _reflection.GeneratedProtocolMessageType('TCPSettings', (_message.Message,), {

    'TcpKeepalive' : _reflection.GeneratedProtocolMessageType('TcpKeepalive', (_message.Message,), {
      'DESCRIPTOR' : _CONNECTIONPOOLSETTINGS_TCPSETTINGS_TCPKEEPALIVE,
      '__module__' : 'networking.v1beta1.destination_rule_pb2'
      # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.ConnectionPoolSettings.TCPSettings.TcpKeepalive)
      })
    ,
    'DESCRIPTOR' : _CONNECTIONPOOLSETTINGS_TCPSETTINGS,
    '__module__' : 'networking.v1beta1.destination_rule_pb2'
    # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.ConnectionPoolSettings.TCPSettings)
    })
  ,

  'HTTPSettings' : _reflection.GeneratedProtocolMessageType('HTTPSettings', (_message.Message,), {
    'DESCRIPTOR' : _CONNECTIONPOOLSETTINGS_HTTPSETTINGS,
    '__module__' : 'networking.v1beta1.destination_rule_pb2'
    # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.ConnectionPoolSettings.HTTPSettings)
    })
  ,
  'DESCRIPTOR' : _CONNECTIONPOOLSETTINGS,
  '__module__' : 'networking.v1beta1.destination_rule_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.ConnectionPoolSettings)
  })
_sym_db.RegisterMessage(ConnectionPoolSettings)
_sym_db.RegisterMessage(ConnectionPoolSettings.TCPSettings)
_sym_db.RegisterMessage(ConnectionPoolSettings.TCPSettings.TcpKeepalive)
_sym_db.RegisterMessage(ConnectionPoolSettings.HTTPSettings)

OutlierDetection = _reflection.GeneratedProtocolMessageType('OutlierDetection', (_message.Message,), {
  'DESCRIPTOR' : _OUTLIERDETECTION,
  '__module__' : 'networking.v1beta1.destination_rule_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.OutlierDetection)
  })
_sym_db.RegisterMessage(OutlierDetection)

ClientTLSSettings = _reflection.GeneratedProtocolMessageType('ClientTLSSettings', (_message.Message,), {
  'DESCRIPTOR' : _CLIENTTLSSETTINGS,
  '__module__' : 'networking.v1beta1.destination_rule_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.ClientTLSSettings)
  })
_sym_db.RegisterMessage(ClientTLSSettings)

LocalityLoadBalancerSetting = _reflection.GeneratedProtocolMessageType('LocalityLoadBalancerSetting', (_message.Message,), {

  'Distribute' : _reflection.GeneratedProtocolMessageType('Distribute', (_message.Message,), {

    'ToEntry' : _reflection.GeneratedProtocolMessageType('ToEntry', (_message.Message,), {
      'DESCRIPTOR' : _LOCALITYLOADBALANCERSETTING_DISTRIBUTE_TOENTRY,
      '__module__' : 'networking.v1beta1.destination_rule_pb2'
      # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.LocalityLoadBalancerSetting.Distribute.ToEntry)
      })
    ,
    'DESCRIPTOR' : _LOCALITYLOADBALANCERSETTING_DISTRIBUTE,
    '__module__' : 'networking.v1beta1.destination_rule_pb2'
    # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.LocalityLoadBalancerSetting.Distribute)
    })
  ,

  'Failover' : _reflection.GeneratedProtocolMessageType('Failover', (_message.Message,), {
    'DESCRIPTOR' : _LOCALITYLOADBALANCERSETTING_FAILOVER,
    '__module__' : 'networking.v1beta1.destination_rule_pb2'
    # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.LocalityLoadBalancerSetting.Failover)
    })
  ,
  'DESCRIPTOR' : _LOCALITYLOADBALANCERSETTING,
  '__module__' : 'networking.v1beta1.destination_rule_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.LocalityLoadBalancerSetting)
  })
_sym_db.RegisterMessage(LocalityLoadBalancerSetting)
_sym_db.RegisterMessage(LocalityLoadBalancerSetting.Distribute)
_sym_db.RegisterMessage(LocalityLoadBalancerSetting.Distribute.ToEntry)
_sym_db.RegisterMessage(LocalityLoadBalancerSetting.Failover)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\037istio.io/api/networking/v1beta1'
  _DESTINATIONRULE.fields_by_name['host']._options = None
  _DESTINATIONRULE.fields_by_name['host']._serialized_options = b'\342A\001\002'
  _SUBSET_LABELSENTRY._options = None
  _SUBSET_LABELSENTRY._serialized_options = b'8\001'
  _SUBSET.fields_by_name['name']._options = None
  _SUBSET.fields_by_name['name']._serialized_options = b'\342A\001\002'
  _LOADBALANCERSETTINGS_CONSISTENTHASHLB_HTTPCOOKIE.fields_by_name['name']._options = None
  _LOADBALANCERSETTINGS_CONSISTENTHASHLB_HTTPCOOKIE.fields_by_name['name']._serialized_options = b'\342A\001\002'
  _LOADBALANCERSETTINGS_CONSISTENTHASHLB_HTTPCOOKIE.fields_by_name['ttl']._options = None
  _LOADBALANCERSETTINGS_CONSISTENTHASHLB_HTTPCOOKIE.fields_by_name['ttl']._serialized_options = b'\342A\001\002'
  _LOADBALANCERSETTINGS_SIMPLELB.values_by_name["LEAST_CONN"]._options = None
  _LOADBALANCERSETTINGS_SIMPLELB.values_by_name["LEAST_CONN"]._serialized_options = b'\010\001'
  _OUTLIERDETECTION.fields_by_name['consecutive_errors']._options = None
  _OUTLIERDETECTION.fields_by_name['consecutive_errors']._serialized_options = b'\030\001'
  _CLIENTTLSSETTINGS.fields_by_name['mode']._options = None
  _CLIENTTLSSETTINGS.fields_by_name['mode']._serialized_options = b'\342A\001\002'
  _LOCALITYLOADBALANCERSETTING_DISTRIBUTE_TOENTRY._options = None
  _LOCALITYLOADBALANCERSETTING_DISTRIBUTE_TOENTRY._serialized_options = b'8\001'
  _DESTINATIONRULE._serialized_start=233
  _DESTINATIONRULE._serialized_end=445
  _TRAFFICPOLICY._serialized_start=448
  _TRAFFICPOLICY._serialized_end=1308
  _TRAFFICPOLICY_PORTTRAFFICPOLICY._serialized_start=901
  _TRAFFICPOLICY_PORTTRAFFICPOLICY._serialized_end=1308
  _SUBSET._serialized_start=1311
  _SUBSET._serialized_end=1554
  _SUBSET_LABELSENTRY._serialized_start=1497
  _SUBSET_LABELSENTRY._serialized_end=1554
  _LOADBALANCERSETTINGS._serialized_start=1557
  _LOADBALANCERSETTINGS._serialized_end=2519
  _LOADBALANCERSETTINGS_CONSISTENTHASHLB._serialized_start=1953
  _LOADBALANCERSETTINGS_CONSISTENTHASHLB._serialized_end=2392
  _LOADBALANCERSETTINGS_CONSISTENTHASHLB_HTTPCOOKIE._serialized_start=2271
  _LOADBALANCERSETTINGS_CONSISTENTHASHLB_HTTPCOOKIE._serialized_end=2380
  _LOADBALANCERSETTINGS_SIMPLELB._serialized_start=2394
  _LOADBALANCERSETTINGS_SIMPLELB._serialized_end=2506
  _CONNECTIONPOOLSETTINGS._serialized_start=2522
  _CONNECTIONPOOLSETTINGS._serialized_end=3607
  _CONNECTIONPOOLSETTINGS_TCPSETTINGS._serialized_start=2712
  _CONNECTIONPOOLSETTINGS_TCPSETTINGS._serialized_end=3089
  _CONNECTIONPOOLSETTINGS_TCPSETTINGS_TCPKEEPALIVE._serialized_start=2949
  _CONNECTIONPOOLSETTINGS_TCPSETTINGS_TCPKEEPALIVE._serialized_end=3089
  _CONNECTIONPOOLSETTINGS_HTTPSETTINGS._serialized_start=3092
  _CONNECTIONPOOLSETTINGS_HTTPSETTINGS._serialized_end=3607
  _CONNECTIONPOOLSETTINGS_HTTPSETTINGS_H2UPGRADEPOLICY._serialized_start=3544
  _CONNECTIONPOOLSETTINGS_HTTPSETTINGS_H2UPGRADEPOLICY._serialized_end=3607
  _OUTLIERDETECTION._serialized_start=3610
  _OUTLIERDETECTION._serialized_end=4260
  _CLIENTTLSSETTINGS._serialized_start=4263
  _CLIENTTLSSETTINGS._serialized_end=4729
  _CLIENTTLSSETTINGS_TLSMODE._serialized_start=4665
  _CLIENTTLSSETTINGS_TLSMODE._serialized_end=4729
  _LOCALITYLOADBALANCERSETTING._serialized_start=4732
  _LOCALITYLOADBALANCERSETTING._serialized_end=5278
  _LOCALITYLOADBALANCERSETTING_DISTRIBUTE._serialized_start=5053
  _LOCALITYLOADBALANCERSETTING_DISTRIBUTE._serialized_end=5230
  _LOCALITYLOADBALANCERSETTING_DISTRIBUTE_TOENTRY._serialized_start=5177
  _LOCALITYLOADBALANCERSETTING_DISTRIBUTE_TOENTRY._serialized_end=5230
  _LOCALITYLOADBALANCERSETTING_FAILOVER._serialized_start=5232
  _LOCALITYLOADBALANCERSETTING_FAILOVER._serialized_end=5278
# @@protoc_insertion_point(module_scope)
