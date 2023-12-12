package leetcode

import (
	"testing"
)

func TestMaximumSubarray(t *testing.T) {
	var cases = []struct {
		nums []int
		want int
	}{
		{
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			nums: []int{1},
			want: 1,
		},
		{
			nums: []int{5, 4, -1, 7, 8},
			want: 23,
		},
		{
			nums: []int{-1},
			want: -1,
		},
		{
			nums: []int{-2, -1},
			want: -1,
		},
		{
			nums: []int{-1, 1, 2, 1},
			want: 4,
		},
	}

	for _, c := range cases {
		result := maximumSubarray(c.nums)
		if result != c.want {
			t.Errorf("maximumSubarray(%v) = %d, want %d", c.nums, result, c.want)
		}
	}
}
