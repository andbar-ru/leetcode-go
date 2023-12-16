package leetcode

import (
	"testing"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	var cases = []struct {
		s    string
		want int
	}{
		{
			s:    "abcabcbb",
			want: 3,
		},
		{
			s:    "bbbbb",
			want: 1,
		},
		{
			s:    "pwwkew",
			want: 3,
		},
		{
			s:    "abba",
			want: 2,
		},
	}

	for _, c := range cases {
		result := longestSubstringWithoutRepeatingCharacters(c.s)
		if result != c.want {
			t.Errorf("longestSubstringWithoutRepeatingCharacters(%q) = %d, want %d", c.s, result, c.want)
		}
	}
}
