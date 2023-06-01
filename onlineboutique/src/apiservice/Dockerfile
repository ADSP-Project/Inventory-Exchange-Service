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

# Skaffold passes in debug-oriented compiler flags
ARG SKAFFOLD_GO_GCFLAGS
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o /apiservice .

FROM alpine:3.17.3@sha256:124c7d2707904eea7431fffe91522a01e5a861a624ee31d03372cc1d138a3126 AS without-grpc-health-probe-bin
RUN apk add --no-cache ca-certificates

WORKDIR /src
COPY --from=builder /go/bin/apiservice /src/apiservice

# Definition of this variable is used by 'skaffold debug' to identify a golang binary.
# Default behavior - a failure prints a stack trace for the current goroutine.
# See https://golang.org/pkg/runtime/
ENV GOTRACEBACK=single

# Expose the port on which the application will run
EXPOSE 8000
ENTRYPOINT ["/src/apiserver"]

FROM without-grpc-health-probe-bin
# renovate: datasource=github-releases depName=grpc-ecosystem/grpc-health-probe
ENV GRPC_HEALTH_PROBE_VERSION=v0.4.18
RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe