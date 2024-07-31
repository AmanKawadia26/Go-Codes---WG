package main

import "fmt"

func main() {
	a := func(a, b int) int {
		return a + b
	}
	b := a(5, 6)
	fmt.Println(b)
}
