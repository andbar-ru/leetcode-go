package leetcode

import (
	"strings"
	"unicode"
)

/*
125. Valid Palindrome

A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.
*/

var alphaNumeric map[byte]byte = map[byte]byte{
	'A': 'a',
	'B': 'b',
	'C': 'c',
	'D': 'd',
	'E': 'e',
	'F': 'f',
	'G': 'g',
	'H': 'h',
	'I': 'i',
	'J': 'j',
	'K': 'k',
	'L': 'l',
	'M': 'm',
	'N': 'n',
	'O': 'o',
	'P': 'p',
	'Q': 'q',
	'R': 'r',
	'S': 's',
	'T': 't',
	'U': 'u',
	'V': 'v',
	'W': 'w',
	'X': 'x',
	'Y': 'y',
	'Z': 'z',
	'a': 'a',
	'b': 'b',
	'c': 'c',
	'd': 'd',
	'e': 'e',
	'f': 'f',
	'g': 'g',
	'h': 'h',
	'i': 'i',
	'j': 'j',
	'k': 'k',
	'l': 'l',
	'm': 'm',
	'n': 'n',
	'o': 'o',
	'p': 'p',
	'q': 'q',
	'r': 'r',
	's': 's',
	't': 't',
	'u': 'u',
	'v': 'v',
	'w': 'w',
	'x': 'x',
	'y': 'y',
	'z': 'z',
	'0': '0',
	'1': '1',
	'2': '2',
	'3': '3',
	'4': '4',
	'5': '5',
	'6': '6',
	'7': '7',
	'8': '8',
	'9': '9',
}

func validPalindrome(s string) bool {
	chars := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if c, ok := alphaNumeric[s[i]]; ok {
			chars = append(chars, c)
		}
	}
	l := len(chars)
	for i := 0; i < l/2; i++ {
		if chars[i] != chars[l-1-i] {
			return false
		}
	}
	return true
}

func validPalindromeTopRated(s string) bool {
	f := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}

	s = strings.Map(f, s)
	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
