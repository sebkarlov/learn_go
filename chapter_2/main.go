package main

import (
	"log"
	"os"

	"github.com/goinaction/code/chapter2/sample/search"
	_ "github.com/goinaction/code/chapter2/sample/search"
)

// init is called prior to main
func init() {
	// change the device for logging to stdout
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program
func main() {
	// perform the search for the specified term
	search.Run("president")
}
