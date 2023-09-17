package leetcode

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   -1,
		},
	}

	for _, test := range tests {
		result := binarySearch(test.nums, test.target)
		if result != test.want {
			t.Errorf("binarySearch(%v, %d) = %d, want %d", test.nums, test.target, result, test.want)
		}
	}
}
