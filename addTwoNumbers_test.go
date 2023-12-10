package leetcode

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	var tests = []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			l1:   slice2list([]int{2, 4, 3}),
			l2:   slice2list([]int{5, 6, 4}),
			want: slice2list([]int{7, 0, 8}),
		},
		{
			l1:   slice2list([]int{0}),
			l2:   slice2list([]int{0}),
			want: slice2list([]int{0}),
		},
		{
			l1:   slice2list([]int{9, 9, 9, 9, 9, 9, 9}),
			l2:   slice2list([]int{9, 9, 9, 9}),
			want: slice2list([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
		{
			l1:   slice2list([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
			l2:   slice2list([]int{5, 6, 4}),
			want: slice2list([]int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
		},
	}

	for _, test := range tests {
		result := addTwoNumbers(test.l1, test.l2)
		if !listsEqual(result, test.want) {
			t.Errorf("addTwoNumbers(%v, %v) = %v, want %v", test.l1, test.l2, result, test.want)
		}
	}
}
