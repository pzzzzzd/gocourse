package router

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.RootHandler)

	mux.HandleFunc("/teachers/", handlers.TeacherHendler)

	mux.HandleFunc("/students/", handlers.StudentHeandler)

	mux.HandleFunc("/execs/", handlers.ExecHeandler)

	return mux
}
