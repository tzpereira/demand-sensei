FROM golang:1.24.5 AS builder

ARG GOARCH=arm64

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=${GOARCH}
ENV CONFLUENT_KAFKA_GO_DISABLE_VENDOR=1
ENV PKG_CONFIG_PATH=/usr/lib/pkgconfig:/usr/lib/aarch64-linux-gnu/pkgconfig:/usr/lib/x86_64-linux-gnu/pkgconfig

RUN apt-get update && apt-get install -y \
    gcc \
    libc6-dev \
    librdkafka-dev \
    pkg-config \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY backend ./backend
WORKDIR /app/backend

RUN case "${GOARCH}" in arm64|amd64) ;; *) echo "unsupported GOARCH: ${GOARCH}. use arm64 or amd64." >&2; exit 1 ;; esac

RUN go mod download
RUN go build -tags dynamic -o /app/backend/bin/backend ./cmd/api

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    librdkafka1 \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/backend/bin/backend .

EXPOSE 9001
CMD ["./backend"]
