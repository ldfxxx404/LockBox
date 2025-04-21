FROM golang:1.24 AS builder

WORKDIR /app

COPY . ./

RUN go mod tidy

RUN go mod download

RUN go build -o backend ./cmd/main.go

EXPOSE 5000
