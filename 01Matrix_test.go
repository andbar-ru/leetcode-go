package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func genWorst01MatrixAndWant() ([][]int, [][]int) {
	m := 100
	n := 100
	mat := make([][]int, m)
	want := make([][]int, m)
	for i := 0; i < m; i++ {
		row := make([]int, n)
		rowWant := make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				row[j] = 0
			} else {
				row[j] = 1
			}
			rowWant[j] = i + j
		}
		mat[i] = row
		want[i] = rowWant
	}
	return mat, want
}

func Test01Matrix(t *testing.T) {
	worstMat, worstWant := genWorst01MatrixAndWant()

	var cases = []struct {
		mat  [][]int
		want [][]int
	}{
		{
			mat:  [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			want: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		},
		{
			mat:  [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}},
			want: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}},
		},
		{
			mat:  worstMat,
			want: worstWant,
		},
	}

	for _, c := range cases {
		mat := fmt.Sprint(c.mat)
		result := _01Matrix(c.mat)
		if !reflect.DeepEqual(result, c.want) {
			t.Errorf("_01Matrix(%s) = %v, want %v", mat, result, c.want)
		}
	}
}

func BenchmarkWorst01Matrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		worstMat, _ := genWorst01MatrixAndWant()
		_01Matrix(worstMat)
	}
}

func BenchmarkWorst01MatrixTopRated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		worstMat, _ := genWorst01MatrixAndWant()
		_01MatrixTopRated(worstMat)
	}
}
