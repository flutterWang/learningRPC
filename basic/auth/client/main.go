package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/testdata"

	pb "github.com/flutterWang/learningRPC/basic/proto/echo"
)

func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "some-secret-token",
	}
}

func main() {
	perRPC := oauth.NewOauthAccess(fetchToken())
	creds, err := credentials.NewClientTLSFromFile(testdata.Path("ca.pem"), "x.test.youtube.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(perRPC),
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial("localhost:8888", opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn)
	resp, err := client.UnaryEcho(context.Background(), &pb.EchoRequest{
		Message: "hello world",
	})

	if err != nil {
		log.Fatalf("client.UnaryEcho(_) = _, %v: ", err)
	}

	fmt.Println("UnaryEcho: ", resp.Message)
}
