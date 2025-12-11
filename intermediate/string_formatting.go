package main

import "fmt"

func main() {

	num := 42
	fmt.Printf("%05d\n", num)
	num1 := 42232222222
	fmt.Printf("%05d\n", num1) // min len

	msg := "Hello"
	fmt.Printf("|%10s|\n", msg)
	fmt.Printf("|%-10s|\n", msg)

	msg2 := `Hello \nWorld`
	fmt.Println(msg2)

}
