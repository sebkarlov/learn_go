// this sample program demonstrates how to use an unbuffered channel to simulate a game of tennis between two goroutines
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg is used to wait for the programm to finish
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano)
}

// main is the entry point for all Go programs
func main() {
	// create an unbuffered channel
	court := make(chan int)

	// add a count of two, one for each goroutines
	wg.Add(2)

	// launch two players
	go player("Nadal", court)
	go player("Djokovic", court)

	// start the set
	court <- 1

	// Wait for the game to finish
	wg.Wait()
}

// player simulates a person playing the game of tennis
func player(name string, court chan int) {
	// schedule the call to Done to tell main we are done
	defer wg.Done()

	for {
		// wait for the ball to be hit back to us
		ball, ok := <-court
		if !ok {
			// if the channel was closed we won
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// pick a random number and see if we miss the ball
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			//close  the channel to signal we lost
			close(court)
			return
		}

		// display and then increment the hit count by one
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// hit the ball back to the opposing player
		court <- ball
	}
}
