package main

func pascalTriangle(n int) [][]int {
	res := [][]int{}

	for i := 0; i < n; i++ {
		if i == 0 {
			res = append(res, []int{1})
			continue
		}

		j := 0
		temp := []int{1}
		for (j + 1) < len(res[(len(res)-1)]) {
			val := res[len(res)-1][j] + res[len(res)-1][j+1]
			temp = append(temp, val)
			j++
		}
		temp = append(temp, 1)
		res = append(res, temp)
	}
	return res
}
