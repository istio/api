load("@org_pubref_rules_protobuf//gogo:rules.bzl", "gogoslick_proto_library")

gogoslick_proto_library(
    name = "mixer/v1/config/descriptor",
    importmap = {
        "google/protobuf/duration.proto": "github.com/gogo/protobuf/types",
    },
    imports = [
        "external/com_github_google_protobuf/src",
    ],
    inputs = [
        "@com_github_google_protobuf//:well_known_protos",
    ],
    protos = [
        ":protos",
    ],
    verbose = 0,
    visibility = ["//visibility:public"],
    with_grpc = False,
    deps = [
        "@com_github_gogo_protobuf//sortkeys:go_default_library",
        "@com_github_gogo_protobuf//types:go_default_library",
    ],
)

filegroup(
    name = "protos",
    srcs = glob(["*.proto"]),
    visibility = ["//visibility:public"],
)
