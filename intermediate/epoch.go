package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("unixTime Time:", unixTime)
	t := time.Unix(unixTime, 0)
	fmt.Println(t)
	fmt.Println("Time:", t.Format("02-01-2006")) // 02 d | 01 m | 2006 y
}
