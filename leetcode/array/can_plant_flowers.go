package main

func canPlaceFlowers(arr []int, k int) bool {
	if k == 0 {
		return true
	}

	n := len(arr)
	var count int

	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			var prev, next int
			if i == 0 {
				prev = 0
			} else {
				prev = arr[i-1]
			}
			if i == n-1 {
				next = 0
			} else {
				next = arr[i+1]
			}
			if prev == 0 && next == 0 {
				arr[i] = 1
				count++
				if count == k {
					return true
				}
			}
		}
	}
	return count == k

}
