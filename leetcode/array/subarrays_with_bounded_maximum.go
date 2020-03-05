package main

// Approach:
// We can find number of subsets for elements whose values <= R
// Then we find the number of subsets for element whose values are < L
// Answer will be the (no. of <=R subsets) - (no. of < than L subsets)

func numSubarrayBoundedMax(arr []int, L int, R int) int {
	return findSubset(arr, R+1) - findSubset(arr, L)
}

func findSubset(arr []int, val int) int {
	count := 0
	temp := 0

	for i := 0; i < len(arr); i++ {
		// We count elements which are less than the value
		// Once it does not satisfy the condition, we count no of subsets,
		// and then set `temp` to 0 again
		if arr[i] < val {
			temp++
		} else {
			count += temp * (temp + 1) / 2
			temp = 0
		}
	}
	count += temp * (temp + 1) / 2
	return count
}

func numSubsetSlidingWindowSoln(arr []int, L, R int) int {
	n := len(arr)
	i, j := 0, 0
	count := 0
	temp := 0

	for j < n {
		if arr[j] >= L && arr[j] <= R {
			temp = j - i + 1
		} else if arr[j] > R {
			temp = 0
			i = j + 1
		}

		count += temp
		j++
	}
	return count
}
