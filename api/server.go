package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello server")
	})

	const port string = ":3011"

	fmt.Println("Port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error:", err)
	}

}
