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

func subsetRecursive(arr []int) [][]int {
	result := make([][]int, 0)

	temp := []int{}
	dfsSubset(arr, &result, temp, 0)
	return result
}

func dfsSubset(arr []int, result *[][]int, path []int, start int) {
	r := append([]int{}, path...)
	*result = append(*result, r)

	for i := start; i < len(arr); i++ {
		t := append(path, arr[i])
		dfsSubset(arr, result, t, i+1)
	}
}
