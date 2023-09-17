package leetcode

import (
	"slices"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	var tests = []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
		},
		{
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
		},
	}

	for _, test := range tests {
		origNums1 := make([]int, len(test.nums1))
		copy(origNums1, test.nums1)
		mergeSortedArray(test.nums1, test.m, test.nums2, test.n)
		if !intSlicesEqual(test.nums1, test.expected) {
			t.Errorf("nums1 after mergeSortedArray(%v, %v, %v, %v) = %v, expected %v", origNums1, test.m, test.nums2, test.n, test.nums1, test.expected)
		}
	}
}

func TestMergeSortedArrayRand(t *testing.T) {
	numTests := 1000

	for i := 0; i < numTests; i++ {
		seed := 1000 + int64(i)
		nums1, err := randSliceInt(0, 10, -100, 100, seed)
		if err != nil {
			panic(err)
		}
		nums2, err := randSliceInt(0, 10, -100, 100, seed+1000)
		if err != nil {
			panic(err)
		}
		slices.Sort(nums1)
		slices.Sort(nums2)
		m, n := len(nums1), len(nums2)
		nums1 = append(nums1, make([]int, n)...)
		origNums1 := make([]int, len(nums1))
		copy(origNums1, nums1)
		mergeSortedArray(nums1, m, nums2, n)

		if len(nums1) != m+n {
			t.Errorf("seed: %d; nums1 after mergeSortedArray(%v, %v, %v, %v) = %v; len(nums1) != m + n\n", seed, origNums1, m, nums2, n, nums1)
		}
		if !slices.IsSorted(nums1) {
			t.Errorf("seed: %d; nums1 after mergeSortedArray(%v, %v, %v, %v) = %v; nums1 is not sorted\n", seed, origNums1, m, nums2, n, nums1)
		}
	}
}
