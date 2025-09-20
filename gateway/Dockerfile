# syntax=docker/dockerfile:1

FROM golang:1.25.1-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates && update-ca-certificates

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY . .

WORKDIR /app/gateway

RUN sed -i 's|replace mini_messenger/auth => ../auth|replace mini_messenger/auth => ../auth|' go.mod

RUN go mod download

RUN go build -o /out/gateway-server ./cmd/server

FROM alpine:3.22

WORKDIR /
RUN apk add --no-cache ca-certificates \
    && update-ca-certificates \
    && adduser -D -u 10001 app

COPY --from=builder /out/gateway-server /gateway-server

USER app

ENTRYPOINT ["/gateway-server"]
