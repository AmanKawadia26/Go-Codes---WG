package main

import "fmt"

func square(n int) int {
	return n * n
}

func printSquare(f func(int) int, i int) string {
	return fmt.Sprint("The square of ", i, " is ", f(i))
}

func main() {
	a := printSquare(square, 12)
	fmt.Println(a)
}
