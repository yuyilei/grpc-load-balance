FROM golang:latest

WORKDIR $GOPATH/BrokerDemo

ADD . $GOPATH/BrokerDemo

CMD ["cd", "client"]

CMD ["go", "run", "main.go"]