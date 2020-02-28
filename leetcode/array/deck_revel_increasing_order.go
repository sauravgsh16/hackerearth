package main

import (
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
	// Skip logic
	n := len(deck)

	sort.Ints(deck)
	result := make([]int, n)
	j := 0
	result[0] = deck[0]

	for i := 1; i < n; i++ {
		skip := false
		for !skip {
			if result[j] == 0 {
				skip = true
			}
			j++
			if j > n-1 {
				j = 0
			}
		}

		for result[j] != 0 {
			j++
			if j > n-1 {
				j = 0
			}
		}
		result[j] = deck[i]
	}
	return result
}
