package main

import (
	"fmt"
	"time"
)

func main() {

	// make(chan Type, capacity)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	fmt.Println("Receiving from buffer")
	go func() {
		fmt.Println("2 second")
		time.Sleep(2 * time.Second)
		fmt.Println("Received:", <-ch)
	}()
	fmt.Println("Start")
	ch <- 3 // blocks because the buffer is full
	fmt.Println("End")
	// fmt.Println("Received:", <-ch)
	// fmt.Println("Received:", <-ch)

	// fmt.Println("Buffered Channel")

	// ===========================================================================

	// ch := make(chan int, 2)

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	ch <- 1
	// 	ch <- 2
	// }()

	// fmt.Println("Value:", <-ch)
	// fmt.Println("Value:", <-ch)
	// fmt.Println("End")
}
