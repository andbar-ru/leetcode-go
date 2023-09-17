package leetcode

/*
235. Lowest Common Ancestor of a Binary Search Tree

Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Example 1:
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.
        ╭─╮
        │6│
        ╰─╯
      /     \
   ╭─╮       ╭─╮
   │2│       │8│
   ╰─╯       ╰─╯
  /   \     /   \
╭─╮   ╭─╮ ╭─╮   ╭─╮
│0│   │4│ │7│   │9│
╰─╯   ╰─╯ ╰─╯   ╰─╯
     /   \
   ╭─╮   ╭─╮
   │3│   │5│
   ╰─╯   ╰─╯

Example 2:
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.
        ╭─╮
        │6│
        ╰─╯
      /     \
   ╭─╮       ╭─╮
   │2│       │8│
   ╰─╯       ╰─╯
  /   \     /   \
╭─╮   ╭─╮ ╭─╮   ╭─╮
│0│   │4│ │7│   │9│
╰─╯   ╰─╯ ╰─╯   ╰─╯
     /   \
   ╭─╮   ╭─╮
   │3│   │5│
   ╰─╯   ╰─╯

Example 3:
Input: root = [2,1], p = 2, q = 1
Output: 2

Constraints:
The number of nodes in the tree is in the range [2, 10^5].
-10^9 <= Node.val <= 10^9
All Node.val are unique.
p != q
p and q will exist in the BST.
*/

func lowestCommonAncestorOfABinarySearchTree(root, p, q *TreeNode) *TreeNode {
	switch {
	case p == root || q == root || (p.Val < root.Val && root.Val < q.Val) || (q.Val < root.Val && root.Val < p.Val):
		return root
	case p.Val < root.Val && q.Val < root.Val:
		return lowestCommonAncestorOfABinarySearchTree(root.Left, p, q)
	case p.Val > root.Val && q.Val > root.Val:
		return lowestCommonAncestorOfABinarySearchTree(root.Right, p, q)
	default:
		return nil
	}
}

func lowestCommonAncestorOfABinarySearchTreeTopRatedRecursive(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorOfABinarySearchTreeTopRatedRecursive(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorOfABinarySearchTreeTopRatedRecursive(root.Right, p, q)
	}
	return root
}

func lowestCommonAncestorOfABinarySearchTreeTopRatedIterative(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
