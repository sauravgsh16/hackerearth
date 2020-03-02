package main

func maxUncrossedLines(a, b []int) int {
	// Variation of Longest Common Subsequence
	// eg :- 2 5 1 2 5
	//       10 5 2 1 5 2
	// If we find the longest common subsequence- (5, 2, 5) or (5 , 1, 5),
	// we solve the problem

	m := len(a)
	n := len(b)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[m][n]
}
