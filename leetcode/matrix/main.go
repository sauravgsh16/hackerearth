package main

import "fmt"

func main() {
	m := [][]int{
		[]int{2, 5, 7},
		[]int{4, 8, 12},
		[]int{9, 11, 15},
		[]int{14, 17, 20},
	}

	i, j, k := findElement(m, 12)
	fmt.Printf("%d, %d, %t\n", i, j, k)
}
