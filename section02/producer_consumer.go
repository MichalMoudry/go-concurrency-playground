package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	message       string
	pizzaNumber   int
	wasSuccessful bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(number int) *PizzaOrder {
	number += 1
	if number > NumberOfPizzas {
		return &PizzaOrder{pizzaNumber: number}
	}

	delay := rand.Intn(5) + 1
	fmt.Printf("Recevied order #%d!\n", number)
	rnd := rand.Intn(12) + 1 // err probability generation
	msg := ""
	wasSuccess := false

	if rnd < 5 {
		pizzasFailed += 1
	} else {
		pizzasMade += 1
	}
	total += 1
	fmt.Printf("Making pizza #%d. It will take %d seconds...\n", number, delay)
	time.Sleep(time.Duration(delay) * time.Second)

	if rnd <= 2 {
		msg = fmt.Sprintf("*** we ran out of ingredients for pizza #%d!", number)
	} else if rnd <= 4 {
		msg = fmt.Sprintf("*** the coook quit while making pizza #%d!", number)
	} else {
		wasSuccess = true
		msg = fmt.Sprintf("pizza order #%d is ready!", number)
	}

	return &PizzaOrder{
		message:       msg,
		pizzaNumber:   number,
		wasSuccessful: wasSuccess,
	}
}

func pizzeria(pizzaMaker *Producer) {
	i := 0

	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// tried to make pizza (but success is not required)
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func ProducerConsumer() {
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}
	go pizzeria(pizzaJob)

	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.wasSuccessful {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad!")
			}
		} else {
			color.Cyan("Done making pizzas")
			if err := pizzaJob.Close(); err != nil {
				color.Red("*** Error closing channel", err)
			}
		}
	}

	color.Cyan("----------------------------------")
	color.Cyan("Done for the day")

	color.Cyan("We made %d pizzas, but failed to make %d, with %d attempts in total", pizzasMade, pizzasFailed, total)
	switch {
	case pizzasFailed >= 9:
		color.Red("It was an awful day")
	case pizzasFailed >= 6:
		color.Red("It was not a very good day")
	case pizzasFailed >= 4:
		color.Yellow("It was an okay day")
	case pizzasFailed >= 2:
		color.Yellow("It was a pretty good day")
	default:
		color.Green("It was a great day")
	}
}
