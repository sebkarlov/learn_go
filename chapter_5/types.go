// user defines a user in the program
type user struct {
	name       string
	email      string
	ext        int
	priveleged bool
}

// declare a var of type user
var bill user

// declare a var of type user and initialize all the fields
lisa := user{
	name:       "Lisa",
	email:      "lisa@email.com",
	ext:        123,
	priveleged: true,
}

// admin represents an admin user with privileges
type admin struct {
	person user
	level  string
}

// declare a variable of type admin
fred := admin{
	person: user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		priveleged: true,
	},
	level: "super",
}

