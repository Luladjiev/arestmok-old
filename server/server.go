package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/luladjiev/arestmok/analyzer"
	"github.com/luladjiev/arestmok/storage"
)

var count = 1

func handleFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() == "/favicon.ico" {
		http.NotFound(w, r)
		return
	}

	data, err := storage.Get(r.URL.String())

	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data.Structure)
		return
	}

	type Response struct {
		URL    string
		Method string
		Query  map[string][]string
		Body   interface{}
		Count  int
	}

	if r.ParseForm() != nil {
		fmt.Println("Error parsing request form")
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	var body map[string]interface{}

	decoder := json.NewDecoder(r.Body)

	if decoder.Decode(&body) != nil {
		fmt.Println("Error decoding request body")
		http.Error(w, "", http.StatusInternalServerError)
	}

	analyzer.Analyze(body)

	response := Response{
		URL:    r.URL.String(),
		Method: r.Method,
		Query:  r.Form,
		Body:   body,
		Count:  count,
	}

	count++

	route := storage.RouteStruct{Structure: response}

	storage.Set(r.URL.String(), route)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Run arestmok server
func Run(port int) {
	fmt.Println()
	fmt.Println("Starting arestmok server on address", port)

	http.HandleFunc("/", handleFunc)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
