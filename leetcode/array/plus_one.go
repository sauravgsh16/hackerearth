package main

func plusOne(arr []int) []int {
	n := len(arr)
	inc := 1

	for i := n - 1; i >= 0; i-- {
		arr[i] += inc
		if arr[i] >= 10 {
			arr[i] -= 10
			continue
		}
		inc = 0
		break
	}

	var res []int

	if inc == 1 {
		res = []int{1}
		for _, i := range arr {
			res = append(res, i)
		}
	} else {
		res = arr
	}
	return res
}
