package(default_visibility = ["//visibility:public"])

load("@org_pubref_rules_protobuf//gogo:rules.bzl", "gogoslick_proto_library")

gogoslick_proto_library(
    name = "go_default_library",
    protos = [":protos"],
    importmap = {
        "google/protobuf/any.proto": "github.com/gogo/protobuf/types",
        "google/protobuf/duration.proto": "github.com/gogo/protobuf/types",
    },
    imports = [
        "external/com_github_google_protobuf/src",
        "../../external/com_github_google_protobuf/src",
    ],
    inputs = [
        "@com_github_google_protobuf//:well_known_protos",
    ],
    deps = [
        "@com_github_gogo_protobuf//types:go_default_library",
    ],
    verbose = 0,
)

filegroup(
    name = "protos",
    srcs = [
        ":status.proto",
        ":code.proto",
        ":error_details.proto",
    ]
)

genrule(
    name = "status_proto",
    srcs = ["@com_github_googleapis_googleapis//:status_proto"],
    outs = ["status.proto"],
    cmd = "cat $(location @com_github_googleapis_googleapis//:status_proto) > $@",
)

genrule(
    name = "code_proto",
    srcs = ["@com_github_googleapis_googleapis//:code_proto"],
    outs = ["code.proto"],
    cmd = "cat $(location @com_github_googleapis_googleapis//:code_proto) > $@",
)

genrule(
    name = "error_details_proto",
    srcs = ["@com_github_googleapis_googleapis//:error_details_proto"],
    outs = ["error_details.proto"],
    cmd = "cat $(location @com_github_googleapis_googleapis//:error_details_proto) > $@",
)

# load("@org_pubref_rules_protobuf//cpp:rules.bzl", "cc_proto_library")
# 
# cc_proto_library(
#     name = "cc_status_proto",
#     protos = [
#         "google/rpc/status.proto",
#     ],
#     imports = [
#         "../../external/com_github_google_protobuf/src",
#     ],
#     verbose = 0,
# )
# 
# filegroup(
#     name = "status_proto",
#     srcs = [ "google/rpc/status.proto" ],
# )
# 
# filegroup(
#     name = "code_proto",
#     srcs = [ "google/rpc/code.proto" ],
# )
# 
# filegroup(
#     name = "label_proto",
#     srcs = [
#         "google/api/label.proto",
#     ],
# )
# 
# filegroup(
#     name = "metric_proto",
#     srcs = [
#         "google/api/metric.proto",
#     ],
# )
