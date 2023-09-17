package leetcode

/*
21. Merge Two Sorted Lists

You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Example 1:
Input: list1 = [1->2->4], list2 = [1->3->4]
Output: [1->1->2->3->4->4]

Example 2:
Input: list1 = [], list2 = []
Output: []

Example 3:
Input: list1 = [], list2 = [0]
Output: [0]
*/

func mergeTwoSortedLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	var curr *ListNode

	addNode := func(node *ListNode) {
		if curr == nil {
			curr = node
			head = node
		} else {
			curr.Next = node
			curr = curr.Next
		}
	}

	for list1 != nil && list2 != nil {
		var val int
		if list1.Val < list2.Val {
			val = list1.Val
			list1 = list1.Next
		} else {
			val = list2.Val
			list2 = list2.Next
		}
		addNode(&ListNode{Val: val})
	}

	for list1 != nil {
		addNode(&ListNode{Val: list1.Val})
		list1 = list1.Next
	}

	for list2 != nil {
		addNode(&ListNode{Val: list2.Val})
		list2 = list2.Next
	}

	return head
}

func mergeTwoSortedListsTopRated(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoSortedListsTopRated(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoSortedListsTopRated(list1, list2.Next)
	return list2
}
