package leetcode

import (
	"testing"
)

func TestClimbingStairs(t *testing.T) {
	var cases = []struct {
		n    int
		want int
	}{
		{
			n:    2,
			want: 2,
		},
		{
			n:    3,
			want: 3,
		},
		{
			n:    45,
			want: 1836311903,
		},
	}

	for _, c := range cases {
		result := climbingStairs(c.n)
		if result != c.want {
			t.Errorf("climbingStairs(%d) = %d, want %d", c.n, result, c.want)
		}
	}
}
