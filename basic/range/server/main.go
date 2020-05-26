package main

import (
	"crypto/rand"
	"log"
	"net"

	pb "github.com/flutterWang/learningRPC/basic/proto/range"
	"google.golang.org/grpc"
)

const chunkSize = 64 * 1024

type chunkerSrv []byte

func (c chunkerSrv) Range(r *pb.Res, srv pb.RangeChunker_RangeServer) error {
	chunk := &pb.Chunk{}
	ranges := c.parseRanges(r)

	for _, subr := range ranges {
		start, stop := subr[0], subr[1]
		for cur := start; cur < stop; cur += chunkSize {
			if cur+chunkSize > stop {
				chunk.Chunk = c[cur:stop]
			} else {
				chunk.Chunk = c[cur : cur+chunkSize]
			}

			if err := srv.Send(chunk); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c chunkerSrv) parseRanges(r *pb.Res) [][2]int {
	n := len(c)
	ranges := [][2]int{}
	res := r.GetR()
	if len(res) == 0 {
		return [][2]int{[2]int{0, n}}
	}

	for _, subr := range res {
		start, stop := rangeLimit(subr, n)
		if start == -1 {
			return nil
		}
		ranges = append(ranges, [2]int{start, stop})
	}

	return ranges
}

func rangeLimit(r *pb.Range, len int) (int, int) {
	start, stop := int(r.Start), int(r.Stop)+1
	if stop > len || stop == 0 {
		stop = len
	}

	if start < 0 || stop < 0 || start >= stop {
		return -1, -1
	}

	return start, stop
}

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	block := make([]byte, 12*1024*1024)
	rand.Read(block)
	pb.RegisterRangeChunkerServer(s, chunkerSrv(block))
	log.Println("serving on localhost:8888")
	log.Fatal(s.Serve(listen))
}
