package leetcode

func intSlicesEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}

func listsEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 != nil || l2 != nil {
		return false
	}
	return true
}

func mapSlice[T, U any](a []T, f func(T) U) []U {
	n := make([]U, len(a))
	for i, v := range a {
		n[i] = f(v)
	}
	return n
}
