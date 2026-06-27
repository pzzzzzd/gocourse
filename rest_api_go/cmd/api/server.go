package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
	"restapi/internal/api/router"
)

func main() {

	port := ":3011"

	cert := "cert.pem"
	key := "key.pem"

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// rl := mw.NewRateLimiter(5, time.Minute)

	// hppOptions := mw.HPPOptions{
	// 	CheckQuery:              true,
	// 	CheckBody:               true,
	// 	CheckBodyForContentType: "application/x-www-form-urlencoded",
	// 	Whitelist:               []string{"sortBy", "sortOrder", "name", "age"},
	// }

	// secureMux := utils.ApplyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddlewares, rl.Middlewares, mw.Cors)
	router := router.Router()
	secureMux := mw.SecurityHeaders(router)

	server := &http.Server{
		Addr: port,
		// Handler: mux,
		// Handler: middlewares.Cors(mux),
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error:", err)
	}

}
