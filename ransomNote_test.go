package leetcode

import (
	"testing"
)

func TestRansomNote(t *testing.T) {
	var cases = []struct {
		ransomNote string
		magazine   string
		want       bool
	}{
		{
			ransomNote: "a",
			magazine:   "b",
			want:       false,
		},
		{
			ransomNote: "aa",
			magazine:   "ab",
			want:       false,
		},
		{
			ransomNote: "aa",
			magazine:   "aab",
			want:       true,
		},
		{
			ransomNote: "aab",
			magazine:   "baa",
			want:       true,
		},
	}

	for _, c := range cases {
		result := ransomNote(c.ransomNote, c.magazine)
		if result != c.want {
			t.Errorf("ransomNote(%q, %q) = %t, want %t", c.ransomNote, c.magazine, result, c.want)
		}
	}
}
