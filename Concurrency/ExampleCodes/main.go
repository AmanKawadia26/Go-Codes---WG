package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	//var mu sync.Mutex

	gs := 500

	counter := 0

	wg.Add(gs)

	start := time.Now()

	for i := 0; i < gs; i++ {
		go func() {
			var wg2 sync.WaitGroup
			//mu.Lock()
			gs2 := 50
			counter2 := 0
			wg2.Add(gs2)
			start2 := time.Now()
			//wg2.Add(1)
			for i := 0; i < gs2; i++ {
				go func() {
					v2 := counter2
					//runtime.Gosched()
					v2++
					counter2 = v2
					wg2.Done()
				}()
			}
			wg2.Wait()
			elapsed2 := time.Since(start2)
			fmt.Printf("Inner counter: %d, elapsed: %s\n", counter2, elapsed2)
			v := counter
			//runtime.Gosched()
			v++
			counter = v
			fmt.Println("Counter: ", counter)

			wg.Done()
			//mu.Unlock()
		}()
		fmt.Println("Go routines:", runtime.NumGoroutine())
	}
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Println("Time elapsed: ", elapsed)
}
