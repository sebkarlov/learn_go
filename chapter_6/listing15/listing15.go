// this sample program demonstrates how to use the atomic package to provide safe access to numeric types
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// shutdown is a flag to alert running gouroutines to shutdown
	shutdown int64

	// wg is used to wait for the program to finish
	wg sync.WaitGroup
)

func main() {
	// Add a count of two, one for each goroutine
	wg.Add(2)

	// Create two goroutines
	go doWork("A")
	go doWork("B")

	// give the goroutines time to run
	time.Sleep(1 * time.Second)

	// safely flag it is time to shutdown
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	// Wait for the goroutines to finish
	wg.Wait()
}

// doWork simulates a goroutine performing work and checking the Shutdown flag to terminate early
func doWork(name string) {
	// shedule the call to Done tell main we are done
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		// do we need to shutdown
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
