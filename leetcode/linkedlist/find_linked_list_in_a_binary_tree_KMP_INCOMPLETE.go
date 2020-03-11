package main

import "fmt"

func isSubPathKMP(h *ListNode, root *TreeNode) bool {
	arr := convertLinkedListToSlice(h)
	lps := computeLPS(arr)
	return kmpSearchBinaryTree(root, arr, lps, 0)
}

func kmpSearchBinaryTree(root *TreeNode, arr, lps []int, j int) bool {
	if j == len(arr) {
		return true
	}

	if root == nil {
		return false
	}

	for j > 0 && root.Val != arr[j] {
		j = lps[j-1]
	}
	if root.Val == arr[j] {
		j++
	}
	return kmpSearchBinaryTree(root.Left, arr, lps, j) || kmpSearchBinaryTree(root.Right, arr, lps, j)
}

func kmpSearchString(h *ListNode, s []int) {
	arr := convertLinkedListToSlice(h)
	lps := computeLPS(arr)

	j, i := 0, 0
	for i < len(s) {
		if arr[j] == s[i] {
			i++
			j++
		}

		if j == len(arr) {
			fmt.Printf("pattern: %d", i-j)
		} else if i < len(s) && arr[j] != s[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
}

func computeLPS(arr []int) []int {
	n := len(arr)
	lps := make([]int, n)

	lps[0] = 0
	j, i := 0, 1

	for i < n {
		if arr[i] == arr[j] {
			j++
			lps[i] = j
			i++
		} else {
			if j != 0 {
				j = lps[j-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}

func convertLinkedListToSlice(h *ListNode) []int {
	arr := []int{}

	curr := h
	for curr != nil {
		arr = append(arr, h.Val)
		curr = curr.Next
	}
	return arr
}
