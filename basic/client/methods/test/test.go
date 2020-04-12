package test

import (
	"context"
	"log"
	"time"

	"github.com/flutterWang/learningRPC/basic/proto/test"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
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

	testmap := make(map[int32]string)
	testmap[1] = "hello"

	char := []byte("hello world")

	child := &test.TestRequest_Child{
		Name: 10,
	}

	snippets := []string{"hello", "world"}

	oneof := &test.TestRequest_Name{
		Name: "wbofeng",
	}

	// subMessage := &sub.SubMessage{
	// 	Msg: "hello",
	// }

	// r, err := c.ImportData(ctx, &csv.ImportDataRequest{FileName: "./electricData.csv", TableName: "electrion"})
	r, err := c.Test(ctx, &test.TestRequest{
		// Query:     "hello",
		Type:      1,
		TestMap:   testmap,
		Char:      char,
		Snippets:  snippets,
		Child:     child,
		TestOneof: oneof,
		// Submessage: subMessage,
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
