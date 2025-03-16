# Use a minimal Go image
FROM golang:alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Copy go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy source files
COPY . .

# Build the binary
RUN go build -o codeforces-rss ./cmd/codeforces-rss/main.go

# Use a small base image for the final container
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/codeforces-rss .

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./codeforces-rss"]

