FROM golang:latest

WORKDIR $GOPATH/BrokerDemo

ADD . $GOPATH/BrokerDemo

RUN go get google.golang.org/grpc 

# CMD ["ls", "$GOPATH/src/google.golang.org"]  

# RUN pwd

CMD ["cd", "client"]

CMD ["go", "run", "main.go"]
