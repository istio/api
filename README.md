# Istio APIs and Common Configuration Definitions

This repository defines component-level APIs and common configuration formats for the Istio
platform. These definitions are specified using the [protobuf](https://github.com/google/protobuf)
syntax.

All other Istio repositories can take a dependency on the api
repository. This repository *will not* depend on any other repos

## Standard vocabulary

All components of an Istio installation operate on a shared vocabulary of attributes,
as defined in this repo.

## Style guide

When designing proto-based APIs in the Istio project, please follow the
[Istio API Style Guide](./STYLE-GUIDE.md).

## Updating

After the [protobuf](https://github.com/google/protobuf) definitions are updated, the corresponding `*pb.go` and `_pb2.py` files must be generated by running `scripts/generate-protos.sh` and submitted as part of the same PR as the updated definitions.

If releasing a new tagged version, please update python/istio-api/setup.py version to reflect.
