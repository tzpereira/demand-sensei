FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY backend ./backend

WORKDIR /app/backend
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/backend/bin/backend ./cmd/api

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/backend/bin/backend .

RUN mkdir -p /app/data/uploads

EXPOSE 9001
CMD ["./backend"]
