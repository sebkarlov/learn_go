// this sample program demonstrates how the goroutine scheduler will tine slice goroutines on a single thread
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// wg is used to wait for the program to finish
var wg sync.WaitGroup

func main() {
	// allocate 1 logical processor for the scheduler to use
	runtime.GOMAXPROCS(1)

	// Add a count of two, one for each goroutine
	wg.Add(2)

	// create 2 goroutines
	fmt.Println("create goroutines")

	go printPrime("A")
	go printPrime("B")

	// Wait for the gourotines to finish
	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("\nTerminating program")
}

// printPrime displays prime numbers for the first 5000 numbers
func printPrime(prefix string) {
	// schedule the call to Done to tell main we are done
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("completed", prefix)
}
