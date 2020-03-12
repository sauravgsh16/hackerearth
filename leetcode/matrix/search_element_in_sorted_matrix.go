package main

func findElement(matrix [][]int, k int) (int, int, bool) {
	n := len(matrix)
	m := len(matrix[0])

	i := n - 1
	j := 0

	for i >= 0 && j < m {
		val := matrix[i][j]

		if val == k {
			return i, j, true
		} else if val > k {
			i--
		} else {
			j++
		}
	}
	return 0, 0, false
}
