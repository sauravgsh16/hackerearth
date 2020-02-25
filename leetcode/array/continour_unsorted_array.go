package main

import "fmt"

func finUnsortedSubarray(arr []int) int {
	n := len(arr)

	s, e := 0, n-1

	for s < n-1 {
		if arr[s] > arr[s+1] {
			break
		}
		s++
	}

	if s == n-1 {
		return 0
	}

	for e > 0 {
		if arr[e] < arr[e-1] {
			break
		}
		e--
	}

	max := arr[s]
	min := arr[s]

	for i := s + 1; i < e+1; i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}

	var i int
	// find element in [0... s-1], such that, it is
	// greater than min, change s to index of the element
	i = 0
	for i < s {
		if arr[i] > min {
			s = i
			break
		}
		i++
	}

	// find element in [e+1...n-1], such that it it
	// smaller than max, change e to index of the element
	i = n - 1
	for i >= e+1 {
		if arr[i] < max {
			e = i
			break
		}
		i--
	}

	fmt.Printf("%d, %d\n", e, s)

	return e - s + 1
}
