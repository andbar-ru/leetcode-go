package leetcode

/*
57. Insert Interval

You are given an array of non-overlapping intervals intervals where intervals[i] = [start_i, end_i] represent the start and the end of the i^th interval and intervals is sorted in ascending order by start_i. You are also given an interval newInterval = [start, end] that represents the start and end of another interval.

Insert newInterval into intervals such that intervals is still sorted in ascending order by start_i and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).

Return intervals after the insertion.

Example 1:
Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]

Example 2:
Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].

Constraints:
* 0 <= intervals.length <= 10^4
* intervals[i].length == 2
* 0 <= start_i <= end_i <= 10^5
* intervals is sorted by start_i in ascending order.
* newInterval.length == 2
* 0 <= start <= end <= 10^5
*/

func insertInterval(intervals [][]int, newInterval []int) [][]int {
	start, end := newInterval[0], newInterval[1]
	length := len(intervals)

	// Special cases
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	if end < intervals[0][0] {
		return append([][]int{newInterval}, intervals...)
	}
	if start > intervals[length-1][1] {
		return append(intervals, newInterval)
	}

	startIntervalIndex := 0
	for start > intervals[startIntervalIndex][1] {
		startIntervalIndex++
	}

	endIntervalIndex := startIntervalIndex
	for endIntervalIndex < length-1 && end >= intervals[endIntervalIndex+1][0] {
		endIntervalIndex++
	}

	// Special case: new interval between existing
	if startIntervalIndex == endIntervalIndex && end < intervals[startIntervalIndex][0] {
		return append(intervals[:startIntervalIndex], append([][]int{{start, end}}, intervals[startIntervalIndex:]...)...)
	}

	if endIntervalIndex > startIntervalIndex {
		intervals[startIntervalIndex][1] = intervals[endIntervalIndex][1]
		intervals = append(intervals[:startIntervalIndex+1], intervals[endIntervalIndex+1:]...)
	}

	if start < intervals[startIntervalIndex][0] {
		intervals[startIntervalIndex][0] = start
	}
	if end > intervals[startIntervalIndex][1] {
		intervals[startIntervalIndex][1] = end
	}

	return intervals
}

func insertIntervalTopRated(intervals [][]int, newInterval []int) [][]int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	res := make([][]int, 0, len(intervals)+1)
	i := 0

	for ; i < len(intervals) && intervals[i][1] <= newInterval[0]; i++ {
		res = append(res, intervals[i])
	}

	for ; i < len(intervals) && intervals[i][0] <= newInterval[1]; i++ {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
	}

	res = append(res, newInterval)

	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}

	return res
}
