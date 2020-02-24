package main

import "fmt"

func pascalsTriangle(n int) {
	res := [][]int{}
	for i := 0; i <= n; i++ {
		if i == 0 {
			res = append(res, []int{1})
			continue
		}

		temp := []int{1}
		j := 0
		for j+1 < len(res[len(res)-1]) {
			val := res[len(res)-1][j] + res[len(res)-1][j+1]
			temp = append(temp, val)
			j++
		}
		temp = append(temp, 1)
		res = append(res, temp)
	}
	fmt.Printf("%v\n", res[n])
}
