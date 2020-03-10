package main

// O(2n) solution
func middleNodeO2n(head *ListNode) *ListNode {
	var n int

	curr := head

	for curr != nil {
		n++
		curr = curr.Next
	}

	if n == 1 {
		return head
	}

	curr = head
	for i := 0; i < n/2; i++ {
		curr = curr.Next
	}

	return curr
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast := head.Next
	slow := head

	if fast == nil {
		return head
	}

	count := 0
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		count++
	}

	if count%2 == 0 {
		return slow.Next
	}
	return slow
}
