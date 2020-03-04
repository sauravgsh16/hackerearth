package main

import "sort"

func isPossibleDivide(arr []int, q int) bool {
	n := len(arr)
	hash := make(map[int]int, 0)

	for i := 0; i < n; i++ {
		if _, ok := hash[arr[i]]; ok {
			hash[arr[i]]++
		} else {
			hash[arr[i]] = 1
		}
	}

	keys := []int{}
	for k := range hash {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		if hash[k] > 0 {
			minus := hash[k]
			for i := k; i < k+q; i++ {
				if hash[i] < minus {
					return false
				}
				hash[i] -= minus
			}
		}
	}

	return true
}
