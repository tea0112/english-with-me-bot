FROM golang:1.24.1-alpine AS dependencies
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM golang:1.24.1-alpine AS builder
WORKDIR /app
COPY --from=dependencies /go/pkg/mod /go/pkg/mod
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-w -s -extldflags '-static'" -trimpath -o english-with-me-bot .

FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/english-with-me-bot /app/
COPY .env.prod .env.dev ./
ENTRYPOINT ["/app/english-with-me-bot"]