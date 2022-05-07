FROM alpine:3.15.0
COPY gcloc /
ENTRYPOINT ["/gcloc"]
