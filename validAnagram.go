package leetcode

/*
242. Valid Anagram

Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false
*/

func validAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counter := make(map[rune]int, len(s))
	for _, r := range s {
		counter[r]++
	}
	for _, r := range t {
		if _, ok := counter[r]; ok {
			counter[r]--
			if counter[r] == 0 {
				delete(counter, r)
			}
		} else {
			return false
		}
	}
	return len(counter) == 0
}
