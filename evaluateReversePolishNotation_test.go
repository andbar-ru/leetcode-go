package leetcode

import (
	"testing"
)

func TestEvaluateReversePolishNotation(t *testing.T) {
	var cases = []struct {
		tokens []string
		want   int
	}{
		{
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
		{
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
	}

	for _, c := range cases {
		result := evaluateReversePolishNotationV2(c.tokens)
		if result != c.want {
			t.Errorf("evaluateReversePolishNotation(%v) = %d, want %d", c.tokens, result, c.want)
		}
	}
}
