#!/bin/bash
# shellcheck disable=SC2045,SC2016
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

shopt -s globstar

# validate.sh checks protos are valid.
# This is not currently a part of presubmit as there is a circular dependency on istio/istio repo
# for istioctl, which would prevent adding new fields (since they would fail the validation command).

# Check if all examples are valid
for proto in $(find . -name "*.proto" | grep -v ^./common-protos | grep -v ^./envoy); do
    echo "Checking ${proto}..."
    # For each proto, we will remove comments (`//` or `// ` at start of line),
    # then find texts between ```yaml and ```. We separate each by ---, strip any ... (some docs elide resources),
    # then pass to istioctl validate
    < "${proto}" \
      sed 's/^\/\/ //g' | sed 's/^\/\///g' \
    | sed -n '/^```yaml/,/^```/ p' | sed 's/^```yaml/\-\-\-/g' \
    | grep -v '```' | grep -v '\.\.\.' \
    | istioctl validate -f -
done

# Check if schemas are identical
# Note: once we migrate to v1 CRDs we can move to scripts/validate_crds.py
# Note: this misses package wide documentation
for file in $(ls networking/v1beta1/*.gen.json); do
    echo "Comparing $(basename "${file}")"
    diff <(< "${file}" jq --sort-keys) <(< "networking/v1alpha3/$(basename "${file}")" sed 's/v1alpha3/v1beta1/g' | jq --sort-keys)
done
