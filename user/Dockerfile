# syntax=docker/dockerfile:1

FROM golang:1.25.1-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates && update-ca-certificates

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /out/user-server ./cmd/server

FROM alpine:3.22

WORKDIR /
RUN apk add --no-cache ca-certificates \
    && update-ca-certificates \
    && adduser -D -u 10001 app

COPY --from=builder /out/user-server /user-server

USER app

ENTRYPOINT ["/user-server"]
