FROM gcr.io/istio-testing/protolock:2018-10-23 as protolock

FROM gcr.io/istio-testing/protoc:2018-06-12

COPY --from=protolock \
        /usr/bin/protolock \
        /usr/bin/

ENTRYPOINT []
