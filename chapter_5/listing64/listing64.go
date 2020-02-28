// sample programm to show how the program can't access an unexported identifier from another package
package main

import (
	"fmt"
)

// main is the entry point for the app
func main() {
	// create a var of the unexported type and init the value to 10
	counter := counters.alertCounter(10)

	// ./listing64.go:15: cannot refer to unexported name counters.alertCounter
	// ./listing64.go:15: undefined: counters.alertCounter

	fmt.Printf("Counter: %d\n", counter)
}
