package main

import "fmt"

func buyAndSellStock(prices []int) int {
	/*
		Using Kadane's algorithm
	*/

	minBuying := 10000000
	maxProfit := 0
	n := len(prices)

	for i := 0; i < n; i++ {
		current := prices[i]

		if current < minBuying {
			minBuying = current
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
