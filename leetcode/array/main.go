package main

import "fmt"

func main() {
	//a := []int{2, 2, 2, 3}
	words := []string{"a", "bb", "acda", "acae"}
	fmt.Printf("%v\n", numMatchingSubseq("abcfdae", words))

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
