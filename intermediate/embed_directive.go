package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed example.txt
var content string

//go:embed basics
var basicFolder embed.FS

func main() {

	fmt.Println("Embedded content:", content)
	content, err := basicFolder.ReadFile("basics/hello.go")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Embedded file content:", string(content))

	err = fs.WalkDir(basicFolder, "basics", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		fmt.Println(path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}
