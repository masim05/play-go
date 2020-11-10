package main

import (
	"fmt"
	"net/http"
	"time"
	"strconv"
	"math/rand"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	duration := 0
	switch true {
	case len(query["r"]) > 0:
		d, err := strconv.ParseUint(query["r"][0], 10, 16)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error while parsing r parameter."))
			fmt.Printf("Error while parsing r parameter: %s\n", err)
			return
		}
		duration = rand.Intn(int(d))
	case len(query["t"]) > 0:
		d, err := strconv.ParseUint(query["t"][0], 10, 16)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error while parsing t parameter."))
			fmt.Printf("Error while parsing t parameter: %s\n", err)
			return
		}
		duration = int(d)
	default:
		fmt.Printf("neither r nor t.\n")

	}

	if duration > 0 {
		time.Sleep(time.Duration(duration) * time.Millisecond)
	}

	fmt.Fprintf(w, "Pong. Slept for %v milliseconds.\n", duration)
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error while ListenAndServe: %s\n", err)
		return
	}
}