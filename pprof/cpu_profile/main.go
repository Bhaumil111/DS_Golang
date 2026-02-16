package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

// expensive function -> we want profiler to catch this
func doWork() {
	data := make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		data = append(data, rand.Intn(1000))
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 500; i++ {
		doWork()
	}
	fmt.Fprintln(w, "done")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", handler)

	fmt.Println("server started at :8080")
	http.ListenAndServe(":8080", nil)
}

//for pprof
//go tool pprof http://localhost:8080/debug/pprof/profile?seconds=20
//
