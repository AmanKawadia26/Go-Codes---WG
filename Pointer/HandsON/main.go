package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

type youngin interface {
	young()
	Age()
}

type person struct {
	first string
	last  string
	age   int
}

func (p person) young() {
	fmt.Println("This is the young function with value semantic")
}

func (p person) Age() {
	fmt.Println("This is the age function with value semantic")
}

func main() {

	//Hands-On-Exercise-1
	//Creating a variable and printing its address
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	//Hands-On-Exercise-2
	//Dereference an address
	fmt.Printf("%v   %T\n", a, a)
	fmt.Printf("%v   %T\n", b, b)
	fmt.Printf("%v   %T\n", c, c)
	fmt.Printf("%v   %T\n", d, d)

	//Hands-On-Exercise-3
	//Method Sets
	p := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Printf("%#v\n", p)
	fmt.Printf("%v\n", &p)
	fmt.Printf("%+v\n", &p)

	p.young()
	p.Age()

	p1 := &person{
		first: "Aman",
		last:  "Kawadia",
		age:   20,
	}

	fmt.Printf("%#v\n", p)
	fmt.Printf("%v\n", &p)
	fmt.Printf("%+v\n", &p)

	p1.young()
	p1.Age()

	//Just-For-Curiosity
	const f = 9
	fmt.Println(f)

	//Here it will give a compilation error because "constant and literals" are not addressable so their address cannot be printed
	//Neither pointer semantic can be used in that case
	//fmt.Println(&f)

}
