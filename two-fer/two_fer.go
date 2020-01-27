//This is twofer program to print "One for X, one for me." is X is null or empty then the progam returs "One for you, one for me."
package twofer

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	var response string
	if len(name) > 0 {
		response := fmt.Sprintf("One for %s, one for me.", name)
	} else {
		response := fmt.Sprintf("One for you, one for me.")
	}
	return string(response)
}
