package leetcode

import (
	"fmt"
	"slices"
	"sort"
)

/*
15. 3Sum

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

Example 1:
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.

Example 2:
Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.

Example 3:
Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.

Constraints:
* 3 <= nums.length <= 3000
* -10^5 <= nums[i] <= 10^5
*/

func _3sumBruteForce(nums []int) [][]int {
	mapTriplets := make(map[string]struct{})
	triplets := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			for k := 0; k < len(nums); k++ {
				if i != j && i != k && j != k && nums[i]+nums[j]+nums[k] == 0 {
					triplet := []int{nums[i], nums[j], nums[k]}
					slices.Sort(triplet)
					t := fmt.Sprint(triplet)
					if _, ok := mapTriplets[t]; !ok {
						triplets = append(triplets, triplet)
						mapTriplets[t] = struct{}{}
					}
				}
			}
		}
	}

	return triplets
}

func _3sum(nums []int) [][]int {
	triplets := make([][]int, 0)
	mTriplets := make(map[[3]int]struct{})

	if len(nums) < 3 {
		return triplets
	}

	negatives := make([]int, 0, len(nums)/2)
	mNegatives := make(map[int]int)
	nZeros := 0
	positives := make([]int, 0, len(nums)/2)
	mPositives := make(map[int]int)
	for _, n := range nums {
		switch {
		case n < 0:
			negatives = append(negatives, n)
			mNegatives[n]++
		case n > 0:
			positives = append(positives, n)
			mPositives[n]++
		default:
			nZeros++
		}
	}

	appendTriplet := func(triplet [3]int) {
		if _, ok := mTriplets[triplet]; !ok {
			triplets = append(triplets, triplet[:])
			mTriplets[triplet] = struct{}{}
		}
	}

	if nZeros > 2 {
		appendTriplet([3]int{0, 0, 0})
	}
	if len(negatives) == 0 || len(positives) == 0 {
		return triplets
	}

	if nZeros > 0 {
		for k := range mPositives {
			if mNegatives[-k] > 0 {
				appendTriplet([3]int{-k, 0, k})
			}
		}
	}

	slices.SortFunc(negatives, func(a, b int) int { return b - a })
	slices.Sort(positives)

	pHigh := 0
	for _, n := range negatives {
		twoSum := -n
		for pHigh < len(positives) && positives[pHigh] < twoSum {
			pHigh++
		}
		for _, p := range positives[:pHigh] {
			p1 := twoSum - p
			if p1 == p {
				if mPositives[p] > 1 {
					appendTriplet([3]int{n, p, p})
				}
			} else if mPositives[p1] > 0 {
				if p < p1 {
					appendTriplet([3]int{n, p, p1})
				} else {
					appendTriplet([3]int{n, p1, p})
				}
			}
		}
	}

	pHigh = 0
	for _, p := range positives {
		twoSum := -p
		for pHigh < len(negatives) && negatives[pHigh] > twoSum {
			pHigh++
		}
		for _, n := range negatives[:pHigh] {
			n1 := twoSum - n
			if n1 == n {
				if mNegatives[n] > 1 {
					appendTriplet([3]int{n, n, p})
				}
			} else if mNegatives[n1] > 0 {
				if n < n1 {
					appendTriplet([3]int{n, n1, p})
				} else {
					appendTriplet([3]int{n1, n, p})
				}
			}
		}
	}

	return triplets
}

func _3sumV2(nums []int) [][]int {
	triplets := make([][]int, 0)

	if len(nums) < 3 {
		return triplets
	}

	mTriplets := make(map[[3]int]struct{})
	n2n := make(map[int]int, 0)

	appendTriplet := func(triplet [3]int) {
		if _, ok := mTriplets[triplet]; !ok {
			triplets = append(triplets, triplet[:])
			mTriplets[triplet] = struct{}{}
		}
	}

	sortTriplet := func(t []int) { // It is faster than slices.Sort.
		if t[0] > t[1] {
			t[0], t[1] = t[1], t[0]
		}
		if t[1] > t[2] {
			t[1], t[2] = t[2], t[1]
		}
		if t[0] > t[1] {
			t[0], t[1] = t[1], t[0]
		}
	}

	for _, n := range nums {
		n2n[n]++
	}

	// Special case
	if n2n[0] > 2 {
		appendTriplet([3]int{0, 0, 0})
	}

	for n1, c1 := range n2n {
		for n2 := range n2n {
			switch {
			case n1 == 0 && n2 == 0:
				break // processed above
			case n1 == n2:
				if c1 < 2 {
					break
				}
				fallthrough
			default:
				n3 := -(n1 + n2)
				gt := 0 // greater than
				if n3 == n1 || n3 == n2 {
					gt = 1
				}
				if n2n[n3] > gt {
					triplet := []int{n1, n2, n3}
					sortTriplet(triplet)
					appendTriplet([3]int(triplet))
				}
			}
		}
	}

	return triplets
}

func _3sumTopRated(nums []int) [][]int {
	results := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // To prevent the repeat
		}
		target, left, right := -nums[i], i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				results = append(results, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > target {
				right--
			} else if sum < target {
				left++
			}
		}
	}
	return results
}
