package main

import "fmt"

func main() {
	/*

		fmt.Printf("%v\n", removeSubFolders(a))
	*/
	a := []int{9, 10, 4, 5}
	b := []int{1, 0, 1, 1}
	fmt.Printf("%d\n", maxSatisfied(a, b, 1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
