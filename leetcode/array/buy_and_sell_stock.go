package main

import "fmt"

func buyAndSellStock(p []int) int {
	/*
		Using Kadane's algorithm
	*/

	minBuying := 10000000
	maxProfit := 0
	n := len(p)

	for i := 0; i < n; i++ {
		current := p[i]

		if p[i] < minBuying {
			minBuying = p[i]
		}

		var profit int
		if current > minBuying {
			profit = current - minBuying
		}

		if profit > maxProfit {
			maxProfit = profit
		}
	}

	fmt.Printf("%d\n", maxProfit)
	return maxProfit
}
