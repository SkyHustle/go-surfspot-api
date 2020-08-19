package main

import (
	"fmt"
	"net/http"
)

type Surfspot struct {
	Name       string
	Founder    string
	ID         string
	Beach      string
	Difficulty int
}

// surfspotsHandler handles http request and response
func surfspotsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("w \n", w)
	fmt.Println("r \n", r)
	fmt.Printf("w is of type %T \n and r is of type %T ", w, r)
}

func main() {
	// HandleFunc registers surfspotsHandler for "/surfspots"
	http.HandleFunc("/surfspots", surfspotsHandler)

	// Simple http server that takes a port and a default handler
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
