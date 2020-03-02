package main

// 2, 30, 15, 10, 8, 25, 80
func maxProfit(prices []int) int {
	n := len(prices)
	minPrice := 1000000
	profit := []int{}

	for i := 0; i < n; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			profit = append(profit, prices[i]-minPrice)
			minPrice = prices[i]
		}
	}
	s := 0
	for i := 0; i < len(profit); i++ {
		s += profit[i]
	}
	return s
}
