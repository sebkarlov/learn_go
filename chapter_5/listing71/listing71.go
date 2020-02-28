// sample programm to show how unexported fields from an exported struct type can't be accessed directly
package main

import (
	"fmt"

	"github.com/goinaction/code/chapter5/listing71/entities"
)

// main is the entry point for the app
func main() {
	// create a value of type User from the entities package
	u := entities.User{
		Name:  "Bill",
		email: "bill@email.com",
	}

	// ./example69.go:16: unknown entities.User field 'email' in struct literal

	fmt.Printf("User: %v\n", u)
}
