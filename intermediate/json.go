package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string  `json:"first_name"`
	Age       int     `json:"age,omitempty"`
	Email     string  `json:"email,omitempty"`
	Address   Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	person := Person{FirstName: "Jogn"}

	// Marshalling
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))

	person1 := Person{FirstName: "Jane", Age: 33, Email: "ex@ex.x", Address: Address{City: "London", State: "London"}}

	jsonData1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData1))

	jsonData2 := `{"full_name":"Jenny Doe", "emp_id":"0009", "age": 27, "address":{"city":"San Jose", "state":"CA"}}`

	var employeeFromJson Employee
	err = json.Unmarshal([]byte(jsonData2), &employeeFromJson)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(employeeFromJson)
	fmt.Println("Age + 5:", employeeFromJson.Age+5)

	listOfCityState := []Address{
		{City: "New York", State: "NY"},
		{City: "San Jose", State: "CA"},
		{City: "Las Vegas", State: "NV"},
	}

	fmt.Println(listOfCityState)
	jsonList, err := json.Marshal(listOfCityState)
	if err != nil {
		log.Fatalln("Error", err)
	}
	fmt.Println("JSON list", string(jsonList))

	// handling unknown json structures
	jsonData3 := `{"full_name":"John", "age": 44, "address":{"city":"New York", "state":"NY"}}`

	var data map[string]interface{}

	err = json.Unmarshal([]byte(jsonData3), &data)
	if err != nil {
		log.Fatalln("Error", err)
	}
	fmt.Println("Unmarshalled JSON:", data)
	fmt.Println("Unmarshalled JSON Address:", data["address"])
}

type Employee struct {
	FullName string  `json:"full_name"`
	EmpID    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}
