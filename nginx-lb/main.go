package main

import (
	pb "BrokerDemo/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

var (
	brokerPort = "50051"
	clientNum = 800
	c *counter
	wg sync.WaitGroup
)

type counter struct {
	mutex sync.Mutex
	count map[string]int32
}

func init()  {
	c = new(counter)
	c.count = make(map[string]int32)
	wg.Add(clientNum)
}

func (c counter) Add(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if _, exist := c.count[key]; !exist {
		c.count[key] = 0
	}
	c.count[key]++
}

func currentTime() string {
	t := time.Now()
	fmtTime := fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
	return fmtTime
}
func costTime(s string, e string) string {
	res := fmt.Sprintf("%ds", toSecond(e)-toSecond(s))
	return res
}

func toSecond(t string) int {
	var hh, mm, ss int
	fmt.Sscanf(t, "%02d:%02d:%02d", &hh, &mm, &ss)
	return hh*3600 + mm*60 + ss
}


func SayHi(clientId int32) {
	clientConn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%v", brokerPort), grpc.WithInsecure())
	defer wg.Done()
	if err != nil {
		log.Printf("cleint could not sayHi: %v", err)
		return
	}
	hiClient := pb.NewGreeterClient(clientConn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := hiClient.SayHello(ctx, &pb.HelloRequest{ClientId: clientId})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
	c.Add(r.GetNodeName())
}

func main() {
	startTime := currentTime()
	var i int32
	for ; i < int32(clientNum); i++ {
		go func() {
			SayHi(i)
		}()
	}
	wg.Wait()
	endTime := currentTime()
	log.Printf("StartTime= %v, EndTime= %v CostTime= %v", startTime, endTime, costTime(startTime, endTime))
	log.Printf("NodeCount: %v\n", c.count)
}