package leetcode

/*
278. First Bad Version

You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API bool isBadVersion(version) which returns whether version is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.

Example 1:
Input: n = 5, bad = 4
Output: 4
Explanation:
call isBadVersion(3) -> false
call isBadVersion(5) -> true
call isBadVersion(4) -> true
Then 4 is the first bad version.

Example 2:
Input: n = 1, bad = 1
Output: 1

Constraints:
1 <= bad <= n <= 2^31 - 1
*/

var firstBadVersionFirstBad int

func isBadVersion(version int) bool {
	return version >= firstBadVersionFirstBad
}

func firstBadVersion(n int) int {
	if n == 1 || isBadVersion(1) {
		return 1
	}

	// Further firstVersion is always good and lastVersion is always bad.
	firstVersion := 1
	lastVersion := n

	for firstVersion < lastVersion-1 {
		version := (lastVersion + firstVersion) / 2
		if isBadVersion(version) {
			lastVersion = version
		} else {
			firstVersion = version
		}
	}

	return lastVersion
}

func firstBadVersionTopRatedV1(n int) int {
	start := 0
	end := n

	for mid := (end + start) / 2; start < end; mid = (end + start) / 2 {
		if isBadVersion(mid) {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return start
}

func firstBadVersionTopRatedV2(n int) int {
	maybeBad := 1

	for mid := (n + maybeBad) / 2; maybeBad < n; mid = (n + maybeBad) / 2 {
		if isBadVersion(mid) {
			n = mid
		} else {
			maybeBad = mid + 1
		}
	}

	return n
}

func firstBadVersionTopRatedV3(n int) int {
	maybeBad := 1

	for maybeBad < n {
		if isBadVersion((n + maybeBad) / 2) {
			n = (n + maybeBad) / 2
		} else {
			maybeBad = (n+maybeBad)/2 + 1
		}
	}

	return n
}

func firstBadVersionTopRatedV4(n int) int {
	start := 1
	end := n

	for mid := (end + start) / 2; start < end+1; mid = (end + start) / 2 {
		if isBadVersion(mid) {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return start
}

func firstBadVersionTopRatedV5(n int) int {
	left, right := 1, n
	var mid int

	for right >= left {
		mid = left + (right-left)/2
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
