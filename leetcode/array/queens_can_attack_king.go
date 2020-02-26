package main

import "fmt"

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	visited := make([][]bool, 8)
	for i := range visited {
		visited[i] = make([]bool, 8)
	}

	// k1, k2 := king[0], king[1]
	for _, queen := range queens {
		x, y := queen[0], queen[1]
		visited[x][y] = true
	}

	fmt.Printf("%v\n", visited)

	result := [][]int{}
	direction := []int{-1, 0, 1}

	for _, dx := range direction {
		for _, dy := range direction {

			if dx == 0 && dy == 0 {
				continue
			}

			x, y := king[0], king[1]

			for dx+x >= 0 && dx+x < 8 && dy+y >= 0 && dy+y < 8 {
				x += dx
				y += dy

				if visited[x][y] == true {
					result = append(result, []int{x, y})
				}
			}
		}
	}

	return result
}
