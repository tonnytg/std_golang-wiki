// https://go.dev/play/p/UWrfgh6nDAZ

package main

import (
	"fmt"
	"strings"
)

func Encode(text string) string {
	var binary []string
	for _, c := range text {
		binary = append(binary, fmt.Sprintf("%08b", c))
	}
	return strings.Join(binary, "")
}

func Decode(binary string) string {
	var text []byte
	for i := 0; i < len(binary); i += 8 {
		if i+8 > len(binary) {
			break
		}
		var b byte
		fmt.Sscanf(binary[i:i+8], "%b", &b)
		text = append(text, b)
	}
	return string(text)
}

func main() {
	originalText := "Hello, World!"
	fmt.Println("Original text:", originalText)

	binaryText := Encode(originalText)
	fmt.Println("Binary text:", binaryText)

	decodedText := Decode(binaryText)
	fmt.Println("Decoded text:", decodedText)
}

