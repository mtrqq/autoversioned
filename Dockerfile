FROM golang:1.20-alpine AS builder

WORKDIR /usr/app

COPY server /usr/app/server
COPY go.mod /usr/app/go.mod
COPY go.sum /usr/app/go.sum

RUN \
    CGO_ENABLED=0 \
    cat ./server/serve.go && \
    go mod download || true && \
    go build -o ./serve -ldflags "-s -w" ./server/serve.go

FROM alpine:3.17.2

WORKDIR /usr/app
COPY --from=builder /usr/app/serve /usr/app

ENTRYPOINT ["/usr/app/serve"]
