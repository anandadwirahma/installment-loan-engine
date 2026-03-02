# Stage 1: Build the application
FROM golang:1.25.6-alpine AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory
WORKDIR /app

# Copy the source code
COPY . .

# Download dependencies
RUN go mod download

# Build the application
RUN go build -o installment-loan-engine .

# Stage 2: Create the production image
FROM alpine:latest

RUN apk add --no-cache tzdata

# Set working directory
WORKDIR /app

# Create user and group for application
#
# Create a group with GID 1001
RUN addgroup -g 1001 binarygroup
# Create a user with UID 1001 and assign them to the 'binarygroup' group
RUN adduser -D -u 1001 -G binarygroup userapp

# Copy the binary from the builder stage
COPY --from=builder --chown=userapp:binarygroup /app/installment-loan-engine .

# Switch to the userapp user
USER userapp

# Expose port 8082 (matching APP_PORT in .env)
EXPOSE 8082

# Command to run the application
CMD ["./installment-loan-engine"]

