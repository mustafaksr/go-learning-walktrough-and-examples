// 01-twosum_test.go

package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums       []int
		target     int
		expected   []int
		expectFail bool
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}, false},
		{[]int{3, 2, 4}, 6, []int{1, 2}, false},
		{[]int{3, 3}, 6, []int{0, 1}, false},
		{[]int{1, 2, 3, 4, 5}, 10, []int{}, true},
	}

	for _, test := range tests {
		result := twoSum(test.nums, test.target)

		if !test.expectFail && !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For nums=%v and target=%d, expected %v but got %v", test.nums, test.target, test.expected, result)
		}

		if test.expectFail && !reflect.DeepEqual(result, test.expected) {
			t.Logf("For nums=%v and target=%d, expected failure but got %v", test.nums, test.target, result)
		}
	}
}
