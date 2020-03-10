package main

func getIntersectionNode(head1, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		return nil
	}

	// Two pointers.
	pA := head1
	pB := head2

	// Move two pointer, whenever any pointer hits end of the list
	// make the pointer to be the head of the other list

	// Run loop such till pa == pb
	// there are two ways in which we exit the loop
	// 1) when we find match
	// 2) when both pointers hit the end of the list
	for pA != pB {
		if pA == nil {
			pA = head2
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = head1
		} else {
			pB = pB.Next
		}
	}

	return pA
}

func getIntersectionNodeC(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
