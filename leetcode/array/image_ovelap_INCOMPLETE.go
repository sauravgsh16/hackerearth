package main

import "fmt"

func largestOverlap(a, b [][]int) int {
	n := len(a)

	aCoord := [][]int{}
	bCoord := [][]int{}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] == 1 {
				aCoord = append(aCoord, []int{i, j})
			}

			if b[i][j] == 1 {
				bCoord = append(bCoord, []int{i, j})
			}
		}
	}

	max := 0

	hash := make(map[string]int)

	for _, ac := range aCoord {
		for _, bc := range bCoord {
			key := fmt.Sprintf("%d%d", (bc[0] - ac[0]), (bc[1] - ac[1]))

			if _, ok := hash[key]; !ok {
				hash[key] = 1
			} else {
				hash[key]++
			}

			if c, _ := hash[key]; c > max {
				max = c
			}
		}
	}

	for k, v := range hash {
		fmt.Printf("%s --> %d\n", k, v)
	}
	return max
}
