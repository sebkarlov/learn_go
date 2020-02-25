// sample programm to show how to embed a type into another type and
// the relationshop between the inner and outer type
package main

import "fmt"

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

	// we can access to the inner type's method direcly
	ad.user.notify()

	// the inner type's method is promoted
	ad.notify()
}
