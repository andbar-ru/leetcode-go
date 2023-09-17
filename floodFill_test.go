package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	var tests = []struct {
		image [][]int
		sr    int
		sc    int
		color int
		want  [][]int
	}{
		{
			image: [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			sr:    1,
			sc:    1,
			color: 2,
			want:  [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
		{
			image: [][]int{{0, 0, 0}, {0, 0, 0}},
			sr:    0,
			sc:    0,
			color: 0,
			want:  [][]int{{0, 0, 0}, {0, 0, 0}},
		},
	}

	for _, test := range tests {
		origImageStr := fmt.Sprint(test.image)
		result := floodFill(test.image, test.sr, test.sc, test.color)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("floodFill(%s, %d, %d, %d) = %v, want %d", origImageStr, test.sr, test.sc, test.color, result, test.want)
		}
	}
}

func BenchmarkFloodFill(b *testing.B) {
	var tests = []struct {
		image [][]int
		sr    int
		sc    int
		color int
		want  [][]int
	}{
		{
			image: [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			sr:    1,
			sc:    1,
			color: 2,
			want:  [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
		{
			image: [][]int{{0, 0, 0}, {0, 0, 0}},
			sr:    0,
			sc:    0,
			color: 0,
			want:  [][]int{{0, 0, 0}, {0, 0, 0}},
		},
	}
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			origImage := make([][]int, len(test.image))
			for i, row := range test.image {
				origRow := make([]int, len(row))
				copy(origRow, row)
				origImage[i] = origRow
			}
			floodFill(origImage, test.sr, test.sc, test.color)
		}
	}
}
