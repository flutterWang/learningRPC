package service

import (
	"context"
	"database/sql"
	"fmt"

	pb "github.com/flutterWang/learningRPC/basic/proto/test"
)

// TestServer -
type TestServer struct {
	db        *sql.DB
	sqlString string
}

// NewTestServer -
func NewTestServer() *TestServer {
	return &TestServer{}
}

// Test -
func (s *TestServer) Test(ctx context.Context, in *pb.TestRequest) (*pb.TestResponse, error) {
	query := in.GetQuery()
	fmt.Println("string:", query)

	obj := in.GetType()
	fmt.Println("enum type:", obj)

	testMap := in.GetTestMap()
	fmt.Println("map type:", testMap)

	testSubMessage := in.GetChild()
	fmt.Println("subMessage type:", testSubMessage)

	testchar := in.GetChar()
	fmt.Println("char type:", testchar)

	testSlice := in.GetSnippets()
	fmt.Println("slice type:", testSlice)

	return &pb.TestResponse{State: "hello"}, nil
}
