# Backend Dockerfile - Multi-stage build for Go application
FROM golang:1.24.7-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make bash curl wget

# Install buf tool for proto code generation
RUN curl -fsSL https://github.com/bufbuild/buf/releases/download/v1.45.0/buf-Linux-x86_64 -o /usr/local/bin/buf && \
    chmod +x /usr/local/bin/buf

# Set working directory to match project structure
WORKDIR /build

# Copy root go.work files from build context (build context is project root)
COPY go.work go.work.sum ./
COPY go.mod go.sum ./

# Copy parent project dependencies (backend imports from these)
COPY pkg/ ./pkg/
COPY internal/ ./internal/
COPY apis/ ./apis/
COPY cmd/ ./cmd/

# Copy all workspace modules referenced in go.work
COPY app/backend/ ./app/backend/
COPY buf/lint/optional-linter/ ./buf/lint/optional-linter/

# Set working directory to backend
WORKDIR /build/app/backend

# Download dependencies (now parent go.mod is available)
RUN go mod download

# Generate proto code
RUN make generate

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server ./cmd/server

# Final stage - minimal runtime image
FROM alpine:3.19

# Install ca-certificates for HTTPS calls, curl and wget for downloading Pulumi, procps for healthcheck, git for cloning Pulumi modules
RUN apk --no-cache add ca-certificates tzdata curl wget procps git

# Copy Go 1.24.7 from builder stage (Alpine's Go package is too old)
COPY --from=builder /usr/local/go /usr/local/go

# Install Pulumi CLI
ARG PULUMI_VERSION=v3.206.0
RUN wget -q "https://get.pulumi.com/releases/sdk/pulumi-${PULUMI_VERSION}-linux-x64.tar.gz" && \
    tar -xf "pulumi-${PULUMI_VERSION}-linux-x64.tar.gz" && \
    mv pulumi/* /usr/local/bin/ && \
    rm -rf pulumi "pulumi-${PULUMI_VERSION}-linux-x64.tar.gz" && \
    chmod +x /usr/local/bin/pulumi && \
    pulumi version

# Note: Pulumi plugins are NOT pre-installed to save Docker disk space
# Plugins will be automatically downloaded on first use and cached in the go-cache volume

# Create non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /build/app/backend/bin/server .

# Copy entrypoint script (from build context root)
COPY app/backend/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Create Pulumi home directory and Go workspace with cache directories
RUN mkdir -p /home/appuser/.pulumi && \
    mkdir -p /home/appuser/.pulumi/state && \
    mkdir -p /home/appuser/go && \
    mkdir -p /home/appuser/go/cache && \
    mkdir -p /home/appuser/go/tmp

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app && \
    chown -R appuser:appgroup /home/appuser/.pulumi && \
    chown -R appuser:appgroup /home/appuser/go

# Switch to non-root user
USER appuser

# Set Go environment (Go 1.24.7 from builder stage)
ENV GOPATH=/home/appuser/go
ENV GOCACHE=/home/appuser/go/cache
ENV GOTMPDIR=/home/appuser/go/tmp
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

# Set Pulumi environment
ENV PULUMI_HOME=/home/appuser/.pulumi
ENV PULUMI_SKIP_UPDATE_CHECK=true
ENV PULUMI_AUTOMATION_API_SKIP_VERSION_CHECK=true

# Expose port
EXPOSE 50051

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:50051/health || exit 1

# Use entrypoint script
ENTRYPOINT ["/entrypoint.sh"]
CMD ["./server"]
