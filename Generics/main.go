package main

import "fmt"

func abc(a any) {
	fmt.Println(a)
}

func addT[T int | float64](a T, b T) T {
	fmt.Println(a)
	fmt.Println(b)
	return a + b
}

func main() {
	a := 5.3
	var b float64 = 23.2
	fmt.Println(addT(a, b))
	abc("Hello World")
}
