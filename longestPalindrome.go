package leetcode

/*
409. Longest Palindrome

Given a string `s` which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.

Letters are case sensitive, for example, "Aa" is not considered a palindrome here.

Example 1:
Input: s = "abccccdd"
Output: 7
Explanation: One longest palindrome that can be built is "dccaccd", whose length is 7.

Example 2:
Input: s = "a"
Output: 1
Explanation: The longest palindrome that can be built is "a", whose length is 1.

Constraints:
* 1 <= s.length <= 2000
* `s` consists of lowercase and/or uppercase English letters only.
*/

func longestPalindrome(s string) int {
	c2n := make(map[byte]int)
	var sum, oddCount int

	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		n := c2n[c] + 1
		if n&1 == 0 {
			oddCount--
			sum += 2
		} else {
			oddCount++
		}
		c2n[c] = n
	}

	if oddCount > 0 {
		sum++
	}
	return sum
}

func longestPalindromeTopRated(s string) int {
	count := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if _, ok := count[ch]; !ok {
			count[ch] = 1
		} else {
			count[ch] += 1
		}
	}

	ans := 0
	for _, val := range count {
		ans += val / 2 * 2
		if ans%2 == 0 && val%2 == 1 {
			ans += 1
		}
	}

	return ans
}
