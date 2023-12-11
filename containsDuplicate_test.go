package leetcode

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	var cases = []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{1, 2, 3, 1},
			want: true,
		},
		{
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want: true,
		},
	}

	for _, c := range cases {
		result := containsDuplicate(c.nums)
		if result != c.want {
			t.Errorf("containsDuplicate(%v) = %t, want %t", c.nums, result, c.want)
		}
	}
}
