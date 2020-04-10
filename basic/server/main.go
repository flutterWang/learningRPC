package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"

	pb "github.com/flutterWang/learningRPC/basic/proto/csv"
	"github.com/flutterWang/learningRPC/basic/service"
	"github.com/flutterWang/learningRPC/tools/mysql"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var (
	db        *sql.DB
	dburl     = "root:111111@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local&allowAllFiles=true"
	sqlString = `LOAD DATA INFILE "%v" INTO TABLE %v FIELDS TERMINATED BY ',' ENCLOSED BY '"' LINES TERMINATED BY '\n'  IGNORE 1 LINES`
)

func main() {
	db = mysql.InitDB(dburl)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterCsvServer(s, service.NewImportDataServer(db, sqlString))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
