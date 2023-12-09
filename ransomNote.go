package leetcode

import "strings"

/*
383. Ransom Note

Given two strings "ransomNote" and "magazine", return true if "ransomNote" can be constructed by using the letters from "magazine" and false otherwise.

Each letter in "magazine" can only be used once in "ransomNote".

Example 1:
Input: ransomNote = "a", magazine = "b"
Output: false

Example 2:
Input: ransomNote = "aa", magazine = "ab"
Output: false

Example 3:
Input: ransomNote = "aa", magazine = "aab"
Output: true

Constraints:
* 1 <= ransomNote.length, magazine.length <= 105
* ransomNote and magazine consist of lowercase English letters.
*/

func ransomNote(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}
	if len(ransomNote) == 0 {
		return true
	}

	c2n := make(map[byte]int, 26)
	for i := len(magazine) - 1; i >= 0; i-- {
		c2n[magazine[i]]++
	}
	for i := len(ransomNote) - 1; i >= 0; i-- {
		c := ransomNote[i]
		n := c2n[c]
		if n == 0 {
			return false
		}
		c2n[c]--
	}
	return true
}

func ransomNoteTopRated(ransomNote string, magazine string) bool {
	for _, v := range ransomNote {
		if strings.Count(ransomNote, string(v)) > strings.Count(magazine, string(v)) {
			return false
		}
	}
	return true
}

func ransomNoteTopRated2(ransomNote string, magazine string) bool {
	counts := [26]int{}
	for i := 0; i < len(magazine); i++ {
		counts[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		counts[ransomNote[i]-'a']--
		if counts[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
