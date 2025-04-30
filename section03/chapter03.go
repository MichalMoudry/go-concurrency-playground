package main

import (
	"fmt"
	"time"
)

func Chapter03() {
	ch := make(chan int, 10)
	go listenToChan(ch)

	for i := range 100 {
		fmt.Println("Sending", i, "to channel...")
		ch <- i
		fmt.Println("Sent", i, "to channel")
	}

	fmt.Println("Done!")
	close(ch)
}

func listenToChan(ch chan int) {
	for {
		// print a got data message
		i := <-ch
		fmt.Println("Got:", i, "from channel")

		// simulate doing a lot of work
		time.Sleep(1 * time.Second)
	}
}
