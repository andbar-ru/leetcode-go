package leetcode

import (
	"testing"
)

func TestMiddleOfTheLinkedList(t *testing.T) {
	var cases = []struct {
		head *ListNode
		want *ListNode
	}{
		{
			head: slice2list([]int{1}),
			want: slice2list([]int{1}),
		},
		{
			head: slice2list([]int{1, 2}),
			want: slice2list([]int{2}),
		},
		{
			head: slice2list([]int{1, 2, 3}),
			want: slice2list([]int{2, 3}),
		},
		{
			head: slice2list([]int{1, 2, 3, 4}),
			want: slice2list([]int{3, 4}),
		},
		{
			head: slice2list([]int{1, 2, 3, 4, 5}),
			want: slice2list([]int{3, 4, 5}),
		},
		{
			head: slice2list([]int{1, 2, 3, 4, 5, 6}),
			want: slice2list([]int{4, 5, 6}),
		},
	}

	for _, c := range cases {
		result := middleOfTheLinkedList(c.head)
		if !listsEqual(result, c.want) {
			t.Errorf("middleOfTheLinkedList(%v) = %v, want %v", c.head, result, c.want)
		}
	}
}
