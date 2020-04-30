// this program demonstrates how to use a channel to monitor the amount of time the program is running and terminate the program if it runs too long
package main

import (
	"log"
	"time"

	"https://github.com/goinaction/code/chapter7/patterns/runner"
)

// timeout is the number of second the program has to finish
const timeout = 3 * time.Second

// main is the entry point for the program
func main() {
	log.Println("Starting work.")

	// create a new timer value for this run
	r := runner.New(timeout)

	// Add the tasks to be run
	r.Add(createTask(), createTask(), createTask())

	// run the tasks and handle the result
	if err := r.Start(); err != nill {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}

	log.Println("Process ended")
}

// createTask retuens an example task that sleeps for the specified number of seconds based on the id
func createTask() func(int) {
	return func(id int) {
		log.Println("Processoe - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}