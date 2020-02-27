package main

import (
	"math"
)

func subsets(arr []int) [][]int {
	n := len(arr)
	size := int(math.Pow(float64(2), float64(n)))

	result := [][]int{}

	for counter := 0; counter < size; counter++ {
		t := []int{}
		for j := 0; j < n; j++ {
			if counter&(1<<uint(j)) > 0 {
				t = append(t, arr[j])
			}
		}
		result = append(result, t)
	}

	return result
}
