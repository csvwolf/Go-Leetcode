package main

/**
  https://leetcode-cn.com/problems/find-pivot-index/
*/
func pivotIndex(nums []int) int {
	left := 0
	sum := 0

	// 计算总和
	for _, num := range nums {
		sum += num
	}

	for index, num := range nums {
		if left == sum-left-num {
			return index
		}
		left += num
	}
	return -1
}
