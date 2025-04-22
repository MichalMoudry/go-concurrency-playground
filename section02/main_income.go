package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount float64
}

func IncomeExercise() {
	bankBalance := float64(0)
	balanceMut := sync.Mutex{}
	fmt.Printf("Bank balance: %f.00\n", bankBalance)

	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))
	for i, income := range incomes {
		//fmt.Printf("%d: %v\n", i, income)
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balanceMut.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balanceMut.Unlock()
				fmt.Printf("On week %d, you earned $%f from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}
	wg.Wait()

	fmt.Printf("Bank balance: %f\n", bankBalance)
}

/*import (
	"fmt"
	"sync"
)

var (
	msg string
	wg  sync.WaitGroup
)

func updateMessage(s string, mut *sync.Mutex) {
	defer wg.Done()
	mut.Lock()
	msg = s
	mut.Unlock()
}

func main() {
	msg = "Hello, world!"
	var mut sync.Mutex

	wg.Add(2)
	go updateMessage("Hello, universe!", &mut)
	go updateMessage("Hello, cosmos!", &mut)
	wg.Wait()

	fmt.Println(msg)
}*/
