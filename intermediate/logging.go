package main

import (
	"log"
	"os"
)

func main() {

	log.Println("This is a log msg")

	log.SetPrefix("INFO: ")
	log.Println("This is an info msg")

	// log flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("This is a log msg with only date, time, file name")

	infoLogger.Println("This is an info msg")
	warnLogger.Println("This is a warning msg")
	errLorogger.Println("This is an error msg")

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file:", err)
	}
	defer file.Close()

	infoLogger1 := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger1 := log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errLorogger1 := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	debugLogger := log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	debugLogger.Println("This is a debug msg")
	infoLogger1.Println("This is an info msg")
	warnLogger1.Println("This is a warning msg")
	errLorogger1.Println("This is an error msg")

}

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errLorogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)
