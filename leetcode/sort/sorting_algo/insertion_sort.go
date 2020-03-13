package main

/*
Time Complexity: O(n*2)

Auxiliary Space: O(1)

Boundary Cases: Insertion sort takes maximum time to sort if elements are sorted in reverse order.
	        And it takes minimum time (Order of n) when elements are already sorted.

Algorithmic Paradigm: Incremental Approach

Sorting In Place: Yes

Stable: Yes

Online: Yes

Uses: Insertion sort is used when number of elements is small.
It can also be useful when input array is almost sorted,
only few elements are misplaced in complete big array.
*/

func insertionSort(arr []int) []int {
	n := len(arr)

	for i := 1; i < n; i++ {

		key := arr[i]
		j := i - 1

		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}

	return arr
}
