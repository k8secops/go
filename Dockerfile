FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o products-service .

FROM alpine:3.19
RUN adduser -D appuser
USER appuser
WORKDIR /app
COPY --from=builder /app/products-service .
EXPOSE 8081
CMD ["./products-service"]
