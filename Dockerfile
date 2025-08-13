# Step 1: Build the Go app
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app

# Step 2: Run the app in a small image
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/app .

# You can set default flags or pass via Kubernetes config
CMD ["./app", "--producers=5", "--consumers=10", "--rps=50", "--duration=10"]
