package main

import (
	"context"
	"io"
	"log"

	echo "github.com/flutterWang/learningRPC/basic/proto/echo"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func printClient(conn *grpc.ClientConn) error {
	client := echo.NewEchoClient(conn)
	stream, _ := client.ClientStreamingEcho(context.Background())

	for n := 0; n < 6; n++ {
		err := stream.Send(&echo.EchoRequest{
			Message: "data from client",
		})
		if err != nil {
			return err
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Printf("resp: Data: %s", resp.Message)

	return nil
}

func printServer(conn *grpc.ClientConn) error {
	client := echo.NewEchoClient(conn)
	stream, err := client.ServerStreamingEcho(context.Background(), &echo.EchoRequest{
		Message: "data from client",
	})

	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: data: %s", resp.Message)
	}

	return nil
}

func printAll(conn *grpc.ClientConn) error {
	client := echo.NewEchoClient(conn)
	stream, err := client.BidirectionalStreamingEcho(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n <= 6; n++ {
		err = stream.Send(&echo.EchoRequest{
			Message: "data from client",
		})
		if err != nil {
			return err
		}

		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: data: %s", resp.Message)
	}

	stream.CloseSend()
	return nil
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	printAll(conn)
}
