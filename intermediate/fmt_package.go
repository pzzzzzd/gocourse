package main

import "fmt"

func main() {

	// // Printing Func
	// fmt.Print("Hello ")
	// fmt.Print("World!")
	// fmt.Print(12, 355)

	// fmt.Println("\nHello ")
	// fmt.Println("World!")
	// fmt.Println(12, 355)

	// name := "John"
	// age := 25
	// fmt.Printf("Name: %s, Age: %d\n", name, age)
	// fmt.Printf("Binary: %b, Hex: %X\n", age, age)

	// // Formatting Func
	// s := fmt.Sprint("Hello", "World!", 123, 456)
	// fmt.Println(s)

	// s = fmt.Sprintln("Hello", "World!", 123, 456)
	// fmt.Println(s)

	// Scanning Func
	var name string
	var age int

	fmt.Print("Enter ur name and age:")
	// fmt.Scan(&name, &age)
	// fmt.Scanln(&name, &age)
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Error Formatting Func

	err := checkAge(age)
	if err != nil {
		fmt.Println("Error:", err)
	}

}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive", age)
	}
	return nil
}
