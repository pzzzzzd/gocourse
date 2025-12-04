package main

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	person // Embedded struct
	empId  string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
}

func (e Employee) introduce() {
	fmt.Printf("Hi, I'm %s, employee ID: %s, and I earn %.2f.\n", e.name, e.empId, e.salary)
}

func main() {

	emp := Employee{
		person: person{name: "Egor", age: 18},
		empId:  "YY01", salary: 33555,
	}

	fmt.Println("Name:", emp.name) // Accessinbg the embadded struct field emp.person.name
	fmt.Println("Age:", emp.age)   // Same as above
	fmt.Println("Emp ID:", emp.empId)
	fmt.Println("Salart:", emp.salary)

	emp.introduce()

}
