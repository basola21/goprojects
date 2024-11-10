package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/basola21/02-backend-api/calculator"
	"github.com/basola21/02-backend-api/middleware"
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
	rateLimitInterval := 200 * time.Millisecond

	rateLimitedHandler := middleware.SimpleRateLimiter(
		http.HandlerFunc(RequestHandler),
		rateLimitInterval,
	)
	mux := http.NewServeMux()
	mux.Handle("/add", rateLimitedHandler)
	mux.Handle("/subtract", rateLimitedHandler)
	mux.Handle("/multiply", rateLimitedHandler)
	mux.Handle("/divide", rateLimitedHandler)
	mux.Handle("/sum", rateLimitedHandler)

	err := http.ListenAndServe("localhost:3000", mux)
	if err != nil {
		log.Fatal("server error")
	}
}
