// sample programm to show how the program can access a value of an unexported identifier from another package
package main

import (
	"fmt"

	"github.com/goinaction/code/chapter5/listing68/counters"
)

// main is the entry point for the app
func main() {
	// create a var of the unexported type using the exported New func from the package counters
	counter := counters.New(10)

	fmt.Printf("Counter: %d\n", counter)
}
