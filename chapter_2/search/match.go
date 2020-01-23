package search

import (
	"log"
)

// Result contains the result of a search
type Result struct {
	Field   string
	Content string
}

// Matcher defines the behavior requared by types that want
// to implement a new search type
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match is launched as a goroutine for each individual feed to run
// searches concurrently
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	//perform the search against the specified matcher
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// write the results to the channel
	for _, result := range searchResults {
		results <- result
	}
}

// Display writes results to the console window as they
// are received by the individual goroutines
func Display(results chan *Result) {
	// the channel blocks untill a result is written to the channel
	// once the channel is closed the for loop terminares
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
