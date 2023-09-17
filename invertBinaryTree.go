package leetcode

/*
226. Invert Binary Tree

Given the root of a binary tree, invert the tree, and return its root.

Example 1:
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]
        ╭─╮                  ╭─╮
        │4│                  │4│
        ╰─╯                  ╰─╯
      /     \              /     \
   ╭─╮       ╭─╮        ╭─╮       ╭─╮
   │2│       │7│   =>   │7│       │2│
   ╰─╯       ╰─╯        ╰─╯       ╰─╯
  /   \     /   \      /   \     /   \
╭─╮   ╭─╮ ╭─╮   ╭─╮  ╭─╮   ╭─╮ ╭─╮   ╭─╮
│1│   │3│ │6│   │9│  │9│   │6│ │3│   │1│
╰─╯   ╰─╯ ╰─╯   ╰─╯  ╰─╯   ╰─╯ ╰─╯   ╰─╯

Example 2:
Input: root = [2,1,3]
Output: [2,3,1]
        ╭─╮                  ╭─╮
        │2│                  │2│
        ╰─╯                  ╰─╯
      /     \      =>      /     \
   ╭─╮       ╭─╮        ╭─╮       ╭─╮
   │1│       │3│        │3│       │1│
   ╰─╯       ╰─╯        ╰─╯       ╰─╯

Example 3:
Input: root = []
Output: []
*/

func invertBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left != nil || root.Right != nil {
		root.Left, root.Right = root.Right, root.Left
	}
	invertBinaryTree(root.Left)
	invertBinaryTree(root.Right)
	return root
}
