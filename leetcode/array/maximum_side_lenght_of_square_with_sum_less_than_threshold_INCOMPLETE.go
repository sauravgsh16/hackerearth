package main

func maxSideLenght(mat [][]int, threshold int) {
	m := len(mat)
	if m == 0 {
		return
	}
	n := len(mat[0])

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
		}
	}
}
