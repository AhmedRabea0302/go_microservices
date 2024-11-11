# Base go image
FROM golang:1.23-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

# Build the binary
RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

RUN chmod +x /app/brokerApp


# Build a tiny docker image
FROM alpine:latest

RUN mkdir /app/

COPY --from=builder /app/brokerApp /app

CMD ["/app/brokerApp"]