package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	n := len(cost)

	if n == 0 {
		return 0
	}
	if n == 1 {
		return cost[0]
	}

	dp := make([]int, n)

	dp[0], dp[1] = cost[0], cost[1]

	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-2]+cost[i], dp[i-1]+cost[i])
	}

	return min(dp[n-2], dp[n-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isToeplitzMatrix(matrix [][]int) bool {
	r, c := len(matrix), len(matrix[0])

	for i := 0; i < r-1; i++ {
		for j := 0; j < c-1; j++ {

			fmt.Printf("%d, %d --> %d %d\n", i, j, i+1, j+1)

			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true
}
