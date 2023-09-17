package leetcode

import (
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	var tests = []struct {
		s    string
		want bool
	}{
		{
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			s:    "race a car",
			want: false,
		},
		{
			s:    " ",
			want: true,
		},
	}

	for _, test := range tests {
		result := validPalindrome(test.s)
		if result != test.want {
			t.Errorf("validPalindrome(%s) = %t, want %t", test.s, result, test.want)
		}
	}
}

func BenchmarkValidPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = validPalindrome("A man, a plan, a canal: Panama")
		_ = validPalindrome("race a car")
		_ = validPalindrome(" ")
	}
}

func BenchmarkValidPalindromeTopRated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = validPalindromeTopRated("A man, a plan, a canal: Panama")
		_ = validPalindromeTopRated("race a car")
		_ = validPalindromeTopRated(" ")
	}
}
