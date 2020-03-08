package main

import "fmt"

/*
	Start from the last Index = (m + n) - 1
	Fill array from last
*/

func mergeSortedArray(arr1, arr2 []int, m, n int) {
	i, j := m-1, n-1
	lastIndex := m + n - 1

	for j >= 0 {
		if i >= 0 && arr1[i] > arr2[j] {
			arr1[lastIndex] = arr1[i]
			i--
		} else {
			arr1[lastIndex] = arr2[j]
			j--
		}
		lastIndex--
	}

	for i >= 0 {
		arr1[lastIndex] = arr1[i]
		i--
		lastIndex--
	}

	for j >= 0 {
		arr1[lastIndex] = arr2[j]
		j++
		lastIndex--
	}

	fmt.Printf("%v\n", arr1)
}
