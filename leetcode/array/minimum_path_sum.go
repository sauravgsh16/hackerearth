package main

func minPathSum(arr [][]int) int {
	m := len(arr)
	n := len(arr[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// fill first cell
	dp[0][0] = arr[0][0]

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if i > 0 && j > 0 {
				dp[i][j] = arr[i][j] + min(dp[i][j-1], dp[i-1][j])

			} else if i > 0 {
				dp[i][j] = arr[i][j] + dp[i-1][j]

			} else if j > 0 {
				dp[i][j] = arr[i][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
