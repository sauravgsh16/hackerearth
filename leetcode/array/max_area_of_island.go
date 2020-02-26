package main

func dfs(i, j, r, c int, grid [][]int, visited *[][]bool) int {
	if i < 0 || i >= r || j < 0 || j >= c {
		return 0
	}

	if (*visited)[i][j] == true {
		return 0
	}
	(*visited)[i][j] = true

	if grid[i][j] == 0 {
		return 0
	}

	// Recur for all 4 directions of current location:
	// ie. (i-1, j), (i+1, j), (i, j-1), (i, j+1)
	return dfs(i-1, j, r, c, grid, visited) + dfs(i+1, j, r, c, grid, visited) +
		dfs(i, j-1, r, c, grid, visited) + dfs(i, j+1, r, c, grid, visited) + 1
}

func maxAreaOfIsland(grid [][]int) int {
	r := len(grid)
	c := len(grid[0])

	// create visited array
	visited := make([][]bool, r)
	for i := range visited {
		visited[i] = make([]bool, c)
	}

	maxArea := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			// if not visited and grid value is 1
			if !visited[i][j] && grid[i][j] == 1 {
				c := dfs(i, j, r, c, grid, &visited)
				if c > maxArea {
					maxArea = c
				}
			}
		}
	}
	return maxArea
}
