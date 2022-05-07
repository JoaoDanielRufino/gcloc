FROM alpine:3.15.0
COPY gcloc /usr/bin
ENTRYPOINT ["/usr/bin/gcloc"]
