package main

import (
	"context"
	"log"
	"time"

	pb "github.com/flutterWang/learningRPC/basic/proto/csv"
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
	c := pb.NewCsvClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ImportData(ctx, &pb.ImportDataRequest{FileName: "./electricData.csv", TableName: "electrion"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())
}
