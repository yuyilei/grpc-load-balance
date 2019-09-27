FROM golang:latest

WORKDIR $GOPATH/src/github.com/yuyilei/BrokerDemo

ADD . $GOPATH/src/github.com/yuyilei/BrokerDemo

CMD ["cd", "client"]

CMD ["go", "run", "main.go"]