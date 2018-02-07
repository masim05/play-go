package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

func main() {

	connectionsTotal := 0

	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println("net.Listen() error:", err)
		os.Exit(1)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("ln.Accept() error:", err)
			continue
		}

		fmt.Println("ln.Accept() success.")

		connectionsTotal = connectionsTotal + 1
		conn.Write([]byte(fmt.Sprintf("Connection #%d\n", connectionsTotal)))

		go blockMe()
	}
}

func blockMe() {
	fmt.Printf("In blockMe(), %d goroutines.\n", runtime.NumGoroutine())
	for {
	}
}
