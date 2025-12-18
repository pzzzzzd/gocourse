package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// reader := bufio.NewReader(strings.NewReader("Hello, bufio packageee\n????"))

	// data := make([]byte, 20)
	// n, err := reader.Read(data)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	// line, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Read string:", line)

	writer := bufio.NewWriter(os.Stdout)

	data := []byte("Hello, bufio package\n")
	n, err := writer.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	str := "This is a string"
	writer.WriteString(str)
	n, err = writer.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}
