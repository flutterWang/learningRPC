package main

import (
	"log"

	"github.com/flutterWang/learningRPC/basic/client/methods/test"
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
	// c := csv.NewCsvClient(conn)

	test.Test(conn)
}
