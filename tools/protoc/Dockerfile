FROM golang:1.12 as build-env

RUN mkdir -p /out_build_env/usr/bin
RUN go get -u -v -ldflags '-w -s' github.com/golang/protobuf/protoc-gen-go
# this repo 404s...
#RUN go get -u -v -ldflags '-w -s' github.com/golang/protobuf/protoc-gen-gofast
RUN go get -u -v -ldflags '-w -s' github.com/gogo/protobuf/protoc-gen-gogo
RUN go get -u -v -ldflags '-w -s' github.com/gogo/protobuf/protoc-gen-gogofaster
RUN go get -u -v -ldflags '-w -s' github.com/gogo/protobuf/protoc-gen-gogoslick
RUN go get -u -v -ldflags '-w -s' istio.io/tools/protoc-gen-docs
RUN cp -aR ${GOPATH}/bin/protoc-gen* /out_build_env/usr/bin/

#FROM znly/upx as packer
#COPY --from=protoc_builder /out_packer/ /out_packer/
#RUN upx --lzma \
#        /out_packer/usr/bin/protoc \
#        /out_packer/usr/bin/grpc_* \
#        /out_packer/usr/bin/protoc-gen-*

FROM gcr.io/distroless/cc

#COPY --from=packer /out_packer/ /
COPY --from=build-env /out_build_env/ /

ENTRYPOINT ["/usr/bin/protoc", "-I/protobuf"]
