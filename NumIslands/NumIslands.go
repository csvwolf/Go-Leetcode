package NumIslands

func NumIslands(grid [][]byte) int {
	var count int
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == '1' {
				traverse(grid, i, j)
				count++
			}
		}
	}
	return count
}

func traverse(grid [][]byte, i int, j int) {
	grid[i][j] = '0'
	if i+1 < len(grid) && grid[i+1][j] == '1' {
		traverse(grid, i+1, j)
	}
	if i-1 >= 0 && grid[i-1][j] == '1' {
		traverse(grid, i-1, j)
	}
	if j+1 < len(grid[i]) && grid[i][j+1] == '1' {
		traverse(grid, i, j+1)
	}

	if j-1 >= 0 && grid[i][j-1] == '1' {
		traverse(grid, i, j-1)
	}
}
