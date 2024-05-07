package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	config := RESTNodeClientConfig{
		Hosts: strings.Split(os.Getenv("API_HOSTS"), ","),
	}
	client, err := New(config)
	if err != nil {
		panic(err)
	}
	r, err := client.GetActiveSet()
	if err != nil {
		panic(err)
	}
	fmt.Println("======")
	fmt.Println(len(r), r[0])
}

type NodeClient interface {
	GetActiveSet() ([]Validator, error)
}
