package main

import (
	"fmt"
)

func findPairs(nums []int, k int) (total int) {
	if k < 0 {
		return
	}
	data := make(map[int]int, len(nums))
	for _, num := range nums {
		data[num]++
	}

	fmt.Printf("%v\n", data)

	for key, value := range data {
		if k == 0 {
			if value >= 2 {
				total++
			}
		} else {
			if data[key+k] != 0 {
				fmt.Printf("k: %d, key: %d, key+k: %d, data[key+k]: %d\n", k, key, key+k, data[key+k])

				total++
			}
		}
	}
	return
}
