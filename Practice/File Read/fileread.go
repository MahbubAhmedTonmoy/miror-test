package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome file read write")
	writeFile()
	readFile("./file.text")
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("text data in file: \n", string(databyte))
}
func writeFile() {
	content := "hello mahbub good to khow you start learn go lang"

	file, err := os.Create("./file.text")
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is :", length)
	defer file.Close()
}
