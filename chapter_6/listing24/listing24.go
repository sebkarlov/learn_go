// this sample program demonstrates how to use a buffered channel to work on multiple tasks with a predefined number of goroutines
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // number of goroutines to use
	taskLoad         = 10 // amount of work to process
)

// wg is used to wait for the programm to finish
var wg sync.WaitGroup

// init is called to initialize the package by the Go runtime prior to any other code being executed
func init() {
	// Seed the random number generator
	rand.Seed(time.Now().Unix())
}

// main is the entry point for all Go programs
func main() {
	// create a buffered channel to manage the task load
	tasks := make(chan string, taskLoad)

	// launch goroutines to handle the work
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// add a bunch of work to get done
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// close the channel so the goroutines will quit when all the work is done
	close(tasks)

	// Wait for all work to get done
	wg.Wait()
}

// worker is launched as a goroutine to process work from the buffered channel
func worker(tasks chan string, worker int) {
	// report that we just returned
	defer wg.Done()

	for {
		// wait for work to be assigned
		task, ok := <-tasks
		if !ok {
			// this means the channel is empty and close
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		// display we are starting the work
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		// randonly wait to simulate work time
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// display we finished the work
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
