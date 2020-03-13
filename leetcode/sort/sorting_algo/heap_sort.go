package main

/*
Heap sort is an in-place algorithm.
Its typical implementation is not stable, but can be made stable (See this)

Time Complexity: Time complexity of heapify is O(Logn).
Time complexity of createAndBuildHeap() is O(n) and overall time complexity of Heap Sort is O(nLogn)

*/

func isleaf(pos, n int) bool {
	if pos > ((n-1)/2) && pos <= n-1 {
		return true
	}
	return false
}

func leftChild(pos, n int) int {
	left := (2 * pos) + 1
	if left <= n-1 {
		return left
	}
	return -1
}

func rightChild(pos, n int) int {
	right := (2 * pos) + 2
	if right <= n-1 {
		return right
	}
	return -1
}

func heapify(arr []int, pos int) {
	n := len(arr)
	if isleaf(pos, n) {
		return
	}

	left := leftChild(pos, n)
	right := rightChild(pos, n)

	if left != -1 && right != -1 {
		if arr[pos] > arr[left] || arr[pos] > arr[right] {
			if arr[left] < arr[right] {
				arr[left], arr[pos] = arr[pos], arr[left]
				heapify(arr, left)
			} else {
				arr[right], arr[pos] = arr[pos], arr[right]
				heapify(arr, right)
			}
		}
	} else if left != -1 {
		if arr[pos] > arr[left] {
			arr[left], arr[pos] = arr[pos], arr[left]
			heapify(arr, left)
		}
	}

}

func parent(idx int) int {
	parent := (idx - 1) / 2
	if parent < 0 {
		return 0
	}
	return parent
}

func insert(arr []int, val int) []int {
	arr = append(arr, val)
	if len(arr) == 1 {
		return arr
	}

	cur := len(arr) - 1

	p := parent(cur)

	for arr[cur] < arr[p] {
		arr[cur], arr[p] = arr[p], arr[cur]
		cur = parent(cur)
		p = parent(cur)
	}

	return arr
}

func extractMin(arr []int) ([]int, int) {
	arr[0], arr[len(arr)-1] = arr[len(arr)-1], arr[0]

	min := arr[len(arr)-1]
	arr = arr[:len(arr)-1]

	heapify(arr, 0)
	return arr, min
}

func heapSort(arr []int) []int {
	var res []int
	for _, val := range arr {
		res = insert(res, val)
	}

	var val int
	for i := 0; i < len(arr); i++ {
		res, val = extractMin(res)
		arr[i] = val
	}
	return arr
}
