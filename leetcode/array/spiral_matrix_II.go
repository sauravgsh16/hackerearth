package main

func generateSpiralMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	k, l := 0, 0
	m := n

	d := 1

	for k < m && l < n {

		// fill up first column of the remaining
		for i := l; i < n; i++ {
			matrix[k][i] = d
			d++
		}
		k++

		// fill up last column of the remaining
		for i := k; i < m; i++ {
			matrix[i][n-1] = d
			d++
		}
		n--

		// fill up last row of the remaining
		if k < m {
			for i := (n - 1); i > (l - 1); i-- {
				matrix[m-1][i] = d
				d++
			}
			m--
		}

		// fill up first column of the remaining
		if l < n {
			for i := (m - 1); i > (k - 1); i-- {
				matrix[i][l] = d
				d++
			}
			l++
		}
	}

	return matrix
}
