package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", helloFunction)
	http.ListenAndServe(":8080", nil)
}

func helloFunction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Next 2019!!!!\n"))
}
