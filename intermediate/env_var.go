package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("User env var:", user)
	fmt.Println("Home env var:", home)

	err := os.Setenv("FRUIT", "APPLE")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("FRUIT env var:", os.Getenv("FRUIT"))

	// for _, e := range os.Environ() {
	// 	kvpair := strings.SplitN(e, "=", 2)
	// 	fmt.Println(kvpair[0])
	// }

	err = os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Unset env var done on key FRUIT")
	fmt.Println("FRUIT env var:", os.Getenv("Fruit"))

	str := "a=b=c=d"
	// "a=b=c=d"
	// n = 1 return "a=b=c=d"
	// n = 2 return "a" and "b=c=d"
	// n = 3 return "a" and "b" and "c=d"
	// n = 4 return "a" and "b" and "c" and "d"
	fmt.Println(strings.SplitN(str, "=", -1))
	fmt.Println(strings.SplitN(str, "=", 0))
	fmt.Println(strings.SplitN(str, "=", 1))
	fmt.Println(strings.SplitN(str, "=", 2))
	fmt.Println(strings.SplitN(str, "=", 3))
	fmt.Println(strings.SplitN(str, "=", 4))
}
