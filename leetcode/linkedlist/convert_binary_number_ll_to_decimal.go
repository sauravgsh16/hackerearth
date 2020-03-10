package main

func getDecimalValue(head *ListNode) int {
	if head == nil {
		return 0
	}

	val := 0
	for head != nil {
		val = val << 1
		val += head.Val
		head = head.Next
	}

	return val
}
