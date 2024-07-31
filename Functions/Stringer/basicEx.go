package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

func changeValues(p *person) {
	p.age = p.age + 1
	//return p
}

func main() {
	p := person{"Bob", 20}
	fmt.Println(p)
	changeValues(&p)
	fmt.Println(p)
}
