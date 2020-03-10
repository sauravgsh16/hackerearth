package main

func longestOness(arr []int, k int) int {
	var maxLength, numZeros int
	n := len(arr)
	i, j := 0, 0

	for i < n {
		if arr[i] == 0 {
			numZeros++
		}
		for numZeros > k {
			if arr[j] == 0 {
				numZeros--
			}
			j++
		}
		if maxLength < (i - j + 1) {
			maxLength = i - j + 1
		}
		i++
	}
	return maxLength
}

func longestOnes(arr []int, k int) int {
	var maxLength, count int
	n := len(arr)
	i, j := 0, 0

	for j < n {
		if arr[j] == 1 || k > 0 {
			if arr[j] == 0 {
				k--
			}
			j++
			count++

		} else {
			if maxLength < count {
				maxLength = count
			}

			for k <= 0 {
				if arr[i] == 0 {
					k++
				}
				i++
			}

			count = j - i
		}
	}

	if maxLength < count {
		maxLength = count
	}

	return maxLength
}
