package leetcode

/*
70. Climbing Stairs

You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

Constraints:
* 1 <= n <= 45
*/

var climbingStairsCache = make(map[int]int)

func climbingStairs(n int) int {
	if v, ok := climbingStairsCache[n]; ok {
		return v
	}

	var result int
	switch n {
	case 1:
		result = 1
	case 2:
		result = 2
	default:
		result = climbingStairs(n-1) + climbingStairs(n-2)
	}
	climbingStairsCache[n] = result
	return result
}

func climbingStairsTopRated(n int) int {
	next, secondNext := 0, 1
	for ; n > 0; n-- {
		next, secondNext = secondNext, next+secondNext
	}
	return secondNext
}
