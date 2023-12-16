package leetcode

/*
3. Longest Substring Without Repeating Characters

Given a string s, find the length of the longest substring without repeating characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Constraints:
* 0 <= s.length <= 5 * 10^4
* s consists of English letters, digits, symbols and spaces.
*/

func longestSubstringWithoutRepeatingCharacters(s string) int {
	b2i := make([]int, 127) // 126 - max code of printable char ('~')
	var maxLen, curLen int
	for i := 0; i < len(s); i++ {
		if b2i[s[i]] != 0 {
			if curLen > maxLen {
				maxLen = curLen
			}
			diff := i - b2i[s[i]]
			if diff < curLen {
				curLen = diff
			}
		}
		b2i[s[i]] = i + 1
		curLen++
	}
	if curLen > maxLen {
		return curLen
	}
	return maxLen
}

func longestSubstringWithoutRepeatingCharactersTopRated(s string) int {
	dict := [128]bool{}
	length, max := 0, 0
	for i, j := 0, 0; i < len(s); i++ {
		index := s[i]
		if dict[index] {
			for ; dict[index]; j++ {
				length--
				dict[s[j]] = false
			}
		}

		dict[index] = true
		length++
		if length > max {
			max = length
		}
	}
	return max
}
