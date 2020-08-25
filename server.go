package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

// Surfspot is a data structure that holds characteristics of a surfspot
type Surfspot struct {
	ID         string
	Name       string
	Founder    string
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
	// Convert raw request body to bytes
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)))
		return
	}

	// UnMarshal bodyBytes into a Body object
	var surfspot Surfspot
	err = json.Unmarshal(bodyBytes, &surfspot)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	// Set a unique ID
	surfspot.ID = fmt.Sprintf("%d", time.Now().UnixNano())

	// Lock the store when we write
	h.Lock()
	h.store[surfspot.ID] = surfspot
	defer h.Unlock()
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

// getRandomSurfspot gets a random surfspot from the store
func (h *surfspotHandlers) getRandomSurfspot(w http.ResponseWriter, r *http.Request) {
	ids := make([]string, len(h.store))

	h.Lock()
	i := 0
	for id := range h.store {
		ids[i] = id
		i++
	}
	defer h.Unlock()

	var target string
	if len(ids) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if len(ids) == 1 {
		target = ids[0]
	} else {
		rand.Seed(time.Now().UnixNano())
		target = ids[rand.Intn(len(ids))]

	}

	fmt.Println(target)
}

// getSurfSpot retrieves surfspot by ID
func (h *surfspotHandlers) getSurfspot(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if parts[2] == "random" {
		h.getRandomSurfspot(w, r)
		return
	}

	h.Lock()
	surfspot, ok := h.store[parts[2]]
	h.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(surfspot)
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
			// "id1": Surfspot{
			// 	ID:         "id1",
			// 	Name:       "Pipeline",
			// 	Founder:    "Jerry Lopez",
			// 	Beach:      "Ehukai",
			// 	Difficulty: 10,
			// },
		},
	}
}

type adminPortal struct {
	password string
}

// newAdmin Portal returns an instance of adminPortal
func newAdminPortal() *adminPortal {
	password := os.Getenv("ADMIN_PASSWORD")
	if password == "" {
		panic("required env var ADMIN_PASSWORD not set")
	}
	return &adminPortal{password: password}
}

func (a adminPortal) handler(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok || user != "admin" || pass != a.password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - unauthorized"))
		return
	}

	w.Write([]byte("<html><h2>Secret Admin Portal</h2></html>"))
}

func main() {
	admin := newAdminPortal()
	http.HandleFunc("/admin", admin.handler)

	surfspotHandlers := newSurfspotHandlers()
	fmt.Println(surfspotHandlers.store)

	// HandleFunc registers surfspotHandlers for "/surfspots"
	http.HandleFunc("/surfspots", surfspotHandlers.surfspots)

	http.HandleFunc("/surfspots/", surfspotHandlers.getSurfspot)

	// Simple http server that takes a port and a default handler
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
