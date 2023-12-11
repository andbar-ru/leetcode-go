package leetcode

import (
	"testing"
)

func TestMaximumDepthOfBinaryTree(t *testing.T) {
	var cases = []struct {
		root *TreeNode
		want int
	}{
		{
			root: slice2tree([]*Int{{3}, {9}, {20}, nil, nil, {15}, {7}}),
			want: 3,
		},
		{
			root: slice2tree([]*Int{{1}, nil, {2}}),
			want: 2,
		},
	}

	for _, c := range cases {
		result := maximumDepthOfBinaryTree(c.root)
		if result != c.want {
			t.Errorf("maximumDepthOfBinaryTree(%v) = %d, want %d", c.root, result, c.want)
		}
	}
}
