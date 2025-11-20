<<<<<<< HEAD
package main

import "fmt"

func main() {
	// // func <name>(parameters list) returnType {
	// // code block
	// // return value
	// // }

	// // sum := add(1, 2)
	// fmt.Println(add(2, 3))

	// greet := func() {
	// 	fmt.Println("Hello Anonymous Functions")
	// }

	// greet()

	// operation := add

	// result := operation(3, 5)

	// fmt.Println(result)

	// Passsing a function as an argument
	result := applyOreration(5, 3, add)
	fmt.Println("5 + 3 =", result)

	// Returning and using function
	multiplyBy2 := createMultilier(2)
	fmt.Println("6 * 2 =", multiplyBy2(6))
}

func add(a, b int) int {
	return a + b
}

// Function that takes a function as an argument
func applyOreration(x int, y int, operatin func(int, int) int) int {
	return operatin(x, y)
}

// Function than return a function
func createMultilier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
=======
package main

import "fmt"

func main() {
	// // func <name>(parameters list) returnType {
	// // code block
	// // return value
	// // }

	// // sum := add(1, 2)
	// fmt.Println(add(2, 3))

	// greet := func() {
	// 	fmt.Println("Hello Anonymous Functions")
	// }

	// greet()

	// operation := add

	// result := operation(3, 5)

	// fmt.Println(result)

	// Passsing a function as an argument
	result := applyOreration(5, 3, add)
	fmt.Println("5 + 3 =", result)

	// Returning and using function
	multiplyBy2 := createMultilier(2)
	fmt.Println("6 * 2 =", multiplyBy2(6))
}

func add(a, b int) int {
	return a + b
}

// Function that takes a function as an argument
func applyOreration(x int, y int, operatin func(int, int) int) int {
	return operatin(x, y)
}

// Function than return a function
func createMultilier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
>>>>>>> fb36f3b (wls_push)
