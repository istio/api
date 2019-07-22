FROM golang:alpine as cue_builder
RUN apk update
RUN apk add --no-cache build-base git

ENV OUTDIR=/out
RUN mkdir -p ${OUTDIR}/usr/bin/
ENV GOPATH=/go \
    PATH=/go/bin/:$PATH
RUN go get -u -v -ldflags '-w -s' \
        istio.io/tools/openapi/cue \
        && install -c ${GOPATH}/bin/cue ${OUTDIR}/usr/bin/ && \
    apk del git

FROM alpine:3.7
COPY --from=cue_builder /out/ /

RUN apk add --no-cache curl && \
    mkdir -p /protobuf/google/protobuf && \
        for f in any duration descriptor empty struct timestamp wrappers; do \
            curl -L -o /protobuf/google/protobuf/${f}.proto https://raw.githubusercontent.com/google/protobuf/master/src/google/protobuf/${f}.proto; \
        done && \
    mkdir -p /protobuf/google/rpc && \
        for f in code error_details status http; do \
            curl -L -o /protobuf/google/rpc/${f}.proto https://raw.githubusercontent.com/istio/gogo-genproto/master/googleapis/google/rpc/${f}.proto; \
            # cue generator requires go_package to be specified in the proto files.
            sed -i '/option java_multiple_files = true;/ioption go_package = "google.golang.org/genproto/googleapis/rpc/'${f}';'${f}'";' /protobuf/google/rpc/${f}.proto; \
        done && \
    mkdir -p /protobuf/gogoproto && \
        curl -L -o /protobuf/gogoproto/gogo.proto https://raw.githubusercontent.com/gogo/protobuf/master/gogoproto/gogo.proto && \
   apk del curl

ENTRYPOINT ["/usr/bin/cue", "--paths=/protobuf,istio/api"]