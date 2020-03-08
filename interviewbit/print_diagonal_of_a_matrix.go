package main

func diagonal(matrix [][]int) [][]int {
	result := [][]int{}
	n := len(matrix)

	// first we print the upper half diagonals.
	for i := 0; i < n; i++ {
		row := 0
		col := i

		temp := []int{}
		for col >= 0 {
			temp = append(temp, matrix[row][col])
			row++
			col--
		}
		result = append(result, temp)
	}

	// we now print the lower half from the main diagonal
	for i := 1; i < n; i++ {
		row := i
		col := n - 1

		temp := []int{}
		for row < n {
			temp = append(temp, matrix[row][col])
			row++
			col--
		}
		result = append(result, temp)
	}

	return result
}
