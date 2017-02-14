package server

import (
	"encoding/json"
	"fmt"
	"net/http"

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
		json.NewEncoder(w).Encode(data)
		return
	}

	type Response struct {
		URL    string
		Method string
		Params map[string][]string
		Count  int
	}

	if r.ParseForm() != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	response := Response{
		URL:    r.URL.String(),
		Method: r.Method,
		Params: r.Form,
		Count:  count,
	}

	count++

	storage.Set(r.URL.String(), response)

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
