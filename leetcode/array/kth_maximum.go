package main

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

func kthMaximun(arr []int) int {
	n := len(arr)

	if n == 1 {
		return arr[0]
	}

	if n == 2 {
		if arr[0] > arr[1] {
			return arr[0]
		}
		return arr[1]
	}

	first, second, third := minInt, minInt, minInt
	for i := 0; i < n; i++ {
		if arr[i] > first {
			third = second
			second = first
			first = arr[i]
		} else if arr[i] < first && arr[i] > second {
			third = second
			second = arr[i]
		} else if arr[i] < first && arr[i] < second && arr[i] > third {
			third = arr[i]
		}
	}
	if third == 0 {
		return first
	}
	return third
}
