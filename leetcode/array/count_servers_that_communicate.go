package main

func countServers(grid [][]int) int {
	r := len(grid)
	c := len(grid[0])

	rowCount := make([]int, r)
	colCount := make([]int, c)

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 1 {
				rowCount[i]++
				colCount[j]++
			}
		}
	}

	count := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 1 && (rowCount[i] > 1 || colCount[j] > 1) {
				count++
			}
		}
	}

	return count
}
