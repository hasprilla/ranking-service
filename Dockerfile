# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
# go.sum might not exist if created manually without go tool
COPY go.sum* ./
# Manual copy of go.sum might fail if it doesn't exist yet, but Railway will handle go mod tidy
RUN go mod download || true

COPY . .

RUN go build -o main .

# Run stage
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
