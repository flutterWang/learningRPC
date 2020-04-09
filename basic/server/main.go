package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "flutterWang/learningRPC/basicUse/pbMysql"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedSqlServiceServer
}

var (
	db    *sql.DB
	dburl = "root:111111@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local&allowAllFiles=true"
)

func init() {
	var err error
	db, err = sql.Open("mysql", dburl)
	if err != nil {
		panic(err)
	}
}

func (s *server) ImportData(ctx context.Context, in *pb.ImportDataRequest) (*pb.ImportDataReply, error) {
	sqlString := in.getUrl()
	_, err := db.Exec(sqlString)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		return &pb.ImportDataReply{Message: "init mysql fail"}, nil
	}

	return &pb.ImportDataReply{Message: "init mysql succeed"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSqlServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
