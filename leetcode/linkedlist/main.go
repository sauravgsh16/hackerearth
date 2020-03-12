package main

import "fmt"

func main() {
	/*
		l7 := &ListNode{
			Val: 1,
			//Next: l5,
		}
		l6 := &ListNode{
			Val:  5,
			Next: l7,
		}
		l5 := &ListNode{
			Val:  2,
			Next: l6,
		}
	*/
	l4 := &ListNode{
		Val: 5,
	}

	l3 := &ListNode{
		Val:  4,
		Next: l4,
	}

	l12 := &ListNode{
		Val:  3,
		Next: l3,
	}
	l11 := &ListNode{
		Val:  2,
		Next: l12,
	}

	h1 := &ListNode{
		Val:  1,
		Next: l11,
	}

	l4.Next = l11

	fmt.Printf("%v\n", detectCycle(h1))
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
