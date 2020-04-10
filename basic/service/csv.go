package service

import (
	"database/sql"
	"fmt"
	"context"

	pb "github.com/flutterWang/learningRPC/basic/proto/csv"	
)

// ImportdataServer -
type ImportdataServer struct{
	db *sql.DB
	sqlString string
}

// NewImportDataServer -
func NewImportDataServer(db *sql.DB, sqlString string) *ImportdataServer {
	return &ImportdataServer{
		db: db,
		sqlString: sqlString,
	}
}

// ImportData -
func (s *ImportdataServer) ImportData(ctx context.Context, in *pb.ImportDataRequest) (*pb.ImportDataReply, error) {
	fileName := in.GetFileName()
	tableName := in.GetTableName()
	execString := fmt.Sprintf(s.sqlString, fileName, tableName)
	_, err := s.db.Exec(execString)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		return &pb.ImportDataReply{Message: "init mysql fail"}, nil
	}

	return &pb.ImportDataReply{Message: "init mysql succeed"}, nil
}