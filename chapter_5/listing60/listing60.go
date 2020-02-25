// sample programm to show how to embed a type into another type and
// the relationshop between the inner and outer type
package main

import "fmt"

// notifier is an interface that defined notification type behavior
type notifier interface {
	notify()
}

// user defines a user in the program
type user struct {
	name  string
	email string
}

// notify implements a method that can be called via a value of type user
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// admin represents an admin user with privileges
type admin struct {
	user  // Embedded Type
	level string
}

// notify implements a method that can be called via a value of type admin
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name,
		a.email)
}

// main is the entry point for the app
func main() {
	// create an admin user
	ad := admin{
		user: user{
			name:  "jonh",
			email: "john@email.com",
		},
		level: "super",
	}

	// send the admin user a notification
	// the embedded inner type's implementation of the interface is NOT "promoted" to the outer type
	sendNotification(&ad)

	// we can access to the inner type's method direcly
	ad.user.notify()

	// the inner type's method is NOT promoted
	ad.notify()
}

// sendNotification accepts values that implement the notifier interface and sends notifications
func sendNotification(n notifier) {
	n.notify()
}
