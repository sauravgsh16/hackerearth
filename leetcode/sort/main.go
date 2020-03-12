package main

import "fmt"

func main() {
	/*
			root1 := &TreeNode{
				Val: 2,
			}

			l1 := &TreeNode{
				Val: 1,
			}

			r1 := &TreeNode{
				Val: 4,
			}

			root2 := &TreeNode{
				Val: 1,
			}

			l2 := &TreeNode{
				Val: 0,
			}

			r2 := &TreeNode{
				Val: 3,
			}

			root1.Left = l1
			root1.Right = r1

			root2.Left = l2
			root2.Right = r2

		arr := [][]int{
			[]int{3, 3},
			[]int{5, -1},
			[]int{-2, 4},
		}

		fmt.Printf("%v\n", kClosest(arr, 2))
	*/
	a := []int{4, 2, 1, 3}

	fmt.Printf("%v\n", insertionSort(a))
}

// TreeNode of a tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
