package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const (
	seatingCapacity int           = 10
	arrivalRate     int           = 100
	cutDuration     time.Duration = 1000 * time.Millisecond
	timeOpen        time.Duration = 10 * time.Second
)

func SleepingBarber() {
	// print a welcome message
	color.Cyan("The sleeping barber problem")

	// create channels if we need any
	clientChannel := make(chan string, seatingCapacity)
	doneChannel := make(chan bool)

	// create the barbershop
	barbershop := Barbershop{
		Capacity:        seatingCapacity,
		HaircutDuration: cutDuration,
		NumOfBarbers:    0,
		BarbersDoneChan: doneChannel,
		ClientsChan:     clientChannel,
		IsOpen:          true,
	}

	color.Green("The shop is open for the day!")

	// add barbers
	barbershop.AddBarber("Frank")
	barbershop.AddBarber("Lucy")
	barbershop.AddBarber("Kelly")
	barbershop.AddBarber("Pat")
	barbershop.AddBarber("Alice")
	barbershop.AddBarber("Bob")

	// start the barbershop as a goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		barbershop.Close()
		closed <- true
	}()

	// add clients
	i := 1
	go func() {
		for {
			// get a random with an average arrival rate
			rand := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(rand)):
				barbershop.AddCustomer(fmt.Sprintf("Client #%d", i))
				i += 1
			}
		}
	}()

	// block until the barbershop is closed
	<-closed
}
