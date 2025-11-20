package main

import "fmt"

func main() {
	// panic(interface{})

	//Example of a value input
	process(10)

	//Example of an invalue input
	process(-3)
}

func process(intput int) {

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if intput < 0 {
		fmt.Println("Before Panic")
		panic("input mast be a non-negativ number")
		// fmt.Println("After Panic")
		// defer fmt.Println("Deferred 3")
	}
	fmt.Println("Processing input:", intput)
}
