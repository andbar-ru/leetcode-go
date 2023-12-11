package leetcode

import "container/list"

/*
104. Maximum Depth of Binary Tree

Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: 3

Example 2:
Input: root = [1,null,2]
Output: 2

Constraints:
* The number of nodes in the tree is in the range [0, 10^4].
* -100 <= Node.val <= 100
*/

func maximumDepthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		dl := depth(node.Left)
		dr := depth(node.Right)
		if dl > dr {
			return dl + 1
		} else {
			return dr + 1
		}
	}

	return depth(root)
}

func maximumDepthOfBinaryTreeTopRatedRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	dl := maximumDepthOfBinaryTreeTopRatedRecursive(root.Left)
	dr := maximumDepthOfBinaryTreeTopRatedRecursive(root.Right)
	if dl > dr {
		return dl + 1
	} else {
		return dr + 1
	}
}

func maximumDepthOfBinaryTreeTopRatedIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	level := 0
	for queue.Len() > 0 {
		n := queue.Len()
		for i := 0; i < n; i++ {
			e := queue.Front()
			queue.Remove(e)
			node := e.Value.(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		level++
	}
	return level
}
