package main

import (
	"fmt"
	"time"
)

func main() {

	// ch := make(chan int)

	// non blocking receiver operation
	// select {
	// case msg := <-ch:
	// 	fmt.Println("Received:", msg)
	// default:
	// 	fmt.Println("No msg available")
	// }

	// non nlocking send operation
	// select {
	// case ch <- 1:
	// 	fmt.Println("Sent msg")
	// default:
	// 	fmt.Println("Channel is not ready to received")
	// }

	// non blocking operation in real time system
	data := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case d := <-data:
				fmt.Println("Data received:", d)
			case <-quit:
				fmt.Println("Stopping")
				return
			default:
				fmt.Println("Wating for data")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(time.Second)
	}

	quit <- true
}
