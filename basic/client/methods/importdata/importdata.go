package importdata

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/flutterWang/learningRPC/basic/proto/csv"
	"google.golang.org/grpc"
)

// Upload -
func Upload(conn *grpc.ClientConn, filepath string, filename string) {
	c := csv.NewCsvClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fd, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return
	}

	r, err := c.Upload(ctx, &csv.UploadRequest{
		File: fd,
		Name: "electron",
	})

	log.Printf("Detail:%+v", r)
}
