package main

func findDuplicates(arr []int) int {
	hash := make(map[int]int)
	n := len(arr)

	for i := 0; i < n; i++ {
		if hash[arr[i]] > 0 {
			return arr[i]
		}
		hash[arr[i]]++
	}
	return -1
}
