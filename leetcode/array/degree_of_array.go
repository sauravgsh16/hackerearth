package main

import "fmt"

func findShortestSubArray(arr []int) int {
	n := len(arr)
	// map storing {element: [count, firstindex, distance]}
	hash := make(map[int][]int, 0)

	for i := 0; i < n; i++ {
		ele := arr[i]
		if _, ok := hash[ele]; ok {
			hash[ele][0]++
			hash[ele][2] = i - hash[ele][1]
		} else {
			hash[ele] = []int{1, i, 0}
		}
	}

	max := 0
	dist := 0

	fmt.Printf("%v\n", hash)

	for _, v := range hash {
		if v[0] > max {
			max = v[0]
			dist = v[2]
		} else if v[0] == max && v[2] < dist {
			dist = v[2]
		}
	}

	return dist + 1
}
