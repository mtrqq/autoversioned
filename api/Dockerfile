FROM golang:1.20-alpine AS builder

WORKDIR /usr/app

COPY . /usr/app

RUN \
    --mount=type=cache,id=gomod,target=/go/pkg/mod \
    --mount=type=cache,id=gobuild,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -o ./serve -ldflags "-s -w" ./serve.go

FROM alpine:3.17.2

WORKDIR /usr/app
COPY --from=builder /usr/app/serve /usr/app

ENTRYPOINT ["./serve"]
