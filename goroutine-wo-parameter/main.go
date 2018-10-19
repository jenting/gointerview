package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	runtime.GOMAXPROCS(1)

	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			// i will always be 10.
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}

	wg.Wait()

	// output
	// 10
	// 10
	// 10
	// 10
	// 10
	// 10
	// 10
	// 10
	// 10
	// 10
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
}
