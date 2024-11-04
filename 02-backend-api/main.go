package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Numbers struct {
	Number1 string `json:"number1"`
	Number2 string `json:"number2"`
}

type Results struct {
	Result float64 `json:"result"`
}

func calculateTwoNumbers(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	operation := strings.ToLower(path)

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

	var result Results

	switch operation {
	case "add":
		result = Results{Result: number1 + number2}
	case "subtract":
		result = Results{Result: number1 - number2}
	case "multiply":
		result = Results{Result: number1 * number2}
	case "divid":
		if number2 == 0 {
			http.Error(w, "can not divide by 0", http.StatusBadRequest)
			return
		}
		result = Results{Result: number1 / number2}
	default:
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}

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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/add", calculateTwoNumbers)
	mux.HandleFunc("/subtract", calculateTwoNumbers)
	mux.HandleFunc("/multiply", calculateTwoNumbers)
	mux.HandleFunc("/divide", calculateTwoNumbers)

	err := http.ListenAndServe("localhost:3000", mux)
	if err != nil {
		log.Fatal("server error")
	}
}
