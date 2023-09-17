package leetcode

import (
	"testing"
)

func TestLinkedListCycle(t *testing.T) {
	var tests = []struct {
		head *ListNode
		want bool
	}{
		{
			head: slice2ListCycle([]int{3, 2, 0, -4}, 1),
			want: true,
		},
		{
			head: slice2ListCycle([]int{1, 2}, 0),
			want: true,
		},
		{
			head: slice2ListCycle([]int{1}, -1),
			want: false,
		},
	}

	for _, test := range tests {
		result := linkedListCycle(test.head)
		if result != test.want {
			t.Errorf("linkedListCycle(%v) = %t, want %t", test.head, result, test.want)
		}
	}
}
