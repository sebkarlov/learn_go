// this sample program demonstrates how to use the work package to use a pool of goroutines to get work done
package main

import (
	"log"
	"sync"
	"time"
	
	"https://github.com/goinaction/code/chapter7/patterns/work"
)

// names provides a set of names to display
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason"
}

// namePrinter provides special support for the printing names
type namePrinter struct {
	name string
}

// Task implements the Worker interface
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

// main is the entry point for all Go programs
func main() {
	// Create a work pool with 2 goroutines
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i :=0; i < 100; i++ {
		// iterate over the slice of names
		for _, name := range names {
			// create a namePrinter and provide the specific name
			np := namePrinter{
				name: name,
			}

			go func() {
				// submit the task to be worked on. When RunTask returns we know it is being handled
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	// Shutdown the work pool and wait for all existing work to be completed
	p.Shutdown
}