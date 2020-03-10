package main

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	rest := reverseListRecursive(head.Next)
	head.Next.Next = head

	head.Next = nil

	return rest
}

func reverseListIterative(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	head = prev
	return head
}
