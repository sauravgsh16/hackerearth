package main

import "fmt"

const (
	die  = 2
	live = 3
)

func gameOfLife(board [][]int) {
	r := len(board)
	c := len(board[0])

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			neighbours := countNeighbours(board, i, j, r, c)

			if board[i][j] == 0 && neighbours == live {
				board[i][j] = live
			} else if board[i][j] == 1 {
				if neighbours == 2 || neighbours == 3 {
					continue
				}
				if neighbours < 2 || neighbours > 3 {
					board[i][j] = die
				}
			}

		}
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {

			switch board[i][j] {
			case die:
				board[i][j] = 0
			case live:
				board[i][j] = 1
			}
		}
	}

	fmt.Printf("%v\n", board)
}

func countNeighbours(board [][]int, i, j, r, c int) int {
	directions := [][]int{
		[]int{1, 0},
		[]int{-1, 0},
		[]int{0, 1},
		[]int{0, -1},
		[]int{1, 1},
		[]int{1, -1},
		[]int{-1, 1},
		[]int{-1, -1},
	}

	count := 0
	for _, d := range directions {
		x := i + d[0]
		y := j + d[1]

		if x >= 0 && y >= 0 && x < r && y < c {
			if board[x][y] == 1 || board[x][y] == die {
				count++
			}
		}
	}
	return count
}
