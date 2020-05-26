package main

import (
	"context"
	"io"
	"log"
	"net"

	echo "github.com/flutterWang/learningRPC/basic/proto/echo"
	"google.golang.org/grpc"
)

// EchoServer -
type EchoServer struct{}

// NewEchoServere -
func NewEchoServere() *EchoServer {
	return &EchoServer{}
}

// UnaryEcho -
func (s *EchoServer) UnaryEcho(ctx context.Context, r *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{
		Message: "echo",
	}, nil
}

// ClientStreamingEcho -
func (s *EchoServer) ClientStreamingEcho(stream echo.Echo_ClientStreamingEchoServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&echo.EchoResponse{Message: "data from server"})
		}
		if err != nil {
			return err
		}

		log.Printf("stream.Recv Data: %s", r.Message)
	}
}

// ServerStreamingEcho -
func (s *EchoServer) ServerStreamingEcho(r *echo.EchoRequest, stream echo.Echo_ServerStreamingEchoServer) error {
	for n := 0; n <= 6; n++ {
		err := stream.Send(&echo.EchoResponse{
			Message: "data from server",
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// BidirectionalStreamingEcho -
func (s *EchoServer) BidirectionalStreamingEcho(stream echo.Echo_BidirectionalStreamingEchoServer) error {
	n := 0
	for {
		err := stream.Send(&echo.EchoResponse{
			Message: "data from server",
		})
		if err != nil {
			return err
		}

		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		n++

		log.Printf("stream.Recv data: %s", r.Message)
	}
}

const (
	port = ":50051"
)

func main() {
	s := grpc.NewServer()

	echo.RegisterEchoServer(s, NewEchoServere())

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
