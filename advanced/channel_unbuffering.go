package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	// go func() {
	// 	// ch <- 1
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("2 second")
	// }()
	// go func() {
	// 	// ch <- 1
	// 	time.Sleep(3 * time.Second)
	// 	fmt.Println("3 second")
	// }()
	go func() {
		time.Sleep(3 * time.Second)
		// receiver := <-ch
		// fmt.Println(receiver)
		fmt.Println(<-ch)
		fmt.Println("3 second")
	}()
	ch <- 1

	fmt.Println("end")

}
