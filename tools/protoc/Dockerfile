# Based on: https://github.com/znly/docker-protobuf/blob/master/Dockerfile
# Modified to slim down the image and adjust for Istio

FROM golang:1.12 as build-env

RUN mkdir -p /out_build_env/usr/bin
# tools repo should pull these deps in automagically, time will tell once its working
#RUN go get -u -v github.com/golang/protobuf/proto
#RUN go get -u -v gopkg.in/russross/blackfriday.v2
RUN ls -lR ${GOPATH}
RUN go get -u -v -ldflags '-w -s' github.com/golang/protobuf/protoc-gen-go
#RUN go get -u -v -ldflags '-w -s' github.com/golang/protobuf/protoc-gen-gofast
RUN go get -u -v -ldflags '-w -s' github.com/gogo/protobuf/protoc-gen-gogo
RUN go get -u -v -ldflags '-w -s' github.com/gogo/protobuf/protoc-gen-gogofaster
RUN go get -u -v -ldflags '-w -s' github.com/gogo/protobuf/protoc-gen-gogoslick
RUN go get -u -v -ldflags '-w -s' istio.io/tools/protoc-gen-docs
RUN cp -aR ${GOPATH}/bin/protoc-gen* /out_build_env/usr/bin/

FROM znly/upx as packer
COPY --from=protoc_builder /out_packer/ /out_packer/
RUN upx --lzma \
        /out_packer/usr/bin/protoc \
        /out_packer/usr/bin/grpc_* \
        /out_packer/usr/bin/protoc-gen-*

FROM gcr.io/distroless/cc

COPY --from=packer /out_packer/ /
COPY --from=build-env /out_build_env/ /

ENTRYPOINT ["/usr/bin/protoc", "-I/protobuf"]
