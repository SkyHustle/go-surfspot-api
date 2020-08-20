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
// And is the method receiver surfspotHandlers
func (h *surfspotHandlers) get(w http.ResponseWriter, r *http.Request) {

}

// newSurfspotHandlers is a contructor function that does not take any arguments
// returns surfspotHandlers object
func newSurfspotHandlers() *surfspotHandlers {
	return &surfspotHandlers{
		store: map[string]Surfspot{
			"id1": Surfspot{
				Name:       "Pipeline",
				Founder:    "Jerry Lopez",
				ID:         "id1",
				Beach:      "Ehukai",
				Difficulty: 10,
			},
		},
	}
}

func main() {
	surfspotHandlers := newSurfspotHandlers()
	fmt.Println(surfspotHandlers)
	fmt.Println(surfspotHandlers.store)
	fmt.Println(surfspotHandlers.store["id1"])
	fmt.Println(surfspotHandlers.store["id1"].Name)

	// HandleFunc registers surfspotHandlers for "/surfspots"
	http.HandleFunc("/surfspots", surfspotHandlers.get)

	// Simple http server that takes a port and a default handler
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
