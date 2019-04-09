package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	log.Print("Listening...")
	http.HandleFunc("/colors", colors)
	http.ListenAndServe(":8888", nil)
}

type Response struct {
	Colors Colors `json:"colors,omitempty"`
}

type Colors struct {
	X string `json:"x,omitempty"`
	O string `json:"o,omitempty"`
}

func colors(w http.ResponseWriter, r *http.Request) {
	log.Print("got something")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	resp := &Response{
		Colors: Colors{
			X: "black",
			O: "black",
		},
	}
	contents, err := json.Marshal(resp)
	if err != nil {
		log.Printf("error marshaling json: %v", err)
		return
	}
	w.Write(contents)
}
