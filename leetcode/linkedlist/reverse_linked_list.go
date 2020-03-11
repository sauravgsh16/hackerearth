package main

func reverseLinkedList(h *ListNode) *ListNode {
	var prev *ListNode
	curr := h

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	h = prev

	return h
}

func reverseLinkedListRecursive(h *ListNode) *ListNode {
	if h == nil || h.Next == nil {
		return h
	}

	rest := reverseLinkedList(h.Next)
	h.Next.Next = h

	h.Next = nil

	return rest
}
