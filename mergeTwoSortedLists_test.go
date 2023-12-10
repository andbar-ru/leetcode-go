package leetcode

import (
	"math/rand"
	"slices"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	var tests = []struct {
		list1    *ListNode
		list2    *ListNode
		expected *ListNode
	}{
		{
			list1:    slice2list([]int{1, 2, 4}),
			list2:    slice2list([]int{1, 3, 4}),
			expected: slice2list([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			list1:    slice2list([]int{}),
			list2:    slice2list([]int{}),
			expected: slice2list([]int{}),
		},
		{
			list1:    slice2list([]int{}),
			list2:    slice2list([]int{0}),
			expected: slice2list([]int{0}),
		},
	}

	for _, test := range tests {
		result := mergeTwoSortedLists(test.list1, test.list2)
		if !listsEqual(result, test.expected) {
			t.Errorf("mergeTwoSortedLists(%s, %s) = %s, expected %s", test.list1, test.list2, result, test.expected)
		}
	}
}

func TestMergeTwoSortedListsRand(t *testing.T) {
	numTests := 1000

	for i := 0; i < numTests; i++ {
		seed := 1000 + int64(i)
		nums, err := randSliceInt(0, 100, -100, 100, seed)
		if err != nil {
			panic(err)
		}
		slices.Sort(nums)
		r := rand.New(rand.NewSource(seed + 1000))
		nums1 := make([]int, 0, len(nums))
		nums2 := make([]int, 0, len(nums))
		for _, n := range nums {
			if r.Float32() < 0.5 {
				nums1 = append(nums1, n)
			} else {
				nums2 = append(nums2, n)
			}
		}
		list := slice2list(nums)
		list1 := slice2list(nums1)
		list2 := slice2list(nums2)
		result := mergeTwoSortedLists(list1, list2)
		if !listsEqual(result, list) {
			t.Errorf("mergeTwoSortedLists(%s, %s) = %s, expected %s", list1, list2, result, list)
		}
	}
}
