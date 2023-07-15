FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /go/bin/newfederationservice .

WORKDIR /app
COPY --from=builder /go/bin/newfederationservice /app/server

EXPOSE 3000
ENTRYPOINT ["/app/server"]

# CMD ["./app"]
