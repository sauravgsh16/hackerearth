package main

import "fmt"

func main() {
	a := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}

	fmt.Printf("%v\n", minPathSum(a))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
