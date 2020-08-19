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

// Temporary data store
type surfspotHandlers struct {
	store map[string]Surfspot
}

// surfspotHandlers handles http request and response
func (h *surfspotHandlers) get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("w \n", w)
	fmt.Println("r \n", r)
	fmt.Printf("w is of type %T \n and r is of type %T ", w, r)
}

func newSurfspotHandlers() *surfspotHandlers {
	return &surfspotHandlers{
		store: map[string]Surfspot{},
	}
}

func main() {
	surfspotHandlers := newSurfspotHandlers()
	fmt.Println(surfspotHandlers)
	// HandleFunc registers surfspotHandlers for "/surfspots"
	http.HandleFunc("/surfspots", surfspotHandlers.get)

	// Simple http server that takes a port and a default handler
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
