# Use the official Golang image to create a build artifact.
FROM golang:1.23.5-alpine AS builder

# Set the working directory inside the container.
WORKDIR /app

# Copy the Go module files.
COPY go.mod go.sum ./

# Download the dependencies.
RUN go mod download

# Copy the source code.
COPY . .

# Build the Go application.
RUN go build -o main .

# Use a minimal Alpine image for the final stage.
FROM alpine:latest

# Set the working directory.
WORKDIR /root/

# Copy the binary from the builder stage.
COPY --from=builder /app/main .
# Copy the templates directory.
COPY --from=builder /app/templates ./templates
# Copy the static directory.
COPY --from=builder /app/static ./static

# Expose the application port.
EXPOSE 3000

# Run the application.
CMD ["./main"]