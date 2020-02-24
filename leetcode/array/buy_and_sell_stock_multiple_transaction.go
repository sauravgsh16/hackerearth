package main

// 2, 30, 15, 10, 8, 25, 80
func maxProfit(p []int) int {
	n := len(p)
	minPrice := 1000000
	profit := []int{}

	for i := 0; i < n; i++ {
		if p[i] < minPrice {
			minPrice = p[i]
		} else {
			profit = append(profit, p[i]-minPrice)
			minPrice = p[i]
		}
	}
	s := 0
	for i := 0; i < len(profit); i++ {
		s += profit[i]
	}
	return s
}
