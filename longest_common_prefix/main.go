package main

/**
https://leetcode-cn.com/problems/longest-common-prefix/
无论怎么二分和分治，都不会比扫描更省
*/
func longestCommonPrefix(strs []string) string {
	listLength := len(strs)

	if listLength == 0 {
		return ""
	}

	firstStr := strs[0]
	resultIdx := -1

	// 扫描法，顺便用了一次骚包 goto 请勿模仿
a:
	for index, char := range firstStr {
		for _, str := range strs {
			if len(str) <= index || rune(str[index]) != char {
				break a
			}
		}
		resultIdx = index
	}
	return firstStr[:resultIdx+1]
}
