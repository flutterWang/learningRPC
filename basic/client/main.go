package main

import (
	"log"

	"github.com/flutterWang/learningRPC/basic/client/methods/stream"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// importdata.Upload(conn, "/Users/bofeng/Documents/electricData.csv", "electrion")
	// stream.PrintClient(conn)
	// stream.PrintServer(conn)
	stream.PrintAll(conn)
}
