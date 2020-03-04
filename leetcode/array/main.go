package main

import "fmt"

func main() {
	a := []int{3, 1, 1, 3}
	// votes := []string{"ABC", "ACB", "ABC", "ACB", "ACB"}
	fmt.Printf("%v\n", prevPermutation(a))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
