package main

func maxSubArray(arr []int) int {
	n := len(arr)
	maxSoFar := -99999999
	maxEndingHere := 0

	for i := 0; i < n; i++ {
		maxEndingHere = arr[i] + maxEndingHere

		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}

		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
	}
	return maxSoFar
}
