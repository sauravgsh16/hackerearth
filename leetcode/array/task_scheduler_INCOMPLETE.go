package main

import (
	"math"
)

func leastInterval(tasks []byte, k int) int {
	n := len(tasks)
	count := make([]int, 26)

	max := 0
	maxCounter := 0

	for i := 0; i < n; i++ {
		count[tasks[i]-'A']++

		if count[tasks[i]-'A'] == max {
			maxCounter++
		} else if max < count[tasks[i]-'A'] {
			max = count[tasks[i]-'A']
			maxCounter = 1
		}
	}

	partialCount := max - 1
	partialLenght := n - (maxCounter - 1)
	emptySlots := partialCount * partialLenght
	availableTasks := n - max*maxCounter

	idles := math.Max(0, float64(emptySlots-availableTasks))

	return int(idles)
}
