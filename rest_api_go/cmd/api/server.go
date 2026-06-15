package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
	"strings"
)

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello Root Route")
	w.Write([]byte("Hello Root Route"))
	fmt.Println("Hello Root Route")
}

func treacherHendler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method)
	switch r.Method {
	case http.MethodGet:
		// // request details
		// fmt.Println("Body:", r.Body)
		// fmt.Println("Form:", r.Form)
		// fmt.Println("Header:", r.Header)
		// fmt.Println("Context:", r.Context())
		// fmt.Println("ContextLength:", r.ContentLength)
		// fmt.Println("Host:", r.Host)
		// fmt.Println("Method:", r.Method)
		// fmt.Println("Protocol:", r.Proto)
		// fmt.Println("Remote Addr:", r.RemoteAddr)
		// fmt.Println("Request URI:", r.RequestURI)
		// fmt.Println("TSL:", r.TLS)
		// fmt.Println("Trailer:", r.Trailer)
		// fmt.Println("Transfer Encoding:", r.TransferEncoding)
		// fmt.Println("URL:", r.URL)
		// fmt.Println("User Agent:", r.UserAgent())
		// fmt.Println("Port:", r.URL.Port())
		// fmt.Println("Schame:", r.URL.Scheme)

		fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		userID := strings.TrimSuffix(path, "/")

		fmt.Println("ID:", userID)

		fmt.Println("Query params:", r.URL.Query())
		queryParams := r.URL.Query()
		sortby := queryParams.Get("sortby")
		key := queryParams.Get("key")
		sortorder := queryParams.Get("sortorder")

		if sortorder == "" {
			sortorder = "DESC"
		}

		fmt.Printf("Sortby: %v, Sortorder: %v, Key: %v", sortby, sortorder, key)

		w.Write([]byte("Hello GET method on Teaschers Route"))
		// fmt.Println("Hello GET method on Teaschers Route")
	case http.MethodPost:

		// // parse x-www-...
		// err := r.ParseForm()
		// if err != nil {
		// 	http.Error(w, "Error parsing form", http.StatusBadRequest)
		// 	return
		// }

		// fmt.Println("Form:", r.Form)

		// // response data
		// response := make(map[string]interface{})
		// for key, value := range r.Form {
		// 	response[key] = value[0]
		// }

		// fmt.Println("Response Map:", response)

		// // parse RAW
		// body, err := io.ReadAll(r.Body)
		// if err != nil {
		// 	return
		// }
		// defer r.Body.Close()

		// fmt.Println("RAW body:", body)
		// fmt.Println("RAW body:", string(body))

		// var userInstance user
		// err = json.Unmarshal(body, &userInstance)
		// if err != nil {
		// 	return
		// }

		// fmt.Println(userInstance)
		// fmt.Println("Name:", userInstance.Name)

		// // response data
		// response1 := make(map[string]interface{})
		// for key, value := range r.Form {
		// 	response1[key] = value[0]
		// }

		// err = json.Unmarshal(body, &response1)
		// if err != nil {
		// 	return
		// }

		// fmt.Println(response1)
		// // fmt.Println("Name:", response1)

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
	// w.Write([]byte("Hello Teaschers Route"))
	// fmt.Println("Hello Teaschers Route")

}

func studentHeandler(w http.ResponseWriter, r *http.Request) {

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

	// w.Write([]byte("Hello Students Route"))
	// fmt.Println("Hello Students Route")
}

func execHeandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on Execs Route"))
		fmt.Println("Hello GET method on Execs Route")
	case http.MethodPost:
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

	// w.Write([]byte("Hello Execs Route"))
	// fmt.Println("Hello Execs Route")
}

func main() {

	port := ":3011"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/teachers/", treacherHendler)

	mux.HandleFunc("/students/", studentHeandler)

	mux.HandleFunc("/execs/", execHeandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr: port,
		// Handler: mux,
		// Handler: middlewares.Cors(mux),
		Handler:   mw.Compression(mw.ResponseTimeMiddlewares(mw.SecurityHeaders(mw.Cors(mux)))),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error:", err)
	}

}
