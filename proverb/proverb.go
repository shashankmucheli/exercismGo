// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import (
	"fmt"
)

// Proverb should have a comment documenting it. Proverb practice problem in exercism.io
func Proverb(rhyme []string) []string {

	s := make([]string, len(rhyme))
	if len(rhyme) <= 0 {
		return []string{}
	}
	lastLine := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	if len(rhyme) == 1 {
		s[0] = lastLine
		return s
	}

	for i:=0; i<len(rhyme)-1 ;i++  {
		tmp := fmt.Sprintf("For want of a %s the %s was lost.",rhyme[i],rhyme[i+1])
		s[i] = tmp
	}
	s[len(rhyme)-1] = lastLine
	return s
}
