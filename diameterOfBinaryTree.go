package leetcode

/*
543. Diameter of Binary Tree

Given the root of a binary tree, return the length of the diameter of the tree.

The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

The length of a path between two nodes is represented by the number of edges between them.

Example 1:
Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].

Example 2:
Input: root = [1,2]
Output: 1

Constraints:
* The number of nodes in the tree is in the range [1, 10^4].
* -100 <= Node.val <= 100
*/

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var diameter int

	var height func(*TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		hl, hr := height(node.Left), height(node.Right)
		sum := hl + hr
		if sum > diameter {
			diameter = sum
		}
		return max(hl, hr) + 1
	}

	height(root)

	return diameter
}

func diameterOfBinaryTreeTopRated(root *TreeNode) int {
	// post order traverse
	// for each recursion we want to get what is the largest number of nodes single sided
	res := 0
	var postorder func(*TreeNode) int
	postorder = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := postorder(node.Left)
		right := postorder(node.Right)
		if left+right > res {
			res = left + right
		}
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}
	postorder(root)
	return res
}
