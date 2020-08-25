package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findSubsequences(nums []int) [][]int {
	// 去重用 Mapping
	mapping := make(map[string]struct{})
	length := len(nums)

	if length == 0 {
		return [][]int{}
	}
	for _, num := range nums {
		// 为了不让 mapping 不断增加，需要先找个临时存放地址
		tmp := make(map[string]struct{})
		for key := range mapping {
			result := strings.Split(key, ",")
			// 算法题，所以转换一定会成功，可以忽略 err
			n, _ := strconv.Atoi(string(result[len(result)-1]))
			if n <= num {
				tmp[fmt.Sprintf("%s,%d", key, num)] = struct{}{}
			}
		}
		// 整完了，放进去吧
		for key := range tmp {
			mapping[key] = struct{}{}
		}
		// 把自己记进去
		mapping[fmt.Sprintf("%d", num)] = struct{}{}
	}
	filtered := make([][]int, 0)
	// 去掉长度为 1 的
	for key := range mapping {
		result := strings.Split(key, ",")
		if len(result) != 1 {
			r := make([]int, len(result))
			for i, substring := range result {
				r[i], _ = strconv.Atoi(substring)
			}
			filtered = append(filtered, r)
		}
	}
	return filtered
}
