package main

func flip(a string) []int {
	n := len(a)
	maxDiff := 0
	diff := 0
	start := 0

	result := make([]int, 2)

	for i := 0; i < n; i++ {
		ch := string(a[i])

		if ch == "0" {
			diff++
		} else {
			diff--
		}

		if diff < 0 {
			diff = 0
			start = i + 1
			continue
		}

		if maxDiff < diff {
			maxDiff = diff
			result[0] = start + 1
			result[1] = i + 1
		}
	}
	if result[0] == 0 && result[1] == 0 {
		return []int{}
	}
	return result
}
