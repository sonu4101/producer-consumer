# -------------------------
# Step 1: Build the Go app
# -------------------------
FROM golang:1.23 AS builder

WORKDIR /app

# Copy and download dependencies first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/app

# -------------------------
# Step 2: Create minimal runtime image
# -------------------------
FROM alpine:latest

WORKDIR /root/

# Copy compiled binary from builder
COPY --from=builder /app/app .

# Default command
CMD ["/root/app"]
