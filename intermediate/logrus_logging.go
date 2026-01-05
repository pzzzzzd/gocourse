package main

import "github.com/sirupsen/logrus"

func main() {

	log := logrus.New()

	// set log lvl
	log.SetLevel(logrus.InfoLevel)

	// set log format
	log.SetFormatter(&logrus.JSONFormatter{})

	// log example
	log.Info("This is an info msg")
	log.Warn("This is a warning msg")
	log.Error("This is an error msg")

	log.WithFields(logrus.Fields{
		"username": "John",
		"method":   "GET",
	}).Info("User logged in")

}
