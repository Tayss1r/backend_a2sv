package main

import (
	"fmt"
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	var cleaned []rune
	s = strings.ToLower(s)

	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			cleaned = append(cleaned, ch)
		}
	}

	n := len(cleaned)
	for i := 0; i < n/2; i++ {
		if cleaned[i] != cleaned[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome("radar"))
	fmt.Println(IsPalindrome("tayssir"))
}
