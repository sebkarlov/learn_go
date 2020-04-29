// this sample program demonstrates how to use an unbuffered channel to simulate a relay race between four goroutines
package main

import (
	"fmt"
	"sync"
	"time"
)

// wg is used to wait for the programm to finish
var wg sync.WaitGroup

// main is the entry point for all Go programs
func main() {
	// create an unbuffered channel
	baton := make(chan int)

	// add a count of one for the last runner
	wg.Add(1)

	// first runner to his mark
	go Runner(baton)

	// start the race
	court <- 1

	// Wait for the game to finish
	wg.Wait()
}

// Runner simulates a person running in the relay race
func Runner(baton chan int) {
	var newRunner int

	// wait to receive the baton
	runner := <-baton

	// start running around the track
	fmt.Printf("Runner %d Running With Baton\n", runner)

	// new runner to the line
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}
	// running arount the track
	time.Sleep(100 * time.Millisecond)

	// is the race over
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}
	// exchange the baton for the next runner
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)

	baton <- newRunner
}
