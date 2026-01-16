package main

import (
	"fmt"
	"time"
)

// func main() {

// 	done := make(chan int)

// 	go func() {

// 		fmt.Println("Working...")
// 		time.Sleep(2 * time.Second)
// 		done <- 0

// 	}()
// 	<-done
// 	fmt.Println("Finished")

// }

// ==========================================================================

// func main() {

// 	ch := make(chan int)

// 	go func() {
// 		ch <- 9
// 		time.Sleep(1 * time.Second)
// 		fmt.Println("Sent value")

// 	}()

// 	value := <-ch
// 	fmt.Println(value)
// 	time.Sleep(2 * time.Second)

// }

// ==========================================================================

// func main() {

// 	numGoroutines := 3
// 	done := make(chan int, 3)

// 	for i := range numGoroutines {
// 		// time.Sleep(1 * time.Second)
// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working... \n", id)
// 			time.Sleep(1 * time.Second)
// 			done <- id
// 		}(i)
// 	}
// 	for range numGoroutines {
// 		<-done // wait for each goroutines to finish
// 	}

// 	fmt.Println("All goroutines are complited")

// }

// ==========================================================================

func main() {

	data := make(chan string)

	go func() {
		for i := range 5 {
			data <- "hello " + string('0'+i)
			time.Sleep(100 * time.Millisecond)
		}
		close(data)
	}()

	for value := range data {
		fmt.Printf("Received value: %s | %s\n", value, time.Now())
	}

}
