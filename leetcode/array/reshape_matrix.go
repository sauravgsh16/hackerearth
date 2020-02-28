package main

func reshape(arr [][]int, r, c int) [][]int {
	or := len(arr)
	oc := len(arr[0])

	total := or * oc

	// return original if the total number does not match
	if total != r*c {
		return arr
	}

	result := make([][]int, r)
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
	}

	for i := 0; i < total; i++ {
		oldRow := i / oc
		oldCol := i % oc
		newRow := i / c
		newCol := i % c

		result[newRow][newCol] = arr[oldRow][oldCol]
	}

	return result
}
