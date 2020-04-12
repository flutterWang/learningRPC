package importdata

import (
	"context"
	"log"
	"time"

	"github.com/flutterWang/learningRPC/basic/proto/csv"
	"google.golang.org/grpc"
)

// ImportData -
func ImportData(conn *grpc.ClientConn) {
	c := csv.NewTestClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ImportData(ctx, &csv.ImportDataRequest{
		FileName:  "./electricData.csv",
		TableName: "electrion",
	})

	log.Printf("Detail:%+v", r)
}
