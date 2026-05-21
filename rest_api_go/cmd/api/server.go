package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	port := ":3011"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello Root Route")
		w.Write([]byte("Hello Root Route"))
		fmt.Println("Hello Root Route")
	})

	http.HandleFunc("/teaschers", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.Method)
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Hello GET method on Teaschers Route"))
			fmt.Println("Hello GET method on Teaschers Route")
		case http.MethodPost:
			w.Write([]byte("Hello POST method on Teaschers Route"))
			fmt.Println("Hello POST method on Teaschers Route")
		case http.MethodPut:
			w.Write([]byte("Hello PUT method on Teaschers Route"))
			fmt.Println("Hello PUT method on Teaschers Route")
		case http.MethodPatch:
			w.Write([]byte("Hello PATCH method on Teaschers Route"))
			fmt.Println("Hello PATCH method on Teaschers Route")
		case http.MethodDelete:
			w.Write([]byte("Hello DELETE method on Teaschers Route"))
			fmt.Println("Hello PATCH method on Teaschers Route")
		}

		// if r.Method == http.MethodGet {
		// 	w.Write([]byte("Hello GET method on Teaschers Route"))
		// 	fmt.Println("Hello GET method on Teaschers Route")
		// 	return
		// }
		w.Write([]byte("Hello Teaschers Route"))
		fmt.Println("Hello Teaschers Route")

	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Students Route"))
		fmt.Println("Hello Students Route")
	})

	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Execs Route"))
		fmt.Println("Hello Execs Route")
	})

	fmt.Println("Server port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error:", err)
	}

}
