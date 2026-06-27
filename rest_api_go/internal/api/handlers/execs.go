package handlers

import (
	"fmt"
	"net/http"
)

func ExecHeandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on Execs Route"))
		fmt.Println("Hello GET method on Execs Route")
	case http.MethodPost:
		fmt.Println("Query:", r.URL.Query())
		fmt.Println("name:", r.URL.Query().Get("name"))

		err := r.ParseForm()
		if err != nil {
			return
		}

		fmt.Println("Form from POST methods:", r.Form)

		w.Write([]byte("Hello POST method on Execs Route"))
		fmt.Println("Hello POST method on Execs Route")
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Execs Route"))
		fmt.Println("Hello PUT method on Execs Route")
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Execs Route"))
		fmt.Println("Hello PATCH method on Execs Route")
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Execs Route"))
		fmt.Println("Hello PATCH method on Execs Route")
	}
}
