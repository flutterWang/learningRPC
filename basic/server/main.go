package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	csv "github.com/flutterWang/learningRPC/basic/proto/csv"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	csv.UnimplementedSqlServiceServer
}

var (
	db        *sql.DB
	dburl     = "root:111111@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local&allowAllFiles=true"
	sqlString = `LOAD DATA INFILE "%v" INTO TABLE %v FIELDS TERMINATED BY ',' ENCLOSED BY '"' LINES TERMINATED BY '\n'  IGNORE 1 LINES`
)

func init() {
	var err error
	db, err = sql.Open("mysql", dburl)
	if err != nil {
		panic(err)
	}
}

func (s *server) ImportData(ctx context.Context, in *csv.ImportDataRequest) (*csv.ImportDataReply, error) {
	fileName := in.GetFileName()
	tableName := in.GetTableName()
	_, err := db.Exec(sqlString)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		return &csv.ImportDataReply{Message: "init mysql fail"}, nil
	}

	return &csv.ImportDataReply{Message: "init mysql succeed"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	csv.RegisterSqlServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
