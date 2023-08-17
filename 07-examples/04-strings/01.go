// strings.Contains(s, substr string) bool: Checks if the string s contains the substring substr.

// strings.HasPrefix(s, prefix string) bool: Checks if the string s starts with the prefix prefix.

// strings.HasSuffix(s, suffix string) bool: Checks if the string s ends with the suffix suffix.

// strings.Index(s, substr string) int: Returns the index of the first occurrence of substr in s, or -1 if not found.

// strings.LastIndex(s, substr string) int: Returns the index of the last occurrence of substr in s, or -1 if not found.

// strings.Replace(s, old, new string, n int) string: Replaces old with new in the string s, up to n occurrences.

// strings.Count(s, substr string) int: Returns the number of non-overlapping occurrences of substr in s.

// strings.ToLower(s string) string: Converts the string s to lowercase.

// strings.ToUpper(s string) string: Converts the string s to uppercase.

// strings.TrimSpace(s string) string: Removes leading and trailing white spaces from the string s.

// strings.Split(s, sep string) []string: Splits the string s into a slice of strings using sep as the delimiter.

// strings.Join(a []string, sep string) string: Joins the elements of a string slice a into a single string with sep as the separator.
package main

import (
	"fmt"
	"strings"
)

func split(x string) []string {
	splits := strings.Split(x, " ")
	return splits
}
func join(x []string, start int, end int, sep string) string {
	joins := strings.Join(x[start:end], sep)
	return joins
}
func count(x []string) int {
	num_words := len(x)
	return num_words
}
func replace(x string, old string, new string, n int) string {
	new_string := strings.Replace(x, old, new, n)
	return new_string
}

func main() {

	s := "Given a non-negative integer x, return the square root of x rounded down to the nearest integer."
	splits := split(s)
	fmt.Println("split slice : ", splits[:5])
	fmt.Println("join split slices : ", join(splits, 5, 10, " "))
	fmt.Println("total word of s string is : ", count(splits))
	fmt.Println("replace string : ", replace(s, "integer", "float", 5))
	fmt.Printf("to upper : %v ,\nto lower : %v ,\nto title: %v", strings.ToUpper(s), strings.ToLower(s), strings.ToTitle(s))

}
