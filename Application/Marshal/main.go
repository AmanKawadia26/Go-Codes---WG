package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

type Internship struct {
	Batch    int
	Language string
	Name     string
}

func main() {
	var a int = 8
	bs, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))

	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}
	people := []person{p1, p2}
	fmt.Println(people)
	bs2, err2 := json.Marshal(people)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(string(bs2))

	var b bool
	fmt.Println(b)

	var a1 int

	err = json.Unmarshal(bs, &a1)
	if err != nil {
		panic(err)
	}
	fmt.Println(a1)

	var people1 []person
	err = json.Unmarshal(bs2, &people1)
	if err != nil {
		panic(err)
	}
	fmt.Println(people1)

	aman := Internship{
		Batch:    4,
		Language: "Go",
		Name:     "AlphaGo",
	}
	bs3, err3 := json.Marshal(aman)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println(string(bs3))
}
