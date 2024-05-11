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
	activeSetOne, err := client.GetActiveSet()
	if err != nil {
		panic(err)
	}
	fmt.Println("======")
	fmt.Println(len(activeSetOne))
	fmt.Println("head: ", activeSetOne[0:2])
	fmt.Println("tail: ", activeSetOne[len(activeSetOne)-2:])
	/*
	   activeSetTwo, err := client.GetActiveSet()

	   	if err != nil {
	   		panic(err)
	   	}

	   diff, err := sliceDiff(activeSetOne, activeSetTwo)

	   	if err != nil {
	   		panic(err)
	   	}

	   fmt.Println("diff: ", diff)
	*/
}

type NodeClient interface {
	GetActiveSet() ([]Validator, error)
}

func sliceDiff(s1 []any, s2 []any) (d []any, e error) {
	return nil, nil
}
