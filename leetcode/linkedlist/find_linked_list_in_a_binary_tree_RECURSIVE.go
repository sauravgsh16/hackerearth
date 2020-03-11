package main

// TreeNode for binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Tree binary
type Tree struct {
	root *TreeNode
}

func createTree() *Tree {
	t := &Tree{
		root: &TreeNode{
			Val: 1,
		},
	}

	t.root.Left = &TreeNode{
		Val: 4,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
		},
	}

	t.root.Right = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
	}

	return t
}

func dfs(h *ListNode, root *TreeNode) bool {
	if h == nil {
		return true
	}

	if root == nil {
		return false
	}

	return h.Val == root.Val && (dfs(h.Next, root.Left) || dfs(h.Next, root.Right))
}

func isSubPath(h *ListNode, root *TreeNode) bool {
	if h == nil {
		return true
	}

	if root == nil {
		return false
	}

	return dfs(h, root) || isSubPath(h, root.Left) || isSubPath(h, root.Right)
}
