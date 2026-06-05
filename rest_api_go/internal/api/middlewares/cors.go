package middlewares

import (
	"fmt"
	"net/http"
)

var allowedOrigins = []string{
	"https://my-origin-url.com",
	"https://www.myfrontend.com",
	"https://localhost:3011",
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		fmt.Println(origin)

		if isOriginAllowed(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			http.Error(w, "Not allowed by CORS", http.StatusForbidden)
			return
		}

		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Autorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "1234")

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)

	})
}

func isOriginAllowed(origin string) bool {
	for _, allowedOrigins := range allowedOrigins {
		if origin == allowedOrigins {
			return true
		}
	}
	return false
}
