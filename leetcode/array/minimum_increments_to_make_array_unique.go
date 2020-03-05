package main

import "sort"

/*
	A = [1, 1, 2, 2, 3, 7]

	A = [1,2,2,2,3,7]
	res= 1

	A = [1,2,3,2,3,7]
	res= 2

	A = [1,2,3,4,3,7]
	res= 4

	A = [1,2,3,4,5,7]
	res= 6
*/
func minIncrementForUnique(arr []int) int {
	n := len(arr)
	count := 0

	sort.Ints(arr)

	for i := 1; i < n; i++ {
		if arr[i] <= arr[i-1] {
			// count the difference between the previous number and the current number
			count += arr[i-1] + 1 - arr[i]
			// increment the current number to the next
			arr[i] = arr[i-1] + 1
		}
	}

	return count
}
