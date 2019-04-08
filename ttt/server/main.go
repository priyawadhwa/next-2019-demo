package main

import (
	"log"
	"net/http"
)

const port = "8080"

func main() {
	fs := http.FileServer(http.Dir("/dist"))
	http.Handle("/", fs)

	log.Printf("Serving on HTTP port: %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
