package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(powinator)

	x := powinator(2)

	//If printing a function identifier, then it will give the address of the function
	fmt.Println(x)

	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

func powinator(a float64) func() float64 {
	var c float64
	return func() float64 {
		c++
		return math.Pow(a, c)
	}
}
