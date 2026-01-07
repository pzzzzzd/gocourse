package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name" db:"firstn" xml:"first"`
	LastName  string `json:"last_name,omitempty"`
	Age       int    `json:"age,omitempty"`
}

func main() {
	// Person{FirstName: "John", LastName: "Don", Age: 22}
	person := Person{FirstName: "John", Age: 0}

	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	fmt.Println(string(jsonData))

}
