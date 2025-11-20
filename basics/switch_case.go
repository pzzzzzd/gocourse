<<<<<<< HEAD
package main

import "fmt"

func main() {
	// switch expressiom{
	// case value1:
	// // если значение 1
	// fallthroungh есл хотим перейти в след кейс
	// case value2:
	// // если значение 2
	// case value3:
	// // если значение 3
	// default:
	// // если ничего не подходит
	// }

	// fruit := "pineapple"

	// switch fruit {
	// case "apple":
	// 	fmt.Println("It's an apple")
	// case "banana":
	// 	fmt.Println("It's a banana")
	// default:
	// 	fmt.Println("Unknown fruit")
	// }

	// day := "Mondey"

	// switch day {
	// case "Mondey", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("It's a weekday")
	// case "Sunday":
	// 	fmt.Println("It's a weekend")
	// default:
	// 	fmt.Println("Invalid day")
	// }

	// number := 15
	// switch {
	// case number < 10:
	// 	fmt.Println("Number is less then 10")
	// case number >= 10 && number < 20:
	// 	fmt.Println("Number is between 10 and 19")
	// default:
	// 	fmt.Println("Number is 20 or more")
	// }

	// num := 2
	// switch {
	// case num > 1:
	// 	fmt.Println("Greater than 1")
	// 	fallthrough
	// case num == 2:
	// 	fmt.Println("Number is 2")
	// default:
	// 	fmt.Println("Not 2")
	// }

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("It's float")
	case string:
		fmt.Println("It's string")
	default:
		fmt.Println("Unknown type")
	}
}
=======
package main

import "fmt"

func main() {
	// switch expressiom{
	// case value1:
	// // если значение 1
	// fallthroungh есл хотим перейти в след кейс
	// case value2:
	// // если значение 2
	// case value3:
	// // если значение 3
	// default:
	// // если ничего не подходит
	// }

	// fruit := "pineapple"

	// switch fruit {
	// case "apple":
	// 	fmt.Println("It's an apple")
	// case "banana":
	// 	fmt.Println("It's a banana")
	// default:
	// 	fmt.Println("Unknown fruit")
	// }

	// day := "Mondey"

	// switch day {
	// case "Mondey", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("It's a weekday")
	// case "Sunday":
	// 	fmt.Println("It's a weekend")
	// default:
	// 	fmt.Println("Invalid day")
	// }

	// number := 15
	// switch {
	// case number < 10:
	// 	fmt.Println("Number is less then 10")
	// case number >= 10 && number < 20:
	// 	fmt.Println("Number is between 10 and 19")
	// default:
	// 	fmt.Println("Number is 20 or more")
	// }

	// num := 2
	// switch {
	// case num > 1:
	// 	fmt.Println("Greater than 1")
	// 	fallthrough
	// case num == 2:
	// 	fmt.Println("Number is 2")
	// default:
	// 	fmt.Println("Not 2")
	// }

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("It's float")
	case string:
		fmt.Println("It's string")
	default:
		fmt.Println("Unknown type")
	}
}
>>>>>>> fb36f3b (wls_push)
