package MergeIntervals

import "sort"

// https://leetcode-cn.com/problems/merge-intervals/
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)

	for index, interval := range intervals {
		p := len(result) - 1
		if index == 0 || interval[0] > result[p][1] {
			result = append(result, interval)
		} else {
			if interval[1] >= result[p][1] {
				result[p][1] = interval[1]
			}
		}
	}
	return result
}
