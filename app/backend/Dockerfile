# Backend Dockerfile - Multi-stage build for Go application
FROM golang:1.24.7-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make curl

# Install buf tool for proto code generation
RUN curl -fsSL https://github.com/bufbuild/buf/releases/download/v1.45.0/buf-Linux-x86_64 -o /usr/local/bin/buf && \
    chmod +x /usr/local/bin/buf

# Set working directory
WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN make generate && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server ./cmd/server

# Final stage - minimal runtime image
FROM alpine:3.19

# Install ca-certificates for HTTPS calls and procps for healthcheck
RUN apk --no-cache add ca-certificates tzdata procps

# Create non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/bin/server .

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 50051

# Health check - gRPC services don't have HTTP endpoints, so we check if process is running

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD pgrep -f server || exit 1

# Run the application
CMD ["./server"]
