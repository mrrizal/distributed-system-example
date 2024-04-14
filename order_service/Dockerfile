# Dockerfile for order-service
FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o order-service .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/order-service .
COPY --from=builder /app/wait-for-rabbitmq.sh .

CMD ["./order-service"]
