package leetcode

import (
	"reflect"
	"testing"
)

func TestBinaryTreeLevelOrderTraversal(t *testing.T) {
	var cases = []struct {
		root *TreeNode
		want [][]int
	}{
		{
			root: slice2tree([]*Int{{3}, {9}, {20}, nil, nil, {15}, {7}}),
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			root: slice2tree([]*Int{{1}}),
			want: [][]int{{1}},
		},
		{
			root: slice2tree([]*Int{}),
			want: [][]int{},
		},
	}

	for _, c := range cases {
		result := binaryTreeLevelOrderTraversalV2(c.root)
		if !reflect.DeepEqual(result, c.want) {
			t.Errorf("binaryTreeLevelOrderTraversal(%v) = %v, want %v", c.root, result, c.want)
		}
	}
}
