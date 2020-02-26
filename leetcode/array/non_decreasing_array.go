package main

func checkPossibility(arr []int) bool {
	n := len(arr)
	var changed bool

	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			if changed == true {
				return false
			}
			changed = true

			if i-2 >= 0 && arr[i] < arr[i-2] {
				arr[i] = arr[i-1]
			}
		}
	}
	return true

}
