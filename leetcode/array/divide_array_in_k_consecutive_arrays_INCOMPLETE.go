package main

func isPossibleDivide(arr []int, k int) bool {
	n := len(arr)
	hash := make(map[int]int, 0)

	// fill up hash
	for i := 0; i < n; i++ {
		if _, ok := hash[arr[i]]; ok {
			hash[arr[i]]++
		} else {
			hash[arr[i]] = 1
		}
	}

	for _, num := range arr {
		start := num

		// find lowest number for the start of the streak
		for hash[start-1] > 0 {
			start--
		}

	}

	return true
}
