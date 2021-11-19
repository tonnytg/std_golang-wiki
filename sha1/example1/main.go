package main

import (
"crypto/sha1"
"fmt"
)

func main() {
	s := "this is my password in sha1"
	s1 := "this is my password in sha1"
	s2 := "this is my password in sha1 v2"

	h := sha1.New()
	h1 := sha1.New()
	h2 := sha1.New()

	h.Write([]byte(s))
	h1.Write([]byte(s))
	h2.Write([]byte(s))

	bs := h.Sum(nil)
	bs1 := h1.Sum(nil)
	bs2 := h2.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	fmt.Println(s1)
	fmt.Printf("%x\n", bs1)

	fmt.Println(s2)
	fmt.Printf("%x\n", bs2)
}
