package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	pb "BrokerDemo/proto"
)

type nodeConn struct {
	host string
	port string
	nodeId int32
}

type nodeList struct {
	mutex   sync.Mutex
	nodes   map[int32]*nodeConn
}

type broker struct {
	allNodes nodeList
	server *grpc.Server
}


func (n *nodeConn) connectToNode() (*pb.GreeterClient, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", n.host, n.port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect to node, nodeId= %v, error= %v", n.nodeId, err)
		return nil, err
	}
	client := pb.NewGreeterClient(conn)
	return &client, nil
}

func (n* nodeList) existNode(id int32) bool{
	n.mutex.Lock()
	defer n.mutex.Unlock()
	if _, exist := n.nodes[id]; exist {
		return true
	}
	return false
}

func (n* nodeList) getRandomNode() (*pb.GreeterClient, error) {
	n.mutex.Lock()
	index :=  rand.Intn(len(n.nodes))
	var node *nodeConn
	for value := range n.nodes {
		if index == 0 {
			node = n.nodes[value]
			break
		}
		index--
	}
	n.mutex.Unlock()
	if node == nil {
		return nil, nil
	}
	return node.connectToNode()
}

func (n* nodeList) addNode(id int32, host string, port string) bool {
	newNode := &nodeConn{host:host, port:port, nodeId: id}
	if n.existNode(id) {
		return false
	}
	n.mutex.Lock()
	n.nodes[id] = newNode
	n.mutex.Unlock()
	return true
}

func (b *broker) RegisterNode(ctx context.Context, req *pb.RegisterNodeRequest) (*pb.RegisterNodeReply, error) {
	if b.allNodes.existNode(req.GetNodeId()) {
		return &pb.RegisterNodeReply{Message: "node already registered!"}, nil
	}
	if b.allNodes.addNode(req.GetNodeId(), req.GetHost(), req.GetPort()) {
		log.Printf("node registered successfully, nodeId= %v", req.GetNodeId())
	}
	return &pb.RegisterNodeReply{Message: "node register success"}, nil
}

func (b *broker)  SayHello(srv pb.Broker_SayHelloServer) error {

	finished := make(chan bool)
	replyFromNode := make(chan *pb.HelloReply)
	go func() {
		for {
			select {
			case res := <-finished:
				if res {
					return
				}
			case reply := <-replyFromNode:
				go func() {
					err := srv.Send(reply)
					if err != nil {
						log.Fatalf("return helloReply to client failed!, error= %v", err)
					}
				}()
			}
		}
	}()

	defer func() {
		finished <- true
	}()

	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			// log.Fatalf("Broker get request from stream failed, error= %v", err)
			return err
		}
		client, err := b.allNodes.getRandomNode()
		if err != nil {
			log.Fatalf("Broker connect to Greeter failed, error= %v", err)
			return err
		}

		go func(c pb.GreeterClient, id int32) {
			ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
			// defer cancel()
			reply, err := c.SayHello(ctx, &pb.HelloRequest{ClientId: id})
			if err != nil {
				log.Fatalf("send hello failed， error= %v", err)
				return
			}
			replyFromNode <- reply
		}(*client, req.GetClientId())
	}
	return nil
}

func (b *broker) Start(port string) error {
	serverListen, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatalf("broker try to listen %v failed", port)
		return err
	}
	b.server = grpc.NewServer()
	pb.RegisterBrokerServer(b.server, b)
	// go func() {
		if err := b.server.Serve(serverListen); err != nil {
			log.Fatalf("broker failed to serve, port= %v, error= %v", port, err)
		} else {
			log.Printf("start broker success, port= %v", port)
		}
	// }()
	return nil
}

func newBroker() *broker {
	b := new(broker)
	b.allNodes.nodes = make(map[int32]*nodeConn)
	return b
}

func InitGreeterNode(port string)  {
	s := newBroker()
	err := s.Start(port)
	if err != nil {
		log.Fatalf("start broker failed, port= %v, error= %v", port, err)
	} else {
		log.Printf("start broker success, port= %v", port)
	}
}

