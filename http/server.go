package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"sync"
	"time"
)

func sleepHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s /sleep\n", time.Now().String())

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

var mu sync.Mutex

func mutexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s /mutex\n", time.Now().String())

	query := r.URL.Query()
	d, err := strconv.ParseUint(query["r"][0], 10, 16)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while parsing r parameter."))
		fmt.Printf("Error while parsing r parameter: %s\n", err)
		return
	}
	duration := rand.Intn(int(d))

	mu.Lock()
	// time.Sleep(time.Duration(duration) * time.Millisecond)
	time.Sleep(200 * time.Millisecond)
	mu.Unlock()

	fmt.Fprintf(w, "Pong. Slept for %v milliseconds.\n", duration)
}

func main() {
	http.HandleFunc("/sleep", sleepHandler)
	http.HandleFunc("/mutex", mutexHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error while ListenAndServe: %s\n", err)
		return
	}
}
