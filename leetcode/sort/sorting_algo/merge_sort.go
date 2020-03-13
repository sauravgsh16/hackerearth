/*
Time Complexity: Sorting arrays on different machines. Merge Sort is a recursive algorithm and time complexity can be expressed as following recurrence relation.
T(n) = 2T(n/2) + O(n)
The above recurrence can be solved either using Recurrence Tree method or Master method.
It falls in case II of Master Method and solution of the recurrence is O(nLogn).
Time complexity of Merge Sort is O(nLogn) in all 3 cases (worst, average and best)
as merge sort always divides the array into two halves and take linear time to merge two halves.

Auxiliary Space: O(n)

Algorithmic Paradigm: Divide and Conquer

Sorting In Place: No in a typical implementation

Stable: Yes

*/
package main

func mergeSort(arr []int) []int {
	n := len(arr)

	if n <= 1 {
		return arr
	}

	mid := n / 2

	left := arr[:mid]
	right := arr[mid:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	res := make([]int, len(left)+len(right))

	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res[k] = left[i]
			i++
		} else {
			res[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		res[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		res[k] = right[j]
		j++
		k++
	}

	return res
}
