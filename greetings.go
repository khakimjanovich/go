package grettings

import "fmt"

// Hello returns a greetings for the named person.
func Hello(name string) string {
	// := is a shortcut for declaring a variable and assigning to some sort of value
	// in this case it is equal to
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome", name)
	message := fmt.Sprintf("Hi, %v. Welcome", name)
	return message
}
