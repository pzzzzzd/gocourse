package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		log.Println("Error in inializing Zap logger:", err)
	}

	defer logger.Sync()

	logger.Info("This is an info msg")

	logger.Info("User logged in", zap.String("username", "John"), zap.String("method", "GET"))

}
