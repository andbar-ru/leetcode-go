package leetcode

/*
2. Add Two Numbers

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:
Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]

Constraints:
The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, curNode *ListNode
	var carry int

	for l1 != nil || l2 != nil {
		var n1, n2 int
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		node := &ListNode{Val: sum % 10}
		carry = sum / 10
		if head == nil {
			head = node
			curNode = node
		} else {
			curNode.Next = node
			curNode = curNode.Next
		}
	}

	if carry != 0 {
		curNode.Next = &ListNode{Val: carry}
	}

	return head
}

func addTwoNumbersTopRated(l1 *ListNode, l2 *ListNode) *ListNode {
	carry, dummy := 0, new(ListNode)

	for node := dummy; l1 != nil || l2 != nil || carry > 0; node = node.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		node.Next = &ListNode{carry % 10, nil}
		carry /= 10
	}

	return dummy.Next
}
