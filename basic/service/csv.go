package service

import (
	"context"
	"database/sql"
	"io/ioutil"

	pb "github.com/flutterWang/learningRPC/basic/proto/csv"
)

// CsvServer -
type CsvServer struct {
	db *sql.DB
}

// NewCsvServer -
func NewCsvServer(db *sql.DB) *CsvServer {
	return &CsvServer{
		db: db,
	}
}

// Upload -
func (s *CsvServer) Upload(ctx context.Context, in *pb.UploadRequest) (*pb.UploadResponse, error) {
	file := in.GetFile()
	// name := in.GetName()
	// fileName := "mysqldata.csv"

	err := ioutil.WriteFile("../../data/mysqldata.csv", file, 0644)

	// err := mysql.ImportData(s.db, fileName, name)

	if err != nil {
		return &pb.UploadResponse{Code: 500, Message: "init mysql fail"}, nil
	}

	return &pb.UploadResponse{Code: 200, Message: "init mysql succeed"}, nil
}
