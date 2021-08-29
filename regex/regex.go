package regex

import (
	"regexp"
)

func myRegex(text, look string) bool {
	// https://regexr.com
	// https://gobyexample.com/regular-expressions

	r, _ := regexp.Compile("p([a-z]+)ch")
	return r.Match([]byte("peach"))
}