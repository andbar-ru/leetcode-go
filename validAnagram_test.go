package leetcode

import (
	"testing"
)

func TestValidAnagram(t *testing.T) {
	var tests = []struct {
		s    string
		t    string
		want bool
	}{
		{
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			s:    "rat",
			t:    "car",
			want: false,
		},
	}

	for _, test := range tests {
		result := validAnagram(test.s, test.t)
		if result != test.want {
			t.Errorf("validAnagram(%q) = %t, want %t", test.s, result, test.want)
		}
	}
}
