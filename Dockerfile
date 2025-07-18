# Dockerfile
# Use Multi-Stage Build

FROM golang:1.24.1-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/app ./cmd/main.go

FROM alpine:3.22

# Uncomment if need bash
# RUN apk add --no-cache bash

WORKDIR /app

COPY --from=builder /app/bin/app .
COPY ./.env ./.env

EXPOSE 3000

CMD ["./app"]