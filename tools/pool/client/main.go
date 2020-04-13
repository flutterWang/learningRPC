package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/flutterWang/learningRPC/tools/pool"
)

const (
	//模拟的最大goroutine
	maxGoroutine = 5
	//资源池的大小
	poolRes = 2
)

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutine)

	p, err := pool.New(createConnection, poolRes)
	if err != nil {
		log.Println(err)
		return
	}

	for query := 0; query < maxGoroutine; query++ {
		go func(q int) {
			dbQuery(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("start closing pool")
	p.Close()
}

func dbQuery(query int, pool *pool.Pool) {
	conn, err := pool.Get()
	if err != nil {
		log.Println(err)
		return
	}

	defer pool.Putback(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("query: %d，conneciton ID: %d", query, conn.(*dbConnection).ID)
}

type dbConnection struct {
	ID int32
}

func (db *dbConnection) Close() error {
	log.Println("close connection", db.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	return &dbConnection{id}, nil
}
