package main

import "fmt"

func main() {
	//b := []int{2, 5}
	a := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}
	fmt.Printf("%v\n", diagonal(a))
}
