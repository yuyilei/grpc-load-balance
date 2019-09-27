package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"

	pb "BrokerDemo/proto"
)

type server struct{
	nodeId int32
	name string
	mutex sync.Mutex
	greeter *grpc.Server
}

func currentTime() string {
	t := time.Now()
	fmtTime := fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
	return fmtTime
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	// s.mutex.Lock()
	// defer s.mutex.Unlock()
	startTime := currentTime()
	log.Printf("Received request from client, nodeName= %v, currentTime= %v", s.name, startTime)
	// time.Sleep(100 * time.Millisecond)
	endTime := currentTime()
	log.Printf("Finished some work, nodeId= %v, currentTime= %v", s.nodeId, endTime)
	return &pb.HelloReply{Message: fmt.Sprintf("node say Hello to client, nodeName= %v, startTime= %v, endTime= %v", s.name, startTime, endTime), NodeName: s.name}, nil
}
/*
func (s *server) SayHi(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	log.Printf("Received from cilent, clientId= %v", in.GetClientId())
	return &pb.HelloReply{Message: fmt.Sprintf("node say Hi to client, nodeNode= %v", s.name), NodeName: s.name}, nil
}
*/
func (s *server) Start(id int32, name string, port string) error {
	s.name = name
	serverListen, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatalf("server try to listen %v failed", port)
		return err
	}
	s.greeter = grpc.NewServer()
	pb.RegisterGreeterServer(s.greeter, s)
	go func() {
		if err := s.greeter.Serve(serverListen); err != nil {
			log.Fatalf("server failed to serve, nodeId= %v, port= %v, error= %v", s.nodeId,  port, err)
		}
	}()
	return nil
}

func newGreeterNode() *server {
	s := new(server)
	return s
}

func InitGreeterNode(id int32, name string, port string)  {
	s := newGreeterNode()
	s.nodeId = id
	s.name = name
	err := s.Start(id, name, port)
	if err != nil {
		log.Fatalf("start node failed, nodeId= %v, nodeName= %v, port= %v, error= %v", id, name, port, err)
	} else {
		log.Printf("start node success, nodeId= %v, nodeName= %v, port= %v", id, name, port)
	}
}
