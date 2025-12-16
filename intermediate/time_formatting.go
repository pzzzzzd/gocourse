package main

import (
	"fmt"
	"time"
)

func main() {

	// Mon Jan 2 15:04:05 MST 2006
	layout := "2006-01-02T15:04:05Z07:00" // base
	str := "2025-12-16T15:17:29Z"

	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(t)

	str1 := "Dec 16, 2025 03:23 PM"
	layout1 := "Jan 02, 2006 03:04 PM" // base
	t1, err := time.Parse(layout1, str1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(t1)

}
