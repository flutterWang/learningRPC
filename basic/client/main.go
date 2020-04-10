package main

import (
	"context"
	"log"
	"time"

	"github.com/flutterWang/learningRPC/basic/proto/test"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

const (
	HTML test.DocumentExtType = 0
	XML  test.DocumentExtType = 1
	JSON test.DocumentExtType = 2
	PDF  test.DocumentExtType = 3
)

// Test -
func Test(conn *grpc.ClientConn) {
	c := test.NewTestClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	testmap := make(map[string]string)
	testmap["one"] = "hello"

	char := []byte("hello world")

	child := &test.TestRequest_Child{
		Age: 10,
	}

	snippets := []string{"hello", "world"}

	oneof := &test.TestRequest_Name{
		Name: "wbofeng",
	}

	// r, err := c.ImportData(ctx, &csv.ImportDataRequest{FileName: "./electricData.csv", TableName: "electrion"})
	r, err := c.Test(ctx, &test.TestRequest{
		Query:     "hello",
		Type:      JSON,
		TestMap:   testmap,
		Child:     child,
		Char:      char,
		Snippets:  snippets,
		TestOneof: oneof,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	var detail test.TestDetail
	err = ptypes.UnmarshalAny(r.GetDetails(), &detail)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Detail:%+v", detail)
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// c := csv.NewCsvClient(conn)

	Test(conn)
}
