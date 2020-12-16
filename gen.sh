# Generate all protos
prototool generate networking
prototool generate security
prototool generate type
prototool generate analysis
prototool generate authentication
prototool generate meta

# These folders do not have the full plugins used, as they are not full CRDs.
# We pass them a custom configuration to exclude the non-required files
prototool generate operator --config-data "`cat prototool.noncrd.json`"
prototool generate mcp --config-data "`cat prototool.noncrd.json`"
prototool generate mesh --config-data "`cat prototool.noncrd.json`"

# Custom hacks to post-process some outputs
go run ./operator/fixup_structs/main.go -f operator/v1alpha1/operator.pb.go
go run ./operator/fixup_structs/main.go -f mesh/v1alpha1/config.pb.go
go run ./operator/fixup_structs/main.go -f mesh/v1alpha1/network.pb.go
go run ./operator/fixup_structs/main.go -f mesh/v1alpha1/proxy.pb.go

# Generate CRDs and open-api schema
cue-gen -paths=common-protos -f=./cue.yaml
cue-gen -paths=common-protos -f=./cue.yaml --crd=true -snake=jwksUri,apiKeys,apiSpecs,includedPaths,jwtHeaders,triggerRules,excludedPaths,mirrorPercent
