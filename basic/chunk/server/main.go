package main

import (
	"crypto/rand"
	"log"
	"net"

	pb "github.com/flutterWang/learningRPC/basic/proto/chunk"
	"google.golang.org/grpc"
)

const chunkSize = 64 * 1024

type chunkServe []byte

func (c chunkServe) Chunker(_ *pb.Empty, srv pb.Chunker_ChunkerServer) error {
	chunk := &pb.Chunk{}
	n := len(c)
	for cur := 0; cur < n; cur += chunkSize {
		if cur+chunkSize > n {
			chunk.Chunk = c[cur:n]
		} else {
			chunk.Chunk = c[cur : cur+chunkSize]
		}
		if err := srv.Send(chunk); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()

	block := make([]byte, 16*1024*1024)
	rand.Read(block)

	pb.RegisterChunkerServer(s, chunkServe(block))

	log.Println("serving on localhost:8888")
	log.Fatal(s.Serve(listen))
}
