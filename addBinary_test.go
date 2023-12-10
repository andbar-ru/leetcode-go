package leetcode

import (
	"testing"
)

func TestAddBinary(t *testing.T) {
	var cases = []struct {
		a    string
		b    string
		want string
	}{
		{
			a:    "11",
			b:    "1",
			want: "100",
		},
		{
			a:    "1010",
			b:    "1011",
			want: "10101",
		},
	}

	for _, c := range cases {
		result := addBinary(c.a, c.b)
		if result != c.want {
			t.Errorf("addBinary(%q, %q) = %q, want %q", c.a, c.b, result, c.want)
		}
	}
}
