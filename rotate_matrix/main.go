package main

/**
https://leetcode-cn.com/problems/rotate-matrix-lcci/
*/
// 弱鸡写法
func rotate(matrix [][]int) {
	// 先来一个同等大小
	size := len(matrix)
	newMatrix := make([][]int, size)

	// 还要初始化
	for i, _ := range matrix {
		newMatrix[i] = make([]int, size)
	}

	// 开始转
	for i, line := range matrix {
		for j, elem := range line {
			newMatrix[j][size-i-1] = elem
		}
	}

	// 整回去
	for i, line := range matrix {
		for j, _ := range line {
			matrix[j][i] = newMatrix[j][i]
		}
	}
}

// 旋转法
func rotate2(matrix [][]int) {
	size := len(matrix)
	// 水平翻转数组
	for i := 0; i < size/2; i++ {
		for j := 0; j < size; j++ {
			matrix[i][j], matrix[size-i-1][j] = matrix[size-i-1][j], matrix[i][j]
		}
	}

	// 对角线转
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// 超级旋转法
func rotate3(matrix [][]int) {
	size := len(matrix)

	for i := 0; i < size-1; i++ {
		for j := i; j < size-i-1; j++ {
			matrix[i][j], matrix[size-1-j][i], matrix[size-i-1][size-1-j], matrix[j][size-1-i] =
				matrix[size-1-j][i], matrix[size-i-1][size-1-j], matrix[j][size-1-i], matrix[i][j]
		}
	}
}
