package main

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

func ProducerConsumer() {

}
