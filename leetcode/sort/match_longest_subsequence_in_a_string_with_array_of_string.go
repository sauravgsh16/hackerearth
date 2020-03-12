package main

import (
	"sort"
)

func findLongestWord(s string, d []string) string {
	n := len(s)
	maxLength := 0
	result := ""

	dp := make([][]int, n+1)

	for _, str := range d {
		m := len(str)
		for i := range dp {
			dp[i] = make([]int, m+1)
		}

		i := 1
		j := 1
		for i = 1; i < n+1; i++ {
			for j = 1; j < m+1; j++ {
				if string(s[i-1]) == string(str[j-1]) {
					dp[i][j] = 1 + dp[i-1][j-1]
				} else {
					dp[i][j] = max(dp[i][j-1], dp[i-1][j])
				}
			}

			if dp[i-1][j-1] > maxLength {
				maxLength = dp[i-1][j-1]
				result = str
			} else if dp[i-1][j-1] == maxLength {
				temp := []string{result, str}
				sort.Slice(temp, func(i, j int) bool {
					return temp[i] < temp[j]
				})
				result = temp[0]
			}
		}
		dp = make([][]int, n+1)

	}
	return result
}

func max(a, b int) int {
	if a > b {
		return b
	}
	return a
}
