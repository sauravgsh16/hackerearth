package main

import "fmt"

func main() {
	var head *ListNode

	head = push(head, 5)
	head = push(head, 20)
	head = push(head, 4)
	head = push(head, 3)
	head = push(head, 30)

	head = insertionSort(head)

	for head != nil {
		fmt.Printf("%d ", head.val)
	}
}

// ListNode of a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode for binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Tree binary
type Tree struct {
	root *TreeNode
}

func (t *Tree) printTree() {
	inorder(t.root)
}

func inorder(node *TreeNode) {
	if node == nil {
		return
	}

	inorder(node.Left)
	fmt.Printf("%d ", node.Val)
	inorder(node.Right)
}

func createTree() *Tree {
	t := &Tree{
		root: &TreeNode{
			Val: 1,
		},
	}

	t.root.Left = &TreeNode{
		Val: 4,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
		},
	}

	t.root.Right = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
	}

	return t
}
