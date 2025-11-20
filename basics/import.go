<<<<<<< HEAD
package main

import (
	"fmt"
	foo "net/http"
)

func main() {
	fmt.Println("Hello, Go Standart Library")

	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("HTTP Response Status", resp.Status)
}
=======
package main

import (
	"fmt"
	foo "net/http"
)

func main() {
	fmt.Println("Hello, Go Standart Library")

	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("HTTP Response Status", resp.Status)
}
>>>>>>> fb36f3b (wls_push)
