# TODO -> Make this multistage, only copy compiled binary
FROM golang:1.24.12-alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .
CMD ["./main"]