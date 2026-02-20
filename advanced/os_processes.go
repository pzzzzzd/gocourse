package main

import (
	"fmt"
	"os/exec"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}

}

func main() {

	cmd := exec.Command("ls", "-l")
	output, err := cmd.CombinedOutput()
	checkErr(err)
	fmt.Println("Output:", string(output))

	// pr, pw := io.Pipe()
	// cmd := exec.Command("grep", "foo")
	// cmd.Stdin = pr
	// go func() {
	// 	defer pw.Close()
	// 	pw.Write([]byte("foo3333\ndddddd\ndddfdsf"))
	// }()
	// output, err := cmd.Output()
	// checkErr(err)
	// fmt.Println("Output:", string(output))

	// cmd := exec.Command("printenv", "SHELL")
	// output, err := cmd.Output()
	// checkErr(err)
	// fmt.Println("Output:", string(output))

	// cmd := exec.Command("sleep", "44")
	// err := cmd.Start()
	// checkErr(err)
	// time.Sleep(2 * time.Second)
	// err = cmd.Process.Kill()
	// checkErr(err)
	// fmt.Println("Process is killed")
	// err = cmd.Wait()
	// checkErr(err)
	// fmt.Println("Process is complited")

	// cmd := exec.Command("grep", "foo")
	// cmd.Stdin = strings.NewReader("foo idddddddddddddd\nfff\nddd")
	// output, err := cmd.Output()
	// checkErr(err)
	// fmt.Println("Output:", string(output))

	// cmd := exec.Command("echo", "hello world")
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Output:", string(output))

}
