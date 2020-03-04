package main

func lenLongestFibSubseq(arr []int) int {
	n := len(arr)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	max := 0

	for i := 1; i < n; i++ {
		l, r := 0, i-1
		for l < r {
			sum := arr[l] + arr[r]
			if sum > arr[i] {
				r--
			} else if sum < arr[i] {
				l++
			} else {
				dp[r][i] = dp[l][r] + 1

				if max < dp[r][i] {
					max = dp[r][i]
				}
				l++
				r--
			}
		}
	}

	if max == 0 {
		return 0
	}
	return max + 2
}
