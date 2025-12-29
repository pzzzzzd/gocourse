package main

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// tempFileName := "tempFile"

	// tempFile, err := os.CreateTemp("", tempFileName)
	// checkError(err)

	// fmt.Println("Temporary file created:", tempFile.Name())

	// defer tempFile.Close()
	// defer os.Remove(tempFile.Name())

	tempDir, err := os.MkdirTemp("", "GoCourseTempDir")
	checkError(err)

	defer os.Remove(tempDir)

	fmt.Println("Temporary dir created:", tempDir)

}
