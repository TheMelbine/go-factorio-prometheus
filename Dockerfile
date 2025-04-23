# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/go-factorio-prometheus

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/go-factorio-prometheus .

# Expose port if your application needs it
# EXPOSE 8080

# Use ENTRYPOINT to set the binary as the entry point
ENTRYPOINT ["./go-factorio-prometheus"]

# Default command
CMD ["server"] 