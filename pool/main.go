package main

import (
	"fmt"
	"time"
)

func newPool(f func()) func() {
	var ch chan int = make(chan int, 3)

	return func() {
		ch <- 1

		defer func() {
			<-ch
		}()

		f()
	}
}

func show() {
	time.Sleep(1 * time.Second)
	fmt.Println(time.Now())
}

func main() {
	pool := newPool(show)
	go pool()
	go pool()
	go pool()
	go pool()
	go pool()
	go pool()
	go pool()
	go pool()
	go pool()

	fmt.Scanln()
}
