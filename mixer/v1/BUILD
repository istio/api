load("@org_pubref_rules_protobuf//gogo:rules.bzl", "gogoslick_proto_library")
load("@com_google_protobuf//:protobuf.bzl", "cc_proto_library")

gogoslick_proto_library(
    name = "go_default_library",
    importmap = {
        "gogoproto/gogo.proto": "github.com/gogo/protobuf/gogoproto",
        "google/protobuf/any.proto": "github.com/gogo/protobuf/types",
        "google/protobuf/timestamp.proto": "github.com/gogo/protobuf/types",
        "google/protobuf/duration.proto": "github.com/gogo/protobuf/types",
    },
    imports = [
        "external/com_github_gogo_protobuf",
        "external/com_github_google_protobuf/src",
    ],
    inputs = [
        "@com_github_google_protobuf//:well_known_protos",
        "@com_github_gogo_protobuf//gogoproto:go_default_library_protos",
    ],
    protos = [
        "attributes.proto",
        "check.proto",
        "code.proto",
        "error_details.proto",
        "report.proto",
        "service.proto",
        "status.proto",
    ],
    verbose = 0,
    visibility = ["//visibility:public"],
    with_grpc = True,
    deps = [
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
    ],
)
