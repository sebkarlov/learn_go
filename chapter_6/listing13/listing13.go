// this sample program demonstrates how to use the atomic package to provide safe access to numeric types
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// conter is a variable incremented bu all goroutines
	counter int64

	// wg is used to wait for the program to finish
	wg sync.WaitGroup
)

func main() {
	// Add a count of two, one for each goroutine
	wg.Add(2)

	// Create two goroutines
	go incCounter(1)
	go incCounter(2)

	// Wait for the goroutines to finish
	wg.Wait()

	// display the final value
	fmt.Println("Final Counter:", counter)
}

// incCounter increments the package level counter variable
func incCounter(id int) {
	//schedule the call to Done to tell main we are done
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// safely Add 1 to Counter
		atomic.AddInt64(&counter, 1)

		// yield the thread and be placed back in queue
		runtime.Gosched()
	}
}
