package main

import (
	"fmt"
	//	"time"
)

func read(c chan int, p *int) {
	*p = <-c
}

func write(c chan int, i int) {
	c <- i
	fmt.Println(len(c), "len")
}

func main() {
	c := make(chan int, 2)

	//i := 0

	fmt.Println(cap(c), "cap")
	go write(c, 5)

	go write(c, 6)

	go write(c, 7)

	go write(c, 8)

	go write(c, 9)

	//go read(c, &i)
	fmt.Println(<-c)
	fmt.Println(len(c), "len")
	fmt.Println(<-c)
	fmt.Println(len(c), "len")
	fmt.Println(<-c)
	fmt.Println(len(c), "len")
	fmt.Println(<-c)
	fmt.Println(len(c), "len")

	fmt.Println("===")
}
