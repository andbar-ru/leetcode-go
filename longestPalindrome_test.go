package leetcode

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	var cases = []struct {
		s    string
		want int
	}{
		{
			s: "abccccdd",
			// One longest palindrome that can be built is "dccaccd", whose length is 7.
			want: 7,
		},
		{
			s: "a",
			// The longest palindrome that can be built is "a", whose length is 1.
			want: 1,
		},
		{
			s:    "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth",
			want: 983,
		},
	}

	for _, c := range cases {
		result := longestPalindrome(c.s)
		if result != c.want {
			t.Errorf("longestPalindrome(%q) = %d, want %d", c.s, result, c.want)
		}
	}
}
