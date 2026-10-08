package main

import (
	"crypto/tls"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	mw "restapi/internal/api/middlewares"
	"restapi/internal/api/router"
	"restapi/pkg/utils"

	"github.com/joho/godotenv"
)

//go:embed all:.env
var envFile embed.FS

func loadENVFromEmbeddedFile() {
	content, err := envFile.ReadFile(".env")
	if err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}

	tempfile, err := os.CreateTemp("", ".env")
	if err != nil {
		log.Fatalf("Error creating temp .env file: %v", err)
	}
	defer os.Remove(tempfile.Name())

	_, err = tempfile.Write(content)
	if err != nil {
		log.Fatalf("Error writing to temp .env file: %v", err)
	}

	err = tempfile.Close()
	if err != nil {
		log.Fatalf("Error closing temp file: %v", err)
	}

	err = godotenv.Load(tempfile.Name())
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

}

func main() {

	// err := godotenv.Load()
	// if err != nil {
	// 	return
	// }

	loadENVFromEmbeddedFile()

	// fmt.Println("CERT_FILE:", os.Getenv("CERT_FILE"))

	port := os.Getenv("API_PORT")

	// cert := "cert.pem"
	// key := "key.pem"

	cert := os.Getenv("CERT_FILE")
	key := os.Getenv("KEY_FILE")

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS10,
	}

	// rl := mw.NewRateLimiter(5, time.Minute)

	hppOptions := mw.HPPOptions{
		CheckQuery:              true,
		CheckBody:               true,
		CheckBodyForContentType: "application/x-www-form-urlencoded",
		Whitelist:               []string{"sortBy", "sortOrder", "name", "age"},
	}

	router := router.MainRouter()
	jwtMiddleware := mw.MiddlewaresExcludePaths(mw.JWTMiddleware, "/execs/login", "/execs/forgotpassword", "/execs/resetpassword/reset")
	// secureMux := utils.ApplyMiddlewares(router, mw.SecurityHeaders, mw.Compression, jwtMiddleware, mw.Hpp(hppOptions),
	// mw.XSSMiddleware, mw.ResponseTimeMiddlewares, rl.Middlewares, mw.Cors)
	secureMux := utils.ApplyMiddlewares(router, mw.SecurityHeaders, mw.Compression, jwtMiddleware, mw.Hpp(hppOptions),
		mw.XSSMiddleware, mw.ResponseTimeMiddlewares, mw.Cors)
	// secureMux := jwtMiddleware(mw.SecurityHeaders(router))
	// secureMux := mw.XSSMiddleware(router)
	// secureMux := mw.SecurityHeaders(router)

	server := &http.Server{
		Addr: port,
		// Handler: mux,
		// Handler: middlewares.Cors(mux),
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server port", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error:", err)
	}

}
