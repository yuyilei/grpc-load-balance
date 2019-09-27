package main

import (
	node "BrokerDemo/node/server"
	"fmt"
	"time"
)

var  (
	brokerPort = "8080"
)

func main()  {
	node.InitGreeterNode(1, "A", brokerPort)

	/*
	registerNode(1, "127.0.0.1", "50055")
	registerNode(2, "127.0.0.1", "50056")
	registerNode(3, "127.0.0.1", "50057")
	*/

	for {
		time.Sleep(60 * time.Second)
		fmt.Printf("ten nodes listening...\n")
	}
}
/*
func registerNode(id int32, host string, port string) {
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v",host, brokerPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("register get connection failed, error= %v", err)
		return
	}
	if conn == nil {
		return
	}
	defer conn.Close()
	client := pb.NewBrokerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	reply, err := client.RegisterNode(ctx, &pb.RegisterNodeRequest{Host: host, NodeId: id, Port: port})
	if err != nil {
		log.Fatalf("register failed, nodeId= %v, error= %v", id, err)
	} else {
		log.Printf("register return message= %v", reply.GetMessage())
	}
}
*/
