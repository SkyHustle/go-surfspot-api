package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Surfspot is a data structure that holds characteristics of a surfspot
type Surfspot struct {
	Name       string
	Founder    string
	ID         string
	Beach      string
	Difficulty int
}

// surfspotHandler is a data structure that is used as a temporary data store
type surfspotHandlers struct {
	// sync.Mutex Handles concurrent requests
	// incase there's a get and post request in parallel
	sync.Mutex
	store map[string]Surfspot
}

// surfspots is a method for surfspothandlers
// that checks what type of HTTP request is made
func (h *surfspotHandlers) surfspots(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
}

// post handles http POST request and response
// get is a method for surfspotHandlers
func (h *surfspotHandlers) post(w http.ResponseWriter, r *http.Request) {
}

// get handles http GET request and response
// get is a method for surfspotHandlers
func (h *surfspotHandlers) get(w http.ResponseWriter, r *http.Request) {
	surfspots := make([]Surfspot, len(h.store))

	// Lock the store when we read
	h.Lock()
	i := 0
	for _, surfspot := range h.store {
		surfspots[i] = surfspot
		i++
	}
	// Unlock the store when we finish reading
	h.Unlock()

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
