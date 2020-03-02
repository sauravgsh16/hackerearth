package main

func combinationSum(arr []int, target int) [][]int {
	result := make([][]int, 0)
	// sort.Ints(arr)

	temp := []int{}
	backTrack(arr, temp, &result, target, 0)
	return result
}

func backTrack(arr []int, temp []int, result *[][]int, remain, start int) {
	if remain < 0 {
		return
	}
	if remain == 0 {
		r := append([]int{}, temp...)
		*result = append(*result, r)
		return
	}

	for i := start; i < len(arr); i++ {
		t := append(temp, arr[i])
		backTrack(arr, t, result, remain-arr[i], i)
	}
}
