package main

import (
	broker "BrokerDemo/broker/server"
)

func main() {
	broker.InitGreeterNode("50051")
}
