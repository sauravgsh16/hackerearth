package main

func getAllElements(r1, r2 *TreeNode) []int {
	res := []int{}

	arr1 := []int{}
	arr2 := []int{}
	inorder(r1, &arr1)
	inorder(r2, &arr2)

	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		res = append(res, arr1[i])
		i++
	}

	for j < len(arr2) {
		res = append(res, arr2[j])
		j++
	}

	return res
}

func inorder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	inorder(node.Left, res)
	*res = append(*res, node.Val)
	inorder(node.Right, res)
}
