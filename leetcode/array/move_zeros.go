package main

func moveZeroes(arr []int) {
	n := len(arr)

	j := 0
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			continue
		}
		arr[j] = arr[i]
		j++
	}

	for j < n {
		arr[j] = 0
		j++
	}
}
