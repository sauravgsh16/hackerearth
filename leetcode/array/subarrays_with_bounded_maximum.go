package main

import (
	"fmt"
	"sort"
)

func numSubarrayBoundedMax(arr []int, l, r int) int {
	i := 0

	for i < len(arr) {
		if arr[i] < l || arr[i] > r {
			arr = append(arr[:i], arr[i+1:]...)
			continue
		}
		i++
	}

	result := make([][]int, 0)
	sort.Ints(arr)

	dfsboundedSubsets(arr, &result, []int{}, 0)

	fmt.Printf("%v\n", result)
	return len(result)
}

func dfsboundedSubsets(arr []int, result *[][]int, path []int, start int) {
	if len(path) != 0 {
		*result = append(*result, path)
	}

	for i := start; i < len(arr); i++ {
		if i > start && arr[i] == arr[i-1] {
			continue
		}

		temp := append(path, arr[i])
		dfsboundedSubsets(arr, result, temp, i+1)
	}
}
