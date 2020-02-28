// sample programm to show how unexported fields from an exported struct type can't be accessed directly
package main

import (
	"fmt"

	"github.com/goinaction/code/chapter5/listing74/entities"
)

// main is the entry point for the app
func main() {
	// create a value of type Admin from the entities package
	a := entities.Admin{
		Rights: 10,
	}

	// set the exported fields from the unexported inner type
	a.Name = "Bill"
	a.email = "bill@email.com"

	fmt.Printf("User: %v\n", u)
}
