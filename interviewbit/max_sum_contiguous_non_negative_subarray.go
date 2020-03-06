package main

func maxSubArrayNonNegative(arr []int) []int {
	n := len(arr)
	maxSoFar := -999999
	maxEndingHere := 0
	start, end := 0, 0

	res := make([]int, 3)

	for i := 0; i < n; i++ {
		if arr[i] < 0 {

			if maxEndingHere > maxSoFar {
				maxSoFar = maxEndingHere
				res[0] = maxSoFar
				res[1] = start
				res[2] = end
			}
			start = i + 1
			maxEndingHere = 0
			continue
		} else {
			maxEndingHere = arr[i] + maxEndingHere
			end = i + 1
		}
	}

	if maxEndingHere > res[0] {
		res[0], res[1], res[2] = maxEndingHere, start, end
	} else if maxEndingHere == res[0] {
		if end-start > res[2]-res[1] {
			res[0], res[1], res[2] = maxEndingHere, start, end
		} else if res[2]-res[1] == end-start && start < res[1] {
			res[0], res[1], res[2] = maxEndingHere, start, end
		}
	}

	if res[1] > n {
		return nil
	}

	return arr[res[1]:res[2]]
}
