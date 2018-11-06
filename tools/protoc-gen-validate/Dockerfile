FROM alpine:3.7 as builder
RUN apk update
RUN apk add --no-cache go>1.10 git build-base
ENV GOPATH=/go \
    PATH=/go/bin/:$PATH \
    OUTDIR=/out
RUN mkdir -p ${OUTDIR}/usr/bin/
RUN go get -u -v -ldflags '-w -s' \
    github.com/lyft/protoc-gen-validate \
    && install ${GOPATH}/bin/protoc-gen-validate ${OUTDIR}/usr/bin/

FROM znly/upx as packer
COPY --from=builder /out/ /out/
RUN upx --lzma \
        /out/usr/bin/protoc-gen-validate

FROM alpine:3.7
COPY --from=packer /out/ /
ENTRYPOINT ["/usr/bin/protoc-gen-validate"]
