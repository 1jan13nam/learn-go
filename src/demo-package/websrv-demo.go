package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func live(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PONG")
}

func main() {
	http.HandleFunc("/ping", greet)
	http.HandleFunc("/live", live)
	http.ListenAndServe(":8080", nil)
}
