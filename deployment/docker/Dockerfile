# Multi-stage build with minimal Rocky Linux
FROM rockylinux:9 AS builder

# Install build dependencies
RUN dnf update -y && \
    dnf groupinstall -y "Development Tools" && \
    dnf install -y wget git libpcap && \
    dnf clean all

# Install Go
ENV GO_VERSION=1.23.2
RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /build

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code
COPY ../../../. .

# Build with optimizations
RUN go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o app ./cmd/honeypotter/.

# Final stage - minimal runtime
FROM rockylinux:9-minimal

# Install minimal runtime dependencies
RUN microdnf update -y && \
    microdnf install -y ca-certificates tzdata && \
    microdnf clean all

# Create app user
RUN adduser --system --no-create-home --shell /sbin/nologin --uid 1001 appuser

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /build/app .

# Set ownership and permissions
RUN chown appuser:appuser /app/app && \
    chmod +x /app/app

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8080

# Add labels for better maintainability
LABEL maintainer=""
LABEL version="0.1"
LABEL description="Honeypotter app in Docker"

# Run the application
ENTRYPOINT ["./app"]