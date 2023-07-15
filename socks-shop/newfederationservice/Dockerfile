FROM golang:1.20.4-alpine as builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /go/bin/newfederationservice .

WORKDIR /src
COPY --from=builder /go/bin/newfederationservice /src/server

EXPOSE 3000
ENTRYPOINT ["/src/server"]

# CMD ["./app"]
