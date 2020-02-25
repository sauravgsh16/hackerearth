package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(arr []int) {
	sort.Ints(arr)
	n := len(arr)
	sum := 0

	for i := 0; i < n; i += 2 {
		sum += arr[i]
	}

	fmt.Printf("%d\n", sum)
}
