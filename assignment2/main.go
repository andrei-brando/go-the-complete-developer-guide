package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args

	saveToFile(args[1], "blablabla")
	content := readFile(args[1])

	fmt.Println(content)
}

func saveToFile(filename string, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0666)
}

func readFile(filename string) string {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return string(bs)
}
