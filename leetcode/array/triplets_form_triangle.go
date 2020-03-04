package main

import (
	"sort"
)

// We solve this using the 3 sum algorithm
// We keep one pointer fixed and move the remaining two to find the combinations
func triangleNumber(arr []int) int {
	n := len(arr)
	count := 0

	sort.Ints(arr)

	for k := n - 1; k > 1; k-- {
		j := k - 1
		i := 0
		for i < j {
			if (arr[i] + arr[j]) > arr[k] {
				count += j - i
				j--
			} else {
				i++
			}
		}
	}

	return count
}
