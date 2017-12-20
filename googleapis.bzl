# Copyright 2017 Istio Authors. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
################################################################################
#

def googleapis_repositories(bind=True):
    GOOGLEAPIS_BUILD_FILE = """
package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_prefix")
go_prefix("github.com/googleapis/googleapis")

load("@org_pubref_rules_protobuf//cpp:rules.bzl", "cc_proto_library")

cc_proto_library(
    name = "cc_status_proto",
    protos = [
        "google/rpc/status.proto",
    ],
    imports = [
        "../../external/com_github_google_protobuf/src",
    ],
    verbose = 0,
)

filegroup(
    name = "status_proto",
    srcs = [ "google/rpc/status.proto" ],
)

filegroup(
    name = "code_proto",
    srcs = [ "google/rpc/code.proto" ],
)

filegroup(
    name = "label_proto",
    srcs = [
        "google/api/label.proto",
    ],
)

filegroup(
    name = "metric_proto",
    srcs = [
        "google/api/metric.proto",
    ],
)

load("@com_google_protobuf//:protobuf.bzl", "cc_proto_library")

cc_proto_library(
    name = "rpc_status_proto",
    srcs = [
        "google/rpc/status.proto",
    ],
    visibility = ["//visibility:public"],
    protoc = "//external:protoc",
    default_runtime = "//external:protobuf",
    deps = [
        "//external:cc_wkt_protos",
    ],
)

"""
    native.new_git_repository(
        name = "com_github_googleapis_googleapis",
        build_file_content = GOOGLEAPIS_BUILD_FILE,
        commit = "13ac2436c5e3d568bd0e938f6ed58b77a48aba15", # Oct 21, 2016 (only release pre-dates sha)
        remote = "https://github.com/googleapis/googleapis.git",
    )

    if bind:
        native.bind(
            name = "rpc_status_proto",
            actual = "@com_github_googleapis_googleapis//:rpc_status_proto",
        )
        native.bind(
            name = "rpc_status_proto_genproto",
            actual = "@com_github_googleapis_googleapis//:rpc_status_proto_genproto",
        )

