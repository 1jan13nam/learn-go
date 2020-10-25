package main

import (
	"io"
	"net/http"
)

type serve int

func (d serve) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Pong")
}

type health int

func (c health) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Health Check Passed , Version 1.0")
}

func main() {
	var d serve
	var c health

	mux := http.NewServeMux()
	mux.Handle("/ping", d)
	mux.Handle("/healthz", c)

	http.ListenAndServe(":8080", mux)
}
