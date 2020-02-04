package twofer

import "fmt"

// ShareWith is a simple function that prints "One for X, one for me." if X is null or empty then the program returns "One for you, one for me."
func ShareWith(name string) string {
	var response string
	if name == "" {
		name = "you"
	}
	response = fmt.Sprintf("One for %s, one for me.", name)
	return response
}
