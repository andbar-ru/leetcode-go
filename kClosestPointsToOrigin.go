package leetcode

import (
	"container/heap"
	"slices"
	"sort"
)

/*
973. K Closest Points to Origin

Given an array of points where points[i] = [x_i, y_i] represents a point on the X-Y plane and an integer k, return the k closest points to the origin (0, 0).

The distance between two points on the X-Y plane is the Euclidean distance (i.e., √(x_1 - x_2)^2 + (y_1 - y_2)^2).

You may return the answer in any order. The answer is guaranteed to be unique (except for the order that it is in).

Example 1:
Input: points = [[1,3],[-2,2]], k = 1
Output: [[-2,2]]
Explanation:
The distance between (1, 3) and the origin is sqrt(10).
The distance between (-2, 2) and the origin is sqrt(8).
Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
We only want the closest k = 1 points from the origin, so the answer is just [[-2,2]].

Example 2:
Input: points = [[3,3],[5,-1],[-2,4]], k = 2
Output: [[3,3],[-2,4]]
Explanation: The answer [[-2,4],[3,3]] would also be accepted.

Constraints:
* 1 <= k <= points.length <= 10^4
* -10^4 <= x_i, y_i <= 10^4
*/

func kClosestPointsToOriginV1(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return (points[i][0]*points[i][0] + points[i][1]*points[i][1]) < (points[j][0]*points[j][0] + points[j][1]*points[j][1])
	})
	return points[:k]
}

func kClosestPointsToOriginV2(points [][]int, k int) [][]int {
	slices.SortFunc(points, func(p1, p2 []int) int {
		return (p1[0]*p1[0] + p1[1]*p1[1]) - (p2[0]*p2[0] + p2[1]*p2[1])
	})
	return points[:k]
}

/* V3 with my implementation of priority queue. <<END */
type KClosestPointDistance struct {
	point    []int
	distance int
}

type KClosestHeap struct {
	array []KClosestPointDistance
	size  int
}

func NewKClosestHeap() *KClosestHeap {
	return &KClosestHeap{}
}

func (h *KClosestHeap) Push(value KClosestPointDistance) {
	h.array = append(h.array, value)
	h.size++
	if h.size > 1 {
		h.percolateUp(h.size - 1)
	}
}

func (h *KClosestHeap) Pop() KClosestPointDistance {
	if h.size == 0 {
		panic("Heap is empty")
	}
	max := h.array[0]
	h.array[0] = h.array[h.size-1]
	h.array = h.array[:h.size-1]
	h.size--
	h.percolateDown(0)
	return max
}

func (h *KClosestHeap) Pick() KClosestPointDistance {
	if h.size == 0 {
		panic("Heap is empty")
	}
	return h.array[0]
}

func (h *KClosestHeap) Array() []KClosestPointDistance {
	return h.array
}

func (h *KClosestHeap) percolateUp(i int) {
	pi := (i - 1) / 2 // Parent index
	for i > 0 && h.array[i].distance > h.array[pi].distance {
		h.array[i], h.array[pi] = h.array[pi], h.array[i]
		i = pi
		pi = (i - 1) / 2
	}
}

func (h *KClosestHeap) percolateDown(i int) {
	for {
		lci := 2*i + 1 // left child index
		rci := 2*i + 2 // right child index
		gi := i        // greatest index

		if lci < h.size && h.array[lci].distance > h.array[gi].distance {
			gi = lci
		}
		if rci < h.size && h.array[rci].distance > h.array[gi].distance {
			gi = rci
		}
		if gi == i {
			break
		}

		h.array[i], h.array[gi] = h.array[gi], h.array[i]
		i = gi
	}
}

func kClosestPointsToOriginV3(points [][]int, k int) [][]int {
	h := NewKClosestHeap()

	for _, point := range points {
		distance := point[0]*point[0] + point[1]*point[1]
		if h.size < k {
			h.Push(KClosestPointDistance{point, distance})
			continue
		}
		if distance < h.Pick().distance {
			h.Pop()
			h.Push(KClosestPointDistance{point, distance})
		}
	}

	res := make([][]int, k)
	for i, pd := range h.Array() {
		res[i] = pd.point
	}

	return res
}

/* END */

/* Top rated with Go implmentation of priority queue. <<END */
func kClosestDistance(point []int) int { return point[0]*point[0] + point[1]*point[1] }

type kClosestMaxHeap [][]int

func (h kClosestMaxHeap) Len() int           { return len(h) }
func (h kClosestMaxHeap) Less(i, j int) bool { return kClosestDistance(h[i]) > kClosestDistance(h[j]) }
func (h kClosestMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *kClosestMaxHeap) Push(x any)        { *h = append(*h, x.([]int)) }
func (h *kClosestMaxHeap) Pop() any {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func kClosestPointsToOriginTopRated(points [][]int, k int) [][]int {
	max := kClosestMaxHeap{}
	for _, point := range points {
		heap.Push(&max, point)
		if len(max) > k {
			heap.Pop(&max)
		}
	}
	return max
}

/* END */

func kClosestPointsToOriginTopRatedV2(points [][]int, k int) [][]int {
	square := func(n int) int {
		return n * n
	}

	for i := 1; i < len(points); i++ {
		if i >= 1 {
			distanceOfIMinusOne := square(points[i-1][0]) + square(points[i-1][1])
			distanceOfI := square(points[i][0]) + square(points[i][1])
			if distanceOfIMinusOne > distanceOfI {
				points[i-1], points[i] = points[i], points[i-1]
				i -= 2
			}
		}
	}
	return points[:k]
}
