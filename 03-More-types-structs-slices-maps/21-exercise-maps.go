package main

import (
	"strings"

	"wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s) // Split the string into words
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++ // Increment the count for each word
	}

	return wordCount
}

func main() {
	wc.Test(WordCount)
}
