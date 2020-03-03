package main

// Given an array, print matrix in spiral order

func spiralOrder(matrix [][]int) []int {
	// We use four loops
	// First loop represents movement from left to right
	// Second loop represents movement from top to bottom
	// Third loop represents movement from right to left
	// Fourth loop represents movement from bottom to top

	// k - starting row index
	// m - ending row index
	// l - starting column index
	// n - ending column index

	k, l := 0, 0
	m := len(matrix)

	if m == 0 {
		return []int{}
	}

	n := len(matrix[0])

	result := []int{}

	for k < m && l < n {

		// Print first row from remaining rows
		for i := l; i < n; i++ {
			result = append(result, matrix[k][i])
		}
		k++

		// Print last column from the remaining columns
		for i := k; i < m; i++ {
			result = append(result, matrix[i][n-1])
		}
		n--

		// Print last row from the remaining rows
		if k < m {
			for i := (n - 1); i > (l - 1); i-- {
				// last index (m-1)
				result = append(result, matrix[m-1][i])
			}
			m--
		}

		if l < n {
			for i := (m - 1); i > (k - 1); i-- {
				// last index(l)
				result = append(result, matrix[i][l])
			}
			l++
		}
	}
	return result
}
