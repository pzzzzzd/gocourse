package main

import (
	"fmt"
	"time"
)

func main() {
	var err error

	fmt.Println("Beginning program")
	go sayHello()
	fmt.Println("After sayHello func")

	go func() {
		err = doWork()
	}()

	// err = doWork()
	go printNumbers()
	go printLetters()

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Work completed")
	}

	time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Work completed")
	}
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println(i, ":", time.Now())
		time.Sleep(100 * time.Microsecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println(string(letter), ":", time.Now())
		time.Sleep(200 * time.Microsecond)
	}
}

func doWork() error {
	time.Sleep(1 * time.Second)
	return fmt.Errorf("an error occured in dowork")
}
