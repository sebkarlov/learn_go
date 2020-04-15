// this sample program demonstrates how to use a mutex to define critical sections of code that need syncronous access
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// conter is a variable incremented bu all goroutines
	counter int

	// wg is used to wait for the program to finish
	wg sync.WaitGroup

	// mutex is used to define a critical section of code
	mutex sync.Mutex
)

func main() {
	// Add a count of two, one for each goroutine
	wg.Add(2)

	// Create two goroutines
	go incCounter(1)
	go incCounter(2)

	// Wait for the goroutines to finish
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

// incCounter increments the package level counter variable using the Mutex to syncronize and provide safe access
func incCounter(id int) {
	//schedule the call to Done to tell main we are done
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// only allow one goroutine throught this critical section at a time
		mutex.Lock()
		{
			// capture the value of Counter
			value := counter

			// yield the thread and be placed back in queue
			runtime.Gosched()

			// increment our local value of Counter
			value++

			// store the value back into Counter
			counter = value
		}
		mutex.Unlock()
		// release the lock and allow any waiting gouroutine through
	}
}
