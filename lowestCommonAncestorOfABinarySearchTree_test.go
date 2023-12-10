package leetcode

import (
	"testing"
)

func TestLowestCommonAncestorOfABinarySearchTree(t *testing.T) {
	bigTree := slice2tree([]*Int{{6}, {2}, {8}, {0}, {4}, {7}, {9}, nil, nil, {3}, {5}})
	smallTree := slice2tree([]*Int{{2}, {1}})

	var tests = []struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		{
			root: bigTree,
			p:    bigTree.Left,
			q:    bigTree.Right,
			want: bigTree,
		},
		{
			root: bigTree,
			p:    bigTree.Left,
			q:    bigTree.Left.Right,
			want: bigTree.Left,
		},
		{
			root: smallTree,
			p:    smallTree,
			q:    smallTree.Left,
			want: smallTree,
		},
	}

	for _, test := range tests {
		result := lowestCommonAncestorOfABinarySearchTree(test.root, test.p, test.q)
		if result != test.want {
			t.Errorf("lowestCommonAncestorOfABinarySearchTree(%v, %v, %v) = %v, want %v", test.root, test.p, test.q, result, test.want)
		}
	}
}
