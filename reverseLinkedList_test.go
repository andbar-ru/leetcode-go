package leetcode

import (
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	var cases = []struct {
		head []int
		want []int
	}{
		{
			head: []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			head: []int{1, 2},
			want: []int{2, 1},
		},
		{
			head: []int{},
			want: []int{},
		},
	}

	for _, c := range cases {
		head := slice2list(c.head)
		want := slice2list(c.want)
		result := reverseLinkedList(head)
		if !listsEqual(result, want) {
			t.Errorf("reverseLinkedList(%v) = %v, want %v", c.head, list2slice(result), c.want)
		}
	}
}
