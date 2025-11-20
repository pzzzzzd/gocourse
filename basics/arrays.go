<<<<<<< HEAD
package main

import "fmt"

func main() {
	// var arrayName [size]elementType

	// var numbers [5]int
	// fmt.Println(numbers)

	// numbers[4] = 20
	// fmt.Println(numbers)

	// numbers[0] = 9
	// fmt.Println(numbers)

	// fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	// fmt.Println("Fruits array:", fruits)

	// fmt.Println("Third element:", fruits[2])

	// originalArray := [3]int{1, 2, 3}
	// copiedArray := originalArray
	// copiedArray[0] = 10
	// fmt.Println("Originl array:", originalArray)
	// fmt.Println("Copird array:", copiedArray)

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println("Element at index", i, ":", numbers[i])
	// }

	// // можно заменить index и value на i и v
	// // можно заменить не исп эл на _
	// for i, v := range numbers {
	// 	fmt.Printf("Index: %d, Value %d\n", i, v)
	// }

	// // можно заменить не исп эл на _
	// // a, b := someFunction()
	// // fmt.Println(a)
	// // fmt.Println(b)

	// a, _ := someFunction()
	// fmt.Println(a)

	// array1 := [3]int{1, 2, 3}
	// array2 := [3]int{10, 2, 3}
	// fmt.Println("Array 1 is equal to Array 2:", array1 == array2)

	// var matrix [3][3]int = [3][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	// fmt.Println(matrix)

	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int

	copiedArray = &originalArray
	copiedArray[0] = 100

	fmt.Println("Originl array:", originalArray)
	// fmt.Println("Copird array:", *copiedArray)

}

func someFunction() (int, int) {
	return 1, 2
}
=======
package main

import "fmt"

func main() {
	// var arrayName [size]elementType

	// var numbers [5]int
	// fmt.Println(numbers)

	// numbers[4] = 20
	// fmt.Println(numbers)

	// numbers[0] = 9
	// fmt.Println(numbers)

	// fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	// fmt.Println("Fruits array:", fruits)

	// fmt.Println("Third element:", fruits[2])

	// originalArray := [3]int{1, 2, 3}
	// copiedArray := originalArray
	// copiedArray[0] = 10
	// fmt.Println("Originl array:", originalArray)
	// fmt.Println("Copird array:", copiedArray)

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println("Element at index", i, ":", numbers[i])
	// }

	// // можно заменить index и value на i и v
	// // можно заменить не исп эл на _
	// for i, v := range numbers {
	// 	fmt.Printf("Index: %d, Value %d\n", i, v)
	// }

	// // можно заменить не исп эл на _
	// // a, b := someFunction()
	// // fmt.Println(a)
	// // fmt.Println(b)

	// a, _ := someFunction()
	// fmt.Println(a)

	// array1 := [3]int{1, 2, 3}
	// array2 := [3]int{10, 2, 3}
	// fmt.Println("Array 1 is equal to Array 2:", array1 == array2)

	// var matrix [3][3]int = [3][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	// fmt.Println(matrix)

	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int

	copiedArray = &originalArray
	copiedArray[0] = 100

	fmt.Println("Originl array:", originalArray)
	// fmt.Println("Copird array:", *copiedArray)

}

func someFunction() (int, int) {
	return 1, 2
}
>>>>>>> fb36f3b (wls_push)
