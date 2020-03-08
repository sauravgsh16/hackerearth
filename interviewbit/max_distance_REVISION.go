/*
Given an array A of integers, find the maximum of j - i subjected to the constraint of A[i] <= A[j].

If there is no solution possible, return -1.

Example :

A : [3 5 4 2]

Output : 2
for the pair (3, 4)
*/

package main

import "fmt"

func maxDistance(arr []int) int {
	n := len(arr)

	if n == 0 {
		return -1
	}
	if n == 1 {
		return 0
	}

	lmin := []int{}
	rmax := make([]int, n)

	min := arr[0]
	max := arr[n-1]

	for i := 0; i < n; i++ {
		if arr[i] <= min {
			min = arr[i]
		}
		lmin = append(lmin, min)
	}

	for i := n - 1; i >= 0; i-- {
		if arr[i] >= max {
			max = arr[i]
		}
		rmax[i] = max
	}

	sol := -1
	i, j := 0, 0
	for i < n && j < n {
		if lmin[i] <= rmax[j] {
			if sol < j-i {
				sol = j - i
			}
			j++
		} else {
			i++
		}
	}

	fmt.Printf("%v\n", lmin)
	fmt.Printf("%v\n", rmax)
	/*
		[3 3 3 2]
		[5 5 4 2]
		2
	*/

	return sol
}
