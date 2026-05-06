package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

var personData = map[string]Person{
	"1": {Name: "JD", Age: 30},
	"2": {Name: "AW", Age: 18},
	"3": {Name: "FG", Age: 52},
}

func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is missing", http.StatusBadRequest)
		return
	}

	person, exists := personData[id]
	if !exists {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(person); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func main() {

	port := 3011

	fmt.Printf("Server started on port %d\n", port)

	http.HandleFunc("/person", getPersonHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}
