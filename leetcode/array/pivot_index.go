package main

func pivotIndex(arr []int) int {
	n := len(arr)
	var left, right int

	for i := 0; i < n; i++ {
		right += arr[i]
	}

	for i := 0; i < n; i++ {
		right -= arr[i]
		if left == right {
			return i
		}
		left += arr[i]
	}
	return -1
}
