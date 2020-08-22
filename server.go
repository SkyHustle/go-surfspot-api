package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Surfspot struct defines the characteristics of a surfspot
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

func (h *surfspotHandlers) surfspots(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	}
}

// surfspotHandlers handles http request and response
// And is the method receiver surfspotHandlers
func (h *surfspotHandlers) get(w http.ResponseWriter, r *http.Request) {
	surfspots := make([]Surfspot, len(h.store))

	i := 0
	for _, surfspot := range h.store {
		surfspots[i] = surfspot
		i++
	}

	jsonBytes, err := json.Marshal(surfspots)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
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
	fmt.Println(surfspotHandlers.store["id1"])

	// HandleFunc registers surfspotHandlers for "/surfspots"
	http.HandleFunc("/surfspots", surfspotHandlers.surfspots)

	// Simple http server that takes a port and a default handler
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
