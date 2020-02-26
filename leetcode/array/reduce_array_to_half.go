package main

import (
	"sort"
)

func minSetSize(arr []int) int {
	sort.Ints(arr)

	freq := []int{}
	count := 1

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			count++
		} else {
			freq = append(freq, count)
			count = 1
		}
	}
	// last element
	freq = append(freq, count)

	// reverse sort freq
	sort.Sort(sort.Reverse(sort.IntSlice(freq)))

	sum := 0
	res := 0
	for sum < len(arr)/2 {
		sum += freq[res]
		res++
	}

	return res
}
