load("@org_pubref_rules_protobuf//gogo:rules.bzl", "gogoslick_proto_library")
load("@com_google_protobuf//:protobuf.bzl", "cc_proto_library")

gogoslick_proto_library(
    name = "go_default_library",
    importmap = {
        "gogoproto/gogo.proto": "github.com/gogo/protobuf/gogoproto",
        "google/protobuf/any.proto": "github.com/gogo/protobuf/types",
        "google/protobuf/timestamp.proto": "github.com/gogo/protobuf/types",
        "google/protobuf/duration.proto": "github.com/gogo/protobuf/types",
        "google/rpc/status.proto": "github.com/googleapis/googleapis/google/rpc",
        "google/rpc/code.proto": "github.com/googleapis/googleapis/google/rpc",
        "google/rpc/error_details.proto": "github.com/googleapis/googleapis/google/rpc",
    },
    imports = [
        "external/com_github_gogo_protobuf",
        "external/com_github_google_protobuf/src",
        "vendor/github.com/googleapis/googleapis",
    ],
    inputs = [
        "//vendor/github.com/googleapis/googleapis/google/rpc:rpc_status_protos",
        "@com_github_google_protobuf//:well_known_protos",
        "@com_github_gogo_protobuf//gogoproto:go_default_library_protos",
    ],
    protos = [
        "attributes.proto",
        "check.proto",
        "report.proto",
        "service.proto",
    ],
    verbose = 0,
    visibility = ["//visibility:public"],
    with_grpc = True,
    deps = [
        "//vendor/github.com/googleapis/googleapis/google/rpc:go_default_library",
        "@com_github_gogo_protobuf//gogoproto:go_default_library",
        "@com_github_gogo_protobuf//sortkeys:go_default_library",
        "@com_github_gogo_protobuf//types:go_default_library",
    ],
)

filegroup(
    name = "attributes",
    srcs = ["attributes.proto"],
    visibility = ["//visibility:public"],
)

filegroup(
    name = "attributes_file",
    srcs = ["global_dictionary.yaml"],
    visibility = ["//visibility:public"],
)

cc_proto_library(
    name = "cc_proto",
    srcs = glob(
        ["*.proto"],
    ),
    default_runtime = "//external:protobuf",
    protoc = "//external:protoc",
    visibility = ["//visibility:public"],
    deps = [
        "//external:cc_gogoproto",
        "//external:cc_wkt_protos",
        "//external:rpc_status_proto",
    ],
)
