package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type person struct {
	name string
}

func (p person) writeOut(w io.Writer) {
	_, err := w.Write([]byte(p.name))
	if err != nil {
		fmt.Println("Error writing to writer:", err)
	}
}

func main() {

	p := person{
		"Aman",
	}

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	var b bytes.Buffer
	p.writeOut(f)
	p.writeOut(&b)
	fmt.Println(p)
	fmt.Println(b.String())

}
