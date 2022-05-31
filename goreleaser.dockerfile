FROM alpine:3.15.0
COPY gcloc /usr/bin
RUN apk add --no-cache git
ENTRYPOINT ["/usr/bin/gcloc"]
