package leetcode

/*
20. Valid Parenttheses

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false
*/

func validParentheses(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 == 1 {
		return false
	}

	c2o := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]rune, 0, len(s))

	for _, r := range s {
		if o, ok := c2o[r]; ok {
			lastI := len(stack) - 1
			if lastI < 0 || stack[lastI] != o {
				return false
			}
			stack = stack[:lastI]
		} else {
			stack = append(stack, r)
		}
	}

	return len(stack) == 0
}
