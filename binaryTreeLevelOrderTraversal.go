package leetcode

/*
102. Binary Tree Level Order Traversal

Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Example 2:
Input: root = [1]
Output: [[1]]

Example 3:
Input: root = []
Output: []

Constraints:
* The number of nodes in the tree is in the range [0, 2000].
* -1000 <= Node.val <= 1000
*/

func binaryTreeLevelOrderTraversal(root *TreeNode) [][]int {
	result := [][]int{}

	var traverse func(node *TreeNode, level int)
	traverse = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(result) < level+1 {
			result = append(result, []int{})
		}
		result[level] = append(result[level], node.Val)
		traverse(node.Left, level+1)
		traverse(node.Right, level+1)
	}

	traverse(root, 0)
	return result
}

func binaryTreeLevelOrderTraversalV2(root *TreeNode) [][]int {
	result := [][]int{}

	type NodeLevel struct {
		node  *TreeNode
		level int
	}

	queue := []NodeLevel{{root, 0}}
	for len(queue) > 0 {
		nl := queue[0]
		queue = queue[1:]
		if nl.node == nil {
			continue
		}
		if len(result) < nl.level+1 {
			result = append(result, []int{})
		}
		result[nl.level] = append(result[nl.level], nl.node.Val)
		if nl.node.Left != nil {
			queue = append(queue, NodeLevel{nl.node.Left, nl.level + 1})
		}
		if nl.node.Right != nil {
			queue = append(queue, NodeLevel{nl.node.Right, nl.level + 1})
		}
	}

	return result
}

func binaryTreeLevelOrderTraversalTopRated(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		var level []int
		for i := 0; i < qlen; i++ {
			node := queue[0]
			level = append(level, node.Val)
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
