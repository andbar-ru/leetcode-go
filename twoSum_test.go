package leetcode

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 4, 2, 1},
			target:   10,
			expected: nil,
		},
		{
			nums:     []int{},
			target:   8,
			expected: nil,
		},
	}

	for _, test := range tests {
		result := twoSum(test.nums, test.target)
		if result != nil {
			slices.Sort(result)
		}
		if !intSlicesEqual(result, test.expected) {
			t.Errorf("twoSum(%v, %v) = %v, expected %v", test.nums, test.target, result, test.expected)
		}
	}
}
