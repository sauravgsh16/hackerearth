package main

import (
	"math"
	"sort"
)

func threeSumClosest(arr []int, target int) int {
	n := len(arr)

	if n < 3 {
		return 0
	}

	closest := 99999999
	sort.Ints(arr)

	for i := 0; i < n-2; i++ {

		j := i + 1
		k := n - 1

		for j < k {
			sum := arr[i] + arr[j] + arr[k]

			if sum == target {
				return sum
			}

			if int(math.Abs(float64(target-sum))) < int(math.Abs(float64(target-closest))) {
				closest = sum
			}

			if sum > target {
				k--
			} else if sum < target {
				j++
			}
		}
	}
	return closest
}
