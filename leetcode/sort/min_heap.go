package main

import "fmt"

type minHeap struct {
	heap []int
	size int
}

func (m *minHeap) parent(idx int) int {
	parent := (idx - 1) / 2
	if parent < 0 {
		return 0
	}
	return parent
}

func (m *minHeap) swap(idx1, idx2 int) {
	temp := m.heap[idx1]
	m.heap[idx1] = m.heap[idx2]
	m.heap[idx2] = temp
}

func (m *minHeap) insert(val int) {
	m.heap = append(m.heap, val)
	m.size++

	if m.size == 1 {
		return
	}

	cur := m.size - 1

	for m.heap[cur] < m.heap[m.parent(cur)] {
		m.swap(cur, m.parent(cur))
		cur = m.parent(cur)
	}
}

func (m *minHeap) isleaf(pos int) bool {
	if pos > ((m.size-1)/2) && pos <= (m.size-1) {
		return true
	}
	return false
}

func (m *minHeap) leftChild(pos int) int {
	left := (2 * pos) + 1
	if left <= m.size-1 {
		return left
	}
	return -1
}

func (m *minHeap) rightChild(pos int) int {
	right := (2 * pos) + 2
	if right <= m.size-1 {
		return right
	}
	return -1
}

func (m *minHeap) heapify(pos int) {
	if m.isleaf(pos) {
		return
	}

	left := m.leftChild(pos)
	right := m.rightChild(pos)

	if left != -1 && right != -1 {
		if m.heap[pos] > m.heap[left] || m.heap[pos] > m.heap[right] {
			if m.heap[left] < m.heap[right] || m.heap[right] >= 0 {
				m.swap(pos, left)
				m.heapify(left)
			} else {
				m.swap(pos, right)
				m.heapify(right)
			}
		}
	} else if left != -1 {
		if m.heap[pos] > m.heap[left] {
			m.swap(pos, left)
			m.heapify(left)
		}
	}
}

func (m *minHeap) extractMin() int {
	m.swap(0, m.size-1)
	val := m.heap[len(m.heap)-1]
	m.heap = m.heap[:len(m.heap)-1]
	m.size--
	m.heapify(0)
	return val
}

func run() {
	h := &minHeap{
		heap: make([]int, 0),
	}
	h.insert(5)
	h.insert(10)
	h.insert(12)
	h.insert(7)
	h.insert(3)
	h.insert(15)

	fmt.Printf("%v\n", h.heap)

	for i := 0; i < 6; i++ {
		fmt.Printf("%d\n", h.extractMin())
	}
}
