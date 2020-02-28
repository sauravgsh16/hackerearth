package main

func countSquares(arr [][]int) int {
	r := len(arr)
	c := len(arr[0])

	dp := make([][]int, r+1)
	for i := range dp {
		dp[i] = make([]int, c+1)
	}

	// Solution is as below:
	// if arr[i-1][j-1] == 1, then
	// if we have a square of value k, where (i, j) is the square just below it,
	// then, dp[i-1][j] is at least of size k-1, same is for dp[i][j-1] and dp[i-1][j-1].
	// Thus we take min of (dp[i-1][j-1], dp[i][j-1], dp[i-1][j])

	sum := 0
	for i := 1; i < r+1; i++ {
		for j := 1; j < c+1; j++ {
			if arr[i-1][j-1] == 1 {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
				sum += dp[i][j]
			}
		}
	}

	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
