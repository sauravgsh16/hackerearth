package main

type Node struct {
	val  int
	next *Node
}

func insertionSort(head *Node) *Node {
	var result *Node

	curr := head

	for curr != nil {
		next := curr.next

		result = sortedInsert(result, curr)

		curr = next
	}
	head = result
	return head

}

func sortedInsert(head, newNode *Node) *Node {
	var curr *Node

	if head == nil || head.val >= newNode.val {
		newNode.next = head
		head = newNode
	} else {
		curr = head

		for curr.next != nil && curr.next.val < newNode.val {
			curr = curr.next
		}

		newNode.next = curr.next
		curr.next = newNode
	}
	return head
}

func push(head *Node, val int) *Node {
	newNode := &Node{
		val: val,
	}

	newNode.next = head
	head = newNode

	return head
}
