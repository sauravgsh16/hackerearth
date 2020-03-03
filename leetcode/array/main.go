package main

import "fmt"

func main() {
	//a := []int{3, 3, 2, 2, 1, 1}
	votes := []string{"ABC", "ACB", "ABC", "ACB", "ACB"}
	fmt.Printf("%s\n", rankTeams(votes))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
