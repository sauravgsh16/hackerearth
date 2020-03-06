package main

import (
	"math"
)

func repeatedNumber(arr []int) []int {
	n := len(arr)
	sum := 0

	repeated := 0

	for i := 0; i < n; i++ {
		idx := int(math.Abs(float64(arr[i]))) - 1
		if arr[idx] > 0 {
			arr[idx] = -arr[idx]
		} else {
			repeated = int(math.Abs(float64(arr[i])))
		}

		val := int(math.Abs(float64(arr[i])))
		sum += val
	}

	expected := n * (n + 1) / 2
	missing := expected - (sum - repeated)

	ans := []int{repeated, missing}

	return ans
}
