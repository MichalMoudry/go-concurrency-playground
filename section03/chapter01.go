package main

import (
	"fmt"
	"strings"
)

func Chapter01() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press ENTER")
	for {
		var input string
		fmt.Print("Input: ")
		_, _ = fmt.Scanln(&input)
		if input == "q" || input == "Q" {
			break
		} else if input != "" {
			ping <- input
			response := <-pong
			fmt.Println("Response is:", response)
		}
	}

	fmt.Println("All done closing channels")
	close(ping)
	close(pong)
}

func shout(ping <-chan string, pong chan<- string) {
	for {
		str := <-ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(str))
	}
}
