package main

func findMaxAverage(arr []int, k int) float64 {
	sum, max := 0, 0
	i, n := 0, len(arr)

	for i = 0; i < k; i++ {
		sum += arr[i]
	}

	max = sum

	for i = k; i < n; i++ {
		sum = sum - arr[i-k] + arr[i]
		if sum > max {
			max = sum
		}
	}
	return float64(max) / float64(k)
}
