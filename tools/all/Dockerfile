FROM alpine:3.7 as prototool_download

RUN apk update
RUN apk add curl libc6-compat
RUN curl -sSL https://github.com/uber/prototool/releases/download/v1.3.0/prototool-Linux-x86_64 -o /usr/bin/prototool
RUN chmod +x /usr/bin/prototool

FROM gcr.io/istio-testing/protolock:2019-03-11 as protolock
FROM gcr.io/istio-testing/protoc:2019-07-25

COPY --from=protolock /usr/bin/protolock /usr/bin/
COPY --from=prototool_download /usr/bin/prototool /usr/bin/

ENTRYPOINT []
