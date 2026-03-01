#!/bin/bash

# Load environment variables from .env file
if [ -f .env ]; then
    export $(cat .env | xargs)
fi

DB_NAME=${DB_NAME:-billenginedb}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-postgres}
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}

echo "Step 1: Checking and creating database if not exists..."
# Using PGPASSWORD to avoid prompt
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -tc "SELECT 1 FROM pg_database WHERE datname = '$DB_NAME'" | grep -q 1 || \
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -c "CREATE DATABASE $DB_NAME"

if [ $? -eq 0 ]; then
    echo "Database '$DB_NAME' is ready."
else
    echo "Failed to create or check database '$DB_NAME'. Please ensure PostgreSQL is running and credentials are correct."
    exit 1
fi

echo "Step 2: Running go mod tidy..."
go mod tidy

echo "Step 3: Running the application..."
# Run the application
go run main.go

if [ $? -ne 0 ]; then
    echo "Failed to run the application."
    exit 1
fi
