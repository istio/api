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

set -eEou pipefail

# sync.sh keeps pairs of files in sync. Specifically, this is used for multi version protobuf.
# These files have a unique (comment) header per version, but should have identical proto definitions.
# To pair two files, The +istio.io/sync-{from,start} tags can be added.
# For example: In v1beta1/service_entry.proto, we can add `+istio.io/sync-from:networking/v1alpha3/service_entry.proto`.
# Next, we add `+istio.io/sync-start` to the `v1alpha3/service_entry.proto` file

FROM_TAG="+istio.io/sync-from"
START_TAG="+istio.io/sync-start"

BIG_NUMBER=100000 # If our files are longer than this we have bigger issues..

find . -name '*.proto'  -not -path "./common-protos/*" -print0 | while read -r -d $'\0' file; do
  res="$(grep "${FROM_TAG}" "${file}" || true)"
  if [[ "${res}" != "" ]]; then # We need to sync this file
    replacement="$(echo "${res}" | cut -d: -f2)"
    echo "Syncing ${file} from ${replacement}"
    # First we retain the top section of the file, everything before FROM_TAG
    header="$(grep "${FROM_TAG}" "${file}" -B "${BIG_NUMBER}")"
    # Then we copy the bottom section of the replacement file, everything after START_TAG
    body="$(grep "${START_TAG}" "${replacement}" -A "${BIG_NUMBER}")"
    # And merge them into a single file
    echo "${header}" > "${file}"
    # We skip the first line of the replacement to avoid copying the start tag
    echo "${body}" | tail -n +2 >> "${file}"
  fi
done
