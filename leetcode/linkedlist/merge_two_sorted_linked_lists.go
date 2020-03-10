package main

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	curr := &ListNode{}
	head := curr

	for l1 != nil && l2 != nil {

		if l1.Val <= l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}

		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	}

	if l2 != nil {
		curr.Next = l2
	}

	return head.Next
}
