package leetcode

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	if n == nil {
		return "nil"
	}

	visited := make(map[*ListNode]struct{})
	var s string

	for n != nil {
		s += fmt.Sprint(n.Val)
		if _, ok := visited[n]; ok {
			s += " ...cycled"
			break
		}
		visited[n] = struct{}{}
		if n.Next != nil {
			s += " -> "
		}
		n = n.Next
	}
	return s
}

func slice2List(slice []int) *ListNode {
	dummy := new(ListNode)
	curNode := dummy

	for _, num := range slice {
		curNode.Next = &ListNode{Val: num}
		curNode = curNode.Next
	}

	return dummy.Next
}

func slice2ListCycle(slice []int, pos int) *ListNode {
	dummy := new(ListNode)
	curNode := dummy
	var nodeToCycle *ListNode

	for i, num := range slice {
		node := &ListNode{Val: num}
		curNode.Next = node
		if i == pos {
			nodeToCycle = node
		}
		curNode = curNode.Next
	}

	if nodeToCycle != nil {
		curNode.Next = nodeToCycle
	}

	return dummy.Next
}
