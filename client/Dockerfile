FROM golang:latest

WORKDIR $GOPATH/BrokerDemo

ADD . $GOPATH/BrokerDemo

CMD ["go", "get", "google.golang.org/grpc"]

CMD ["cd", "client"]

CMD ["go", "run", "main.go"]