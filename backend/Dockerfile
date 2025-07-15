FROM golang:1.24.1-alpine AS builder

WORKDIR /app

RUN apk add --no-cache \
    gcc \
    bash \
    musl-dev \
    libc-dev \
    make \
    git

ENV CGO_ENABLED=1 \
    GO111MODULE=on \
    PATH="/go/bin:${PATH}"

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN go build -o backend ./cmd/lock_box/main.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates bash

WORKDIR /app

COPY --from=builder /app/backend .
COPY --from=builder /go/bin/ /usr/local/bin/
COPY ./migrations /app/migrations

EXPOSE 5000
