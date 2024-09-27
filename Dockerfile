# Build Stage
FROM golang:1.23-alpine AS builder

# Create and set app directory
WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy other files
COPY . .

# Get build of service with static binary (smaller size)
RUN CGO_ENABLED=0 GOOS=linux go build -o pastebin-clone ./cmd

# Final Stage
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from build stage
COPY --from=builder /app/pastebin-clone .

# Expose port 8080
EXPOSE 8080

# Run the service
CMD ["./pastebin-clone"]
