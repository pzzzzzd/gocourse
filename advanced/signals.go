package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	pid := os.Getpid()
	fmt.Println("Process ID:", pid)
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGHUP)

	go func() {
		sig := <-sigs
		fmt.Println("Received signal:", sig)
		done <- true
	}()

	go func() {

		for {
			select {
			case <-done:
				fmt.Println("Stopping work due to signal.")
				// os.Exit(0)
				return
			default:
				fmt.Println("Working...")
				time.Sleep(time.Second)
			}
		}
		// sig := <-sigs
		// for sig := range sigs {
		// 	switch sig {
		// 	case syscall.SIGINT:
		// 		fmt.Println("Received SIGINT (Interrupt)")
		// 	// case syscall.SIGTERM:
		// 	// 	fmt.Println("Received SIGTERM (Terminate)")
		// 	case syscall.SIGHUP:
		// 		fmt.Println("Received SIGHUP (Hangup)")
		// 	case syscall.SIGUSR1:
		// 		fmt.Println("Received SIGNUSR1 (User defined Signal 1)")
		// 		fmt.Println("User define function is executed")
		// 		// continue
		// 	}
		// 	// fmt.Println("Graceful exit.")
		// 	// os.Exit(0)
		// }
	}()

	// fmt.Println("Working...")
	for {
		time.Sleep(time.Second)
	}
}
