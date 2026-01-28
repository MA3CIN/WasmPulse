FROM golang:1.24.12-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

# Networking capabilities for port exposure, post hardening cause host pid (see docker compose)
# RUN setcap 'cap_net_bind_service=+ep' /main

CMD ["./main"]