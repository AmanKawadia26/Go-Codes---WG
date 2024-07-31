package main

import "fmt"

type dog struct {
	first string
}

type cat struct {
	first string
}

type animal interface {
	walk()
	run()
}

func (d *dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}

func (c cat) walk() {
	fmt.Println("My name is", c.first, "and I'm walking.")
}

func (d *dog) run() {
	d.first = "Rover"
	fmt.Println("My name is", d.first, "and I'm running.")
}

func (c cat) run() {
	c.first = "Rover"
	fmt.Println("My name is", c.first, "and I'm running.")
}

func main() {
	d1 := dog{"Henry"}
	d1.walk()
	d1.run()

	var b animal = &d1
	b.walk()
	b.run()

	d2 := &dog{"Padget"}
	fmt.Printf("%T, %v\n", d2, d2)
	d2.walk()
	d2.run()

	c := cat{
		first: "Tom",
	}
	var a animal = c
	a.walk()
	a.run()
	fmt.Println(c)
}
