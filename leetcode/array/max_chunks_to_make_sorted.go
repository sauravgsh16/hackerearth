package main

func maxChunksToSorted(arr []int) int {
	// idea is to maintain a max array.
	// Since the arr elements are [0, 1, .... (n-1)], where n is lenght of array
	// We compare the element in the max array with the index.
	// If element is same as index, we increment count.

	n := len(arr)
	max := make([]int, n)

	max[0] = arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < max[i-1] {
			max[i] = max[i-1]
		} else {
			max[i] = arr[i]
		}
	}

	var result int
	for i := 0; i < n; i++ {
		if max[i] == i {
			result++
		}
	}
	return result
}

func maxChunksToSortedShorter(arr []int) int {
	n := len(arr)
	max := 0
	result := 0

	for i := 0; i < n; i++ {
		if max < arr[i] {
			max = arr[i]
		}

		if max == i {
			result++
		}
	}

	return result
}
