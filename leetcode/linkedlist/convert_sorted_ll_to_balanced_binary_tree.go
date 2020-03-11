package main

func sortedListToBST(h *ListNode) *TreeNode {
	if h == nil {
		return nil
	}

	return toBST(h, nil)
}

func toBST(head, tail *ListNode) *TreeNode {
	if head == tail {
		return nil
	}

	slow := head
	fast := head

	// Move till middle of the list
	for fast != tail && fast.Next != tail {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// tree node
	tNode := &TreeNode{
		Val: slow.Val,
	}

	tNode.Left = toBST(head, slow)
	tNode.Right = toBST(slow.Next, tail)
	return tNode
}
