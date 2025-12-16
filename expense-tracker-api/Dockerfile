FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o ./bin/api main.go


# Install the goose CLI binary for SQL migrations
RUN go install github.com/pressly/goose/v3/cmd/goose@v3.26.0

FROM alpine:3.20 AS runtime
WORKDIR /app

# Copy only necessary files from builder stage
COPY --from=builder /app/bin/api ./bin/api
# Copy goose binary or custom migrator if built
COPY --from=builder /go/bin/goose /usr/local/bin/goose
# Copy database migrations
COPY db/migrations ./db/migrations

# Copy the entrypoint script
COPY docker-entrypoint.sh .
RUN apk --no-cache add wget && chmod +x docker-entrypoint.sh

ENTRYPOINT [ "./docker-entrypoint.sh" ]
CMD [ "/app/bin/api" ]
