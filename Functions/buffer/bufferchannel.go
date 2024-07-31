package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3) // Buffered channel with a capacity of 3

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Printf("Sent %d\n", i)
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Printf("Received %d\n", v)
	}
}
