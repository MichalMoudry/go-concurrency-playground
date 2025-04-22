package main

import (
	"fmt"
	"sync"
)

func printSomthing(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"test 1",
		"test 2",
		"test 3",
		"test 4",
	}

	wg.Add(len(words))

	for i, v := range words {
		go printSomthing(fmt.Sprintf("%d: %s\n", i, v), &wg)
	}

	wg.Wait()

	wg.Add(1)
	printSomthing("Print something", &wg)

	fmt.Println("----------------")

	challange()
}
