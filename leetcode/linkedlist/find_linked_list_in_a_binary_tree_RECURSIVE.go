package main

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
