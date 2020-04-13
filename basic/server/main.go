package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"

	stream "github.com/flutterWang/learningRPC/basic/proto/stream"
	"github.com/flutterWang/learningRPC/basic/service"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var (
	db    *sql.DB
	dburl = "root:111111@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local&allowAllFiles=true"
)

func main() {
	// db = mysql.InitDB(dburl)
	s := grpc.NewServer()

	stream.RegisterStreamServer(s, service.NewStreamServer())

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
