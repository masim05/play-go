package main

import (
	"fmt"
	//	"time"
)

func write(a chan int, i int) {
	a <- i
}

func sel(a, b chan int) {
	fmt.Println("SEL")

	for {
		select {
		case <-a:
			fmt.Println("a")
		case <-b:
			fmt.Println("b")
			return
		}
	}
}

func main() {
	fmt.Println("====")
	a := make(chan int)
	b := make(chan int)

	go write(b, 5)

	go write(a, 1)
	go write(a, 2)
	go write(a, 3)
	go write(a, 4)
	go write(a, 5)

	sel(a, b)

}
