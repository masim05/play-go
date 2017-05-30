package main

import (
	"fmt"
	_ "github.com/masim05/play-go/import/one"
	_ "github.com/masim05/play-go/import/two"
)

func main() {
	fmt.Println("Import!")
	a := 1
	a, b := 2, 3
	fmt.Println(a, b)
}
