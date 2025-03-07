# Use Go 1.24.1 as the builder
FROM golang:1.24.1 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Ensure the Go binary is statically linked (important for Alpine)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o orderpack cmd/main.go

# Use a smaller base image for the final container
FROM alpine:latest

# Set the working directory in the new container
WORKDIR /root/

# Install required dependencies for running Go binaries in Alpine
RUN apk add --no-cache ca-certificates

# Copy the compiled binary from the builder
COPY --from=builder /app/orderpack .

# Make sure the binary is executable
RUN chmod +x orderpack

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./orderpack"]
