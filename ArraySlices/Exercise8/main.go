package main

import (
	"fmt"
)

func reverseCorrectly(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)

}

func main() {

	test := []string{
		"Hello, world!",
		"Hello, 世界",
		"Todo está bien",
	}

	for _, t := range test {
		fmt.Printf("Input: %q, Correct reverse: %q\n", t, reverseCorrectly(t))
	}
	fmt.Println("")

}
