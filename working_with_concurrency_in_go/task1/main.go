package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string) {
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	msg = "Hello, world"
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		updateMessage("Hello, universe")
		printMessage()

		go func() {
			defer wg.Done()
			updateMessage("Hello, cosmos")
			printMessage()

			go func() {
				defer wg.Done()
				updateMessage("Hello, world")
				printMessage()
			}()
		}()
	}()

	wg.Wait()
}
