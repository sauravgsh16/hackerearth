package main

import "sort"

func rankTeams(votes []string) string {
	hmap := make(map[rune]map[int]int)
	n := len(votes[0])
	for _, vote := range votes {
		for j, v := range vote {
			if hmap[v] == nil {
				hmap[v] = make(map[int]int)
			}
			hmap[v][j]++
		}
	}

	resArr := []rune(votes[0])
	sort.Slice(resArr, func(i, j int) bool {
		for k := 0; k < n; k++ {
			if hmap[resArr[i]][k] != hmap[resArr[j]][k] {
				return hmap[resArr[i]][k] > hmap[resArr[j]][k]
			}
		}
		return resArr[i] < resArr[j]
	})

	return string(resArr)
}
