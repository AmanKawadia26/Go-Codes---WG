package main

import "fmt"

func get42() func() int {
	return func() int {
		return 42
	}
}

func main() {
	a := get42()
	fmt.Println(a())
}
