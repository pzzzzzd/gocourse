package handlers

import (
	"fmt"
	"net/http"
)

func StudentHeandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on Students Route"))
		fmt.Println("Hello GET method on Students Route")
	case http.MethodPost:
		w.Write([]byte("Hello POST method on Students Route"))
		fmt.Println("Hello POST method on Students Route")
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Students Route"))
		fmt.Println("Hello PUT method on Students Route")
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Students Route"))
		fmt.Println("Hello PATCH method on Students Route")
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Students Route"))
		fmt.Println("Hello PATCH method on Students Route")
	}
}
