package leetcode

/*
110. Balanced Binary Tree

Given a binary tree, determine if it is height-balanced. A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: true
        ╭─╮
        │3│
        ╰─╯
      /     \
   ╭─╮       ╭──╮
   │9│       │20│
   ╰─╯       ╰──╯
            /   \
          ╭──╮   ╭─╮
          │15│   │7│
          ╰──╯   ╰─╯
Example 2:
Input: root = [1,2,2,3,3,null,null,4,4]
Output: false
                ╭─╮
                │1│
                ╰─╯
              /     \
           ╭─╮       ╭─╮
           │2│       │2│
           ╰─╯       ╰─╯
          /   \
        ╭─╮   ╭─╮
        │3│   │3│
        ╰─╯   ╰─╯
       /   \
	 ╭─╮   ╭─╮
     │4│   │4│
     ╰─╯   ╰─╯

Example 3:
Input: root = []
Output: true

Constraints:
The number of nodes in the tree is in the range [0, 5000].
-10^4 <= Node.val <= 10^4
*/

func balancedBinaryTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	var height func(*TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		hl, hr := height(node.Left), height(node.Right)
		if hl > hr {
			return hl + 1
		}
		return hr + 1
	}

	var balanced func(*TreeNode, *TreeNode) bool
	balanced = func(leftNode, rightNode *TreeNode) bool {
		if leftNode == nil && rightNode == nil {
			return true
		}
		hl := height(leftNode)
		hr := height(rightNode)
		if abs(hl-hr) > 1 {
			return false
		}
		leftBalanced := true
		if leftNode != nil {
			leftBalanced = balanced(leftNode.Left, leftNode.Right)
		}
		rightBalanced := true
		if rightNode != nil {
			rightBalanced = balanced(rightNode.Left, rightNode.Right)
		}
		return leftBalanced && rightBalanced
	}

	return balanced(root.Left, root.Right)
}

func balancedBinaryTreeTopRated(root *TreeNode) bool {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var dfs func(*TreeNode) (bool, int)
	dfs = func(root *TreeNode) (bool, int) {
		if root == nil {
			return true, 0
		}

		isLeftBalanced, leftHeight := dfs(root.Left)
		isRightBalanced, rightHeight := dfs(root.Right)
		diff := abs(leftHeight - rightHeight)
		if isLeftBalanced && isRightBalanced && diff <= 1 {
			return true, 1 + max(leftHeight, rightHeight)
		}
		return false, -1
	}

	res, _ := dfs(root)
	return res
}
