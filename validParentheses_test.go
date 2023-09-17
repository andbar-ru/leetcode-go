package leetcode

import (
	"testing"
)

func TestValidParentheses(t *testing.T) {
	var tests = []struct {
		s        string
		expected bool
	}{
		{
			s:        "()",
			expected: true,
		},
		{
			s:        "()[]{}",
			expected: true,
		},
		{
			s:        "(]",
			expected: false,
		},
		{
			s:        "{[()]}",
			expected: true,
		},
		{
			s:        "",
			expected: true,
		},
		{
			s:        "{{[(]})}",
			expected: false,
		},
		{
			s:        "{[()]",
			expected: false,
		},
	}

	for _, test := range tests {
		result := validParentheses(test.s)
		if result != test.expected {
			t.Errorf("validParentheses(%q) = %v, expected %v", test.s, result, test.expected)
		}
	}
}
