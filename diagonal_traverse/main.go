package main

/**
https://leetcode-cn.com/problems/diagonal-traverse/
*/
func findDiagonalOrder(matrix [][]int) []int {
	curRow := 0
	curCol := 0
	rows := len(matrix)
	if rows == 0 {
		return []int{}
	}

	cols := len(matrix[0])
	if cols == 0 {
		return []int{}
	}

	result := make([]int, rows*cols)
	currentId := 0
	// 向上走的线
	up := true

	// 如果到最后一个元素证明可以结束遍历了
	for curRow <= rows-1 && curCol <= cols-1 {
		result[currentId] = matrix[curRow][curCol]
		// 为下一个元素做准备拉
		currentId++

		// 如果是顶层元素，得拐弯了
		if up && (curRow == 0 || curCol == cols-1) {
			// 再加 col 要溢出了
			if curCol == cols-1 {
				curRow++
			} else {
				curCol++
			}
			// 走不上去拉向下吧
			up = false
		} else if up {
			curCol++
			curRow--
		} else if !up && (curCol == 0 || curRow == rows-1) {
			if curRow == rows-1 {
				// 再加 row 溢出
				curCol += 1
			} else {
				curRow++
			}
			// 走不下去了向上看
			up = true
		} else if !up {
			curRow++
			curCol--
		}
	}
	return result
}
