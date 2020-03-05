package main

import "sort"

// Combination sum - without any duplicated present.
// The input array can have duplicates
// Key here is to sort the array, so that when we select i, we can escape having the same combinations.

func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	temp := []int{}

	sort.Ints(candidates)
	backTrack2(candidates, temp, &result, target, 0)

	return result
}

func backTrack2(candidates []int, temp []int, result *[][]int, remain, curr int) {
	if remain < 0 {
		return
	}

	if remain == 0 {
		r := append([]int{}, temp...)
		*result = append(*result, r)
		return
	}

	for i := curr; i < len(candidates); i++ {
		// check for duplicates, so that we avoid repeating the combinations.
		if i > curr && candidates[i] == candidates[i-1] {
			continue
		}

		t := append(temp, candidates[i])
		// **We recur using i+1, since we can select one number only once
		backTrack2(candidates, t, result, remain-candidates[i], i+1)
	}
}
