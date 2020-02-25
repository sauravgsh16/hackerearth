package main

import (
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
	n := len(deck)
	res := make([]int, n)

	sort.Ints(deck)

	for len(deck) > 0 {
		// mid := low + (high-low)/2

	}
}
