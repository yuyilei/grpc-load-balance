FROM golang:latest

WORKDIR $GOPATH/BrokerDemo/node

ADD . $GOPATH/BrokerDemo/node

RUN mkdir -p $GOPATH/src/BrokerDemo/node/proto

COPY proto/* $GOPATH/src/BrokerDemo/node/proto/

RUN mkdir -p $GOPATH/src/BrokerDemo/node/server

COPY server/* $GOPATH/src/BrokerDemo/node/server

RUN go get google.golang.org/grpc

EXPOSE 8080

CMD ["cd", "node"]

CMD ["go", "run", "main.go"]
