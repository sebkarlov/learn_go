// this sample program demonstrates how to create goroutines and how the scheduler behaves
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// allocate 1 logical processor for the scheduler to use
	runtime.GOMAXPROCS(1)

	// wg is used to wait for the program to finish
	// Add a count of two, one for each goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start goroutines")

	// declare an anon func and create a goroutine
	go func() {
		// schedule the call to Done to tell main we are done
		defer wg.Done()

		// display tha alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// declare an anon func and create a goroutine
	go func() {
		// schedule the call to Done to tell main we are done
		defer wg.Done()

		// display tha alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// Wait for the gourotines to finish
	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("\nTerminating program")
}
