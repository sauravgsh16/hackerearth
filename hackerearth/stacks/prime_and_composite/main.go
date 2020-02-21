package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')

	txt, _ := reader.ReadString('\n')

	arr := []int{}
	for _, f := range strings.Fields(txt) {
		i, _ := strconv.Atoi(string(f))
		arr = append(arr, i)
	}

	formArray(arr)
}

const (
	max = 1000001
)

func primeSeive() []bool {
	sieve := make([]bool, max)
	for i := 0; i < max; i++ {
		sieve[i] = true
	}

	i := 2

	for i*i <= max {
		if sieve[i] == true {
			for j := i * i; j < max; j += i {
				sieve[j] = false
			}
		}
		i++
	}
	return sieve
}

func formArray(arr []int) {
	prime := primeSeive()

	stackPrime := []int{}
	stackcomp := []int{}

	for i := 0; i < len(arr); i++ {
		if prime[arr[i]] == true {
			stackPrime = append(stackPrime, arr[i])
		} else {
			stackcomp = append(stackcomp, arr[i])
		}
	}

	for _, p := range stackPrime {
		fmt.Printf("%d ", p)
	}

	fmt.Println()

	for i := len(stackcomp) - 1; i >= 0; i-- {
		fmt.Printf("%d ", stackcomp[i])
	}

	fmt.Println()
}
