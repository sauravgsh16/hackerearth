package main

func findLengthOfLCIS(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	count, max := 1, 1

	for i := 0; i < n-1; i++ {
		if arr[i] < arr[i+1] {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 1
		}
	}
	return max
}
