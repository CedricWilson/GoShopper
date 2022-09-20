# Build Stage
FROM golang:1.19 AS builder

WORKDIR /app
COPY . .
RUN go build -o main main.go

# Execute Stage
FROM alpine:3.10
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["/app/main"]