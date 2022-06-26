package algorithm

import "sort"

func Merge(intervals [][]int) [][]int {
	// 排序，按数组第一个点的值，顺序排序
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}
	start := intervals[0][0]
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if (intervals[i][0] >= start && intervals[i][0] <= end) ||
			(intervals[i][1] >= start && intervals[i][1] <= end) ||
			(intervals[i][0] < start && intervals[i][1] > end) {
			start = min(start, intervals[i][0])
			end = max(end, intervals[i][1])
		} else {
			result = append(result, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		}
	}
	result = append(result, []int{start, end})
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
