package main

func main() {
	/*

		fmt.Printf("%v\n", removeSubFolders(a))
	*/
	a := []int{1, 3, 2, 8, 4, 9}
	//fmt.Printf("%d\n", )
	maxProfitWithFee(a, 2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
