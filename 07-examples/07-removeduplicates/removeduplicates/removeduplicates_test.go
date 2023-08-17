// mysqrt_test.go

package main

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
		result   []int
	}{
		{[]int{1, 1, 2}, 2, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{[]int{}, 0, []int{}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
	}

	for i, test := range tests {
		result := removeDuplicates(test.input)

		if result == test.expected {
			fmt.Printf("Test case %d: Input = %d, Expected = %d, Result = %d (PASS)\n", i+1, test.input, test.expected, result)
		} else {
			fmt.Printf("Test case %d: Input = %d, Expected = %d, Result = %d (FAIL)\n", i+1, test.input, test.expected, result)
			t.Errorf("For input=%d, expected %d but got %d", test.input, test.expected, result)
		}
	}
}
