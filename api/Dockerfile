FROM golang:1.20-alpine AS builder

WORKDIR /usr/app

COPY . /usr/app

RUN \
    CGO_ENABLED=0 \
    go mod download || true && \
    go build -o ./serve -ldflags "-s -w" ./serve.go

FROM alpine:3.17.2

WORKDIR /usr/app
COPY --from=builder /usr/app/serve /usr/app

ENTRYPOINT ["/usr/app/serve"]
