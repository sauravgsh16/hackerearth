package main

import "sort"

// Thing to remember her is the check for duplicate processing.
// i.e, we need to sort the array before finding subsets.

// Before calling the recursive function, we need to check,
// if `i` is greater than the starting index. If yes, we check whether
// `arr[i]` is same as `arr[i-1]`. If true, we skip the recursive call for that
// array element, as it would result in duplication, since we have already called
// the recursive function for the first instance of the element.

func subsetsWithDup(arr []int) [][]int {
	result := make([][]int, 0)

	// sort array before starting the process, so that we get all same element grouped together.
	sort.Ints(arr)
	temp := []int{}
	dfsSubsetWithOutDup(arr, &result, temp, 0)

	return result
}

func dfsSubsetWithOutDup(arr []int, result *[][]int, path []int, start int) {
	r := append([]int{}, path...)
	*result = append(*result, r)

	for i := start; i < len(arr); i++ {
		if i > start && arr[i] == arr[i-1] {
			continue
		}

		temp := append(path, arr[i])
		dfsSubsetWithOutDup(arr, result, temp, i+1)
	}
}
