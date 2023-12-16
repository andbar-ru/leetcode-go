package leetcode

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func genKClosestPointsToOriginPoints() [][]int {
	r := rand.New(rand.NewSource(0))
	n := 10000
	points := make([][]int, n)
	for i := 0; i < n; i++ {
		x := r.Intn(20001) - 10000
		y := r.Intn(20001) - 10000
		points[i] = []int{x, y}
	}
	return points
}

func TestKClosestPointsToOrigin(t *testing.T) {
	var cases = []struct {
		points [][]int
		k      int
		want   [][]int // must be presorted
	}{
		{
			points: [][]int{{1, 3}, {-2, 2}},
			k:      1,
			want:   [][]int{{-2, 2}},
		},
		{
			points: [][]int{{3, 3}, {5, -1}, {-2, 4}},
			k:      2,
			want:   [][]int{{-2, 4}, {3, 3}},
		},
	}

	for _, c := range cases {
		points := fmt.Sprint(c.points)
		result := kClosestPointsToOriginV2(c.points, c.k)
		sort.Slice(result, func(i, j int) bool {
			if result[i][0] == result[j][0] {
				return result[i][1] < result[j][1]
			}
			return result[i][0] < result[j][0]
		})
		if !reflect.DeepEqual(result, c.want) {
			t.Errorf("kClosestPointsToOrigin(%s, %d) = %v, want %v", points, c.k, result, c.want)
		}
	}
}

func BenchmarkKClosestPointsToOriginV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		points := genKClosestPointsToOriginPoints()
		kClosestPointsToOriginV1(points, 5000)
	}
}

func BenchmarkKClosestPointsToOriginV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		points := genKClosestPointsToOriginPoints()
		kClosestPointsToOriginV2(points, 5000)
	}
}

func BenchmarkKClosestPointsToOriginV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		points := genKClosestPointsToOriginPoints()
		kClosestPointsToOriginV3(points, 5000)
	}
}

func BenchmarkKClosestPointsToOriginTopRated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		points := genKClosestPointsToOriginPoints()
		kClosestPointsToOriginTopRated(points, 5000)
	}
}

func BenchmarkKClosestPointsToOriginTopRatedV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		points := genKClosestPointsToOriginPoints()
		kClosestPointsToOriginTopRatedV2(points, 5000)
	}
}
