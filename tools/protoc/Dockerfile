FROM golang:1.12 as golang-build-env

# Build gRPC from source
RUN apt-get update
RUN apt-get install -y build-essential unzip autoconf libtool zlibc zlib1g-dev libc-ares-dev libssl-dev

# Install protoc
RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.9.0/protobuf-all-3.9.0.tar.gz
run tar -xzvf protobuf-all-3.9.0.tar.gz
WORKDIR protobuf-3.9.0
RUN aclocal
RUN automake
RUN ./autogen.sh
RUN ./configure --prefix=/usr
RUN make
RUN make install

# Install gRPC
RUN curl -LO https://github.com/grpc/grpc/archive/v1.22.0.tar.gz
RUN tar -xzvf v1.22.0.tar.gz
WORKDIR grpc-1.22.0
RUN make install

# Curl mysterious protobuf files
RUN mkdir -p /protobuf/google/protobuf
RUN mkdir -p /protobuf/google/rpc
RUN mkdir -p /protobuf/gogoproto
RUN for f in any duration descriptor empty struct timestamp wrappers; do \
            curl -L -o /protobuf/google/protobuf/${f}.proto https://raw.githubusercontent.com/google/protobuf/master/src/google/protobuf/${f}.proto; \
        done
RUN for f in code error_details status http; do \
            curl -L -o /protobuf/google/rpc/${f}.proto https://raw.githubusercontent.com/istio/gogo-genproto/master/googleapis/google/rpc/${f}.proto; \
        done
RUN curl -L -o /protobuf/gogoproto/gogo.proto https://raw.githubusercontent.com/gogo/protobuf/master/gogoproto/gogo.proto

# Install protoc plugins
RUN go get -u -v -ldflags '-w -s' github.com/gogo/protobuf/protoc-gen-gogofast
RUN go get -u -v -ldflags '-w -s' github.com/gogo/protobuf/protoc-gen-gogofaster
RUN go get -u -v -ldflags '-w -s' github.com/gogo/protobuf/protoc-gen-gogoslick
RUN go get -u -v -ldflags '-w -s' istio.io/tools/protoc-gen-docs

# And finally, install other tools needed by the API repo
RUN go get -u -v -ldflags '-w -s' istio.io/tools/cmd/annotations_prep

ENTRYPOINT ["/usr/bin/protoc", "-I/protobuf"]
