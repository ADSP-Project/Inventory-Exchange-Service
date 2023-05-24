FROM golang:1.20.4-alpine@sha256:4ee203ff3933e7a6f18d3574fd6661a73b58c60f028d2927274400f4774aaa41 as builder
RUN apk add --no-cache ca-certificates git
RUN apk add build-base
WORKDIR /src

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

WORKDIR /src
COPY --from=builder /go/bin/apiservice /src/apiservice

# Definition of this variable is used by 'skaffold debug' to identify a golang binary.
# Default behavior - a failure prints a stack trace for the current goroutine.
# See https://golang.org/pkg/runtime/
ENV GOTRACEBACK=single

# Expose the port on which the application will run
EXPOSE 8000
ENTRYPOINT ["/src/apiserver"]