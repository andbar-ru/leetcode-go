package leetcode

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) String() string {
	return fmt.Sprint(tree2Slice(n))
}

func slice2Tree(slice []*Int) *TreeNode {
	var insertLevelOrder func([]*Int, int) *TreeNode
	insertLevelOrder = func(slice []*Int, i int) *TreeNode {
		if i >= len(slice) || slice[i] == nil {
			return nil
		}
		root := &TreeNode{Val: slice[i].v}
		root.Left = insertLevelOrder(slice, 2*i+1)
		root.Right = insertLevelOrder(slice, 2*i+2)
		return root
	}

	if slice == nil || len(slice) == 0 {
		return nil
	}

	return insertLevelOrder(slice, 0)
}

func tree2Slice(root *TreeNode) []*Int {
	if root == nil {
		return nil
	}

	i2v := make(map[int]*Int)
	maxI := -1

	var walk func(*TreeNode, int)
	walk = func(node *TreeNode, i int) {
		if node == nil {
			return
		}
		i2v[i] = &Int{v: node.Val}
		if i > maxI {
			maxI = i
		}
		walk(node.Left, 2*i+1)
		walk(node.Right, 2*i+2)
	}

	walk(root, 0)

	slice := make([]*Int, maxI+1)
	for i, v := range i2v {
		slice[i] = v
	}

	return slice
}

func invertSlicedTree(slice []*Int) {
	low, high := 0, 0
	for high < len(slice) {
		for i, j := low, high; i < j; i, j = i+1, j-1 {
			slice[i], slice[j] = slice[j], slice[i]
		}
		low = 2*low + 1
		high = 2*high + 2
	}
}
