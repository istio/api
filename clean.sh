PATTERNS="_deepcopy.gen.go .gen.json .pb.go .pb.html _json.gen.go .gen.yaml"
# PATTERNS="_deepcopy.gen.go .pb.go .pb.html _json.gen.go"
shopt -s globstar

for p in $PATTERNS; do
    rm -f **/*$p
done
rm -r python/istio_api
