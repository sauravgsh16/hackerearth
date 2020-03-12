package main

import (
	"sort"
)

type dist struct {
	point    []int
	position int
}

func kClosest(points [][]int, k int) [][]int {
	distances := make([]*dist, 0)

	for _, p := range points {
		pos := p[0]*p[0] + p[1]*p[1]
		distances = append(distances, &dist{p, pos})
	}

	sort.Slice(distances[:], func(i, j int) bool {
		return distances[i].position < distances[j].position
	})

	res := make([][]int, k)

	for i := 0; i < k; i++ {
		res[i] = distances[i].point
	}

	return res
}
