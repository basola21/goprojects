package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/add", add)

	err := http.ListenAndServe("localhost:3000", mux)
	if err != nil {
		log.Fatal("server error")
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	number1 := r.FormValue("nubmer1")
	number2 := r.FormValue("nubmer2")

	w.Write([]byte(fmt.Sprintf("recived these values %s %s", number1, number2)))
}
