package main

import (
	"math"
)

func findDuplicates(arr []int) []int {
	n := len(arr)
	res := []int{}

	for i := 0; i < n; i++ {
		// get the absolute value, used to search for index
		val := math.Abs(float64(arr[i]))

		// need to substract 1, from value because of zero indexing
		if arr[int(val)-1] >= 0 {
			arr[int(val)-1] = -arr[int(val)-1]
		} else {
			res = append(res, int(val))
		}
	}

	return res
}
