#!/bin/bash

set -o errexit
set -o nounset

dep ensure --vendor-only || (echo "failed to update vendor directory" && exit -1)
make clean-kube-apis generate-kube-apis
