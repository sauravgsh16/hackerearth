package main

import "fmt"

func maxProfitWithFee(prices []int, fee int) {
	minPrice := 100000
	n := len(prices)

	profit := []int{}

	for i := 0; i < n; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			profit = append(profit, prices[i]-minPrice-fee)
			minPrice = prices[i]
		}
	}

	s := 0
	for _, p := range profit {
		s += p
	}
	fmt.Printf("%d\n", s)
}
