FROM golang:1.24-bullseye AS builder

RUN apt-get update && apt-get install -y \
    gcc \
    libc6-dev \
    librdkafka-dev \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY backend ./backend
WORKDIR /app/backend

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o /app/backend/bin/backend ./cmd/api

FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y \
    librdkafka1 \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/backend/bin/backend .

EXPOSE 9001
CMD ["./backend"]
