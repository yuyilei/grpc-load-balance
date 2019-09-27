package main

import (
	 pb "BrokerDemo/client/proto"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"context"
	"sync"
	"time"
	"os"
)

var (
	brokerPort = os.Getenv("ECHO_SERVICE_PORT")
	brokerHost = os.Getenv("ECHO_SERVICE_HOST")
)

type counter struct {
	mutex sync.Mutex
	count map[string]int32
}

func (c counter) Add(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if _, exist := c.count[key]; !exist {
		c.count[key] = 0
	}
	c.count[key]++
}

func main() {
	SayHello(10, 10)
}

func currentTime() string {
	t := time.Now()
	fmtTime := fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
	return fmtTime
}

func toSecond(t string) int {
	var hh, mm, ss int
	fmt.Sscanf(t, "%02d:%02d:%02d", &hh, &mm, &ss)
	return hh*3600 + mm*60 + ss
}

func costTime(s string, e string) string {
	res := fmt.Sprintf("%ds", toSecond(e)-toSecond(s))
	return res
}

func SayHello(n int, id int32) {

	clientConn, err := grpc.Dial(fmt.Sprintf("%v:%v", brokerHost, brokerPort), grpc.WithInsecure())
	helloClient := pb.NewBrokerClient(clientConn)
	stream, err := helloClient.SayHello(context.Background())
	if err != nil {
		log.Fatalf("cleint could not sayHello: %v", err)
		return
	}
	startTime := currentTime()
	c := new(counter)
	c.count = make(map[string]int32)
	var wg sync.WaitGroup
	wg.Add(n)
	go func() {
		for i := 0; i < n; i++ {
			reply, err := stream.Recv()
			if err != nil {
				log.Fatalf("get hello failed, index= %v, error= %v", i, err)
			}
			log.Printf("get Hello reply, message= {%v}", reply.GetMessage())
			// replyFromBroker <- reply
			c.Add(reply.GetNodeName())
			wg.Done()
		}
	}()
	for i := 0; i < n; i++ {
		if err := stream.Send(&pb.HelloRequest{ClientId: id}); err != nil {
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("counld not send hello index= %v", i)
			}
		}
	}

	wg.Wait()
	endTime := currentTime()
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("client close stream failed, error= %v", err)
	}
	log.Printf("StartTime= %v, EndTime= %v CostTime= %v", startTime, endTime, costTime(startTime, endTime))
	log.Printf("NodeCount: %v\n", c.count)
}


