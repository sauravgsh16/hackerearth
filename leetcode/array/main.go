package main

import "fmt"

func main() {
	/*
		a1 := []int{1, 2, 3, 0, 0, 0, 0}
		a2 := []int{2, 4, 6, 7}
	*/
	//fmt.Printf("%d\n", missingNumber([]int{2}))
	a := []int{3, 1, 0, 1, 5}
	//a := []int{0, 0, 1, 0, 0}
	fmt.Printf("%d\n", findPairs(a, 2))
}
