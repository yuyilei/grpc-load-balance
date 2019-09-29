package main

import (
	node "BrokerDemo/node/server"
	"fmt"
	"time"
)

var  (
	brokerPort = "8080"
)

func use_nginx()  {
	node.InitGreeterNode(1, "A", "50055")
	node.InitGreeterNode(2, "B", "50056")
	node.InitGreeterNode(3, "C", "50057")
	node.InitGreeterNode(4, "E", "50058")
	node.InitGreeterNode(5, "F", "50059")
	node.InitGreeterNode(6, "G", "50060")
	node.InitGreeterNode(7, "H", "50061")
	node.InitGreeterNode(8, "I", "50062")
	node.InitGreeterNode(9, "G", "50063")
	node.InitGreeterNode(10, "K", "50064")
	node.InitGreeterNode(11, "L", "50065")
	node.InitGreeterNode(12, "M", "50066")
	node.InitGreeterNode(13, "N", "50067")
	node.InitGreeterNode(14, "O", "50068")
	node.InitGreeterNode(15, "P", "50069")
	node.InitGreeterNode(16, "Q", "50070")
	node.InitGreeterNode(17, "R", "50071")
	node.InitGreeterNode(18, "S", "50072")
	node.InitGreeterNode(19, "T", "50073")
	node.InitGreeterNode(20, "U", "50074")
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

func main()  {
	use_nginx()
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
