// samlpe program to show how to use an interface in Go
package main

import (
	"fmt"
)

// notifier is an interface that defined notification type behavior
type notifier interface {
	notify()
}

// user defines a use in the program
type user struct {
	name  string
	email string
}

// notify implements a method with a pointer receiver
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// main is the entry point for the app
func main() {
	// create a value of type User and send a notification
	u := user{"Bill", "bill@email.com"}

	// sendNotification(u)

	// ./listing36.go:31: cannot use u (type user) as type notifier in argument to sendNotication:
	// user does not implement notifier (notify method has pointer reciever)

	sendNotification(&u)

	// PASSED THE ADDRESS AND NO MORE ERROR
}

// sendNotification accepts values that implement the notifier interface and sends notifications
func sendNotification(n notifier) {
	n.notify()
}
