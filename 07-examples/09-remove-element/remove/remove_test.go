package main

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		input    []int
		val      int
		expected int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
		{[]int{1}, 1, 0},
		{[]int{}, 5, 0},
		{[]int{7, 7, 7, 7}, 7, 0},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test case %d", i+1), func(t *testing.T) {
			result := removeElement(test.input, test.val)
			if result == test.expected {
				fmt.Printf("Test case %d: Input = %v, Expected = %d, Result = %d (PASS)\n", i+1, test.input, test.expected, result)
			} else {
				fmt.Printf("Test case %d: Input = %v, Expected = %d, Result = %d (FAIL)\n", i+1, test.input, test.expected, result)
			}
		})
	}
}
