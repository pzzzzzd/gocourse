package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(1 * time.Second)
	timer2 := time.NewTimer(2 * time.Second)

	select {
	case <-timer1.C:
		fmt.Println("Timer1 expired")
	case <-timer2.C:
		fmt.Println("Timer2 expired")
	}

}

// func main() {
// 	timer := time.NewTimer(2 * time.Second)
// 	go func() {
// 		<-timer.C
// 		fmt.Println("Delayed operation executed")
// 	}()
// 	fmt.Println("Wating")
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("End")
// }

// ---- Time out ----
// func longRunningOperation() {
// 	for i := range 20 {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	timeout := time.After(2 * time.Second)
// 	done := make(chan bool)
// 	go func() {
// 		longRunningOperation()
// 		done <- true
// 	}()
// 	select {
// 	case <-timeout:
// 		fmt.Println("time out")
// 	case <-done:
// 		fmt.Println("Operation complited")
// 	}
// }

// ---- Timer ----
// func main() {
// 	time.Sleep(time.Second)
// 	fmt.Println("Starting")
// 	timer := time.NewTimer(2 * time.Second)
// 	fmt.Println("Wating for timer.C")
// 	stopped := timer.Stop()
// 	if stopped {
// 		fmt.Println("Timer stopped")
// 	}
// 	timer.Reset(time.Second)
// 	fmt.Println("Timer reset")
// 	<-timer.C
// 	fmt.Println("Timer expired")
// }
