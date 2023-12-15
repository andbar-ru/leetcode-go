package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsertInterval(t *testing.T) {
	var cases = []struct {
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			want:        [][]int{{1, 5}, {6, 9}},
		},
		{
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			want:        [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			intervals:   [][]int{},
			newInterval: []int{2, 5},
			want:        [][]int{{2, 5}},
		},
		{
			intervals:   [][]int{{6, 12}},
			newInterval: []int{1, 4},
			want:        [][]int{{1, 4}, {6, 12}},
		},
		{
			intervals:   [][]int{{6, 12}},
			newInterval: []int{15, 18},
			want:        [][]int{{6, 12}, {15, 18}},
		},
		{
			intervals:   [][]int{{6, 12}},
			newInterval: []int{4, 15},
			want:        [][]int{{4, 15}},
		},
		{
			intervals:   [][]int{{6, 12}},
			newInterval: []int{6, 12},
			want:        [][]int{{6, 12}},
		},
		{
			intervals:   [][]int{{6, 12}},
			newInterval: []int{8, 10},
			want:        [][]int{{6, 12}},
		},
		{
			intervals:   [][]int{{6, 12}},
			newInterval: []int{4, 12},
			want:        [][]int{{4, 12}},
		},
		{
			intervals:   [][]int{{6, 12}},
			newInterval: []int{6, 15},
			want:        [][]int{{6, 15}},
		},
		{
			intervals:   [][]int{{6, 12}},
			newInterval: []int{4, 10},
			want:        [][]int{{4, 12}},
		},
		{
			intervals:   [][]int{{6, 12}},
			newInterval: []int{10, 15},
			want:        [][]int{{6, 15}},
		},
		{
			intervals:   [][]int{{0, 4}, {15, 18}},
			newInterval: []int{6, 12},
			want:        [][]int{{0, 4}, {6, 12}, {15, 18}},
		},
		{
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{0, 20},
			want:        [][]int{{0, 20}},
		},
		{
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{11, 11},
			want:        [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {11, 11}, {12, 16}},
		},
		{
			intervals:   [][]int{{0, 5}, {9, 12}},
			newInterval: []int{7, 16},
			want:        [][]int{{0, 5}, {7, 16}},
		},
	}

	for _, c := range cases {
		origIntervals := fmt.Sprint(c.intervals)
		result := insertInterval(c.intervals, c.newInterval)
		if !reflect.DeepEqual(result, c.want) {
			t.Errorf("insertInterval(%s, %v) = %v, want %v", origIntervals, c.newInterval, result, c.want)
		}
	}
}
