FROM golang:1.24-alpine AS builder

# Install git (required for go install)
RUN apk add --no-cache git

# Install goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM alpine:latest

# Install ca-certificates for SSL connections
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy goose binary from builder
COPY --from=builder /go/bin/goose /usr/local/bin/goose

# Copy migration files
COPY db/migrations ./migrations

# Run migrations
CMD ["goose", "-dir", "./migrations", "postgres", "${GOOSE_DBSTRING}", "up"]
