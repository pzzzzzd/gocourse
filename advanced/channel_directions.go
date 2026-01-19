package main

import "fmt"

func main() {

	ch := make(chan int)
	// go func(ch chan<- int) {
	// 	for i := range 5 {
	// 		ch <- i
	// 	}
	// 	close(ch)
	// }(ch)

	producer(ch)

	// for value := range ch {
	// 	fmt.Println("Received:", value)
	// }

	receiveData(ch)
}

// receive only channel

func producer(ch chan<- int) {
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
}

func receiveData(ch <-chan int) {
	for value := range ch {
		fmt.Println("Received:", value)
	}
}
