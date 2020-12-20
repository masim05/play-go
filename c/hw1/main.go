package main

import (
	"fmt"
	"os"
)

type config struct {
	showFiles bool
}

type state struct {
	depth int
}

func main() {
	cfg := config{true}
	st := state{1}
	notLast := []bool{false}
	res, err := tree(cfg, st, ".", notLast)
	if err != nil {
		fmt.Println("ERR:", err)
	}
	fmt.Println(res)
}

func tree(cfg config, st state, path string, notLast []bool) (result string, err error) {
	dir, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer dir.Close()

	names, err := dir.Readdir(-1)
	if err != nil {
		return "", err
	}

	for i, f := range names {
		printName(cfg, st, f.Name(), i, len(names), notLast)

		newSt := state{depth: st.depth + 1}
		notLast = append(notLast, false)
		if f.IsDir() {
			_, err := tree(cfg, newSt, path+"/"+f.Name(), notLast)
			if err != nil {
				return "", err
			}
		}
	}

	return result, nil
}

func printName(cf config, st state, name string, idx, length int, notLast []bool) {
	for i := 0; i < st.depth-1; i++ {
		if notLast[i] {
			fmt.Print("|   ")
		} else {
			fmt.Print("    ")
		}
	}
	if idx < length-1 && length > 1 {
		fmt.Printf("├───%s\n", name)
		notLast[st.depth-1] = true
	}
	if idx == length-1 {
		fmt.Printf("└───%s\n", name)
		notLast[st.depth-1] = false
	}
}
