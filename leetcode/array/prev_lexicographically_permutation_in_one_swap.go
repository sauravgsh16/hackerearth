package main

// We need to find the largest lexicographically permulation
// which is smaller than the given permutation, in just one swap
// eg : A = [3, 1, 1, 3], o/p -> [1, 3, 1, 3]
// eg : A = [3, 2, 1] o/p -> [3, 1, 2]
// eg : A = [1, 9, 4, 6, 7] -> [1, 7, 4, 6, 9]

func prevPermutation(arr []int) []int {
	n := len(arr)

	if n < 2 || isSorted(arr) {
		return arr
	}

	candidate1 := n - 2

	// find the point where, array is no longer sorted in ascending order
	for candidate1 >= 0 && arr[candidate1] <= arr[candidate1+1] {
		candidate1--
	}

	candidate2 := toSwap(arr, candidate1+1)

	swapelements(arr, candidate1, candidate2)
	return arr
}

func toSwap(arr []int, left int) int {
	right := len(arr) - 1
	val := arr[left-1]
	result := left

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] < val {
			if arr[result] < arr[mid] {
				result = mid
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

func swapelements(arr []int, a, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}

func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
