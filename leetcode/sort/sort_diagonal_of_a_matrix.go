package main

import (
	"sort"
)

func diagonalSort(arr [][]int) [][]int {
	n := len(arr)
	m := len(arr[0])

	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		row := n - 1
		col := i
		temp := []int{}
		for col >= 0 && row >= 0 {
			temp = append(temp, arr[row][col])
			row--
			col--
		}

		sort.Ints(temp)
		row = n - 1
		col = i
		j := len(temp) - 1
		for col >= 0 && j >= 0 {
			res[row][col] = temp[j]
			j--
			row--
			col--
		}
	}

	for i := n - 1; i >= 0; i-- {
		row := i
		col := m - 1

		temp := []int{}
		for row >= 0 && col >= 0 {
			temp = append(temp, arr[row][col])
			col--
			row--
		}

		sort.Ints(temp)
		row = i
		col = m - 1
		j := len(temp) - 1
		for row >= 0 && j >= 0 {
			res[row][col] = temp[j]
			j--
			col--
			row--
		}
	}

	return res
}
