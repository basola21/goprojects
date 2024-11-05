package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/basola21/02-backend-api/calculator"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	operation := strings.ToLower(path)

	var numbers calculator.Numbers
	if err := json.NewDecoder(r.Body).Decode(&numbers); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := calculator.CalculateTwoNumbers(operation, numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
    return
	}

	json.NewEncoder(w).Encode(&result)

	w.Header().Set("Content-Type", "application/json")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/add", RequestHandler)
	mux.HandleFunc("/subtract", RequestHandler)
	mux.HandleFunc("/multiply", RequestHandler)
	mux.HandleFunc("/divide", RequestHandler)
	mux.HandleFunc("/sum", RequestHandler)

	err := http.ListenAndServe("localhost:3000", mux)
	if err != nil {
		log.Fatal("server error")
	}
}
