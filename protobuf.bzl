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

def protobuf_repositories(load_repo=True, bind=True):
    if load_repo:
        native.git_repository(
            name = "com_google_protobuf",
            commit = "593e917c176b5bc5aafa57bf9f6030d749d91cd5",  # v3.2.0
            remote = "https://github.com/google/protobuf.git",
        )

    if bind:
        native.bind(
            name = "protoc",
            actual = "@com_google_protobuf//:protoc",
        )

        native.bind(
            name = "protocol_compiler",
            actual = "@com_google_protobuf//:protoc",
        )

        native.bind(
            name = "protobuf",
            actual = "@com_google_protobuf//:protobuf",
        )

        native.bind(
            name = "cc_wkt_protos",
            actual = "@com_google_protobuf//:cc_wkt_protos",
        )

        native.bind(
            name = "cc_wkt_protos_genproto",
            actual = "@com_google_protobuf//:cc_wkt_protos_genproto",
        )

        native.bind(
            name = "protobuf_compiler",
            actual = "@com_google_protobuf//:protoc_lib",
        )

        native.bind(
            name = "protobuf_clib",
            actual = "@com_google_protobuf//:protoc_lib",
        )

