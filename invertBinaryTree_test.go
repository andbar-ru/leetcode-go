package leetcode

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"testing"
)

func TestInvertBinaryTree(t *testing.T) {
	f := func(n int) *Int {
		return &Int{v: n}
	}

	var tests = []struct {
		root []*Int
		want []*Int
	}{
		{
			root: mapSlice([]int{4, 2, 7, 1, 3, 6, 9}, f),
			want: mapSlice([]int{4, 7, 2, 9, 6, 3, 1}, f),
		},
		{
			root: mapSlice([]int{2, 1, 3}, f),
			want: mapSlice([]int{2, 3, 1}, f),
		},
		{
			root: mapSlice([]int{}, f),
			want: mapSlice([]int{}, f),
		},
	}

	for _, test := range tests {
		tree := slice2tree(test.root)
		result := invertBinaryTree(tree)
		resultSlice := tree2slice(result)
		if resultSlice == nil {
			resultSlice = []*Int{}
		}
		if !reflect.DeepEqual(resultSlice, test.want) {
			t.Errorf("invertBinaryTree(%#v) = %#v, want %#v", test.root, resultSlice, test.want)
		}
	}
}

func TestInvertBinaryTreeRand(t *testing.T) {
	f := func(n int) *Int {
		return &Int{v: n}
	}

	numTests := 1
	r := rand.New(rand.NewSource(2000))

	for i := 0; i < numTests; i++ {
		length := int(math.Pow(2, float64(r.Intn(6))) - 1) // choice([0 1 3 7 15 31 63])
		nums, err := randSliceInt(length, length, -100, 100, int64(i))
		if err != nil {
			panic(err)
		}
		slicedTree := mapSlice(nums, f)

		tree := slice2tree(slicedTree)
		invertedTree := invertBinaryTree(tree)
		invertedTreeSlice := tree2slice(invertedTree)
		if invertedTreeSlice == nil {
			invertedTreeSlice = []*Int{}
		}
		slicedTreeStr := fmt.Sprint(slicedTree)
		invertSlicedTree(slicedTree)
		if !reflect.DeepEqual(invertedTreeSlice, slicedTree) {
			t.Errorf("invertBinaryTree(%v) = %v, want %s", slicedTree, invertedTreeSlice, slicedTreeStr)
		}
	}
}
