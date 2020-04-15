// sending values into a channel

// unbuffered channel of integers
unbuffered := make(chan int)
// buffered channel of strings
buffered := make(chan string, 10)

// sending values into a channel

//bufferd channel of strings
buffered := make(chan string, 10)
// send a string through the channel
buffered <- "Gopher"

// receive a string from the channell
value := <-buffered