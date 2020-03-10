package main

func deleteNode(node *ListNode) {
	*node = *node.Next

	// OR
	/*
		node.Val = node.Next.Val
		node.Next = node.Next.Next
	*/
}
