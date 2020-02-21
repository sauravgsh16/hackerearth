package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		a []int
		b []int
	)

	reader := bufio.NewReader(os.Stdin)
	txt, _ := reader.ReadString('\n')
	for _, f := range strings.Fields(txt) {
		i, _ := strconv.Atoi(string(f))
		a = append(a, i)
	}

	txt, _ = reader.ReadString('\n')
	for _, f := range strings.Fields(txt) {
		i, _ := strconv.Atoi(string(f))
		b = append(b, i)
	}

	findMax(b, a[1])

}

func findMax(arr []int, k int) {
	sum := 0
	n := len(arr)

	for i := 0; i < k; i++ {
		sum += arr[i]
	}

	maxSum := sum

	for i := 0; i < k-1; i++ {
		sum = sum + arr[n-1-i] - arr[k-1-i]
		if sum >= maxSum {
			maxSum = sum
		}
	}
	fmt.Printf("%d\n", maxSum)
}

/*

124 50
26 26 37 38 10 69 25 48 80 43 54 71 40 96 5 95 82 74 12 35 32 68 94 11 80 58 12 8 30 74 72 10 86 71 42 17 84 9 46 67 21 5 28 49 91 10 84 5 70 69 66 37 22 61 39 59 65 20 79 56 89 59 23 62 39 85 27 90 59 48 20 96 64 100 58 99 25 44 13 93 77 90 87 99 80 35 23 37 63 37 42 82 14 95 42 55 74 75 60 86 100 90 18 33 37 56 55 97 81 58 99 75 44 14 54 10 81 95 46 31 40 42 68 65

*/
