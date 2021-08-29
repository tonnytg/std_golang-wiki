package regex

import (
	"fmt"
	"testing"
)

func TestMyRegex(t *testing.T) {
	if MyRegex("antonio@gmail.com") {
		fmt.Println("email valid")
	} else {
		t.Fatal("must be true")
	}

	if ! MyRegex("antoniogmail.com") {
		fmt.Println("email invalid")
	} else {
		t.Fatal("must be false")
	}
}