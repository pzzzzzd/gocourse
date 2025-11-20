package main

import "fmt"

func main() {

	process(10)

}

func process(i int) {
	defer fmt.Println("Deffered i value:", i)
	defer fmt.Println("Firtst defferred statement executed")
	defer fmt.Println("Second defferred statement executed")
	defer fmt.Println("Third defferred statement executed")
	i++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of i:", i)
}
