#!/usr/bin/env bash
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

branch="${1:?branch to compare against}"

# buf breaking --against ".git#branch=${branch}" does not work due to https://github.com/bufbuild/buf/issues/1003. Workaround it.

d="$(mktemp -d )"
trap '
rm -rf "${d}"
' EXIT

cwd="${PWD}"
pushd "${d}" > /dev/null || exit
git clone "${cwd}" -q -b "${branch}" api > /dev/null
cd api
buf build -o proto.bin
popd  > /dev/null|| exit

buf breaking --against "${d}/api/proto.bin"
