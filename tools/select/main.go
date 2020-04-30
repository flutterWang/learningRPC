package main

import "fmt"

func testClose() {
	ch := make(chan int)
	select {
	case <-ch:
		fmt.Println("get")
	default:
		fmt.Println("close")
		close(ch)
	}
}

// func isClosed(ch <-chan int) {
// 	close(ch)
// }

func readChannel() {
	done := make(chan int, 1)
	done <- 1

	close(done)
	for i := 1; i <= 3; i++ {
		t, ok := <-done
		fmt.Println(i, ":", t, ok)
	}
}

func main() {
	// c := make(chan int)

	// testClose()
	// isClosed(c)
	// fmt.Println(isClosed(c))
	// close(c)
	// fmt.Println(isClosed(c))
	// readChannel()

	count := 0
	closed := make(chan int)
	buffered := make(chan int, 10)

	for i := 0; i < 10; i++ {
		buffered <- i
	}

	close(closed)

	select {
	case <-closed:
		count++
	case <-buffered:
	}

	fmt.Println(count, 10-count)
}
