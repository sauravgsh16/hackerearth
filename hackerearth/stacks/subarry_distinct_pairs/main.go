package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var loop int
	fmt.Scanf("%d", &loop)

	reader := bufio.NewReader(os.Stdin)
	txt, _ := reader.ReadString('\n')

	arr := []int{}
	for _, t := range strings.Fields(txt) {
		t := strings.TrimSpace(string(t))
		i, _ := strconv.Atoi(t)
		arr = append(arr, i)
	}

	printPairs(arr)
}

func printPairs(arr []int) {
	n := len(arr)

	var count int

	for i := 0; i < n; i++ {
		left := i - 1
		right := i + 1

		for left >= 0 && arr[left] < arr[i] {
			left--
		}

		for right < n && arr[right] < arr[i] {
			right++
		}

		if left != -1 {
			count++
		}
		if right != n {
			count++
		}
	}

	fmt.Printf("%d\n", count)
}

/*

Given a permutation of number from 1 to N. Among all the subarrays, find the number of unique pairs such that

and a is maximum and b is second maximum in that subarray.

Input:
First line contains an integer, N. Second line contains N space separated distinct integers, , denoting the permutation.


arr = [7 5 6 2 3 4 1]

7
7 5
7 5 6
7 5 6 2
7 5 6 2 3
7 5 6 2 3 4
7 5 6 2 3 4 1

5
5 6
5 6 2
5 6 2 3
5 6 2 3 4
5 6 2 3 4 1

6
6 2
6 2 3
6 2 3 4
6 2 3 4 1

2
2 3
2 3 4
2 3 4 1

3
3 4
3 4 1

4
4 1

1

*/
