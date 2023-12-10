package leetcode

/*
67. Add Binary

Given two binary strings a and b, return their sum as a binary string.

Example 1:
Input: a = "11", b = "1"
Output: "100"

Example 2:
Input: a = "1010", b = "1011"
Output: "10101"

Constraints:
* 1 <= a.length, b.length <= 10^4
* a and b consist only of '0' or '1' characters.
* Each string does not contain leading zeros except for the zero itself.
*/

func addBinary(a string, b string) string {
	var l int
	if len(a) > len(b) {
		l = len(a)
	} else {
		l = len(b)
	}

	sum := make([]byte, 0, l+1)

	var carry byte
	for i := 0; i < l; i++ {
		ai := len(a) - 1 - i
		bi := len(b) - 1 - i
		var ab, bb byte

		if ai >= 0 {
			if a[ai] == '1' {
				ab = 1
			} else {
				ab = 0
			}
		} else {
			ab = 0
		}

		if bi >= 0 {
			if b[bi] == '1' {
				bb = 1
			} else {
				bb = 0
			}
		} else {
			bb = 0
		}

		s := ab + bb + carry // 0..3
		if s >= 2 {
			carry = 1
		} else {
			carry = 0
		}
		s = s % 2 // 0 or 1
		sum = append(sum, '0'+s)
	}
	if carry == 1 {
		sum = append(sum, '1')
	}

	// Reverse
	for i, l := 0, len(sum); i < l/2; i++ {
		sum[i], sum[l-i-1] = sum[l-i-1], sum[i]
	}

	return string(sum)
}

func addBinaryTopRated(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	ans := make([]byte, len(a)+1)
	var carry byte
	for i, j := len(a), len(b); i >= 1 || j >= 1; {
		i, j = i-1, j-1
		var a2Digit, b2Digit byte
		if i >= 0 {
			a2Digit = a[i] - '0'
		}
		if j >= 0 {
			b2Digit = b[j] - '0'
		}
		// sum and carry of full adder
		sum := a2Digit ^ b2Digit ^ carry
		carry = a2Digit&b2Digit | carry&(a2Digit^b2Digit)
		ans[i+1] = sum + '0'
	}
	if carry == 1 {
		ans[0] = '1'
		return string(ans)
	}
	return string(ans[1:])
}
