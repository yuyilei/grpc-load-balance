FROM golang:latest

WORKDIR $GOPATH/BrokerDemo/client

ADD . $GOPATH/BrokerDemo/client

RUN mkdir -p $GOPATH/src/BrokerDemo/client/proto

COPY proto/* $GOPATH/src/BrokerDemo/client/proto/


RUN go get google.golang.org/grpc 

EXPOSE 8080

CMD ["cd", "client"]

CMD ["go", "run", "main.go"]
