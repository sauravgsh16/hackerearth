package main

func maxSumTwoNoOverlap(arr []int, l, m int) int {
	return max(maxTwoNoOverlap(arr, l, m), maxTwoNoOverlap(arr, m, l))
}

func maxTwoNoOverlap(arr []int, l, m int) int {
	res := 0
	n := len(arr)
	left := make([]int, n+1)
	right := make([]int, n+1)
	sl, sr := 0, 0

	j := n - 1
	for i := 0; i < n; i++ {
		sl += arr[i]
		sr += arr[j]

		left[i+1] = max(sl, left[i])
		right[j] = max(sr, right[j+1])

		if i+1 >= l {
			sl -= arr[i+1-l]
		}
		if i+1 >= m {
			sr -= arr[j+m-1]
		}
		j--
	}

	for i := 0; i < n-1; i++ {
		res = max(res, left[i]+right[i])
	}

	return res
}
