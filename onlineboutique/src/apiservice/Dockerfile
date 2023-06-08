# Start from the official Golang image with a specific version
FROM golang:1.20.4

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 9090
EXPOSE 9090

# Set the entry point for the container
CMD ["./main"]
