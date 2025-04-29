package main

import (
	"fmt"
	"time"
)

func Chapter02() {
	fmt.Println("Select with channels")
	fmt.Println("--------------------")
	chan1 := make(chan string)
	chan2 := make(chan string)

	go server1(chan1)
	go server2(chan2)
}

func server1(ch chan string) {
	for {
		time.Sleep(6 * time.Second)
		ch <- "this is from server 1"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "this is from server 2"
	}
}
