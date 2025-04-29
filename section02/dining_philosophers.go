package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
)

// Is a struct which stores information about a philosopher
type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

type RestaurantLog struct {
	Items []string
	Mut   *sync.Mutex
}

var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

// how many times a person eats
const hunger int = 3
const eatTime time.Duration = 1 * time.Second
const thinkTime time.Duration = 3 * time.Second
const sleepTime time.Duration = 1 * time.Second

func DiningPhilosophers() {
	color.Cyan("Dining philosophers problem")
	color.Cyan("---------------------------")
	fmt.Println("The table is empty...")

	dine()

	color.Green("The table is empty...")
}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seatedWg := &sync.WaitGroup{}
	seatedWg.Add(len(philosophers))

	// a map of forks
	forks := make(map[int]*sync.Mutex)
	for i := range philosophers {
		forks[i] = &sync.Mutex{}
	}

	log := RestaurantLog{
		Items: make([]string, 0, len(philosophers)),
		Mut:   &sync.Mutex{},
	}
	// start the meal
	for _, v := range philosophers {
		go diningProblem(v, wg, forks, seatedWg, &log)
	}

	wg.Wait()
	color.Cyan("Customer log:")
	for _, v := range log.Items {
		fmt.Printf("\t- %s\n", v)
	}
}

func diningProblem(p Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seatedWg *sync.WaitGroup, log *RestaurantLog) {
	defer wg.Done()

	// seat the philosopher
	fmt.Printf("%s is seated at the table\n", p.name)
	seatedWg.Done()
	seatedWg.Wait()

	// eat three times
	for range hunger {
		// get a lock on both forks
		if p.leftFork > p.rightFork {
			forks[p.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork\n", p.name)
			forks[p.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork\n", p.name)
		} else {
			forks[p.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork\n", p.name)
			forks[p.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork\n", p.name)
		}

		fmt.Printf("\t%s has both forks and is eating\n", p.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking\n", p.name)
		time.Sleep(thinkTime)

		forks[p.leftFork].Unlock()
		forks[p.rightFork].Unlock()
		fmt.Printf("\t%s put down both forks\n", p.name)
	}

	fmt.Println(p.name, "is satisfied")
	fmt.Println(p.name, "left the table")
	log.Mut.Lock()
	log.Items = append(log.Items, p.name)
	log.Mut.Unlock()
}
