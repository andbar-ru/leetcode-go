package leetcode

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	var cases = []struct {
		nums []int
		want int
	}{
		{
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
	}

	for _, c := range cases {
		result := majorityElement(c.nums)
		if result != c.want {
			t.Errorf("majorityElement(%v) = %d, want %d", c.nums, result, c.want)
		}
	}
}
