package main

func findLength(a, b []int) int {
	m := len(a)
	n := len(b)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	max := 0
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			}
			if dp[i][j] > max {
				max = dp[i][j]
			}

		}
	}

	return max
}
