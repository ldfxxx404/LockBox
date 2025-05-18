FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest

ENV PATH="/go/bin:${PATH}"

RUN swag init -g /cmd/lock_box/main.go --output /cmd/swagger

RUN go build -o backend ./cmd/lock_box/main.go

EXPOSE 5000
