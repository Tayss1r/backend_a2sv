package main

import (
	"fmt"
	"strings"
	"unicode"
)

func WordFrequencyCount(s string) map[string]int {
	var word strings.Builder
	freq := make(map[string]int)

	s = strings.ToLower(s)

	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			word.WriteRune(ch)
		} else if word.Len() > 0 {
			freq[word.String()]++
			word.Reset()
		}
	}
	if word.Len() > 0 {
		freq[word.String()]++
	}

	return freq
}

func main() {
	text := "Hi my name is tayssir, tayssir"
	fmt.Println(WordFrequencyCount(text))
}
