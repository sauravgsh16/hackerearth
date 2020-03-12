package main

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return findCycleStart(head, slow)
		}
	}
	return nil
}

func findCycleStart(head, loopNode *ListNode) *ListNode {
	ptr1 := head
	var ptr2 *ListNode

	for {
		ptr2 = loopNode

		// We start from the loopNode and check if we reach ptr1
		for ptr2.Next != loopNode && ptr2.Next != ptr1 {
			ptr2 = ptr2.Next
		}

		// We check if the previous loop broke because of
		// ptr2.Next == ptr1
		if ptr2.Next == ptr1 {
			break
		}

		// We increment ptr1
		ptr1 = ptr1.Next
	}

	return ptr2.Next
}
