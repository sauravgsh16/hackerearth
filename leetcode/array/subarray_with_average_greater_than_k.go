package main

func numOfSubarrays(arr []int, k, threshold int) int {
	n := len(arr)
	res := 0
	sum := 0

	var i int
	for i = 0; i < k; i++ {
		sum += arr[i]
	}

	if sum/k >= threshold {
		res++
	}

	for j := i; j < n; j++ {
		sum = sum - arr[j-k] + arr[j]
		if (sum / k) >= threshold {
			res++
		}
	}

	return res
}
