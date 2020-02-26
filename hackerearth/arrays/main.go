package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	txt, _ := reader.ReadString('\n')
	test, _ := strconv.Atoi(strings.TrimSpace(txt))

	for j := 0; j < test; j++ {
		arr := make([]int, 0)
		txt, _ := reader.ReadString('\n')
		n, _ := strconv.Atoi(strings.TrimSpace(txt))

		t, _, _ := reader.ReadLine()

		for _, s := range strings.Fields(string(t)) {
			i, _ := strconv.Atoi(s)

			arr = append(arr, i)
		}

		min1, min2 := math.MaxInt32, math.MaxInt32
		max1, max2 := -1, -1
		for i := 0; i < n; i++ {
			if arr[i]+i > max1 {
				max1 = arr[i] + i
			}
			if arr[i]+1 < min1 {
				min1 = arr[i] + i
			}

			if arr[i]-1 > max2 {
				max2 = arr[i] - i
			}

			if arr[i]-1 < min2 {
				min2 = arr[i] - i
			}
		}

		fmt.Printf("%d", getMax(int(max1-min1), int(max2-min2)))
	}
}

func getMax(i, j int) int {
	if i > j {
		return i
	}
	return j
}

/*
func main() {
	// a := []int{5, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8}
	// key := 8
	//a := []int{5, 7, 7, 8, 8, 10}
	a := []int{2, 3, 6, 7}
	key := 7
	fmt.Printf("%v\n", combinationSum(a, key))
}
*/

func searchInsert(arr []int, target int) int {
	n := len(arr)

	if target > arr[n-1] {
		return n
	}
	if target < arr[0] {
		return 0
	}

	return binarySearch(arr, 0, len(arr)-1, target)
}

func binarySearch(arr []int, low, high, key int) int {

	for low < high {

		var mid int

		mid = low + ((high - low + 1) / 2)

		if arr[mid] == key {
			return mid
		}

		if arr[mid] < key {
			if mid+1 <= len(arr)-1 && arr[mid+1] >= key {
				return mid + 1
			}
			low = mid + 1
		} else {

			if mid-1 >= 0 && arr[mid-1] < key {
				return mid
			}
			high = mid - 1
		}
	}
	return 0
}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	result = dfs(result, candidates, target, 0, []int{})
	return result

}

func dfs(result [][]int, candidates []int, target, index int, path []int) [][]int {
	if target < 0 {
		return nil
	}
	if target == 0 {
		result = append(result, path)
	}

	for i := index; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		path = append(path, candidates[i])
		result = append(result, dfs(result, candidates, target-candidates[i], i, path)...)
	}
	return result
}
