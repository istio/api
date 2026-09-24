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

# Tests for check-release-channel.sh

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CHECK_SCRIPT="${SCRIPT_DIR}/check-release-channel.sh"
TMPDIR=$(mktemp -d)
trap 'rm -rf "${TMPDIR}"' EXIT

pass=0
fail=0

run_test() {
  local name="$1" want_exit="$2"
  shift 2
  if output=$("$@" 2>&1); then
    got_exit=0
  else
    got_exit=$?
  fi
  if [[ ${got_exit} -eq ${want_exit} ]]; then
    echo "PASS: ${name}"
    pass=$((pass + 1))
  else
    echo "FAIL: ${name} (want exit=${want_exit}, got exit=${got_exit})"
    echo "${output}"
    fail=$((fail + 1))
  fi
}

# Set up a temp git repo with a base proto on 'base' branch,
# then switch to a 'work' branch for changes.
setup_repo() {
  rm -rf "${TMPDIR}/repo"
  mkdir -p "${TMPDIR}/repo/security/v1beta1"
  cd "${TMPDIR}/repo"
  git init -q
  git checkout -q -b base
  cat > security/v1beta1/test.proto << 'EOF'
syntax = "proto3";
package istio.security.v1beta1;

message Source {
  string principal = 1;
  repeated string namespaces = 2;
}
EOF
  cp "${CHECK_SCRIPT}" ./check.sh
  chmod +x ./check.sh
  git add -A && git commit -q -m "base"
  git checkout -q -b work
}

# Test 1: no proto changes -> pass
setup_repo
run_test "no changes" 0 ./check.sh base

# Test 2: new field without annotation -> fail
setup_repo
cat >> security/v1beta1/test.proto << 'EOF'

message Rule {
  // some comment
  string bad_field = 1;
}
EOF
git add -A && git commit -q -m "add bad field"
run_test "field without annotation fails" 1 ./check.sh base

# Test 3: new field with annotation -> pass
setup_repo
cat >> security/v1beta1/test.proto << 'EOF'

message Rule {
  // +cue-gen:Rule:releaseChannel:extended
  string good_field = 1;
}
EOF
git add -A && git commit -q -m "add good field"
run_test "field with annotation passes" 0 ./check.sh base

# Test 4: repeated field without annotation -> fail
setup_repo
cat >> security/v1beta1/test.proto << 'EOF'

message Rule {
  repeated string bad_repeated = 1;
}
EOF
git add -A && git commit -q -m "add bad repeated"
run_test "repeated field without annotation fails" 1 ./check.sh base

# Test 5: map field without annotation -> fail
setup_repo
cat >> security/v1beta1/test.proto << 'EOF'

message Rule {
  map<string, string> bad_map = 1;
}
EOF
git add -A && git commit -q -m "add bad map"
run_test "map field without annotation fails" 1 ./check.sh base

# Test 6: enum value (not a field) without annotation -> pass
setup_repo
cat >> security/v1beta1/test.proto << 'EOF'

enum Mode {
  DEFAULT = 0;
  STRICT = 1;
}
EOF
git add -A && git commit -q -m "add enum"
run_test "enum values ignored" 0 ./check.sh base

# Test 7: mixed good and bad fields -> fail
setup_repo
cat >> security/v1beta1/test.proto << 'EOF'

message Rule {
  // +cue-gen:Rule:releaseChannel:extended
  string good = 1;

  string bad = 2;
}
EOF
git add -A && git commit -q -m "add mixed"
run_test "mixed fields catches bad only" 1 ./check.sh base

# Test 8: common-protos changes are ignored
setup_repo
mkdir -p common-protos/google
cat > common-protos/google/test.proto << 'EOF'
syntax = "proto3";
message Foo {
  string no_annotation = 1;
}
EOF
git add -A && git commit -q -m "add common-proto"
run_test "common-protos ignored" 0 ./check.sh base

echo ""
echo "Results: ${pass} passed, ${fail} failed"
[[ ${fail} -eq 0 ]]
