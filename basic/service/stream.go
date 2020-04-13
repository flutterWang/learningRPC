package service

import (
	"io"
	"log"

	pb "github.com/flutterWang/learningRPC/basic/proto/stream"
)

// StreamServer -
type StreamServer struct{}

// NewStreamServer -
func NewStreamServer() *StreamServer {
	return &StreamServer{}
}

// PutStream -
func (s *StreamServer) PutStream(stream pb.Stream_PutStreamServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.StreamResp{Data: "data from server"})
		}
		if err != nil {
			return err
		}

		log.Printf("stream.Recv Data: %s", r.Data)
	}

	return nil
}

// GetStream -
func (s *StreamServer) GetStream(r *pb.StreamReq, stream pb.Stream_GetStreamServer) error {
	for n := 0; n <= 6; n++ {
		err := stream.Send(&pb.StreamResp{
			Data: "data from server",
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// AllStream -
func (s *StreamServer) AllStream(stream pb.Stream_AllStreamServer) error {
	n := 0
	for {
		err := stream.Send(&pb.StreamResp{
			Data: "data from server",
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

		log.Printf("stream.Recv data: %s", r.Data)
	}

	return nil
}
