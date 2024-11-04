package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Numbers struct {
	Number1 string `json:"number1"`
	Number2 string `json:"number2"`
}

type Results struct {
	Result float64 `json:"result"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/add", add)

	err := http.ListenAndServe("localhost:3000", mux)
	if err != nil {
		log.Fatal("server error")
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	var numbers Numbers

	if err := json.NewDecoder(r.Body).Decode(&numbers); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	number1, err := castString(numbers.Number1)
	number2, err := castString(numbers.Number2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sum := number1 + number2

	result := Results{Result: sum}

	json.NewEncoder(w).Encode(&result)

	w.Header().Set("Content-Type", "application/json")

}

func castString(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}
