package main

func pancakeSort(arr []int) []int {
	n := len(arr)

	currSize := n

	for currSize > 1 {

		max := findMax(arr, currSize)

		if max != currSize-1 {
			// To move at the end,
			// first move maximum number to beginning
			flip(arr, max)

			// Now move the maximum number to end by
			// reversing current array
			flip(arr, currSize-1)
		}

		currSize--
	}

	return arr
}

func findMax(arr []int, n int) int {
	var max int
	for i := 0; i < n; i++ {
		if arr[i] > arr[max] {
			max = i
		}
	}
	return max
}

func flip(arr []int, mx int) {
	start := 0

	for start < mx {
		temp := arr[start]
		arr[start] = arr[mx]
		arr[mx] = temp
		start++
		mx--
	}
}
