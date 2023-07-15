FROM golang:1.20.4-alpine

WORKDIR /src

# Restore dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build the application
COPY . .
RUN go build -o /go/bin/newfederationservice

# Set the entrypoint and expose the port
EXPOSE 3000
ENTRYPOINT ["/go/bin/newfederationservice"]

