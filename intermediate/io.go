package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	checkError(err)
	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	checkError(err)
}

func closeResource(c io.Closer) {
	err := c.Close()
	checkError(err)
}

func bufferExample() {
	var buf bytes.Buffer
	buf.WriteString("Hello Buffer")
	fmt.Println(buf.String())
}

func multiReaderExample() {
	r1 := strings.NewReader("Hello")
	r2 := strings.NewReader(" world")
	mr := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(mr)
	checkError(err)
	fmt.Println(buf.String())
}

func writeToFile(filepath string, data string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkError(err)
	defer closeResource(file)

	_, err = file.Write([]byte(data))
	checkError(err)

	// // Type(value)
	// writer := io.Writer(file)
	// _, err = writer.Write([]byte(data))
	// checkError(err)
}

func pipeExample() {
	pr, pw := io.Pipe()
	go func() {
		pw.Write([]byte("Hello Pipe"))
		pw.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)
	fmt.Println(buf.String())
}

func main() {

	fmt.Println("--- Read from Reader ---")
	readFromReader(strings.NewReader("Hello Reader"))

	fmt.Println("--- Write from Writer ---")
	var writer bytes.Buffer
	writeToWriter(&writer, "Hello writer")
	fmt.Println(writer.String())

	fmt.Println("--- Buffer Example ---")
	bufferExample()

	fmt.Println("--- Multi Reader Example ---")
	multiReaderExample()

	fmt.Println("--- Pipe Example ---")
	pipeExample()

	filepath := "io.txt"
	writeToFile(filepath, "hello file")

	resourse := &MyResource{name: "TestResource"}
	closeResource(resourse)

}

type MyResource struct {
	name string
}

func (m MyResource) Close() error {
	fmt.Println("Closing resource:", m.name)
	return nil
}
