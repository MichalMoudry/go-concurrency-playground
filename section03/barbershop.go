package main

import (
	"time"

	"github.com/fatih/color"
)

type Barbershop struct {
	Capacity        int
	HaircutDuration time.Duration
	NumOfBarbers    int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	IsOpen          bool
}

func (s *Barbershop) AddBarber(barber string) {
	s.NumOfBarbers += 1
	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients...", barber)
		for {
			if len(s.ClientsChan) == 0 {
				color.Yellow("There is nothing to do, %s takes a nap...", barber)
				isSleeping = true
			}

			client, shopOpen := <-s.ClientsChan
			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up", client, barber)
					isSleeping = false
				}
				// cut hair
				s.CutHair(barber, client)
			} else {
				// shop is closed, so send the barber home and close this goroutine
				s.SendBarberHome(barber)
				return
			}
		}
	}()
}

func (s *Barbershop) CutHair(barber, client string) {
	color.Green("%s is cutting %s's hair...", barber, client)
	time.Sleep(s.HaircutDuration)
	color.Green("%s is finished cutting %s's hair...", barber, client)
}

func (s *Barbershop) SendBarberHome(barber string) {
	color.Magenta("%s is going home", barber)
	s.BarbersDoneChan <- true
}

func (s *Barbershop) Close() {
	color.Black("Closing the shop for the day")
	close(s.ClientsChan)
	s.IsOpen = false

	for a := 1; a <= s.NumOfBarbers; a++ {
		<-s.BarbersDoneChan
	}

	close(s.BarbersDoneChan)
	color.Green("--------------------------------------------------------------------")
	color.Green("The barbershop is now closed for the day, and everyone has gone home")
}

func (s *Barbershop) AddCustomer(name string) {
	color.Black("*** %s arrives", name)
	if s.IsOpen {
		select {
		case s.ClientsChan <- name:
			color.Blue("%s takes a seat in the waiting room", name)
		default:
			color.Red("The waiting room, so %s leaves!", name)
		}
	} else {
		color.Red("The shop is closed, so %s leaves!", name)
		return
	}
}
