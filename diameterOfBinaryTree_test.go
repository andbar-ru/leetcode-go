package leetcode

import (
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	var cases = []struct {
		root *TreeNode
		want int
	}{
		{
			root: slice2tree([]*Int{{1}, {2}, {3}, {4}, {5}}),
			want: 3,
		},
		{
			root: slice2tree([]*Int{{1}, {2}}),
			want: 1,
		},
	}

	for _, c := range cases {
		result := diameterOfBinaryTree(c.root)
		if result != c.want {
			t.Errorf("diameterOfBinaryTree(%v) = %d, want %d", c.root, result, c.want)
		}
	}
}
