package main

import (
	"context"
	"io"
	"log"

	pb "github.com/flutterWang/learningRPC/basic/proto/chunk"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewChunkerClient(conn)
	stream, err := client.Chunker(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	var block []byte
	for {
		c, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("Transfer of %d bytes successful", len(block))
				return
			}
			log.Fatal(err)
		}
		block = append(block, c.Chunk...)
	}
}
