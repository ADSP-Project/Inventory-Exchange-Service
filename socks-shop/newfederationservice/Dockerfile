FROM golang:1.20.4-alpine
WORKDIR /src

# restore dependencies
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /go/bin/federationservice .

WORKDIR /src
COPY --from=builder /go/bin/federationservice /src/server

EXPOSE 3000
ENTRYPOINT ["/src/server"]
