FROM golang:1.17.7-alpine as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download -x

COPY . .

RUN go build -o bin/gcloc cmd/main.go

FROM alpine:3.15.0

WORKDIR /app

COPY --from=build /app/bin/gcloc ./

ENTRYPOINT [ "/app/gcloc" ]
