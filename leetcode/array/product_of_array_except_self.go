package main

func productExceptSelf(arr []int) []int {
	n := len(arr)
	t1 := make([]int, n)
	t2 := make([]int, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			t1[i] = 1
			continue
		}
		t1[i] = t1[i-1] * arr[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			t2[i] = 1
			continue
		}
		t2[i] = t2[i+1] * arr[i+1]
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = t1[i] * t2[i]
	}
	return res
}
