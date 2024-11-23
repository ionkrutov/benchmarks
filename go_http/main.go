package main

import (
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!\n")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
