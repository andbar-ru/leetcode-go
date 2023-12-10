package leetcode

/*
206. Reverse Linked List

Given the head of a singly linked list, reverse the list, and return the reversed list.

Example 1:
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Example 2:
Input: head = [1,2]
Output: [2,1]

Example 3:
Input: head = []
Output: []

Constraints:
* The number of nodes in the list is the range [0, 5000].
* -5000 <= Node.val <= 5000
*/

func reverseLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prevNode *ListNode
	curNode := head
	for curNode != nil {
		nextNode := curNode.Next
		curNode.Next = prevNode
		prevNode = curNode
		curNode = nextNode
	}

	return prevNode
}

func reverseLinkedListTopRated(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}
