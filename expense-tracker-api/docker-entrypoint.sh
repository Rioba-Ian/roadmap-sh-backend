#!/bin/sh
# docker-entrypoint.sh

# Wait for the database to be ready (optional, but recommended in a compose environment)
# A simple sleep might work, or a more robust solution like wait-for-it.sh
sleep 5

# Run the goose migrations
echo "Running database migrations with Goose..."
export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable
goose -version
echo "Using connection string: ${GOOSE_DBSTRING}"
goose status
goose up
echo "Migrations applied."

# Run the main application command
exec "$@"
