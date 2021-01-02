package main

import (
	"fmt"
	"time"

	"github.com/masim05/play-go/calculator"
)

func main() {
	calc := calculator.NewCalculator()
	fmt.Println("calc.Add(2,3): ", calc.Add(2, 3))
	adderWithRL := calculator.NewAdderWithConcurrencyLimit(calc, 3)
	for i := 0; i < 20; i++ {
		go func(i int) {
			fmt.Println(adderWithRL.Add(i, 10), time.Now())

		}(i)
	}

	fmt.Scanln()

}
