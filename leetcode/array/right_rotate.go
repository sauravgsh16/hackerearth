package main

import (
	"fmt"
	"math"
)

func reverse(a []int) {
	i, j := 0, len(a)-1

	var temp int
	for i < j {
		temp = a[i]
		a[i] = a[j]
		a[j] = temp
		i++
		j--
	}
}

func leftRotate(arr []int, k int) {
	n := len(arr)

	if k > n {
		k = n % k
	}

	reverse(arr[:k])
	reverse(arr[k:])
	reverse(arr)
}

func rightRotate(arr []int, k int) {
	n := len(arr)

	if k > n {
		k = k % n
	}

	k = int(math.Abs(float64(n - k)))

	if k == 0 {
		return
	}

	leftRotate(arr, k)
	fmt.Printf("%v\n", arr)
}
