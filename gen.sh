#!/bin/bash

# Copyright Istio Authors

#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.

set -eu

# Generate all protos
buf generate \
  --path networking \
  --path security \
  --path type \
  --path analysis \
  --path authentication \
  --path meta \
  --path telemetry \
  --path extensions

# These folders do not have the full plugins used, as they are not full CRDs.
# We pass them a custom configuration to exclude the non-required files
buf generate --template buf.gen-noncrd.yaml \
  --path mcp \
  --path mesh

# These plugins are sent to Envoy, which uses golang/protobuf, so do not use gogo
buf generate --template buf.gen-golang.yaml \
  --path envoy
