package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("readFile %s: %v", filename, err)
	}
	return data, nil
}
func main() {
	data, err := readFile("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Printf("Data: %s\n", data)
}
