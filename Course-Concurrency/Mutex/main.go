package main

import (
	"fmt"
	"sync"
)

type Income struct {
	source string
	amount int
}

var wg sync.WaitGroup

func main() {
	var bankBalance int
	var Balance sync.Mutex

	fmt.Printf("Initial account balance : $%d.00\n", bankBalance) // Changed from fmt.Println to fmt.Printf
	fmt.Println()

	incomes := []Income{
		{source: "Main job", amount: 50000},
		{source: "Gifts", amount: 1000},
		{source: "Part time job", amount: 5000},
		{source: "Investment", amount: 10000},
	}

	wg.Add(4)

	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for j := 1; j <= 52; j++ {
				Balance.Lock()
				temp := bankBalance
				temp += income.amount
				bankBalance = temp
				Balance.Unlock()

				fmt.Printf("On week %d, you earned $%d.00, from %s\n", j, income.amount, income.source)
			}
		}(i, income)
	}
	wg.Wait()

	fmt.Printf("Final Bank Balance : $%d.00\n", bankBalance) // Changed from fmt.Println to fmt.Printf
}
