// mysqrt_test.go

package main

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"II", 2},
		{"VIII", 8},
		{"XVI", 16},
		{"C", 100},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"MCMXCIV", 1994},
		{"MMMDCCXXIV", 3724},
	}

	for i, test := range tests {
		result := romanToInt(test.input)

		if result == test.expected {
			fmt.Printf("Test case %d: Input = %s, Expected = %d, Result = %d (PASS)\n", i+1, test.input, test.expected, result)
		} else {
			fmt.Printf("Test case %d: Input = %s, Expected = %d, Result = %d (FAIL)\n", i+1, test.input, test.expected, result)
			t.Errorf("For input=%s, expected %d but got %d", test.input, test.expected, result)
		}
	}
}
