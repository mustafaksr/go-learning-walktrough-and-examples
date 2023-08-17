// palindrome_test.go

package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{12321, true},
		{123456, false},
		{3, true},
	}

	for i, test := range tests {
		result := isPalindrome(test.input)

		if result == test.expected {
			fmt.Printf("Test case %d: Input = %d, Expected = %v, Result = %v (PASS)\n", i+1, test.input, test.expected, result)
		} else {
			fmt.Printf("Test case %d: Input = %d, Expected = %v, Result = %v (FAIL)\n", i+1, test.input, test.expected, result)
			t.Errorf("For input=%d, expected %v but got %v", test.input, test.expected, result)
		}
	}
}
