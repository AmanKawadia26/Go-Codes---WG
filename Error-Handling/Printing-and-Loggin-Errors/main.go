package main

import (
	"fmt"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("test.log")
	if err != nil {
		//fmt.Println("Error Happened", err)   //Only Prints the error with modification feature and continues the execution
		//log.Println("Error Happened", err)   //Prints the error with time stamp and allow modification and continues the execution
		//fmt.Fatalln("Error Happened", err)   //Kills the execution when the error is passed
		//log.Panicln("Error Happened", err)   //Stops the execution of the current go routine but not others, but it can be recovered using "recover"
		panic(err) //Stops the execution of the current go routine but not others
	}

	fmt.Println("Program Exiting")
}

func foo() {
	fmt.Println("It is a deferred function.")
}
