package leetcode

import (
	"testing"
)

func TestBalancedBinaryTree(t *testing.T) {
	var tests = []struct {
		root *TreeNode
		want bool
	}{
		{
			root: slice2tree([]*Int{{3}, {9}, {20}, nil, nil, {15}, {7}}),
			want: true,
		},
		{
			root: slice2tree([]*Int{{1}, {2}, {2}, {3}, {3}, nil, nil, {4}, {4}}),
			want: false,
		},
		{
			root: slice2tree([]*Int{}),
			want: true,
		},
	}

	for _, test := range tests {
		result := balancedBinaryTree(test.root)
		if result != test.want {
			t.Errorf("balancedBinaryTree(%v) = %t, want %t", test.root, result, test.want)
		}
	}
}
