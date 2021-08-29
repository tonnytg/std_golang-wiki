package regex

import (
	"regexp"
)

func MyRegex(email string) bool {
	// https://regexr.com
	// https://gobyexample.com/regular-expressions

	r := regexp.MustCompile("^[a-z0-9]+@+[a-z0-9]+\\.[a-z0-9.][a-z0-9]")
	return r.Match([]byte(email))
}