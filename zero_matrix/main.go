package main

/**
https://leetcode-cn.com/problems/zero-matrix-lcci/
*/
// 普通方式，时间复杂度为 O(MxN) 空间复杂度为 O(n)
func setZeroes(matrix [][]int) {
	colKeys := make(map[int]struct{})
	rowKeys := make(map[int]struct{})

	rowNum := len(matrix)
	if rowNum == 0 {
		return
	}

	colNum := len(matrix[0])

	for i, row := range matrix {
		for j, elem := range row {
			if elem == 0 {
				colKeys[j] = struct{}{}
				rowKeys[i] = struct{}{}
			}
		}
	}

	for row, _ := range rowKeys {
		for i := 0; i < colNum; i++ {
			matrix[row][i] = 0
		}
	}

	for col, _ := range colKeys {
		for i := 0; i < rowNum; i++ {
			matrix[i][col] = 0
		}
	}
}

// 使用行列中的第一个元素作为标识符的特殊做法
func setZeroesBetter(matrix [][]int) {
	isFirstRowZero := false
	isFirstColZero := false

	rowLen := len(matrix)
	if rowLen == 0 {
		return
	}
	colLen := len(matrix[0])

	if colLen == 0 {
		return
	}

	for i, row := range matrix {
		// 展开的 row
		for j, elem := range row {
			// 对第一行做特殊处理，因为被用于标识符，这些行在后续扫描中不能正确的被判断
			if elem == 0 && i == 0 {
				isFirstRowZero = true
			}
			if elem == 0 && j == 0 {
				isFirstColZero = true
			}
			// 零作为标记
			if elem == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 行扫荡：标记行对应全置 0
	// 第一行先不变 最后处理，不然列处理就 GG
	for i := 1; i < rowLen; i++ {
		row := matrix[i]
		if row[0] == 0 {
			for j, _ := range row {
				matrix[i][j] = 0
			}
		}
	}

	// 列扫荡：标记列对应全置 0
	// 第一列先不变 最后处理，不然列处理就 GG
	for j := 1; j < colLen; j++ {
		elem := matrix[0][j]
		if elem == 0 {
			for i, _ := range matrix {
				if i == 0 {
					continue
				}
				matrix[i][j] = 0
			}
		}
	}

	// 特殊处理
	if isFirstRowZero {
		for i := 0; i < colLen; i++ {
			matrix[0][i] = 0
		}
	}

	if isFirstColZero {
		for i := 0; i < rowLen; i++ {
			matrix[i][0] = 0
		}
	}
}
