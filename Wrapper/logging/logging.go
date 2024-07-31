package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

func writeLog(s fmt.Stringer) {
	log.Println(s.String())
}

func main() {
	b := book{
		title: "West With The Night",
	}

	var c count = 42

	writeLog(b)
	writeLog(c)

}
