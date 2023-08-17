package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "Please contact us at john@example.com or jane@example.org for more information."

	// Define a regex pattern to match email addresses
	emailPattern := `\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`

	// Compile the regex pattern
	re := regexp.MustCompile(emailPattern)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)

	// Print the matched email addresses
	fmt.Println("Matched email addresses:")
	for _, match := range matches {
		fmt.Println(match)
	}
}
