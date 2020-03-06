package main

// Space complexity should be not more than O(k)
func getRow(n int) []int {
	if n == 0 {
		return []int{1}
	}
	if n == 1 {
		return []int{1, 1}
	}

	temp := []int{1}
	for i := 2; i < n; i++ {
		j := 0
		for j+1 < len(temp) {

		}
		temp = append(temp, 1)
	}
	return temp
}
