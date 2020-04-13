package stream

import (
	"context"
	"io"
	"log"

	pb "github.com/flutterWang/learningRPC/basic/proto/stream"
	"google.golang.org/grpc"
)

// PrintClient -
func PrintClient(conn *grpc.ClientConn) error {
	client := pb.NewStreamClient(conn)
	stream, _ := client.PutStream(context.Background())

	for n := 0; n < 6; n++ {
		err := stream.Send(&pb.StreamReq{
			Data: "data from client",
		})
		if err != nil {
			return err
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Printf("resp: Data: %s", resp.Data)

	return nil
}

// PrintServer -
func PrintServer(conn *grpc.ClientConn) error {
	client := pb.NewStreamClient(conn)
	stream, err := client.GetStream(context.Background(), &pb.StreamReq{
		Data: "data from client",
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

		log.Printf("resp: data: %s", resp.Data)
	}

	return nil
}

// PrintAll -
func PrintAll(conn *grpc.ClientConn) error {
	client := pb.NewStreamClient(conn)
	stream, err := client.AllStream(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n <= 6; n++ {
		err = stream.Send(&pb.StreamReq{
			Data: "data from client",
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

		log.Printf("resp: data: %s", resp.Data)
	}

	stream.CloseSend()
	return nil
}
