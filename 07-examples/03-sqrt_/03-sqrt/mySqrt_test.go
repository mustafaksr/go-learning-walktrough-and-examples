// mysqrt_test.go

package main

import (
	"fmt"
	"testing"
)

func TestMySqrt(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{4, 2},
		{8, 2},
		{16, 4},
		{100, 10},
	}

	for i, test := range tests {
		result := mySqrt(test.input)

		if result == test.expected {
			fmt.Printf("Test case %d: Input = %d, Expected = %d, Result = %d (PASS)\n", i+1, test.input, test.expected, result)
		} else {
			fmt.Printf("Test case %d: Input = %d, Expected = %d, Result = %d (FAIL)\n", i+1, test.input, test.expected, result)
			t.Errorf("For input=%d, expected %d but got %d", test.input, test.expected, result)
		}
	}
}
