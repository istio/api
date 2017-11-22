#!/bin/bash

set -o errexit
set -o nounset

bazel build //...

genfiles=$(bazel info bazel-genfiles)
files=$(find -L ${genfiles} -name "*.pb.go")

for src in ${files}; do
    dst=${src##${genfiles}/}
    if [ -d "$(dirname ${dst})" ]; then
        install -m 0640 ${src} ${dst}
    fi
done

# Delete rewritten source
rm mixer/v1/config/cfg.pb.go

# Copy vendorized gogoprotos (not compatible with google.golang.org/api/)
external=${genfiles}/external/com_github_googleapis_googleapis
googleapis=$(find -L ${external} -name "*.pb.go")
for src in ${googleapis}; do
    dst=vendor/github.com/googleapis/googleapis/${src##${external}/}
    if [ -d "$(dirname ${dst})" ]; then
        install -m 0640 ${src} ${dst}
    fi
done
