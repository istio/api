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

# check-release-channel.sh
#
# Ensures that newly added proto fields include a
# +cue-gen:<APIName>:releaseChannel:extended annotation in their comment block.
#
# Per GUIDELINES.md, new fields added to stable v1 APIs must go through the
# "extended" release channel before being promoted to stable. This is indicated
# by a comment annotation above the field definition. Existing stable fields
# intentionally lack this annotation, so we use a diff-based approach: only
# fields added relative to the base branch are checked.
#
# How it works:
#   1. Find proto files changed vs the base branch (excluding common-protos/).
#   2. Parse the git diff to extract added lines with their line numbers.
#   3. Identify lines that are proto field definitions (type name = N;),
#      skipping enum values, reserved statements, and options.
#   4. For each new field, walk upward through the contiguous comment block
#      immediately above it, looking for the releaseChannel annotation.
#   5. Report any fields missing the annotation and exit non-zero.
#
# Usage:
#   ./scripts/check-release-channel.sh <base-branch>
#   ./scripts/check-release-channel.sh master
#
# See https://github.com/istio/api/issues/3147

set -euo pipefail

branch="${1:-master}"
errors=0

# Step 1: Find changed/added proto files, excluding third-party common-protos.
changed_protos=$(git diff "${branch}" --name-only --diff-filter=AM -- '*.proto' | grep -v common-protos || true)
if [[ -z "${changed_protos}" ]]; then
  exit 0
fi

# check_comment_block: starting from the line above a field definition,
# walk upward through comment lines (// ...) and blank lines. If we find
# a +cue-gen:*:releaseChannel: annotation, return 0 (success). If we hit
# a non-comment, non-blank line (previous field, message boundary, etc.),
# stop and return 1 (not found). This prevents matching annotations that
# belong to a different field.
check_comment_block() {
  local file="$1" lineno="$2"
  local i=$((lineno - 1))
  while [[ ${i} -ge 1 ]]; do
    local line
    line=$(sed -n "${i}p" "${file}")
    # Comment line — check for the annotation
    if echo "${line}" | grep -qP '^\s*//'; then
      if echo "${line}" | grep -q '+cue-gen:.*:releaseChannel:'; then
        return 0
      fi
      i=$((i - 1))
      continue
    fi
    # Blank line — skip (comments may have a gap before the field)
    if echo "${line}" | grep -qP '^\s*$'; then
      i=$((i - 1))
      continue
    fi
    # Any other line (field, message, etc.) — stop searching
    break
  done
  return 1
}

for file in ${changed_protos}; do
  # Step 2: Parse "git diff -U0" to get added lines with their new-file line numbers.
  # -U0 means no context lines, so we only see actual additions.
  # The awk script reads @@ hunk headers to get the starting line number,
  # then counts up for each "+" line (skipping the "+++ b/file" header).
  added_lines=$(git diff "${branch}" -U0 -- "${file}" | awk '
    /^@@/ { match($0, /\+([0-9]+)/, a); nr = a[1]; next }
    /^\+/ && !/^\+\+\+/ { print nr ":" substr($0, 2); nr++ }
  ')
  [[ -z "${added_lines}" ]] && continue

  while IFS= read -r entry; do
    [[ -z "${entry}" ]] && continue
    lineno="${entry%%:*}"
    content="${entry#*:}"

    # Step 3: Match proto field definitions.
    # Pattern: [repeated|optional] type name = N; or map<K,V> name = N;
    # This matches:
    #   string foo = 1;
    #   repeated string foo = 2;
    #   google.protobuf.Duration timeout = 3;
    #   map<string, string> labels = 4;
    # But NOT enum values (single word before =): SOME_VALUE = 0;
    echo "${content}" | grep -qP '^\s*((repeated|optional)\s+)?((\w[\w.]*\s+\w+)|(map<.*>\s+\w+))\s*=\s*\d+\s*[;\[]' || continue
    # Skip reserved/option/comment-only lines that might match the pattern
    echo "${content}" | grep -qP '^\s*(reserved|option|//)\s' && continue

    # Step 4: Walk upward through the comment block above this field.
    if ! check_comment_block "${file}" "${lineno}"; then
      # Step 5: Report the error.
      echo "ERROR: ${file}:${lineno}: new field missing releaseChannel annotation"
      echo "  ${content}"
      errors=$((errors + 1))
    fi
  done <<< "${added_lines}"
done

if [[ ${errors} -gt 0 ]]; then
  echo ""
  echo "Found ${errors} new field(s) without releaseChannel annotation."
  echo "Add: // +cue-gen:<APIName>:releaseChannel:extended"
  echo "See GUIDELINES.md for details."
  exit 1
fi
